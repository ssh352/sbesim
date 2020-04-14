// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type NegotiationResponse501 struct {
	ServerFlow                  [11]byte
	UUID                        uint64
	RequestTimestamp            uint64
	SecretKeySecureIDExpiration uint16
	FaultToleranceIndicator     FTIEnum
	SplitMsg                    SplitMsgEnum
	PreviousSeqNo               uint32
	PreviousUUID                uint64
	Credentials                 []byte
}

func (n *NegotiationResponse501) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := n.RangeCheck(n.SbeSchemaVersion(), n.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteUint64(_w, n.UUID); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, n.RequestTimestamp); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, n.SecretKeySecureIDExpiration); err != nil {
		return err
	}
	if err := n.FaultToleranceIndicator.Encode(_m, _w); err != nil {
		return err
	}
	if err := n.SplitMsg.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, n.PreviousSeqNo); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, n.PreviousUUID); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, uint16(len(n.Credentials))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, n.Credentials); err != nil {
		return err
	}
	return nil
}

func (n *NegotiationResponse501) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	copy(n.ServerFlow[:], "RECOVERABLE")
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
	if !n.SecretKeySecureIDExpirationInActingVersion(actingVersion) {
		n.SecretKeySecureIDExpiration = n.SecretKeySecureIDExpirationNullValue()
	} else {
		if err := _m.ReadUint16(_r, &n.SecretKeySecureIDExpiration); err != nil {
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
	if !n.PreviousSeqNoInActingVersion(actingVersion) {
		n.PreviousSeqNo = n.PreviousSeqNoNullValue()
	} else {
		if err := _m.ReadUint32(_r, &n.PreviousSeqNo); err != nil {
			return err
		}
	}
	if !n.PreviousUUIDInActingVersion(actingVersion) {
		n.PreviousUUID = n.PreviousUUIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &n.PreviousUUID); err != nil {
			return err
		}
	}
	if actingVersion > n.SbeSchemaVersion() && blockLength > n.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-n.SbeBlockLength()))
	}

	if n.CredentialsInActingVersion(actingVersion) {
		var CredentialsLength uint16
		if err := _m.ReadUint16(_r, &CredentialsLength); err != nil {
			return err
		}
		if cap(n.Credentials) < int(CredentialsLength) {
			n.Credentials = make([]byte, CredentialsLength)
		}
		if err := _m.ReadBytes(_r, n.Credentials); err != nil {
			return err
		}
	}
	if doRangeCheck {
		if err := n.RangeCheck(actingVersion, n.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (n *NegotiationResponse501) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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
	if n.SecretKeySecureIDExpirationInActingVersion(actingVersion) {
		if n.SecretKeySecureIDExpiration != n.SecretKeySecureIDExpirationNullValue() && (n.SecretKeySecureIDExpiration < n.SecretKeySecureIDExpirationMinValue() || n.SecretKeySecureIDExpiration > n.SecretKeySecureIDExpirationMaxValue()) {
			return fmt.Errorf("Range check failed on n.SecretKeySecureIDExpiration (%v < %v > %v)", n.SecretKeySecureIDExpirationMinValue(), n.SecretKeySecureIDExpiration, n.SecretKeySecureIDExpirationMaxValue())
		}
	}
	if err := n.FaultToleranceIndicator.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := n.SplitMsg.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if n.PreviousSeqNoInActingVersion(actingVersion) {
		if n.PreviousSeqNo < n.PreviousSeqNoMinValue() || n.PreviousSeqNo > n.PreviousSeqNoMaxValue() {
			return fmt.Errorf("Range check failed on n.PreviousSeqNo (%v < %v > %v)", n.PreviousSeqNoMinValue(), n.PreviousSeqNo, n.PreviousSeqNoMaxValue())
		}
	}
	if n.PreviousUUIDInActingVersion(actingVersion) {
		if n.PreviousUUID < n.PreviousUUIDMinValue() || n.PreviousUUID > n.PreviousUUIDMaxValue() {
			return fmt.Errorf("Range check failed on n.PreviousUUID (%v < %v > %v)", n.PreviousUUIDMinValue(), n.PreviousUUID, n.PreviousUUIDMaxValue())
		}
	}
	return nil
}

func NegotiationResponse501Init(n *NegotiationResponse501) {
	copy(n.ServerFlow[:], "RECOVERABLE")
	n.SecretKeySecureIDExpiration = 65535
	return
}

func (*NegotiationResponse501) SbeBlockLength() (blockLength uint16) {
	return 32
}

func (*NegotiationResponse501) SbeTemplateId() (templateId uint16) {
	return 501
}

func (*NegotiationResponse501) SbeSchemaId() (schemaId uint16) {
	return 8
}

func (*NegotiationResponse501) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}

func (*NegotiationResponse501) SbeSemanticType() (semanticType []byte) {
	return []byte("NegotiationResponse")
}

func (*NegotiationResponse501) ServerFlowId() uint16 {
	return 39009
}

func (*NegotiationResponse501) ServerFlowSinceVersion() uint16 {
	return 0
}

func (n *NegotiationResponse501) ServerFlowInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.ServerFlowSinceVersion()
}

func (*NegotiationResponse501) ServerFlowDeprecated() uint16 {
	return 0
}

func (*NegotiationResponse501) ServerFlowMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "String"
	case 4:
		return "constant"
	}
	return ""
}

func (*NegotiationResponse501) ServerFlowMinValue() byte {
	return byte(32)
}

func (*NegotiationResponse501) ServerFlowMaxValue() byte {
	return byte(126)
}

