// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type NegotiationReject502 struct {
	Reason                  [48]byte
	UUID                    uint64
	RequestTimestamp        uint64
	ErrorCodes              uint16
	FaultToleranceIndicator FTIEnum
	SplitMsg                SplitMsgEnum
}

func (n *NegotiationReject502) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := n.RangeCheck(n.SbeSchemaVersion(), n.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteBytes(_w, n.Reason[:]); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, n.UUID); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, n.RequestTimestamp); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, n.ErrorCodes); err != nil {
		return err
	}
	if err := n.FaultToleranceIndicator.Encode(_m, _w); err != nil {
		return err
	}
	if err := n.SplitMsg.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (n *NegotiationReject502) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !n.ReasonInActingVersion(actingVersion) {
		for idx := 0; idx < 48; idx++ {
			n.Reason[idx] = n.ReasonNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, n.Reason[:]); err != nil {
			return err
		}
	}
	if !n.UUIDInActingVersion(actingVersion) {
		n.UUID = n.UUIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &n.UUID); err != nil {
			return err
		}
	}
	if !n.RequestTimestampInActingVersion(actingVersion) {
		n.RequestTimestamp = n.RequestTimestampNullValue()
	} else {
		if err := _m.ReadUint64(_r, &n.RequestTimestamp); err != nil {
			return err
		}
	}
	if !n.ErrorCodesInActingVersion(actingVersion) {
		n.ErrorCodes = n.ErrorCodesNullValue()
	} else {
		if err := _m.ReadUint16(_r, &n.ErrorCodes); err != nil {
			return err
		}
	}
	if n.FaultToleranceIndicatorInActingVersion(actingVersion) {
		if err := n.FaultToleranceIndicator.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if n.SplitMsgInActingVersion(actingVersion) {
		if err := n.SplitMsg.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > n.SbeSchemaVersion() && blockLength > n.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-n.SbeBlockLength()))
	}
	if doRangeCheck {
		if err := n.RangeCheck(actingVersion, n.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (n *NegotiationReject502) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if n.ReasonInActingVersion(actingVersion) {
		for idx := 0; idx < 48; idx++ {
			if n.Reason[idx] < n.ReasonMinValue() || n.Reason[idx] > n.ReasonMaxValue() {
				return fmt.Errorf("Range check failed on n.Reason[%d] (%v < %v > %v)", idx, n.ReasonMinValue(), n.Reason[idx], n.ReasonMaxValue())
			}
		}
	}
	if n.UUIDInActingVersion(actingVersion) {
		if n.UUID < n.UUIDMinValue() || n.UUID > n.UUIDMaxValue() {
			return fmt.Errorf("Range check failed on n.UUID (%v < %v > %v)", n.UUIDMinValue(), n.UUID, n.UUIDMaxValue())
		}
	}
	if n.RequestTimestampInActingVersion(actingVersion) {
		if n.RequestTimestamp < n.RequestTimestampMinValue() || n.RequestTimestamp > n.RequestTimestampMaxValue() {
			return fmt.Errorf("Range check failed on n.RequestTimestamp (%v < %v > %v)", n.RequestTimestampMinValue(), n.RequestTimestamp, n.RequestTimestampMaxValue())
		}
	}
	if n.ErrorCodesInActingVersion(actingVersion) {
		if n.ErrorCodes < n.ErrorCodesMinValue() || n.ErrorCodes > n.ErrorCodesMaxValue() {
			return fmt.Errorf("Range check failed on n.ErrorCodes (%v < %v > %v)", n.ErrorCodesMinValue(), n.ErrorCodes, n.ErrorCodesMaxValue())
		}
	}
	if err := n.FaultToleranceIndicator.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := n.SplitMsg.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	return nil
}

func NegotiationReject502Init(n *NegotiationReject502) {
	for idx := 0; idx < 48; idx++ {
		n.Reason[idx] = 0
	}
	return
}

func (*NegotiationReject502) SbeBlockLength() (blockLength uint16) {
	return 68
}

func (*NegotiationReject502) SbeTemplateId() (templateId uint16) {
	return 502
}

func (*NegotiationReject502) SbeSchemaId() (schemaId uint16) {
	return 8
}

func (*NegotiationReject502) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}

