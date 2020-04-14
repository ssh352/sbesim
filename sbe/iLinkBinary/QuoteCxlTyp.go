// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"reflect"
)

type QuoteCxlTypEnum uint8
type QuoteCxlTypValues struct {
	CancelperInstrument      QuoteCxlTypEnum
	CancelperInstrumentgroup QuoteCxlTypEnum
	Cancelallquotes          QuoteCxlTypEnum
	CancelperQuoteSet        QuoteCxlTypEnum
	NullValue                QuoteCxlTypEnum
}

var QuoteCxlTyp = QuoteCxlTypValues{1, 3, 4, 100, 255}

func (q QuoteCxlTypEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(q)); err != nil {
		return err
	}
	return nil
}

func (q *QuoteCxlTypEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(q)); err != nil {
		return err
	}
	return nil
}

func (q QuoteCxlTypEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(QuoteCxlTyp)
	for idx := 0; idx < value.NumField(); idx++ {
		if q == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on QuoteCxlTyp, unknown enumeration value %d", q)
}

func (*QuoteCxlTypEnum) EncodedLength() int64 {
	return 1
}

func (*QuoteCxlTypEnum) CancelperInstrumentSinceVersion() uint16 {
	return 0
}

func (q *QuoteCxlTypEnum) CancelperInstrumentInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.CancelperInstrumentSinceVersion()
}

func (*QuoteCxlTypEnum) CancelperInstrumentDeprecated() uint16 {
	return 0
}

func (*QuoteCxlTypEnum) CancelperInstrumentgroupSinceVersion() uint16 {
	return 0
}

func (q *QuoteCxlTypEnum) CancelperInstrumentgroupInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.CancelperInstrumentgroupSinceVersion()
}

func (*QuoteCxlTypEnum) CancelperInstrumentgroupDeprecated() uint16 {
	return 0
}

func (*QuoteCxlTypEnum) CancelallquotesSinceVersion() uint16 {
	return 0
}

func (q *QuoteCxlTypEnum) CancelallquotesInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.CancelallquotesSinceVersion()
}

func (*QuoteCxlTypEnum) CancelallquotesDeprecated() uint16 {
	return 0
}

func (*QuoteCxlTypEnum) CancelperQuoteSetSinceVersion() uint16 {
	return 0
}

func (q *QuoteCxlTypEnum) CancelperQuoteSetInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.CancelperQuoteSetSinceVersion()
}

func (*QuoteCxlTypEnum) CancelperQuoteSetDeprecated() uint16 {
	return 0
}
