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
class CDonation : public CBaseNode {
  public:
    blknum_t block;
    blknum_t tx_id;
    blknum_t log_id;
    string_q date;
    string_q token;
    string_q amount;
    address_t recipient;
    address_t donor;

  public:
    CDonation(void);
    CDonation(const CDonation& dona);
    virtual ~CDonation(void);
    CDonation& operator=(const CDonation& dona);

    DECLARE_NODE(CDonation);

    // EXISTING_CODE
    // EXISTING_CODE
    bool operator==(const CDonation& it) const;
    bool operator!=(const CDonation& it) const {
        return !operator==(it);
    }
    friend bool operator<(const CDonation& v1, const CDonation& v2);
    friend ostream& operator<<(ostream& os, const CDonation& it);

  protected:
    void clear(void);
    void initialize(void);
    void duplicate(const CDonation& dona);
    bool readBackLevel(CArchive& archive) override;

    // EXISTING_CODE
    // EXISTING_CODE
};

//--------------------------------------------------------------------------
inline CDonation::CDonation(void) {
    initialize();
    // EXISTING_CODE
    // EXISTING_CODE
}

//--------------------------------------------------------------------------
inline CDonation::CDonation(const CDonation& dona) {
    // EXISTING_CODE
    // EXISTING_CODE
    duplicate(dona);
}

// EXISTING_CODE
// EXISTING_CODE

//--------------------------------------------------------------------------
inline CDonation::~CDonation(void) {
    clear();
    // EXISTING_CODE
    // EXISTING_CODE
}

//--------------------------------------------------------------------------
inline void CDonation::clear(void) {
    // EXISTING_CODE
    // EXISTING_CODE
}

//--------------------------------------------------------------------------
inline void CDonation::initialize(void) {
    CBaseNode::initialize();

    block = 0;
    tx_id = 0;
    log_id = 0;
    date = "";
    token = "";
    amount = "";
    recipient = "";
    donor = "";

    // EXISTING_CODE
    // EXISTING_CODE
}

//--------------------------------------------------------------------------
inline void CDonation::duplicate(const CDonation& dona) {
    clear();
    CBaseNode::duplicate(dona);

    block = dona.block;
    tx_id = dona.tx_id;
    log_id = dona.log_id;
    date = dona.date;
    token = dona.token;
    amount = dona.amount;
    recipient = dona.recipient;
    donor = dona.donor;

    // EXISTING_CODE
    // EXISTING_CODE
}

//--------------------------------------------------------------------------
inline CDonation& CDonation::operator=(const CDonation& dona) {
    duplicate(dona);
    // EXISTING_CODE
    // EXISTING_CODE
    return *this;
}

//-------------------------------------------------------------------------
inline bool CDonation::operator==(const CDonation& it) const {
    // EXISTING_CODE
    // EXISTING_CODE
    // No default equal operator in class definition, assume none are equal (so find fails)
    return false;
}

//-------------------------------------------------------------------------
inline bool operator<(const CDonation& v1, const CDonation& v2) {
    // EXISTING_CODE
    // EXISTING_CODE
    // No default sort defined in class definition, assume already sorted, preserve ordering
    return true;
}

//---------------------------------------------------------------------------
typedef vector<CDonation> CDonationArray;
extern CArchive& operator>>(CArchive& archive, CDonationArray& array);
extern CArchive& operator<<(CArchive& archive, const CDonationArray& array);

//---------------------------------------------------------------------------
extern const char* STR_DISPLAY_DONATION;

//---------------------------------------------------------------------------
// EXISTING_CODE
// EXISTING_CODE
}  // namespace qblocks
