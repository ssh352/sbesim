// Generated SBE (Simple Binary Encoding) message codec

package fix

import (
	"fmt"
	"io"
	"math"
)

type OptionalPrice struct {
	Mantissa int64
	Exponent int8
}

func (o *OptionalPrice) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt64(_w, o.Mantissa); err != nil {
		return err
	}
	if err := _m.WriteInt8(_w, o.Exponent); err != nil {
		return err
	}
	return nil
}

func (o *OptionalPrice) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if !o.MantissaInActingVersion(actingVersion) {
		o.Mantissa = o.MantissaNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.Mantissa); err != nil {
			return err
		}
	}
	if !o.ExponentInActingVersion(actingVersion) {
		o.Exponent = o.ExponentNullValue()
	} else {
		if err := _m.ReadInt8(_r, &o.Exponent); err != nil {
			return err
		}
	}
	return nil
}

func (o *OptionalPrice) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if o.MantissaInActingVersion(actingVersion) {
		if o.Mantissa < o.MantissaMinValue() || o.Mantissa > o.MantissaMaxValue() {
			return fmt.Errorf("Range check failed on o.Mantissa (%v < %v > %v)", o.MantissaMinValue(), o.Mantissa, o.MantissaMaxValue())
		}
	}
	if o.ExponentInActingVersion(actingVersion) {
		if o.Exponent < o.ExponentMinValue() || o.Exponent > o.ExponentMaxValue() {
			return fmt.Errorf("Range check failed on o.Exponent (%v < %v > %v)", o.ExponentMinValue(), o.Exponent, o.ExponentMaxValue())
		}
	}
	return nil
}

func OptionalPriceInit(o *OptionalPrice) {
	return
}

func (*OptionalPrice) EncodedLength() int64 {
	return 9
}

func (*OptionalPrice) MantissaMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OptionalPrice) MantissaMaxValue() int64 {
	return math.MaxInt64
}

func (*OptionalPrice) MantissaNullValue() int64 {
	return math.MinInt64
}

func (*OptionalPrice) MantissaSinceVersion() uint16 {
	return 0
}

func (o *OptionalPrice) MantissaInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.MantissaSinceVersion()
}

func (*OptionalPrice) MantissaDeprecated() uint16 {
	return 0
}

func (*OptionalPrice) ExponentMinValue() int8 {
	return math.MinInt8 + 1
}

func (*OptionalPrice) ExponentMaxValue() int8 {
	return math.MaxInt8
}

func (*OptionalPrice) ExponentNullValue() int8 {
	return math.MinInt8
}

func (*OptionalPrice) ExponentSinceVersion() uint16 {
	return 0
}

func (o *OptionalPrice) ExponentInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ExponentSinceVersion()
}

func (*OptionalPrice) ExponentDeprecated() uint16 {
	return 0
}
