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
#include "logentry_min.h"
#include "etherlib.h"

namespace qblocks {

//---------------------------------------------------------------------------
IMPLEMENT_NODE(CLogEntry_min, CBaseNode);

//---------------------------------------------------------------------------
extern string_q nextLogentry_MinChunk(const string_q& fieldIn, const void* dataPtr);
static string_q nextLogentry_MinChunk_custom(const string_q& fieldIn, const void* dataPtr);

//---------------------------------------------------------------------------
void CLogEntry_min::Format(ostream& ctx, const string_q& fmtIn, void* dataPtr) const {
    if (!m_showing)
        return;

    // EXISTING_CODE
    // EXISTING_CODE

    string_q fmt = (fmtIn.empty() ? expContext().fmtMap["logentry_min_fmt"] : fmtIn);
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
        ctx << getNextChunk(fmt, nextLogentry_MinChunk, this);
}

//---------------------------------------------------------------------------
string_q nextLogentry_MinChunk(const string_q& fieldIn, const void* dataPtr) {
    if (dataPtr)
        return reinterpret_cast<const CLogEntry_min*>(dataPtr)->getValueByName(fieldIn);

    // EXISTING_CODE
    // EXISTING_CODE

    return fldNotFound(fieldIn);
}

//---------------------------------------------------------------------------
string_q CLogEntry_min::getValueByName(const string_q& fieldName) const {
    // Give customized code a chance to override first
    string_q ret = nextLogentry_MinChunk_custom(fieldName, this);
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
            if (fieldName % "articulatedLog") {
                if (articulatedLog == CFunction())
                    return "{}";
                return articulatedLog.Format();
            }
            break;
        case 'b':
            if (fieldName % "blockNumber") {
                return uint_2_Str(blockNumber);
            }
            break;
        case 'c':
            if (fieldName % "compressedLog") {
                return compressedLog;
            }
            break;
        case 'd':
            if (fieldName % "data") {
                return data;
            }
            break;
        case 'l':
            if (fieldName % "logIndex") {
                return uint_2_Str(logIndex);
            }
            break;
        case 't':
            if (fieldName % "topics" || fieldName % "topicsCnt") {
                size_t cnt = topics.size();
                if (endsWith(toLower(fieldName), "cnt"))
                    return uint_2_Str(cnt);
                if (!cnt)
                    return "";
                string_q retS;
                for (size_t i = 0; i < cnt; i++) {
                    retS += ("\"" + topic_2_Str(topics[i]) + "\"");
                    retS += ((i < cnt - 1) ? ",\n" + indentStr() : "\n");
                }
                return retS;
            }
            if (fieldName % "transactionIndex") {
                return uint_2_Str(transactionIndex);
            }
            break;
        default:
            break;
    }

    // EXISTING_CODE
    // EXISTING_CODE

    string_q s;
    s = toUpper(string_q("articulatedLog")) + "::";
    if (contains(fieldName, s)) {
        string_q f = fieldName;
        replaceAll(f, s, "");
        f = articulatedLog.getValueByName(f);
        return f;
    }

    // Finally, give the parent class a chance
    return CBaseNode::getValueByName(fieldName);
}

//---------------------------------------------------------------------------------------------------
bool CLogEntry_min::setValueByName(const string_q& fieldNameIn, const string_q& fieldValueIn) {
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
            if (fieldName % "articulatedLog") {
                return articulatedLog.parseJson3(fieldValue);
            }
            break;
        case 'b':
            if (fieldName % "blockNumber") {
                blockNumber = str_2_Uint(fieldValue);
                return true;
            }
            break;
        case 'c':
            if (fieldName % "compressedLog") {
                compressedLog = fieldValue;
                return true;
            }
            break;
        case 'd':
            if (fieldName % "data") {
                data = fieldValue;
                return true;
            }
            break;
        case 'l':
            if (fieldName % "logIndex") {
                logIndex = str_2_Uint(fieldValue);
                return true;
            }
            break;
        case 't':
            if (fieldName % "topics") {
                string_q str = fieldValue;
                while (!str.empty()) {
                    topics.push_back(str_2_Topic(nextTokenClear(str, ',')));
                }
                return true;
            }
            if (fieldName % "transactionIndex") {
                transactionIndex = str_2_Uint(fieldValue);
                return true;
            }
            break;
        default:
            break;
    }
    return false;
}

