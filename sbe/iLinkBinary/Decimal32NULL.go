// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"math"
)

type Decimal32NULL struct {
	Mantissa int32
	Exponent int8
}

func (d *Decimal32NULL) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt32(_w, d.Mantissa); err != nil {
		return err
	}
	if err := _m.WriteInt8(_w, d.Exponent); err != nil {
		return err
	}
	return nil
}

func (d *Decimal32NULL) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if !d.MantissaInActingVersion(actingVersion) {
		d.Mantissa = d.MantissaNullValue()
	} else {
		if err := _m.ReadInt32(_r, &d.Mantissa); err != nil {
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

func (d *Decimal32NULL) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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

func Decimal32NULLInit(d *Decimal32NULL) {
	d.Mantissa = 2147483647
	d.Exponent = 127
	return
}

func (*Decimal32NULL) EncodedLength() int64 {
	return 5
}

func (*Decimal32NULL) MantissaMinValue() int32 {
	return math.MinInt32 + 1
}

func (*Decimal32NULL) MantissaMaxValue() int32 {
	return math.MaxInt32
}

func (*Decimal32NULL) MantissaNullValue() int32 {
	return 2147483647
}

func (*Decimal32NULL) MantissaSinceVersion() uint16 {
	return 0
}

func (d *Decimal32NULL) MantissaInActingVersion(actingVersion uint16) bool {
	return actingVersion >= d.MantissaSinceVersion()
}

func (*Decimal32NULL) MantissaDeprecated() uint16 {
	return 0
}

func (*Decimal32NULL) ExponentMinValue() int8 {
	return math.MinInt8 + 1
}

func (*Decimal32NULL) ExponentMaxValue() int8 {
	return math.MaxInt8
}

func (*Decimal32NULL) ExponentNullValue() int8 {
	return 127
}

func (*Decimal32NULL) ExponentSinceVersion() uint16 {
	return 0
}

func (d *Decimal32NULL) ExponentInActingVersion(actingVersion uint16) bool {
	return actingVersion >= d.ExponentSinceVersion()
}

func (*Decimal32NULL) ExponentDeprecated() uint16 {
	return 0
}
