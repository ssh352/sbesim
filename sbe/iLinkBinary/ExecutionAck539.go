// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type ExecutionAck539 struct {
	PartyDetailsListReqID uint64
	OrderID               uint64
	ExecAckStatus         ExecAckStatusEnum
	SeqNum                uint32
	ClOrdID               [20]byte
	SecExecID             uint64
	LastPx                PRICE9
	SecurityID            int32
	LastQty               uint32
	DKReason              DKReasonEnum
	Side                  SideReqEnum
	SenderID              [20]byte
	SendingTimeEpoch      uint64
	Location              [5]byte
	ManualOrderIndicator  ManualOrdIndReqEnum
}

func (e *ExecutionAck539) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := e.RangeCheck(e.SbeSchemaVersion(), e.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteUint64(_w, e.PartyDetailsListReqID); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, e.OrderID); err != nil {
		return err
	}
	if err := e.ExecAckStatus.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, e.SeqNum); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, e.ClOrdID[:]); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, e.SecExecID); err != nil {
		return err
	}
	if err := e.LastPx.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, e.SecurityID); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, e.LastQty); err != nil {
		return err
	}
	if err := e.DKReason.Encode(_m, _w); err != nil {
		return err
	}
	if err := e.Side.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, e.SenderID[:]); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, e.SendingTimeEpoch); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, e.Location[:]); err != nil {
		return err
	}
	if err := e.ManualOrderIndicator.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (e *ExecutionAck539) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !e.PartyDetailsListReqIDInActingVersion(actingVersion) {
		e.PartyDetailsListReqID = e.PartyDetailsListReqIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &e.PartyDetailsListReqID); err != nil {
			return err
		}
	}
	if !e.OrderIDInActingVersion(actingVersion) {
		e.OrderID = e.OrderIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &e.OrderID); err != nil {
			return err
		}
	}
	if e.ExecAckStatusInActingVersion(actingVersion) {
		if err := e.ExecAckStatus.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !e.SeqNumInActingVersion(actingVersion) {
		e.SeqNum = e.SeqNumNullValue()
	} else {
		if err := _m.ReadUint32(_r, &e.SeqNum); err != nil {
			return err
		}
	}
	if !e.ClOrdIDInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			e.ClOrdID[idx] = e.ClOrdIDNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, e.ClOrdID[:]); err != nil {
			return err
		}
	}
	if !e.SecExecIDInActingVersion(actingVersion) {
		e.SecExecID = e.SecExecIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &e.SecExecID); err != nil {
			return err
		}
	}
	if e.LastPxInActingVersion(actingVersion) {
		if err := e.LastPx.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !e.SecurityIDInActingVersion(actingVersion) {
		e.SecurityID = e.SecurityIDNullValue()
	} else {
		if err := _m.ReadInt32(_r, &e.SecurityID); err != nil {
			return err
		}
	}
	if !e.LastQtyInActingVersion(actingVersion) {
		e.LastQty = e.LastQtyNullValue()
	} else {
		if err := _m.ReadUint32(_r, &e.LastQty); err != nil {
			return err
		}
	}
	if e.DKReasonInActingVersion(actingVersion) {
		if err := e.DKReason.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if e.SideInActingVersion(actingVersion) {
		if err := e.Side.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !e.SenderIDInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			e.SenderID[idx] = e.SenderIDNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, e.SenderID[:]); err != nil {
			return err
		}
	}
	if !e.SendingTimeEpochInActingVersion(actingVersion) {
		e.SendingTimeEpoch = e.SendingTimeEpochNullValue()
	} else {
		if err := _m.ReadUint64(_r, &e.SendingTimeEpoch); err != nil {
			return err
		}
	}
	if !e.LocationInActingVersion(actingVersion) {
		for idx := 0; idx < 5; idx++ {
			e.Location[idx] = e.LocationNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, e.Location[:]); err != nil {
			return err
		}
	}
	if e.ManualOrderIndicatorInActingVersion(actingVersion) {
		if err := e.ManualOrderIndicator.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > e.SbeSchemaVersion() && blockLength > e.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-e.SbeBlockLength()))
	}
	if doRangeCheck {
		if err := e.RangeCheck(actingVersion, e.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (e *ExecutionAck539) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if e.PartyDetailsListReqIDInActingVersion(actingVersion) {
		if e.PartyDetailsListReqID < e.PartyDetailsListReqIDMinValue() || e.PartyDetailsListReqID > e.PartyDetailsListReqIDMaxValue() {
			return fmt.Errorf("Range check failed on e.PartyDetailsListReqID (%v < %v > %v)", e.PartyDetailsListReqIDMinValue(), e.PartyDetailsListReqID, e.PartyDetailsListReqIDMaxValue())
		}
	}
	if e.OrderIDInActingVersion(actingVersion) {
		if e.OrderID < e.OrderIDMinValue() || e.OrderID > e.OrderIDMaxValue() {
			return fmt.Errorf("Range check failed on e.OrderID (%v < %v > %v)", e.OrderIDMinValue(), e.OrderID, e.OrderIDMaxValue())
		}
	}
	if err := e.ExecAckStatus.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if e.SeqNumInActingVersion(actingVersion) {
		if e.SeqNum < e.SeqNumMinValue() || e.SeqNum > e.SeqNumMaxValue() {
			return fmt.Errorf("Range check failed on e.SeqNum (%v < %v > %v)", e.SeqNumMinValue(), e.SeqNum, e.SeqNumMaxValue())
		}
	}
	if e.ClOrdIDInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			if e.ClOrdID[idx] < e.ClOrdIDMinValue() || e.ClOrdID[idx] > e.ClOrdIDMaxValue() {
				return fmt.Errorf("Range check failed on e.ClOrdID[%d] (%v < %v > %v)", idx, e.ClOrdIDMinValue(), e.ClOrdID[idx], e.ClOrdIDMaxValue())
			}
		}
	}
	if e.SecExecIDInActingVersion(actingVersion) {
		if e.SecExecID < e.SecExecIDMinValue() || e.SecExecID > e.SecExecIDMaxValue() {
			return fmt.Errorf("Range check failed on e.SecExecID (%v < %v > %v)", e.SecExecIDMinValue(), e.SecExecID, e.SecExecIDMaxValue())
		}
	}
	if e.SecurityIDInActingVersion(actingVersion) {
		if e.SecurityID < e.SecurityIDMinValue() || e.SecurityID > e.SecurityIDMaxValue() {
			return fmt.Errorf("Range check failed on e.SecurityID (%v < %v > %v)", e.SecurityIDMinValue(), e.SecurityID, e.SecurityIDMaxValue())
		}
	}
	if e.LastQtyInActingVersion(actingVersion) {
		if e.LastQty < e.LastQtyMinValue() || e.LastQty > e.LastQtyMaxValue() {
			return fmt.Errorf("Range check failed on e.LastQty (%v < %v > %v)", e.LastQtyMinValue(), e.LastQty, e.LastQtyMaxValue())
		}
	}
	if err := e.DKReason.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := e.Side.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if e.SenderIDInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			if e.SenderID[idx] < e.SenderIDMinValue() || e.SenderID[idx] > e.SenderIDMaxValue() {
				return fmt.Errorf("Range check failed on e.SenderID[%d] (%v < %v > %v)", idx, e.SenderIDMinValue(), e.SenderID[idx], e.SenderIDMaxValue())
			}
		}
	}
	if e.SendingTimeEpochInActingVersion(actingVersion) {
		if e.SendingTimeEpoch < e.SendingTimeEpochMinValue() || e.SendingTimeEpoch > e.SendingTimeEpochMaxValue() {
			return fmt.Errorf("Range check failed on e.SendingTimeEpoch (%v < %v > %v)", e.SendingTimeEpochMinValue(), e.SendingTimeEpoch, e.SendingTimeEpochMaxValue())
		}
	}
	if e.LocationInActingVersion(actingVersion) {
		for idx := 0; idx < 5; idx++ {
			if e.Location[idx] < e.LocationMinValue() || e.Location[idx] > e.LocationMaxValue() {
				return fmt.Errorf("Range check failed on e.Location[%d] (%v < %v > %v)", idx, e.LocationMinValue(), e.Location[idx], e.LocationMaxValue())
			}
		}
	}
	if err := e.ManualOrderIndicator.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	return nil
}

