// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"reflect"
)

type OrdStatusTrdCxlEnum byte
type OrdStatusTrdCxlValues struct {
	TradeCorrection OrdStatusTrdCxlEnum
	TradeCancel     OrdStatusTrdCxlEnum
	NullValue       OrdStatusTrdCxlEnum
}

var OrdStatusTrdCxl = OrdStatusTrdCxlValues{71, 72, 0}

func (o OrdStatusTrdCxlEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, byte(o)); err != nil {
		return err
	}
	return nil
}

func (o *OrdStatusTrdCxlEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*byte)(o)); err != nil {
		return err
	}
	return nil
}

func (o OrdStatusTrdCxlEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(OrdStatusTrdCxl)
	for idx := 0; idx < value.NumField(); idx++ {
		if o == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on OrdStatusTrdCxl, unknown enumeration value %d", o)
}

func (*OrdStatusTrdCxlEnum) EncodedLength() int64 {
	return 1
}

func (*OrdStatusTrdCxlEnum) TradeCorrectionSinceVersion() uint16 {
	return 0
}

func (o *OrdStatusTrdCxlEnum) TradeCorrectionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.TradeCorrectionSinceVersion()
}

func (*OrdStatusTrdCxlEnum) TradeCorrectionDeprecated() uint16 {
	return 0
}

func (*OrdStatusTrdCxlEnum) TradeCancelSinceVersion() uint16 {
	return 0
}

func (o *OrdStatusTrdCxlEnum) TradeCancelInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.TradeCancelSinceVersion()
}

func (*OrdStatusTrdCxlEnum) TradeCancelDeprecated() uint16 {
	return 0
}
