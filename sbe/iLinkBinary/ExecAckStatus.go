// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"reflect"
)

type ExecAckStatusEnum uint8
type ExecAckStatusValues struct {
	Accepted  ExecAckStatusEnum
	Rejected  ExecAckStatusEnum
	NullValue ExecAckStatusEnum
}

var ExecAckStatus = ExecAckStatusValues{1, 2, 255}

func (e ExecAckStatusEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(e)); err != nil {
		return err
	}
	return nil
}

func (e *ExecAckStatusEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(e)); err != nil {
		return err
	}
	return nil
}

func (e ExecAckStatusEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(ExecAckStatus)
	for idx := 0; idx < value.NumField(); idx++ {
		if e == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on ExecAckStatus, unknown enumeration value %d", e)
}

func (*ExecAckStatusEnum) EncodedLength() int64 {
	return 1
}

func (*ExecAckStatusEnum) AcceptedSinceVersion() uint16 {
	return 0
}

func (e *ExecAckStatusEnum) AcceptedInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.AcceptedSinceVersion()
}

func (*ExecAckStatusEnum) AcceptedDeprecated() uint16 {
	return 0
}

func (*ExecAckStatusEnum) RejectedSinceVersion() uint16 {
	return 0
}

func (e *ExecAckStatusEnum) RejectedInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.RejectedSinceVersion()
}

func (*ExecAckStatusEnum) RejectedDeprecated() uint16 {
	return 0
}
