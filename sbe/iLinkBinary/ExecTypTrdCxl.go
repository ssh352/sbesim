// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"reflect"
)

type ExecTypTrdCxlEnum byte
type ExecTypTrdCxlValues struct {
	TradeCorrection ExecTypTrdCxlEnum
	TradeCancel     ExecTypTrdCxlEnum
	NullValue       ExecTypTrdCxlEnum
}

var ExecTypTrdCxl = ExecTypTrdCxlValues{71, 72, 0}

func (e ExecTypTrdCxlEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, byte(e)); err != nil {
		return err
	}
	return nil
}

func (e *ExecTypTrdCxlEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*byte)(e)); err != nil {
		return err
	}
	return nil
}

func (e ExecTypTrdCxlEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(ExecTypTrdCxl)
	for idx := 0; idx < value.NumField(); idx++ {
		if e == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on ExecTypTrdCxl, unknown enumeration value %d", e)
}

func (*ExecTypTrdCxlEnum) EncodedLength() int64 {
	return 1
}

func (*ExecTypTrdCxlEnum) TradeCorrectionSinceVersion() uint16 {
	return 0
}

func (e *ExecTypTrdCxlEnum) TradeCorrectionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TradeCorrectionSinceVersion()
}

func (*ExecTypTrdCxlEnum) TradeCorrectionDeprecated() uint16 {
	return 0
}

func (*ExecTypTrdCxlEnum) TradeCancelSinceVersion() uint16 {
	return 0
}

func (e *ExecTypTrdCxlEnum) TradeCancelInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TradeCancelSinceVersion()
}

func (*ExecTypTrdCxlEnum) TradeCancelDeprecated() uint16 {
	return 0
}
