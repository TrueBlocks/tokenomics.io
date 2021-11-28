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
#include "utillib.h"
#include "abi.h"

namespace qblocks {

// EXISTING_CODE
// EXISTING_CODE

//--------------------------------------------------------------------------
class CLogEntry_min : public CBaseNode {
  public:
    address_t address;
    blknum_t blockNumber;
    blknum_t logIndex;
    CTopicArray topics;
    string_q data;
    CFunction articulatedLog;
    string_q compressedLog;
    blknum_t transactionIndex;

  public:
    CLogEntry_min(void);
    CLogEntry_min(const CLogEntry_min& lo);
    virtual ~CLogEntry_min(void);
    CLogEntry_min& operator=(const CLogEntry_min& lo);

    DECLARE_NODE(CLogEntry_min);

    const CBaseNode* getObjectAt(const string_q& fieldName, size_t index) const override;
    const string_q getStringAt(const string_q& fieldName, size_t i) const override;

    // EXISTING_CODE
    CLogEntry_min& operator=(const CLogEntry& lo);
    // EXISTING_CODE
    bool operator==(const CLogEntry_min& it) const;
    bool operator!=(const CLogEntry_min& it) const {
        return !operator==(it);
    }
    friend bool operator<(const CLogEntry_min& v1, const CLogEntry_min& v2);
    friend ostream& operator<<(ostream& os, const CLogEntry_min& it);

  protected:
    void clear(void);
    void initialize(void);
    void duplicate(const CLogEntry_min& lo);
    bool readBackLevel(CArchive& archive) override;

    // EXISTING_CODE
    // EXISTING_CODE
};

//--------------------------------------------------------------------------
inline CLogEntry_min::CLogEntry_min(void) {
    initialize();
    // EXISTING_CODE
    // EXISTING_CODE
}

//--------------------------------------------------------------------------
inline CLogEntry_min::CLogEntry_min(const CLogEntry_min& lo) {
    // EXISTING_CODE
    // EXISTING_CODE
    duplicate(lo);
}

// EXISTING_CODE
// EXISTING_CODE

//--------------------------------------------------------------------------
inline CLogEntry_min::~CLogEntry_min(void) {
    clear();
    // EXISTING_CODE
    // EXISTING_CODE
}

//--------------------------------------------------------------------------
inline void CLogEntry_min::clear(void) {
    // EXISTING_CODE
    // EXISTING_CODE
}

//--------------------------------------------------------------------------
inline void CLogEntry_min::initialize(void) {
    CBaseNode::initialize();

    address = "";
    blockNumber = 0;
    logIndex = 0;
    topics.clear();
    data = "";
    articulatedLog = CFunction();
    compressedLog = "";
    transactionIndex = 0;

    // EXISTING_CODE
    // EXISTING_CODE
}

//--------------------------------------------------------------------------
inline void CLogEntry_min::duplicate(const CLogEntry_min& lo) {
    clear();
    CBaseNode::duplicate(lo);

    address = lo.address;
    blockNumber = lo.blockNumber;
    logIndex = lo.logIndex;
    topics = lo.topics;
    data = lo.data;
    articulatedLog = lo.articulatedLog;
    compressedLog = lo.compressedLog;
    transactionIndex = lo.transactionIndex;

    // EXISTING_CODE
    // EXISTING_CODE
}

//--------------------------------------------------------------------------
inline CLogEntry_min& CLogEntry_min::operator=(const CLogEntry_min& lo) {
    duplicate(lo);
    // EXISTING_CODE
    // EXISTING_CODE
    return *this;
}

//-------------------------------------------------------------------------
inline bool CLogEntry_min::operator==(const CLogEntry_min& it) const {
    // EXISTING_CODE
    // EXISTING_CODE
    // No default equal operator in class definition, assume none are equal (so find fails)
    return false;
}

//-------------------------------------------------------------------------
inline bool operator<(const CLogEntry_min& v1, const CLogEntry_min& v2) {
    // EXISTING_CODE
    // EXISTING_CODE
    // No default sort defined in class definition, assume already sorted, preserve ordering
    return true;
}

//---------------------------------------------------------------------------
typedef vector<CLogEntry_min> CLogEntry_minArray;
extern CArchive& operator>>(CArchive& archive, CLogEntry_minArray& array);
extern CArchive& operator<<(CArchive& archive, const CLogEntry_minArray& array);

//---------------------------------------------------------------------------
extern CArchive& operator<<(CArchive& archive, const CLogEntry_min& log);
extern CArchive& operator>>(CArchive& archive, CLogEntry_min& log);

//---------------------------------------------------------------------------
extern const char* STR_DISPLAY_LOGENTRY_MIN;

//---------------------------------------------------------------------------
// EXISTING_CODE
// EXISTING_CODE
}  // namespace qblocks
