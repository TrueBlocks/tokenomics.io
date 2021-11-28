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

namespace qblocks {

// EXISTING_CODE
// EXISTING_CODE

//--------------------------------------------------------------------------
class CPayout : public CBaseNode {
  public:
    address_t address;
    blknum_t bn;
    blknum_t txid;
    blknum_t logid;
    string_q type;
    double amount;

  public:
    CPayout(void);
    CPayout(const CPayout& pa);
    virtual ~CPayout(void);
    CPayout& operator=(const CPayout& pa);

    DECLARE_NODE(CPayout);

    // EXISTING_CODE
    explicit CPayout(string_q& line);
    // EXISTING_CODE
    bool operator==(const CPayout& it) const;
    bool operator!=(const CPayout& it) const {
        return !operator==(it);
    }
    friend bool operator<(const CPayout& v1, const CPayout& v2);
    friend ostream& operator<<(ostream& os, const CPayout& it);

  protected:
    void clear(void);
    void initialize(void);
    void duplicate(const CPayout& pa);
    bool readBackLevel(CArchive& archive) override;

    // EXISTING_CODE
    // EXISTING_CODE
};

//--------------------------------------------------------------------------
inline CPayout::CPayout(void) {
    initialize();
    // EXISTING_CODE
    // EXISTING_CODE
}

//--------------------------------------------------------------------------
inline CPayout::CPayout(const CPayout& pa) {
    // EXISTING_CODE
    // EXISTING_CODE
    duplicate(pa);
}

// EXISTING_CODE
// EXISTING_CODE

//--------------------------------------------------------------------------
inline CPayout::~CPayout(void) {
    clear();
    // EXISTING_CODE
    // EXISTING_CODE
}

//--------------------------------------------------------------------------
inline void CPayout::clear(void) {
    // EXISTING_CODE
    // EXISTING_CODE
}

//--------------------------------------------------------------------------
inline void CPayout::initialize(void) {
    CBaseNode::initialize();

    address = "";
    bn = 0;
    txid = 0;
    logid = 0;
    type = "";
    amount = 0.0;

    // EXISTING_CODE
    // EXISTING_CODE
}

//--------------------------------------------------------------------------
inline void CPayout::duplicate(const CPayout& pa) {
    clear();
    CBaseNode::duplicate(pa);

    address = pa.address;
    bn = pa.bn;
    txid = pa.txid;
    logid = pa.logid;
    type = pa.type;
    amount = pa.amount;

    // EXISTING_CODE
    // EXISTING_CODE
}

//--------------------------------------------------------------------------
inline CPayout& CPayout::operator=(const CPayout& pa) {
    duplicate(pa);
    // EXISTING_CODE
    // EXISTING_CODE
    return *this;
}

//-------------------------------------------------------------------------
inline bool CPayout::operator==(const CPayout& it) const {
    // EXISTING_CODE
    // EXISTING_CODE
    // No default equal operator in class definition, assume none are equal (so find fails)
    return false;
}

//-------------------------------------------------------------------------
inline bool operator<(const CPayout& v1, const CPayout& v2) {
    // EXISTING_CODE
    // EXISTING_CODE
    // No default sort defined in class definition, assume already sorted, preserve ordering
    return true;
}

//---------------------------------------------------------------------------
typedef vector<CPayout> CPayoutArray;
extern CArchive& operator>>(CArchive& archive, CPayoutArray& array);
extern CArchive& operator<<(CArchive& archive, const CPayoutArray& array);

//---------------------------------------------------------------------------
extern const char* STR_DISPLAY_PAYOUT;

//---------------------------------------------------------------------------
// EXISTING_CODE
typedef map<address_t, CPayout> CPayoutMap;
// EXISTING_CODE
}  // namespace qblocks
