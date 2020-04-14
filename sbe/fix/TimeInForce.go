// Generated SBE (Simple Binary Encoding) message codec

package fix

import (
	"fmt"
	"io"
	"reflect"
)

type TimeInForceEnum byte
type TimeInForceValues struct {
	DAY              TimeInForceEnum
	GOOD_TILL_CANCEL TimeInForceEnum
	FILL_AND_KILL    TimeInForceEnum
	GOOD_TILL_DATE   TimeInForceEnum
	NullValue        TimeInForceEnum
}

var TimeInForce = TimeInForceValues{48, 49, 51, 54, 0}

func (t TimeInForceEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, byte(t)); err != nil {
		return err
	}
	return nil
}

func (t *TimeInForceEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*byte)(t)); err != nil {
		return err
	}
	return nil
}

func (t TimeInForceEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(TimeInForce)
	for idx := 0; idx < value.NumField(); idx++ {
		if t == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on TimeInForce, unknown enumeration value %d", t)
}

func (*TimeInForceEnum) EncodedLength() int64 {
	return 1
}

func (*TimeInForceEnum) DAYSinceVersion() uint16 {
	return 0
}

func (t *TimeInForceEnum) DAYInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.DAYSinceVersion()
}

func (*TimeInForceEnum) DAYDeprecated() uint16 {
	return 0
}

func (*TimeInForceEnum) GOOD_TILL_CANCELSinceVersion() uint16 {
	return 0
}

func (t *TimeInForceEnum) GOOD_TILL_CANCELInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.GOOD_TILL_CANCELSinceVersion()
}

func (*TimeInForceEnum) GOOD_TILL_CANCELDeprecated() uint16 {
	return 0
}

func (*TimeInForceEnum) FILL_AND_KILLSinceVersion() uint16 {
	return 0
}

func (t *TimeInForceEnum) FILL_AND_KILLInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.FILL_AND_KILLSinceVersion()
}

func (*TimeInForceEnum) FILL_AND_KILLDeprecated() uint16 {
	return 0
}

func (*TimeInForceEnum) GOOD_TILL_DATESinceVersion() uint16 {
	return 0
}

func (t *TimeInForceEnum) GOOD_TILL_DATEInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.GOOD_TILL_DATESinceVersion()
}

func (*TimeInForceEnum) GOOD_TILL_DATEDeprecated() uint16 {
	return 0
}
