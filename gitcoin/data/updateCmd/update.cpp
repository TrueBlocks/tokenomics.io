/*-------------------------------------------------------------------------------------------
 * qblocks - fast, easily-accessible, fully-decentralized data from blockchains
 * copyright (c) 2016, 2021 TrueBlocks, LLC (http://trueblocks.io)
 *
 * This program is free software: you may redistribute it and/or modify it under the terms
 * of the GNU General Public License as published by the Free Software Foundation, either
 * version 3 of the License, or (at your option) any later version. This program is
 * distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even
 * the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
 * General Public License for more details. You should have received a copy of the GNU General
 * Public License along with this program. If not, see http://www.gnu.org/licenses/.
 *-------------------------------------------------------------------------------------------*/
#include "etherlib.h"
#include "update.h"

//-----------------------------------------------------------------------------------
int main(int argc, const char* argv[]) {
    CStringArray lines;
    asciiFileToLines("./addresses.csv", lines);

    blknum_t lastChunk = getMetaData().finalized;
    bool first = true;

    bool quit = false;
    while (!quit) {
        blknum_t thisChunk = getMetaData().finalized;

        // We only need to process if this is the first run or we have a new chunk...
        if (!first && lastChunk == thisChunk) {
            LOG_INFO("No new chunks: ", lastChunk, " --> ", thisChunk, ". Sleeping for 3 minutes...");
            usleep(1800000000);
            continue;
        }
        first = false;

        timestamp_t start = date_2_Ts(Now());
        size_t nChanged = 0, nProcessed = 0;
        for (auto line : lines) {
            CStringArray parts;
            explode(parts, line, ',');
            address_t addr = toLower(parts[1]);

            // Ignore invalid addresses
            if ((!isAddress(addr) || isZeroAddr(addr)))
                continue;

            // Figure out how many records there are...
            string_q monitorFn = getCachePath("monitors/" + addr + ".acct.bin");
            uint64_t nRecordsBefore = fileSize(monitorFn) / 8;

            // If there are too many, report the same and skip...
            if (nRecordsBefore > 100000) {
                ostringstream os;
                os << "Skipping too-large address: " << addr << " with " << nRecordsBefore << " appearances.";
                LOG_ERR(bRed, os.str(), cOff);
                appendToAsciiFile("./skipped-too-large.txt", os.str());
                continue;
            }

            nProcessed++;

            // Freshen the monitor...
            if (system(substitute(STR_CMD_LIST, "[{ADDR}]", addr).c_str()) != 0) {
                quit = true;
                break;
            }

            // Figure out how many records there are after freshen...
            uint64_t sizeAfter = fileSize(monitorFn) / 8;

            // If there is no transactions file (we can delete that file to force a re-calc) or
            // there are new transactions, re-process. Otherwise, skip...
            if (fileExists("./txs/" + addr + ".csv") && nRecordsBefore == sizeAfter) {
                // There are no new records, we don't have to freshen the rest of the data...
                LOG_INFO(bBlack, "Skip ", substitute(monitorFn, getCachePath(""), "./"), bGreen, " (", nRecordsBefore,
                         " == ", sizeAfter, ")", cOff);

            } else {
                // There are new records, re-write everything...
                LOG_INFO(bYellow, "Call ", substitute(monitorFn, getCachePath(""), "./"), bGreen, " (", nRecordsBefore,
                         " != ", sizeAfter, ")", cOff);

                nChanged++;

                ostringstream oss;
                oss << STR_CMD_TXS << endl;
                oss << STR_CMD_LOGS << endl;
                oss << STR_CMD_NEIGHBORS << endl;
                oss << STR_CMD_STATEMENTS << endl;
                int ret = system(substitute(oss.str(), "[{ADDR}]", addr).c_str());
                if (WIFSIGNALED(ret) && (WTERMSIG(ret) == SIGINT || WTERMSIG(ret) == SIGQUIT)) {
                    cerr << "system call interrupted" << endl;
                    break;
                } else {
                    if (ret != 0 && ret != 256) {
                        cerr << "system call returned " << ret << ". Quitting..." << endl;
                        quit = true;
                    }
                }
            }
        }  // for (auto line : lines)

        if (!quit && nChanged) {
            if (system("./combine_update.sh")) {
            }
            if (system("./update_zips.sh")) {
            }
        }
        lastChunk = thisChunk;

        string_q report = reportTiming(lines.size(), nProcessed, nChanged, start);
        LOG_INFO(report, ". Sleeping...");
        usleep(1800000000);
    }  // while (true)

    return 0;
}

//-----------------------------------------------------------------------------------
string_q reportTiming(size_t nLines, size_t nProc, size_t nChange, timestamp_t start) {
    timestamp_t stop = date_2_Ts(Now());
    ostringstream out;
    out << nLines << "," << nProc << "," << nChange << "," << start << "," << stop << "," << (stop - start) << endl;
    appendToAsciiFile("./timing.txt", out.str());
    return out.str();
}

//-----------------------------------------------------------------------------------
const char* STR_CMD_LIST = "chifra export --appearances --fmt csv [{ADDR}] | cut -f2,3 -d',' >apps/[{ADDR}].csv ; ";
const char* STR_CMD_TXS = "chifra export --articulate --cache --cache_traces --fmt csv [{ADDR}] >txs/[{ADDR}].csv ; ";
const char* STR_CMD_LOGS = "./export_logs.1.sh [{ADDR}] ; ";
const char* STR_CMD_NEIGHBORS = "chifra export --neighbors --deep --fmt csv [{ADDR}] >neighbors/[{ADDR}].csv ; ";
const char* STR_CMD_STATEMENTS = "./export_statements.1.sh [{ADDR}] ; ";
