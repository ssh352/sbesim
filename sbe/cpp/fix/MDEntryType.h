/* Generated SBE (Simple Binary Encoding) message codec */
#ifndef _FIX_MDENTRYTYPE_H_
#define _FIX_MDENTRYTYPE_H_

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


namespace fix {

class MDEntryType
{
public:
    enum Value
    {
        BID = (char)48,
        OFFER = (char)49,
        TRADE = (char)50,
        OPENING_PRICE = (char)52,
        SETTLEMENT_PRICE = (char)54,
        TRADING_SESSION_HIGH_PRICE = (char)55,
        TRADING_SESSION_LOW_PRICE = (char)56,
        TRADE_VOLUME = (char)66,
        OPEN_INTEREST = (char)67,
        SIMULATED_SELL = (char)69,
        SIMULATED_BUY = (char)70,
        EMPTY_THE_BOOK = (char)74,
        SESSION_HIGH_BID = (char)78,
        SESSION_LOW_OFFER = (char)79,
        FIXING_PRICE = (char)87,
        CASH_NOTE = (char)88,
        NULL_VALUE = (char)0
    };

    static MDEntryType::Value get(const char value)
    {
        switch (value)
        {
            case 48: return BID;
            case 49: return OFFER;
            case 50: return TRADE;
            case 52: return OPENING_PRICE;
            case 54: return SETTLEMENT_PRICE;
            case 55: return TRADING_SESSION_HIGH_PRICE;
            case 56: return TRADING_SESSION_LOW_PRICE;
            case 66: return TRADE_VOLUME;
            case 67: return OPEN_INTEREST;
            case 69: return SIMULATED_SELL;
            case 70: return SIMULATED_BUY;
            case 74: return EMPTY_THE_BOOK;
            case 78: return SESSION_HIGH_BID;
            case 79: return SESSION_LOW_OFFER;
            case 87: return FIXING_PRICE;
            case 88: return CASH_NOTE;
            case 0: return NULL_VALUE;
        }

        throw std::runtime_error("unknown value for enum MDEntryType [E103]");
    }

    static const char* c_str(const MDEntryType::Value value)
    {
        switch (value)
        {
            case BID: return "BID";
            case OFFER: return "OFFER";
            case TRADE: return "TRADE";
            case OPENING_PRICE: return "OPENING_PRICE";
            case SETTLEMENT_PRICE: return "SETTLEMENT_PRICE";
            case TRADING_SESSION_HIGH_PRICE: return "TRADING_SESSION_HIGH_PRICE";
            case TRADING_SESSION_LOW_PRICE: return "TRADING_SESSION_LOW_PRICE";
            case TRADE_VOLUME: return "TRADE_VOLUME";
            case OPEN_INTEREST: return "OPEN_INTEREST";
            case SIMULATED_SELL: return "SIMULATED_SELL";
            case SIMULATED_BUY: return "SIMULATED_BUY";
            case EMPTY_THE_BOOK: return "EMPTY_THE_BOOK";
            case SESSION_HIGH_BID: return "SESSION_HIGH_BID";
            case SESSION_LOW_OFFER: return "SESSION_LOW_OFFER";
            case FIXING_PRICE: return "FIXING_PRICE";
            case CASH_NOTE: return "CASH_NOTE";
            case NULL_VALUE: return "NULL_VALUE";
        }

        throw std::runtime_error("unknown value for enum MDEntryType [E103]:");
    }

    template<typename CharT, typename Traits>
    friend std::basic_ostream<CharT, Traits>& operator << (
        std::basic_ostream<CharT, Traits>& os, MDEntryType::Value m)
    {
        return os << MDEntryType::c_str(m);
    }
};

}

#endif
