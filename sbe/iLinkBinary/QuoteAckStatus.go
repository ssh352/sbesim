// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"reflect"
)

type QuoteAckStatusEnum uint8
type QuoteAckStatusValues struct {
	Accepted  QuoteAckStatusEnum
	Rejected  QuoteAckStatusEnum
	NullValue QuoteAckStatusEnum
}

var QuoteAckStatus = QuoteAckStatusValues{0, 5, 255}

func (q QuoteAckStatusEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(q)); err != nil {
		return err
	}
	return nil
}

func (q *QuoteAckStatusEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(q)); err != nil {
		return err
	}
	return nil
}

func (q QuoteAckStatusEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(QuoteAckStatus)
	for idx := 0; idx < value.NumField(); idx++ {
		if q == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on QuoteAckStatus, unknown enumeration value %d", q)
}

func (*QuoteAckStatusEnum) EncodedLength() int64 {
	return 1
}

func (*QuoteAckStatusEnum) AcceptedSinceVersion() uint16 {
	return 0
}

func (q *QuoteAckStatusEnum) AcceptedInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.AcceptedSinceVersion()
}

func (*QuoteAckStatusEnum) AcceptedDeprecated() uint16 {
	return 0
}

func (*QuoteAckStatusEnum) RejectedSinceVersion() uint16 {
	return 0
}

func (q *QuoteAckStatusEnum) RejectedInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.RejectedSinceVersion()
}

func (*QuoteAckStatusEnum) RejectedDeprecated() uint16 {
	return 0
}
