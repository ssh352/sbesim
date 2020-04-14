// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"reflect"
)

type TradeAddendumEnum uint8
type TradeAddendumValues struct {
	PartiallyFilled TradeAddendumEnum
	Filled          TradeAddendumEnum
	TradeCancel     TradeAddendumEnum
	TradeCorrection TradeAddendumEnum
	NullValue       TradeAddendumEnum
}

var TradeAddendum = TradeAddendumValues{4, 5, 100, 101, 255}

func (t TradeAddendumEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(t)); err != nil {
		return err
	}
	return nil
}

func (t *TradeAddendumEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(t)); err != nil {
		return err
	}
	return nil
}

func (t TradeAddendumEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(TradeAddendum)
	for idx := 0; idx < value.NumField(); idx++ {
		if t == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on TradeAddendum, unknown enumeration value %d", t)
}

func (*TradeAddendumEnum) EncodedLength() int64 {
	return 1
}

func (*TradeAddendumEnum) PartiallyFilledSinceVersion() uint16 {
	return 0
}

func (t *TradeAddendumEnum) PartiallyFilledInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.PartiallyFilledSinceVersion()
}

func (*TradeAddendumEnum) PartiallyFilledDeprecated() uint16 {
	return 0
}

func (*TradeAddendumEnum) FilledSinceVersion() uint16 {
	return 0
}

func (t *TradeAddendumEnum) FilledInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.FilledSinceVersion()
}

func (*TradeAddendumEnum) FilledDeprecated() uint16 {
	return 0
}

func (*TradeAddendumEnum) TradeCancelSinceVersion() uint16 {
	return 0
}

func (t *TradeAddendumEnum) TradeCancelInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.TradeCancelSinceVersion()
}

func (*TradeAddendumEnum) TradeCancelDeprecated() uint16 {
	return 0
}

func (*TradeAddendumEnum) TradeCorrectionSinceVersion() uint16 {
	return 0
}

func (t *TradeAddendumEnum) TradeCorrectionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.TradeCorrectionSinceVersion()
}

func (*TradeAddendumEnum) TradeCorrectionDeprecated() uint16 {
	return 0
}
