// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type EstablishmentReject505 struct {
	Reason                  [48]byte
	UUID                    uint64
	RequestTimestamp        uint64
	NextSeqNo               uint32
	ErrorCodes              uint16
	FaultToleranceIndicator FTIEnum
	SplitMsg                SplitMsgEnum
}

func (e *EstablishmentReject505) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := e.RangeCheck(e.SbeSchemaVersion(), e.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteBytes(_w, e.Reason[:]); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, e.UUID); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, e.RequestTimestamp); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, e.NextSeqNo); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, e.ErrorCodes); err != nil {
		return err
	}
	if err := e.FaultToleranceIndicator.Encode(_m, _w); err != nil {
		return err
	}
	if err := e.SplitMsg.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (e *EstablishmentReject505) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !e.ReasonInActingVersion(actingVersion) {
		for idx := 0; idx < 48; idx++ {
			e.Reason[idx] = e.ReasonNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, e.Reason[:]); err != nil {
			return err
		}
	}
	if !e.UUIDInActingVersion(actingVersion) {
		e.UUID = e.UUIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &e.UUID); err != nil {
			return err
		}
	}
	if !e.RequestTimestampInActingVersion(actingVersion) {
		e.RequestTimestamp = e.RequestTimestampNullValue()
	} else {
		if err := _m.ReadUint64(_r, &e.RequestTimestamp); err != nil {
			return err
		}
	}
	if !e.NextSeqNoInActingVersion(actingVersion) {
		e.NextSeqNo = e.NextSeqNoNullValue()
	} else {
		if err := _m.ReadUint32(_r, &e.NextSeqNo); err != nil {
			return err
		}
	}
	if !e.ErrorCodesInActingVersion(actingVersion) {
		e.ErrorCodes = e.ErrorCodesNullValue()
	} else {
		if err := _m.ReadUint16(_r, &e.ErrorCodes); err != nil {
			return err
		}
	}
	if e.FaultToleranceIndicatorInActingVersion(actingVersion) {
		if err := e.FaultToleranceIndicator.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if e.SplitMsgInActingVersion(actingVersion) {
		if err := e.SplitMsg.Decode(_m, _r, actingVersion); err != nil {
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

func (e *EstablishmentReject505) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if e.ReasonInActingVersion(actingVersion) {
		for idx := 0; idx < 48; idx++ {
			if e.Reason[idx] < e.ReasonMinValue() || e.Reason[idx] > e.ReasonMaxValue() {
				return fmt.Errorf("Range check failed on e.Reason[%d] (%v < %v > %v)", idx, e.ReasonMinValue(), e.Reason[idx], e.ReasonMaxValue())
			}
		}
	}
	if e.UUIDInActingVersion(actingVersion) {
		if e.UUID < e.UUIDMinValue() || e.UUID > e.UUIDMaxValue() {
			return fmt.Errorf("Range check failed on e.UUID (%v < %v > %v)", e.UUIDMinValue(), e.UUID, e.UUIDMaxValue())
		}
	}
	if e.RequestTimestampInActingVersion(actingVersion) {
		if e.RequestTimestamp < e.RequestTimestampMinValue() || e.RequestTimestamp > e.RequestTimestampMaxValue() {
			return fmt.Errorf("Range check failed on e.RequestTimestamp (%v < %v > %v)", e.RequestTimestampMinValue(), e.RequestTimestamp, e.RequestTimestampMaxValue())
		}
	}
	if e.NextSeqNoInActingVersion(actingVersion) {
		if e.NextSeqNo < e.NextSeqNoMinValue() || e.NextSeqNo > e.NextSeqNoMaxValue() {
			return fmt.Errorf("Range check failed on e.NextSeqNo (%v < %v > %v)", e.NextSeqNoMinValue(), e.NextSeqNo, e.NextSeqNoMaxValue())
		}
	}
	if e.ErrorCodesInActingVersion(actingVersion) {
		if e.ErrorCodes < e.ErrorCodesMinValue() || e.ErrorCodes > e.ErrorCodesMaxValue() {
			return fmt.Errorf("Range check failed on e.ErrorCodes (%v < %v > %v)", e.ErrorCodesMinValue(), e.ErrorCodes, e.ErrorCodesMaxValue())
		}
	}
	if err := e.FaultToleranceIndicator.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := e.SplitMsg.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	return nil
}

func EstablishmentReject505Init(e *EstablishmentReject505) {
	for idx := 0; idx < 48; idx++ {
		e.Reason[idx] = 0
	}
	return
}

func (*EstablishmentReject505) SbeBlockLength() (blockLength uint16) {
	return 72
}

func (*EstablishmentReject505) SbeTemplateId() (templateId uint16) {
	return 505
}

func (*EstablishmentReject505) SbeSchemaId() (schemaId uint16) {
	return 8
}

func (*EstablishmentReject505) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}

func (*EstablishmentReject505) SbeSemanticType() (semanticType []byte) {
	return []byte("EstablishmentReject")
}

func (*EstablishmentReject505) ReasonId() uint16 {
	return 39011
}

func (*EstablishmentReject505) ReasonSinceVersion() uint16 {
	return 0
}

func (e *EstablishmentReject505) ReasonInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ReasonSinceVersion()
}

