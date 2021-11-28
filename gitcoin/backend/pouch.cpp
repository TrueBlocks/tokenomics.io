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
#include "etherlib.h"
#include "options.h"

extern const char* STR_OUTPUT;
extern const char* STR_BALANCE_OUTPUT;
extern bool saveRecords(const CRecordArray& records);
static uint64_t key = 1;
//----------------------------------------------------------------
int main(int argc, const char* argv[]) {
    etherlib_init(quickQuitHandler);
    CRecord::registerClass();
    CBalance::registerClass();
    CPayout::registerClass();
    CDonation::registerClass();

    // Parse command line, allowing for command files
    COptions options;
    if (!options.prepareArguments(argc, argv))
        return 0;

    for (auto command : options.commandLines) {
        if (!options.parseArguments(command))
            return 1;

        if (!options.loadPayouts())
            return options.usage("Could not load the payouts file.");

        if (!options.loadGrantList())
            return options.usage("Could not load grants list. Are you in the pouch folder?");

        options.freshen_loop();
    }

    etherlib_cleanup();

    return 0;
}

// //----------------------------------------------------------------
// class CMonitorUpdater {
//   public:
//     blkrange_t updateRange;
//     CAddressBoolMap monitorMap;
//     CAddressBoolMap updateMap;
//     CMonitorUpdater(void) {
//         updateRange = make_pair(NOPOS, NOPOS);
//     }
// };

// //----------------------------------------------------------------
// bool visitAddrs(const CAppearance& item, void* data) {
//     cerr << "Checking address " << item.addr << "\r";
//     cerr.flush();
//     CMonitorUpdater* check = (CMonitorUpdater*)data;
//     if (check->monitorMap[item.addr] && !check->updateMap[item.addr]) {
//         check->updateMap[item.addr] = true;
//         cout << "Address " << item.addr << " needs update." << endl;
//     }
//     return true;
// }

// //----------------------------------------------------------------
// bool getLatestMonitoredBlock(const string_q& path, void* data) {
//     if (endsWith(path, '/')) {
//         return forEveryFileInFolder(path + "*", getLatestMonitoredBlock, data);

//     } else {
//         CMonitorUpdater* checkup = (CMonitorUpdater*)data;
//         if (contains(path, ".acct.bin")) {
//             CMonitor m;
//             m.address = substitute(substitute(path, m.getMonitorPath(""), ""), ".acct.bin", "");
//             checkup->updateRange.first = min(checkup->updateRange.first, m.getLastVisitedBlock());
//             checkup->monitorMap[m.address] = true;
//         }
//     }
//     return true;
// }

//----------------------------------------------------------------
bool COptions::freshen_loop(void) {
    // Before entering the loop, we figure out the list of monitored address and where we need to start the update
    // for (auto grant : grants) {
    //     cout << 

    // }
    // uint64_t sizeBefore = 
    // CMonitor m;
    // CMonitorUpdater checkup;
    // forEveryFileInFolder(m.getMonitorPath(""), getLatestMonitoredBlock, &checkup);
    // //    for (auto mon : checkup.monitorMap)
    // //        cout << mon.first << endl;
    // checkup.updateRange.second = getBlockProgress(BP_FINAL).finalized;

    // //    while (!shouldQuit()) {
    // //        cout << checkup.updateRange.first << " " << checkup.updateRange.second << endl;
    // //        LOG_INFO("Sleeping for 30 seconds...");
    // //        usleep(15 * 100000);
    // //        blknum_t latest = getBlockProgress(BP_FINAL).finalized;
    // //        checkup.updateRange.first = min(checkup.updateRange.second + 1, latest);
    // //        checkup.updateRange.second = latest;
    // //    }
    // return true;
            if (!loadRecords()) {
                LOG_INFO(cTeal, "Refreshing records...", cOff);
                key = 1;
                if (!updateAll())
                    return usage("Could not load records.");
            }

            while (!shouldQuit()) {
                ostringstream os;
                os << "export const grantsData = [\n";
                for (auto record : records) {
                    ostringstream oss;
                    bool first = true;
                    for (auto balance : record.balances) {
                        if (!first)
                            oss << ",";
                        oss << "[" << balance.Format(STR_BALANCE_OUTPUT) << "]";
                        first = false;
                    }
                    os << substitute(record.Format(STR_OUTPUT), "++BALANCES++", oss.str()) << endl;
                }
                os << "];";
                stringToAsciiFile("../src/grants-data.js", os.str());
                //            cerr << "Sleeping for 28 seconds";
                //            size_t cnt = 0;
                //            while (++cnt < 28 && !shouldQuit())
                //            {
                //                cerr << ".";
                //                cerr.flush();
                //                sleep(1);
                //            }
                //            cerr << endl;
                key = 1;
                return 0;
            }
            return true;
}

//----------------------------------------------------------------
bool saveRecords(const CRecordArray& records) {
    CArchive archive(WRITING_ARCHIVE);
    if (archive.Lock("./data/records.bin", modeWriteCreate, LOCK_WAIT)) {
        archive << records;
        archive.Release();
        return true;
    }
    return false;
}

/*
 CMonitorUpdater checkup
checkup.monitorMap
forEveryAddressInIndex
*/

//----------------------------------------------------------------
bool COptions::updateAll(void) {
    records.clear();
    for (auto grant : grants) {
        CRecord record;
        if (updateOne(record, grant))
            records.push_back(record);
    }
    return saveRecords(records) && records.size();
}

