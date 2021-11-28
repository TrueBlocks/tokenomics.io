/*-------------------------------------------------------------------------------------------
 * qblocks - fast, easily-accessible, fully-decentralized data from blockchains
 * copyright (c) 2018, 2019 TrueBlocks, LLC (http://trueblocks.io)
 *
 * This program is free software: you may redistribute it and/or modify it under the terms
 * of the GNU General Public License as published by the Free Software Foundation, either
 * version 3 of the License, or (at your option) any later version. This program is
 * distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even
 * the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
 * General Public License for more details. You should have received a copy of the GNU General
 * Public License along with this program. If not, see http://www.gnu.org/licenses/.
 *-------------------------------------------------------------------------------------------*/
#include "options.h"
#include "classes/logentry_min.h"

bool visitFile_json_2_csv(const string_q& path, void* data);
//--------------------------------------------------------------------
bool COptions::handle_json_2_csv(void) {
    expContext().exportFmt = CSV1;
    configureDisplay("getLogs", "CLogEntry_min", STR_DISPLAY_LOGENTRY_MIN);
    forEveryFileInFolder("./data/", visitFile_json_2_csv, nullptr);
    return false;
}

void cleanLogs(string_q& contents);
//--------------------------------------------------------------------
bool visitFile_json_2_csv(const string_q& path, void* data) {
    if (endsWith(path, '/')) {
        return forEveryFileInFolder(path + "*", visitFile_json_2_csv, data);
    } else {
        cerr << "Visiting " << path << endl;
        if (endsWith(path, ".json")) {
            string_q contents = asciiFileToString(path);
            cleanLogs(contents);

            CLogEntry_min log;

            ostringstream os;
            bool first = true;
            while (log.parseJson3(contents)) {
                if (first)
                    os << exportPreamble(expContext().fmtMap["header"], GETRUNTIME_CLASS(CLogEntry_min));
                first = false;
                os << log.Format("\"" + substitute(STR_DISPLAY_LOGENTRY_MIN, "\t", "\",\"") + "\"") << endl;
                log = CLogEntry_min();
            }
            if (first)
                os << exportPreamble(expContext().fmtMap["header"], GETRUNTIME_CLASS(CLogEntry_min));
            stringToAsciiFile(substitute(path, ".json", ".csv"), os.str());
        }
    }
    return true;
}

//--------------------------------------------------------------------
void cleanLogs(string_q& contents) {
    replace(contents, "{ \"data\": [", "");
    replaceReverse(contents, "] }", "");
    CStringArray lines;
    explode(lines, contents, '\n');
    ostringstream os;
    for (auto line : lines) {
        if (!contains(line, "compressedLog"))
            os << line << endl;
    }
    contents = os.str();
}
