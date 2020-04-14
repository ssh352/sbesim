// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"reflect"
)

type BooleanNULLEnum uint8
type BooleanNULLValues struct {
	False     BooleanNULLEnum
	True      BooleanNULLEnum
	NullValue BooleanNULLEnum
}

var BooleanNULL = BooleanNULLValues{0, 1, 255}

func (b BooleanNULLEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(b)); err != nil {
		return err
	}
	return nil
}

func (b *BooleanNULLEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(b)); err != nil {
		return err
	}
	return nil
}

func (b BooleanNULLEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(BooleanNULL)
	for idx := 0; idx < value.NumField(); idx++ {
		if b == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on BooleanNULL, unknown enumeration value %d", b)
}

func (*BooleanNULLEnum) EncodedLength() int64 {
	return 1
}

func (*BooleanNULLEnum) FalseSinceVersion() uint16 {
	return 0
}

func (b *BooleanNULLEnum) FalseInActingVersion(actingVersion uint16) bool {
	return actingVersion >= b.FalseSinceVersion()
}

func (*BooleanNULLEnum) FalseDeprecated() uint16 {
	return 0
}

func (*BooleanNULLEnum) TrueSinceVersion() uint16 {
	return 0
}

func (b *BooleanNULLEnum) TrueInActingVersion(actingVersion uint16) bool {
	return actingVersion >= b.TrueSinceVersion()
}

func (*BooleanNULLEnum) TrueDeprecated() uint16 {
	return 0
}
