// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"reflect"
)

type QuoteTypEnum uint8
type QuoteTypValues struct {
	Tradeable QuoteTypEnum
	NullValue QuoteTypEnum
}

var QuoteTyp = QuoteTypValues{1, 255}

func (q QuoteTypEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(q)); err != nil {
		return err
	}
	return nil
}

func (q *QuoteTypEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(q)); err != nil {
		return err
	}
	return nil
}

func (q QuoteTypEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(QuoteTyp)
	for idx := 0; idx < value.NumField(); idx++ {
		if q == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on QuoteTyp, unknown enumeration value %d", q)
}

func (*QuoteTypEnum) EncodedLength() int64 {
	return 1
}

func (*QuoteTypEnum) TradeableSinceVersion() uint16 {
	return 0
}

func (q *QuoteTypEnum) TradeableInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.TradeableSinceVersion()
}

func (*QuoteTypEnum) TradeableDeprecated() uint16 {
	return 0
}
