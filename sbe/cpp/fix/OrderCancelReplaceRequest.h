/* Generated SBE (Simple Binary Encoding) message codec */
#ifndef _FIX_ORDERCANCELREPLACEREQUEST_H_
#define _FIX_ORDERCANCELREPLACEREQUEST_H_

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


#include "MatchEventIndicator.h"
#include "IntQty32.h"
#include "Decimal64.h"
#include "MDUpdateAction.h"
#include "GroupSizeEncoding.h"
#include "Side.h"
#include "OptionalPrice.h"
#include "TickDirection.h"
#include "QuoteCondition.h"
#include "TimeInForce.h"
#include "MDEntryType.h"
#include "CtiCode.h"
#include "HandInst.h"
#include "NoAllocs.h"
#include "MMProtectionReset.h"
#include "MessageHeader.h"
#include "OrdType.h"
#include "MDQuoteType.h"
#include "BooleanType.h"
#include "SecurityIDSource.h"
#include "OFMOverride.h"
#include "CustOrderHandlingInst.h"
#include "TradeCondition.h"
#include "OpenCloseSettleFlag.h"
#include "CustomerOrFirm.h"
#include "MarketStateIdentifier.h"

namespace fix {

class OrderCancelReplaceRequest
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

    OrderCancelReplaceRequest() = default;

