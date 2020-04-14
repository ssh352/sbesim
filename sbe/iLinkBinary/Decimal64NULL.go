// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"math"
)

type Decimal64NULL struct {
	Mantissa int64
	Exponent int8
}

func (d *Decimal64NULL) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt64(_w, d.Mantissa); err != nil {
		return err
	}
	if err := _m.WriteInt8(_w, d.Exponent); err != nil {
		return err
	}
	return nil
}

func (d *Decimal64NULL) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if !d.MantissaInActingVersion(actingVersion) {
		d.Mantissa = d.MantissaNullValue()
	} else {
		if err := _m.ReadInt64(_r, &d.Mantissa); err != nil {
			return err
		}
	}
	if !d.ExponentInActingVersion(actingVersion) {
		d.Exponent = d.ExponentNullValue()
	} else {
		if err := _m.ReadInt8(_r, &d.Exponent); err != nil {
			return err
		}
	}
	return nil
}

func (d *Decimal64NULL) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if d.MantissaInActingVersion(actingVersion) {
		if d.Mantissa != d.MantissaNullValue() && (d.Mantissa < d.MantissaMinValue() || d.Mantissa > d.MantissaMaxValue()) {
			return fmt.Errorf("Range check failed on d.Mantissa (%v < %v > %v)", d.MantissaMinValue(), d.Mantissa, d.MantissaMaxValue())
		}
	}
	if d.ExponentInActingVersion(actingVersion) {
		if d.Exponent != d.ExponentNullValue() && (d.Exponent < d.ExponentMinValue() || d.Exponent > d.ExponentMaxValue()) {
			return fmt.Errorf("Range check failed on d.Exponent (%v < %v > %v)", d.ExponentMinValue(), d.Exponent, d.ExponentMaxValue())
		}
	}
	return nil
}

func Decimal64NULLInit(d *Decimal64NULL) {
	d.Mantissa = int64(9223372036854775807)
	d.Exponent = 127
	return
}

func (*Decimal64NULL) EncodedLength() int64 {
	return 9
}

func (*Decimal64NULL) MantissaMinValue() int64 {
	return math.MinInt64 + 1
}

func (*Decimal64NULL) MantissaMaxValue() int64 {
	return math.MaxInt64
}

func (*Decimal64NULL) MantissaNullValue() int64 {
	return int64(9223372036854775807)
}

func (*Decimal64NULL) MantissaSinceVersion() uint16 {
	return 0
}

func (d *Decimal64NULL) MantissaInActingVersion(actingVersion uint16) bool {
	return actingVersion >= d.MantissaSinceVersion()
}

func (*Decimal64NULL) MantissaDeprecated() uint16 {
	return 0
}

func (*Decimal64NULL) ExponentMinValue() int8 {
	return math.MinInt8 + 1
}

func (*Decimal64NULL) ExponentMaxValue() int8 {
	return math.MaxInt8
}

func (*Decimal64NULL) ExponentNullValue() int8 {
	return 127
}

func (*Decimal64NULL) ExponentSinceVersion() uint16 {
	return 0
}

func (d *Decimal64NULL) ExponentInActingVersion(actingVersion uint16) bool {
	return actingVersion >= d.ExponentSinceVersion()
}

func (*Decimal64NULL) ExponentDeprecated() uint16 {
	return 0
}
