// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"reflect"
)

type BooleanFlagEnum uint8
type BooleanFlagValues struct {
	False     BooleanFlagEnum
	True      BooleanFlagEnum
	NullValue BooleanFlagEnum
}

var BooleanFlag = BooleanFlagValues{0, 1, 255}

func (b BooleanFlagEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(b)); err != nil {
		return err
	}
	return nil
}

func (b *BooleanFlagEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(b)); err != nil {
		return err
	}
	return nil
}

func (b BooleanFlagEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(BooleanFlag)
	for idx := 0; idx < value.NumField(); idx++ {
		if b == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on BooleanFlag, unknown enumeration value %d", b)
}

func (*BooleanFlagEnum) EncodedLength() int64 {
	return 1
}

func (*BooleanFlagEnum) FalseSinceVersion() uint16 {
	return 0
}

func (b *BooleanFlagEnum) FalseInActingVersion(actingVersion uint16) bool {
	return actingVersion >= b.FalseSinceVersion()
}

func (*BooleanFlagEnum) FalseDeprecated() uint16 {
	return 0
}

func (*BooleanFlagEnum) TrueSinceVersion() uint16 {
	return 0
}

func (b *BooleanFlagEnum) TrueInActingVersion(actingVersion uint16) bool {
	return actingVersion >= b.TrueSinceVersion()
}

func (*BooleanFlagEnum) TrueDeprecated() uint16 {
	return 0
}
