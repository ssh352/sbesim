// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"reflect"
)

type FTIEnum uint8
type FTIValues struct {
	Backup    FTIEnum
	Primary   FTIEnum
	NullValue FTIEnum
}

var FTI = FTIValues{0, 1, 255}

func (f FTIEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(f)); err != nil {
		return err
	}
	return nil
}

func (f *FTIEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(f)); err != nil {
		return err
	}
	return nil
}

func (f FTIEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(FTI)
	for idx := 0; idx < value.NumField(); idx++ {
		if f == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on FTI, unknown enumeration value %d", f)
}

func (*FTIEnum) EncodedLength() int64 {
	return 1
}

func (*FTIEnum) BackupSinceVersion() uint16 {
	return 0
}

func (f *FTIEnum) BackupInActingVersion(actingVersion uint16) bool {
	return actingVersion >= f.BackupSinceVersion()
}

func (*FTIEnum) BackupDeprecated() uint16 {
	return 0
}

func (*FTIEnum) PrimarySinceVersion() uint16 {
	return 0
}

func (f *FTIEnum) PrimaryInActingVersion(actingVersion uint16) bool {
	return actingVersion >= f.PrimarySinceVersion()
}

func (*FTIEnum) PrimaryDeprecated() uint16 {
	return 0
}