//----------------------------------------------------------------
bool COptions::updateOne(CRecord& record, CAccountName& grant) {
    record.key = key++;
    record.grant_id = str_2_Uint(substitute(grant.name, "Grant ", ""));
    nextTokenClear(grant.name, ' ');
    nextTokenClear(grant.name, ' ');
    record.name = substitute(grant.name.substr(0, 60), "'", "&#39;");

    record.type = "logs";  // types[key % 3];
    record.address = grant.address;
    record.slug = grant.source;
    record.core = contains(grant.tags, ":Core");
    ostringstream cmd;
    cmd << "tail -1 data/" << record.address << ".csv | sed 's/\\\"//g' | cut -f1 -d, | sed 's/blocknumber/0/'";
    record.last_block = str_2_Uint(doCommand(cmd.str()));
    if (record.last_block == 0) {
        record.date = "n/a";
        record.last_ts = 0;
        return false;
    }
    if (record.last_block * 2 > (tsCnt * 2) + 2) {
        usage("Last block * 2 (" + uint_2_Str(record.last_block * 2) + ") greater than tsCnt (" + uint_2_Str(tsCnt) +
              ")");
        quickQuitHandler(1);
    }
    record.last_ts = tsMemMap[(record.last_block * 2) + 1];
    record.date = ts_2_Date(record.last_ts).Format(FMT_JSON);
    record.matched = matches[record.address].amount;
    record.claimed = claims[record.address].amount;

    static blknum_t latest = NOPOS;
    if (latest == NOPOS)
        latest = getBlockProgress(BP_FINAL).finalized;
    CBalance bal;
    bal.asset = "ETH";
    wei_t balance = getBalanceAt(grant.address, latest);
    bal.balance = wei_2_Ether(balance, 18);
    bal.balance = double_2_Str(str_2_Double(bal.balance), 12);
    record.balances.push_back(bal);

    string_q jsonFile = "./data/" + record.address + ".json";
    string_q csvFile = "./data/" + record.address + ".csv";
    string_q monFile = getCachePath("monitors/" + toLower(grant.address) + ".acct.bin");
    record.tx_cnt = (fileExists(monFile) ? (fileSize(monFile) / sizeof(CAppearance_base)) : 0);

    if (fileExists(csvFile)) {
        record.log_cnt = str_2_Uint(doCommand("wc " + csvFile));
        if (record.log_cnt > 0)
            record.log_cnt -= 1;
        if (record.address == STR_PAYOUT) {
            record.donation_cnt = str_2_Uint(doCommand("cat " + csvFile + " | grep Payout | wc"));
        } else if (record.address == STR_ROUND5) {
            record.donation_cnt = str_2_Uint(doCommand("cat " + csvFile + " | grep " + STR_ROUND5 + " | wc"));
        } else {
            record.donation_cnt = str_2_Uint(doCommand("cat " + csvFile + " | grep Donation | wc"));
        }
    } else {
        record.log_cnt = record.tx_cnt;
        record.donation_cnt = 0;
    }

    LOG_INFO("  processing grant ", grant.address, " ", grant.name.substr(0, 60));
    return true;
}

//----------------------------------------------------------------
 bool COptions::loadGrantList(void) {
    CAccountName name;
    string_q contents = asciiFileToString("./app-data/grants.json");
    while (name.parseJson3(contents)) {
        if (name.address ==
            "0x322d58b9e75a6918f7e7849aee0ff09369977e08")  // Skip this. It's both inactive and really big
            continue;
        name.address = toLower(name.address);
        grants.push_back(name);
        name = CAccountName();
    }
    return grants.size();
}

//----------------------------------------------------------------
bool COptions::loadPayouts(void) {
    CStringArray lines;
    asciiFileToLines("./app-data/payouts.csv", lines);
    for (auto line : lines) {
        replaceAll(line, "(", ",");
        replaceAny(line, ";\") ", "");
        CPayout payout(line);
        if (payout.type == "PayoutAdded")
            matches[payout.address] = payout;
        else if (payout.type == "PayoutClaimed")
            claims[payout.address] = payout;
        else {
            LOG_ERR("Invalid payout type: ", payout.type);
            return false;
        }
    }
    return true;
}

//----------------------------------------------------------------
bool COptions::loadRecords(void) {
    CArchive archive(READING_ARCHIVE);
    if (archive.Lock("./data/records.bin", modeReadOnly, LOCK_NOWAIT)) {
        archive >> records;
        archive.Release();
        // check to see if we need to update anything
        time_q recordsTime = fileLastModifyDate("./data/records.bin");
        for (auto record : records) {
            time_q csvTime = fileLastModifyDate("./data/" + record.address + ".csv");
            if (csvTime >= recordsTime)
                return false;  // we need to refresh
        }
        return true;
    }
    return false;
}

//----------------------------------------------------------------
const char* STR_OUTPUT =
    "  {\n"
    "    key: [{KEY}],\n"
    "    date: '[{DATE}]',\n"
    "    last_block: '[{LAST_BLOCK}]',\n"
    "    last_ts: '[{LAST_TS}]',\n"
    "    type: '[{TYPE}]',\n"
    "    grant_id: [{GRANT_ID}],\n"
    "    address: '[{ADDRESS}]',\n"
    "    name: '[{NAME}]',\n"
    "    slug: '[{SLUG}]',\n"
    "    tx_cnt: [{TX_CNT}],\n"
    "    log_cnt: [{LOG_CNT}],\n"
    "    donation_cnt: [{DONATION_CNT}],\n"
    "    matched: [{MATCHED}],\n"
    "    claimed: [{CLAIMED}],\n"
    "    balances: ++BALANCES++,\n"
    "    core: [{CORE}],\n"
    "  },";

//----------------------------------------------------------------
const char* STR_BALANCE_OUTPUT =
    "{\n"
    "      asset: 'ETH',\n"
    "      balance: '[{BALANCE}]'\n"
    "    }";
