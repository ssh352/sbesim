/* Generated SBE (Simple Binary Encoding) message codec */
#ifndef _ILINKBINARY_EXECUTIONREPORTSTATUS532_H_
#define _ILINKBINARY_EXECUTIONREPORTSTATUS532_H_

#if defined(SBE_HAVE_CMATH)
/* cmath needed for std::numeric_limits<double>::quiet_NaN() */
#  include <cmath>
#  define SBE_FLOAT_NAN std::numeric_limits<float>::quiet_NaN()
#  define SBE_DOUBLE_NAN std::numeric_limits<double>::quiet_NaN()
#else
/* math.h needed for NAN */
#  include <math.h>
#  define SBE_FLOAT_NAN NAN
#  define SBE_DOUBLE_NAN NAN
#endif

#if __cplusplus >= 201103L
#  define SBE_CONSTEXPR constexpr
#  define SBE_NOEXCEPT noexcept
#else
#  define SBE_CONSTEXPR
#  define SBE_NOEXCEPT
#endif

#if __cplusplus >= 201402L
#  define SBE_CONSTEXPR_14 constexpr
#else
#  define SBE_CONSTEXPR_14
#endif

#if __cplusplus >= 201703L
#  include <string_view>
#  define SBE_NODISCARD [[nodiscard]]
#else
#  define SBE_NODISCARD
#endif

#if !defined(__STDC_LIMIT_MACROS)
#  define __STDC_LIMIT_MACROS 1
#endif

#include <cstdint>
#include <cstring>
#include <iomanip>
#include <limits>
#include <ostream>
#include <stdexcept>
#include <sstream>
#include <string>
#include <vector>

#if defined(WIN32) || defined(_WIN32)
#  define SBE_BIG_ENDIAN_ENCODE_16(v) _byteswap_ushort(v)
#  define SBE_BIG_ENDIAN_ENCODE_32(v) _byteswap_ulong(v)
#  define SBE_BIG_ENDIAN_ENCODE_64(v) _byteswap_uint64(v)
#  define SBE_LITTLE_ENDIAN_ENCODE_16(v) (v)
#  define SBE_LITTLE_ENDIAN_ENCODE_32(v) (v)
#  define SBE_LITTLE_ENDIAN_ENCODE_64(v) (v)
#elif __BYTE_ORDER__ == __ORDER_LITTLE_ENDIAN__
#  define SBE_BIG_ENDIAN_ENCODE_16(v) __builtin_bswap16(v)
#  define SBE_BIG_ENDIAN_ENCODE_32(v) __builtin_bswap32(v)
#  define SBE_BIG_ENDIAN_ENCODE_64(v) __builtin_bswap64(v)
#  define SBE_LITTLE_ENDIAN_ENCODE_16(v) (v)
#  define SBE_LITTLE_ENDIAN_ENCODE_32(v) (v)
#  define SBE_LITTLE_ENDIAN_ENCODE_64(v) (v)
#elif __BYTE_ORDER__ == __ORDER_BIG_ENDIAN__
#  define SBE_LITTLE_ENDIAN_ENCODE_16(v) __builtin_bswap16(v)
#  define SBE_LITTLE_ENDIAN_ENCODE_32(v) __builtin_bswap32(v)
#  define SBE_LITTLE_ENDIAN_ENCODE_64(v) __builtin_bswap64(v)
#  define SBE_BIG_ENDIAN_ENCODE_16(v) (v)
#  define SBE_BIG_ENDIAN_ENCODE_32(v) (v)
#  define SBE_BIG_ENDIAN_ENCODE_64(v) (v)
#else
#  error "Byte Ordering of platform not determined. Set __BYTE_ORDER__ manually before including this file."
#endif

#if defined(SBE_NO_BOUNDS_CHECK)
#  define SBE_BOUNDS_CHECK_EXPECT(exp,c) (false)
#elif defined(_MSC_VER)
#  define SBE_BOUNDS_CHECK_EXPECT(exp,c) (exp)
#else
#  define SBE_BOUNDS_CHECK_EXPECT(exp,c) (__builtin_expect(exp,c))
#endif

#define SBE_NULLVALUE_INT8 (std::numeric_limits<std::int8_t>::min)()
#define SBE_NULLVALUE_INT16 (std::numeric_limits<std::int16_t>::min)()
#define SBE_NULLVALUE_INT32 (std::numeric_limits<std::int32_t>::min)()
#define SBE_NULLVALUE_INT64 (std::numeric_limits<std::int64_t>::min)()
#define SBE_NULLVALUE_UINT8 (std::numeric_limits<std::uint8_t>::max)()
#define SBE_NULLVALUE_UINT16 (std::numeric_limits<std::uint16_t>::max)()
#define SBE_NULLVALUE_UINT32 (std::numeric_limits<std::uint32_t>::max)()
#define SBE_NULLVALUE_UINT64 (std::numeric_limits<std::uint64_t>::max)()


#include "MaturityMonthYear.h"
#include "Decimal64NULL.h"
#include "OrderTypeReq.h"
#include "BooleanNULL.h"
#include "AvgPxInd.h"
#include "OrderEventType.h"
#include "CmtaGiveUpCD.h"
#include "PRICENULL9.h"
#include "BooleanFlag.h"
#include "OrderType.h"
#include "QuoteCxlStatus.h"
#include "TimeInForce.h"
#include "PRICE9.h"
#include "MassStatusOrdTyp.h"
#include "ExecInst.h"
#include "Decimal32NULL.h"
#include "SplitMsg.h"
#include "SideReq.h"
#include "MassActionResponse.h"
#include "CustOrderCapacity.h"
#include "SideTimeInForce.h"
#include "SLEDS.h"
#include "ReqResult.h"
#include "ClearingAcctType.h"
#include "KeepAliveLapsed.h"
#include "ExecAckStatus.h"
#include "DATA.h"
#include "OFMOverrideReq.h"
#include "SideNULL.h"
#include "GroupSize.h"
#include "ListUpdAct.h"
#include "FTI.h"
#include "ExecReason.h"
#include "MassStatusReqTyp.h"
#include "DKReason.h"
#include "SecRspTyp.h"
#include "MassActionOrdTyp.h"
#include "QuoteAckStatus.h"
#include "ExpCycle.h"
#include "MassCancelTIF.h"
#include "OrderStatus.h"
#include "ShortSaleType.h"
#include "PartyDetailRole.h"
#include "MassActionScope.h"
#include "RFQSide.h"
#include "OrdStatusTrdCxl.h"
#include "TradeAddendum.h"
#include "MessageHeader.h"
#include "ManualOrdInd.h"
#include "QuoteCxlTyp.h"
#include "MassCxlReqTyp.h"
#include "ExecTypTrdCxl.h"
#include "ManualOrdIndReq.h"
#include "CustOrdHandlInst.h"
#include "OrdStatusTrd.h"
#include "MassStatusTIF.h"
#include "SMPI.h"
#include "ExecMode.h"
#include "QuoteTyp.h"

namespace iLinkBinary {

class ExecutionReportStatus532
{
private:
    char *m_buffer = nullptr;
    std::uint64_t m_bufferLength = 0;
    std::uint64_t m_offset = 0;
    std::uint64_t m_position = 0;
    std::uint64_t m_actingVersion = 0;

    inline std::uint64_t *sbePositionPtr() SBE_NOEXCEPT
    {
        return &m_position;
    }

public:
    enum MetaAttribute
    {
        EPOCH, TIME_UNIT, SEMANTIC_TYPE, PRESENCE
    };

    union sbe_float_as_uint_u
    {
        float fp_value;
        std::uint32_t uint_value;
    };

    union sbe_double_as_uint_u
    {
        double fp_value;
        std::uint64_t uint_value;
    };

    ExecutionReportStatus532() = default;

    ExecutionReportStatus532(
        char *buffer,
        const std::uint64_t offset,
        const std::uint64_t bufferLength,
        const std::uint64_t actingBlockLength,
        const std::uint64_t actingVersion) :
        m_buffer(buffer),
        m_bufferLength(bufferLength),
        m_offset(offset),
        m_position(sbeCheckPosition(offset + actingBlockLength)),
        m_actingVersion(actingVersion)
    {
    }

    ExecutionReportStatus532(char *buffer, const std::uint64_t bufferLength) :
        ExecutionReportStatus532(buffer, 0, bufferLength, sbeBlockLength(), sbeSchemaVersion())
    {
    }

