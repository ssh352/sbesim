/* Generated SBE (Simple Binary Encoding) message codec */
#ifndef _FIX_MASSQUOTE_H_
#define _FIX_MASSQUOTE_H_

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

class MassQuote
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

    MassQuote() = default;

    MassQuote(
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

    MassQuote(char *buffer, const std::uint64_t bufferLength) :
        MassQuote(buffer, 0, bufferLength, sbeBlockLength(), sbeSchemaVersion())
    {
    }

    MassQuote(
        char *buffer,
        const std::uint64_t bufferLength,
        const std::uint64_t actingBlockLength,
        const std::uint64_t actingVersion) :
        MassQuote(buffer, 0, bufferLength, actingBlockLength, actingVersion)
    {
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint16_t sbeBlockLength() SBE_NOEXCEPT
    {
        return (std::uint16_t)62;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint16_t sbeBlockAndHeaderLength() SBE_NOEXCEPT
    {
        return MessageHeader::encodedLength() + sbeBlockLength();
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint16_t sbeTemplateId() SBE_NOEXCEPT
    {
        return (std::uint16_t)105;
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
        return "i";
    }

    SBE_NODISCARD std::uint64_t offset() const SBE_NOEXCEPT
    {
        return m_offset;
    }

    MassQuote &wrapForEncode(char *buffer, const std::uint64_t offset, const std::uint64_t bufferLength)
    {
        return *this = MassQuote(buffer, offset, bufferLength, sbeBlockLength(), sbeSchemaVersion());
    }

    MassQuote &wrapAndApplyHeader(char *buffer, const std::uint64_t offset, const std::uint64_t bufferLength)
    {
        MessageHeader hdr(buffer, offset, bufferLength, sbeSchemaVersion());

        hdr
            .blockLength(sbeBlockLength())
            .templateId(sbeTemplateId())
            .schemaId(sbeSchemaId())
            .version(sbeSchemaVersion());

        return *this = MassQuote(
            buffer,
            offset + MessageHeader::encodedLength(),
            bufferLength,
            sbeBlockLength(),
            sbeSchemaVersion());
    }

    MassQuote &wrapForDecode(
        char *buffer,
        const std::uint64_t offset,
        const std::uint64_t actingBlockLength,
        const std::uint64_t actingVersion,
        const std::uint64_t bufferLength)
    {
        return *this = MassQuote(buffer, offset, bufferLength, actingBlockLength, actingVersion);
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
        MassQuote skipper(m_buffer, m_offset,
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

    SBE_NODISCARD static const char * QuoteReqIDMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::SEMANTIC_TYPE: return "String";
            case MetaAttribute::PRESENCE: return "required";
            default: return "";
        }
    }

    static SBE_CONSTEXPR std::uint16_t quoteReqIDId() SBE_NOEXCEPT
    {
        return 131;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t quoteReqIDSinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool quoteReqIDInActingVersion() SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= quoteReqIDSinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t quoteReqIDEncodingOffset() SBE_NOEXCEPT
    {
        return 0;
    }

    static SBE_CONSTEXPR char quoteReqIDNullValue() SBE_NOEXCEPT
    {
        return (char)0;
    }

    static SBE_CONSTEXPR char quoteReqIDMinValue() SBE_NOEXCEPT
    {
        return (char)32;
    }

    static SBE_CONSTEXPR char quoteReqIDMaxValue() SBE_NOEXCEPT
    {
        return (char)126;
    }

    static SBE_CONSTEXPR std::size_t quoteReqIDEncodingLength() SBE_NOEXCEPT
    {
        return 23;
    }

    static SBE_CONSTEXPR std::uint64_t quoteReqIDLength() SBE_NOEXCEPT
    {
        return 23;
    }

    SBE_NODISCARD const char *quoteReqID() const SBE_NOEXCEPT
    {
        return m_buffer + m_offset + 0;
    }

    SBE_NODISCARD char *quoteReqID() SBE_NOEXCEPT
    {
        return m_buffer + m_offset + 0;
    }

    SBE_NODISCARD char quoteReqID(const std::uint64_t index) const
    {
        if (index >= 23)
        {
            throw std::runtime_error("index out of range for quoteReqID [E104]");
        }

        char val;
        std::memcpy(&val, m_buffer + m_offset + 0 + (index * 1), sizeof(char));
        return (val);
    }

    MassQuote &quoteReqID(const std::uint64_t index, const char value)
    {
        if (index >= 23)
        {
            throw std::runtime_error("index out of range for quoteReqID [E105]");
        }

        char val = (value);
        std::memcpy(m_buffer + m_offset + 0 + (index * 1), &val, sizeof(char));
        return *this;
    }

    std::uint64_t getQuoteReqID(char *const dst, const std::uint64_t length) const
    {
        if (length > 23)
        {
            throw std::runtime_error("length too large for getQuoteReqID [E106]");
        }

        std::memcpy(dst, m_buffer + m_offset + 0, sizeof(char) * static_cast<size_t>(length));
        return length;
    }

    MassQuote &putQuoteReqID(const char *const src) SBE_NOEXCEPT
    {
        std::memcpy(m_buffer + m_offset + 0, src, sizeof(char) * 23);
        return *this;
    }

    SBE_NODISCARD std::string getQuoteReqIDAsString() const
    {
        const char *buffer = m_buffer + m_offset + 0;
        size_t length = 0;

        for (; length < 23 && *(buffer + length) != '\0'; ++length);
        std::string result(buffer, length);

        return result;
    }

    std::string getQuoteReqIDAsJsonEscapedString()
    {
        std::ostringstream oss;
        std::string s = getQuoteReqIDAsString();

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
    SBE_NODISCARD std::string_view getQuoteReqIDAsStringView() const SBE_NOEXCEPT
    {
        const char *buffer = m_buffer + m_offset + 0;
        size_t length = 0;

        for (; length < 23 && *(buffer + length) != '\0'; ++length);
        std::string_view result(buffer, length);

        return result;
    }
    #endif

    #if __cplusplus >= 201703L
    MassQuote &putQuoteReqID(const std::string_view str)
    {
        const size_t srcLength = str.length();
        if (srcLength > 23)
        {
            throw std::runtime_error("string too large for putQuoteReqID [E106]");
        }

        std::memcpy(m_buffer + m_offset + 0, str.data(), srcLength);
        for (size_t start = srcLength; start < 23; ++start)
        {
            m_buffer[m_offset + 0 + start] = 0;
        }

        return *this;
    }
    #else
    MassQuote &putQuoteReqID(const std::string& str)
    {
        const size_t srcLength = str.length();
        if (srcLength > 23)
        {
            throw std::runtime_error("string too large for putQuoteReqID [E106]");
        }

        std::memcpy(m_buffer + m_offset + 0, str.c_str(), srcLength);
        for (size_t start = srcLength; start < 23; ++start)
        {
            m_buffer[m_offset + 0 + start] = 0;
        }

        return *this;
    }
    #endif

    SBE_NODISCARD static const char * QuoteIDMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::SEMANTIC_TYPE: return "String";
            case MetaAttribute::PRESENCE: return "required";
            default: return "";
        }
    }

    static SBE_CONSTEXPR std::uint16_t quoteIDId() SBE_NOEXCEPT
    {
        return 117;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t quoteIDSinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool quoteIDInActingVersion() SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= quoteIDSinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t quoteIDEncodingOffset() SBE_NOEXCEPT
    {
        return 23;
    }

    static SBE_CONSTEXPR char quoteIDNullValue() SBE_NOEXCEPT
    {
        return (char)0;
    }

    static SBE_CONSTEXPR char quoteIDMinValue() SBE_NOEXCEPT
    {
        return (char)32;
    }

    static SBE_CONSTEXPR char quoteIDMaxValue() SBE_NOEXCEPT
    {
        return (char)126;
    }

    static SBE_CONSTEXPR std::size_t quoteIDEncodingLength() SBE_NOEXCEPT
    {
        return 10;
    }

    static SBE_CONSTEXPR std::uint64_t quoteIDLength() SBE_NOEXCEPT
    {
        return 10;
    }

    SBE_NODISCARD const char *quoteID() const SBE_NOEXCEPT
    {
        return m_buffer + m_offset + 23;
    }

    SBE_NODISCARD char *quoteID() SBE_NOEXCEPT
    {
        return m_buffer + m_offset + 23;
    }

    SBE_NODISCARD char quoteID(const std::uint64_t index) const
    {
        if (index >= 10)
        {
            throw std::runtime_error("index out of range for quoteID [E104]");
        }

        char val;
        std::memcpy(&val, m_buffer + m_offset + 23 + (index * 1), sizeof(char));
        return (val);
    }

    MassQuote &quoteID(const std::uint64_t index, const char value)
    {
        if (index >= 10)
        {
            throw std::runtime_error("index out of range for quoteID [E105]");
        }

        char val = (value);
        std::memcpy(m_buffer + m_offset + 23 + (index * 1), &val, sizeof(char));
        return *this;
    }

    std::uint64_t getQuoteID(char *const dst, const std::uint64_t length) const
    {
        if (length > 10)
        {
            throw std::runtime_error("length too large for getQuoteID [E106]");
        }

        std::memcpy(dst, m_buffer + m_offset + 23, sizeof(char) * static_cast<size_t>(length));
        return length;
    }

    MassQuote &putQuoteID(const char *const src) SBE_NOEXCEPT
    {
        std::memcpy(m_buffer + m_offset + 23, src, sizeof(char) * 10);
        return *this;
    }

    SBE_NODISCARD std::string getQuoteIDAsString() const
    {
        const char *buffer = m_buffer + m_offset + 23;
        size_t length = 0;

        for (; length < 10 && *(buffer + length) != '\0'; ++length);
        std::string result(buffer, length);

        return result;
    }

    std::string getQuoteIDAsJsonEscapedString()
    {
        std::ostringstream oss;
        std::string s = getQuoteIDAsString();

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
    SBE_NODISCARD std::string_view getQuoteIDAsStringView() const SBE_NOEXCEPT
    {
        const char *buffer = m_buffer + m_offset + 23;
        size_t length = 0;

        for (; length < 10 && *(buffer + length) != '\0'; ++length);
        std::string_view result(buffer, length);

        return result;
    }
    #endif

    #if __cplusplus >= 201703L
    MassQuote &putQuoteID(const std::string_view str)
    {
        const size_t srcLength = str.length();
        if (srcLength > 10)
        {
            throw std::runtime_error("string too large for putQuoteID [E106]");
        }

        std::memcpy(m_buffer + m_offset + 23, str.data(), srcLength);
        for (size_t start = srcLength; start < 10; ++start)
        {
            m_buffer[m_offset + 23 + start] = 0;
        }

        return *this;
    }
    #else
    MassQuote &putQuoteID(const std::string& str)
    {
        const size_t srcLength = str.length();
        if (srcLength > 10)
        {
            throw std::runtime_error("string too large for putQuoteID [E106]");
        }

        std::memcpy(m_buffer + m_offset + 23, str.c_str(), srcLength);
        for (size_t start = srcLength; start < 10; ++start)
        {
            m_buffer[m_offset + 23 + start] = 0;
        }

        return *this;
    }
    #endif

    SBE_NODISCARD static const char * MMAccountMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::SEMANTIC_TYPE: return "String";
            case MetaAttribute::PRESENCE: return "required";
            default: return "";
        }
    }

    static SBE_CONSTEXPR std::uint16_t mMAccountId() SBE_NOEXCEPT
    {
        return 9771;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t mMAccountSinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool mMAccountInActingVersion() SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= mMAccountSinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t mMAccountEncodingOffset() SBE_NOEXCEPT
    {
        return 33;
    }

    static SBE_CONSTEXPR char mMAccountNullValue() SBE_NOEXCEPT
    {
        return (char)0;
    }

    static SBE_CONSTEXPR char mMAccountMinValue() SBE_NOEXCEPT
    {
        return (char)32;
    }

    static SBE_CONSTEXPR char mMAccountMaxValue() SBE_NOEXCEPT
    {
        return (char)126;
    }

    static SBE_CONSTEXPR std::size_t mMAccountEncodingLength() SBE_NOEXCEPT
    {
        return 12;
    }

    static SBE_CONSTEXPR std::uint64_t mMAccountLength() SBE_NOEXCEPT
    {
        return 12;
    }

    SBE_NODISCARD const char *mMAccount() const SBE_NOEXCEPT
    {
        return m_buffer + m_offset + 33;
    }

    SBE_NODISCARD char *mMAccount() SBE_NOEXCEPT
    {
        return m_buffer + m_offset + 33;
    }

    SBE_NODISCARD char mMAccount(const std::uint64_t index) const
    {
        if (index >= 12)
        {
            throw std::runtime_error("index out of range for mMAccount [E104]");
        }

        char val;
        std::memcpy(&val, m_buffer + m_offset + 33 + (index * 1), sizeof(char));
        return (val);
    }

    MassQuote &mMAccount(const std::uint64_t index, const char value)
    {
        if (index >= 12)
        {
            throw std::runtime_error("index out of range for mMAccount [E105]");
        }

        char val = (value);
        std::memcpy(m_buffer + m_offset + 33 + (index * 1), &val, sizeof(char));
        return *this;
    }

    std::uint64_t getMMAccount(char *const dst, const std::uint64_t length) const
    {
        if (length > 12)
        {
            throw std::runtime_error("length too large for getMMAccount [E106]");
        }

        std::memcpy(dst, m_buffer + m_offset + 33, sizeof(char) * static_cast<size_t>(length));
        return length;
    }

    MassQuote &putMMAccount(const char *const src) SBE_NOEXCEPT
    {
        std::memcpy(m_buffer + m_offset + 33, src, sizeof(char) * 12);
        return *this;
    }

    SBE_NODISCARD std::string getMMAccountAsString() const
    {
        const char *buffer = m_buffer + m_offset + 33;
        size_t length = 0;

        for (; length < 12 && *(buffer + length) != '\0'; ++length);
        std::string result(buffer, length);

        return result;
    }

    std::string getMMAccountAsJsonEscapedString()
    {
        std::ostringstream oss;
        std::string s = getMMAccountAsString();

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
    SBE_NODISCARD std::string_view getMMAccountAsStringView() const SBE_NOEXCEPT
    {
        const char *buffer = m_buffer + m_offset + 33;
        size_t length = 0;

        for (; length < 12 && *(buffer + length) != '\0'; ++length);
        std::string_view result(buffer, length);

        return result;
    }
    #endif

    #if __cplusplus >= 201703L
    MassQuote &putMMAccount(const std::string_view str)
    {
        const size_t srcLength = str.length();
        if (srcLength > 12)
        {
            throw std::runtime_error("string too large for putMMAccount [E106]");
        }

        std::memcpy(m_buffer + m_offset + 33, str.data(), srcLength);
        for (size_t start = srcLength; start < 12; ++start)
        {
            m_buffer[m_offset + 33 + start] = 0;
        }

        return *this;
    }
    #else
    MassQuote &putMMAccount(const std::string& str)
    {
        const size_t srcLength = str.length();
        if (srcLength > 12)
        {
            throw std::runtime_error("string too large for putMMAccount [E106]");
        }

        std::memcpy(m_buffer + m_offset + 33, str.c_str(), srcLength);
        for (size_t start = srcLength; start < 12; ++start)
        {
            m_buffer[m_offset + 33 + start] = 0;
        }

        return *this;
    }
    #endif

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
        return 45;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t manualOrderIndicatorEncodingLength() SBE_NOEXCEPT
    {
        return 1;
    }

    SBE_NODISCARD BooleanType::Value manualOrderIndicator() const SBE_NOEXCEPT
    {
        std::uint8_t val;
        std::memcpy(&val, m_buffer + m_offset + 45, sizeof(std::uint8_t));
        return BooleanType::get((val));
    }

    MassQuote &manualOrderIndicator(const BooleanType::Value value) SBE_NOEXCEPT
    {
        std::uint8_t val = (value);
        std::memcpy(m_buffer + m_offset + 45, &val, sizeof(std::uint8_t));
        return *this;
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
        return 46;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t custOrderHandlingInstEncodingLength() SBE_NOEXCEPT
    {
        return 1;
    }

    SBE_NODISCARD CustOrderHandlingInst::Value custOrderHandlingInst() const SBE_NOEXCEPT
    {
        char val;
        std::memcpy(&val, m_buffer + m_offset + 46, sizeof(char));
        return CustOrderHandlingInst::get((val));
    }

    MassQuote &custOrderHandlingInst(const CustOrderHandlingInst::Value value) SBE_NOEXCEPT
    {
        char val = (value);
        std::memcpy(m_buffer + m_offset + 46, &val, sizeof(char));
        return *this;
    }

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
        return 47;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t customerOrFirmEncodingLength() SBE_NOEXCEPT
    {
        return 1;
    }

    SBE_NODISCARD CustomerOrFirm::Value customerOrFirm() const SBE_NOEXCEPT
    {
        std::uint8_t val;
        std::memcpy(&val, m_buffer + m_offset + 47, sizeof(std::uint8_t));
        return CustomerOrFirm::get((val));
    }

    MassQuote &customerOrFirm(const CustomerOrFirm::Value value) SBE_NOEXCEPT
    {
        std::uint8_t val = (value);
        std::memcpy(m_buffer + m_offset + 47, &val, sizeof(std::uint8_t));
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
        return 48;
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
        return m_buffer + m_offset + 48;
    }

    SBE_NODISCARD char *selfMatchPreventionID() SBE_NOEXCEPT
    {
        return m_buffer + m_offset + 48;
    }

    SBE_NODISCARD char selfMatchPreventionID(const std::uint64_t index) const
    {
        if (index >= 12)
        {
            throw std::runtime_error("index out of range for selfMatchPreventionID [E104]");
        }

        char val;
        std::memcpy(&val, m_buffer + m_offset + 48 + (index * 1), sizeof(char));
        return (val);
    }

    MassQuote &selfMatchPreventionID(const std::uint64_t index, const char value)
    {
        if (index >= 12)
        {
            throw std::runtime_error("index out of range for selfMatchPreventionID [E105]");
        }

        char val = (value);
        std::memcpy(m_buffer + m_offset + 48 + (index * 1), &val, sizeof(char));
        return *this;
    }

    std::uint64_t getSelfMatchPreventionID(char *const dst, const std::uint64_t length) const
    {
        if (length > 12)
        {
            throw std::runtime_error("length too large for getSelfMatchPreventionID [E106]");
        }

        std::memcpy(dst, m_buffer + m_offset + 48, sizeof(char) * static_cast<size_t>(length));
        return length;
    }

    MassQuote &putSelfMatchPreventionID(const char *const src) SBE_NOEXCEPT
    {
        std::memcpy(m_buffer + m_offset + 48, src, sizeof(char) * 12);
        return *this;
    }

    SBE_NODISCARD std::string getSelfMatchPreventionIDAsString() const
    {
        const char *buffer = m_buffer + m_offset + 48;
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
        const char *buffer = m_buffer + m_offset + 48;
        size_t length = 0;

        for (; length < 12 && *(buffer + length) != '\0'; ++length);
        std::string_view result(buffer, length);

        return result;
    }
    #endif

    #if __cplusplus >= 201703L
    MassQuote &putSelfMatchPreventionID(const std::string_view str)
    {
        const size_t srcLength = str.length();
        if (srcLength > 12)
        {
            throw std::runtime_error("string too large for putSelfMatchPreventionID [E106]");
        }

        std::memcpy(m_buffer + m_offset + 48, str.data(), srcLength);
        for (size_t start = srcLength; start < 12; ++start)
        {
            m_buffer[m_offset + 48 + start] = 0;
        }

        return *this;
    }
    #else
    MassQuote &putSelfMatchPreventionID(const std::string& str)
    {
        const size_t srcLength = str.length();
        if (srcLength > 12)
        {
            throw std::runtime_error("string too large for putSelfMatchPreventionID [E106]");
        }

        std::memcpy(m_buffer + m_offset + 48, str.c_str(), srcLength);
        for (size_t start = srcLength; start < 12; ++start)
        {
            m_buffer[m_offset + 48 + start] = 0;
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
        return 60;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t ctiCodeEncodingLength() SBE_NOEXCEPT
    {
        return 1;
    }

    SBE_NODISCARD CtiCode::Value ctiCode() const SBE_NOEXCEPT
    {
        char val;
        std::memcpy(&val, m_buffer + m_offset + 60, sizeof(char));
        return CtiCode::get((val));
    }

    MassQuote &ctiCode(const CtiCode::Value value) SBE_NOEXCEPT
    {
        char val = (value);
        std::memcpy(m_buffer + m_offset + 60, &val, sizeof(char));
        return *this;
    }

    SBE_NODISCARD static const char * MMProtectionResetMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
    {
        switch (metaAttribute)
        {
            case MetaAttribute::PRESENCE: return "optional";
            default: return "";
        }
    }

    static SBE_CONSTEXPR std::uint16_t mMProtectionResetId() SBE_NOEXCEPT
    {
        return 9773;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t mMProtectionResetSinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool mMProtectionResetInActingVersion() SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= mMProtectionResetSinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t mMProtectionResetEncodingOffset() SBE_NOEXCEPT
    {
        return 61;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::size_t mMProtectionResetEncodingLength() SBE_NOEXCEPT
    {
        return 1;
    }

    SBE_NODISCARD MMProtectionReset::Value mMProtectionReset() const SBE_NOEXCEPT
    {
        char val;
        std::memcpy(&val, m_buffer + m_offset + 61, sizeof(char));
        return MMProtectionReset::get((val));
    }

    MassQuote &mMProtectionReset(const MMProtectionReset::Value value) SBE_NOEXCEPT
    {
        char val = (value);
        std::memcpy(m_buffer + m_offset + 61, &val, sizeof(char));
        return *this;
    }

    class QuoteSets
    {
    private:
        char *m_buffer = nullptr;
        std::uint64_t m_bufferLength = 0;
        std::uint64_t m_initialPosition = 0;
        std::uint64_t *m_positionPtr = nullptr;
        std::uint64_t m_blockLength = 0;
        std::uint64_t m_count = 0;
        std::uint64_t m_index = 0;
        std::uint64_t m_offset = 0;
        std::uint64_t m_actingVersion = 0;

        SBE_NODISCARD std::uint64_t *sbePositionPtr() SBE_NOEXCEPT
        {
            return m_positionPtr;
        }

    public:
        inline void wrapForDecode(
            char *buffer,
            std::uint64_t *pos,
            const std::uint64_t actingVersion,
            const std::uint64_t bufferLength)
        {
            GroupSizeEncoding dimensions(buffer, *pos, bufferLength, actingVersion);
            m_buffer = buffer;
            m_bufferLength = bufferLength;
            m_blockLength = dimensions.blockLength();
            m_count = dimensions.numInGroup();
            m_index = 0;
            m_actingVersion = actingVersion;
            m_initialPosition = *pos;
            m_positionPtr = pos;
            *m_positionPtr = *m_positionPtr + 4;
        }

        inline void wrapForEncode(
            char *buffer,
            const std::uint16_t count,
            std::uint64_t *pos,
            const std::uint64_t actingVersion,
            const std::uint64_t bufferLength)
        {
    #if defined(__GNUG__) && !defined(__clang__)
    #pragma GCC diagnostic push
    #pragma GCC diagnostic ignored "-Wtype-limits"
    #endif
            if (count > 65534)
            {
                throw std::runtime_error("count outside of allowed range [E110]");
            }
    #if defined(__GNUG__) && !defined(__clang__)
    #pragma GCC diagnostic pop
    #endif
            m_buffer = buffer;
            m_bufferLength = bufferLength;
            GroupSizeEncoding dimensions(buffer, *pos, bufferLength, actingVersion);
            dimensions.blockLength((std::uint16_t)24);
            dimensions.numInGroup((std::uint16_t)count);
            m_index = 0;
            m_count = count;
            m_blockLength = 24;
            m_actingVersion = actingVersion;
            m_initialPosition = *pos;
            m_positionPtr = pos;
            *m_positionPtr = *m_positionPtr + 4;
        }

        static SBE_CONSTEXPR std::uint64_t sbeHeaderSize() SBE_NOEXCEPT
        {
            return 4;
        }

        static SBE_CONSTEXPR std::uint64_t sbeBlockLength() SBE_NOEXCEPT
        {
            return 24;
        }

        SBE_NODISCARD std::uint64_t sbePosition() const SBE_NOEXCEPT
        {
            return *m_positionPtr;
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
            *m_positionPtr = sbeCheckPosition(position);
        }

        SBE_NODISCARD inline std::uint64_t count() const SBE_NOEXCEPT
        {
            return m_count;
        }

        SBE_NODISCARD inline bool hasNext() const SBE_NOEXCEPT
        {
            return m_index < m_count;
        }

        inline QuoteSets &next()
        {
            if (m_index >= m_count)
            {
                throw std::runtime_error("index >= count [E108]");
            }
            m_offset = *m_positionPtr;
            if (SBE_BOUNDS_CHECK_EXPECT(((m_offset + m_blockLength) > m_bufferLength), false))
            {
                throw std::runtime_error("buffer too short for next group index [E108]");
            }
            *m_positionPtr = m_offset + m_blockLength;
            ++m_index;

            return *this;
        }

        inline std::uint64_t resetCountToIndex() SBE_NOEXCEPT
        {
            m_count = m_index;
            GroupSizeEncoding dimensions(m_buffer, m_initialPosition, m_bufferLength, m_actingVersion);
            dimensions.numInGroup((std::uint16_t)m_count);
            return m_count;
        }

    #if __cplusplus < 201103L
        template<class Func> inline void forEach(Func& func)
        {
            while (hasNext())
            {
                next();
                func(*this);
            }
        }

    #else
        template<class Func> inline void forEach(Func&& func)
        {
            while (hasNext())
            {
                next();
                func(*this);
            }
        }

    #endif

        SBE_NODISCARD static const char * QuoteSetIDMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
        {
            switch (metaAttribute)
            {
                case MetaAttribute::SEMANTIC_TYPE: return "String";
                case MetaAttribute::PRESENCE: return "required";
                default: return "";
            }
        }

        static SBE_CONSTEXPR std::uint16_t quoteSetIDId() SBE_NOEXCEPT
        {
            return 302;
        }

        SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t quoteSetIDSinceVersion() SBE_NOEXCEPT
        {
            return 0;
        }

        SBE_NODISCARD bool quoteSetIDInActingVersion() SBE_NOEXCEPT
        {
    #if defined(__clang__)
    #pragma clang diagnostic push
    #pragma clang diagnostic ignored "-Wtautological-compare"
    #endif
            return m_actingVersion >= quoteSetIDSinceVersion();
    #if defined(__clang__)
    #pragma clang diagnostic pop
    #endif
        }

        SBE_NODISCARD static SBE_CONSTEXPR std::size_t quoteSetIDEncodingOffset() SBE_NOEXCEPT
        {
            return 0;
        }

        static SBE_CONSTEXPR char quoteSetIDNullValue() SBE_NOEXCEPT
        {
            return (char)0;
        }

        static SBE_CONSTEXPR char quoteSetIDMinValue() SBE_NOEXCEPT
        {
            return (char)32;
        }

        static SBE_CONSTEXPR char quoteSetIDMaxValue() SBE_NOEXCEPT
        {
            return (char)126;
        }

        static SBE_CONSTEXPR std::size_t quoteSetIDEncodingLength() SBE_NOEXCEPT
        {
            return 3;
        }

        static SBE_CONSTEXPR std::uint64_t quoteSetIDLength() SBE_NOEXCEPT
        {
            return 3;
        }

        SBE_NODISCARD const char *quoteSetID() const SBE_NOEXCEPT
        {
            return m_buffer + m_offset + 0;
        }

        SBE_NODISCARD char *quoteSetID() SBE_NOEXCEPT
        {
            return m_buffer + m_offset + 0;
        }

        SBE_NODISCARD char quoteSetID(const std::uint64_t index) const
        {
            if (index >= 3)
            {
                throw std::runtime_error("index out of range for quoteSetID [E104]");
            }

            char val;
            std::memcpy(&val, m_buffer + m_offset + 0 + (index * 1), sizeof(char));
            return (val);
        }

        QuoteSets &quoteSetID(const std::uint64_t index, const char value)
        {
            if (index >= 3)
            {
                throw std::runtime_error("index out of range for quoteSetID [E105]");
            }

            char val = (value);
            std::memcpy(m_buffer + m_offset + 0 + (index * 1), &val, sizeof(char));
            return *this;
        }

        std::uint64_t getQuoteSetID(char *const dst, const std::uint64_t length) const
        {
            if (length > 3)
            {
                throw std::runtime_error("length too large for getQuoteSetID [E106]");
            }

            std::memcpy(dst, m_buffer + m_offset + 0, sizeof(char) * static_cast<size_t>(length));
            return length;
        }

        QuoteSets &putQuoteSetID(const char *const src) SBE_NOEXCEPT
        {
            std::memcpy(m_buffer + m_offset + 0, src, sizeof(char) * 3);
            return *this;
        }

        QuoteSets &putQuoteSetID(
            const char value0,
            const char value1,
            const char value2) SBE_NOEXCEPT
        {
            char val0 = (value0);
            std::memcpy(m_buffer + m_offset + 0, &val0, sizeof(char));
            char val1 = (value1);
            std::memcpy(m_buffer + m_offset + 1, &val1, sizeof(char));
            char val2 = (value2);
            std::memcpy(m_buffer + m_offset + 2, &val2, sizeof(char));

            return *this;
        }

        SBE_NODISCARD std::string getQuoteSetIDAsString() const
        {
            const char *buffer = m_buffer + m_offset + 0;
            size_t length = 0;

            for (; length < 3 && *(buffer + length) != '\0'; ++length);
            std::string result(buffer, length);

            return result;
        }

        std::string getQuoteSetIDAsJsonEscapedString()
        {
            std::ostringstream oss;
            std::string s = getQuoteSetIDAsString();

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
        SBE_NODISCARD std::string_view getQuoteSetIDAsStringView() const SBE_NOEXCEPT
        {
            const char *buffer = m_buffer + m_offset + 0;
            size_t length = 0;

            for (; length < 3 && *(buffer + length) != '\0'; ++length);
            std::string_view result(buffer, length);

            return result;
        }
        #endif

        #if __cplusplus >= 201703L
        QuoteSets &putQuoteSetID(const std::string_view str)
        {
            const size_t srcLength = str.length();
            if (srcLength > 3)
            {
                throw std::runtime_error("string too large for putQuoteSetID [E106]");
            }

            std::memcpy(m_buffer + m_offset + 0, str.data(), srcLength);
            for (size_t start = srcLength; start < 3; ++start)
            {
                m_buffer[m_offset + 0 + start] = 0;
            }

            return *this;
        }
        #else
        QuoteSets &putQuoteSetID(const std::string& str)
        {
            const size_t srcLength = str.length();
            if (srcLength > 3)
            {
                throw std::runtime_error("string too large for putQuoteSetID [E106]");
            }

            std::memcpy(m_buffer + m_offset + 0, str.c_str(), srcLength);
            for (size_t start = srcLength; start < 3; ++start)
            {
                m_buffer[m_offset + 0 + start] = 0;
            }

            return *this;
        }
        #endif

        SBE_NODISCARD static const char * UnderlyingSecurityDescMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
        {
            switch (metaAttribute)
            {
                case MetaAttribute::SEMANTIC_TYPE: return "String";
                case MetaAttribute::PRESENCE: return "required";
                default: return "";
            }
        }

        static SBE_CONSTEXPR std::uint16_t underlyingSecurityDescId() SBE_NOEXCEPT
        {
            return 307;
        }

        SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t underlyingSecurityDescSinceVersion() SBE_NOEXCEPT
        {
            return 0;
        }

        SBE_NODISCARD bool underlyingSecurityDescInActingVersion() SBE_NOEXCEPT
        {
    #if defined(__clang__)
    #pragma clang diagnostic push
    #pragma clang diagnostic ignored "-Wtautological-compare"
    #endif
            return m_actingVersion >= underlyingSecurityDescSinceVersion();
    #if defined(__clang__)
    #pragma clang diagnostic pop
    #endif
        }

        SBE_NODISCARD static SBE_CONSTEXPR std::size_t underlyingSecurityDescEncodingOffset() SBE_NOEXCEPT
        {
            return 3;
        }

        static SBE_CONSTEXPR char underlyingSecurityDescNullValue() SBE_NOEXCEPT
        {
            return (char)0;
        }

        static SBE_CONSTEXPR char underlyingSecurityDescMinValue() SBE_NOEXCEPT
        {
            return (char)32;
        }

        static SBE_CONSTEXPR char underlyingSecurityDescMaxValue() SBE_NOEXCEPT
        {
            return (char)126;
        }

        static SBE_CONSTEXPR std::size_t underlyingSecurityDescEncodingLength() SBE_NOEXCEPT
        {
            return 20;
        }

        static SBE_CONSTEXPR std::uint64_t underlyingSecurityDescLength() SBE_NOEXCEPT
        {
            return 20;
        }

        SBE_NODISCARD const char *underlyingSecurityDesc() const SBE_NOEXCEPT
        {
            return m_buffer + m_offset + 3;
        }

        SBE_NODISCARD char *underlyingSecurityDesc() SBE_NOEXCEPT
        {
            return m_buffer + m_offset + 3;
        }

        SBE_NODISCARD char underlyingSecurityDesc(const std::uint64_t index) const
        {
            if (index >= 20)
            {
                throw std::runtime_error("index out of range for underlyingSecurityDesc [E104]");
            }

            char val;
            std::memcpy(&val, m_buffer + m_offset + 3 + (index * 1), sizeof(char));
            return (val);
        }

        QuoteSets &underlyingSecurityDesc(const std::uint64_t index, const char value)
        {
            if (index >= 20)
            {
                throw std::runtime_error("index out of range for underlyingSecurityDesc [E105]");
            }

            char val = (value);
            std::memcpy(m_buffer + m_offset + 3 + (index * 1), &val, sizeof(char));
            return *this;
        }

        std::uint64_t getUnderlyingSecurityDesc(char *const dst, const std::uint64_t length) const
        {
            if (length > 20)
            {
                throw std::runtime_error("length too large for getUnderlyingSecurityDesc [E106]");
            }

            std::memcpy(dst, m_buffer + m_offset + 3, sizeof(char) * static_cast<size_t>(length));
            return length;
        }

        QuoteSets &putUnderlyingSecurityDesc(const char *const src) SBE_NOEXCEPT
        {
            std::memcpy(m_buffer + m_offset + 3, src, sizeof(char) * 20);
            return *this;
        }

        SBE_NODISCARD std::string getUnderlyingSecurityDescAsString() const
        {
            const char *buffer = m_buffer + m_offset + 3;
            size_t length = 0;

            for (; length < 20 && *(buffer + length) != '\0'; ++length);
            std::string result(buffer, length);

            return result;
        }

        std::string getUnderlyingSecurityDescAsJsonEscapedString()
        {
            std::ostringstream oss;
            std::string s = getUnderlyingSecurityDescAsString();

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
        SBE_NODISCARD std::string_view getUnderlyingSecurityDescAsStringView() const SBE_NOEXCEPT
        {
            const char *buffer = m_buffer + m_offset + 3;
            size_t length = 0;

            for (; length < 20 && *(buffer + length) != '\0'; ++length);
            std::string_view result(buffer, length);

            return result;
        }
        #endif

        #if __cplusplus >= 201703L
        QuoteSets &putUnderlyingSecurityDesc(const std::string_view str)
        {
            const size_t srcLength = str.length();
            if (srcLength > 20)
            {
                throw std::runtime_error("string too large for putUnderlyingSecurityDesc [E106]");
            }

            std::memcpy(m_buffer + m_offset + 3, str.data(), srcLength);
            for (size_t start = srcLength; start < 20; ++start)
            {
                m_buffer[m_offset + 3 + start] = 0;
            }

            return *this;
        }
        #else
        QuoteSets &putUnderlyingSecurityDesc(const std::string& str)
        {
            const size_t srcLength = str.length();
            if (srcLength > 20)
            {
                throw std::runtime_error("string too large for putUnderlyingSecurityDesc [E106]");
            }

            std::memcpy(m_buffer + m_offset + 3, str.c_str(), srcLength);
            for (size_t start = srcLength; start < 20; ++start)
            {
                m_buffer[m_offset + 3 + start] = 0;
            }

            return *this;
        }
        #endif

        SBE_NODISCARD static const char * TotQuoteEntriesMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
        {
            switch (metaAttribute)
            {
                case MetaAttribute::SEMANTIC_TYPE: return "int";
                case MetaAttribute::PRESENCE: return "required";
                default: return "";
            }
        }

        static SBE_CONSTEXPR std::uint16_t totQuoteEntriesId() SBE_NOEXCEPT
        {
            return 304;
        }

        SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t totQuoteEntriesSinceVersion() SBE_NOEXCEPT
        {
            return 0;
        }

        SBE_NODISCARD bool totQuoteEntriesInActingVersion() SBE_NOEXCEPT
        {
    #if defined(__clang__)
    #pragma clang diagnostic push
    #pragma clang diagnostic ignored "-Wtautological-compare"
    #endif
            return m_actingVersion >= totQuoteEntriesSinceVersion();
    #if defined(__clang__)
    #pragma clang diagnostic pop
    #endif
        }

        SBE_NODISCARD static SBE_CONSTEXPR std::size_t totQuoteEntriesEncodingOffset() SBE_NOEXCEPT
        {
            return 23;
        }

        static SBE_CONSTEXPR std::uint8_t totQuoteEntriesNullValue() SBE_NOEXCEPT
        {
            return SBE_NULLVALUE_UINT8;
        }

        static SBE_CONSTEXPR std::uint8_t totQuoteEntriesMinValue() SBE_NOEXCEPT
        {
            return (std::uint8_t)0;
        }

        static SBE_CONSTEXPR std::uint8_t totQuoteEntriesMaxValue() SBE_NOEXCEPT
        {
            return (std::uint8_t)254;
        }

        static SBE_CONSTEXPR std::size_t totQuoteEntriesEncodingLength() SBE_NOEXCEPT
        {
            return 1;
        }

        SBE_NODISCARD std::uint8_t totQuoteEntries() const SBE_NOEXCEPT
        {
            std::uint8_t val;
            std::memcpy(&val, m_buffer + m_offset + 23, sizeof(std::uint8_t));
            return (val);
        }

        QuoteSets &totQuoteEntries(const std::uint8_t value) SBE_NOEXCEPT
        {
            std::uint8_t val = (value);
            std::memcpy(m_buffer + m_offset + 23, &val, sizeof(std::uint8_t));
            return *this;
        }

        class QuoteEntries
        {
        private:
            char *m_buffer = nullptr;
            std::uint64_t m_bufferLength = 0;
            std::uint64_t m_initialPosition = 0;
            std::uint64_t *m_positionPtr = nullptr;
            std::uint64_t m_blockLength = 0;
            std::uint64_t m_count = 0;
            std::uint64_t m_index = 0;
            std::uint64_t m_offset = 0;
            std::uint64_t m_actingVersion = 0;

            SBE_NODISCARD std::uint64_t *sbePositionPtr() SBE_NOEXCEPT
            {
                return m_positionPtr;
            }

        public:
            inline void wrapForDecode(
                char *buffer,
                std::uint64_t *pos,
                const std::uint64_t actingVersion,
                const std::uint64_t bufferLength)
            {
                GroupSizeEncoding dimensions(buffer, *pos, bufferLength, actingVersion);
                m_buffer = buffer;
                m_bufferLength = bufferLength;
                m_blockLength = dimensions.blockLength();
                m_count = dimensions.numInGroup();
                m_index = 0;
                m_actingVersion = actingVersion;
                m_initialPosition = *pos;
                m_positionPtr = pos;
                *m_positionPtr = *m_positionPtr + 4;
            }

            inline void wrapForEncode(
                char *buffer,
                const std::uint16_t count,
                std::uint64_t *pos,
                const std::uint64_t actingVersion,
                const std::uint64_t bufferLength)
            {
        #if defined(__GNUG__) && !defined(__clang__)
        #pragma GCC diagnostic push
        #pragma GCC diagnostic ignored "-Wtype-limits"
        #endif
                if (count > 65534)
                {
                    throw std::runtime_error("count outside of allowed range [E110]");
                }
        #if defined(__GNUG__) && !defined(__clang__)
        #pragma GCC diagnostic pop
        #endif
                m_buffer = buffer;
                m_bufferLength = bufferLength;
                GroupSizeEncoding dimensions(buffer, *pos, bufferLength, actingVersion);
                dimensions.blockLength((std::uint16_t)90);
                dimensions.numInGroup((std::uint16_t)count);
                m_index = 0;
                m_count = count;
                m_blockLength = 90;
                m_actingVersion = actingVersion;
                m_initialPosition = *pos;
                m_positionPtr = pos;
                *m_positionPtr = *m_positionPtr + 4;
            }

            static SBE_CONSTEXPR std::uint64_t sbeHeaderSize() SBE_NOEXCEPT
            {
                return 4;
            }

            static SBE_CONSTEXPR std::uint64_t sbeBlockLength() SBE_NOEXCEPT
            {
                return 90;
            }

            SBE_NODISCARD std::uint64_t sbePosition() const SBE_NOEXCEPT
            {
                return *m_positionPtr;
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
                *m_positionPtr = sbeCheckPosition(position);
            }

            SBE_NODISCARD inline std::uint64_t count() const SBE_NOEXCEPT
            {
                return m_count;
            }

            SBE_NODISCARD inline bool hasNext() const SBE_NOEXCEPT
            {
                return m_index < m_count;
            }

            inline QuoteEntries &next()
            {
                if (m_index >= m_count)
                {
                    throw std::runtime_error("index >= count [E108]");
                }
                m_offset = *m_positionPtr;
                if (SBE_BOUNDS_CHECK_EXPECT(((m_offset + m_blockLength) > m_bufferLength), false))
                {
                    throw std::runtime_error("buffer too short for next group index [E108]");
                }
                *m_positionPtr = m_offset + m_blockLength;
                ++m_index;

                return *this;
            }

            inline std::uint64_t resetCountToIndex() SBE_NOEXCEPT
            {
                m_count = m_index;
                GroupSizeEncoding dimensions(m_buffer, m_initialPosition, m_bufferLength, m_actingVersion);
                dimensions.numInGroup((std::uint16_t)m_count);
                return m_count;
            }

        #if __cplusplus < 201103L
            template<class Func> inline void forEach(Func& func)
            {
                while (hasNext())
                {
                    next();
                    func(*this);
                }
            }

        #else
            template<class Func> inline void forEach(Func&& func)
            {
                while (hasNext())
                {
                    next();
                    func(*this);
                }
            }

        #endif

            SBE_NODISCARD static const char * QuoteEntryIDMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
            {
                switch (metaAttribute)
                {
                    case MetaAttribute::SEMANTIC_TYPE: return "String";
                    case MetaAttribute::PRESENCE: return "required";
                    default: return "";
                }
            }

            static SBE_CONSTEXPR std::uint16_t quoteEntryIDId() SBE_NOEXCEPT
            {
                return 299;
            }

            SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t quoteEntryIDSinceVersion() SBE_NOEXCEPT
            {
                return 0;
            }

            SBE_NODISCARD bool quoteEntryIDInActingVersion() SBE_NOEXCEPT
            {
        #if defined(__clang__)
        #pragma clang diagnostic push
        #pragma clang diagnostic ignored "-Wtautological-compare"
        #endif
                return m_actingVersion >= quoteEntryIDSinceVersion();
        #if defined(__clang__)
        #pragma clang diagnostic pop
        #endif
            }

            SBE_NODISCARD static SBE_CONSTEXPR std::size_t quoteEntryIDEncodingOffset() SBE_NOEXCEPT
            {
                return 0;
            }

            static SBE_CONSTEXPR char quoteEntryIDNullValue() SBE_NOEXCEPT
            {
                return (char)0;
            }

            static SBE_CONSTEXPR char quoteEntryIDMinValue() SBE_NOEXCEPT
            {
                return (char)32;
            }

            static SBE_CONSTEXPR char quoteEntryIDMaxValue() SBE_NOEXCEPT
            {
                return (char)126;
            }

            static SBE_CONSTEXPR std::size_t quoteEntryIDEncodingLength() SBE_NOEXCEPT
            {
                return 10;
            }

            static SBE_CONSTEXPR std::uint64_t quoteEntryIDLength() SBE_NOEXCEPT
            {
                return 10;
            }

            SBE_NODISCARD const char *quoteEntryID() const SBE_NOEXCEPT
            {
                return m_buffer + m_offset + 0;
            }

            SBE_NODISCARD char *quoteEntryID() SBE_NOEXCEPT
            {
                return m_buffer + m_offset + 0;
            }

            SBE_NODISCARD char quoteEntryID(const std::uint64_t index) const
            {
                if (index >= 10)
                {
                    throw std::runtime_error("index out of range for quoteEntryID [E104]");
                }

                char val;
                std::memcpy(&val, m_buffer + m_offset + 0 + (index * 1), sizeof(char));
                return (val);
            }

            QuoteEntries &quoteEntryID(const std::uint64_t index, const char value)
            {
                if (index >= 10)
                {
                    throw std::runtime_error("index out of range for quoteEntryID [E105]");
                }

                char val = (value);
                std::memcpy(m_buffer + m_offset + 0 + (index * 1), &val, sizeof(char));
                return *this;
            }

            std::uint64_t getQuoteEntryID(char *const dst, const std::uint64_t length) const
            {
                if (length > 10)
                {
                    throw std::runtime_error("length too large for getQuoteEntryID [E106]");
                }

                std::memcpy(dst, m_buffer + m_offset + 0, sizeof(char) * static_cast<size_t>(length));
                return length;
            }

            QuoteEntries &putQuoteEntryID(const char *const src) SBE_NOEXCEPT
            {
                std::memcpy(m_buffer + m_offset + 0, src, sizeof(char) * 10);
                return *this;
            }

            SBE_NODISCARD std::string getQuoteEntryIDAsString() const
            {
                const char *buffer = m_buffer + m_offset + 0;
                size_t length = 0;

                for (; length < 10 && *(buffer + length) != '\0'; ++length);
                std::string result(buffer, length);

                return result;
            }

            std::string getQuoteEntryIDAsJsonEscapedString()
            {
                std::ostringstream oss;
                std::string s = getQuoteEntryIDAsString();

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
            SBE_NODISCARD std::string_view getQuoteEntryIDAsStringView() const SBE_NOEXCEPT
            {
                const char *buffer = m_buffer + m_offset + 0;
                size_t length = 0;

                for (; length < 10 && *(buffer + length) != '\0'; ++length);
                std::string_view result(buffer, length);

                return result;
            }
            #endif

            #if __cplusplus >= 201703L
            QuoteEntries &putQuoteEntryID(const std::string_view str)
            {
                const size_t srcLength = str.length();
                if (srcLength > 10)
                {
                    throw std::runtime_error("string too large for putQuoteEntryID [E106]");
                }

                std::memcpy(m_buffer + m_offset + 0, str.data(), srcLength);
                for (size_t start = srcLength; start < 10; ++start)
                {
                    m_buffer[m_offset + 0 + start] = 0;
                }

                return *this;
            }
            #else
            QuoteEntries &putQuoteEntryID(const std::string& str)
            {
                const size_t srcLength = str.length();
                if (srcLength > 10)
                {
                    throw std::runtime_error("string too large for putQuoteEntryID [E106]");
                }

                std::memcpy(m_buffer + m_offset + 0, str.c_str(), srcLength);
                for (size_t start = srcLength; start < 10; ++start)
                {
                    m_buffer[m_offset + 0 + start] = 0;
                }

                return *this;
            }
            #endif

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
                return 10;
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
                return m_buffer + m_offset + 10;
            }

            SBE_NODISCARD char *symbol() SBE_NOEXCEPT
            {
                return m_buffer + m_offset + 10;
            }

            SBE_NODISCARD char symbol(const std::uint64_t index) const
            {
                if (index >= 6)
                {
                    throw std::runtime_error("index out of range for symbol [E104]");
                }

                char val;
                std::memcpy(&val, m_buffer + m_offset + 10 + (index * 1), sizeof(char));
                return (val);
            }

            QuoteEntries &symbol(const std::uint64_t index, const char value)
            {
                if (index >= 6)
                {
                    throw std::runtime_error("index out of range for symbol [E105]");
                }

                char val = (value);
                std::memcpy(m_buffer + m_offset + 10 + (index * 1), &val, sizeof(char));
                return *this;
            }

            std::uint64_t getSymbol(char *const dst, const std::uint64_t length) const
            {
                if (length > 6)
                {
                    throw std::runtime_error("length too large for getSymbol [E106]");
                }

                std::memcpy(dst, m_buffer + m_offset + 10, sizeof(char) * static_cast<size_t>(length));
                return length;
            }

            QuoteEntries &putSymbol(const char *const src) SBE_NOEXCEPT
            {
                std::memcpy(m_buffer + m_offset + 10, src, sizeof(char) * 6);
                return *this;
            }

            SBE_NODISCARD std::string getSymbolAsString() const
            {
                const char *buffer = m_buffer + m_offset + 10;
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
                const char *buffer = m_buffer + m_offset + 10;
                size_t length = 0;

                for (; length < 6 && *(buffer + length) != '\0'; ++length);
                std::string_view result(buffer, length);

                return result;
            }
            #endif

            #if __cplusplus >= 201703L
            QuoteEntries &putSymbol(const std::string_view str)
            {
                const size_t srcLength = str.length();
                if (srcLength > 6)
                {
                    throw std::runtime_error("string too large for putSymbol [E106]");
                }

                std::memcpy(m_buffer + m_offset + 10, str.data(), srcLength);
                for (size_t start = srcLength; start < 6; ++start)
                {
                    m_buffer[m_offset + 10 + start] = 0;
                }

                return *this;
            }
            #else
            QuoteEntries &putSymbol(const std::string& str)
            {
                const size_t srcLength = str.length();
                if (srcLength > 6)
                {
                    throw std::runtime_error("string too large for putSymbol [E106]");
                }

                std::memcpy(m_buffer + m_offset + 10, str.c_str(), srcLength);
                for (size_t start = srcLength; start < 6; ++start)
                {
                    m_buffer[m_offset + 10 + start] = 0;
                }

                return *this;
            }
            #endif

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
                return 16;
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
                return m_buffer + m_offset + 16;
            }

            SBE_NODISCARD char *securityDesc() SBE_NOEXCEPT
            {
                return m_buffer + m_offset + 16;
            }

            SBE_NODISCARD char securityDesc(const std::uint64_t index) const
            {
                if (index >= 20)
                {
                    throw std::runtime_error("index out of range for securityDesc [E104]");
                }

                char val;
                std::memcpy(&val, m_buffer + m_offset + 16 + (index * 1), sizeof(char));
                return (val);
            }

            QuoteEntries &securityDesc(const std::uint64_t index, const char value)
            {
                if (index >= 20)
                {
                    throw std::runtime_error("index out of range for securityDesc [E105]");
                }

                char val = (value);
                std::memcpy(m_buffer + m_offset + 16 + (index * 1), &val, sizeof(char));
                return *this;
            }

            std::uint64_t getSecurityDesc(char *const dst, const std::uint64_t length) const
            {
                if (length > 20)
                {
                    throw std::runtime_error("length too large for getSecurityDesc [E106]");
                }

                std::memcpy(dst, m_buffer + m_offset + 16, sizeof(char) * static_cast<size_t>(length));
                return length;
            }

            QuoteEntries &putSecurityDesc(const char *const src) SBE_NOEXCEPT
            {
                std::memcpy(m_buffer + m_offset + 16, src, sizeof(char) * 20);
                return *this;
            }

            SBE_NODISCARD std::string getSecurityDescAsString() const
            {
                const char *buffer = m_buffer + m_offset + 16;
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
                const char *buffer = m_buffer + m_offset + 16;
                size_t length = 0;

                for (; length < 20 && *(buffer + length) != '\0'; ++length);
                std::string_view result(buffer, length);

                return result;
            }
            #endif

            #if __cplusplus >= 201703L
            QuoteEntries &putSecurityDesc(const std::string_view str)
            {
                const size_t srcLength = str.length();
                if (srcLength > 20)
                {
                    throw std::runtime_error("string too large for putSecurityDesc [E106]");
                }

                std::memcpy(m_buffer + m_offset + 16, str.data(), srcLength);
                for (size_t start = srcLength; start < 20; ++start)
                {
                    m_buffer[m_offset + 16 + start] = 0;
                }

                return *this;
            }
            #else
            QuoteEntries &putSecurityDesc(const std::string& str)
            {
                const size_t srcLength = str.length();
                if (srcLength > 20)
                {
                    throw std::runtime_error("string too large for putSecurityDesc [E106]");
                }

                std::memcpy(m_buffer + m_offset + 16, str.c_str(), srcLength);
                for (size_t start = srcLength; start < 20; ++start)
                {
                    m_buffer[m_offset + 16 + start] = 0;
                }

                return *this;
            }
            #endif

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
                return 36;
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
                return m_buffer + m_offset + 36;
            }

            SBE_NODISCARD char *securityType() SBE_NOEXCEPT
            {
                return m_buffer + m_offset + 36;
            }

            SBE_NODISCARD char securityType(const std::uint64_t index) const
            {
                if (index >= 3)
                {
                    throw std::runtime_error("index out of range for securityType [E104]");
                }

                char val;
                std::memcpy(&val, m_buffer + m_offset + 36 + (index * 1), sizeof(char));
                return (val);
            }

            QuoteEntries &securityType(const std::uint64_t index, const char value)
            {
                if (index >= 3)
                {
                    throw std::runtime_error("index out of range for securityType [E105]");
                }

                char val = (value);
                std::memcpy(m_buffer + m_offset + 36 + (index * 1), &val, sizeof(char));
                return *this;
            }

            std::uint64_t getSecurityType(char *const dst, const std::uint64_t length) const
            {
                if (length > 3)
                {
                    throw std::runtime_error("length too large for getSecurityType [E106]");
                }

                std::memcpy(dst, m_buffer + m_offset + 36, sizeof(char) * static_cast<size_t>(length));
                return length;
            }

            QuoteEntries &putSecurityType(const char *const src) SBE_NOEXCEPT
            {
                std::memcpy(m_buffer + m_offset + 36, src, sizeof(char) * 3);
                return *this;
            }

            QuoteEntries &putSecurityType(
                const char value0,
                const char value1,
                const char value2) SBE_NOEXCEPT
            {
                char val0 = (value0);
                std::memcpy(m_buffer + m_offset + 36, &val0, sizeof(char));
                char val1 = (value1);
                std::memcpy(m_buffer + m_offset + 37, &val1, sizeof(char));
                char val2 = (value2);
                std::memcpy(m_buffer + m_offset + 38, &val2, sizeof(char));

                return *this;
            }

            SBE_NODISCARD std::string getSecurityTypeAsString() const
            {
                const char *buffer = m_buffer + m_offset + 36;
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
                const char *buffer = m_buffer + m_offset + 36;
                size_t length = 0;

                for (; length < 3 && *(buffer + length) != '\0'; ++length);
                std::string_view result(buffer, length);

                return result;
            }
            #endif

            #if __cplusplus >= 201703L
            QuoteEntries &putSecurityType(const std::string_view str)
            {
                const size_t srcLength = str.length();
                if (srcLength > 3)
                {
                    throw std::runtime_error("string too large for putSecurityType [E106]");
                }

                std::memcpy(m_buffer + m_offset + 36, str.data(), srcLength);
                for (size_t start = srcLength; start < 3; ++start)
                {
                    m_buffer[m_offset + 36 + start] = 0;
                }

                return *this;
            }
            #else
            QuoteEntries &putSecurityType(const std::string& str)
            {
                const size_t srcLength = str.length();
                if (srcLength > 3)
                {
                    throw std::runtime_error("string too large for putSecurityType [E106]");
                }

                std::memcpy(m_buffer + m_offset + 36, str.c_str(), srcLength);
                for (size_t start = srcLength; start < 3; ++start)
                {
                    m_buffer[m_offset + 36 + start] = 0;
                }

                return *this;
            }
            #endif

            SBE_NODISCARD static const char * SecurityIDMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
            {
                switch (metaAttribute)
                {
                    case MetaAttribute::SEMANTIC_TYPE: return "int";
                    case MetaAttribute::PRESENCE: return "optional";
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
                return 39;
            }

            static SBE_CONSTEXPR std::int64_t securityIDNullValue() SBE_NOEXCEPT
            {
                return SBE_NULLVALUE_INT64;
            }

            static SBE_CONSTEXPR std::int64_t securityIDMinValue() SBE_NOEXCEPT
            {
                return -9223372036854775807L;
            }

            static SBE_CONSTEXPR std::int64_t securityIDMaxValue() SBE_NOEXCEPT
            {
                return 9223372036854775807L;
            }

            static SBE_CONSTEXPR std::size_t securityIDEncodingLength() SBE_NOEXCEPT
            {
                return 8;
            }

            SBE_NODISCARD std::int64_t securityID() const SBE_NOEXCEPT
            {
                std::int64_t val;
                std::memcpy(&val, m_buffer + m_offset + 39, sizeof(std::int64_t));
                return SBE_LITTLE_ENDIAN_ENCODE_64(val);
            }

            QuoteEntries &securityID(const std::int64_t value) SBE_NOEXCEPT
            {
                std::int64_t val = SBE_LITTLE_ENDIAN_ENCODE_64(value);
                std::memcpy(m_buffer + m_offset + 39, &val, sizeof(std::int64_t));
                return *this;
            }

            SBE_NODISCARD static const char * SecurityIDSourceMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
            {
                switch (metaAttribute)
                {
                    case MetaAttribute::PRESENCE: return "optional";
                    default: return "";
                }
            }

            static SBE_CONSTEXPR std::uint16_t securityIDSourceId() SBE_NOEXCEPT
            {
                return 22;
            }

            SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t securityIDSourceSinceVersion() SBE_NOEXCEPT
            {
                return 0;
            }

            SBE_NODISCARD bool securityIDSourceInActingVersion() SBE_NOEXCEPT
            {
        #if defined(__clang__)
        #pragma clang diagnostic push
        #pragma clang diagnostic ignored "-Wtautological-compare"
        #endif
                return m_actingVersion >= securityIDSourceSinceVersion();
        #if defined(__clang__)
        #pragma clang diagnostic pop
        #endif
            }

            SBE_NODISCARD static SBE_CONSTEXPR std::size_t securityIDSourceEncodingOffset() SBE_NOEXCEPT
            {
                return 47;
            }

            SBE_NODISCARD static SBE_CONSTEXPR std::size_t securityIDSourceEncodingLength() SBE_NOEXCEPT
            {
                return 1;
            }

            SBE_NODISCARD SecurityIDSource::Value securityIDSource() const SBE_NOEXCEPT
            {
                char val;
                std::memcpy(&val, m_buffer + m_offset + 47, sizeof(char));
                return SecurityIDSource::get((val));
            }

            QuoteEntries &securityIDSource(const SecurityIDSource::Value value) SBE_NOEXCEPT
            {
                char val = (value);
                std::memcpy(m_buffer + m_offset + 47, &val, sizeof(char));
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
                return 48;
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
                std::memcpy(&val, m_buffer + m_offset + 48, sizeof(std::uint64_t));
                return SBE_LITTLE_ENDIAN_ENCODE_64(val);
            }

            QuoteEntries &transactTime(const std::uint64_t value) SBE_NOEXCEPT
            {
                std::uint64_t val = SBE_LITTLE_ENDIAN_ENCODE_64(value);
                std::memcpy(m_buffer + m_offset + 48, &val, sizeof(std::uint64_t));
                return *this;
            }

            SBE_NODISCARD static const char * BidPxMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
            {
                switch (metaAttribute)
                {
                    case MetaAttribute::SEMANTIC_TYPE: return "Price";
                    case MetaAttribute::PRESENCE: return "required";
                    default: return "";
                }
            }

            static SBE_CONSTEXPR std::uint16_t bidPxId() SBE_NOEXCEPT
            {
                return 132;
            }

            SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t bidPxSinceVersion() SBE_NOEXCEPT
            {
                return 0;
            }

            SBE_NODISCARD bool bidPxInActingVersion() SBE_NOEXCEPT
            {
        #if defined(__clang__)
        #pragma clang diagnostic push
        #pragma clang diagnostic ignored "-Wtautological-compare"
        #endif
                return m_actingVersion >= bidPxSinceVersion();
        #if defined(__clang__)
        #pragma clang diagnostic pop
        #endif
            }

            SBE_NODISCARD static SBE_CONSTEXPR std::size_t bidPxEncodingOffset() SBE_NOEXCEPT
            {
                return 56;
            }

private:
            OptionalPrice m_bidPx;

public:
            SBE_NODISCARD OptionalPrice &bidPx()
            {
                m_bidPx.wrap(m_buffer, m_offset + 56, m_actingVersion, m_bufferLength);
                return m_bidPx;
            }

            SBE_NODISCARD static const char * BidSizeMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
            {
                switch (metaAttribute)
                {
                    case MetaAttribute::SEMANTIC_TYPE: return "int";
                    case MetaAttribute::PRESENCE: return "optional";
                    default: return "";
                }
            }

            static SBE_CONSTEXPR std::uint16_t bidSizeId() SBE_NOEXCEPT
            {
                return 134;
            }

            SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t bidSizeSinceVersion() SBE_NOEXCEPT
            {
                return 0;
            }

            SBE_NODISCARD bool bidSizeInActingVersion() SBE_NOEXCEPT
            {
        #if defined(__clang__)
        #pragma clang diagnostic push
        #pragma clang diagnostic ignored "-Wtautological-compare"
        #endif
                return m_actingVersion >= bidSizeSinceVersion();
        #if defined(__clang__)
        #pragma clang diagnostic pop
        #endif
            }

            SBE_NODISCARD static SBE_CONSTEXPR std::size_t bidSizeEncodingOffset() SBE_NOEXCEPT
            {
                return 65;
            }

            static SBE_CONSTEXPR std::int64_t bidSizeNullValue() SBE_NOEXCEPT
            {
                return SBE_NULLVALUE_INT64;
            }

            static SBE_CONSTEXPR std::int64_t bidSizeMinValue() SBE_NOEXCEPT
            {
                return -9223372036854775807L;
            }

            static SBE_CONSTEXPR std::int64_t bidSizeMaxValue() SBE_NOEXCEPT
            {
                return 9223372036854775807L;
            }

            static SBE_CONSTEXPR std::size_t bidSizeEncodingLength() SBE_NOEXCEPT
            {
                return 8;
            }

            SBE_NODISCARD std::int64_t bidSize() const SBE_NOEXCEPT
            {
                std::int64_t val;
                std::memcpy(&val, m_buffer + m_offset + 65, sizeof(std::int64_t));
                return SBE_LITTLE_ENDIAN_ENCODE_64(val);
            }

            QuoteEntries &bidSize(const std::int64_t value) SBE_NOEXCEPT
            {
                std::int64_t val = SBE_LITTLE_ENDIAN_ENCODE_64(value);
                std::memcpy(m_buffer + m_offset + 65, &val, sizeof(std::int64_t));
                return *this;
            }

            SBE_NODISCARD static const char * OfferPxMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
            {
                switch (metaAttribute)
                {
                    case MetaAttribute::SEMANTIC_TYPE: return "Price";
                    case MetaAttribute::PRESENCE: return "required";
                    default: return "";
                }
            }

            static SBE_CONSTEXPR std::uint16_t offerPxId() SBE_NOEXCEPT
            {
                return 133;
            }

            SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t offerPxSinceVersion() SBE_NOEXCEPT
            {
                return 0;
            }

            SBE_NODISCARD bool offerPxInActingVersion() SBE_NOEXCEPT
            {
        #if defined(__clang__)
        #pragma clang diagnostic push
        #pragma clang diagnostic ignored "-Wtautological-compare"
        #endif
                return m_actingVersion >= offerPxSinceVersion();
        #if defined(__clang__)
        #pragma clang diagnostic pop
        #endif
            }

            SBE_NODISCARD static SBE_CONSTEXPR std::size_t offerPxEncodingOffset() SBE_NOEXCEPT
            {
                return 73;
            }

private:
            OptionalPrice m_offerPx;

public:
            SBE_NODISCARD OptionalPrice &offerPx()
            {
                m_offerPx.wrap(m_buffer, m_offset + 73, m_actingVersion, m_bufferLength);
                return m_offerPx;
            }

            SBE_NODISCARD static const char * OfferSizeMetaAttribute(const MetaAttribute metaAttribute) SBE_NOEXCEPT
            {
                switch (metaAttribute)
                {
                    case MetaAttribute::SEMANTIC_TYPE: return "int";
                    case MetaAttribute::PRESENCE: return "optional";
                    default: return "";
                }
            }

            static SBE_CONSTEXPR std::uint16_t offerSizeId() SBE_NOEXCEPT
            {
                return 135;
            }

            SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t offerSizeSinceVersion() SBE_NOEXCEPT
            {
                return 0;
            }

            SBE_NODISCARD bool offerSizeInActingVersion() SBE_NOEXCEPT
            {
        #if defined(__clang__)
        #pragma clang diagnostic push
        #pragma clang diagnostic ignored "-Wtautological-compare"
        #endif
                return m_actingVersion >= offerSizeSinceVersion();
        #if defined(__clang__)
        #pragma clang diagnostic pop
        #endif
            }

            SBE_NODISCARD static SBE_CONSTEXPR std::size_t offerSizeEncodingOffset() SBE_NOEXCEPT
            {
                return 82;
            }

            static SBE_CONSTEXPR std::int64_t offerSizeNullValue() SBE_NOEXCEPT
            {
                return SBE_NULLVALUE_INT64;
            }

            static SBE_CONSTEXPR std::int64_t offerSizeMinValue() SBE_NOEXCEPT
            {
                return -9223372036854775807L;
            }

            static SBE_CONSTEXPR std::int64_t offerSizeMaxValue() SBE_NOEXCEPT
            {
                return 9223372036854775807L;
            }

            static SBE_CONSTEXPR std::size_t offerSizeEncodingLength() SBE_NOEXCEPT
            {
                return 8;
            }

            SBE_NODISCARD std::int64_t offerSize() const SBE_NOEXCEPT
            {
                std::int64_t val;
                std::memcpy(&val, m_buffer + m_offset + 82, sizeof(std::int64_t));
                return SBE_LITTLE_ENDIAN_ENCODE_64(val);
            }

            QuoteEntries &offerSize(const std::int64_t value) SBE_NOEXCEPT
            {
                std::int64_t val = SBE_LITTLE_ENDIAN_ENCODE_64(value);
                std::memcpy(m_buffer + m_offset + 82, &val, sizeof(std::int64_t));
                return *this;
            }

            template<typename CharT, typename Traits>
            friend std::basic_ostream<CharT, Traits>& operator << (
                std::basic_ostream<CharT, Traits>& builder, QuoteEntries writer)
            {
                builder << '{';
                builder << R"("QuoteEntryID": )";
                builder << '"' <<
                    writer.getQuoteEntryIDAsJsonEscapedString().c_str() << '"';

                builder << ", ";
                builder << R"("Symbol": )";
                builder << '"' <<
                    writer.getSymbolAsJsonEscapedString().c_str() << '"';

                builder << ", ";
                builder << R"("SecurityDesc": )";
                builder << '"' <<
                    writer.getSecurityDescAsJsonEscapedString().c_str() << '"';

                builder << ", ";
                builder << R"("SecurityType": )";
                builder << '"' <<
                    writer.getSecurityTypeAsJsonEscapedString().c_str() << '"';

                builder << ", ";
                builder << R"("SecurityID": )";
                builder << +writer.securityID();

                builder << ", ";
                builder << R"("SecurityIDSource": )";
                builder << '"' << writer.securityIDSource() << '"';

                builder << ", ";
                builder << R"("TransactTime": )";
                builder << +writer.transactTime();

                builder << ", ";
                builder << R"("BidPx": )";
                builder << writer.bidPx();

                builder << ", ";
                builder << R"("BidSize": )";
                builder << +writer.bidSize();

                builder << ", ";
                builder << R"("OfferPx": )";
                builder << writer.offerPx();

                builder << ", ";
                builder << R"("OfferSize": )";
                builder << +writer.offerSize();

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

private:
        QuoteEntries m_quoteEntries;

public:
        SBE_NODISCARD static SBE_CONSTEXPR std::uint16_t QuoteEntriesId() SBE_NOEXCEPT
        {
            return 295;
        }

        SBE_NODISCARD inline QuoteEntries &quoteEntries()
        {
            m_quoteEntries.wrapForDecode(m_buffer, sbePositionPtr(), m_actingVersion, m_bufferLength);
            return m_quoteEntries;
        }

        QuoteEntries &quoteEntriesCount(const std::uint16_t count)
        {
            m_quoteEntries.wrapForEncode(m_buffer, count, sbePositionPtr(), m_actingVersion, m_bufferLength);
            return m_quoteEntries;
        }

        SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t quoteEntriesSinceVersion() SBE_NOEXCEPT
        {
            return 0;
        }

        SBE_NODISCARD bool quoteEntriesInActingVersion() const SBE_NOEXCEPT
        {
    #if defined(__clang__)
    #pragma clang diagnostic push
    #pragma clang diagnostic ignored "-Wtautological-compare"
    #endif
            return m_actingVersion >= quoteEntriesSinceVersion();
    #if defined(__clang__)
    #pragma clang diagnostic pop
    #endif
        }

        template<typename CharT, typename Traits>
        friend std::basic_ostream<CharT, Traits>& operator << (
            std::basic_ostream<CharT, Traits>& builder, QuoteSets writer)
        {
            builder << '{';
            builder << R"("QuoteSetID": )";
            builder << '"' <<
                writer.getQuoteSetIDAsJsonEscapedString().c_str() << '"';

            builder << ", ";
            builder << R"("UnderlyingSecurityDesc": )";
            builder << '"' <<
                writer.getUnderlyingSecurityDescAsJsonEscapedString().c_str() << '"';

            builder << ", ";
            builder << R"("TotQuoteEntries": )";
            builder << +writer.totQuoteEntries();

            builder << ", ";
            {
                bool atLeastOne = false;
                builder << R"("QuoteEntries": [)";
                writer.quoteEntries().forEach([&](QuoteEntries& quoteEntries)
                {
                    if (atLeastOne)
                    {
                        builder << ", ";
                    }
                    atLeastOne = true;
                    builder << quoteEntries;
                });
                builder << ']';
            }

            builder << '}';

            return builder;
        }

        void skip()
        {
            quoteEntries().forEach([](QuoteEntries e)
            {
                e.skip();
            });
        }

        SBE_NODISCARD static SBE_CONSTEXPR bool isConstLength() SBE_NOEXCEPT
        {
            return false;
        }

        SBE_NODISCARD static SBE_CONSTEXPR_14 size_t computeLength(std::size_t QuoteEntriesLength = 0)
        {
#if defined(__GNUG__) && !defined(__clang__)
#pragma GCC diagnostic push
#pragma GCC diagnostic ignored "-Wtype-limits"
#endif
            size_t length = sbeBlockLength();


            length += QuoteEntries::sbeHeaderSize();
            if (QuoteEntriesLength > 65534)
            {
                throw std::runtime_error("QuoteEntriesLength outside of allowed range [E110]");
            }
            length += QuoteEntriesLength * QuoteEntries::sbeBlockLength();
            return length;
#if defined(__GNUG__) && !defined(__clang__)
#pragma GCC diagnostic pop
#endif
        }
    };

private:
    QuoteSets m_quoteSets;

public:
    SBE_NODISCARD static SBE_CONSTEXPR std::uint16_t QuoteSetsId() SBE_NOEXCEPT
    {
        return 296;
    }

    SBE_NODISCARD inline QuoteSets &quoteSets()
    {
        m_quoteSets.wrapForDecode(m_buffer, sbePositionPtr(), m_actingVersion, m_bufferLength);
        return m_quoteSets;
    }

    QuoteSets &quoteSetsCount(const std::uint16_t count)
    {
        m_quoteSets.wrapForEncode(m_buffer, count, sbePositionPtr(), m_actingVersion, m_bufferLength);
        return m_quoteSets;
    }

    SBE_NODISCARD static SBE_CONSTEXPR std::uint64_t quoteSetsSinceVersion() SBE_NOEXCEPT
    {
        return 0;
    }

    SBE_NODISCARD bool quoteSetsInActingVersion() const SBE_NOEXCEPT
    {
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
        return m_actingVersion >= quoteSetsSinceVersion();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
    }

template<typename CharT, typename Traits>
friend std::basic_ostream<CharT, Traits>& operator << (
    std::basic_ostream<CharT, Traits>& builder, MassQuote _writer)
{
    MassQuote writer(_writer.m_buffer, _writer.m_offset,
        _writer.m_bufferLength, _writer.sbeBlockLength(), _writer.m_actingVersion);
    builder << '{';
    builder << R"("Name": "MassQuote", )";
    builder << R"("sbeTemplateId": )";
    builder << writer.sbeTemplateId();
    builder << ", ";

    builder << R"("QuoteReqID": )";
    builder << '"' <<
        writer.getQuoteReqIDAsJsonEscapedString().c_str() << '"';

    builder << ", ";
    builder << R"("QuoteID": )";
    builder << '"' <<
        writer.getQuoteIDAsJsonEscapedString().c_str() << '"';

    builder << ", ";
    builder << R"("MMAccount": )";
    builder << '"' <<
        writer.getMMAccountAsJsonEscapedString().c_str() << '"';

    builder << ", ";
    builder << R"("ManualOrderIndicator": )";
    builder << '"' << writer.manualOrderIndicator() << '"';

    builder << ", ";
    builder << R"("CustOrderHandlingInst": )";
    builder << '"' << writer.custOrderHandlingInst() << '"';

    builder << ", ";
    builder << R"("CustomerOrFirm": )";
    builder << '"' << writer.customerOrFirm() << '"';

    builder << ", ";
    builder << R"("SelfMatchPreventionID": )";
    builder << '"' <<
        writer.getSelfMatchPreventionIDAsJsonEscapedString().c_str() << '"';

    builder << ", ";
    builder << R"("CtiCode": )";
    builder << '"' << writer.ctiCode() << '"';

    builder << ", ";
    builder << R"("MMProtectionReset": )";
    builder << '"' << writer.mMProtectionReset() << '"';

    builder << ", ";
    {
        bool atLeastOne = false;
        builder << R"("QuoteSets": [)";
        writer.quoteSets().forEach([&](QuoteSets& quoteSets)
        {
            if (atLeastOne)
            {
                builder << ", ";
            }
            atLeastOne = true;
            builder << quoteSets;
        });
        builder << ']';
    }

    builder << '}';

    return builder;
}

void skip()
{
    quoteSets().forEach([](QuoteSets e)
    {
        e.skip();
    });
}

SBE_NODISCARD static SBE_CONSTEXPR bool isConstLength() SBE_NOEXCEPT
{
    return false;
}

SBE_NODISCARD static SBE_CONSTEXPR_14 size_t computeLength(const std::vector<std::tuple<std::size_t>>& QuoteSetsItemLengths = {})
{
#if defined(__GNUG__) && !defined(__clang__)
#pragma GCC diagnostic push
#pragma GCC diagnostic ignored "-Wtype-limits"
#endif
    size_t length = sbeBlockLength();


    length += QuoteSets::sbeHeaderSize();
    if (QuoteSetsItemLengths.size() > 65534)
    {
        throw std::runtime_error("QuoteSetsItemLengths.size() outside of allowed range [E110]");
    }
    for (const auto& e: QuoteSetsItemLengths)
    {
        #if __cpluplus >= 201703L
        length += std::apply(QuoteSets::computeLength, e);
        #else
        length += QuoteSets::computeLength(std::get<0>(e));
        #endif
    }
    return length;
#if defined(__GNUG__) && !defined(__clang__)
#pragma GCC diagnostic pop
#endif
}
};
}
#endif
