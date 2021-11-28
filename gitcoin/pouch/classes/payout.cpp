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
#include "payout.h"

namespace qblocks {

//---------------------------------------------------------------------------
IMPLEMENT_NODE(CPayout, CBaseNode);

//---------------------------------------------------------------------------
static string_q nextPayoutChunk(const string_q& fieldIn, const void* dataPtr);
static string_q nextPayoutChunk_custom(const string_q& fieldIn, const void* dataPtr);

//---------------------------------------------------------------------------
void CPayout::Format(ostream& ctx, const string_q& fmtIn, void* dataPtr) const {
    if (!m_showing)
        return;

    // EXISTING_CODE
    // EXISTING_CODE

    string_q fmt = (fmtIn.empty() ? expContext().fmtMap["payout_fmt"] : fmtIn);
    if (fmt.empty()) {
        if (expContext().exportFmt == YAML1) {
            toYaml(ctx);
        } else {
            toJson(ctx);
        }
        return;
    }

    // EXISTING_CODE
    // EXISTING_CODE

    while (!fmt.empty())
        ctx << getNextChunk(fmt, nextPayoutChunk, this);
}

//---------------------------------------------------------------------------
string_q nextPayoutChunk(const string_q& fieldIn, const void* dataPtr) {
    if (dataPtr)
        return reinterpret_cast<const CPayout*>(dataPtr)->getValueByName(fieldIn);

    // EXISTING_CODE
    // EXISTING_CODE

    return fldNotFound(fieldIn);
}

//---------------------------------------------------------------------------
string_q CPayout::getValueByName(const string_q& fieldName) const {
    // Give customized code a chance to override first
    string_q ret = nextPayoutChunk_custom(fieldName, this);
    if (!ret.empty())
        return ret;

    // EXISTING_CODE
    // EXISTING_CODE

    // Return field values
    switch (tolower(fieldName[0])) {
        case 'a':
            if (fieldName % "address") {
                return addr_2_Str(address);
            }
            if (fieldName % "amount") {
                return double_2_Str(amount, 5);
            }
            break;
        case 'b':
            if (fieldName % "bn") {
                return uint_2_Str(bn);
            }
            break;
        case 'l':
            if (fieldName % "logid") {
                return uint_2_Str(logid);
            }
            break;
        case 't':
            if (fieldName % "txid") {
                return uint_2_Str(txid);
            }
            if (fieldName % "type") {
                return type;
            }
            break;
        default:
            break;
    }

    // EXISTING_CODE
    // EXISTING_CODE

    // Finally, give the parent class a chance
    return CBaseNode::getValueByName(fieldName);
}

//---------------------------------------------------------------------------------------------------
bool CPayout::setValueByName(const string_q& fieldNameIn, const string_q& fieldValueIn) {
    string_q fieldName = fieldNameIn;
    string_q fieldValue = fieldValueIn;

    // EXISTING_CODE
    // EXISTING_CODE

    switch (tolower(fieldName[0])) {
        case 'a':
            if (fieldName % "address") {
                address = str_2_Addr(fieldValue);
                return true;
            }
            if (fieldName % "amount") {
                amount = str_2_Double(fieldValue);
                return true;
            }
            break;
        case 'b':
            if (fieldName % "bn") {
                bn = str_2_Uint(fieldValue);
                return true;
            }
            break;
        case 'l':
            if (fieldName % "logid") {
                logid = str_2_Uint(fieldValue);
                return true;
            }
            break;
        case 't':
            if (fieldName % "txid") {
                txid = str_2_Uint(fieldValue);
                return true;
            }
            if (fieldName % "type") {
                type = fieldValue;
                return true;
            }
            break;
        default:
            break;
    }
    return false;
}

//---------------------------------------------------------------------------------------------------
void CPayout::finishParse() {
    // EXISTING_CODE
    // EXISTING_CODE
}

//---------------------------------------------------------------------------------------------------
bool CPayout::Serialize(CArchive& archive) {
    if (archive.isWriting())
        return SerializeC(archive);

    // Always read the base class (it will handle its own backLevels if any, then
    // read this object's back level (if any) or the current version.
    CBaseNode::Serialize(archive);
    if (readBackLevel(archive))
        return true;

    // EXISTING_CODE
    // EXISTING_CODE
    archive >> address;
    archive >> bn;
    archive >> txid;
    archive >> logid;
    archive >> type;
    archive >> amount;
    finishParse();
    return true;
}

//---------------------------------------------------------------------------------------------------
bool CPayout::SerializeC(CArchive& archive) const {
    // Writing always write the latest version of the data
    CBaseNode::SerializeC(archive);

    // EXISTING_CODE
    // EXISTING_CODE
    archive << address;
    archive << bn;
    archive << txid;
    archive << logid;
    archive << type;
    archive << amount;

    return true;
}

//---------------------------------------------------------------------------
CArchive& operator>>(CArchive& archive, CPayoutArray& array) {
    uint64_t count;
    archive >> count;
    array.resize(count);
    for (size_t i = 0; i < count; i++) {
        ASSERT(i < array.capacity());
        array.at(i).Serialize(archive);
    }
    return archive;
}

//---------------------------------------------------------------------------
CArchive& operator<<(CArchive& archive, const CPayoutArray& array) {
    uint64_t count = array.size();
    archive << count;
    for (size_t i = 0; i < array.size(); i++)
        array[i].SerializeC(archive);
    return archive;
}

//---------------------------------------------------------------------------
void CPayout::registerClass(void) {
    // only do this once
    if (HAS_FIELD(CPayout, "schema"))
        return;

    size_t fieldNum = 1000;
    ADD_FIELD(CPayout, "schema", T_NUMBER, ++fieldNum);
    ADD_FIELD(CPayout, "deleted", T_BOOL, ++fieldNum);
    ADD_FIELD(CPayout, "showing", T_BOOL, ++fieldNum);
    ADD_FIELD(CPayout, "cname", T_TEXT, ++fieldNum);
    ADD_FIELD(CPayout, "address", T_ADDRESS | TS_OMITEMPTY, ++fieldNum);
    ADD_FIELD(CPayout, "bn", T_BLOCKNUM, ++fieldNum);
    ADD_FIELD(CPayout, "txid", T_BLOCKNUM, ++fieldNum);
    ADD_FIELD(CPayout, "logid", T_BLOCKNUM, ++fieldNum);
    ADD_FIELD(CPayout, "type", T_TEXT | TS_OMITEMPTY, ++fieldNum);
    ADD_FIELD(CPayout, "amount", T_DOUBLE, ++fieldNum);

    // Hide our internal fields, user can turn them on if they like
    HIDE_FIELD(CPayout, "schema");
    HIDE_FIELD(CPayout, "deleted");
    HIDE_FIELD(CPayout, "showing");
    HIDE_FIELD(CPayout, "cname");

    builtIns.push_back(_biCPayout);

    // EXISTING_CODE
    // EXISTING_CODE
}

//---------------------------------------------------------------------------
string_q nextPayoutChunk_custom(const string_q& fieldIn, const void* dataPtr) {
    const CPayout* pay = reinterpret_cast<const CPayout*>(dataPtr);
    if (pay) {
        switch (tolower(fieldIn[0])) {
            // EXISTING_CODE
            // EXISTING_CODE
            case 'p':
                // Display only the fields of this node, not it's parent type
                if (fieldIn % "parsed")
                    return nextBasenodeChunk(fieldIn, pay);
                // EXISTING_CODE
                // EXISTING_CODE
                break;

            default:
                break;
        }
    }

    return "";
}

//---------------------------------------------------------------------------
bool CPayout::readBackLevel(CArchive& archive) {
    bool done = false;
    // EXISTING_CODE
    // EXISTING_CODE
    return done;
}

//-------------------------------------------------------------------------
ostream& operator<<(ostream& os, const CPayout& it) {
    // EXISTING_CODE
    // EXISTING_CODE

    it.Format(os, "", nullptr);
    os << "\n";
    return os;
}

//---------------------------------------------------------------------------
const char* STR_DISPLAY_PAYOUT =
    "[{ADDRESS}]\t"
    "[{BN}]\t"
    "[{TXID}]\t"
    "[{LOGID}]\t"
    "[{TYPE}]\t"
    "[{AMOUNT}]";

//---------------------------------------------------------------------------
// EXISTING_CODE
CPayout::CPayout(string_q& line) {
    bn = str_2_Uint(nextTokenClear(line, ','));
    txid = str_2_Uint(nextTokenClear(line, ','));
    logid = str_2_Uint(nextTokenClear(line, ','));
    type = nextTokenClear(line, ',');
    address = nextTokenClear(line, ',');
    wei_t wei = str_2_Wei(nextTokenClear(line, ','));
    amount = str_2_Double(wei_2_Ether(wei, 18));
}
// EXISTING_CODE
}  // namespace qblocks