//---------------------------------------------------------------------------------------------------
void CLogEntry_min::finishParse() {
    // EXISTING_CODE
    // EXISTING_CODE
}

//---------------------------------------------------------------------------------------------------
bool CLogEntry_min::Serialize(CArchive& archive) {
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
    archive >> blockNumber;
    archive >> logIndex;
    archive >> topics;
    archive >> data;
    archive >> articulatedLog;
    archive >> compressedLog;
    archive >> transactionIndex;
    finishParse();
    return true;
}

//---------------------------------------------------------------------------------------------------
bool CLogEntry_min::SerializeC(CArchive& archive) const {
    // Writing always write the latest version of the data
    CBaseNode::SerializeC(archive);

    // EXISTING_CODE
    // EXISTING_CODE
    archive << address;
    archive << blockNumber;
    archive << logIndex;
    archive << topics;
    archive << data;
    archive << articulatedLog;
    archive << compressedLog;
    archive << transactionIndex;

    return true;
}

//---------------------------------------------------------------------------
CArchive& operator>>(CArchive& archive, CLogEntry_minArray& array) {
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
CArchive& operator<<(CArchive& archive, const CLogEntry_minArray& array) {
    uint64_t count = array.size();
    archive << count;
    for (size_t i = 0; i < array.size(); i++)
        array[i].SerializeC(archive);
    return archive;
}

//---------------------------------------------------------------------------
void CLogEntry_min::registerClass(void) {
    // only do this once
    if (HAS_FIELD(CLogEntry_min, "schema"))
        return;

    size_t fieldNum = 1000;
    ADD_FIELD(CLogEntry_min, "schema", T_NUMBER, ++fieldNum);
    ADD_FIELD(CLogEntry_min, "deleted", T_BOOL, ++fieldNum);
    ADD_FIELD(CLogEntry_min, "showing", T_BOOL, ++fieldNum);
    ADD_FIELD(CLogEntry_min, "cname", T_TEXT, ++fieldNum);
    ADD_FIELD(CLogEntry_min, "address", T_ADDRESS | TS_OMITEMPTY, ++fieldNum);
    ADD_FIELD(CLogEntry_min, "blockNumber", T_BLOCKNUM, ++fieldNum);
    ADD_FIELD(CLogEntry_min, "logIndex", T_BLOCKNUM, ++fieldNum);
    ADD_FIELD(CLogEntry_min, "topics", T_OBJECT | TS_ARRAY | TS_OMITEMPTY, ++fieldNum);
    ADD_FIELD(CLogEntry_min, "data", T_TEXT | TS_OMITEMPTY, ++fieldNum);
    ADD_OBJECT(CLogEntry_min, "articulatedLog", T_OBJECT | TS_OMITEMPTY, ++fieldNum, GETRUNTIME_CLASS(CFunction));
    ADD_FIELD(CLogEntry_min, "compressedLog", T_TEXT | TS_OMITEMPTY, ++fieldNum);
    ADD_FIELD(CLogEntry_min, "transactionIndex", T_BLOCKNUM, ++fieldNum);

    // Hide our internal fields, user can turn them on if they like
    HIDE_FIELD(CLogEntry_min, "schema");
    HIDE_FIELD(CLogEntry_min, "deleted");
    HIDE_FIELD(CLogEntry_min, "showing");
    HIDE_FIELD(CLogEntry_min, "cname");

    builtIns.push_back(_biCLogEntry_min);

    // EXISTING_CODE
    ADD_FIELD(CLogEntry, "topic0", T_HASH | TS_OMITEMPTY, ++fieldNum);
    ADD_FIELD(CLogEntry, "topic1", T_HASH | TS_OMITEMPTY, ++fieldNum);
    ADD_FIELD(CLogEntry, "topic2", T_HASH | TS_OMITEMPTY, ++fieldNum);
    ADD_FIELD(CLogEntry, "topic3", T_HASH | TS_OMITEMPTY, ++fieldNum);
    HIDE_FIELD(CLogEntry, "topic0");
    HIDE_FIELD(CLogEntry, "topic1");
    HIDE_FIELD(CLogEntry, "topic2");
    HIDE_FIELD(CLogEntry, "topic3");
    // EXISTING_CODE
}

//---------------------------------------------------------------------------
string_q nextLogentry_MinChunk_custom(const string_q& fieldIn, const void* dataPtr) {
    const CLogEntry_min* log = reinterpret_cast<const CLogEntry_min*>(dataPtr);
    if (log) {
        switch (tolower(fieldIn[0])) {
            // EXISTING_CODE
            case 'c':
                if (fieldIn % "compressedLog")
                    return stripWhitespace(log->articulatedLog.compressed());
                break;
            case 't':
                if (fieldIn % "topic0") {
                    return ((log->topics.size() > 0) ? topic_2_Str(log->topics[0]) : "");
                }
                if (fieldIn % "topic1") {
                    return ((log->topics.size() > 1) ? topic_2_Str(log->topics[1]) : "");
                }
                if (fieldIn % "topic2") {
                    return ((log->topics.size() > 2) ? topic_2_Str(log->topics[2]) : "");
                }
                if (fieldIn % "topic3") {
                    return ((log->topics.size() > 3) ? topic_2_Str(log->topics[3]) : "");
                }
                break;
            // EXISTING_CODE
            case 'p':
                // Display only the fields of this node, not it's parent type
                if (fieldIn % "parsed")
                    return nextBasenodeChunk(fieldIn, log);
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
bool CLogEntry_min::readBackLevel(CArchive& archive) {
    bool done = false;
    // EXISTING_CODE
    // EXISTING_CODE
    return done;
}

//---------------------------------------------------------------------------
CArchive& operator<<(CArchive& archive, const CLogEntry_min& log) {
    log.SerializeC(archive);
    return archive;
}

//---------------------------------------------------------------------------
CArchive& operator>>(CArchive& archive, CLogEntry_min& log) {
    log.Serialize(archive);
    return archive;
}

//-------------------------------------------------------------------------
ostream& operator<<(ostream& os, const CLogEntry_min& it) {
    // EXISTING_CODE
    // EXISTING_CODE

    it.Format(os, "", nullptr);
    os << "\n";
    return os;
}

//---------------------------------------------------------------------------
const CBaseNode* CLogEntry_min::getObjectAt(const string_q& fieldName, size_t index) const {
    if (fieldName % "articulatedLog")
        return &articulatedLog;

    return NULL;
}

//---------------------------------------------------------------------------
const string_q CLogEntry_min::getStringAt(const string_q& fieldName, size_t i) const {
    if (fieldName % "topics" && i < topics.size())
        return topic_2_Str(topics[i]);
    return "";
}

//---------------------------------------------------------------------------
const char* STR_DISPLAY_LOGENTRY_MIN =
    "[{BLOCKNUMBER}]\t"
    "[{TRANSACTIONINDEX}]\t"
    "[{LOGINDEX}]\t"
    "[{ADDRESS}]\t"
    "[{TOPIC0}]\t"
    "[{TOPIC1}]\t"
    "[{TOPIC2}]\t"
    "[{TOPIC3}]\t"
    "[{DATA}]\t"
    "[{TYPE}]\t"
    "[{COMPRESSEDLOG}]";

//---------------------------------------------------------------------------
// EXISTING_CODE
CLogEntry_min& CLogEntry_min::operator=(const CLogEntry& lo) {
    clear();
    CBaseNode::duplicate(lo);

    address = lo.address;
    logIndex = lo.logIndex;
    topics = lo.topics;
    data = lo.data;
    articulatedLog = lo.articulatedLog;
    compressedLog = lo.compressedLog;
    return *this;
}
// EXISTING_CODE
}  // namespace qblocks
