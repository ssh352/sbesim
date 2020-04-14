// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"math"
)

type GroupSize struct {
	BlockLength uint16
	NumInGroup  uint8
}

func (g *GroupSize) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint16(_w, g.BlockLength); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, g.NumInGroup); err != nil {
		return err
	}
	return nil
}

func (g *GroupSize) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if !g.BlockLengthInActingVersion(actingVersion) {
		g.BlockLength = g.BlockLengthNullValue()
	} else {
		if err := _m.ReadUint16(_r, &g.BlockLength); err != nil {
			return err
		}
	}
	if !g.NumInGroupInActingVersion(actingVersion) {
		g.NumInGroup = g.NumInGroupNullValue()
	} else {
		if err := _m.ReadUint8(_r, &g.NumInGroup); err != nil {
			return err
		}
	}
	return nil
}

func (g *GroupSize) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if g.BlockLengthInActingVersion(actingVersion) {
		if g.BlockLength < g.BlockLengthMinValue() || g.BlockLength > g.BlockLengthMaxValue() {
			return fmt.Errorf("Range check failed on g.BlockLength (%v < %v > %v)", g.BlockLengthMinValue(), g.BlockLength, g.BlockLengthMaxValue())
		}
	}
	if g.NumInGroupInActingVersion(actingVersion) {
		if g.NumInGroup < g.NumInGroupMinValue() || g.NumInGroup > g.NumInGroupMaxValue() {
			return fmt.Errorf("Range check failed on g.NumInGroup (%v < %v > %v)", g.NumInGroupMinValue(), g.NumInGroup, g.NumInGroupMaxValue())
		}
	}
	return nil
}

func GroupSizeInit(g *GroupSize) {
	return
}

func (*GroupSize) EncodedLength() int64 {
	return 3
}

func (*GroupSize) BlockLengthMinValue() uint16 {
	return 0
}

func (*GroupSize) BlockLengthMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*GroupSize) BlockLengthNullValue() uint16 {
	return math.MaxUint16
}

func (*GroupSize) BlockLengthSinceVersion() uint16 {
	return 0
}

func (g *GroupSize) BlockLengthInActingVersion(actingVersion uint16) bool {
	return actingVersion >= g.BlockLengthSinceVersion()
}

func (*GroupSize) BlockLengthDeprecated() uint16 {
	return 0
}

func (*GroupSize) NumInGroupMinValue() uint8 {
	return 0
}

func (*GroupSize) NumInGroupMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*GroupSize) NumInGroupNullValue() uint8 {
	return math.MaxUint8
}

func (*GroupSize) NumInGroupSinceVersion() uint16 {
	return 0
}

func (g *GroupSize) NumInGroupInActingVersion(actingVersion uint16) bool {
	return actingVersion >= g.NumInGroupSinceVersion()
}

func (*GroupSize) NumInGroupDeprecated() uint16 {
	return 0
}
