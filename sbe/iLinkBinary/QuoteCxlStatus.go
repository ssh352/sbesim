// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"reflect"
)

type QuoteCxlStatusEnum uint8
type QuoteCxlStatusValues struct {
	CancelperInstrument      QuoteCxlStatusEnum
	CancelperInstrumentgroup QuoteCxlStatusEnum
	Cancelallquotes          QuoteCxlStatusEnum
	Rejected                 QuoteCxlStatusEnum
	CancelperQuoteSet        QuoteCxlStatusEnum
	NullValue                QuoteCxlStatusEnum
}

var QuoteCxlStatus = QuoteCxlStatusValues{1, 3, 4, 5, 100, 255}

func (q QuoteCxlStatusEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(q)); err != nil {
		return err
	}
	return nil
}

func (q *QuoteCxlStatusEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(q)); err != nil {
		return err
	}
	return nil
}

func (q QuoteCxlStatusEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(QuoteCxlStatus)
	for idx := 0; idx < value.NumField(); idx++ {
		if q == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on QuoteCxlStatus, unknown enumeration value %d", q)
}

func (*QuoteCxlStatusEnum) EncodedLength() int64 {
	return 1
}

func (*QuoteCxlStatusEnum) CancelperInstrumentSinceVersion() uint16 {
	return 0
}

func (q *QuoteCxlStatusEnum) CancelperInstrumentInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.CancelperInstrumentSinceVersion()
}

func (*QuoteCxlStatusEnum) CancelperInstrumentDeprecated() uint16 {
	return 0
}

func (*QuoteCxlStatusEnum) CancelperInstrumentgroupSinceVersion() uint16 {
	return 0
}

func (q *QuoteCxlStatusEnum) CancelperInstrumentgroupInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.CancelperInstrumentgroupSinceVersion()
}

func (*QuoteCxlStatusEnum) CancelperInstrumentgroupDeprecated() uint16 {
	return 0
}

func (*QuoteCxlStatusEnum) CancelallquotesSinceVersion() uint16 {
	return 0
}

func (q *QuoteCxlStatusEnum) CancelallquotesInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.CancelallquotesSinceVersion()
}

func (*QuoteCxlStatusEnum) CancelallquotesDeprecated() uint16 {
	return 0
}

func (*QuoteCxlStatusEnum) RejectedSinceVersion() uint16 {
	return 0
}

func (q *QuoteCxlStatusEnum) RejectedInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.RejectedSinceVersion()
}

func (*QuoteCxlStatusEnum) RejectedDeprecated() uint16 {
	return 0
}

func (*QuoteCxlStatusEnum) CancelperQuoteSetSinceVersion() uint16 {
	return 0
}

func (q *QuoteCxlStatusEnum) CancelperQuoteSetInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.CancelperQuoteSetSinceVersion()
}

func (*QuoteCxlStatusEnum) CancelperQuoteSetDeprecated() uint16 {
	return 0
}