    OrderCancelReplaceRequest(
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

    OrderCancelReplaceRequest(char *buffer, const std::uint64_t bufferLength) :
        OrderCancelReplaceRequest(buffer, 0, bufferLength, sbeBlockLength(), sbeSchemaVersion())
    {
    }

    OrderCancelReplaceRequest(
        char *buffer,
        const std::uint64_t bufferLength,
        const std::uint64_t actingBlockLength,
        const std::uint64_t actingVersion) :
        OrderCancelReplaceRequest(buffer, 0, bufferLength, actingBlockLength, actingVersion)
    {
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint16_t sbeBlockLength() SBE_NOEXCEPT
    {
        return (std::uint16_t)204;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint16_t sbeBlockAndHeaderLength() SBE_NOEXCEPT
    {
        return MessageHeader::encodedLength() + sbeBlockLength();
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint16_t sbeTemplateId() SBE_NOEXCEPT
    {
        return (std::uint16_t)71;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint16_t sbeSchemaId() SBE_NOEXCEPT
    {
        return (std::uint16_t)1;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint16_t sbeSchemaVersion() SBE_NOEXCEPT
    {
        return (std::uint16_t)1;
    }

    SBE_NODISCARD static SBE_CONSTEXPR const char * sbeSemanticType() SBE_NOEXCEPT
    {
        return "G";
    }

    SBE_NODISCARD std::uint64_t offset() const SBE_NOEXCEPT
    {
        return m_offset;
    }

    OrderCancelReplaceRequest &wrapForEncode(char *buffer, const std::uint64_t offset, const std::uint64_t bufferLength)
    {
        return *this = OrderCancelReplaceRequest(buffer, offset, bufferLength, sbeBlockLength(), sbeSchemaVersion());
    }

    OrderCancelReplaceRequest &wrapAndApplyHeader(char *buffer, const std::uint64_t offset, const std::uint64_t bufferLength)
    {
        MessageHeader hdr(buffer, offset, bufferLength, sbeSchemaVersion());

        hdr
            .blockLength(sbeBlockLength())
            .templateId(sbeTemplateId())
            .schemaId(sbeSchemaId())
            .version(sbeSchemaVersion());

        return *this = OrderCancelReplaceRequest(
            buffer,
            offset + MessageHeader::encodedLength(),
            bufferLength,
            sbeBlockLength(),
            sbeSchemaVersion());
    }

    OrderCancelReplaceRequest &wrapForDecode(
        char *buffer,
        const std::uint64_t offset,
        const std::uint64_t actingBlockLength,
        const std::uint64_t actingVersion,
        const std::uint64_t bufferLength)
    {
        return *this = OrderCancelReplaceRequest(buffer, offset, bufferLength, actingBlockLength, actingVersion);
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
        OrderCancelReplaceRequest skipper(m_buffer, m_offset,
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

    SBE_NODISCARD static const char * AccountMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::SEMANTIC_TYPE: return "String";
            case MetaAttribute::PRESENCE: return "required";
            default: return "";
        }
    }

    static SBE_CONSTEXPR std::uint16_t accountId() SBE_NOEXCEPT
    {
        return 1;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t accountSinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool accountInActingVersion() SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= accountSinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t accountEncodingOffset() SBE_NOEXCEPT
    {
        return 0;
    }

    static SBE_CONSTEXPR char accountNullValue() SBE_NOEXCEPT
    {
        return (char)0;
    }

    static SBE_CONSTEXPR char accountMinValue() SBE_NOEXCEPT
    {
        return (char)32;
    }

    static SBE_CONSTEXPR char accountMaxValue() SBE_NOEXCEPT
    {
        return (char)126;
    }

    static SBE_CONSTEXPR std::size_t accountEncodingLength() SBE_NOEXCEPT
    {
        return 12;
    }

    static SBE_CONSTEXPR std::uint64_t accountLength() SBE_NOEXCEPT
    {
        return 12;
    }

    SBE_NODISCARD const char *account() const SBE_NOEXCEPT
    {
        return m_buffer + m_offset + 0;
    }

    SBE_NODISCARD char *account() SBE_NOEXCEPT
    {
        return m_buffer + m_offset + 0;
    }

    SBE_NODISCARD char account(const std::uint64_t index) const
    {
        if (index >= 12)
        {
            throw std::runtime_error("index out of range for account [E104]");
        }

        char val;
        std::memcpy(&val, m_buffer + m_offset + 0 + (index * 1), sizeof(char));
        return (val);
    }

    OrderCancelReplaceRequest &account(const std::uint64_t index, const char value)
    {
        if (index >= 12)
        {
            throw std::runtime_error("index out of range for account [E105]");
        }

        char val = (value);
        std::memcpy(m_buffer + m_offset + 0 + (index * 1), &val, sizeof(char));
        return *this;
    }

    std::uint64_t getAccount(char *const dst, const std::uint64_t length) const
    {
        if (length > 12)
        {
            throw std::runtime_error("length too large for getAccount [E106]");
        }

        std::memcpy(dst, m_buffer + m_offset + 0, sizeof(char) * static_cast<size_t>(length));
        return length;
    }

    OrderCancelReplaceRequest &putAccount(const char *const src) SBE_NOEXCEPT
    {
        std::memcpy(m_buffer + m_offset + 0, src, sizeof(char) * 12);
        return *this;
    }

    SBE_NODISCARD std::string getAccountAsString() const
    {
        const char *buffer = m_buffer + m_offset + 0;
        size_t length = 0;

        for (; length < 12 && *(buffer + length) != '\0'; ++length);
        std::string result(buffer, length);

        return result;
    }

    std::string getAccountAsJsonEscapedString()
    {
        std::ostringstream oss;
        std::string s = getAccountAsString();

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
    SBE_NODISCARD std::string_view getAccountAsStringView() const SBE_NOEXCEPT
    {
        const char *buffer = m_buffer + m_offset + 0;
        size_t length = 0;

        for (; length < 12 && *(buffer + length) != '\0'; ++length);
        std::string_view result(buffer, length);

        return result;
    }
    #endif

    #if __cplusplus >= 201703L
    OrderCancelReplaceRequest &putAccount(const std::string_view str)
    {
        const size_t srcLength = str.length();
        if (srcLength > 12)
        {
            throw std::runtime_error("string too large for putAccount [E106]");
        }

        std::memcpy(m_buffer + m_offset + 0, str.data(), srcLength);
        for (size_t start = srcLength; start < 12; ++start)
        {
            m_buffer[m_offset + 0 + start] = 0;
        }

        return *this;
    }
    #else
    OrderCancelReplaceRequest &putAccount(const std::string& str)
    {
        const size_t srcLength = str.length();
        if (srcLength > 12)
        {
            throw std::runtime_error("string too large for putAccount [E106]");
        }

        std::memcpy(m_buffer + m_offset + 0, str.c_str(), srcLength);
        for (size_t start = srcLength; start < 12; ++start)
        {
            m_buffer[m_offset + 0 + start] = 0;
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
        return 12;
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
        return m_buffer + m_offset + 12;
    }

    SBE_NODISCARD char *clOrdID() SBE_NOEXCEPT
    {
        return m_buffer + m_offset + 12;
    }

    SBE_NODISCARD char clOrdID(const std::uint64_t index) const
    {
        if (index >= 20)
        {
            throw std::runtime_error("index out of range for clOrdID [E104]");
        }

        char val;
        std::memcpy(&val, m_buffer + m_offset + 12 + (index * 1), sizeof(char));
        return (val);
    }

    OrderCancelReplaceRequest &clOrdID(const std::uint64_t index, const char value)
    {
        if (index >= 20)
        {
            throw std::runtime_error("index out of range for clOrdID [E105]");
        }

        char val = (value);
        std::memcpy(m_buffer + m_offset + 12 + (index * 1), &val, sizeof(char));
        return *this;
    }

    std::uint64_t getClOrdID(char *const dst, const std::uint64_t length) const
    {
        if (length > 20)
        {
            throw std::runtime_error("length too large for getClOrdID [E106]");
        }

        std::memcpy(dst, m_buffer + m_offset + 12, sizeof(char) * static_cast<size_t>(length));
        return length;
    }

    OrderCancelReplaceRequest &putClOrdID(const char *const src) SBE_NOEXCEPT
    {
        std::memcpy(m_buffer + m_offset + 12, src, sizeof(char) * 20);
        return *this;
    }

    SBE_NODISCARD std::string getClOrdIDAsString() const
    {
        const char *buffer = m_buffer + m_offset + 12;
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
        const char *buffer = m_buffer + m_offset + 12;
        size_t length = 0;

        for (; length < 20 && *(buffer + length) != '\0'; ++length);
        std::string_view result(buffer, length);

        return result;
    }
    #endif

    #if __cplusplus >= 201703L
    OrderCancelReplaceRequest &putClOrdID(const std::string_view str)
    {
        const size_t srcLength = str.length();
        if (srcLength > 20)
        {
            throw std::runtime_error("string too large for putClOrdID [E106]");
        }

        std::memcpy(m_buffer + m_offset + 12, str.data(), srcLength);
        for (size_t start = srcLength; start < 20; ++start)
        {
            m_buffer[m_offset + 12 + start] = 0;
        }

        return *this;
    }
    #else
    OrderCancelReplaceRequest &putClOrdID(const std::string& str)
    {
        const size_t srcLength = str.length();
        if (srcLength > 20)
        {
            throw std::runtime_error("string too large for putClOrdID [E106]");
        }

        std::memcpy(m_buffer + m_offset + 12, str.c_str(), srcLength);
        for (size_t start = srcLength; start < 20; ++start)
        {
            m_buffer[m_offset + 12 + start] = 0;
        }

        return *this;
    }
    #endif

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
        return 32;
    }

    static SBE_CONSTEXPR std::int64_t orderIDNullValue() SBE_NOEXCEPT
    {
        return SBE_NULLVALUE_INT64;
    }

    static SBE_CONSTEXPR std::int64_t orderIDMinValue() SBE_NOEXCEPT
    {
        return -9223372036854775807L;
    }

    static SBE_CONSTEXPR std::int64_t orderIDMaxValue() SBE_NOEXCEPT
    {
        return 9223372036854775807L;
    }

    static SBE_CONSTEXPR std::size_t orderIDEncodingLength() SBE_NOEXCEPT
    {
        return 8;
    }

    SBE_NODISCARD std::int64_t orderID() const SBE_NOEXCEPT
    {
        std::int64_t val;
        std::memcpy(&val, m_buffer + m_offset + 32, sizeof(std::int64_t));
        return SBE_LITTLE_ENDIAN_ENCODE_64(val);
    }

    OrderCancelReplaceRequest &orderID(const std::int64_t value) SBE_NOEXCEPT
    {
        std::int64_t val = SBE_LITTLE_ENDIAN_ENCODE_64(value);
        std::memcpy(m_buffer + m_offset + 32, &val, sizeof(std::int64_t));
        return *this;
    }

    SBE_NODISCARD static const char * HandInstMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::SEMANTIC_TYPE: return "char";
            case MetaAttribute::PRESENCE: return "required";
            default: return "";
        }
    }

    static SBE_CONSTEXPR std::uint16_t handInstId() SBE_NOEXCEPT
    {
        return 21;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t handInstSinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool handInstInActingVersion() SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= handInstSinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t handInstEncodingOffset() SBE_NOEXCEPT
    {
        return 40;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t handInstEncodingLength() SBE_NOEXCEPT
    {
        return 1;
    }

    SBE_NODISCARD HandInst::Value handInst() const SBE_NOEXCEPT
    {
        char val;
        std::memcpy(&val, m_buffer + m_offset + 40, sizeof(char));
        return HandInst::get((val));
    }

    OrderCancelReplaceRequest &handInst(const HandInst::Value value) SBE_NOEXCEPT
    {
        char val = (value);
        std::memcpy(m_buffer + m_offset + 40, &val, sizeof(char));
        return *this;
    }

    SBE_NODISCARD static const char * OrderQtyMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::SEMANTIC_TYPE: return "Qty";
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
        return 41;
    }

private:
    IntQty32 m_orderQty;

public:
    SBE_NODISCARD IntQty32 &orderQty()
    {
        m_orderQty.wrap(m_buffer, m_offset + 41, m_actingVersion, m_bufferLength);
        return m_orderQty;
    }

    SBE_NODISCARD static const char * CustOrderHandlingInstMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::SEMANTIC_TYPE: return "char";
            case MetaAttribute::PRESENCE: return "optional";
            default: return "";
        }
    }

    static SBE_CONSTEXPR std::uint16_t custOrderHandlingInstId() SBE_NOEXCEPT
    {
        return 1031;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t custOrderHandlingInstSinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool custOrderHandlingInstInActingVersion() SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= custOrderHandlingInstSinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t custOrderHandlingInstEncodingOffset() SBE_NOEXCEPT
    {
        return 45;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t custOrderHandlingInstEncodingLength() SBE_NOEXCEPT
    {
        return 1;
    }

    SBE_NODISCARD CustOrderHandlingInst::Value custOrderHandlingInst() const SBE_NOEXCEPT
    {
        char val;
        std::memcpy(&val, m_buffer + m_offset + 45, sizeof(char));
        return CustOrderHandlingInst::get((val));
    }

    OrderCancelReplaceRequest &custOrderHandlingInst(const CustOrderHandlingInst::Value value) SBE_NOEXCEPT
    {
        char val = (value);
        std::memcpy(m_buffer + m_offset + 45, &val, sizeof(char));
        return *this;
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
        return 46;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t ordTypeEncodingLength() SBE_NOEXCEPT
    {
        return 1;
    }

    SBE_NODISCARD OrdType::Value ordType() const SBE_NOEXCEPT
    {
        char val;
        std::memcpy(&val, m_buffer + m_offset + 46, sizeof(char));
        return OrdType::get((val));
    }

    OrderCancelReplaceRequest &ordType(const OrdType::Value value) SBE_NOEXCEPT
    {
        char val = (value);
        std::memcpy(m_buffer + m_offset + 46, &val, sizeof(char));
        return *this;
    }

    SBE_NODISCARD static const char * OrigClOrdIDMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::SEMANTIC_TYPE: return "String";
            case MetaAttribute::PRESENCE: return "required";
            default: return "";
        }
    }

    static SBE_CONSTEXPR std::uint16_t origClOrdIDId() SBE_NOEXCEPT
    {
        return 41;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t origClOrdIDSinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool origClOrdIDInActingVersion() SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= origClOrdIDSinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t origClOrdIDEncodingOffset() SBE_NOEXCEPT
    {
        return 47;
    }

    static SBE_CONSTEXPR char origClOrdIDNullValue() SBE_NOEXCEPT
    {
        return (char)0;
    }

    static SBE_CONSTEXPR char origClOrdIDMinValue() SBE_NOEXCEPT
    {
        return (char)32;
    }

    static SBE_CONSTEXPR char origClOrdIDMaxValue() SBE_NOEXCEPT
    {
        return (char)126;
    }

    static SBE_CONSTEXPR std::size_t origClOrdIDEncodingLength() SBE_NOEXCEPT
    {
        return 20;
    }

    static SBE_CONSTEXPR std::uint64_t origClOrdIDLength() SBE_NOEXCEPT
    {
        return 20;
    }

    SBE_NODISCARD const char *origClOrdID() const SBE_NOEXCEPT
    {
        return m_buffer + m_offset + 47;
    }

    SBE_NODISCARD char *origClOrdID() SBE_NOEXCEPT
    {
        return m_buffer + m_offset + 47;
    }

    SBE_NODISCARD char origClOrdID(const std::uint64_t index) const
    {
        if (index >= 20)
        {
            throw std::runtime_error("index out of range for origClOrdID [E104]");
        }

        char val;
        std::memcpy(&val, m_buffer + m_offset + 47 + (index * 1), sizeof(char));
        return (val);
    }

    OrderCancelReplaceRequest &origClOrdID(const std::uint64_t index, const char value)
    {
        if (index >= 20)
        {
            throw std::runtime_error("index out of range for origClOrdID [E105]");
        }

        char val = (value);
        std::memcpy(m_buffer + m_offset + 47 + (index * 1), &val, sizeof(char));
        return *this;
    }

    std::uint64_t getOrigClOrdID(char *const dst, const std::uint64_t length) const
    {
        if (length > 20)
        {
            throw std::runtime_error("length too large for getOrigClOrdID [E106]");
        }

        std::memcpy(dst, m_buffer + m_offset + 47, sizeof(char) * static_cast<size_t>(length));
        return length;
    }

    OrderCancelReplaceRequest &putOrigClOrdID(const char *const src) SBE_NOEXCEPT
    {
        std::memcpy(m_buffer + m_offset + 47, src, sizeof(char) * 20);
        return *this;
    }

    SBE_NODISCARD std::string getOrigClOrdIDAsString() const
    {
        const char *buffer = m_buffer + m_offset + 47;
        size_t length = 0;

        for (; length < 20 && *(buffer + length) != '\0'; ++length);
        std::string result(buffer, length);

        return result;
    }

    std::string getOrigClOrdIDAsJsonEscapedString()
    {
        std::ostringstream oss;
        std::string s = getOrigClOrdIDAsString();

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
    SBE_NODISCARD std::string_view getOrigClOrdIDAsStringView() const SBE_NOEXCEPT
    {
        const char *buffer = m_buffer + m_offset + 47;
        size_t length = 0;

        for (; length < 20 && *(buffer + length) != '\0'; ++length);
        std::string_view result(buffer, length);

        return result;
    }
    #endif

    #if __cplusplus >= 201703L
    OrderCancelReplaceRequest &putOrigClOrdID(const std::string_view str)
    {
        const size_t srcLength = str.length();
        if (srcLength > 20)
        {
            throw std::runtime_error("string too large for putOrigClOrdID [E106]");
        }

        std::memcpy(m_buffer + m_offset + 47, str.data(), srcLength);
        for (size_t start = srcLength; start < 20; ++start)
        {
            m_buffer[m_offset + 47 + start] = 0;
        }

        return *this;
    }
    #else
    OrderCancelReplaceRequest &putOrigClOrdID(const std::string& str)
    {
        const size_t srcLength = str.length();
        if (srcLength > 20)
        {
            throw std::runtime_error("string too large for putOrigClOrdID [E106]");
        }

        std::memcpy(m_buffer + m_offset + 47, str.c_str(), srcLength);
        for (size_t start = srcLength; start < 20; ++start)
        {
            m_buffer[m_offset + 47 + start] = 0;
        }

        return *this;
    }
    #endif

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
        return 67;
    }

private:
    OptionalPrice m_price;

public:
    SBE_NODISCARD OptionalPrice &price()
    {
        m_price.wrap(m_buffer, m_offset + 67, m_actingVersion, m_bufferLength);
        return m_price;
    }

    SBE_NODISCARD static const char * SideMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::SEMANTIC_TYPE: return "char";
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
        return 76;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t sideEncodingLength() SBE_NOEXCEPT
    {
        return 1;
    }

    SBE_NODISCARD Side::Value side() const SBE_NOEXCEPT
    {
        char val;
        std::memcpy(&val, m_buffer + m_offset + 76, sizeof(char));
        return Side::get((val));
    }

    OrderCancelReplaceRequest &side(const Side::Value value) SBE_NOEXCEPT
    {
        char val = (value);
        std::memcpy(m_buffer + m_offset + 76, &val, sizeof(char));
        return *this;
    }

    SBE_NODISCARD static const char * SymbolMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::SEMANTIC_TYPE: return "String";
            case MetaAttribute::PRESENCE: return "required";
            default: return "";
        }
    }

    static SBE_CONSTEXPR std::uint16_t symbolId() SBE_NOEXCEPT
    {
        return 55;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t symbolSinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool symbolInActingVersion() SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= symbolSinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t symbolEncodingOffset() SBE_NOEXCEPT
    {
        return 77;
    }

    static SBE_CONSTEXPR char symbolNullValue() SBE_NOEXCEPT
    {
        return (char)0;
    }

    static SBE_CONSTEXPR char symbolMinValue() SBE_NOEXCEPT
    {
        return (char)32;
    }

    static SBE_CONSTEXPR char symbolMaxValue() SBE_NOEXCEPT
    {
        return (char)126;
    }

    static SBE_CONSTEXPR std::size_t symbolEncodingLength() SBE_NOEXCEPT
    {
        return 6;
    }

    static SBE_CONSTEXPR std::uint64_t symbolLength() SBE_NOEXCEPT
    {
        return 6;
    }

    SBE_NODISCARD const char *symbol() const SBE_NOEXCEPT
    {
        return m_buffer + m_offset + 77;
    }

    SBE_NODISCARD char *symbol() SBE_NOEXCEPT
    {
        return m_buffer + m_offset + 77;
    }

    SBE_NODISCARD char symbol(const std::uint64_t index) const
    {
        if (index >= 6)
        {
            throw std::runtime_error("index out of range for symbol [E104]");
        }

        char val;
        std::memcpy(&val, m_buffer + m_offset + 77 + (index * 1), sizeof(char));
        return (val);
    }

    OrderCancelReplaceRequest &symbol(const std::uint64_t index, const char value)
    {
        if (index >= 6)
        {
            throw std::runtime_error("index out of range for symbol [E105]");
        }

        char val = (value);
        std::memcpy(m_buffer + m_offset + 77 + (index * 1), &val, sizeof(char));
        return *this;
    }

    std::uint64_t getSymbol(char *const dst, const std::uint64_t length) const
    {
        if (length > 6)
        {
            throw std::runtime_error("length too large for getSymbol [E106]");
        }

        std::memcpy(dst, m_buffer + m_offset + 77, sizeof(char) * static_cast<size_t>(length));
        return length;
    }

    OrderCancelReplaceRequest &putSymbol(const char *const src) SBE_NOEXCEPT
    {
        std::memcpy(m_buffer + m_offset + 77, src, sizeof(char) * 6);
        return *this;
    }

    SBE_NODISCARD std::string getSymbolAsString() const
    {
        const char *buffer = m_buffer + m_offset + 77;
        size_t length = 0;

        for (; length < 6 && *(buffer + length) != '\0'; ++length);
        std::string result(buffer, length);

        return result;
    }

    std::string getSymbolAsJsonEscapedString()
    {
        std::ostringstream oss;
        std::string s = getSymbolAsString();

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
    SBE_NODISCARD std::string_view getSymbolAsStringView() const SBE_NOEXCEPT
    {
        const char *buffer = m_buffer + m_offset + 77;
        size_t length = 0;

        for (; length < 6 && *(buffer + length) != '\0'; ++length);
        std::string_view result(buffer, length);

        return result;
    }
    #endif

    #if __cplusplus >= 201703L
    OrderCancelReplaceRequest &putSymbol(const std::string_view str)
    {
        const size_t srcLength = str.length();
        if (srcLength > 6)
        {
            throw std::runtime_error("string too large for putSymbol [E106]");
        }

        std::memcpy(m_buffer + m_offset + 77, str.data(), srcLength);
        for (size_t start = srcLength; start < 6; ++start)
        {
            m_buffer[m_offset + 77 + start] = 0;
        }

        return *this;
    }
    #else
    OrderCancelReplaceRequest &putSymbol(const std::string& str)
    {
        const size_t srcLength = str.length();
        if (srcLength > 6)
        {
            throw std::runtime_error("string too large for putSymbol [E106]");
        }

        std::memcpy(m_buffer + m_offset + 77, str.c_str(), srcLength);
        for (size_t start = srcLength; start < 6; ++start)
        {
            m_buffer[m_offset + 77 + start] = 0;
        }

        return *this;
    }
    #endif

    SBE_NODISCARD static const char * TestMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::SEMANTIC_TYPE: return "String";
            case MetaAttribute::PRESENCE: return "required";
            default: return "";
        }
    }

    static SBE_CONSTEXPR std::uint16_t testId() SBE_NOEXCEPT
    {
        return 58;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t testSinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool testInActingVersion() SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= testSinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t testEncodingOffset() SBE_NOEXCEPT
    {
        return 83;
    }

    static SBE_CONSTEXPR char testNullValue() SBE_NOEXCEPT
    {
        return (char)0;
    }

    static SBE_CONSTEXPR char testMinValue() SBE_NOEXCEPT
    {
        return (char)32;
    }

    static SBE_CONSTEXPR char testMaxValue() SBE_NOEXCEPT
    {
        return (char)126;
    }

    static SBE_CONSTEXPR std::size_t testEncodingLength() SBE_NOEXCEPT
    {
        return 18;
    }

    static SBE_CONSTEXPR std::uint64_t testLength() SBE_NOEXCEPT
    {
        return 18;
    }

    SBE_NODISCARD const char *test() const SBE_NOEXCEPT
    {
        return m_buffer + m_offset + 83;
    }

    SBE_NODISCARD char *test() SBE_NOEXCEPT
    {
        return m_buffer + m_offset + 83;
    }

    SBE_NODISCARD char test(const std::uint64_t index) const
    {
        if (index >= 18)
        {
            throw std::runtime_error("index out of range for test [E104]");
        }

        char val;
        std::memcpy(&val, m_buffer + m_offset + 83 + (index * 1), sizeof(char));
        return (val);
    }

    OrderCancelReplaceRequest &test(const std::uint64_t index, const char value)
    {
        if (index >= 18)
        {
            throw std::runtime_error("index out of range for test [E105]");
        }

        char val = (value);
        std::memcpy(m_buffer + m_offset + 83 + (index * 1), &val, sizeof(char));
        return *this;
    }

    std::uint64_t getTest(char *const dst, const std::uint64_t length) const
    {
        if (length > 18)
        {
            throw std::runtime_error("length too large for getTest [E106]");
        }

        std::memcpy(dst, m_buffer + m_offset + 83, sizeof(char) * static_cast<size_t>(length));
        return length;
    }

    OrderCancelReplaceRequest &putTest(const char *const src) SBE_NOEXCEPT
    {
        std::memcpy(m_buffer + m_offset + 83, src, sizeof(char) * 18);
        return *this;
    }

    SBE_NODISCARD std::string getTestAsString() const
    {
        const char *buffer = m_buffer + m_offset + 83;
        size_t length = 0;

        for (; length < 18 && *(buffer + length) != '\0'; ++length);
        std::string result(buffer, length);

        return result;
    }

    std::string getTestAsJsonEscapedString()
    {
        std::ostringstream oss;
        std::string s = getTestAsString();

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
    SBE_NODISCARD std::string_view getTestAsStringView() const SBE_NOEXCEPT
    {
        const char *buffer = m_buffer + m_offset + 83;
        size_t length = 0;

        for (; length < 18 && *(buffer + length) != '\0'; ++length);
        std::string_view result(buffer, length);

        return result;
    }
    #endif

    #if __cplusplus >= 201703L
    OrderCancelReplaceRequest &putTest(const std::string_view str)
    {
        const size_t srcLength = str.length();
        if (srcLength > 18)
        {
            throw std::runtime_error("string too large for putTest [E106]");
        }

        std::memcpy(m_buffer + m_offset + 83, str.data(), srcLength);
        for (size_t start = srcLength; start < 18; ++start)
        {
            m_buffer[m_offset + 83 + start] = 0;
        }

        return *this;
    }
    #else
    OrderCancelReplaceRequest &putTest(const std::string& str)
    {
        const size_t srcLength = str.length();
        if (srcLength > 18)
        {
            throw std::runtime_error("string too large for putTest [E106]");
        }

        std::memcpy(m_buffer + m_offset + 83, str.c_str(), srcLength);
        for (size_t start = srcLength; start < 18; ++start)
        {
            m_buffer[m_offset + 83 + start] = 0;
        }

        return *this;
    }
    #endif

    SBE_NODISCARD static const char * TimeInForceMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::SEMANTIC_TYPE: return "char";
            case MetaAttribute::PRESENCE: return "optional";
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
        return 101;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t timeInForceEncodingLength() SBE_NOEXCEPT
    {
        return 1;
    }

    SBE_NODISCARD TimeInForce::Value timeInForce() const SBE_NOEXCEPT
    {
        char val;
        std::memcpy(&val, m_buffer + m_offset + 101, sizeof(char));
        return TimeInForce::get((val));
    }

    OrderCancelReplaceRequest &timeInForce(const TimeInForce::Value value) SBE_NOEXCEPT
    {
        char val = (value);
        std::memcpy(m_buffer + m_offset + 101, &val, sizeof(char));
        return *this;
    }

    SBE_NODISCARD static const char * ManualOrderIndicatorMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
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
        return 102;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t manualOrderIndicatorEncodingLength() SBE_NOEXCEPT
    {
        return 1;
    }

    SBE_NODISCARD BooleanType::Value manualOrderIndicator() const SBE_NOEXCEPT
    {
        std::uint8_t val;
        std::memcpy(&val, m_buffer + m_offset + 102, sizeof(std::uint8_t));
        return BooleanType::get((val));
    }

    OrderCancelReplaceRequest &manualOrderIndicator(const BooleanType::Value value) SBE_NOEXCEPT
    {
        std::uint8_t val = (value);
        std::memcpy(m_buffer + m_offset + 102, &val, sizeof(std::uint8_t));
        return *this;
    }

    SBE_NODISCARD static const char * TransactTimeMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::TIME_UNIT: return "nanosecond";
            case MetaAttribute::SEMANTIC_TYPE: return "UTCTimestamp";
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
        return 103;
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
        std::memcpy(&val, m_buffer + m_offset + 103, sizeof(std::uint64_t));
        return SBE_LITTLE_ENDIAN_ENCODE_64(val);
    }

    OrderCancelReplaceRequest &transactTime(const std::uint64_t value) SBE_NOEXCEPT
    {
        std::uint64_t val = SBE_LITTLE_ENDIAN_ENCODE_64(value);
        std::memcpy(m_buffer + m_offset + 103, &val, sizeof(std::uint64_t));
        return *this;
    }

    SBE_NODISCARD static const char * NoAllocsMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::PRESENCE: return "optional";
            default: return "";
        }
    }

    static SBE_CONSTEXPR std::uint16_t noAllocsId() SBE_NOEXCEPT
    {
        return 78;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t noAllocsSinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool noAllocsInActingVersion() SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= noAllocsSinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t noAllocsEncodingOffset() SBE_NOEXCEPT
    {
        return 111;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t noAllocsEncodingLength() SBE_NOEXCEPT
    {
        return 1;
    }

    SBE_NODISCARD NoAllocs::Value noAllocs() const SBE_NOEXCEPT
    {
        char val;
        std::memcpy(&val, m_buffer + m_offset + 111, sizeof(char));
        return NoAllocs::get((val));
    }

    OrderCancelReplaceRequest &noAllocs(const NoAllocs::Value value) SBE_NOEXCEPT
    {
        char val = (value);
        std::memcpy(m_buffer + m_offset + 111, &val, sizeof(char));
        return *this;
    }

    SBE_NODISCARD static const char * AllocAccountMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::SEMANTIC_TYPE: return "String";
            case MetaAttribute::PRESENCE: return "required";
            default: return "";
        }
    }

    static SBE_CONSTEXPR std::uint16_t allocAccountId() SBE_NOEXCEPT
    {
        return 79;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t allocAccountSinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool allocAccountInActingVersion() SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= allocAccountSinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t allocAccountEncodingOffset() SBE_NOEXCEPT
    {
        return 112;
    }

    static SBE_CONSTEXPR char allocAccountNullValue() SBE_NOEXCEPT
    {
        return (char)0;
    }

    static SBE_CONSTEXPR char allocAccountMinValue() SBE_NOEXCEPT
    {
        return (char)32;
    }

    static SBE_CONSTEXPR char allocAccountMaxValue() SBE_NOEXCEPT
    {
        return (char)126;
    }

    static SBE_CONSTEXPR std::size_t allocAccountEncodingLength() SBE_NOEXCEPT
    {
        return 10;
    }

    static SBE_CONSTEXPR std::uint64_t allocAccountLength() SBE_NOEXCEPT
    {
        return 10;
    }

    SBE_NODISCARD const char *allocAccount() const SBE_NOEXCEPT
    {
        return m_buffer + m_offset + 112;
    }

    SBE_NODISCARD char *allocAccount() SBE_NOEXCEPT
    {
        return m_buffer + m_offset + 112;
    }

    SBE_NODISCARD char allocAccount(const std::uint64_t index) const
    {
        if (index >= 10)
        {
            throw std::runtime_error("index out of range for allocAccount [E104]");
        }

        char val;
        std::memcpy(&val, m_buffer + m_offset + 112 + (index * 1), sizeof(char));
        return (val);
    }

    OrderCancelReplaceRequest &allocAccount(const std::uint64_t index, const char value)
    {
        if (index >= 10)
        {
            throw std::runtime_error("index out of range for allocAccount [E105]");
        }

        char val = (value);
        std::memcpy(m_buffer + m_offset + 112 + (index * 1), &val, sizeof(char));
        return *this;
    }

    std::uint64_t getAllocAccount(char *const dst, const std::uint64_t length) const
    {
        if (length > 10)
        {
            throw std::runtime_error("length too large for getAllocAccount [E106]");
        }

        std::memcpy(dst, m_buffer + m_offset + 112, sizeof(char) * static_cast<size_t>(length));
        return length;
    }

    OrderCancelReplaceRequest &putAllocAccount(const char *const src) SBE_NOEXCEPT
    {
        std::memcpy(m_buffer + m_offset + 112, src, sizeof(char) * 10);
        return *this;
    }

    SBE_NODISCARD std::string getAllocAccountAsString() const
    {
        const char *buffer = m_buffer + m_offset + 112;
        size_t length = 0;

        for (; length < 10 && *(buffer + length) != '\0'; ++length);
        std::string result(buffer, length);

        return result;
    }

    std::string getAllocAccountAsJsonEscapedString()
    {
        std::ostringstream oss;
        std::string s = getAllocAccountAsString();

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
    SBE_NODISCARD std::string_view getAllocAccountAsStringView() const SBE_NOEXCEPT
    {
        const char *buffer = m_buffer + m_offset + 112;
        size_t length = 0;

        for (; length < 10 && *(buffer + length) != '\0'; ++length);
        std::string_view result(buffer, length);

        return result;
    }
    #endif

    #if __cplusplus >= 201703L
    OrderCancelReplaceRequest &putAllocAccount(const std::string_view str)
    {
        const size_t srcLength = str.length();
        if (srcLength > 10)
        {
            throw std::runtime_error("string too large for putAllocAccount [E106]");
        }

        std::memcpy(m_buffer + m_offset + 112, str.data(), srcLength);
        for (size_t start = srcLength; start < 10; ++start)
        {
            m_buffer[m_offset + 112 + start] = 0;
        }

        return *this;
    }
    #else
    OrderCancelReplaceRequest &putAllocAccount(const std::string& str)
    {
        const size_t srcLength = str.length();
        if (srcLength > 10)
        {
            throw std::runtime_error("string too large for putAllocAccount [E106]");
        }

        std::memcpy(m_buffer + m_offset + 112, str.c_str(), srcLength);
        for (size_t start = srcLength; start < 10; ++start)
        {
            m_buffer[m_offset + 112 + start] = 0;
        }

        return *this;
    }
    #endif

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
        return 122;
    }

private:
    OptionalPrice m_stopPx;

public:
    SBE_NODISCARD OptionalPrice &stopPx()
    {
        m_stopPx.wrap(m_buffer, m_offset + 122, m_actingVersion, m_bufferLength);
        return m_stopPx;
    }

    SBE_NODISCARD static const char * SecurityDescMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::SEMANTIC_TYPE: return "String";
            case MetaAttribute::PRESENCE: return "required";
            default: return "";
        }
    }

    static SBE_CONSTEXPR std::uint16_t securityDescId() SBE_NOEXCEPT
    {
        return 107;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t securityDescSinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool securityDescInActingVersion() SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= securityDescSinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t securityDescEncodingOffset() SBE_NOEXCEPT
    {
        return 131;
    }

    static SBE_CONSTEXPR char securityDescNullValue() SBE_NOEXCEPT
    {
        return (char)0;
    }

    static SBE_CONSTEXPR char securityDescMinValue() SBE_NOEXCEPT
    {
        return (char)32;
    }

    static SBE_CONSTEXPR char securityDescMaxValue() SBE_NOEXCEPT
    {
        return (char)126;
    }

    static SBE_CONSTEXPR std::size_t securityDescEncodingLength() SBE_NOEXCEPT
    {
        return 20;
    }

    static SBE_CONSTEXPR std::uint64_t securityDescLength() SBE_NOEXCEPT
    {
        return 20;
    }

    SBE_NODISCARD const char *securityDesc() const SBE_NOEXCEPT
    {
        return m_buffer + m_offset + 131;
    }

    SBE_NODISCARD char *securityDesc() SBE_NOEXCEPT
    {
        return m_buffer + m_offset + 131;
    }

    SBE_NODISCARD char securityDesc(const std::uint64_t index) const
    {
        if (index >= 20)
        {
            throw std::runtime_error("index out of range for securityDesc [E104]");
        }

        char val;
        std::memcpy(&val, m_buffer + m_offset + 131 + (index * 1), sizeof(char));
        return (val);
    }

    OrderCancelReplaceRequest &securityDesc(const std::uint64_t index, const char value)
    {
        if (index >= 20)
        {
            throw std::runtime_error("index out of range for securityDesc [E105]");
        }

        char val = (value);
        std::memcpy(m_buffer + m_offset + 131 + (index * 1), &val, sizeof(char));
        return *this;
    }

    std::uint64_t getSecurityDesc(char *const dst, const std::uint64_t length) const
    {
        if (length > 20)
        {
            throw std::runtime_error("length too large for getSecurityDesc [E106]");
        }

        std::memcpy(dst, m_buffer + m_offset + 131, sizeof(char) * static_cast<size_t>(length));
        return length;
    }

    OrderCancelReplaceRequest &putSecurityDesc(const char *const src) SBE_NOEXCEPT
    {
        std::memcpy(m_buffer + m_offset + 131, src, sizeof(char) * 20);
        return *this;
    }

    SBE_NODISCARD std::string getSecurityDescAsString() const
    {
        const char *buffer = m_buffer + m_offset + 131;
        size_t length = 0;

        for (; length < 20 && *(buffer + length) != '\0'; ++length);
        std::string result(buffer, length);

        return result;
    }

    std::string getSecurityDescAsJsonEscapedString()
    {
        std::ostringstream oss;
        std::string s = getSecurityDescAsString();

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
    SBE_NODISCARD std::string_view getSecurityDescAsStringView() const SBE_NOEXCEPT
    {
        const char *buffer = m_buffer + m_offset + 131;
        size_t length = 0;

        for (; length < 20 && *(buffer + length) != '\0'; ++length);
        std::string_view result(buffer, length);

        return result;
    }
    #endif

    #if __cplusplus >= 201703L
    OrderCancelReplaceRequest &putSecurityDesc(const std::string_view str)
    {
        const size_t srcLength = str.length();
        if (srcLength > 20)
        {
            throw std::runtime_error("string too large for putSecurityDesc [E106]");
        }

        std::memcpy(m_buffer + m_offset + 131, str.data(), srcLength);
        for (size_t start = srcLength; start < 20; ++start)
        {
            m_buffer[m_offset + 131 + start] = 0;
        }

        return *this;
    }
    #else
    OrderCancelReplaceRequest &putSecurityDesc(const std::string& str)
    {
        const size_t srcLength = str.length();
        if (srcLength > 20)
        {
            throw std::runtime_error("string too large for putSecurityDesc [E106]");
        }

        std::memcpy(m_buffer + m_offset + 131, str.c_str(), srcLength);
        for (size_t start = srcLength; start < 20; ++start)
        {
            m_buffer[m_offset + 131 + start] = 0;
        }

        return *this;
    }
    #endif

    SBE_NODISCARD static const char * MinQtyMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::SEMANTIC_TYPE: return "Qty";
            case MetaAttribute::PRESENCE: return "required";
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
        return 151;
    }

private:
    IntQty32 m_minQty;

public:
    SBE_NODISCARD IntQty32 &minQty()
    {
        m_minQty.wrap(m_buffer, m_offset + 151, m_actingVersion, m_bufferLength);
        return m_minQty;
    }

    SBE_NODISCARD static const char * SecurityTypeMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::SEMANTIC_TYPE: return "String";
            case MetaAttribute::PRESENCE: return "required";
            default: return "";
        }
    }

    static SBE_CONSTEXPR std::uint16_t securityTypeId() SBE_NOEXCEPT
    {
        return 167;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t securityTypeSinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool securityTypeInActingVersion() SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= securityTypeSinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t securityTypeEncodingOffset() SBE_NOEXCEPT
    {
        return 155;
    }

    static SBE_CONSTEXPR char securityTypeNullValue() SBE_NOEXCEPT
    {
        return (char)0;
    }

    static SBE_CONSTEXPR char securityTypeMinValue() SBE_NOEXCEPT
    {
        return (char)32;
    }

    static SBE_CONSTEXPR char securityTypeMaxValue() SBE_NOEXCEPT
    {
        return (char)126;
    }

    static SBE_CONSTEXPR std::size_t securityTypeEncodingLength() SBE_NOEXCEPT
    {
        return 3;
    }

    static SBE_CONSTEXPR std::uint64_t securityTypeLength() SBE_NOEXCEPT
    {
        return 3;
    }

    SBE_NODISCARD const char *securityType() const SBE_NOEXCEPT
    {
        return m_buffer + m_offset + 155;
    }

    SBE_NODISCARD char *securityType() SBE_NOEXCEPT
    {
        return m_buffer + m_offset + 155;
    }

    SBE_NODISCARD char securityType(const std::uint64_t index) const
    {
        if (index >= 3)
        {
            throw std::runtime_error("index out of range for securityType [E104]");
        }

        char val;
        std::memcpy(&val, m_buffer + m_offset + 155 + (index * 1), sizeof(char));
        return (val);
    }

    OrderCancelReplaceRequest &securityType(const std::uint64_t index, const char value)
    {
        if (index >= 3)
        {
            throw std::runtime_error("index out of range for securityType [E105]");
        }

        char val = (value);
        std::memcpy(m_buffer + m_offset + 155 + (index * 1), &val, sizeof(char));
        return *this;
    }

    std::uint64_t getSecurityType(char *const dst, const std::uint64_t length) const
    {
        if (length > 3)
        {
            throw std::runtime_error("length too large for getSecurityType [E106]");
        }

        std::memcpy(dst, m_buffer + m_offset + 155, sizeof(char) * static_cast<size_t>(length));
        return length;
    }

    OrderCancelReplaceRequest &putSecurityType(const char *const src) SBE_NOEXCEPT
    {
        std::memcpy(m_buffer + m_offset + 155, src, sizeof(char) * 3);
        return *this;
    }

    OrderCancelReplaceRequest &putSecurityType(
        const char value0,
        const char value1,
        const char value2) SBE_NOEXCEPT
    {
        char val0 = (value0);
        std::memcpy(m_buffer + m_offset + 155, &val0, sizeof(char));
        char val1 = (value1);
        std::memcpy(m_buffer + m_offset + 156, &val1, sizeof(char));
        char val2 = (value2);
        std::memcpy(m_buffer + m_offset + 157, &val2, sizeof(char));

        return *this;
    }

    SBE_NODISCARD std::string getSecurityTypeAsString() const
    {
        const char *buffer = m_buffer + m_offset + 155;
        size_t length = 0;

        for (; length < 3 && *(buffer + length) != '\0'; ++length);
        std::string result(buffer, length);

        return result;
    }

    std::string getSecurityTypeAsJsonEscapedString()
    {
        std::ostringstream oss;
        std::string s = getSecurityTypeAsString();

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
    SBE_NODISCARD std::string_view getSecurityTypeAsStringView() const SBE_NOEXCEPT
    {
        const char *buffer = m_buffer + m_offset + 155;
        size_t length = 0;

        for (; length < 3 && *(buffer + length) != '\0'; ++length);
        std::string_view result(buffer, length);

        return result;
    }
    #endif

    #if __cplusplus >= 201703L
    OrderCancelReplaceRequest &putSecurityType(const std::string_view str)
    {
        const size_t srcLength = str.length();
        if (srcLength > 3)
        {
            throw std::runtime_error("string too large for putSecurityType [E106]");
        }

        std::memcpy(m_buffer + m_offset + 155, str.data(), srcLength);
        for (size_t start = srcLength; start < 3; ++start)
        {
            m_buffer[m_offset + 155 + start] = 0;
        }

        return *this;
    }
    #else
    OrderCancelReplaceRequest &putSecurityType(const std::string& str)
    {
        const size_t srcLength = str.length();
        if (srcLength > 3)
        {
            throw std::runtime_error("string too large for putSecurityType [E106]");
        }

        std::memcpy(m_buffer + m_offset + 155, str.c_str(), srcLength);
        for (size_t start = srcLength; start < 3; ++start)
        {
            m_buffer[m_offset + 155 + start] = 0;
        }

        return *this;
    }
    #endif

    SBE_NODISCARD static const char * CustomerOrFirmMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::PRESENCE: return "required";
            default: return "";
        }
    }

    static SBE_CONSTEXPR std::uint16_t customerOrFirmId() SBE_NOEXCEPT
    {
        return 204;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t customerOrFirmSinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool customerOrFirmInActingVersion() SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= customerOrFirmSinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t customerOrFirmEncodingOffset() SBE_NOEXCEPT
    {
        return 158;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t customerOrFirmEncodingLength() SBE_NOEXCEPT
    {
        return 1;
    }

    SBE_NODISCARD CustomerOrFirm::Value customerOrFirm() const SBE_NOEXCEPT
    {
        std::uint8_t val;
        std::memcpy(&val, m_buffer + m_offset + 158, sizeof(std::uint8_t));
        return CustomerOrFirm::get((val));
    }

    OrderCancelReplaceRequest &customerOrFirm(const CustomerOrFirm::Value value) SBE_NOEXCEPT
    {
        std::uint8_t val = (value);
        std::memcpy(m_buffer + m_offset + 158, &val, sizeof(std::uint8_t));
        return *this;
    }

    SBE_NODISCARD static const char * MaxShowMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::SEMANTIC_TYPE: return "Qty";
            case MetaAttribute::PRESENCE: return "required";
            default: return "";
        }
    }

    static SBE_CONSTEXPR std::uint16_t maxShowId() SBE_NOEXCEPT
    {
        return 210;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t maxShowSinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool maxShowInActingVersion() SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= maxShowSinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t maxShowEncodingOffset() SBE_NOEXCEPT
    {
        return 159;
    }

private:
    IntQty32 m_maxShow;

public:
    SBE_NODISCARD IntQty32 &maxShow()
    {
        m_maxShow.wrap(m_buffer, m_offset + 159, m_actingVersion, m_bufferLength);
        return m_maxShow;
    }

    SBE_NODISCARD static const char * ExpireDateMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::PRESENCE: return "required";
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
        return 163;
    }

    static SBE_CONSTEXPR std::uint16_t expireDateNullValue() SBE_NOEXCEPT
    {
        return SBE_NULLVALUE_UINT16;
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
        std::memcpy(&val, m_buffer + m_offset + 163, sizeof(std::uint16_t));
        return SBE_LITTLE_ENDIAN_ENCODE_16(val);
    }

    OrderCancelReplaceRequest &expireDate(const std::uint16_t value) SBE_NOEXCEPT
    {
        std::uint16_t val = SBE_LITTLE_ENDIAN_ENCODE_16(value);
        std::memcpy(m_buffer + m_offset + 163, &val, sizeof(std::uint16_t));
        return *this;
    }

    SBE_NODISCARD static const char * SelfMatchPreventionIDMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::SEMANTIC_TYPE: return "String";
            case MetaAttribute::PRESENCE: return "required";
            default: return "";
        }
    }

