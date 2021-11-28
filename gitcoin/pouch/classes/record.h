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
 * This file was generated with makeClass. Edit only those parts of the code inside
 * of 'EXISTING_CODE' tags.
 */
#include "etherlib.h"
#include "balance.h"

namespace qblocks {

// EXISTING_CODE
// EXISTING_CODE

//--------------------------------------------------------------------------
class CRecord : public CBaseNode {
  public:
    uint64_t key;
    string_q date;
    blknum_t last_block;
    timestamp_t last_ts;
    string_q type;
    uint64_t grant_id;
    address_t address;
    string_q name;
    string_q slug;
    uint64_t tx_cnt;
    uint64_t log_cnt;
    uint64_t core;
    uint64_t donation_cnt;
    double matched;
    double claimed;
    CBalanceArray balances;

  public:
    CRecord(void);
    CRecord(const CRecord& re);
    virtual ~CRecord(void);
    CRecord& operator=(const CRecord& re);

    DECLARE_NODE(CRecord);

    const CBaseNode* getObjectAt(const string_q& fieldName, size_t index) const override;

    // EXISTING_CODE
    // EXISTING_CODE
    bool operator==(const CRecord& it) const;
    bool operator!=(const CRecord& it) const {
        return !operator==(it);
    }
    friend bool operator<(const CRecord& v1, const CRecord& v2);
    friend ostream& operator<<(ostream& os, const CRecord& it);

  protected:
    void clear(void);
    void initialize(void);
    void duplicate(const CRecord& re);
    bool readBackLevel(CArchive& archive) override;

    // EXISTING_CODE
    // EXISTING_CODE
};

//--------------------------------------------------------------------------
inline CRecord::CRecord(void) {
    initialize();
    // EXISTING_CODE
    // EXISTING_CODE
}

//--------------------------------------------------------------------------
inline CRecord::CRecord(const CRecord& re) {
    // EXISTING_CODE
    // EXISTING_CODE
    duplicate(re);
}

// EXISTING_CODE
// EXISTING_CODE

//--------------------------------------------------------------------------
inline CRecord::~CRecord(void) {
    clear();
    // EXISTING_CODE
    // EXISTING_CODE
}

//--------------------------------------------------------------------------
inline void CRecord::clear(void) {
    // EXISTING_CODE
    // EXISTING_CODE
}

//--------------------------------------------------------------------------
inline void CRecord::initialize(void) {
    CBaseNode::initialize();

    key = 0;
    date = "";
    last_block = 0;
    last_ts = 0;
    type = "";
    grant_id = 0;
    address = "";
    name = "";
    slug = "";
    tx_cnt = 0;
    log_cnt = 0;
    core = 0;
    donation_cnt = 0;
    matched = 0.0;
    claimed = 0.0;
    balances.clear();

    // EXISTING_CODE
    // EXISTING_CODE
}

//--------------------------------------------------------------------------
inline void CRecord::duplicate(const CRecord& re) {
    clear();
    CBaseNode::duplicate(re);

    key = re.key;
    date = re.date;
    last_block = re.last_block;
    last_ts = re.last_ts;
    type = re.type;
    grant_id = re.grant_id;
    address = re.address;
    name = re.name;
    slug = re.slug;
    tx_cnt = re.tx_cnt;
    log_cnt = re.log_cnt;
    core = re.core;
    donation_cnt = re.donation_cnt;
    matched = re.matched;
    claimed = re.claimed;
    balances = re.balances;

    // EXISTING_CODE
    // EXISTING_CODE
}

//--------------------------------------------------------------------------
inline CRecord& CRecord::operator=(const CRecord& re) {
    duplicate(re);
    // EXISTING_CODE
    // EXISTING_CODE
    return *this;
}

//-------------------------------------------------------------------------
inline bool CRecord::operator==(const CRecord& it) const {
    // EXISTING_CODE
    // EXISTING_CODE
    // No default equal operator in class definition, assume none are equal (so find fails)
    return false;
}

//-------------------------------------------------------------------------
inline bool operator<(const CRecord& v1, const CRecord& v2) {
    // EXISTING_CODE
    // EXISTING_CODE
    // No default sort defined in class definition, assume already sorted, preserve ordering
    return true;
}

//---------------------------------------------------------------------------
typedef vector<CRecord> CRecordArray;
extern CArchive& operator>>(CArchive& archive, CRecordArray& array);
extern CArchive& operator<<(CArchive& archive, const CRecordArray& array);

//---------------------------------------------------------------------------
extern const char* STR_DISPLAY_RECORD;

//---------------------------------------------------------------------------
// EXISTING_CODE
// EXISTING_CODE
}  // namespace qblocks