func ExecutionAck539Init(e *ExecutionAck539) {
	return
}

func (*ExecutionAck539) SbeBlockLength() (blockLength uint16) {
	return 101
}

func (*ExecutionAck539) SbeTemplateId() (templateId uint16) {
	return 539
}

func (*ExecutionAck539) SbeSchemaId() (schemaId uint16) {
	return 8
}

func (*ExecutionAck539) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}

func (*ExecutionAck539) SbeSemanticType() (semanticType []byte) {
	return []byte("BN")
}

func (*ExecutionAck539) PartyDetailsListReqIDId() uint16 {
	return 1505
}

func (*ExecutionAck539) PartyDetailsListReqIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionAck539) PartyDetailsListReqIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.PartyDetailsListReqIDSinceVersion()
}

func (*ExecutionAck539) PartyDetailsListReqIDDeprecated() uint16 {
	return 0
}

func (*ExecutionAck539) PartyDetailsListReqIDMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "int"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionAck539) PartyDetailsListReqIDMinValue() uint64 {
	return 0
}

func (*ExecutionAck539) PartyDetailsListReqIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionAck539) PartyDetailsListReqIDNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionAck539) OrderIDId() uint16 {
	return 37
}

func (*ExecutionAck539) OrderIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionAck539) OrderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderIDSinceVersion()
}