    static SBE_CONSTEXPR std::uint16_t selfMatchPreventionIDId() SBE_NOEXCEPT
    {
        return 7928;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t selfMatchPreventionIDSinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool selfMatchPreventionIDInActingVersion() SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= selfMatchPreventionIDSinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t selfMatchPreventionIDEncodingOffset() SBE_NOEXCEPT
    {
        return 165;
    }

    static SBE_CONSTEXPR char selfMatchPreventionIDNullValue() SBE_NOEXCEPT
    {
        return (char)0;
    }

    static SBE_CONSTEXPR char selfMatchPreventionIDMinValue() SBE_NOEXCEPT
    {
        return (char)32;
    }

    static SBE_CONSTEXPR char selfMatchPreventionIDMaxValue() SBE_NOEXCEPT
    {
        return (char)126;
    }

    static SBE_CONSTEXPR std::size_t selfMatchPreventionIDEncodingLength() SBE_NOEXCEPT
    {
        return 12;
    }

    static SBE_CONSTEXPR std::uint64_t selfMatchPreventionIDLength() SBE_NOEXCEPT
    {
        return 12;
    }

    SBE_NODISCARD const char *selfMatchPreventionID() const SBE_NOEXCEPT
    {
        return m_buffer + m_offset + 165;
    }

    SBE_NODISCARD char *selfMatchPreventionID() SBE_NOEXCEPT
    {
        return m_buffer + m_offset + 165;
    }

    SBE_NODISCARD char selfMatchPreventionID(const std::uint64_t index) const
    {
        if (index >= 12)
        {
            throw std::runtime_error("index out of range for selfMatchPreventionID [E104]");
        }

        char val;
        std::memcpy(&val, m_buffer + m_offset + 165 + (index * 1), sizeof(char));
        return (val);
    }

    OrderCancelReplaceRequest &selfMatchPreventionID(const std::uint64_t index, const char value)
    {
        if (index >= 12)
        {
            throw std::runtime_error("index out of range for selfMatchPreventionID [E105]");
        }

        char val = (value);
        std::memcpy(m_buffer + m_offset + 165 + (index * 1), &val, sizeof(char));
        return *this;
    }

    std::uint64_t getSelfMatchPreventionID(char *const dst, const std::uint64_t length) const
    {
        if (length > 12)
        {
            throw std::runtime_error("length too large for getSelfMatchPreventionID [E106]");
        }

        std::memcpy(dst, m_buffer + m_offset + 165, sizeof(char) * static_cast<size_t>(length));
        return length;
    }

    OrderCancelReplaceRequest &putSelfMatchPreventionID(const char *const src) SBE_NOEXCEPT
    {
        std::memcpy(m_buffer + m_offset + 165, src, sizeof(char) * 12);
        return *this;
    }

    SBE_NODISCARD std::string getSelfMatchPreventionIDAsString() const
    {
        const char *buffer = m_buffer + m_offset + 165;
        size_t length = 0;

        for (; length < 12 && *(buffer + length) != '\0'; ++length);
        std::string result(buffer, length);

        return result;
    }

    std::string getSelfMatchPreventionIDAsJsonEscapedString()
    {
        std::ostringstream oss;
        std::string s = getSelfMatchPreventionIDAsString();

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
    SBE_NODISCARD std::string_view getSelfMatchPreventionIDAsStringView() const SBE_NOEXCEPT
    {
        const char *buffer = m_buffer + m_offset + 165;
        size_t length = 0;

        for (; length < 12 && *(buffer + length) != '\0'; ++length);
        std::string_view result(buffer, length);

        return result;
    }
    #endif

    #if __cplusplus >= 201703L
    OrderCancelReplaceRequest &putSelfMatchPreventionID(const std::string_view str)
    {
        const size_t srcLength = str.length();
        if (srcLength > 12)
        {
            throw std::runtime_error("string too large for putSelfMatchPreventionID [E106]");
        }

        std::memcpy(m_buffer + m_offset + 165, str.data(), srcLength);
        for (size_t start = srcLength; start < 12; ++start)
        {
            m_buffer[m_offset + 165 + start] = 0;
        }

        return *this;
    }
    #else
    OrderCancelReplaceRequest &putSelfMatchPreventionID(const std::string& str)
    {
        const size_t srcLength = str.length();
        if (srcLength > 12)
        {
            throw std::runtime_error("string too large for putSelfMatchPreventionID [E106]");
        }

        std::memcpy(m_buffer + m_offset + 165, str.c_str(), srcLength);
        for (size_t start = srcLength; start < 12; ++start)
        {
            m_buffer[m_offset + 165 + start] = 0;
        }

        return *this;
    }
    #endif

    SBE_NODISCARD static const char * CtiCodeMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::PRESENCE: return "required";
            default: return "";
        }
    }

