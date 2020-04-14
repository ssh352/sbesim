// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"reflect"
)

type OrdStatusTrdEnum uint8
type OrdStatusTrdValues struct {
	PartiallyFilled OrdStatusTrdEnum
	Filled          OrdStatusTrdEnum
	NullValue       OrdStatusTrdEnum
}

var OrdStatusTrd = OrdStatusTrdValues{1, 2, 255}

func (o OrdStatusTrdEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(o)); err != nil {
		return err
	}
	return nil
}

func (o *OrdStatusTrdEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(o)); err != nil {
		return err
	}
	return nil
}

func (o OrdStatusTrdEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(OrdStatusTrd)
	for idx := 0; idx < value.NumField(); idx++ {
		if o == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on OrdStatusTrd, unknown enumeration value %d", o)
}

func (*OrdStatusTrdEnum) EncodedLength() int64 {
	return 1
}

func (*OrdStatusTrdEnum) PartiallyFilledSinceVersion() uint16 {
	return 0
}

func (o *OrdStatusTrdEnum) PartiallyFilledInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.PartiallyFilledSinceVersion()
}

func (*OrdStatusTrdEnum) PartiallyFilledDeprecated() uint16 {
	return 0
}

func (*OrdStatusTrdEnum) FilledSinceVersion() uint16 {
	return 0
}

func (o *OrdStatusTrdEnum) FilledInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.FilledSinceVersion()
}

func (*OrdStatusTrdEnum) FilledDeprecated() uint16 {
	return 0
}