func (*NegotiationResponse501) ServerFlowNullValue() byte {
	return 0
}

func (*NegotiationResponse501) UUIDId() uint16 {
	return 39001
}

func (*NegotiationResponse501) UUIDSinceVersion() uint16 {
	return 0
}

func (n *NegotiationResponse501) UUIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.UUIDSinceVersion()
}

func (*NegotiationResponse501) UUIDDeprecated() uint16 {
	return 0
}

func (*NegotiationResponse501) UUIDMetaAttribute(meta int) string {
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

func (*NegotiationResponse501) UUIDMinValue() uint64 {
	return 0
}

func (*NegotiationResponse501) UUIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*NegotiationResponse501) UUIDNullValue() uint64 {
	return math.MaxUint64
}

func (*NegotiationResponse501) RequestTimestampId() uint16 {
	return 39002
}

func (*NegotiationResponse501) RequestTimestampSinceVersion() uint16 {
	return 0
}

func (n *NegotiationResponse501) RequestTimestampInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.RequestTimestampSinceVersion()
}

func (*NegotiationResponse501) RequestTimestampDeprecated() uint16 {
	return 0
}

func (*NegotiationResponse501) RequestTimestampMetaAttribute(meta int) string {
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

func (*NegotiationResponse501) RequestTimestampMinValue() uint64 {
	return 0
}

func (*NegotiationResponse501) RequestTimestampMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*NegotiationResponse501) RequestTimestampNullValue() uint64 {
	return math.MaxUint64
}

func (*NegotiationResponse501) SecretKeySecureIDExpirationId() uint16 {
	return 39022
}

func (*NegotiationResponse501) SecretKeySecureIDExpirationSinceVersion() uint16 {
	return 0
}

func (n *NegotiationResponse501) SecretKeySecureIDExpirationInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.SecretKeySecureIDExpirationSinceVersion()
}

func (*NegotiationResponse501) SecretKeySecureIDExpirationDeprecated() uint16 {
	return 0
}

func (*NegotiationResponse501) SecretKeySecureIDExpirationMetaAttribute(meta int) string {
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

func (*NegotiationResponse501) SecretKeySecureIDExpirationMinValue() uint16 {
	return 0
}

func (*NegotiationResponse501) SecretKeySecureIDExpirationMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*NegotiationResponse501) SecretKeySecureIDExpirationNullValue() uint16 {
	return 65535
}

func (*NegotiationResponse501) FaultToleranceIndicatorId() uint16 {
	return 39010
}

func (*NegotiationResponse501) FaultToleranceIndicatorSinceVersion() uint16 {
	return 0
}

func (n *NegotiationResponse501) FaultToleranceIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.FaultToleranceIndicatorSinceVersion()
}

func (*NegotiationResponse501) FaultToleranceIndicatorDeprecated() uint16 {
	return 0
}

func (*NegotiationResponse501) FaultToleranceIndicatorMetaAttribute(meta int) string {
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

func (*NegotiationResponse501) SplitMsgId() uint16 {
	return 9553
}

func (*NegotiationResponse501) SplitMsgSinceVersion() uint16 {
	return 0
}

func (n *NegotiationResponse501) SplitMsgInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.SplitMsgSinceVersion()
}

func (*NegotiationResponse501) SplitMsgDeprecated() uint16 {
	return 0
}

func (*NegotiationResponse501) SplitMsgMetaAttribute(meta int) string {
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

func (*NegotiationResponse501) PreviousSeqNoId() uint16 {
	return 39021
}

func (*NegotiationResponse501) PreviousSeqNoSinceVersion() uint16 {
	return 0
}

func (n *NegotiationResponse501) PreviousSeqNoInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.PreviousSeqNoSinceVersion()
}

func (*NegotiationResponse501) PreviousSeqNoDeprecated() uint16 {
	return 0
}

func (*NegotiationResponse501) PreviousSeqNoMetaAttribute(meta int) string {
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

func (*NegotiationResponse501) PreviousSeqNoMinValue() uint32 {
	return 0
}

func (*NegotiationResponse501) PreviousSeqNoMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*NegotiationResponse501) PreviousSeqNoNullValue() uint32 {
	return math.MaxUint32
}

func (*NegotiationResponse501) PreviousUUIDId() uint16 {
	return 39015
}

func (*NegotiationResponse501) PreviousUUIDSinceVersion() uint16 {
	return 0
}

func (n *NegotiationResponse501) PreviousUUIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.PreviousUUIDSinceVersion()
}

func (*NegotiationResponse501) PreviousUUIDDeprecated() uint16 {
	return 0
}

func (*NegotiationResponse501) PreviousUUIDMetaAttribute(meta int) string {
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

func (*NegotiationResponse501) PreviousUUIDMinValue() uint64 {
	return 0
}

func (*NegotiationResponse501) PreviousUUIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*NegotiationResponse501) PreviousUUIDNullValue() uint64 {
	return math.MaxUint64
}

func (*NegotiationResponse501) CredentialsMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "data"
	case 4:
		return "required"
	}
	return ""
}

func (*NegotiationResponse501) CredentialsSinceVersion() uint16 {
	return 0
}

func (n *NegotiationResponse501) CredentialsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.CredentialsSinceVersion()
}

func (*NegotiationResponse501) CredentialsDeprecated() uint16 {
	return 0
}

func (NegotiationResponse501) CredentialsCharacterEncoding() string {
	return "US-ASCII"
}

func (NegotiationResponse501) CredentialsHeaderLength() uint64 {
	return 2
}
