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

address_t only;
bool fixHeader(const string_q &path, void *data);
//----------------------------------------------------------------
int main(int argc, const char *argv[])
{
    etherlib_init(quickQuitHandler);
    if (argc > 1)
    {
        only = argv[1];
    }
    forEveryFileInFolder("./", fixHeader, nullptr);
    etherlib_cleanup();
    return 1;
}

//----------------------------------------------------------------
bool fixHeader(const string_q &path, void *data)
{
    if (endsWith(path, "/"))
    {
        return forEveryFileInFolder(path + "*", fixHeader, data);
    }

    if (!endsWith(path, ".csv")) {
        return true;
    }

    if ((only.empty() && !contains(path, "combined")) || (contains(path, only))) {
        cerr << "Fixing " << path << "                                                \r";
        cerr.flush();
        CStringArray lines;
        asciiFileToLines(path, lines);
        if (lines.size() > 0) {
            replace(lines[0], "\"blocknumber\"", "\"blockNumber\"");
            replace(lines[0], "\"transactionindex\"", "\"transactionIndex\"");
            replace(lines[0], "\"logindex\"", "\"logIndex\"");
            replace(lines[0], "\"compressedlog\"", "\"compressedLog\"");
            replace(lines[0], "\"bn\"", "\"blockNumber\"");
            replace(lines[0], "\"tx\"", "\"transactionIndex\"");
            replace(lines[0], "\"tc\"", "\"traceIndex\"");
            replace(lines[0], "\"addr\"", "\"neighbor\"");
            replace(lines[0], "\"ethergasprice\"", "\"ethGasPrice\"");
            replace(lines[0], "\"gasused\"", "\"gasUsed\"");
            replace(lines[0], "\"iserror\"", "\"isError\"");
            replace(lines[0], "\"compressedtx\"", "\"compressedTx\"");
        }
        ostringstream os;
        for (auto line : lines) {
            os << line << endl;
        }
        stringToAsciiFile(path, os.str());
    }
    return true;
}
