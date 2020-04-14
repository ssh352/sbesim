// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"reflect"
)

type ListUpdActEnum byte
type ListUpdActValues struct {
	Add       ListUpdActEnum
	Delete    ListUpdActEnum
	NullValue ListUpdActEnum
}

var ListUpdAct = ListUpdActValues{65, 68, 0}

func (l ListUpdActEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, byte(l)); err != nil {
		return err
	}
	return nil
}

func (l *ListUpdActEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*byte)(l)); err != nil {
		return err
	}
	return nil
}

func (l ListUpdActEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(ListUpdAct)
	for idx := 0; idx < value.NumField(); idx++ {
		if l == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on ListUpdAct, unknown enumeration value %d", l)
}

func (*ListUpdActEnum) EncodedLength() int64 {
	return 1
}

func (*ListUpdActEnum) AddSinceVersion() uint16 {
	return 0
}

func (l *ListUpdActEnum) AddInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.AddSinceVersion()
}

func (*ListUpdActEnum) AddDeprecated() uint16 {
	return 0
}

func (*ListUpdActEnum) DeleteSinceVersion() uint16 {
	return 0
}

func (l *ListUpdActEnum) DeleteInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.DeleteSinceVersion()
}

func (*ListUpdActEnum) DeleteDeprecated() uint16 {
	return 0
}