    static SBE_CONSTEXPR std::uint16_t ctiCodeId() SBE_NOEXCEPT
    {
        return 9702;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t ctiCodeSinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool ctiCodeInActingVersion() SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= ctiCodeSinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t ctiCodeEncodingOffset() SBE_NOEXCEPT
    {
        return 177;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t ctiCodeEncodingLength() SBE_NOEXCEPT
    {
        return 1;
    }

    SBE_NODISCARD CtiCode::Value ctiCode() const SBE_NOEXCEPT
    {
        char val;
        std::memcpy(&val, m_buffer + m_offset + 177, sizeof(char));
        return CtiCode::get((val));
    }

    OrderCancelReplaceRequest &ctiCode(const CtiCode::Value value) SBE_NOEXCEPT
    {
        char val = (value);
        std::memcpy(m_buffer + m_offset + 177, &val, sizeof(char));
        return *this;
    }

    SBE_NODISCARD static const char * GiveUpFirmMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::SEMANTIC_TYPE: return "String";
            case MetaAttribute::PRESENCE: return "required";
            default: return "";
        }
    }

    static SBE_CONSTEXPR std::uint16_t giveUpFirmId() SBE_NOEXCEPT
    {
        return 9707;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t giveUpFirmSinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool giveUpFirmInActingVersion() SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= giveUpFirmSinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t giveUpFirmEncodingOffset() SBE_NOEXCEPT
    {
        return 178;
    }

    static SBE_CONSTEXPR char giveUpFirmNullValue() SBE_NOEXCEPT
    {
        return (char)0;
    }

    static SBE_CONSTEXPR char giveUpFirmMinValue() SBE_NOEXCEPT
    {
        return (char)32;
    }

    static SBE_CONSTEXPR char giveUpFirmMaxValue() SBE_NOEXCEPT
    {
        return (char)126;
    }

    static SBE_CONSTEXPR std::size_t giveUpFirmEncodingLength() SBE_NOEXCEPT
    {
        return 3;
    }

    static SBE_CONSTEXPR std::uint64_t giveUpFirmLength() SBE_NOEXCEPT
    {
        return 3;
    }

    SBE_NODISCARD const char *giveUpFirm() const SBE_NOEXCEPT
    {
        return m_buffer + m_offset + 178;
    }

    SBE_NODISCARD char *giveUpFirm() SBE_NOEXCEPT
    {
        return m_buffer + m_offset + 178;
    }

    SBE_NODISCARD char giveUpFirm(const std::uint64_t index) const
    {
        if (index >= 3)
        {
            throw std::runtime_error("index out of range for giveUpFirm [E104]");
        }

        char val;
        std::memcpy(&val, m_buffer + m_offset + 178 + (index * 1), sizeof(char));
        return (val);
    }

    OrderCancelReplaceRequest &giveUpFirm(const std::uint64_t index, const char value)
    {
        if (index >= 3)
        {
            throw std::runtime_error("index out of range for giveUpFirm [E105]");
        }

        char val = (value);
        std::memcpy(m_buffer + m_offset + 178 + (index * 1), &val, sizeof(char));
        return *this;
    }

    std::uint64_t getGiveUpFirm(char *const dst, const std::uint64_t length) const
    {
        if (length > 3)
        {
            throw std::runtime_error("length too large for getGiveUpFirm [E106]");
        }

        std::memcpy(dst, m_buffer + m_offset + 178, sizeof(char) * static_cast<size_t>(length));
        return length;
    }

    OrderCancelReplaceRequest &putGiveUpFirm(const char *const src) SBE_NOEXCEPT
    {
        std::memcpy(m_buffer + m_offset + 178, src, sizeof(char) * 3);
        return *this;
    }

    OrderCancelReplaceRequest &putGiveUpFirm(
        const char value0,
        const char value1,
        const char value2) SBE_NOEXCEPT
    {
        char val0 = (value0);
        std::memcpy(m_buffer + m_offset + 178, &val0, sizeof(char));
        char val1 = (value1);
        std::memcpy(m_buffer + m_offset + 179, &val1, sizeof(char));
        char val2 = (value2);
        std::memcpy(m_buffer + m_offset + 180, &val2, sizeof(char));

        return *this;
    }

    SBE_NODISCARD std::string getGiveUpFirmAsString() const
    {
        const char *buffer = m_buffer + m_offset + 178;
        size_t length = 0;

        for (; length < 3 && *(buffer + length) != '\0'; ++length);
        std::string result(buffer, length);

        return result;
    }

    std::string getGiveUpFirmAsJsonEscapedString()
    {
        std::ostringstream oss;
        std::string s = getGiveUpFirmAsString();

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
    SBE_NODISCARD std::string_view getGiveUpFirmAsStringView() const SBE_NOEXCEPT
    {
        const char *buffer = m_buffer + m_offset + 178;
        size_t length = 0;

        for (; length < 3 && *(buffer + length) != '\0'; ++length);
        std::string_view result(buffer, length);

        return result;
    }
    #endif

    #if __cplusplus >= 201703L
    OrderCancelReplaceRequest &putGiveUpFirm(const std::string_view str)
    {
        const size_t srcLength = str.length();
        if (srcLength > 3)
        {
            throw std::runtime_error("string too large for putGiveUpFirm [E106]");
        }

        std::memcpy(m_buffer + m_offset + 178, str.data(), srcLength);
        for (size_t start = srcLength; start < 3; ++start)
        {
            m_buffer[m_offset + 178 + start] = 0;
        }

        return *this;
    }
    #else
    OrderCancelReplaceRequest &putGiveUpFirm(const std::string& str)
    {
        const size_t srcLength = str.length();
        if (srcLength > 3)
        {
            throw std::runtime_error("string too large for putGiveUpFirm [E106]");
        }

        std::memcpy(m_buffer + m_offset + 178, str.c_str(), srcLength);
        for (size_t start = srcLength; start < 3; ++start)
        {
            m_buffer[m_offset + 178 + start] = 0;
        }

        return *this;
    }
    #endif

    SBE_NODISCARD static const char * CmtaGiveupCDMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::SEMANTIC_TYPE: return "String";
            case MetaAttribute::PRESENCE: return "required";
            default: return "";
        }
    }

