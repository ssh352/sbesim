// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"reflect"
)

type OFMOverrideReqEnum uint8
type OFMOverrideReqValues struct {
	Disabled  OFMOverrideReqEnum
	Enabled   OFMOverrideReqEnum
	NullValue OFMOverrideReqEnum
}

var OFMOverrideReq = OFMOverrideReqValues{0, 1, 255}

func (o OFMOverrideReqEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(o)); err != nil {
		return err
	}
	return nil
}

func (o *OFMOverrideReqEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(o)); err != nil {
		return err
	}
	return nil
}

func (o OFMOverrideReqEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(OFMOverrideReq)
	for idx := 0; idx < value.NumField(); idx++ {
		if o == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on OFMOverrideReq, unknown enumeration value %d", o)
}

func (*OFMOverrideReqEnum) EncodedLength() int64 {
	return 1
}

func (*OFMOverrideReqEnum) DisabledSinceVersion() uint16 {
	return 0
}

func (o *OFMOverrideReqEnum) DisabledInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.DisabledSinceVersion()
}

func (*OFMOverrideReqEnum) DisabledDeprecated() uint16 {
	return 0
}

func (*OFMOverrideReqEnum) EnabledSinceVersion() uint16 {
	return 0
}

func (o *OFMOverrideReqEnum) EnabledInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.EnabledSinceVersion()
}

func (*OFMOverrideReqEnum) EnabledDeprecated() uint16 {
	return 0
}
