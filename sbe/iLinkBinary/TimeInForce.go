// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"reflect"
)

type TimeInForceEnum uint8
type TimeInForceValues struct {
	Day            TimeInForceEnum
	GoodTillCancel TimeInForceEnum
	FillAndKill    TimeInForceEnum
	FillOrKill     TimeInForceEnum
	GoodTillDate   TimeInForceEnum
	NullValue      TimeInForceEnum
}

var TimeInForce = TimeInForceValues{0, 1, 3, 4, 6, 255}

func (t TimeInForceEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(t)); err != nil {
		return err
	}
	return nil
}

func (t *TimeInForceEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(t)); err != nil {
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

func (*TimeInForceEnum) DaySinceVersion() uint16 {
	return 0
}

func (t *TimeInForceEnum) DayInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.DaySinceVersion()
}

func (*TimeInForceEnum) DayDeprecated() uint16 {
	return 0
}

func (*TimeInForceEnum) GoodTillCancelSinceVersion() uint16 {
	return 0
}

func (t *TimeInForceEnum) GoodTillCancelInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.GoodTillCancelSinceVersion()
}

func (*TimeInForceEnum) GoodTillCancelDeprecated() uint16 {
	return 0
}

func (*TimeInForceEnum) FillAndKillSinceVersion() uint16 {
	return 0
}

func (t *TimeInForceEnum) FillAndKillInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.FillAndKillSinceVersion()
}

func (*TimeInForceEnum) FillAndKillDeprecated() uint16 {
	return 0
}

func (*TimeInForceEnum) FillOrKillSinceVersion() uint16 {
	return 0
}

func (t *TimeInForceEnum) FillOrKillInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.FillOrKillSinceVersion()
}

func (*TimeInForceEnum) FillOrKillDeprecated() uint16 {
	return 0
}

func (*TimeInForceEnum) GoodTillDateSinceVersion() uint16 {
	return 0
}

func (t *TimeInForceEnum) GoodTillDateInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.GoodTillDateSinceVersion()
}

func (*TimeInForceEnum) GoodTillDateDeprecated() uint16 {
	return 0
}