    static SBE_CONSTEXPR std::uint16_t cmtaGiveupCDId() SBE_NOEXCEPT
    {
        return 9708;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t cmtaGiveupCDSinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool cmtaGiveupCDInActingVersion() SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= cmtaGiveupCDSinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t cmtaGiveupCDEncodingOffset() SBE_NOEXCEPT
    {
        return 181;
    }

    static SBE_CONSTEXPR char cmtaGiveupCDNullValue() SBE_NOEXCEPT
    {
        return (char)0;
    }

    static SBE_CONSTEXPR char cmtaGiveupCDMinValue() SBE_NOEXCEPT
    {
        return (char)32;
    }

    static SBE_CONSTEXPR char cmtaGiveupCDMaxValue() SBE_NOEXCEPT
    {
        return (char)126;
    }

    static SBE_CONSTEXPR std::size_t cmtaGiveupCDEncodingLength() SBE_NOEXCEPT
    {
        return 2;
    }

    static SBE_CONSTEXPR std::uint64_t cmtaGiveupCDLength() SBE_NOEXCEPT
    {
        return 2;
    }

    SBE_NODISCARD const char *cmtaGiveupCD() const SBE_NOEXCEPT
    {
        return m_buffer + m_offset + 181;
    }

    SBE_NODISCARD char *cmtaGiveupCD() SBE_NOEXCEPT
    {
        return m_buffer + m_offset + 181;
    }

    SBE_NODISCARD char cmtaGiveupCD(const std::uint64_t index) const
    {
        if (index >= 2)
        {
            throw std::runtime_error("index out of range for cmtaGiveupCD [E104]");
        }

        char val;
        std::memcpy(&val, m_buffer + m_offset + 181 + (index * 1), sizeof(char));
        return (val);
    }

    OrderCancelReplaceRequest &cmtaGiveupCD(const std::uint64_t index, const char value)
    {
        if (index >= 2)
        {
            throw std::runtime_error("index out of range for cmtaGiveupCD [E105]");
        }

        char val = (value);
        std::memcpy(m_buffer + m_offset + 181 + (index * 1), &val, sizeof(char));
        return *this;
    }

    std::uint64_t getCmtaGiveupCD(char *const dst, const std::uint64_t length) const
    {
        if (length > 2)
        {
            throw std::runtime_error("length too large for getCmtaGiveupCD [E106]");
        }

        std::memcpy(dst, m_buffer + m_offset + 181, sizeof(char) * static_cast<size_t>(length));
        return length;
    }

    OrderCancelReplaceRequest &putCmtaGiveupCD(const char *const src) SBE_NOEXCEPT
    {
        std::memcpy(m_buffer + m_offset + 181, src, sizeof(char) * 2);
        return *this;
    }

    OrderCancelReplaceRequest &putCmtaGiveupCD(
        const char value0,
        const char value1) SBE_NOEXCEPT
    {
        char val0 = (value0);
        std::memcpy(m_buffer + m_offset + 181, &val0, sizeof(char));
        char val1 = (value1);
        std::memcpy(m_buffer + m_offset + 182, &val1, sizeof(char));

        return *this;
    }

    SBE_NODISCARD std::string getCmtaGiveupCDAsString() const
    {
        const char *buffer = m_buffer + m_offset + 181;
        size_t length = 0;

        for (; length < 2 && *(buffer + length) != '\0'; ++length);
        std::string result(buffer, length);

        return result;
    }

    std::string getCmtaGiveupCDAsJsonEscapedString()
    {
        std::ostringstream oss;
        std::string s = getCmtaGiveupCDAsString();

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
    SBE_NODISCARD std::string_view getCmtaGiveupCDAsStringView() const SBE_NOEXCEPT
    {
        const char *buffer = m_buffer + m_offset + 181;
        size_t length = 0;

        for (; length < 2 && *(buffer + length) != '\0'; ++length);
        std::string_view result(buffer, length);

        return result;
    }
    #endif

    #if __cplusplus >= 201703L
    OrderCancelReplaceRequest &putCmtaGiveupCD(const std::string_view str)
    {
        const size_t srcLength = str.length();
        if (srcLength > 2)
        {
            throw std::runtime_error("string too large for putCmtaGiveupCD [E106]");
        }

        std::memcpy(m_buffer + m_offset + 181, str.data(), srcLength);
        for (size_t start = srcLength; start < 2; ++start)
        {
            m_buffer[m_offset + 181 + start] = 0;
        }

        return *this;
    }
    #else
    OrderCancelReplaceRequest &putCmtaGiveupCD(const std::string& str)
    {
        const size_t srcLength = str.length();
        if (srcLength > 2)
        {
            throw std::runtime_error("string too large for putCmtaGiveupCD [E106]");
        }

        std::memcpy(m_buffer + m_offset + 181, str.c_str(), srcLength);
        for (size_t start = srcLength; start < 2; ++start)
        {
            m_buffer[m_offset + 181 + start] = 0;
        }

        return *this;
    }
    #endif

    SBE_NODISCARD static const char * CorrelationClOrdIDMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::SEMANTIC_TYPE: return "String";
            case MetaAttribute::PRESENCE: return "required";
            default: return "";
        }
    }

