// Generated SBE (Simple Binary Encoding) message codec

package fix

import (
	"fmt"
	"io"
	"reflect"
)

type BooleanTypeEnum uint8
type BooleanTypeValues struct {
	FIX_FALSE BooleanTypeEnum
	FIX_TRUE  BooleanTypeEnum
	NullValue BooleanTypeEnum
}

var BooleanType = BooleanTypeValues{0, 1, 255}

func (b BooleanTypeEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(b)); err != nil {
		return err
	}
	return nil
}

func (b *BooleanTypeEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(b)); err != nil {
		return err
	}
	return nil
}

func (b BooleanTypeEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(BooleanType)
	for idx := 0; idx < value.NumField(); idx++ {
		if b == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on BooleanType, unknown enumeration value %d", b)
}

func (*BooleanTypeEnum) EncodedLength() int64 {
	return 1
}

func (*BooleanTypeEnum) FIX_FALSESinceVersion() uint16 {
	return 0
}

func (b *BooleanTypeEnum) FIX_FALSEInActingVersion(actingVersion uint16) bool {
	return actingVersion >= b.FIX_FALSESinceVersion()
}

func (*BooleanTypeEnum) FIX_FALSEDeprecated() uint16 {
	return 0
}

func (*BooleanTypeEnum) FIX_TRUESinceVersion() uint16 {
	return 0
}

func (b *BooleanTypeEnum) FIX_TRUEInActingVersion(actingVersion uint16) bool {
	return actingVersion >= b.FIX_TRUESinceVersion()
}

func (*BooleanTypeEnum) FIX_TRUEDeprecated() uint16 {
	return 0
}