func (*EstablishmentReject505) ReasonDeprecated() uint16 {
	return 0
}

func (*EstablishmentReject505) ReasonMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "String"
	case 4:
		return "optional"
	}
	return ""
}

func (*EstablishmentReject505) ReasonMinValue() byte {
	return byte(32)
}

func (*EstablishmentReject505) ReasonMaxValue() byte {
	return byte(126)
}

func (*EstablishmentReject505) ReasonNullValue() byte {
	return 0
}

func (e *EstablishmentReject505) ReasonCharacterEncoding() string {
	return "US-ASCII"
}

func (*EstablishmentReject505) UUIDId() uint16 {
	return 39001
}

func (*EstablishmentReject505) UUIDSinceVersion() uint16 {
	return 0
}

func (e *EstablishmentReject505) UUIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.UUIDSinceVersion()
}

func (*EstablishmentReject505) UUIDDeprecated() uint16 {
	return 0
}

func (*EstablishmentReject505) UUIDMetaAttribute(meta int) string {
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

func (*EstablishmentReject505) UUIDMinValue() uint64 {
	return 0
}

func (*EstablishmentReject505) UUIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*EstablishmentReject505) UUIDNullValue() uint64 {
	return math.MaxUint64
}

func (*EstablishmentReject505) RequestTimestampId() uint16 {
	return 39002
}

func (*EstablishmentReject505) RequestTimestampSinceVersion() uint16 {
	return 0
}

func (e *EstablishmentReject505) RequestTimestampInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.RequestTimestampSinceVersion()
}

func (*EstablishmentReject505) RequestTimestampDeprecated() uint16 {
	return 0
}

func (*EstablishmentReject505) RequestTimestampMetaAttribute(meta int) string {
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

func (*EstablishmentReject505) RequestTimestampMinValue() uint64 {
	return 0
}

func (*EstablishmentReject505) RequestTimestampMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*EstablishmentReject505) RequestTimestampNullValue() uint64 {
	return math.MaxUint64
}

func (*EstablishmentReject505) NextSeqNoId() uint16 {
	return 39013
}

func (*EstablishmentReject505) NextSeqNoSinceVersion() uint16 {
	return 0
}

func (e *EstablishmentReject505) NextSeqNoInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.NextSeqNoSinceVersion()
}

func (*EstablishmentReject505) NextSeqNoDeprecated() uint16 {
	return 0
}

func (*EstablishmentReject505) NextSeqNoMetaAttribute(meta int) string {
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

func (*EstablishmentReject505) NextSeqNoMinValue() uint32 {
	return 0
}

func (*EstablishmentReject505) NextSeqNoMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*EstablishmentReject505) NextSeqNoNullValue() uint32 {
	return math.MaxUint32
}

func (*EstablishmentReject505) ErrorCodesId() uint16 {
	return 39012
}

func (*EstablishmentReject505) ErrorCodesSinceVersion() uint16 {
	return 0
}

func (e *EstablishmentReject505) ErrorCodesInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ErrorCodesSinceVersion()
}

func (*EstablishmentReject505) ErrorCodesDeprecated() uint16 {
	return 0
}

func (*EstablishmentReject505) ErrorCodesMetaAttribute(meta int) string {
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

func (*EstablishmentReject505) ErrorCodesMinValue() uint16 {
	return 0
}

func (*EstablishmentReject505) ErrorCodesMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*EstablishmentReject505) ErrorCodesNullValue() uint16 {
	return math.MaxUint16
}

func (*EstablishmentReject505) FaultToleranceIndicatorId() uint16 {
	return 39010
}

func (*EstablishmentReject505) FaultToleranceIndicatorSinceVersion() uint16 {
	return 0
}

func (e *EstablishmentReject505) FaultToleranceIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.FaultToleranceIndicatorSinceVersion()
}

func (*EstablishmentReject505) FaultToleranceIndicatorDeprecated() uint16 {
	return 0
}

func (*EstablishmentReject505) FaultToleranceIndicatorMetaAttribute(meta int) string {
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

func (*EstablishmentReject505) SplitMsgId() uint16 {
	return 9553
}

func (*EstablishmentReject505) SplitMsgSinceVersion() uint16 {
	return 0
}

func (e *EstablishmentReject505) SplitMsgInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SplitMsgSinceVersion()
}

func (*EstablishmentReject505) SplitMsgDeprecated() uint16 {
	return 0
}

func (*EstablishmentReject505) SplitMsgMetaAttribute(meta int) string {
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
