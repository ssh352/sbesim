// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"reflect"
)

type ExpCycleEnum uint8
type ExpCycleValues struct {
	ExpireOnTradingSessionClose ExpCycleEnum
	Expirationatgivendate       ExpCycleEnum
	NullValue                   ExpCycleEnum
}

var ExpCycle = ExpCycleValues{0, 2, 255}

func (e ExpCycleEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(e)); err != nil {
		return err
	}
	return nil
}

func (e *ExpCycleEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(e)); err != nil {
		return err
	}
	return nil
}

func (e ExpCycleEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(ExpCycle)
	for idx := 0; idx < value.NumField(); idx++ {
		if e == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on ExpCycle, unknown enumeration value %d", e)
}

func (*ExpCycleEnum) EncodedLength() int64 {
	return 1
}

func (*ExpCycleEnum) ExpireOnTradingSessionCloseSinceVersion() uint16 {
	return 0
}

func (e *ExpCycleEnum) ExpireOnTradingSessionCloseInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExpireOnTradingSessionCloseSinceVersion()
}

func (*ExpCycleEnum) ExpireOnTradingSessionCloseDeprecated() uint16 {
	return 0
}

func (*ExpCycleEnum) ExpirationatgivendateSinceVersion() uint16 {
	return 0
}

func (e *ExpCycleEnum) ExpirationatgivendateInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExpirationatgivendateSinceVersion()
}

func (*ExpCycleEnum) ExpirationatgivendateDeprecated() uint16 {
	return 0
}