    ExecutionReportStatus532(
        char *buffer,
        const std::uint64_t bufferLength,
        const std::uint64_t actingBlockLength,
        const std::uint64_t actingVersion) :
        ExecutionReportStatus532(buffer, 0, bufferLength, actingBlockLength, actingVersion)
    {
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint16_t sbeBlockLength() SBE_NOEXCEPT
    {
        return (std::uint16_t)480;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint16_t sbeBlockAndHeaderLength() SBE_NOEXCEPT
    {
        return MessageHeader::encodedLength() + sbeBlockLength();
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint16_t sbeTemplateId() SBE_NOEXCEPT
    {
        return (std::uint16_t)532;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint16_t sbeSchemaId() SBE_NOEXCEPT
    {
        return (std::uint16_t)8;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint16_t sbeSchemaVersion() SBE_NOEXCEPT
    {
        return (std::uint16_t)5;
    }

    SBE_NODISCARD static SBE_CONSTEXPR const char * sbeSemanticType() SBE_NOEXCEPT
    {
        return "8";
    }

    SBE_NODISCARD std::uint64_t offset() const SBE_NOEXCEPT
    {
        return m_offset;
    }

    ExecutionReportStatus532 &wrapForEncode(char *buffer, const std::uint64_t offset, const std::uint64_t bufferLength)
    {
        return *this = ExecutionReportStatus532(buffer, offset, bufferLength, sbeBlockLength(), sbeSchemaVersion());
    }

    ExecutionReportStatus532 &wrapAndApplyHeader(char *buffer, const std::uint64_t offset, const std::uint64_t bufferLength)
    {
        MessageHeader hdr(buffer, offset, bufferLength, sbeSchemaVersion());

        hdr
            .blockLength(sbeBlockLength())
            .templateId(sbeTemplateId())
            .schemaId(sbeSchemaId())
            .version(sbeSchemaVersion());

        return *this = ExecutionReportStatus532(
            buffer,
            offset + MessageHeader::encodedLength(),
            bufferLength,
            sbeBlockLength(),
            sbeSchemaVersion());
    }

    ExecutionReportStatus532 &wrapForDecode(
        char *buffer,
        const std::uint64_t offset,
        const std::uint64_t actingBlockLength,
        const std::uint64_t actingVersion,
        const std::uint64_t bufferLength)
    {
        return *this = ExecutionReportStatus532(buffer, offset, bufferLength, actingBlockLength, actingVersion);
    }

    SBE_NODISCARD std::uint64_t sbePosition() const SBE_NOEXCEPT
    {
        return m_position;
    }

    // NOLINTNEXTLINE(readability-convert-member-functions-to-static)
    std::uint64_t sbeCheckPosition(const std::uint64_t position)
    {
        if (SBE_BOUNDS_CHECK_EXPECT((position > m_bufferLength), false))
        {
            throw std::runtime_error("buffer too short [E100]");
        }
        return position;
    }

    void sbePosition(const std::uint64_t position)
    {
        m_position = sbeCheckPosition(position);
    }

    SBE_NODISCARD std::uint64_t encodedLength() const SBE_NOEXCEPT
    {
        return sbePosition() - m_offset;
    }

    SBE_NODISCARD std::uint64_t decodeLength() const
    {
        ExecutionReportStatus532 skipper(m_buffer, m_offset,
            m_bufferLength, sbeBlockLength(), m_actingVersion);
        skipper.skip();
        return skipper.encodedLength();
    }

    SBE_NODISCARD const char * buffer() const SBE_NOEXCEPT
    {
        return m_buffer;
    }

    SBE_NODISCARD char * buffer() SBE_NOEXCEPT
    {
        return m_buffer;
    }

    SBE_NODISCARD std::uint64_t bufferLength() const SBE_NOEXCEPT
    {
        return m_bufferLength;
    }

    SBE_NODISCARD std::uint64_t actingVersion() const SBE_NOEXCEPT
    {
        return m_actingVersion;
    }

    SBE_NODISCARD static const char * SeqNumMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::SEMANTIC_TYPE: return "int";
            case MetaAttribute::PRESENCE: return "required";
            default: return "";
        }
    }

    static SBE_CONSTEXPR std::uint16_t seqNumId() SBE_NOEXCEPT
    {
        return 9726;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t seqNumSinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool seqNumInActingVersion() SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= seqNumSinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t seqNumEncodingOffset() SBE_NOEXCEPT
    {
        return 0;
    }

    static SBE_CONSTEXPR std::uint32_t seqNumNullValue() SBE_NOEXCEPT
    {
        return SBE_NULLVALUE_UINT32;
    }

    static SBE_CONSTEXPR std::uint32_t seqNumMinValue() SBE_NOEXCEPT
    {
        return 0;
    }

    static SBE_CONSTEXPR std::uint32_t seqNumMaxValue() SBE_NOEXCEPT
    {
        return 4294967294;
    }

    static SBE_CONSTEXPR std::size_t seqNumEncodingLength() SBE_NOEXCEPT
    {
        return 4;
    }

    SBE_NODISCARD std::uint32_t seqNum() const SBE_NOEXCEPT
    {
        std::uint32_t val;
        std::memcpy(&val, m_buffer + m_offset + 0, sizeof(std::uint32_t));
        return SBE_LITTLE_ENDIAN_ENCODE_32(val);
    }

    ExecutionReportStatus532 &seqNum(const std::uint32_t value) SBE_NOEXCEPT
    {
        std::uint32_t val = SBE_LITTLE_ENDIAN_ENCODE_32(value);
        std::memcpy(m_buffer + m_offset + 0, &val, sizeof(std::uint32_t));
        return *this;
    }

    SBE_NODISCARD static const char * UUIDMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::SEMANTIC_TYPE: return "int";
            case MetaAttribute::PRESENCE: return "required";
            default: return "";
        }
    }

    static SBE_CONSTEXPR std::uint16_t uUIDId() SBE_NOEXCEPT
    {
        return 39001;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t uUIDSinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool uUIDInActingVersion() SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= uUIDSinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t uUIDEncodingOffset() SBE_NOEXCEPT
    {
        return 4;
    }

    static SBE_CONSTEXPR std::uint64_t uUIDNullValue() SBE_NOEXCEPT
    {
        return SBE_NULLVALUE_UINT64;
    }

    static SBE_CONSTEXPR std::uint64_t uUIDMinValue() SBE_NOEXCEPT
    {
        return 0x0L;
    }

    static SBE_CONSTEXPR std::uint64_t uUIDMaxValue() SBE_NOEXCEPT
    {
        return 0xfffffffffffffffeL;
    }

    static SBE_CONSTEXPR std::size_t uUIDEncodingLength() SBE_NOEXCEPT
    {
        return 8;
    }

    SBE_NODISCARD std::uint64_t uUID() const SBE_NOEXCEPT
    {
        std::uint64_t val;
        std::memcpy(&val, m_buffer + m_offset + 4, sizeof(std::uint64_t));
        return SBE_LITTLE_ENDIAN_ENCODE_64(val);
    }

    ExecutionReportStatus532 &uUID(const std::uint64_t value) SBE_NOEXCEPT
    {
        std::uint64_t val = SBE_LITTLE_ENDIAN_ENCODE_64(value);
        std::memcpy(m_buffer + m_offset + 4, &val, sizeof(std::uint64_t));
        return *this;
    }

    SBE_NODISCARD static const char * TextMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::SEMANTIC_TYPE: return "String";
            case MetaAttribute::PRESENCE: return "optional";
            default: return "";
        }
    }

    static SBE_CONSTEXPR std::uint16_t textId() SBE_NOEXCEPT
    {
        return 58;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t textSinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool textInActingVersion() SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= textSinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t textEncodingOffset() SBE_NOEXCEPT
    {
        return 12;
    }

    static SBE_CONSTEXPR char textNullValue() SBE_NOEXCEPT
    {
        return (char)0;
    }

    static SBE_CONSTEXPR char textMinValue() SBE_NOEXCEPT
    {
        return (char)32;
    }

    static SBE_CONSTEXPR char textMaxValue() SBE_NOEXCEPT
    {
        return (char)126;
    }

    static SBE_CONSTEXPR std::size_t textEncodingLength() SBE_NOEXCEPT
    {
        return 256;
    }

    static SBE_CONSTEXPR std::uint64_t textLength() SBE_NOEXCEPT
    {
        return 256;
    }

    SBE_NODISCARD const char *text() const SBE_NOEXCEPT
    {
        return m_buffer + m_offset + 12;
    }

    SBE_NODISCARD char *text() SBE_NOEXCEPT
    {
        return m_buffer + m_offset + 12;
    }

    SBE_NODISCARD char text(const std::uint64_t index) const
    {
        if (index >= 256)
        {
            throw std::runtime_error("index out of range for text [E104]");
        }

        char val;
        std::memcpy(&val, m_buffer + m_offset + 12 + (index * 1), sizeof(char));
        return (val);
    }

    ExecutionReportStatus532 &text(const std::uint64_t index, const char value)
    {
        if (index >= 256)
        {
            throw std::runtime_error("index out of range for text [E105]");
        }

        char val = (value);
        std::memcpy(m_buffer + m_offset + 12 + (index * 1), &val, sizeof(char));
        return *this;
    }

    std::uint64_t getText(char *const dst, const std::uint64_t length) const
    {
        if (length > 256)
        {
            throw std::runtime_error("length too large for getText [E106]");
        }

        std::memcpy(dst, m_buffer + m_offset + 12, sizeof(char) * static_cast<size_t>(length));
        return length;
    }

    ExecutionReportStatus532 &putText(const char *const src) SBE_NOEXCEPT
    {
        std::memcpy(m_buffer + m_offset + 12, src, sizeof(char) * 256);
        return *this;
    }

    SBE_NODISCARD std::string getTextAsString() const
    {
        const char *buffer = m_buffer + m_offset + 12;
        size_t length = 0;

        for (; length < 256 && *(buffer + length) != '\0'; ++length);
        std::string result(buffer, length);

        return result;
    }

    std::string getTextAsJsonEscapedString()
    {
        std::ostringstream oss;
        std::string s = getTextAsString();

        for (const auto c : s)
        {
            switch (c)
            {
                case '"': oss << "\\\""; break;
                case '\\': oss << "\\\\"; break;
                case '\b': oss << "\\b"; break;
                case '\f': oss << "\\f"; break;
                case '\n': oss << "\\n"; break;
                case '\r': oss << "\\r"; break;
                case '\t': oss << "\\t"; break;

                default:
                    if ('\x00' <= c && c <= '\x1f')
                    {
                        oss << "\\u" << std::hex << std::setw(4)
                            << std::setfill('0') << (int)(c);
                    }
                    else
                    {
                        oss << c;
                    }
            }
        }

        return oss.str();
    }

    #if __cplusplus >= 201703L
    SBE_NODISCARD std::string_view getTextAsStringView() const SBE_NOEXCEPT
    {
        const char *buffer = m_buffer + m_offset + 12;
        size_t length = 0;

        for (; length < 256 && *(buffer + length) != '\0'; ++length);
        std::string_view result(buffer, length);

        return result;
    }
    #endif

    #if __cplusplus >= 201703L
    ExecutionReportStatus532 &putText(const std::string_view str)
    {
        const size_t srcLength = str.length();
        if (srcLength > 256)
        {
            throw std::runtime_error("string too large for putText [E106]");
        }

        std::memcpy(m_buffer + m_offset + 12, str.data(), srcLength);
        for (size_t start = srcLength; start < 256; ++start)
        {
            m_buffer[m_offset + 12 + start] = 0;
        }

        return *this;
    }
    #else
    ExecutionReportStatus532 &putText(const std::string& str)
    {
        const size_t srcLength = str.length();
        if (srcLength > 256)
        {
            throw std::runtime_error("string too large for putText [E106]");
        }

        std::memcpy(m_buffer + m_offset + 12, str.c_str(), srcLength);
        for (size_t start = srcLength; start < 256; ++start)
        {
            m_buffer[m_offset + 12 + start] = 0;
        }

        return *this;
    }
    #endif

    SBE_NODISCARD static const char * ExecIDMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::SEMANTIC_TYPE: return "String";
            case MetaAttribute::PRESENCE: return "required";
            default: return "";
        }
    }

    static SBE_CONSTEXPR std::uint16_t execIDId() SBE_NOEXCEPT
    {
        return 17;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t execIDSinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool execIDInActingVersion() SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= execIDSinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t execIDEncodingOffset() SBE_NOEXCEPT
    {
        return 268;
    }

    static SBE_CONSTEXPR char execIDNullValue() SBE_NOEXCEPT
    {
        return (char)0;
    }

    static SBE_CONSTEXPR char execIDMinValue() SBE_NOEXCEPT
    {
        return (char)32;
    }

    static SBE_CONSTEXPR char execIDMaxValue() SBE_NOEXCEPT
    {
        return (char)126;
    }

    static SBE_CONSTEXPR std::size_t execIDEncodingLength() SBE_NOEXCEPT
    {
        return 40;
    }

    static SBE_CONSTEXPR std::uint64_t execIDLength() SBE_NOEXCEPT
    {
        return 40;
    }

    SBE_NODISCARD const char *execID() const SBE_NOEXCEPT
    {
        return m_buffer + m_offset + 268;
    }

    SBE_NODISCARD char *execID() SBE_NOEXCEPT
    {
        return m_buffer + m_offset + 268;
    }

    SBE_NODISCARD char execID(const std::uint64_t index) const
    {
        if (index >= 40)
        {
            throw std::runtime_error("index out of range for execID [E104]");
        }

        char val;
        std::memcpy(&val, m_buffer + m_offset + 268 + (index * 1), sizeof(char));
        return (val);
    }

    ExecutionReportStatus532 &execID(const std::uint64_t index, const char value)
    {
        if (index >= 40)
        {
            throw std::runtime_error("index out of range for execID [E105]");
        }

        char val = (value);
        std::memcpy(m_buffer + m_offset + 268 + (index * 1), &val, sizeof(char));
        return *this;
    }

    std::uint64_t getExecID(char *const dst, const std::uint64_t length) const
    {
        if (length > 40)
        {
            throw std::runtime_error("length too large for getExecID [E106]");
        }

        std::memcpy(dst, m_buffer + m_offset + 268, sizeof(char) * static_cast<size_t>(length));
        return length;
    }

    ExecutionReportStatus532 &putExecID(const char *const src) SBE_NOEXCEPT
    {
        std::memcpy(m_buffer + m_offset + 268, src, sizeof(char) * 40);
        return *this;
    }

    SBE_NODISCARD std::string getExecIDAsString() const
    {
        const char *buffer = m_buffer + m_offset + 268;
        size_t length = 0;

        for (; length < 40 && *(buffer + length) != '\0'; ++length);
        std::string result(buffer, length);

        return result;
    }

    std::string getExecIDAsJsonEscapedString()
    {
        std::ostringstream oss;
        std::string s = getExecIDAsString();

        for (const auto c : s)
        {
            switch (c)
            {
                case '"': oss << "\\\""; break;
                case '\\': oss << "\\\\"; break;
                case '\b': oss << "\\b"; break;
                case '\f': oss << "\\f"; break;
                case '\n': oss << "\\n"; break;
                case '\r': oss << "\\r"; break;
                case '\t': oss << "\\t"; break;

                default:
                    if ('\x00' <= c && c <= '\x1f')
                    {
                        oss << "\\u" << std::hex << std::setw(4)
                            << std::setfill('0') << (int)(c);
                    }
                    else
                    {
                        oss << c;
                    }
            }
        }

        return oss.str();
    }

    #if __cplusplus >= 201703L
    SBE_NODISCARD std::string_view getExecIDAsStringView() const SBE_NOEXCEPT
    {
        const char *buffer = m_buffer + m_offset + 268;
        size_t length = 0;

        for (; length < 40 && *(buffer + length) != '\0'; ++length);
        std::string_view result(buffer, length);

        return result;
    }
    #endif

    #if __cplusplus >= 201703L
    ExecutionReportStatus532 &putExecID(const std::string_view str)
    {
        const size_t srcLength = str.length();
        if (srcLength > 40)
        {
            throw std::runtime_error("string too large for putExecID [E106]");
        }

        std::memcpy(m_buffer + m_offset + 268, str.data(), srcLength);
        for (size_t start = srcLength; start < 40; ++start)
        {
            m_buffer[m_offset + 268 + start] = 0;
        }

        return *this;
    }
    #else
    ExecutionReportStatus532 &putExecID(const std::string& str)
    {
        const size_t srcLength = str.length();
        if (srcLength > 40)
        {
            throw std::runtime_error("string too large for putExecID [E106]");
        }

        std::memcpy(m_buffer + m_offset + 268, str.c_str(), srcLength);
        for (size_t start = srcLength; start < 40; ++start)
        {
            m_buffer[m_offset + 268 + start] = 0;
        }

        return *this;
    }
    #endif

    SBE_NODISCARD static const char * SenderIDMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::SEMANTIC_TYPE: return "String";
            case MetaAttribute::PRESENCE: return "required";
            default: return "";
        }
    }

