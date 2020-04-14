// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type EstablishmentAck504 struct {
	UUID                        uint64
	RequestTimestamp            uint64
	NextSeqNo                   uint32
	PreviousSeqNo               uint32
	PreviousUUID                uint64
	KeepAliveInterval           uint16
	SecretKeySecureIDExpiration uint16
	FaultToleranceIndicator     FTIEnum
	SplitMsg                    SplitMsgEnum
}

func (e *EstablishmentAck504) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := e.RangeCheck(e.SbeSchemaVersion(), e.SbeSchemaVersion()); err != nil {
			return err
		}
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
	if err := _m.WriteUint32(_w, e.PreviousSeqNo); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, e.PreviousUUID); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, e.KeepAliveInterval); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, e.SecretKeySecureIDExpiration); err != nil {
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

func (e *EstablishmentAck504) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
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
	if !e.PreviousSeqNoInActingVersion(actingVersion) {
		e.PreviousSeqNo = e.PreviousSeqNoNullValue()
	} else {
		if err := _m.ReadUint32(_r, &e.PreviousSeqNo); err != nil {
			return err
		}
	}
	if !e.PreviousUUIDInActingVersion(actingVersion) {
		e.PreviousUUID = e.PreviousUUIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &e.PreviousUUID); err != nil {
			return err
		}
	}
	if !e.KeepAliveIntervalInActingVersion(actingVersion) {
		e.KeepAliveInterval = e.KeepAliveIntervalNullValue()
	} else {
		if err := _m.ReadUint16(_r, &e.KeepAliveInterval); err != nil {
			return err
		}
	}
	if !e.SecretKeySecureIDExpirationInActingVersion(actingVersion) {
		e.SecretKeySecureIDExpiration = e.SecretKeySecureIDExpirationNullValue()
	} else {
		if err := _m.ReadUint16(_r, &e.SecretKeySecureIDExpiration); err != nil {
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

func (e *EstablishmentAck504) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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
	if e.PreviousSeqNoInActingVersion(actingVersion) {
		if e.PreviousSeqNo < e.PreviousSeqNoMinValue() || e.PreviousSeqNo > e.PreviousSeqNoMaxValue() {
			return fmt.Errorf("Range check failed on e.PreviousSeqNo (%v < %v > %v)", e.PreviousSeqNoMinValue(), e.PreviousSeqNo, e.PreviousSeqNoMaxValue())
		}
	}
	if e.PreviousUUIDInActingVersion(actingVersion) {
		if e.PreviousUUID < e.PreviousUUIDMinValue() || e.PreviousUUID > e.PreviousUUIDMaxValue() {
			return fmt.Errorf("Range check failed on e.PreviousUUID (%v < %v > %v)", e.PreviousUUIDMinValue(), e.PreviousUUID, e.PreviousUUIDMaxValue())
		}
	}
	if e.KeepAliveIntervalInActingVersion(actingVersion) {
		if e.KeepAliveInterval < e.KeepAliveIntervalMinValue() || e.KeepAliveInterval > e.KeepAliveIntervalMaxValue() {
			return fmt.Errorf("Range check failed on e.KeepAliveInterval (%v < %v > %v)", e.KeepAliveIntervalMinValue(), e.KeepAliveInterval, e.KeepAliveIntervalMaxValue())
		}
	}
	if e.SecretKeySecureIDExpirationInActingVersion(actingVersion) {
		if e.SecretKeySecureIDExpiration != e.SecretKeySecureIDExpirationNullValue() && (e.SecretKeySecureIDExpiration < e.SecretKeySecureIDExpirationMinValue() || e.SecretKeySecureIDExpiration > e.SecretKeySecureIDExpirationMaxValue()) {
			return fmt.Errorf("Range check failed on e.SecretKeySecureIDExpiration (%v < %v > %v)", e.SecretKeySecureIDExpirationMinValue(), e.SecretKeySecureIDExpiration, e.SecretKeySecureIDExpirationMaxValue())
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

func EstablishmentAck504Init(e *EstablishmentAck504) {
	e.SecretKeySecureIDExpiration = 65535
	return
}

func (*EstablishmentAck504) SbeBlockLength() (blockLength uint16) {
	return 38
}

func (*EstablishmentAck504) SbeTemplateId() (templateId uint16) {
	return 504
}

func (*EstablishmentAck504) SbeSchemaId() (schemaId uint16) {
	return 8
}

func (*EstablishmentAck504) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}

func (*EstablishmentAck504) SbeSemanticType() (semanticType []byte) {
	return []byte("EstablishmentAck")
}

func (*EstablishmentAck504) UUIDId() uint16 {
	return 39001
}

func (*EstablishmentAck504) UUIDSinceVersion() uint16 {
	return 0
}

func (e *EstablishmentAck504) UUIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.UUIDSinceVersion()
}

func (*EstablishmentAck504) UUIDDeprecated() uint16 {
	return 0
}

func (*EstablishmentAck504) UUIDMetaAttribute(meta int) string {
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

func (*EstablishmentAck504) UUIDMinValue() uint64 {
	return 0
}

func (*EstablishmentAck504) UUIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*EstablishmentAck504) UUIDNullValue() uint64 {
	return math.MaxUint64
}

func (*EstablishmentAck504) RequestTimestampId() uint16 {
	return 39002
}

func (*EstablishmentAck504) RequestTimestampSinceVersion() uint16 {
	return 0
}

func (e *EstablishmentAck504) RequestTimestampInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.RequestTimestampSinceVersion()
}

func (*EstablishmentAck504) RequestTimestampDeprecated() uint16 {
	return 0
}

func (*EstablishmentAck504) RequestTimestampMetaAttribute(meta int) string {
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

func (*EstablishmentAck504) RequestTimestampMinValue() uint64 {
	return 0
}

func (*EstablishmentAck504) RequestTimestampMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*EstablishmentAck504) RequestTimestampNullValue() uint64 {
	return math.MaxUint64
}

func (*EstablishmentAck504) NextSeqNoId() uint16 {
	return 39013
}

func (*EstablishmentAck504) NextSeqNoSinceVersion() uint16 {
	return 0
}

func (e *EstablishmentAck504) NextSeqNoInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.NextSeqNoSinceVersion()
}