func (*ExecutionAck539) OrderIDDeprecated() uint16 {
	return 0
}

func (*ExecutionAck539) OrderIDMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "int"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionAck539) OrderIDMinValue() uint64 {
	return 0
}

func (*ExecutionAck539) OrderIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionAck539) OrderIDNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionAck539) ExecAckStatusId() uint16 {
	return 1036
}

func (*ExecutionAck539) ExecAckStatusSinceVersion() uint16 {
	return 0
}

func (e *ExecutionAck539) ExecAckStatusInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExecAckStatusSinceVersion()
}

func (*ExecutionAck539) ExecAckStatusDeprecated() uint16 {
	return 0
}

func (*ExecutionAck539) ExecAckStatusMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "int"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionAck539) SeqNumId() uint16 {
	return 9726
}

func (*ExecutionAck539) SeqNumSinceVersion() uint16 {
	return 0
}

func (e *ExecutionAck539) SeqNumInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SeqNumSinceVersion()
}

func (*ExecutionAck539) SeqNumDeprecated() uint16 {
	return 0
}

func (*ExecutionAck539) SeqNumMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "int"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionAck539) SeqNumMinValue() uint32 {
	return 0
}

func (*ExecutionAck539) SeqNumMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionAck539) SeqNumNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionAck539) ClOrdIDId() uint16 {
	return 11
}

func (*ExecutionAck539) ClOrdIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionAck539) ClOrdIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ClOrdIDSinceVersion()
}

func (*ExecutionAck539) ClOrdIDDeprecated() uint16 {
	return 0
}

func (*ExecutionAck539) ClOrdIDMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "String"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionAck539) ClOrdIDMinValue() byte {
	return byte(32)
}

func (*ExecutionAck539) ClOrdIDMaxValue() byte {
	return byte(126)
}

func (*ExecutionAck539) ClOrdIDNullValue() byte {
	return 0
}

func (e *ExecutionAck539) ClOrdIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*ExecutionAck539) SecExecIDId() uint16 {
	return 527
}

func (*ExecutionAck539) SecExecIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionAck539) SecExecIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SecExecIDSinceVersion()
}

func (*ExecutionAck539) SecExecIDDeprecated() uint16 {
	return 0
}

func (*ExecutionAck539) SecExecIDMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "int"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionAck539) SecExecIDMinValue() uint64 {
	return 0
}

func (*ExecutionAck539) SecExecIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionAck539) SecExecIDNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionAck539) LastPxId() uint16 {
	return 31
}

func (*ExecutionAck539) LastPxSinceVersion() uint16 {
	return 0
}

func (e *ExecutionAck539) LastPxInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LastPxSinceVersion()
}

func (*ExecutionAck539) LastPxDeprecated() uint16 {
	return 0
}

func (*ExecutionAck539) LastPxMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "Price"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionAck539) SecurityIDId() uint16 {
	return 48
}

func (*ExecutionAck539) SecurityIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionAck539) SecurityIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SecurityIDSinceVersion()
}

func (*ExecutionAck539) SecurityIDDeprecated() uint16 {
	return 0
}

