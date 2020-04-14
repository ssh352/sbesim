// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"math"
)

type PRICE9 struct {
	Mantissa int64
	Exponent int8
}

func (p *PRICE9) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt64(_w, p.Mantissa); err != nil {
		return err
	}
	return nil
}

func (p *PRICE9) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
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

func (p *PRICE9) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if p.MantissaInActingVersion(actingVersion) {
		if p.Mantissa < p.MantissaMinValue() || p.Mantissa > p.MantissaMaxValue() {
			return fmt.Errorf("Range check failed on p.Mantissa (%v < %v > %v)", p.MantissaMinValue(), p.Mantissa, p.MantissaMaxValue())
		}
	}
	return nil
}

func PRICE9Init(p *PRICE9) {
	p.Exponent = -9
	return
}

func (*PRICE9) EncodedLength() int64 {
	return 8
}

func (*PRICE9) MantissaMinValue() int64 {
	return math.MinInt64 + 1
}

func (*PRICE9) MantissaMaxValue() int64 {
	return math.MaxInt64
}

func (*PRICE9) MantissaNullValue() int64 {
	return math.MinInt64
}

func (*PRICE9) MantissaSinceVersion() uint16 {
	return 0
}

func (p *PRICE9) MantissaInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.MantissaSinceVersion()
}

func (*PRICE9) MantissaDeprecated() uint16 {
	return 0
}

func (*PRICE9) ExponentMinValue() int8 {
	return math.MinInt8 + 1
}

func (*PRICE9) ExponentMaxValue() int8 {
	return math.MaxInt8
}

func (*PRICE9) ExponentNullValue() int8 {
	return math.MinInt8
}

func (*PRICE9) ExponentSinceVersion() uint16 {
	return 0
}

func (p *PRICE9) ExponentInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.ExponentSinceVersion()
}

func (*PRICE9) ExponentDeprecated() uint16 {
	return 0
}