    static SBE_CONSTEXPR std::uint16_t senderIDId() SBE_NOEXCEPT
    {
        return 5392;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t senderIDSinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool senderIDInActingVersion() SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= senderIDSinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t senderIDEncodingOffset() SBE_NOEXCEPT
    {
        return 308;
    }

    static SBE_CONSTEXPR char senderIDNullValue() SBE_NOEXCEPT
    {
        return (char)0;
    }

    static SBE_CONSTEXPR char senderIDMinValue() SBE_NOEXCEPT
    {
        return (char)32;
    }

    static SBE_CONSTEXPR char senderIDMaxValue() SBE_NOEXCEPT
    {
        return (char)126;
    }

    static SBE_CONSTEXPR std::size_t senderIDEncodingLength() SBE_NOEXCEPT
    {
        return 20;
    }

    static SBE_CONSTEXPR std::uint64_t senderIDLength() SBE_NOEXCEPT
    {
        return 20;
    }

    SBE_NODISCARD const char *senderID() const SBE_NOEXCEPT
    {
        return m_buffer + m_offset + 308;
    }

    SBE_NODISCARD char *senderID() SBE_NOEXCEPT
    {
        return m_buffer + m_offset + 308;
    }

    SBE_NODISCARD char senderID(const std::uint64_t index) const
    {
        if (index >= 20)
        {
            throw std::runtime_error("index out of range for senderID [E104]");
        }

        char val;
        std::memcpy(&val, m_buffer + m_offset + 308 + (index * 1), sizeof(char));
        return (val);
    }

    ExecutionReportStatus532 &senderID(const std::uint64_t index, const char value)
    {
        if (index >= 20)
        {
            throw std::runtime_error("index out of range for senderID [E105]");
        }

        char val = (value);
        std::memcpy(m_buffer + m_offset + 308 + (index * 1), &val, sizeof(char));
        return *this;
    }

    std::uint64_t getSenderID(char *const dst, const std::uint64_t length) const
    {
        if (length > 20)
        {
            throw std::runtime_error("length too large for getSenderID [E106]");
        }

        std::memcpy(dst, m_buffer + m_offset + 308, sizeof(char) * static_cast<size_t>(length));
        return length;
    }

    ExecutionReportStatus532 &putSenderID(const char *const src) SBE_NOEXCEPT
    {
        std::memcpy(m_buffer + m_offset + 308, src, sizeof(char) * 20);
        return *this;
    }

    SBE_NODISCARD std::string getSenderIDAsString() const
    {
        const char *buffer = m_buffer + m_offset + 308;
        size_t length = 0;

        for (; length < 20 && *(buffer + length) != '\0'; ++length);
        std::string result(buffer, length);

        return result;
    }

    std::string getSenderIDAsJsonEscapedString()
    {
        std::ostringstream oss;
        std::string s = getSenderIDAsString();

        for (const auto c : s)
        {
            switch (c)
            {
                case '"': oss << "\\\""; break;
                case '\\': oss << "\\\\"; break;
                case '\b': oss << "\\b"; break;
                case '\f': oss << "\\f"; break;
                case '\n': oss << "\\n"; break;
                case '\r': oss << "\\r"; break;
                case '\t': oss << "\\t"; break;

                default:
                    if ('\x00' <= c && c <= '\x1f')
                    {
                        oss << "\\u" << std::hex << std::setw(4)
                            << std::setfill('0') << (int)(c);
                    }
                    else
                    {
                        oss << c;
                    }
            }
        }

        return oss.str();
    }

    #if __cplusplus >= 201703L
    SBE_NODISCARD std::string_view getSenderIDAsStringView() const SBE_NOEXCEPT
    {
        const char *buffer = m_buffer + m_offset + 308;
        size_t length = 0;

        for (; length < 20 && *(buffer + length) != '\0'; ++length);
        std::string_view result(buffer, length);

        return result;
    }
    #endif

    #if __cplusplus >= 201703L
    ExecutionReportStatus532 &putSenderID(const std::string_view str)
    {
        const size_t srcLength = str.length();
        if (srcLength > 20)
        {
            throw std::runtime_error("string too large for putSenderID [E106]");
        }

        std::memcpy(m_buffer + m_offset + 308, str.data(), srcLength);
        for (size_t start = srcLength; start < 20; ++start)
        {
            m_buffer[m_offset + 308 + start] = 0;
        }

        return *this;
    }
    #else
    ExecutionReportStatus532 &putSenderID(const std::string& str)
    {
        const size_t srcLength = str.length();
        if (srcLength > 20)
        {
            throw std::runtime_error("string too large for putSenderID [E106]");
        }

        std::memcpy(m_buffer + m_offset + 308, str.c_str(), srcLength);
        for (size_t start = srcLength; start < 20; ++start)
        {
            m_buffer[m_offset + 308 + start] = 0;
        }

        return *this;
    }
    #endif

    SBE_NODISCARD static const char * ClOrdIDMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::SEMANTIC_TYPE: return "String";
            case MetaAttribute::PRESENCE: return "required";
            default: return "";
        }
    }

    static SBE_CONSTEXPR std::uint16_t clOrdIDId() SBE_NOEXCEPT
    {
        return 11;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t clOrdIDSinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool clOrdIDInActingVersion() SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= clOrdIDSinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t clOrdIDEncodingOffset() SBE_NOEXCEPT
    {
        return 328;
    }

    static SBE_CONSTEXPR char clOrdIDNullValue() SBE_NOEXCEPT
    {
        return (char)0;
    }

    static SBE_CONSTEXPR char clOrdIDMinValue() SBE_NOEXCEPT
    {
        return (char)32;
    }

    static SBE_CONSTEXPR char clOrdIDMaxValue() SBE_NOEXCEPT
    {
        return (char)126;
    }

    static SBE_CONSTEXPR std::size_t clOrdIDEncodingLength() SBE_NOEXCEPT
    {
        return 20;
    }

    static SBE_CONSTEXPR std::uint64_t clOrdIDLength() SBE_NOEXCEPT
    {
        return 20;
    }

    SBE_NODISCARD const char *clOrdID() const SBE_NOEXCEPT
    {
        return m_buffer + m_offset + 328;
    }

    SBE_NODISCARD char *clOrdID() SBE_NOEXCEPT
    {
        return m_buffer + m_offset + 328;
    }

    SBE_NODISCARD char clOrdID(const std::uint64_t index) const
    {
        if (index >= 20)
        {
            throw std::runtime_error("index out of range for clOrdID [E104]");
        }

        char val;
        std::memcpy(&val, m_buffer + m_offset + 328 + (index * 1), sizeof(char));
        return (val);
    }

    ExecutionReportStatus532 &clOrdID(const std::uint64_t index, const char value)
    {
        if (index >= 20)
        {
            throw std::runtime_error("index out of range for clOrdID [E105]");
        }

        char val = (value);
        std::memcpy(m_buffer + m_offset + 328 + (index * 1), &val, sizeof(char));
        return *this;
    }

    std::uint64_t getClOrdID(char *const dst, const std::uint64_t length) const
    {
        if (length > 20)
        {
            throw std::runtime_error("length too large for getClOrdID [E106]");
        }

        std::memcpy(dst, m_buffer + m_offset + 328, sizeof(char) * static_cast<size_t>(length));
        return length;
    }

    ExecutionReportStatus532 &putClOrdID(const char *const src) SBE_NOEXCEPT
    {
        std::memcpy(m_buffer + m_offset + 328, src, sizeof(char) * 20);
        return *this;
    }

    SBE_NODISCARD std::string getClOrdIDAsString() const
    {
        const char *buffer = m_buffer + m_offset + 328;
        size_t length = 0;

        for (; length < 20 && *(buffer + length) != '\0'; ++length);
        std::string result(buffer, length);

        return result;
    }

    std::string getClOrdIDAsJsonEscapedString()
    {
        std::ostringstream oss;
        std::string s = getClOrdIDAsString();

        for (const auto c : s)
        {
            switch (c)
            {
                case '"': oss << "\\\""; break;
                case '\\': oss << "\\\\"; break;
                case '\b': oss << "\\b"; break;
                case '\f': oss << "\\f"; break;
                case '\n': oss << "\\n"; break;
                case '\r': oss << "\\r"; break;
                case '\t': oss << "\\t"; break;

                default:
                    if ('\x00' <= c && c <= '\x1f')
                    {
                        oss << "\\u" << std::hex << std::setw(4)
                            << std::setfill('0') << (int)(c);
                    }
                    else
                    {
                        oss << c;
                    }
            }
        }

        return oss.str();
    }

    #if __cplusplus >= 201703L
    SBE_NODISCARD std::string_view getClOrdIDAsStringView() const SBE_NOEXCEPT
    {
        const char *buffer = m_buffer + m_offset + 328;
        size_t length = 0;

        for (; length < 20 && *(buffer + length) != '\0'; ++length);
        std::string_view result(buffer, length);

        return result;
    }
    #endif

    #if __cplusplus >= 201703L
    ExecutionReportStatus532 &putClOrdID(const std::string_view str)
    {
        const size_t srcLength = str.length();
        if (srcLength > 20)
        {
            throw std::runtime_error("string too large for putClOrdID [E106]");
        }

        std::memcpy(m_buffer + m_offset + 328, str.data(), srcLength);
        for (size_t start = srcLength; start < 20; ++start)
        {
            m_buffer[m_offset + 328 + start] = 0;
        }

        return *this;
    }
    #else
    ExecutionReportStatus532 &putClOrdID(const std::string& str)
    {
        const size_t srcLength = str.length();
        if (srcLength > 20)
        {
            throw std::runtime_error("string too large for putClOrdID [E106]");
        }

        std::memcpy(m_buffer + m_offset + 328, str.c_str(), srcLength);
        for (size_t start = srcLength; start < 20; ++start)
        {
            m_buffer[m_offset + 328 + start] = 0;
        }

        return *this;
    }
    #endif

    SBE_NODISCARD static const char * PartyDetailsListReqIDMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::SEMANTIC_TYPE: return "int";
            case MetaAttribute::PRESENCE: return "required";
            default: return "";
        }
    }

    static SBE_CONSTEXPR std::uint16_t partyDetailsListReqIDId() SBE_NOEXCEPT
    {
        return 1505;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t partyDetailsListReqIDSinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool partyDetailsListReqIDInActingVersion() SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= partyDetailsListReqIDSinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t partyDetailsListReqIDEncodingOffset() SBE_NOEXCEPT
    {
        return 348;
    }

    static SBE_CONSTEXPR std::uint64_t partyDetailsListReqIDNullValue() SBE_NOEXCEPT
    {
        return SBE_NULLVALUE_UINT64;
    }

    static SBE_CONSTEXPR std::uint64_t partyDetailsListReqIDMinValue() SBE_NOEXCEPT
    {
        return 0x0L;
    }

    static SBE_CONSTEXPR std::uint64_t partyDetailsListReqIDMaxValue() SBE_NOEXCEPT
    {
        return 0xfffffffffffffffeL;
    }

    static SBE_CONSTEXPR std::size_t partyDetailsListReqIDEncodingLength() SBE_NOEXCEPT
    {
        return 8;
    }

    SBE_NODISCARD std::uint64_t partyDetailsListReqID() const SBE_NOEXCEPT
    {
        std::uint64_t val;
        std::memcpy(&val, m_buffer + m_offset + 348, sizeof(std::uint64_t));
        return SBE_LITTLE_ENDIAN_ENCODE_64(val);
    }

    ExecutionReportStatus532 &partyDetailsListReqID(const std::uint64_t value) SBE_NOEXCEPT
    {
        std::uint64_t val = SBE_LITTLE_ENDIAN_ENCODE_64(value);
        std::memcpy(m_buffer + m_offset + 348, &val, sizeof(std::uint64_t));
        return *this;
    }

    SBE_NODISCARD static const char * OrderIDMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::SEMANTIC_TYPE: return "int";
            case MetaAttribute::PRESENCE: return "required";
            default: return "";
        }
    }

    static SBE_CONSTEXPR std::uint16_t orderIDId() SBE_NOEXCEPT
    {
        return 37;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t orderIDSinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool orderIDInActingVersion() SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= orderIDSinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t orderIDEncodingOffset() SBE_NOEXCEPT
    {
        return 356;
    }

    static SBE_CONSTEXPR std::uint64_t orderIDNullValue() SBE_NOEXCEPT
    {
        return SBE_NULLVALUE_UINT64;
    }

    static SBE_CONSTEXPR std::uint64_t orderIDMinValue() SBE_NOEXCEPT
    {
        return 0x0L;
    }

    static SBE_CONSTEXPR std::uint64_t orderIDMaxValue() SBE_NOEXCEPT
    {
        return 0xfffffffffffffffeL;
    }

    static SBE_CONSTEXPR std::size_t orderIDEncodingLength() SBE_NOEXCEPT
    {
        return 8;
    }

    SBE_NODISCARD std::uint64_t orderID() const SBE_NOEXCEPT
    {
        std::uint64_t val;
        std::memcpy(&val, m_buffer + m_offset + 356, sizeof(std::uint64_t));
        return SBE_LITTLE_ENDIAN_ENCODE_64(val);
    }

    ExecutionReportStatus532 &orderID(const std::uint64_t value) SBE_NOEXCEPT
    {
        std::uint64_t val = SBE_LITTLE_ENDIAN_ENCODE_64(value);
        std::memcpy(m_buffer + m_offset + 356, &val, sizeof(std::uint64_t));
        return *this;
    }

    SBE_NODISCARD static const char * PriceMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::SEMANTIC_TYPE: return "Price";
            case MetaAttribute::PRESENCE: return "required";
            default: return "";
        }
    }

    static SBE_CONSTEXPR std::uint16_t priceId() SBE_NOEXCEPT
    {
        return 44;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t priceSinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool priceInActingVersion() SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= priceSinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t priceEncodingOffset() SBE_NOEXCEPT
    {
        return 364;
    }

private:
    PRICENULL9 m_price;

public:
    SBE_NODISCARD PRICENULL9 &price()
    {
        m_price.wrap(m_buffer, m_offset + 364, m_actingVersion, m_bufferLength);
        return m_price;
    }

    SBE_NODISCARD static const char * StopPxMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::SEMANTIC_TYPE: return "Price";
            case MetaAttribute::PRESENCE: return "required";
            default: return "";
        }
    }

    static SBE_CONSTEXPR std::uint16_t stopPxId() SBE_NOEXCEPT
    {
        return 99;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t stopPxSinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool stopPxInActingVersion() SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= stopPxSinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t stopPxEncodingOffset() SBE_NOEXCEPT
    {
        return 372;
    }

private:
    PRICENULL9 m_stopPx;

public:
    SBE_NODISCARD PRICENULL9 &stopPx()
    {
        m_stopPx.wrap(m_buffer, m_offset + 372, m_actingVersion, m_bufferLength);
        return m_stopPx;
    }

    SBE_NODISCARD static const char * TransactTimeMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::SEMANTIC_TYPE: return "int";
            case MetaAttribute::PRESENCE: return "required";
            default: return "";
        }
    }

    static SBE_CONSTEXPR std::uint16_t transactTimeId() SBE_NOEXCEPT
    {
        return 60;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t transactTimeSinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool transactTimeInActingVersion() SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= transactTimeSinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t transactTimeEncodingOffset() SBE_NOEXCEPT
    {
        return 380;
    }

    static SBE_CONSTEXPR std::uint64_t transactTimeNullValue() SBE_NOEXCEPT
    {
        return SBE_NULLVALUE_UINT64;
    }

    static SBE_CONSTEXPR std::uint64_t transactTimeMinValue() SBE_NOEXCEPT
    {
        return 0x0L;
    }

    static SBE_CONSTEXPR std::uint64_t transactTimeMaxValue() SBE_NOEXCEPT
    {
        return 0xfffffffffffffffeL;
    }

    static SBE_CONSTEXPR std::size_t transactTimeEncodingLength() SBE_NOEXCEPT
    {
        return 8;
    }

    SBE_NODISCARD std::uint64_t transactTime() const SBE_NOEXCEPT
    {
        std::uint64_t val;
        std::memcpy(&val, m_buffer + m_offset + 380, sizeof(std::uint64_t));
        return SBE_LITTLE_ENDIAN_ENCODE_64(val);
    }

    ExecutionReportStatus532 &transactTime(const std::uint64_t value) SBE_NOEXCEPT
    {
        std::uint64_t val = SBE_LITTLE_ENDIAN_ENCODE_64(value);
        std::memcpy(m_buffer + m_offset + 380, &val, sizeof(std::uint64_t));
        return *this;
    }

    SBE_NODISCARD static const char * SendingTimeEpochMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::SEMANTIC_TYPE: return "int";
            case MetaAttribute::PRESENCE: return "required";
            default: return "";
        }
    }

    static SBE_CONSTEXPR std::uint16_t sendingTimeEpochId() SBE_NOEXCEPT
    {
        return 5297;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t sendingTimeEpochSinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool sendingTimeEpochInActingVersion() SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= sendingTimeEpochSinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t sendingTimeEpochEncodingOffset() SBE_NOEXCEPT
    {
        return 388;
    }

    static SBE_CONSTEXPR std::uint64_t sendingTimeEpochNullValue() SBE_NOEXCEPT
    {
        return SBE_NULLVALUE_UINT64;
    }

    static SBE_CONSTEXPR std::uint64_t sendingTimeEpochMinValue() SBE_NOEXCEPT
    {
        return 0x0L;
    }

    static SBE_CONSTEXPR std::uint64_t sendingTimeEpochMaxValue() SBE_NOEXCEPT
    {
        return 0xfffffffffffffffeL;
    }

    static SBE_CONSTEXPR std::size_t sendingTimeEpochEncodingLength() SBE_NOEXCEPT
    {
        return 8;
    }

    SBE_NODISCARD std::uint64_t sendingTimeEpoch() const SBE_NOEXCEPT
    {
        std::uint64_t val;
        std::memcpy(&val, m_buffer + m_offset + 388, sizeof(std::uint64_t));
        return SBE_LITTLE_ENDIAN_ENCODE_64(val);
    }

    ExecutionReportStatus532 &sendingTimeEpoch(const std::uint64_t value) SBE_NOEXCEPT
    {
        std::uint64_t val = SBE_LITTLE_ENDIAN_ENCODE_64(value);
        std::memcpy(m_buffer + m_offset + 388, &val, sizeof(std::uint64_t));
        return *this;
    }

    SBE_NODISCARD static const char * OrderRequestIDMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::SEMANTIC_TYPE: return "int";
            case MetaAttribute::PRESENCE: return "required";
            default: return "";
        }
    }

    static SBE_CONSTEXPR std::uint16_t orderRequestIDId() SBE_NOEXCEPT
    {
        return 2422;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t orderRequestIDSinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool orderRequestIDInActingVersion() SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= orderRequestIDSinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t orderRequestIDEncodingOffset() SBE_NOEXCEPT
    {
        return 396;
    }

    static SBE_CONSTEXPR std::uint64_t orderRequestIDNullValue() SBE_NOEXCEPT
    {
        return SBE_NULLVALUE_UINT64;
    }

    static SBE_CONSTEXPR std::uint64_t orderRequestIDMinValue() SBE_NOEXCEPT
    {
        return 0x0L;
    }

    static SBE_CONSTEXPR std::uint64_t orderRequestIDMaxValue() SBE_NOEXCEPT
    {
        return 0xfffffffffffffffeL;
    }

    static SBE_CONSTEXPR std::size_t orderRequestIDEncodingLength() SBE_NOEXCEPT
    {
        return 8;
    }

    SBE_NODISCARD std::uint64_t orderRequestID() const SBE_NOEXCEPT
    {
        std::uint64_t val;
        std::memcpy(&val, m_buffer + m_offset + 396, sizeof(std::uint64_t));
        return SBE_LITTLE_ENDIAN_ENCODE_64(val);
    }

    ExecutionReportStatus532 &orderRequestID(const std::uint64_t value) SBE_NOEXCEPT
    {
        std::uint64_t val = SBE_LITTLE_ENDIAN_ENCODE_64(value);
        std::memcpy(m_buffer + m_offset + 396, &val, sizeof(std::uint64_t));
        return *this;
    }

    SBE_NODISCARD static const char * OrdStatusReqIDMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::SEMANTIC_TYPE: return "int";
            case MetaAttribute::PRESENCE: return "optional";
            default: return "";
        }
    }

    static SBE_CONSTEXPR std::uint16_t ordStatusReqIDId() SBE_NOEXCEPT
    {
        return 790;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t ordStatusReqIDSinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool ordStatusReqIDInActingVersion() SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= ordStatusReqIDSinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t ordStatusReqIDEncodingOffset() SBE_NOEXCEPT
    {
        return 404;
    }

    static SBE_CONSTEXPR std::uint64_t ordStatusReqIDNullValue() SBE_NOEXCEPT
    {
        return 0xffffffffffffffffL;
    }

    static SBE_CONSTEXPR std::uint64_t ordStatusReqIDMinValue() SBE_NOEXCEPT
    {
        return 0x0L;
    }

    static SBE_CONSTEXPR std::uint64_t ordStatusReqIDMaxValue() SBE_NOEXCEPT
    {
        return 0xfffffffffffffffeL;
    }

    static SBE_CONSTEXPR std::size_t ordStatusReqIDEncodingLength() SBE_NOEXCEPT
    {
        return 8;
    }

    SBE_NODISCARD std::uint64_t ordStatusReqID() const SBE_NOEXCEPT
    {
        std::uint64_t val;
        std::memcpy(&val, m_buffer + m_offset + 404, sizeof(std::uint64_t));
        return SBE_LITTLE_ENDIAN_ENCODE_64(val);
    }

    ExecutionReportStatus532 &ordStatusReqID(const std::uint64_t value) SBE_NOEXCEPT
    {
        std::uint64_t val = SBE_LITTLE_ENDIAN_ENCODE_64(value);
        std::memcpy(m_buffer + m_offset + 404, &val, sizeof(std::uint64_t));
        return *this;
    }

    SBE_NODISCARD static const char * MassStatusReqIDMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::SEMANTIC_TYPE: return "int";
            case MetaAttribute::PRESENCE: return "optional";
            default: return "";
        }
    }

    static SBE_CONSTEXPR std::uint16_t massStatusReqIDId() SBE_NOEXCEPT
    {
        return 584;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t massStatusReqIDSinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool massStatusReqIDInActingVersion() SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= massStatusReqIDSinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t massStatusReqIDEncodingOffset() SBE_NOEXCEPT
    {
        return 412;
    }

    static SBE_CONSTEXPR std::uint64_t massStatusReqIDNullValue() SBE_NOEXCEPT
    {
        return 0xffffffffffffffffL;
    }

    static SBE_CONSTEXPR std::uint64_t massStatusReqIDMinValue() SBE_NOEXCEPT
    {
        return 0x0L;
    }

    static SBE_CONSTEXPR std::uint64_t massStatusReqIDMaxValue() SBE_NOEXCEPT
    {
        return 0xfffffffffffffffeL;
    }

    static SBE_CONSTEXPR std::size_t massStatusReqIDEncodingLength() SBE_NOEXCEPT
    {
        return 8;
    }

    SBE_NODISCARD std::uint64_t massStatusReqID() const SBE_NOEXCEPT
    {
        std::uint64_t val;
        std::memcpy(&val, m_buffer + m_offset + 412, sizeof(std::uint64_t));
        return SBE_LITTLE_ENDIAN_ENCODE_64(val);
    }

    ExecutionReportStatus532 &massStatusReqID(const std::uint64_t value) SBE_NOEXCEPT
    {
        std::uint64_t val = SBE_LITTLE_ENDIAN_ENCODE_64(value);
        std::memcpy(m_buffer + m_offset + 412, &val, sizeof(std::uint64_t));
        return *this;
    }

    SBE_NODISCARD static const char * CrossIDMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::SEMANTIC_TYPE: return "int";
            case MetaAttribute::PRESENCE: return "optional";
            default: return "";
        }
    }

    static SBE_CONSTEXPR std::uint16_t crossIDId() SBE_NOEXCEPT
    {
        return 548;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t crossIDSinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool crossIDInActingVersion() SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= crossIDSinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t crossIDEncodingOffset() SBE_NOEXCEPT
    {
        return 420;
    }

    static SBE_CONSTEXPR std::uint64_t crossIDNullValue() SBE_NOEXCEPT
    {
        return 0xffffffffffffffffL;
    }

    static SBE_CONSTEXPR std::uint64_t crossIDMinValue() SBE_NOEXCEPT
    {
        return 0x0L;
    }

    static SBE_CONSTEXPR std::uint64_t crossIDMaxValue() SBE_NOEXCEPT
    {
        return 0xfffffffffffffffeL;
    }

    static SBE_CONSTEXPR std::size_t crossIDEncodingLength() SBE_NOEXCEPT
    {
        return 8;
    }

    SBE_NODISCARD std::uint64_t crossID() const SBE_NOEXCEPT
    {
        std::uint64_t val;
        std::memcpy(&val, m_buffer + m_offset + 420, sizeof(std::uint64_t));
        return SBE_LITTLE_ENDIAN_ENCODE_64(val);
    }

    ExecutionReportStatus532 &crossID(const std::uint64_t value) SBE_NOEXCEPT
    {
        std::uint64_t val = SBE_LITTLE_ENDIAN_ENCODE_64(value);
        std::memcpy(m_buffer + m_offset + 420, &val, sizeof(std::uint64_t));
        return *this;
    }

    SBE_NODISCARD static const char * HostCrossIDMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::SEMANTIC_TYPE: return "int";
            case MetaAttribute::PRESENCE: return "optional";
            default: return "";
        }
    }

    static SBE_CONSTEXPR std::uint16_t hostCrossIDId() SBE_NOEXCEPT
    {
        return 961;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t hostCrossIDSinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool hostCrossIDInActingVersion() SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= hostCrossIDSinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t hostCrossIDEncodingOffset() SBE_NOEXCEPT
    {
        return 428;
    }

    static SBE_CONSTEXPR std::uint64_t hostCrossIDNullValue() SBE_NOEXCEPT
    {
        return 0xffffffffffffffffL;
    }

    static SBE_CONSTEXPR std::uint64_t hostCrossIDMinValue() SBE_NOEXCEPT
    {
        return 0x0L;
    }

    static SBE_CONSTEXPR std::uint64_t hostCrossIDMaxValue() SBE_NOEXCEPT
    {
        return 0xfffffffffffffffeL;
    }

    static SBE_CONSTEXPR std::size_t hostCrossIDEncodingLength() SBE_NOEXCEPT
    {
        return 8;
    }

    SBE_NODISCARD std::uint64_t hostCrossID() const SBE_NOEXCEPT
    {
        std::uint64_t val;
        std::memcpy(&val, m_buffer + m_offset + 428, sizeof(std::uint64_t));
        return SBE_LITTLE_ENDIAN_ENCODE_64(val);
    }

    ExecutionReportStatus532 &hostCrossID(const std::uint64_t value) SBE_NOEXCEPT
    {
        std::uint64_t val = SBE_LITTLE_ENDIAN_ENCODE_64(value);
        std::memcpy(m_buffer + m_offset + 428, &val, sizeof(std::uint64_t));
        return *this;
    }

    SBE_NODISCARD static const char * LocationMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::SEMANTIC_TYPE: return "String";
            case MetaAttribute::PRESENCE: return "required";
            default: return "";
        }
    }

    static SBE_CONSTEXPR std::uint16_t locationId() SBE_NOEXCEPT
    {
        return 9537;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t locationSinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool locationInActingVersion() SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= locationSinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t locationEncodingOffset() SBE_NOEXCEPT
    {
        return 436;
    }

    static SBE_CONSTEXPR char locationNullValue() SBE_NOEXCEPT
    {
        return (char)0;
    }

    static SBE_CONSTEXPR char locationMinValue() SBE_NOEXCEPT
    {
        return (char)32;
    }

    static SBE_CONSTEXPR char locationMaxValue() SBE_NOEXCEPT
    {
        return (char)126;
    }

    static SBE_CONSTEXPR std::size_t locationEncodingLength() SBE_NOEXCEPT
    {
        return 5;
    }

    static SBE_CONSTEXPR std::uint64_t locationLength() SBE_NOEXCEPT
    {
        return 5;
    }

    SBE_NODISCARD const char *location() const SBE_NOEXCEPT
    {
        return m_buffer + m_offset + 436;
    }

    SBE_NODISCARD char *location() SBE_NOEXCEPT
    {
        return m_buffer + m_offset + 436;
    }

    SBE_NODISCARD char location(const std::uint64_t index) const
    {
        if (index >= 5)
        {
            throw std::runtime_error("index out of range for location [E104]");
        }

        char val;
        std::memcpy(&val, m_buffer + m_offset + 436 + (index * 1), sizeof(char));
        return (val);
    }

    ExecutionReportStatus532 &location(const std::uint64_t index, const char value)
    {
        if (index >= 5)
        {
            throw std::runtime_error("index out of range for location [E105]");
        }

        char val = (value);
        std::memcpy(m_buffer + m_offset + 436 + (index * 1), &val, sizeof(char));
        return *this;
    }

    std::uint64_t getLocation(char *const dst, const std::uint64_t length) const
    {
        if (length > 5)
        {
            throw std::runtime_error("length too large for getLocation [E106]");
        }

        std::memcpy(dst, m_buffer + m_offset + 436, sizeof(char) * static_cast<size_t>(length));
        return length;
    }

    ExecutionReportStatus532 &putLocation(const char *const src) SBE_NOEXCEPT
    {
        std::memcpy(m_buffer + m_offset + 436, src, sizeof(char) * 5);
        return *this;
    }

    SBE_NODISCARD std::string getLocationAsString() const
    {
        const char *buffer = m_buffer + m_offset + 436;
        size_t length = 0;

        for (; length < 5 && *(buffer + length) != '\0'; ++length);
        std::string result(buffer, length);

        return result;
    }

    std::string getLocationAsJsonEscapedString()
    {
        std::ostringstream oss;
        std::string s = getLocationAsString();

        for (const auto c : s)
        {
            switch (c)
            {
                case '"': oss << "\\\""; break;
                case '\\': oss << "\\\\"; break;
                case '\b': oss << "\\b"; break;
                case '\f': oss << "\\f"; break;
                case '\n': oss << "\\n"; break;
                case '\r': oss << "\\r"; break;
                case '\t': oss << "\\t"; break;

                default:
                    if ('\x00' <= c && c <= '\x1f')
                    {
                        oss << "\\u" << std::hex << std::setw(4)
                            << std::setfill('0') << (int)(c);
                    }
                    else
                    {
                        oss << c;
                    }
            }
        }

        return oss.str();
    }

    #if __cplusplus >= 201703L
    SBE_NODISCARD std::string_view getLocationAsStringView() const SBE_NOEXCEPT
    {
        const char *buffer = m_buffer + m_offset + 436;
        size_t length = 0;

        for (; length < 5 && *(buffer + length) != '\0'; ++length);
        std::string_view result(buffer, length);

        return result;
    }
    #endif

    #if __cplusplus >= 201703L
    ExecutionReportStatus532 &putLocation(const std::string_view str)
    {
        const size_t srcLength = str.length();
        if (srcLength > 5)
        {
            throw std::runtime_error("string too large for putLocation [E106]");
        }

        std::memcpy(m_buffer + m_offset + 436, str.data(), srcLength);
        for (size_t start = srcLength; start < 5; ++start)
        {
            m_buffer[m_offset + 436 + start] = 0;
        }

        return *this;
    }
    #else
    ExecutionReportStatus532 &putLocation(const std::string& str)
    {
        const size_t srcLength = str.length();
        if (srcLength > 5)
        {
            throw std::runtime_error("string too large for putLocation [E106]");
        }

        std::memcpy(m_buffer + m_offset + 436, str.c_str(), srcLength);
        for (size_t start = srcLength; start < 5; ++start)
        {
            m_buffer[m_offset + 436 + start] = 0;
        }

        return *this;
    }
    #endif

    SBE_NODISCARD static const char * SecurityIDMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::SEMANTIC_TYPE: return "int";
            case MetaAttribute::PRESENCE: return "required";
            default: return "";
        }
    }

    static SBE_CONSTEXPR std::uint16_t securityIDId() SBE_NOEXCEPT
    {
        return 48;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t securityIDSinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool securityIDInActingVersion() SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= securityIDSinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t securityIDEncodingOffset() SBE_NOEXCEPT
    {
        return 441;
    }

    static SBE_CONSTEXPR std::int32_t securityIDNullValue() SBE_NOEXCEPT
    {
        return SBE_NULLVALUE_INT32;
    }

    static SBE_CONSTEXPR std::int32_t securityIDMinValue() SBE_NOEXCEPT
    {
        return -2147483647;
    }

    static SBE_CONSTEXPR std::int32_t securityIDMaxValue() SBE_NOEXCEPT
    {
        return 2147483647;
    }

    static SBE_CONSTEXPR std::size_t securityIDEncodingLength() SBE_NOEXCEPT
    {
        return 4;
    }

    SBE_NODISCARD std::int32_t securityID() const SBE_NOEXCEPT
    {
        std::int32_t val;
        std::memcpy(&val, m_buffer + m_offset + 441, sizeof(std::int32_t));
        return SBE_LITTLE_ENDIAN_ENCODE_32(val);
    }

    ExecutionReportStatus532 &securityID(const std::int32_t value) SBE_NOEXCEPT
    {
        std::int32_t val = SBE_LITTLE_ENDIAN_ENCODE_32(value);
        std::memcpy(m_buffer + m_offset + 441, &val, sizeof(std::int32_t));
        return *this;
    }

    SBE_NODISCARD static const char * OrderQtyMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::SEMANTIC_TYPE: return "int";
            case MetaAttribute::PRESENCE: return "required";
            default: return "";
        }
    }

    static SBE_CONSTEXPR std::uint16_t orderQtyId() SBE_NOEXCEPT
    {
        return 38;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t orderQtySinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool orderQtyInActingVersion() SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= orderQtySinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t orderQtyEncodingOffset() SBE_NOEXCEPT
    {
        return 445;
    }

    static SBE_CONSTEXPR std::uint32_t orderQtyNullValue() SBE_NOEXCEPT
    {
        return SBE_NULLVALUE_UINT32;
    }

    static SBE_CONSTEXPR std::uint32_t orderQtyMinValue() SBE_NOEXCEPT
    {
        return 0;
    }

    static SBE_CONSTEXPR std::uint32_t orderQtyMaxValue() SBE_NOEXCEPT
    {
        return 4294967294;
    }

    static SBE_CONSTEXPR std::size_t orderQtyEncodingLength() SBE_NOEXCEPT
    {
        return 4;
    }

    SBE_NODISCARD std::uint32_t orderQty() const SBE_NOEXCEPT
    {
        std::uint32_t val;
        std::memcpy(&val, m_buffer + m_offset + 445, sizeof(std::uint32_t));
        return SBE_LITTLE_ENDIAN_ENCODE_32(val);
    }

    ExecutionReportStatus532 &orderQty(const std::uint32_t value) SBE_NOEXCEPT
    {
        std::uint32_t val = SBE_LITTLE_ENDIAN_ENCODE_32(value);
        std::memcpy(m_buffer + m_offset + 445, &val, sizeof(std::uint32_t));
        return *this;
    }

    SBE_NODISCARD static const char * CumQtyMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::SEMANTIC_TYPE: return "int";
            case MetaAttribute::PRESENCE: return "required";
            default: return "";
        }
    }

    static SBE_CONSTEXPR std::uint16_t cumQtyId() SBE_NOEXCEPT
    {
        return 14;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t cumQtySinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool cumQtyInActingVersion() SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= cumQtySinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t cumQtyEncodingOffset() SBE_NOEXCEPT
    {
        return 449;
    }

    static SBE_CONSTEXPR std::uint32_t cumQtyNullValue() SBE_NOEXCEPT
    {
        return SBE_NULLVALUE_UINT32;
    }

    static SBE_CONSTEXPR std::uint32_t cumQtyMinValue() SBE_NOEXCEPT
    {
        return 0;
    }

    static SBE_CONSTEXPR std::uint32_t cumQtyMaxValue() SBE_NOEXCEPT
    {
        return 4294967294;
    }

    static SBE_CONSTEXPR std::size_t cumQtyEncodingLength() SBE_NOEXCEPT
    {
        return 4;
    }

    SBE_NODISCARD std::uint32_t cumQty() const SBE_NOEXCEPT
    {
        std::uint32_t val;
        std::memcpy(&val, m_buffer + m_offset + 449, sizeof(std::uint32_t));
        return SBE_LITTLE_ENDIAN_ENCODE_32(val);
    }

    ExecutionReportStatus532 &cumQty(const std::uint32_t value) SBE_NOEXCEPT
    {
        std::uint32_t val = SBE_LITTLE_ENDIAN_ENCODE_32(value);
        std::memcpy(m_buffer + m_offset + 449, &val, sizeof(std::uint32_t));
        return *this;
    }

    SBE_NODISCARD static const char * LeavesQtyMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::SEMANTIC_TYPE: return "int";
            case MetaAttribute::PRESENCE: return "required";
            default: return "";
        }
    }

    static SBE_CONSTEXPR std::uint16_t leavesQtyId() SBE_NOEXCEPT
    {
        return 151;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t leavesQtySinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool leavesQtyInActingVersion() SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= leavesQtySinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t leavesQtyEncodingOffset() SBE_NOEXCEPT
    {
        return 453;
    }

    static SBE_CONSTEXPR std::uint32_t leavesQtyNullValue() SBE_NOEXCEPT
    {
        return SBE_NULLVALUE_UINT32;
    }

    static SBE_CONSTEXPR std::uint32_t leavesQtyMinValue() SBE_NOEXCEPT
    {
        return 0;
    }

    static SBE_CONSTEXPR std::uint32_t leavesQtyMaxValue() SBE_NOEXCEPT
    {
        return 4294967294;
    }

    static SBE_CONSTEXPR std::size_t leavesQtyEncodingLength() SBE_NOEXCEPT
    {
        return 4;
    }

    SBE_NODISCARD std::uint32_t leavesQty() const SBE_NOEXCEPT
    {
        std::uint32_t val;
        std::memcpy(&val, m_buffer + m_offset + 453, sizeof(std::uint32_t));
        return SBE_LITTLE_ENDIAN_ENCODE_32(val);
    }

    ExecutionReportStatus532 &leavesQty(const std::uint32_t value) SBE_NOEXCEPT
    {
        std::uint32_t val = SBE_LITTLE_ENDIAN_ENCODE_32(value);
        std::memcpy(m_buffer + m_offset + 453, &val, sizeof(std::uint32_t));
        return *this;
    }

    SBE_NODISCARD static const char * MinQtyMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::SEMANTIC_TYPE: return "int";
            case MetaAttribute::PRESENCE: return "optional";
            default: return "";
        }
    }

    static SBE_CONSTEXPR std::uint16_t minQtyId() SBE_NOEXCEPT
    {
        return 110;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t minQtySinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool minQtyInActingVersion() SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= minQtySinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t minQtyEncodingOffset() SBE_NOEXCEPT
    {
        return 457;
    }

    static SBE_CONSTEXPR std::uint32_t minQtyNullValue() SBE_NOEXCEPT
    {
        return 4294967295;
    }

    static SBE_CONSTEXPR std::uint32_t minQtyMinValue() SBE_NOEXCEPT
    {
        return 0;
    }

    static SBE_CONSTEXPR std::uint32_t minQtyMaxValue() SBE_NOEXCEPT
    {
        return 4294967294;
    }

    static SBE_CONSTEXPR std::size_t minQtyEncodingLength() SBE_NOEXCEPT
    {
        return 4;
    }

    SBE_NODISCARD std::uint32_t minQty() const SBE_NOEXCEPT
    {
        std::uint32_t val;
        std::memcpy(&val, m_buffer + m_offset + 457, sizeof(std::uint32_t));
        return SBE_LITTLE_ENDIAN_ENCODE_32(val);
    }

    ExecutionReportStatus532 &minQty(const std::uint32_t value) SBE_NOEXCEPT
    {
        std::uint32_t val = SBE_LITTLE_ENDIAN_ENCODE_32(value);
        std::memcpy(m_buffer + m_offset + 457, &val, sizeof(std::uint32_t));
        return *this;
    }

    SBE_NODISCARD static const char * DisplayQtyMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::SEMANTIC_TYPE: return "int";
            case MetaAttribute::PRESENCE: return "optional";
            default: return "";
        }
    }

    static SBE_CONSTEXPR std::uint16_t displayQtyId() SBE_NOEXCEPT
    {
        return 1138;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t displayQtySinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool displayQtyInActingVersion() SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= displayQtySinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t displayQtyEncodingOffset() SBE_NOEXCEPT
    {
        return 461;
    }

    static SBE_CONSTEXPR std::uint32_t displayQtyNullValue() SBE_NOEXCEPT
    {
        return 4294967295;
    }

    static SBE_CONSTEXPR std::uint32_t displayQtyMinValue() SBE_NOEXCEPT
    {
        return 0;
    }

    static SBE_CONSTEXPR std::uint32_t displayQtyMaxValue() SBE_NOEXCEPT
    {
        return 4294967294;
    }

    static SBE_CONSTEXPR std::size_t displayQtyEncodingLength() SBE_NOEXCEPT
    {
        return 4;
    }

    SBE_NODISCARD std::uint32_t displayQty() const SBE_NOEXCEPT
    {
        std::uint32_t val;
        std::memcpy(&val, m_buffer + m_offset + 461, sizeof(std::uint32_t));
        return SBE_LITTLE_ENDIAN_ENCODE_32(val);
    }

    ExecutionReportStatus532 &displayQty(const std::uint32_t value) SBE_NOEXCEPT
    {
        std::uint32_t val = SBE_LITTLE_ENDIAN_ENCODE_32(value);
        std::memcpy(m_buffer + m_offset + 461, &val, sizeof(std::uint32_t));
        return *this;
    }

    SBE_NODISCARD static const char * ExpireDateMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::SEMANTIC_TYPE: return "LocalMktDate";
            case MetaAttribute::PRESENCE: return "optional";
            default: return "";
        }
    }

    static SBE_CONSTEXPR std::uint16_t expireDateId() SBE_NOEXCEPT
    {
        return 432;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t expireDateSinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool expireDateInActingVersion() SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= expireDateSinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t expireDateEncodingOffset() SBE_NOEXCEPT
    {
        return 465;
    }

    static SBE_CONSTEXPR std::uint16_t expireDateNullValue() SBE_NOEXCEPT
    {
        return (std::uint16_t)65535;
    }

    static SBE_CONSTEXPR std::uint16_t expireDateMinValue() SBE_NOEXCEPT
    {
        return (std::uint16_t)0;
    }

    static SBE_CONSTEXPR std::uint16_t expireDateMaxValue() SBE_NOEXCEPT
    {
        return (std::uint16_t)65534;
    }

    static SBE_CONSTEXPR std::size_t expireDateEncodingLength() SBE_NOEXCEPT
    {
        return 2;
    }

    SBE_NODISCARD std::uint16_t expireDate() const SBE_NOEXCEPT
    {
        std::uint16_t val;
        std::memcpy(&val, m_buffer + m_offset + 465, sizeof(std::uint16_t));
        return SBE_LITTLE_ENDIAN_ENCODE_16(val);
    }

    ExecutionReportStatus532 &expireDate(const std::uint16_t value) SBE_NOEXCEPT
    {
        std::uint16_t val = SBE_LITTLE_ENDIAN_ENCODE_16(value);
        std::memcpy(m_buffer + m_offset + 465, &val, sizeof(std::uint16_t));
        return *this;
    }

    SBE_NODISCARD static const char * OrdStatusMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::SEMANTIC_TYPE: return "char";
            case MetaAttribute::PRESENCE: return "required";
            default: return "";
        }
    }

    static SBE_CONSTEXPR std::uint16_t ordStatusId() SBE_NOEXCEPT
    {
        return 39;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t ordStatusSinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool ordStatusInActingVersion() SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= ordStatusSinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t ordStatusEncodingOffset() SBE_NOEXCEPT
    {
        return 467;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t ordStatusEncodingLength() SBE_NOEXCEPT
    {
        return 1;
    }

    SBE_NODISCARD OrderStatus::Value ordStatus() const SBE_NOEXCEPT
    {
        char val;
        std::memcpy(&val, m_buffer + m_offset + 467, sizeof(char));
        return OrderStatus::get((val));
    }

    ExecutionReportStatus532 &ordStatus(const OrderStatus::Value value) SBE_NOEXCEPT
    {
        char val = (value);
        std::memcpy(m_buffer + m_offset + 467, &val, sizeof(char));
        return *this;
    }

    SBE_NODISCARD static const char * ExecTypeMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::SEMANTIC_TYPE: return "char";
            case MetaAttribute::PRESENCE: return "constant";
            default: return "";
        }
    }

    static SBE_CONSTEXPR std::uint16_t execTypeId() SBE_NOEXCEPT
    {
        return 150;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t execTypeSinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool execTypeInActingVersion() SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= execTypeSinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t execTypeEncodingOffset() SBE_NOEXCEPT
    {
        return 468;
    }

    static SBE_CONSTEXPR char execTypeNullValue() SBE_NOEXCEPT
    {
        return (char)0;
    }

    static SBE_CONSTEXPR char execTypeMinValue() SBE_NOEXCEPT
    {
        return (char)32;
    }

    static SBE_CONSTEXPR char execTypeMaxValue() SBE_NOEXCEPT
    {
        return (char)126;
    }

    static SBE_CONSTEXPR std::size_t execTypeEncodingLength() SBE_NOEXCEPT
    {
        return 0;
    }

    static SBE_CONSTEXPR std::uint64_t execTypeLength() SBE_NOEXCEPT
    {
        return 1;
    }

    SBE_NODISCARD const char *execType() const
    {
        static const std::uint8_t execTypeValues[] = { 73 };

        return (const char *)execTypeValues;
    }

    SBE_NODISCARD char execType(const std::uint64_t index) const
    {
        static const std::uint8_t execTypeValues[] = { 73 };

        return (char)execTypeValues[index];
    }

    std::uint64_t getExecType(char *dst, const std::uint64_t length) const
    {
        static std::uint8_t execTypeValues[] = { 73 };
        std::uint64_t bytesToCopy = length < sizeof(execTypeValues) ? length : sizeof(execTypeValues);

        std::memcpy(dst, execTypeValues, static_cast<size_t>(bytesToCopy));
        return bytesToCopy;
    }

    std::string getExecTypeAsString() const
    {
        static const std::uint8_t ExecTypeValues[] = { 73 };

        return std::string((const char *)ExecTypeValues, 1);
    }

    std::string getExecTypeAsJsonEscapedString()
    {
        std::ostringstream oss;
        std::string s = getExecTypeAsString();

        for (const auto c : s)
        {
            switch (c)
            {
                case '"': oss << "\\\""; break;
                case '\\': oss << "\\\\"; break;
                case '\b': oss << "\\b"; break;
                case '\f': oss << "\\f"; break;
                case '\n': oss << "\\n"; break;
                case '\r': oss << "\\r"; break;
                case '\t': oss << "\\t"; break;

                default:
                    if ('\x00' <= c && c <= '\x1f')
                    {
                        oss << "\\u" << std::hex << std::setw(4)
                            << std::setfill('0') << (int)(c);
                    }
                    else
                    {
                        oss << c;
                    }
            }
        }

        return oss.str();
    }

    SBE_NODISCARD static const char * OrdTypeMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::SEMANTIC_TYPE: return "char";
            case MetaAttribute::PRESENCE: return "required";
            default: return "";
        }
    }

    static SBE_CONSTEXPR std::uint16_t ordTypeId() SBE_NOEXCEPT
    {
        return 40;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t ordTypeSinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool ordTypeInActingVersion() SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= ordTypeSinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t ordTypeEncodingOffset() SBE_NOEXCEPT
    {
        return 468;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t ordTypeEncodingLength() SBE_NOEXCEPT
    {
        return 1;
    }

    SBE_NODISCARD OrderType::Value ordType() const SBE_NOEXCEPT
    {
        char val;
        std::memcpy(&val, m_buffer + m_offset + 468, sizeof(char));
        return OrderType::get((val));
    }

    ExecutionReportStatus532 &ordType(const OrderType::Value value) SBE_NOEXCEPT
    {
        char val = (value);
        std::memcpy(m_buffer + m_offset + 468, &val, sizeof(char));
        return *this;
    }

    SBE_NODISCARD static const char * SideMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::SEMANTIC_TYPE: return "int";
            case MetaAttribute::PRESENCE: return "required";
            default: return "";
        }
    }

    static SBE_CONSTEXPR std::uint16_t sideId() SBE_NOEXCEPT
    {
        return 54;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t sideSinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool sideInActingVersion() SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= sideSinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t sideEncodingOffset() SBE_NOEXCEPT
    {
        return 469;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t sideEncodingLength() SBE_NOEXCEPT
    {
        return 1;
    }

    SBE_NODISCARD SideReq::Value side() const SBE_NOEXCEPT
    {
        std::uint8_t val;
        std::memcpy(&val, m_buffer + m_offset + 469, sizeof(std::uint8_t));
        return SideReq::get((val));
    }

    ExecutionReportStatus532 &side(const SideReq::Value value) SBE_NOEXCEPT
    {
        std::uint8_t val = (value);
        std::memcpy(m_buffer + m_offset + 469, &val, sizeof(std::uint8_t));
        return *this;
    }

    SBE_NODISCARD static const char * TimeInForceMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::SEMANTIC_TYPE: return "int";
            case MetaAttribute::PRESENCE: return "required";
            default: return "";
        }
    }

    static SBE_CONSTEXPR std::uint16_t timeInForceId() SBE_NOEXCEPT
    {
        return 59;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t timeInForceSinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool timeInForceInActingVersion() SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= timeInForceSinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t timeInForceEncodingOffset() SBE_NOEXCEPT
    {
        return 470;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t timeInForceEncodingLength() SBE_NOEXCEPT
    {
        return 1;
    }

    SBE_NODISCARD TimeInForce::Value timeInForce() const SBE_NOEXCEPT
    {
        std::uint8_t val;
        std::memcpy(&val, m_buffer + m_offset + 470, sizeof(std::uint8_t));
        return TimeInForce::get((val));
    }

    ExecutionReportStatus532 &timeInForce(const TimeInForce::Value value) SBE_NOEXCEPT
    {
        std::uint8_t val = (value);
        std::memcpy(m_buffer + m_offset + 470, &val, sizeof(std::uint8_t));
        return *this;
    }

    SBE_NODISCARD static const char * ManualOrderIndicatorMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::SEMANTIC_TYPE: return "int";
            case MetaAttribute::PRESENCE: return "required";
            default: return "";
        }
    }

    static SBE_CONSTEXPR std::uint16_t manualOrderIndicatorId() SBE_NOEXCEPT
    {
        return 1028;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t manualOrderIndicatorSinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool manualOrderIndicatorInActingVersion() SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= manualOrderIndicatorSinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t manualOrderIndicatorEncodingOffset() SBE_NOEXCEPT
    {
        return 471;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t manualOrderIndicatorEncodingLength() SBE_NOEXCEPT
    {
        return 1;
    }

    SBE_NODISCARD ManualOrdIndReq::Value manualOrderIndicator() const SBE_NOEXCEPT
    {
        std::uint8_t val;
        std::memcpy(&val, m_buffer + m_offset + 471, sizeof(std::uint8_t));
        return ManualOrdIndReq::get((val));
    }

    ExecutionReportStatus532 &manualOrderIndicator(const ManualOrdIndReq::Value value) SBE_NOEXCEPT
    {
        std::uint8_t val = (value);
        std::memcpy(m_buffer + m_offset + 471, &val, sizeof(std::uint8_t));
        return *this;
    }

    SBE_NODISCARD static const char * PossRetransFlagMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::SEMANTIC_TYPE: return "int";
            case MetaAttribute::PRESENCE: return "required";
            default: return "";
        }
    }

    static SBE_CONSTEXPR std::uint16_t possRetransFlagId() SBE_NOEXCEPT
    {
        return 9765;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t possRetransFlagSinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool possRetransFlagInActingVersion() SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= possRetransFlagSinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t possRetransFlagEncodingOffset() SBE_NOEXCEPT
    {
        return 472;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t possRetransFlagEncodingLength() SBE_NOEXCEPT
    {
        return 1;
    }

    SBE_NODISCARD BooleanFlag::Value possRetransFlag() const SBE_NOEXCEPT
    {
        std::uint8_t val;
        std::memcpy(&val, m_buffer + m_offset + 472, sizeof(std::uint8_t));
        return BooleanFlag::get((val));
    }

    ExecutionReportStatus532 &possRetransFlag(const BooleanFlag::Value value) SBE_NOEXCEPT
    {
        std::uint8_t val = (value);
        std::memcpy(m_buffer + m_offset + 472, &val, sizeof(std::uint8_t));
        return *this;
    }

    SBE_NODISCARD static const char * LastRptRequestedMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::SEMANTIC_TYPE: return "int";
            case MetaAttribute::PRESENCE: return "required";
            default: return "";
        }
    }

    static SBE_CONSTEXPR std::uint16_t lastRptRequestedId() SBE_NOEXCEPT
    {
        return 912;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t lastRptRequestedSinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool lastRptRequestedInActingVersion() SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= lastRptRequestedSinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t lastRptRequestedEncodingOffset() SBE_NOEXCEPT
    {
        return 473;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t lastRptRequestedEncodingLength() SBE_NOEXCEPT
    {
        return 1;
    }

    SBE_NODISCARD BooleanNULL::Value lastRptRequested() const SBE_NOEXCEPT
    {
        std::uint8_t val;
        std::memcpy(&val, m_buffer + m_offset + 473, sizeof(std::uint8_t));
        return BooleanNULL::get((val));
    }

    ExecutionReportStatus532 &lastRptRequested(const BooleanNULL::Value value) SBE_NOEXCEPT
    {
        std::uint8_t val = (value);
        std::memcpy(m_buffer + m_offset + 473, &val, sizeof(std::uint8_t));
        return *this;
    }

    SBE_NODISCARD static const char * CrossTypeMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::SEMANTIC_TYPE: return "int";
            case MetaAttribute::PRESENCE: return "optional";
            default: return "";
        }
    }

    static SBE_CONSTEXPR std::uint16_t crossTypeId() SBE_NOEXCEPT
    {
        return 549;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t crossTypeSinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool crossTypeInActingVersion() SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= crossTypeSinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t crossTypeEncodingOffset() SBE_NOEXCEPT
    {
        return 474;
    }

    static SBE_CONSTEXPR std::uint8_t crossTypeNullValue() SBE_NOEXCEPT
    {
        return (std::uint8_t)255;
    }

    static SBE_CONSTEXPR std::uint8_t crossTypeMinValue() SBE_NOEXCEPT
    {
        return (std::uint8_t)0;
    }

    static SBE_CONSTEXPR std::uint8_t crossTypeMaxValue() SBE_NOEXCEPT
    {
        return (std::uint8_t)254;
    }

    static SBE_CONSTEXPR std::size_t crossTypeEncodingLength() SBE_NOEXCEPT
    {
        return 1;
    }

    SBE_NODISCARD std::uint8_t crossType() const SBE_NOEXCEPT
    {
        std::uint8_t val;
        std::memcpy(&val, m_buffer + m_offset + 474, sizeof(std::uint8_t));
        return (val);
    }

    ExecutionReportStatus532 &crossType(const std::uint8_t value) SBE_NOEXCEPT
    {
        std::uint8_t val = (value);
        std::memcpy(m_buffer + m_offset + 474, &val, sizeof(std::uint8_t));
        return *this;
    }

    SBE_NODISCARD static const char * ExecInstMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::SEMANTIC_TYPE: return "MultipleCharValue";
            case MetaAttribute::PRESENCE: return "required";
            default: return "";
        }
    }

    static SBE_CONSTEXPR std::uint16_t execInstId() SBE_NOEXCEPT
    {
        return 18;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t execInstSinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool execInstInActingVersion() SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= execInstSinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t execInstEncodingOffset() SBE_NOEXCEPT
    {
        return 475;
    }

private:
    ExecInst m_execInst;

public:
    SBE_NODISCARD ExecInst &execInst()
    {
        m_execInst.wrap(m_buffer, m_offset + 475, m_actingVersion, m_bufferLength);
        return m_execInst;
    }

    static SBE_CONSTEXPR std::size_t execInstEncodingLength() SBE_NOEXCEPT
    {
        return 1;
    }

    SBE_NODISCARD static const char * ExecutionModeMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::SEMANTIC_TYPE: return "char";
            case MetaAttribute::PRESENCE: return "required";
            default: return "";
        }
    }

    static SBE_CONSTEXPR std::uint16_t executionModeId() SBE_NOEXCEPT
    {
        return 5906;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t executionModeSinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool executionModeInActingVersion() SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= executionModeSinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t executionModeEncodingOffset() SBE_NOEXCEPT
    {
        return 476;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t executionModeEncodingLength() SBE_NOEXCEPT
    {
        return 1;
    }

    SBE_NODISCARD ExecMode::Value executionMode() const SBE_NOEXCEPT
    {
        char val;
        std::memcpy(&val, m_buffer + m_offset + 476, sizeof(char));
        return ExecMode::get((val));
    }

    ExecutionReportStatus532 &executionMode(const ExecMode::Value value) SBE_NOEXCEPT
    {
        char val = (value);
        std::memcpy(m_buffer + m_offset + 476, &val, sizeof(char));
        return *this;
    }

    SBE_NODISCARD static const char * LiquidityFlagMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::SEMANTIC_TYPE: return "int";
            case MetaAttribute::PRESENCE: return "required";
            default: return "";
        }
    }

    static SBE_CONSTEXPR std::uint16_t liquidityFlagId() SBE_NOEXCEPT
    {
        return 9373;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t liquidityFlagSinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool liquidityFlagInActingVersion() SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= liquidityFlagSinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t liquidityFlagEncodingOffset() SBE_NOEXCEPT
    {
        return 477;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t liquidityFlagEncodingLength() SBE_NOEXCEPT
    {
        return 1;
    }

    SBE_NODISCARD BooleanNULL::Value liquidityFlag() const SBE_NOEXCEPT
    {
        std::uint8_t val;
        std::memcpy(&val, m_buffer + m_offset + 477, sizeof(std::uint8_t));
        return BooleanNULL::get((val));
    }

    ExecutionReportStatus532 &liquidityFlag(const BooleanNULL::Value value) SBE_NOEXCEPT
    {
        std::uint8_t val = (value);
        std::memcpy(m_buffer + m_offset + 477, &val, sizeof(std::uint8_t));
        return *this;
    }

    SBE_NODISCARD static const char * ManagedOrderMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::SEMANTIC_TYPE: return "int";
            case MetaAttribute::PRESENCE: return "required";
            default: return "";
        }
    }

    static SBE_CONSTEXPR std::uint16_t managedOrderId() SBE_NOEXCEPT
    {
        return 6881;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t managedOrderSinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool managedOrderInActingVersion() SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= managedOrderSinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t managedOrderEncodingOffset() SBE_NOEXCEPT
    {
        return 478;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t managedOrderEncodingLength() SBE_NOEXCEPT
    {
        return 1;
    }

    SBE_NODISCARD BooleanNULL::Value managedOrder() const SBE_NOEXCEPT
    {
        std::uint8_t val;
        std::memcpy(&val, m_buffer + m_offset + 478, sizeof(std::uint8_t));
        return BooleanNULL::get((val));
    }

    ExecutionReportStatus532 &managedOrder(const BooleanNULL::Value value) SBE_NOEXCEPT
    {
        std::uint8_t val = (value);
        std::memcpy(m_buffer + m_offset + 478, &val, sizeof(std::uint8_t));
        return *this;
    }

    SBE_NODISCARD static const char * ShortSaleTypeMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::SEMANTIC_TYPE: return "int";
            case MetaAttribute::PRESENCE: return "required";
            default: return "";
        }
    }

    static SBE_CONSTEXPR std::uint16_t shortSaleTypeId() SBE_NOEXCEPT
    {
        return 5409;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t shortSaleTypeSinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool shortSaleTypeInActingVersion() SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= shortSaleTypeSinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t shortSaleTypeEncodingOffset() SBE_NOEXCEPT
    {
        return 479;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t shortSaleTypeEncodingLength() SBE_NOEXCEPT
    {
        return 1;
    }

    SBE_NODISCARD ShortSaleType::Value shortSaleType() const SBE_NOEXCEPT
    {
        std::uint8_t val;
        std::memcpy(&val, m_buffer + m_offset + 479, sizeof(std::uint8_t));
        return ShortSaleType::get((val));
    }

    ExecutionReportStatus532 &shortSaleType(const ShortSaleType::Value value) SBE_NOEXCEPT
    {
        std::uint8_t val = (value);
        std::memcpy(m_buffer + m_offset + 479, &val, sizeof(std::uint8_t));
        return *this;
    }

template<typename CharT, typename Traits>
friend std::basic_ostream<CharT, Traits>& operator << (
    std::basic_ostream<CharT, Traits>& builder, ExecutionReportStatus532 _writer)
{
    ExecutionReportStatus532 writer(_writer.m_buffer, _writer.m_offset,
        _writer.m_bufferLength, _writer.sbeBlockLength(), _writer.m_actingVersion);
    builder << '{';
    builder << R"("Name": "ExecutionReportStatus532", )";
    builder << R"("sbeTemplateId": )";
    builder << writer.sbeTemplateId();
    builder << ", ";

    builder << R"("SeqNum": )";
    builder << +writer.seqNum();

    builder << ", ";
    builder << R"("UUID": )";
    builder << +writer.uUID();

    builder << ", ";
    builder << R"("Text": )";
    builder << '"' <<
        writer.getTextAsJsonEscapedString().c_str() << '"';

    builder << ", ";
    builder << R"("ExecID": )";
    builder << '"' <<
        writer.getExecIDAsJsonEscapedString().c_str() << '"';

    builder << ", ";
    builder << R"("SenderID": )";
    builder << '"' <<
        writer.getSenderIDAsJsonEscapedString().c_str() << '"';

    builder << ", ";
    builder << R"("ClOrdID": )";
    builder << '"' <<
        writer.getClOrdIDAsJsonEscapedString().c_str() << '"';

    builder << ", ";
    builder << R"("PartyDetailsListReqID": )";
    builder << +writer.partyDetailsListReqID();

    builder << ", ";
    builder << R"("OrderID": )";
    builder << +writer.orderID();

    builder << ", ";
    builder << R"("Price": )";
    builder << writer.price();

    builder << ", ";
    builder << R"("StopPx": )";
    builder << writer.stopPx();

    builder << ", ";
    builder << R"("TransactTime": )";
    builder << +writer.transactTime();

    builder << ", ";
    builder << R"("SendingTimeEpoch": )";
    builder << +writer.sendingTimeEpoch();

    builder << ", ";
    builder << R"("OrderRequestID": )";
    builder << +writer.orderRequestID();

    builder << ", ";
    builder << R"("OrdStatusReqID": )";
    builder << +writer.ordStatusReqID();

    builder << ", ";
    builder << R"("MassStatusReqID": )";
    builder << +writer.massStatusReqID();

    builder << ", ";
    builder << R"("CrossID": )";
    builder << +writer.crossID();

    builder << ", ";
    builder << R"("HostCrossID": )";
    builder << +writer.hostCrossID();

    builder << ", ";
    builder << R"("Location": )";
    builder << '"' <<
        writer.getLocationAsJsonEscapedString().c_str() << '"';

    builder << ", ";
    builder << R"("SecurityID": )";
    builder << +writer.securityID();

    builder << ", ";
    builder << R"("OrderQty": )";
    builder << +writer.orderQty();

    builder << ", ";
    builder << R"("CumQty": )";
    builder << +writer.cumQty();

    builder << ", ";
    builder << R"("LeavesQty": )";
    builder << +writer.leavesQty();

    builder << ", ";
    builder << R"("MinQty": )";
    builder << +writer.minQty();

    builder << ", ";
    builder << R"("DisplayQty": )";
    builder << +writer.displayQty();

    builder << ", ";
    builder << R"("ExpireDate": )";
    builder << +writer.expireDate();

    builder << ", ";
    builder << R"("OrdStatus": )";
    builder << '"' << writer.ordStatus() << '"';

    builder << ", ";
    builder << R"("OrdType": )";
    builder << '"' << writer.ordType() << '"';

    builder << ", ";
    builder << R"("Side": )";
    builder << '"' << writer.side() << '"';

    builder << ", ";
    builder << R"("TimeInForce": )";
    builder << '"' << writer.timeInForce() << '"';

    builder << ", ";
    builder << R"("ManualOrderIndicator": )";
    builder << '"' << writer.manualOrderIndicator() << '"';

    builder << ", ";
    builder << R"("PossRetransFlag": )";
    builder << '"' << writer.possRetransFlag() << '"';

    builder << ", ";
    builder << R"("LastRptRequested": )";
    builder << '"' << writer.lastRptRequested() << '"';

    builder << ", ";
    builder << R"("CrossType": )";
    builder << +writer.crossType();

    builder << ", ";
    builder << R"("ExecInst": )";
    builder << writer.execInst();

    builder << ", ";
    builder << R"("ExecutionMode": )";
    builder << '"' << writer.executionMode() << '"';

    builder << ", ";
    builder << R"("LiquidityFlag": )";
    builder << '"' << writer.liquidityFlag() << '"';

    builder << ", ";
    builder << R"("ManagedOrder": )";
    builder << '"' << writer.managedOrder() << '"';

    builder << ", ";
    builder << R"("ShortSaleType": )";
    builder << '"' << writer.shortSaleType() << '"';

    builder << '}';

    return builder;
}

void skip()
{
}

SBE_NODISCARD static SBE_CONSTEXPR bool isConstLength() SBE_NOEXCEPT
{
    return true;
}

SBE_NODISCARD static SBE_CONSTEXPR_14 size_t computeLength()
{
#if defined(__GNUG__) && !defined(__clang__)
#pragma GCC diagnostic push
#pragma GCC diagnostic ignored "-Wtype-limits"
#endif
    size_t length = sbeBlockLength();

    return length;
#if defined(__GNUG__) && !defined(__clang__)
#pragma GCC diagnostic pop
#endif
}
};
}
#endif
