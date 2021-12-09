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
const char* STR_EMIT = "--emitter 0xdf869fad6db91f437b59f1edefab319493d4c4ce --emitter 0xf2354570be2fb420832fb7ff6ff0ae0df80cf2c6 --emitter 0x7d655c57f71464b6f83811c55d84009cd9f5221c";

const char* STR_CMD1 = "chifra export --appearances --fmt csv [{ADDR}] | cut -f2,3 -d',' >apps/[{ADDR}].csv ; cd apps ; ../fixHeaders [{ADDR}] ; cd - ; ";
const char* STR_CMD3 = "chifra export --articulate --cache --cache_traces  --fmt csv [{ADDR}] >txs/[{ADDR}].csv ; cd txs ; ../fixHeaders [{ADDR}] ; cd - ; ";
const char* STR_CMD2 = "chifra export --balances --fmt csv [{ADDR}] >bals/[{ADDR}].csv ; cd bals ; ../fixHeaders [{ADDR}] ; cd - ; ";
const char* STR_CMD4 = "chifra export --logs --articulate --relevant [{EMITTERS}] --fmt csv [{ADDR}] >logs/[{ADDR}].csv ; cd logs ; ../fixHeaders [{ADDR}] ; cd - ; ";
const char* STR_CMD5 = "chifra export --neighbors --fmt csv [{ADDR}] >neighbors/[{ADDR}].csv ; cd neighbors ; ../fixHeaders [{ADDR}] ; cd - ; ";

const char* STR_CMD6 = "./fixHeaders [{ADDR}] ; ";
const char* STR_CMD7 = "chifra list [{ADDR}] >/dev/null";
// clang-format on

//----------------------------------------------------------------
int main(int argc, const char* argv[]) {
    bool quit = false;
    while (!quit) {
        CStringArray lines;
        asciiFileToLines("./addresses.csv", lines);
        for (auto line : lines) {
            CStringArray parts;
            explode(parts, line, ',');
            address_t addr = toLower(parts[1]);

            // Ignore invalid addresses
            if ((!isAddress(addr) || isZeroAddr(addr)))
                continue;

            // Someone used UniSwap v2 as thier grant address, ignore it
            if (addr == "0x7a250d5630b4cf539739df2c5dacb4c659f2488d")
                continue;

            // Figure out how many records there are...
            string_q monitorFn = getCachePath("monitors/" + addr + ".acct.bin");
            uint64_t nRecordsBefore = fileSize(monitorFn) / 8;

            // Freshen the monitor
            if (system(substitute(STR_CMD1, "[{ADDR}]", addr).c_str()) != 0) {
                quit = true;
                break;
            }

            // Figure out how many records there are after freshen
            uint64_t sizeAfter = fileSize(monitorFn) / 8;

            if (fileExists("./apps/" + addr + ".csv") && nRecordsBefore == sizeAfter) {
                // If there are no new records, we don't have to freshen the rest of the data
                cerr << bBlack << "Skip " << monitorFn;
                cerr << bGreen << " (" << nRecordsBefore << " == " << sizeAfter << ")" << cOff << endl;

            } else {
                // There are new records, freshen everything
                cerr << bYellow << "Call " << monitorFn;
                cerr << " (" << nRecordsBefore << " == " << sizeAfter << ")" << cOff << endl;

                ostringstream oss;
                oss << substitute(STR_CMD1, "[{ADDR}]", addr) << endl;
                // oss << substitute(STR_CMD2, "[{ADDR}]", addr) << endl;
                oss << substitute(STR_CMD3, "[{ADDR}]", addr) << endl;
                oss << substitute(substitute(STR_CMD4, "[{ADDR}]", addr), "[{EMITTERS}]", STR_EMIT) << endl;
                oss << substitute(STR_CMD5, "[{ADDR}]", addr) << endl;
                oss << substitute(STR_CMD6, "[{ADDR}]", addr) << endl;
                int ret = system(oss.str().c_str());
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
    }      // while (true)

    return 0;
}
