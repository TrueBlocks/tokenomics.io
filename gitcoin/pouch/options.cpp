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
/*
 * Parts of this file were generated with makeClass. Edit only those parts of the code
 * outside of the BEG_CODE/END_CODE sections
 */
#include "options.h"

//---------------------------------------------------------------------------------------------------
static const COption params[] = {
    // BEG_CODE_OPTIONS
    // clang-format off
    COption("freshen", "f", "", OPT_SWITCH, "freshen the database"),
    COption("json2csv", "j", "", OPT_SWITCH, "convert non-destructivly all json files into csv"),
    COption("csv2json", "c", "", OPT_SWITCH, "convert non-destructivly all csv files into json"),
    COption("summarize", "s", "", OPT_SWITCH, "if true produce summary data as well"),
    COption("bucketSize", "b", "<uint64>", OPT_FLAG, "the size of the buckets to summarize"),
    COption("audit", "a", "", OPT_SWITCH, "audit the data"),
    COption("", "", "", OPT_DESCRIPTION, "Handle pouch data in various ways."),
    // clang-format on
    // END_CODE_OPTIONS
};
static const size_t nParams = sizeof(params) / sizeof(COption);

//---------------------------------------------------------------------------------------------------
bool COptions::parseArguments(string_q& command) {
    if (!standardOptions(command))
        return false;

    // BEG_CODE_LOCAL_INIT
    bool freshen = false;
    bool json2csv = false;
    bool csv2json = false;
    bool audit = false;
    // END_CODE_LOCAL_INIT

    Init();
    explode(arguments, command, ' ');
    for (auto arg : arguments) {
        string_q orig = arg;
        if (false) {
            // do nothing -- make auto code generation easier
            // BEG_CODE_AUTO
        } else if (arg == "-f" || arg == "--freshen") {
            freshen = true;

        } else if (arg == "-j" || arg == "--json2csv") {
            json2csv = true;

        } else if (arg == "-c" || arg == "--csv2json") {
            csv2json = true;

        } else if (arg == "-s" || arg == "--summarize") {
            summarize = true;

        } else if (startsWith(arg, "-b:") || startsWith(arg, "--bucketSize:")) {
            if (!confirmUint("bucketSize", bucketSize, arg))
                return false;

        } else if (arg == "-a" || arg == "--audit") {
            audit = true;

        } else if (startsWith(arg, '-')) {  // do not collapse

            if (!builtInCmd(arg)) {
                return usage("Invalid option: " + arg);
            }

            // END_CODE_AUTO
        }
    }

    // BEG_DEBUG_DISPLAY
    LOG_TEST_BOOL("freshen", freshen);
    LOG_TEST_BOOL("json2csv", json2csv);
    LOG_TEST_BOOL("csv2json", csv2json);
    LOG_TEST_BOOL("summarize", summarize);
    LOG_TEST("bucketSize", bucketSize, (bucketSize == 5000));
    LOG_TEST_BOOL("audit", audit);
    // END_DEBUG_DISPLAY

    if (json2csv && csv2json)
        return usage("Choose on of --json2csv or --csv2json, not both.");

    if (!loadTimestamps())
        return usage("Could not load timestamps file");

    if (json2csv)
        return handle_json_2_csv();

    if (csv2json)
        return handle_csv_2_json();

    if (summarize)
        return handle_summarize();

    if (audit)
        return handle_audit();

    return true;
}

//---------------------------------------------------------------------------------------------------
void COptions::Init(void) {
    registerOptions(nParams, params);

    // BEG_CODE_INIT
    summarize = false;
    bucketSize = 5000;
    // END_CODE_INIT

    if (tsMemMap)
        delete[] tsMemMap;
    tsMemMap = nullptr;
    tsCnt = 0;
}

//---------------------------------------------------------------------------------------------------
COptions::COptions(void) {
    tsMemMap = nullptr;
    tsCnt = 0;
    Init();

    // BEG_CODE_NOTES
    // clang-format off
    // clang-format on
    // END_CODE_NOTES

    // BEG_ERROR_MSG
    // END_ERROR_MSG
}

//--------------------------------------------------------------------------------
COptions::~COptions(void) {
}

//--------------------------------------------------------------------------------
bool COptions::loadTimestamps(void) {
    if (tsMemMap)
        return true;

    if (!freshenTimestamps(getBlockProgress(BP_CLIENT).client)) {
        LOG_INFO("Failed to freshen timestamps");
        return false;
    }

    // Note that we do not need to allocate memory since this is a memory mapped file
    if (!::loadTimestamps(&tsMemMap, tsCnt)) {
        LOG_INFO("Failed to load timestamps");
        return false;
    }

    LOG_INFO("Loaded ", tsCnt, " timestamps to ", tsMemMap);
    return true;
}
