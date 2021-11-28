#pragma once
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
#include "acctlib.h"
#include "classes/record.h"
#include "classes/payout.h"
#include "classes/donation.h"

// BEG_ERROR_DEFINES
// END_ERROR_DEFINES

//-----------------------------------------------------------------------------
class COptions : public COptionsBase {
  public:
    // BEG_CODE_DECLARE
    bool summarize;
    uint64_t bucketSize;
    // END_CODE_DECLARE

    CPayoutMap matches;
    CPayoutMap claims;
    CAccountNameArray grants;
    CRecordArray records;
    uint32_t* tsMemMap; // not allocated
    size_t tsCnt;

    COptions(void);
    ~COptions(void);

    bool parseArguments(string_q& command);
    void Init(void);

    bool freshen_loop(void);

    bool handle_json_2_csv(void);
    bool handle_csv_2_json(void);
    bool handle_audit(void);
    bool handle_summarize(void);

    bool loadGrantList(void);
    bool loadPayouts(void);
    bool loadRecords(void);
    bool loadTimestamps(void);
    bool updateOne(CRecord& record, CAccountName& grant);
    bool updateAll(void);
};

//-----------------------------------------------------------------------------
extern bool visitNonEmptyBlock(CBlock& node, void* data);
extern bool visitEmptyBlock(CBlock& node, void* data);
extern const string_q STR_PAYOUT;
extern const string_q STR_ROUND8;
extern const string_q STR_ROUND5;
