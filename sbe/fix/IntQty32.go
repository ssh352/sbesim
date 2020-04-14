// Generated SBE (Simple Binary Encoding) message codec

package fix

import (
	"fmt"
	"io"
	"math"
)

type IntQty32 struct {
	Mantissa int32
	Exponent int8
}

func (i *IntQty32) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt32(_w, i.Mantissa); err != nil {
		return err
	}
	return nil
}

func (i *IntQty32) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if !i.MantissaInActingVersion(actingVersion) {
		i.Mantissa = i.MantissaNullValue()
	} else {
		if err := _m.ReadInt32(_r, &i.Mantissa); err != nil {
			return err
		}
	}
	i.Exponent = 0
	return nil
}

func (i *IntQty32) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if i.MantissaInActingVersion(actingVersion) {
		if i.Mantissa < i.MantissaMinValue() || i.Mantissa > i.MantissaMaxValue() {
			return fmt.Errorf("Range check failed on i.Mantissa (%v < %v > %v)", i.MantissaMinValue(), i.Mantissa, i.MantissaMaxValue())
		}
	}
	return nil
}

func IntQty32Init(i *IntQty32) {
	i.Exponent = 0
	return
}

func (*IntQty32) EncodedLength() int64 {
	return 4
}

func (*IntQty32) MantissaMinValue() int32 {
	return math.MinInt32 + 1
}

func (*IntQty32) MantissaMaxValue() int32 {
	return math.MaxInt32
}

func (*IntQty32) MantissaNullValue() int32 {
	return math.MinInt32
}

func (*IntQty32) MantissaSinceVersion() uint16 {
	return 0
}

func (i *IntQty32) MantissaInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.MantissaSinceVersion()
}

func (*IntQty32) MantissaDeprecated() uint16 {
	return 0
}

func (*IntQty32) ExponentMinValue() int8 {
	return math.MinInt8 + 1
}

func (*IntQty32) ExponentMaxValue() int8 {
	return math.MaxInt8
}

func (*IntQty32) ExponentNullValue() int8 {
	return math.MinInt8
}

func (*IntQty32) ExponentSinceVersion() uint16 {
	return 0
}

func (i *IntQty32) ExponentInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.ExponentSinceVersion()
}

func (*IntQty32) ExponentDeprecated() uint16 {
	return 0
}
