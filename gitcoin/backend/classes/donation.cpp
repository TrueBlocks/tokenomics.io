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
#include "donation.h"

namespace qblocks {

//---------------------------------------------------------------------------
IMPLEMENT_NODE(CDonation, CBaseNode);

//---------------------------------------------------------------------------
static string_q nextDonationChunk(const string_q& fieldIn, const void* dataPtr);
static string_q nextDonationChunk_custom(const string_q& fieldIn, const void* dataPtr);

//---------------------------------------------------------------------------
void CDonation::Format(ostream& ctx, const string_q& fmtIn, void* dataPtr) const {
    if (!m_showing)
        return;

    // EXISTING_CODE
    // EXISTING_CODE

    string_q fmt = (fmtIn.empty() ? expContext().fmtMap["donation_fmt"] : fmtIn);
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
        ctx << getNextChunk(fmt, nextDonationChunk, this);
}

//---------------------------------------------------------------------------
string_q nextDonationChunk(const string_q& fieldIn, const void* dataPtr) {
    if (dataPtr)
        return reinterpret_cast<const CDonation*>(dataPtr)->getValueByName(fieldIn);

    // EXISTING_CODE
    // EXISTING_CODE

    return fldNotFound(fieldIn);
}

//---------------------------------------------------------------------------
string_q CDonation::getValueByName(const string_q& fieldName) const {
    // Give customized code a chance to override first
    string_q ret = nextDonationChunk_custom(fieldName, this);
    if (!ret.empty())
        return ret;

    // EXISTING_CODE
    // EXISTING_CODE

    // Return field values
    switch (tolower(fieldName[0])) {
        case 'a':
            if (fieldName % "amount") {
                return amount;
            }
            break;
        case 'b':
            if (fieldName % "block") {
                return uint_2_Str(block);
            }
            break;
        case 'd':
            if (fieldName % "date") {
                return date;
            }
            if (fieldName % "donor") {
                return addr_2_Str(donor);
            }
            break;
        case 'l':
            if (fieldName % "log_id") {
                return uint_2_Str(log_id);
            }
            break;
        case 'r':
            if (fieldName % "recipient") {
                return addr_2_Str(recipient);
            }
            break;
        case 't':
            if (fieldName % "tx_id") {
                return uint_2_Str(tx_id);
            }
            if (fieldName % "token") {
                return token;
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
bool CDonation::setValueByName(const string_q& fieldNameIn, const string_q& fieldValueIn) {
    string_q fieldName = fieldNameIn;
    string_q fieldValue = fieldValueIn;

    // EXISTING_CODE
    // EXISTING_CODE

    switch (tolower(fieldName[0])) {
        case 'a':
            if (fieldName % "amount") {
                amount = fieldValue;
                return true;
            }
            break;
        case 'b':
            if (fieldName % "block") {
                block = str_2_Uint(fieldValue);
                return true;
            }
            break;
        case 'd':
            if (fieldName % "date") {
                date = fieldValue;
                return true;
            }
            if (fieldName % "donor") {
                donor = str_2_Addr(fieldValue);
                return true;
            }
            break;
        case 'l':
            if (fieldName % "log_id") {
                log_id = str_2_Uint(fieldValue);
                return true;
            }
            break;
        case 'r':
            if (fieldName % "recipient") {
                recipient = str_2_Addr(fieldValue);
                return true;
            }
            break;
        case 't':
            if (fieldName % "tx_id") {
                tx_id = str_2_Uint(fieldValue);
                return true;
            }
            if (fieldName % "token") {
                token = fieldValue;
                return true;
            }
            break;
        default:
            break;
    }
    return false;
}

//---------------------------------------------------------------------------------------------------
void CDonation::finishParse() {
    // EXISTING_CODE
    // EXISTING_CODE
}

//---------------------------------------------------------------------------------------------------
bool CDonation::Serialize(CArchive& archive) {
    if (archive.isWriting())
        return SerializeC(archive);

    // Always read the base class (it will handle its own backLevels if any, then
    // read this object's back level (if any) or the current version.
    CBaseNode::Serialize(archive);
    if (readBackLevel(archive))
        return true;

    // EXISTING_CODE
    // EXISTING_CODE
    archive >> block;
    archive >> tx_id;
    archive >> log_id;
    archive >> date;
    archive >> token;
    archive >> amount;
    archive >> recipient;
    archive >> donor;
    finishParse();
    return true;
}

//---------------------------------------------------------------------------------------------------
bool CDonation::SerializeC(CArchive& archive) const {
    // Writing always write the latest version of the data
    CBaseNode::SerializeC(archive);

    // EXISTING_CODE
    // EXISTING_CODE
    archive << block;
    archive << tx_id;
    archive << log_id;
    archive << date;
    archive << token;
    archive << amount;
    archive << recipient;
    archive << donor;

    return true;
}

//---------------------------------------------------------------------------
CArchive& operator>>(CArchive& archive, CDonationArray& array) {
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
CArchive& operator<<(CArchive& archive, const CDonationArray& array) {
    uint64_t count = array.size();
    archive << count;
    for (size_t i = 0; i < array.size(); i++)
        array[i].SerializeC(archive);
    return archive;
}

//---------------------------------------------------------------------------
void CDonation::registerClass(void) {
    // only do this once
    if (HAS_FIELD(CDonation, "schema"))
        return;

    size_t fieldNum = 1000;
    ADD_FIELD(CDonation, "schema", T_NUMBER, ++fieldNum);
    ADD_FIELD(CDonation, "deleted", T_BOOL, ++fieldNum);
    ADD_FIELD(CDonation, "showing", T_BOOL, ++fieldNum);
    ADD_FIELD(CDonation, "cname", T_TEXT, ++fieldNum);
    ADD_FIELD(CDonation, "block", T_BLOCKNUM, ++fieldNum);
    ADD_FIELD(CDonation, "tx_id", T_BLOCKNUM, ++fieldNum);
    ADD_FIELD(CDonation, "log_id", T_BLOCKNUM, ++fieldNum);
    ADD_FIELD(CDonation, "date", T_TEXT | TS_OMITEMPTY, ++fieldNum);
    ADD_FIELD(CDonation, "token", T_TEXT | TS_OMITEMPTY, ++fieldNum);
    ADD_FIELD(CDonation, "amount", T_TEXT | TS_OMITEMPTY, ++fieldNum);
    ADD_FIELD(CDonation, "recipient", T_ADDRESS | TS_OMITEMPTY, ++fieldNum);
    ADD_FIELD(CDonation, "donor", T_ADDRESS | TS_OMITEMPTY, ++fieldNum);

    // Hide our internal fields, user can turn them on if they like
    HIDE_FIELD(CDonation, "schema");
    HIDE_FIELD(CDonation, "deleted");
    HIDE_FIELD(CDonation, "showing");
    HIDE_FIELD(CDonation, "cname");

    builtIns.push_back(_biCDonation);

    // EXISTING_CODE
    // EXISTING_CODE
}

//---------------------------------------------------------------------------
string_q nextDonationChunk_custom(const string_q& fieldIn, const void* dataPtr) {
    const CDonation* dona = reinterpret_cast<const CDonation*>(dataPtr);
    if (dona) {
        switch (tolower(fieldIn[0])) {
            // EXISTING_CODE
            // EXISTING_CODE
            case 'p':
                // Display only the fields of this node, not it's parent type
                if (fieldIn % "parsed")
                    return nextBasenodeChunk(fieldIn, dona);
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
bool CDonation::readBackLevel(CArchive& archive) {
    bool done = false;
    // EXISTING_CODE
    // EXISTING_CODE
    return done;
}

//-------------------------------------------------------------------------
ostream& operator<<(ostream& os, const CDonation& it) {
    // EXISTING_CODE
    // EXISTING_CODE

    it.Format(os, "", nullptr);
    os << "\n";
    return os;
}

//---------------------------------------------------------------------------
const char* STR_DISPLAY_DONATION = "";

//---------------------------------------------------------------------------
// EXISTING_CODE
// EXISTING_CODE
}  // namespace qblocks
