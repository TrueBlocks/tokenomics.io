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

//----------------------------------------------------------------
int main(int argc, const char *argv[])
{
    CStringArray lines;
    asciiFileToLines("./addresses.csv", lines);

    bool quit = false;

    while (!quit)
    {
        for (auto line : lines)
        {
            CStringArray parts;
            explode(parts, line, ',');
            address_t addr = toLower(parts[1]);
            address_t uniSwap = "0x7a250d5630b4cf539739df2c5dacb4c659f2488d"; // someone used UniSwap v2 as thier grant address
            if ((!isAddress(addr) || isZeroAddr(addr)) && addr == uniSwap)
                continue;

            string_q fileName = getCachePath("monitors/" + addr + ".acct.bin");
            uint64_t sizeBefore = fileSize(fileName) / 8;
            ostringstream os;
            os << "chifra list " << addr << " >/dev/null";
            if (system(os.str().c_str()) != 0)
            {
                quit = true;
                break;
            }
            uint64_t sizeAfter = fileSize(fileName) / 8;
            if (sizeBefore == sizeAfter)
            {
                cerr << bBlack << "Skipping " << fileName << bGreen << " (" << sizeBefore << " == " << sizeAfter << ")" << cOff << endl;
            } else {
                cerr << bYellow << "Calling " << fileName << " (" << sizeBefore << " == " << sizeAfter << ")" << cOff << endl;
                ostringstream oss;
                oss << "chifra export --appearances --fmt csv " << addr << " | cut -f2,3 -d',' >apps/" << addr << ".csv ; " << endl;
                // oss << "chifra export --balances --fmt csv " << addr << " >bals/" << addr << ".csv ; " << endl;
                oss << "chifra export --articulate --cache --cache_traces --fmt csv " << addr << " >txs/" << addr << ".csv ; " << endl;
                oss << "chifra export --logs --articulate --relevant "
                       "--emitter 0xdf869fad6db91f437b59f1edefab319493d4c4ce "
                       "--emitter 0xf2354570be2fb420832fb7ff6ff0ae0df80cf2c6 "
                       "--emitter 0x7d655c57f71464b6f83811c55d84009cd9f5221c --fmt csv "
                    << addr << " >logs/" << addr << ".csv ; " << endl;
                oss << "chifra export --neighbors --cache --cache_traces --fmt csv " << addr << " >neighbors/" << addr << ".csv;" << endl;
                oss << "./fixHeaders " << addr << endl;
                if (system(oss.str().c_str()) != 0)
                {
                    //                    quit = true;
                    //                    break;
                }
            }
        }
    }
    return 0;
}
