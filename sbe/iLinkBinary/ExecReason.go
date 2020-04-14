// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"reflect"
)

type ExecReasonEnum uint8
type ExecReasonValues struct {
	MarketExchangeOption                                         ExecReasonEnum
	CancelledNotBest                                             ExecReasonEnum
	CancelOnDisconnect                                           ExecReasonEnum
	SelfMatchPreventionOldestOrderCancelled                      ExecReasonEnum
	CancelOnGlobexCreditControlsViolation                        ExecReasonEnum
	CancelFromFirmsoft                                           ExecReasonEnum
	CancelFromRiskManagementAPI                                  ExecReasonEnum
	SelfMatchPreventionNewestOrderCancelled                      ExecReasonEnum
	Cancelduetovolquotedoptionorderrestedqtylessthanminordersize ExecReasonEnum
	NullValue                                                    ExecReasonEnum
}

var ExecReason = ExecReasonValues{8, 9, 100, 103, 104, 105, 106, 107, 108, 255}

func (e ExecReasonEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(e)); err != nil {
		return err
	}
	return nil
}

func (e *ExecReasonEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(e)); err != nil {
		return err
	}
	return nil
}

func (e ExecReasonEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(ExecReason)
	for idx := 0; idx < value.NumField(); idx++ {
		if e == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on ExecReason, unknown enumeration value %d", e)
}

func (*ExecReasonEnum) EncodedLength() int64 {
	return 1
}

func (*ExecReasonEnum) MarketExchangeOptionSinceVersion() uint16 {
	return 0
}

func (e *ExecReasonEnum) MarketExchangeOptionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.MarketExchangeOptionSinceVersion()
}

func (*ExecReasonEnum) MarketExchangeOptionDeprecated() uint16 {
	return 0
}

func (*ExecReasonEnum) CancelledNotBestSinceVersion() uint16 {
	return 0
}

func (e *ExecReasonEnum) CancelledNotBestInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.CancelledNotBestSinceVersion()
}

func (*ExecReasonEnum) CancelledNotBestDeprecated() uint16 {
	return 0
}

func (*ExecReasonEnum) CancelOnDisconnectSinceVersion() uint16 {
	return 0
}

func (e *ExecReasonEnum) CancelOnDisconnectInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.CancelOnDisconnectSinceVersion()
}

func (*ExecReasonEnum) CancelOnDisconnectDeprecated() uint16 {
	return 0
}

func (*ExecReasonEnum) SelfMatchPreventionOldestOrderCancelledSinceVersion() uint16 {
	return 0
}

func (e *ExecReasonEnum) SelfMatchPreventionOldestOrderCancelledInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SelfMatchPreventionOldestOrderCancelledSinceVersion()
}

func (*ExecReasonEnum) SelfMatchPreventionOldestOrderCancelledDeprecated() uint16 {
	return 0
}

func (*ExecReasonEnum) CancelOnGlobexCreditControlsViolationSinceVersion() uint16 {
	return 0
}

func (e *ExecReasonEnum) CancelOnGlobexCreditControlsViolationInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.CancelOnGlobexCreditControlsViolationSinceVersion()
}

func (*ExecReasonEnum) CancelOnGlobexCreditControlsViolationDeprecated() uint16 {
	return 0
}

func (*ExecReasonEnum) CancelFromFirmsoftSinceVersion() uint16 {
	return 0
}

func (e *ExecReasonEnum) CancelFromFirmsoftInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.CancelFromFirmsoftSinceVersion()
}

func (*ExecReasonEnum) CancelFromFirmsoftDeprecated() uint16 {
	return 0
}

func (*ExecReasonEnum) CancelFromRiskManagementAPISinceVersion() uint16 {
	return 0
}

func (e *ExecReasonEnum) CancelFromRiskManagementAPIInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.CancelFromRiskManagementAPISinceVersion()
}

func (*ExecReasonEnum) CancelFromRiskManagementAPIDeprecated() uint16 {
	return 0
}

func (*ExecReasonEnum) SelfMatchPreventionNewestOrderCancelledSinceVersion() uint16 {
	return 0
}

func (e *ExecReasonEnum) SelfMatchPreventionNewestOrderCancelledInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SelfMatchPreventionNewestOrderCancelledSinceVersion()
}

func (*ExecReasonEnum) SelfMatchPreventionNewestOrderCancelledDeprecated() uint16 {
	return 0
}

func (*ExecReasonEnum) CancelduetovolquotedoptionorderrestedqtylessthanminordersizeSinceVersion() uint16 {
	return 0
}

func (e *ExecReasonEnum) CancelduetovolquotedoptionorderrestedqtylessthanminordersizeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.CancelduetovolquotedoptionorderrestedqtylessthanminordersizeSinceVersion()
}

func (*ExecReasonEnum) CancelduetovolquotedoptionorderrestedqtylessthanminordersizeDeprecated() uint16 {
	return 0
}