func (*EstablishmentAck504) NextSeqNoDeprecated() uint16 {
	return 0
}

func (*EstablishmentAck504) NextSeqNoMetaAttribute(meta int) string {
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

func (*EstablishmentAck504) NextSeqNoMinValue() uint32 {
	return 0
}

func (*EstablishmentAck504) NextSeqNoMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*EstablishmentAck504) NextSeqNoNullValue() uint32 {
	return math.MaxUint32
}

func (*EstablishmentAck504) PreviousSeqNoId() uint16 {
	return 39021
}

func (*EstablishmentAck504) PreviousSeqNoSinceVersion() uint16 {
	return 0
}

func (e *EstablishmentAck504) PreviousSeqNoInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.PreviousSeqNoSinceVersion()
}

func (*EstablishmentAck504) PreviousSeqNoDeprecated() uint16 {
	return 0
}

func (*EstablishmentAck504) PreviousSeqNoMetaAttribute(meta int) string {
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

func (*EstablishmentAck504) PreviousSeqNoMinValue() uint32 {
	return 0
}

func (*EstablishmentAck504) PreviousSeqNoMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*EstablishmentAck504) PreviousSeqNoNullValue() uint32 {
	return math.MaxUint32
}

func (*EstablishmentAck504) PreviousUUIDId() uint16 {
	return 39015
}

func (*EstablishmentAck504) PreviousUUIDSinceVersion() uint16 {
	return 0
}

func (e *EstablishmentAck504) PreviousUUIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.PreviousUUIDSinceVersion()
}

func (*EstablishmentAck504) PreviousUUIDDeprecated() uint16 {
	return 0
}

func (*EstablishmentAck504) PreviousUUIDMetaAttribute(meta int) string {
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

func (*EstablishmentAck504) PreviousUUIDMinValue() uint64 {
	return 0
}

func (*EstablishmentAck504) PreviousUUIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*EstablishmentAck504) PreviousUUIDNullValue() uint64 {
	return math.MaxUint64
}

func (*EstablishmentAck504) KeepAliveIntervalId() uint16 {
	return 39014
}

func (*EstablishmentAck504) KeepAliveIntervalSinceVersion() uint16 {
	return 0
}

func (e *EstablishmentAck504) KeepAliveIntervalInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.KeepAliveIntervalSinceVersion()
}

func (*EstablishmentAck504) KeepAliveIntervalDeprecated() uint16 {
	return 0
}

func (*EstablishmentAck504) KeepAliveIntervalMetaAttribute(meta int) string {
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

func (*EstablishmentAck504) KeepAliveIntervalMinValue() uint16 {
	return 0
}

func (*EstablishmentAck504) KeepAliveIntervalMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*EstablishmentAck504) KeepAliveIntervalNullValue() uint16 {
	return math.MaxUint16
}

func (*EstablishmentAck504) SecretKeySecureIDExpirationId() uint16 {
	return 39022
}

func (*EstablishmentAck504) SecretKeySecureIDExpirationSinceVersion() uint16 {
	return 0
}

func (e *EstablishmentAck504) SecretKeySecureIDExpirationInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SecretKeySecureIDExpirationSinceVersion()
}

func (*EstablishmentAck504) SecretKeySecureIDExpirationDeprecated() uint16 {
	return 0
}

func (*EstablishmentAck504) SecretKeySecureIDExpirationMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "int"
	case 4:
		return "optional"
	}
	return ""
}

func (*EstablishmentAck504) SecretKeySecureIDExpirationMinValue() uint16 {
	return 0
}

func (*EstablishmentAck504) SecretKeySecureIDExpirationMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*EstablishmentAck504) SecretKeySecureIDExpirationNullValue() uint16 {
	return 65535
}

func (*EstablishmentAck504) FaultToleranceIndicatorId() uint16 {
	return 39010
}

func (*EstablishmentAck504) FaultToleranceIndicatorSinceVersion() uint16 {
	return 0
}

func (e *EstablishmentAck504) FaultToleranceIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.FaultToleranceIndicatorSinceVersion()
}

func (*EstablishmentAck504) FaultToleranceIndicatorDeprecated() uint16 {
	return 0
}

func (*EstablishmentAck504) FaultToleranceIndicatorMetaAttribute(meta int) string {
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

func (*EstablishmentAck504) SplitMsgId() uint16 {
	return 9553
}

func (*EstablishmentAck504) SplitMsgSinceVersion() uint16 {
	return 0
}

func (e *EstablishmentAck504) SplitMsgInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SplitMsgSinceVersion()
}

func (*EstablishmentAck504) SplitMsgDeprecated() uint16 {
	return 0
}

func (*EstablishmentAck504) SplitMsgMetaAttribute(meta int) string {
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