    static SBE_CONSTEXPR std::uint16_t correlationClOrdIDId() SBE_NOEXCEPT
    {
        return 9717;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t correlationClOrdIDSinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool correlationClOrdIDInActingVersion() SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= correlationClOrdIDSinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t correlationClOrdIDEncodingOffset() SBE_NOEXCEPT
    {
        return 183;
    }

    static SBE_CONSTEXPR char correlationClOrdIDNullValue() SBE_NOEXCEPT
    {
        return (char)0;
    }

    static SBE_CONSTEXPR char correlationClOrdIDMinValue() SBE_NOEXCEPT
    {
        return (char)32;
    }

    static SBE_CONSTEXPR char correlationClOrdIDMaxValue() SBE_NOEXCEPT
    {
        return (char)126;
    }

    static SBE_CONSTEXPR std::size_t correlationClOrdIDEncodingLength() SBE_NOEXCEPT
    {
        return 20;
    }

    static SBE_CONSTEXPR std::uint64_t correlationClOrdIDLength() SBE_NOEXCEPT
    {
        return 20;
    }

    SBE_NODISCARD const char *correlationClOrdID() const SBE_NOEXCEPT
    {
        return m_buffer + m_offset + 183;
    }

    SBE_NODISCARD char *correlationClOrdID() SBE_NOEXCEPT
    {
        return m_buffer + m_offset + 183;
    }

    SBE_NODISCARD char correlationClOrdID(const std::uint64_t index) const
    {
        if (index >= 20)
        {
            throw std::runtime_error("index out of range for correlationClOrdID [E104]");
        }

        char val;
        std::memcpy(&val, m_buffer + m_offset + 183 + (index * 1), sizeof(char));
        return (val);
    }

    OrderCancelReplaceRequest &correlationClOrdID(const std::uint64_t index, const char value)
    {
        if (index >= 20)
        {
            throw std::runtime_error("index out of range for correlationClOrdID [E105]");
        }

        char val = (value);
        std::memcpy(m_buffer + m_offset + 183 + (index * 1), &val, sizeof(char));
        return *this;
    }

    std::uint64_t getCorrelationClOrdID(char *const dst, const std::uint64_t length) const
    {
        if (length > 20)
        {
            throw std::runtime_error("length too large for getCorrelationClOrdID [E106]");
        }

        std::memcpy(dst, m_buffer + m_offset + 183, sizeof(char) * static_cast<size_t>(length));
        return length;
    }

    OrderCancelReplaceRequest &putCorrelationClOrdID(const char *const src) SBE_NOEXCEPT
    {
        std::memcpy(m_buffer + m_offset + 183, src, sizeof(char) * 20);
        return *this;
    }

    SBE_NODISCARD std::string getCorrelationClOrdIDAsString() const
    {
        const char *buffer = m_buffer + m_offset + 183;
        size_t length = 0;

        for (; length < 20 && *(buffer + length) != '\0'; ++length);
        std::string result(buffer, length);

        return result;
    }

    std::string getCorrelationClOrdIDAsJsonEscapedString()
    {
        std::ostringstream oss;
        std::string s = getCorrelationClOrdIDAsString();

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
    SBE_NODISCARD std::string_view getCorrelationClOrdIDAsStringView() const SBE_NOEXCEPT
    {
        const char *buffer = m_buffer + m_offset + 183;
        size_t length = 0;

        for (; length < 20 && *(buffer + length) != '\0'; ++length);
        std::string_view result(buffer, length);

        return result;
    }
    #endif

    #if __cplusplus >= 201703L
    OrderCancelReplaceRequest &putCorrelationClOrdID(const std::string_view str)
    {
        const size_t srcLength = str.length();
        if (srcLength > 20)
        {
            throw std::runtime_error("string too large for putCorrelationClOrdID [E106]");
        }

        std::memcpy(m_buffer + m_offset + 183, str.data(), srcLength);
        for (size_t start = srcLength; start < 20; ++start)
        {
            m_buffer[m_offset + 183 + start] = 0;
        }

        return *this;
    }
    #else
    OrderCancelReplaceRequest &putCorrelationClOrdID(const std::string& str)
    {
        const size_t srcLength = str.length();
        if (srcLength > 20)
        {
            throw std::runtime_error("string too large for putCorrelationClOrdID [E106]");
        }

        std::memcpy(m_buffer + m_offset + 183, str.c_str(), srcLength);
        for (size_t start = srcLength; start < 20; ++start)
        {
            m_buffer[m_offset + 183 + start] = 0;
        }

        return *this;
    }
    #endif

    SBE_NODISCARD static const char * OFMOverrideMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::PRESENCE: return "required";
            default: return "";
        }
    }

