
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

class CBucket {
  public:
    CAddressUintMap donors;
    CAddressUintMap recipients;
    CAddressUintMap tokens;
    CDonationArray donations;
    blknum_t seen;
    blknum_t curBucket;
    blknum_t bucketSize;
    size_t cnt;
    map<int, bool> fieldMap;
    CBucket(void) : seen(0), curBucket(10244000), bucketSize(1000), cnt(0) {
        fieldMap[0] = true;
        fieldMap[1] = true;
        fieldMap[2] = true;
        fieldMap[11] = true;
        fieldMap[12] = true;
        fieldMap[13] = true;
        fieldMap[14] = true;
    }
};

//--------------------------------------------------------------------
bool countByBucket(const char* line, void* data) {
    static bool first = true;
    if (first) {  // header
        first = false;
        return true;
    }
    
    CBucket* bucket = (CBucket*)data;
    bucket->seen++;
    
    blknum_t thisBlock = str_2_Uint(&line[1]);
    blknum_t thisBucket = (thisBlock / bucket->bucketSize) * bucket->bucketSize;
    blknum_t prevBucket = bucket->curBucket;
    //    cout << " seen: " << padNum9T(bucket->seen);
    //    cout << " thisBlock: " << padNum9T(thisBlock);
    //    cout << " thisBucket: " << padNum9T(thisBucket);
    //    cout << " prevBucket: " << padNum9T(prevBucket);
    //    cout << " size: " << padNum9T(bucket->bucketSize);
    if (thisBucket > prevBucket) {
        blknum_t diff = (thisBucket - prevBucket);
        if (diff > bucket->bucketSize) {
            // There was a gap (more than one bucket), render zeros
            while (diff > bucket->bucketSize) {
                cout << prevBucket << "\t" << 0 << endl;
                prevBucket += bucket->bucketSize;
                diff -= bucket->bucketSize;
            }
        }
        cout << prevBucket << "\t" << bucket->seen << endl;
        bucket->curBucket = thisBucket;
        bucket->seen = 0;
        //         cout << " cur: " << padNum9T(thisBucket);
        //         cout << " prevBucket: " << padNum9T(prevBucket);
        //         cout << " cur: " << padNum9T(thisBucket);
        //         cout << " diff: " << padNum9T(diff);
        //         cout << endl;
    } else {
        //         cout << "\r";
        //         cout.flush();
    }
    return !shouldQuit();
}

//--------------------------------------------------------------------
bool listDonations(const char* line, void* data) {
    static bool first = true;
    if (first) {  // header
        first = false;
        return true;
    }

    if (!strstr(line, "DonationSent"))
        return true;

    CBucket* bucket = (CBucket*)data;

    char buffer[4096];
    int pos = 0;
    const char* s = line;
    int cur = 0;
    while (*s) {
        switch (*s) {
            case '\"':
            case ' ':
            case ')':
            case ';':
                break;
            case '(':
            case ',':
                cur++;
                if (bucket->fieldMap[cur]) {
                    buffer[pos++] = ',';
                }
                break;
            default:
                if (bucket->fieldMap[cur] && *s != '\n') {
                    buffer[pos++] = *s;
                }
        }
        s++;
    }
    buffer[pos] = '\0';
    // cout << buffer << endl;

    CStringArray parts;
    explode(parts, buffer, ',');
    CDonation donation;
    donation.block = str_2_Uint(parts[0]);
    donation.tx_id = str_2_Uint(parts[1]);
    donation.log_id = str_2_Uint(parts[2]);
    // donation.date;
    donation.token = parts[3];
    donation.amount = parts[4];
    donation.recipient = parts[5];
    donation.donor = parts[6];
    bucket->donations.push_back(donation);
    bucket->donors[donation.donor]++;
    bucket->recipients[donation.recipient]++;
    bucket->tokens[donation.token]++;
//    static bool ff = true;
//    if (!ff)
//        cout << ",";
//    cout << donation;
    LOG_INFO(donation.block, "\r");
//    ff = false;
    // 12016322
    // 292
    // 199
    // 0x7d655c57f71464b6f83811c55d84009cd9f5221c
    // 0x3bb7428b25f9bdad9bd2faa4c6a7a9e5d5882657e96c1d24cc41c1d6c1910a98
    // 0x0000000000000000000000006b175474e89094c44da98b954eedeac495271d0f
    // 0x0000000000000000000000000000000000000000000000006124fee993bc0000
    // 0x000000000000000000000000ed647437df2898f681918f3256320064cbae5ca5
    // 0x000000000000000000000000201930f7b5a5fa20496fe90411f7babef35e94dc
    // DonationSent
    // 0x6b175474e89094c44da98b954eedeac495271d0f /*token*/
    // 7000000000000000000 /*amount*/
    // 0x201930f7b5a5fa20496fe90411f7babef35e94dc /*dest*/
    // 0xed647437df2898f681918f3256320064cbae5ca5 /*donor*/

    return true;
}

//--------------------------------------------------------------------
bool COptions::handle_summarize(void) {
    CBucket bucket;
    bucket.bucketSize = bucketSize;
    
//    forEveryLineInAsciiFile("./data/" + STR_ROUND8 + ".csv", countByBucket, &bucket);
//    cout << bucket.curBucket << "\t" << bucket.seen << endl;
    
    bucket = CBucket();
//    cout << "[";
//    forEveryLineInAsciiFile("./data/" + STR_ROUND8 + ".csv", listDonations, &bucket);
//    cout << "]";

    forEveryLineInAsciiFile("./data/" + STR_ROUND8 + ".csv", listDonations, &bucket);
//    cout << "[";
    cout << "type\taddress\tcount" << endl;
    for (auto donor : bucket.donors)
        cout << "donor\t\t\t" << donor.first << "\t\t\t" << donor.second << endl;
//        cout << "{ \"type\": \"donor\", \"address\": \"" << donor.first << "\", \"count\": \"" << donor.second << "\" }" << endl;
    for (auto recipient : bucket.recipients)
        cout << "recipient\t\t\t" << recipient.first << "\t\t\t" << recipient.second << endl;
//        cout << "{ \"type\": \"recipient\", \"address\": \"" << recipient.first << "\", \"count\": \"" << recipient.second << "\" }" << endl;
    for (auto token : bucket.tokens)
        cout << "token\t\t\t" << token.first << "\t\t\t" << token.second << endl;
//        cout << "{ \"type\": \"token\", \"address\": \"" << token.first << "\", \"count\": \"" << token.second << "\" }" << endl;
//    cout << "]";
    return !shouldQuit();
}
