// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"math"
)

type PRICENULL9 struct {
	Mantissa int64
	Exponent int8
}

func (p *PRICENULL9) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt64(_w, p.Mantissa); err != nil {
		return err
	}
	return nil
}

func (p *PRICENULL9) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if !p.MantissaInActingVersion(actingVersion) {
		p.Mantissa = p.MantissaNullValue()
	} else {
		if err := _m.ReadInt64(_r, &p.Mantissa); err != nil {
			return err
		}
	}
	p.Exponent = -9
	return nil
}

func (p *PRICENULL9) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if p.MantissaInActingVersion(actingVersion) {
		if p.Mantissa != p.MantissaNullValue() && (p.Mantissa < p.MantissaMinValue() || p.Mantissa > p.MantissaMaxValue()) {
			return fmt.Errorf("Range check failed on p.Mantissa (%v < %v > %v)", p.MantissaMinValue(), p.Mantissa, p.MantissaMaxValue())
		}
	}
	return nil
}

func PRICENULL9Init(p *PRICENULL9) {
	p.Mantissa = int64(9223372036854775807)
	p.Exponent = -9
	return
}

func (*PRICENULL9) EncodedLength() int64 {
	return 8
}

func (*PRICENULL9) MantissaMinValue() int64 {
	return math.MinInt64 + 1
}

func (*PRICENULL9) MantissaMaxValue() int64 {
	return math.MaxInt64
}

func (*PRICENULL9) MantissaNullValue() int64 {
	return int64(9223372036854775807)
}

func (*PRICENULL9) MantissaSinceVersion() uint16 {
	return 0
}

func (p *PRICENULL9) MantissaInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.MantissaSinceVersion()
}

func (*PRICENULL9) MantissaDeprecated() uint16 {
	return 0
}

func (*PRICENULL9) ExponentMinValue() int8 {
	return math.MinInt8 + 1
}

func (*PRICENULL9) ExponentMaxValue() int8 {
	return math.MaxInt8
}

func (*PRICENULL9) ExponentNullValue() int8 {
	return math.MinInt8
}

func (*PRICENULL9) ExponentSinceVersion() uint16 {
	return 0
}

func (p *PRICENULL9) ExponentInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.ExponentSinceVersion()
}

func (*PRICENULL9) ExponentDeprecated() uint16 {
	return 0
}