func (*ExecutionAck539) SecurityIDMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "int"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionAck539) SecurityIDMinValue() int32 {
	return math.MinInt32 + 1
}

func (*ExecutionAck539) SecurityIDMaxValue() int32 {
	return math.MaxInt32
}

func (*ExecutionAck539) SecurityIDNullValue() int32 {
	return math.MinInt32
}

func (*ExecutionAck539) LastQtyId() uint16 {
	return 32
}

func (*ExecutionAck539) LastQtySinceVersion() uint16 {
	return 0
}

func (e *ExecutionAck539) LastQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LastQtySinceVersion()
}

func (*ExecutionAck539) LastQtyDeprecated() uint16 {
	return 0
}

func (*ExecutionAck539) LastQtyMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "int"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionAck539) LastQtyMinValue() uint32 {
	return 0
}

func (*ExecutionAck539) LastQtyMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionAck539) LastQtyNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionAck539) DKReasonId() uint16 {
	return 127
}

func (*ExecutionAck539) DKReasonSinceVersion() uint16 {
	return 0
}

func (e *ExecutionAck539) DKReasonInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.DKReasonSinceVersion()
}

func (*ExecutionAck539) DKReasonDeprecated() uint16 {
	return 0
}

func (*ExecutionAck539) DKReasonMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "char"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionAck539) SideId() uint16 {
	return 54
}

func (*ExecutionAck539) SideSinceVersion() uint16 {
	return 0
}

func (e *ExecutionAck539) SideInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SideSinceVersion()
}

func (*ExecutionAck539) SideDeprecated() uint16 {
	return 0
}

func (*ExecutionAck539) SideMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "int"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionAck539) SenderIDId() uint16 {
	return 5392
}

func (*ExecutionAck539) SenderIDSinceVersion() uint16 {
	return 4
}

func (e *ExecutionAck539) SenderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SenderIDSinceVersion()
}

func (*ExecutionAck539) SenderIDDeprecated() uint16 {
	return 0
}

func (*ExecutionAck539) SenderIDMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "String"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionAck539) SenderIDMinValue() byte {
	return byte(32)
}

func (*ExecutionAck539) SenderIDMaxValue() byte {
	return byte(126)
}

func (*ExecutionAck539) SenderIDNullValue() byte {
	return 0
}

func (e *ExecutionAck539) SenderIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*ExecutionAck539) SendingTimeEpochId() uint16 {
	return 5297
}

func (*ExecutionAck539) SendingTimeEpochSinceVersion() uint16 {
	return 5
}

func (e *ExecutionAck539) SendingTimeEpochInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SendingTimeEpochSinceVersion()
}

func (*ExecutionAck539) SendingTimeEpochDeprecated() uint16 {
	return 0
}

func (*ExecutionAck539) SendingTimeEpochMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "int"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionAck539) SendingTimeEpochMinValue() uint64 {
	return 0
}

func (*ExecutionAck539) SendingTimeEpochMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionAck539) SendingTimeEpochNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionAck539) LocationId() uint16 {
	return 9537
}

func (*ExecutionAck539) LocationSinceVersion() uint16 {
	return 5
}

func (e *ExecutionAck539) LocationInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LocationSinceVersion()
}

func (*ExecutionAck539) LocationDeprecated() uint16 {
	return 0
}

func (*ExecutionAck539) LocationMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "String"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionAck539) LocationMinValue() byte {
	return byte(32)
}

func (*ExecutionAck539) LocationMaxValue() byte {
	return byte(126)
}

func (*ExecutionAck539) LocationNullValue() byte {
	return 0
}

func (e *ExecutionAck539) LocationCharacterEncoding() string {
	return "US-ASCII"
}

func (*ExecutionAck539) ManualOrderIndicatorId() uint16 {
	return 1028
}

func (*ExecutionAck539) ManualOrderIndicatorSinceVersion() uint16 {
	return 5
}

func (e *ExecutionAck539) ManualOrderIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ManualOrderIndicatorSinceVersion()
}

func (*ExecutionAck539) ManualOrderIndicatorDeprecated() uint16 {
	return 0
}

func (*ExecutionAck539) ManualOrderIndicatorMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "int"
	case 4:
		return "required"
	}
	return ""
}