func (*NegotiationReject502) SbeSemanticType() (semanticType []byte) {
	return []byte("NegotiationReject")
}

func (*NegotiationReject502) ReasonId() uint16 {
	return 39011
}

func (*NegotiationReject502) ReasonSinceVersion() uint16 {
	return 0
}

func (n *NegotiationReject502) ReasonInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.ReasonSinceVersion()
}

func (*NegotiationReject502) ReasonDeprecated() uint16 {
	return 0
}

func (*NegotiationReject502) ReasonMetaAttribute(meta int) string {
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

func (*NegotiationReject502) ReasonMinValue() byte {
	return byte(32)
}

func (*NegotiationReject502) ReasonMaxValue() byte {
	return byte(126)
}

func (*NegotiationReject502) ReasonNullValue() byte {
	return 0
}

func (n *NegotiationReject502) ReasonCharacterEncoding() string {
	return "US-ASCII"
}

func (*NegotiationReject502) UUIDId() uint16 {
	return 39001
}

func (*NegotiationReject502) UUIDSinceVersion() uint16 {
	return 0
}

func (n *NegotiationReject502) UUIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.UUIDSinceVersion()
}

func (*NegotiationReject502) UUIDDeprecated() uint16 {
	return 0
}

func (*NegotiationReject502) UUIDMetaAttribute(meta int) string {
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

func (*NegotiationReject502) UUIDMinValue() uint64 {
	return 0
}

func (*NegotiationReject502) UUIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*NegotiationReject502) UUIDNullValue() uint64 {
	return math.MaxUint64
}

func (*NegotiationReject502) RequestTimestampId() uint16 {
	return 39002
}

func (*NegotiationReject502) RequestTimestampSinceVersion() uint16 {
	return 0
}

func (n *NegotiationReject502) RequestTimestampInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.RequestTimestampSinceVersion()
}

func (*NegotiationReject502) RequestTimestampDeprecated() uint16 {
	return 0
}

func (*NegotiationReject502) RequestTimestampMetaAttribute(meta int) string {
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

func (*NegotiationReject502) RequestTimestampMinValue() uint64 {
	return 0
}

func (*NegotiationReject502) RequestTimestampMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*NegotiationReject502) RequestTimestampNullValue() uint64 {
	return math.MaxUint64
}

func (*NegotiationReject502) ErrorCodesId() uint16 {
	return 39012
}

func (*NegotiationReject502) ErrorCodesSinceVersion() uint16 {
	return 0
}

func (n *NegotiationReject502) ErrorCodesInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.ErrorCodesSinceVersion()
}

func (*NegotiationReject502) ErrorCodesDeprecated() uint16 {
	return 0
}

func (*NegotiationReject502) ErrorCodesMetaAttribute(meta int) string {
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

func (*NegotiationReject502) ErrorCodesMinValue() uint16 {
	return 0
}

func (*NegotiationReject502) ErrorCodesMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*NegotiationReject502) ErrorCodesNullValue() uint16 {
	return math.MaxUint16
}

func (*NegotiationReject502) FaultToleranceIndicatorId() uint16 {
	return 39010
}

func (*NegotiationReject502) FaultToleranceIndicatorSinceVersion() uint16 {
	return 0
}

func (n *NegotiationReject502) FaultToleranceIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.FaultToleranceIndicatorSinceVersion()
}

func (*NegotiationReject502) FaultToleranceIndicatorDeprecated() uint16 {
	return 0
}

func (*NegotiationReject502) FaultToleranceIndicatorMetaAttribute(meta int) string {
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

func (*NegotiationReject502) SplitMsgId() uint16 {
	return 9553
}

func (*NegotiationReject502) SplitMsgSinceVersion() uint16 {
	return 0
}

func (n *NegotiationReject502) SplitMsgInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.SplitMsgSinceVersion()
}

func (*NegotiationReject502) SplitMsgDeprecated() uint16 {
	return 0
}

func (*NegotiationReject502) SplitMsgMetaAttribute(meta int) string {
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
