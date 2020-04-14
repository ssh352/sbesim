// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"math"
)

type DATA struct {
	Length  uint16
	VarData byte
}

func (d *DATA) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint16(_w, d.Length); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, d.VarData); err != nil {
		return err
	}
	return nil
}

func (d *DATA) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if !d.LengthInActingVersion(actingVersion) {
		d.Length = d.LengthNullValue()
	} else {
		if err := _m.ReadUint16(_r, &d.Length); err != nil {
			return err
		}
	}
	if !d.VarDataInActingVersion(actingVersion) {
		d.VarData = d.VarDataNullValue()
	} else {
		if err := _m.ReadUint8(_r, &d.VarData); err != nil {
			return err
		}
	}
	return nil
}

func (d *DATA) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if d.LengthInActingVersion(actingVersion) {
		if d.Length < d.LengthMinValue() || d.Length > d.LengthMaxValue() {
			return fmt.Errorf("Range check failed on d.Length (%v < %v > %v)", d.LengthMinValue(), d.Length, d.LengthMaxValue())
		}
	}
	if d.VarDataInActingVersion(actingVersion) {
		if d.VarData < d.VarDataMinValue() || d.VarData > d.VarDataMaxValue() {
			return fmt.Errorf("Range check failed on d.VarData (%v < %v > %v)", d.VarDataMinValue(), d.VarData, d.VarDataMaxValue())
		}
	}
	return nil
}

func DATAInit(d *DATA) {
	return
}

func (*DATA) EncodedLength() int64 {
	return -1
}

func (*DATA) LengthMinValue() uint16 {
	return 0
}

func (*DATA) LengthMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*DATA) LengthNullValue() uint16 {
	return math.MaxUint16
}

func (*DATA) LengthSinceVersion() uint16 {
	return 0
}

func (d *DATA) LengthInActingVersion(actingVersion uint16) bool {
	return actingVersion >= d.LengthSinceVersion()
}

func (*DATA) LengthDeprecated() uint16 {
	return 0
}

func (*DATA) VarDataMinValue() byte {
	return byte(32)
}

func (*DATA) VarDataMaxValue() byte {
	return byte(126)
}

func (*DATA) VarDataNullValue() byte {
	return 0
}

func (*DATA) VarDataSinceVersion() uint16 {
	return 0
}

func (d *DATA) VarDataInActingVersion(actingVersion uint16) bool {
	return actingVersion >= d.VarDataSinceVersion()
}

func (*DATA) VarDataDeprecated() uint16 {
	return 0
}