    static SBE_CONSTEXPR std::uint16_t oFMOverrideId() SBE_NOEXCEPT
    {
        return 9768;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t oFMOverrideSinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool oFMOverrideInActingVersion() SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= oFMOverrideSinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t oFMOverrideEncodingOffset() SBE_NOEXCEPT
    {
        return 203;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t oFMOverrideEncodingLength() SBE_NOEXCEPT
    {
        return 1;
    }

    SBE_NODISCARD OFMOverride::Value oFMOverride() const SBE_NOEXCEPT
    {
        char val;
        std::memcpy(&val, m_buffer + m_offset + 203, sizeof(char));
        return OFMOverride::get((val));
    }

    OrderCancelReplaceRequest &oFMOverride(const OFMOverride::Value value) SBE_NOEXCEPT
    {
        char val = (value);
        std::memcpy(m_buffer + m_offset + 203, &val, sizeof(char));
        return *this;
    }

template<typename CharT, typename Traits>
friend std::basic_ostream<CharT, Traits>& operator << (
    std::basic_ostream<CharT, Traits>& builder, OrderCancelReplaceRequest _writer)
{
    OrderCancelReplaceRequest writer(_writer.m_buffer, _writer.m_offset,
        _writer.m_bufferLength, _writer.sbeBlockLength(), _writer.m_actingVersion);
    builder << '{';
    builder << R"("Name": "OrderCancelReplaceRequest", )";
    builder << R"("sbeTemplateId": )";
    builder << writer.sbeTemplateId();
    builder << ", ";

    builder << R"("Account": )";
    builder << '"' <<
        writer.getAccountAsJsonEscapedString().c_str() << '"';

    builder << ", ";
    builder << R"("ClOrdID": )";
    builder << '"' <<
        writer.getClOrdIDAsJsonEscapedString().c_str() << '"';

    builder << ", ";
    builder << R"("OrderID": )";
    builder << +writer.orderID();

    builder << ", ";
    builder << R"("HandInst": )";
    builder << '"' << writer.handInst() << '"';

    builder << ", ";
    builder << R"("OrderQty": )";
    builder << writer.orderQty();

    builder << ", ";
    builder << R"("CustOrderHandlingInst": )";
    builder << '"' << writer.custOrderHandlingInst() << '"';

    builder << ", ";
    builder << R"("OrdType": )";
    builder << '"' << writer.ordType() << '"';

    builder << ", ";
    builder << R"("OrigClOrdID": )";
    builder << '"' <<
        writer.getOrigClOrdIDAsJsonEscapedString().c_str() << '"';

    builder << ", ";
    builder << R"("Price": )";
    builder << writer.price();

    builder << ", ";
    builder << R"("Side": )";
    builder << '"' << writer.side() << '"';

    builder << ", ";
    builder << R"("Symbol": )";
    builder << '"' <<
        writer.getSymbolAsJsonEscapedString().c_str() << '"';

    builder << ", ";
    builder << R"("Test": )";
    builder << '"' <<
        writer.getTestAsJsonEscapedString().c_str() << '"';

    builder << ", ";
    builder << R"("TimeInForce": )";
    builder << '"' << writer.timeInForce() << '"';

    builder << ", ";
    builder << R"("ManualOrderIndicator": )";
    builder << '"' << writer.manualOrderIndicator() << '"';

    builder << ", ";
    builder << R"("TransactTime": )";
    builder << +writer.transactTime();

    builder << ", ";
    builder << R"("NoAllocs": )";
    builder << '"' << writer.noAllocs() << '"';

    builder << ", ";
    builder << R"("AllocAccount": )";
    builder << '"' <<
        writer.getAllocAccountAsJsonEscapedString().c_str() << '"';

    builder << ", ";
    builder << R"("StopPx": )";
    builder << writer.stopPx();

    builder << ", ";
    builder << R"("SecurityDesc": )";
    builder << '"' <<
        writer.getSecurityDescAsJsonEscapedString().c_str() << '"';

    builder << ", ";
    builder << R"("MinQty": )";
    builder << writer.minQty();

    builder << ", ";
    builder << R"("SecurityType": )";
    builder << '"' <<
        writer.getSecurityTypeAsJsonEscapedString().c_str() << '"';

    builder << ", ";
    builder << R"("CustomerOrFirm": )";
    builder << '"' << writer.customerOrFirm() << '"';

    builder << ", ";
    builder << R"("MaxShow": )";
    builder << writer.maxShow();

    builder << ", ";
    builder << R"("ExpireDate": )";
    builder << +writer.expireDate();

    builder << ", ";
    builder << R"("SelfMatchPreventionID": )";
    builder << '"' <<
        writer.getSelfMatchPreventionIDAsJsonEscapedString().c_str() << '"';

    builder << ", ";
    builder << R"("CtiCode": )";
    builder << '"' << writer.ctiCode() << '"';

    builder << ", ";
    builder << R"("GiveUpFirm": )";
    builder << '"' <<
        writer.getGiveUpFirmAsJsonEscapedString().c_str() << '"';

    builder << ", ";
    builder << R"("CmtaGiveupCD": )";
    builder << '"' <<
        writer.getCmtaGiveupCDAsJsonEscapedString().c_str() << '"';

    builder << ", ";
    builder << R"("CorrelationClOrdID": )";
    builder << '"' <<
        writer.getCorrelationClOrdIDAsJsonEscapedString().c_str() << '"';

    builder << ", ";
    builder << R"("OFMOverride": )";
    builder << '"' << writer.oFMOverride() << '"';

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
