// Generated SBE (Simple Binary Encoding) message codec

package fix

import (
	"fmt"
	"io"
	"reflect"
)

type HandInstEnum byte
type HandInstValues struct {
	AUTOMATED_EXECUTION HandInstEnum
	NullValue           HandInstEnum
}

var HandInst = HandInstValues{49, 0}

func (h HandInstEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, byte(h)); err != nil {
		return err
	}
	return nil
}

func (h *HandInstEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*byte)(h)); err != nil {
		return err
	}
	return nil
}

func (h HandInstEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(HandInst)
	for idx := 0; idx < value.NumField(); idx++ {
		if h == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on HandInst, unknown enumeration value %d", h)
}

func (*HandInstEnum) EncodedLength() int64 {
	return 1
}

func (*HandInstEnum) AUTOMATED_EXECUTIONSinceVersion() uint16 {
	return 0
}

func (h *HandInstEnum) AUTOMATED_EXECUTIONInActingVersion(actingVersion uint16) bool {
	return actingVersion >= h.AUTOMATED_EXECUTIONSinceVersion()
}

func (*HandInstEnum) AUTOMATED_EXECUTIONDeprecated() uint16 {
	return 0
}
