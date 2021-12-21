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

// clang-format off
const char* STR_CMD_LIST = "chifra export --appearances --fmt csv [{ADDR}] | cut -f2,3 -d',' >apps/[{ADDR}].csv ; ";
const char* STR_CMD_TXS = "chifra export --articulate --cache --cache_traces --fmt csv [{ADDR}] >txs/[{ADDR}].csv ; ";
const char* STR_CMD_LOGS = "./export_logs.1.sh [{ADDR}] ; ";
const char* STR_CMD_NEIGHBORS = "chifra export --neighbors --deep --fmt csv [{ADDR}] >neighbors/[{ADDR}].csv ; ";
const char* STR_CMD_STATEMENTS = "./export_statements.1.sh [{ADDR}] ; ";
// clang-format on

//----------------------------------------------------------------
int main(int argc, const char* argv[]) {
    bool quit = false;
    CMetaData lastChunk = getMetaData();
    while (!quit) {
        CMetaData thisChunk = getMetaData();
        if (lastChunk.finalized == thisChunk.finalized) {
            LOG_INFO("Skipping because no new chunks: ", lastChunk.finalized, " --> ", thisChunk.finalized);
            LOG_INFO("Sleeping for 3 minutes...");
            usleep(1800000000);
            continue;
        }

        CStringArray lines;
        asciiFileToLines("./addresses.csv", lines);

        size_t nChanged = 0, nProcessed = 0;
        timestamp_t start = date_2_Ts(Now());
        for (auto line : lines) {
            CStringArray parts;
            explode(parts, line, ',');
            address_t addr = toLower(parts[1]);

            // Ignore invalid addresses
            if ((!isAddress(addr) || isZeroAddr(addr)))
                continue;

            nProcessed++;

            // Figure out how many records there are...
            string_q monitorFn = getCachePath("monitors/" + addr + ".acct.bin");
            uint64_t nRecordsBefore = fileSize(monitorFn) / 8;
            if (nRecordsBefore > 100000) {
                ostringstream os;
                os << "Skipping massive address: " << addr << " with " << nRecordsBefore << " appearances.";
                stringToAsciiFile("./skipped-too-large.txt", os.str());
                LOG_ERR(bRed, os.str(), cOff);
                continue;
            }

            // Freshen the monitor
            if (system(substitute(STR_CMD_LIST, "[{ADDR}]", addr).c_str()) != 0) {
                quit = true;
                break;
            }

            // Figure out how many records there are after freshen
            uint64_t sizeAfter = fileSize(monitorFn) / 8;

            if (fileExists("./apps/" + addr + ".csv") && nRecordsBefore == sizeAfter) {
                // If there are no new records, we don't have to freshen the rest of the data
                LOG_INFO(bBlack, "Skip ", substitute(monitorFn, getCachePath(""), "./"), bGreen, " (", nRecordsBefore,
                         " == ", sizeAfter, ")", cOff);

            } else {
                nChanged++;

                // There are new records, freshen everything
                LOG_INFO(bYellow, "Call ", substitute(monitorFn, getCachePath(""), "./"), bGreen, " (", nRecordsBefore,
                         " != ", sizeAfter, ")", cOff);

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

        timestamp_t stop = date_2_Ts(Now());
        ostringstream out;
        out << lines.size() << "," << nProcessed << "," << nChanged << "," << start << "," << stop << ","
            << (stop - start) << endl;
        out << "sleeping: " << date_2_Ts(Now()) << endl;
        LOG_INFO(out.str());
        usleep(1800000000);
        out << "starting: " << date_2_Ts(Now()) << endl;
        appendToAsciiFile("./timing.txt", out.str());

        lastChunk = thisChunk;

    }  // while (true)

    return 0;
}
