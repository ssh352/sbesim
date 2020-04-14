// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type Negotiate500 struct {
	CustomerFlow     [10]byte
	HMACVersion      [13]byte
	HMACSignature    [32]byte
	AccessKeyID      [20]byte
	UUID             uint64
	RequestTimestamp uint64
	Session          [3]byte
	Firm             [5]byte
	Credentials      []byte
}

func (n *Negotiate500) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := n.RangeCheck(n.SbeSchemaVersion(), n.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteBytes(_w, n.HMACSignature[:]); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, n.AccessKeyID[:]); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, n.UUID); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, n.RequestTimestamp); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, n.Session[:]); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, n.Firm[:]); err != nil {
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

func (n *Negotiate500) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	copy(n.CustomerFlow[:], "IDEMPOTENT")
	copy(n.HMACVersion[:], "CME-1-SHA-256")
	if !n.HMACSignatureInActingVersion(actingVersion) {
		for idx := 0; idx < 32; idx++ {
			n.HMACSignature[idx] = n.HMACSignatureNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, n.HMACSignature[:]); err != nil {
			return err
		}
	}
	if !n.AccessKeyIDInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			n.AccessKeyID[idx] = n.AccessKeyIDNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, n.AccessKeyID[:]); err != nil {
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
	if !n.SessionInActingVersion(actingVersion) {
		for idx := 0; idx < 3; idx++ {
			n.Session[idx] = n.SessionNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, n.Session[:]); err != nil {
			return err
		}
	}
	if !n.FirmInActingVersion(actingVersion) {
		for idx := 0; idx < 5; idx++ {
			n.Firm[idx] = n.FirmNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, n.Firm[:]); err != nil {
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

func (n *Negotiate500) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if n.HMACSignatureInActingVersion(actingVersion) {
		for idx := 0; idx < 32; idx++ {
			if n.HMACSignature[idx] < n.HMACSignatureMinValue() || n.HMACSignature[idx] > n.HMACSignatureMaxValue() {
				return fmt.Errorf("Range check failed on n.HMACSignature[%d] (%v < %v > %v)", idx, n.HMACSignatureMinValue(), n.HMACSignature[idx], n.HMACSignatureMaxValue())
			}
		}
	}
	if n.AccessKeyIDInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			if n.AccessKeyID[idx] < n.AccessKeyIDMinValue() || n.AccessKeyID[idx] > n.AccessKeyIDMaxValue() {
				return fmt.Errorf("Range check failed on n.AccessKeyID[%d] (%v < %v > %v)", idx, n.AccessKeyIDMinValue(), n.AccessKeyID[idx], n.AccessKeyIDMaxValue())
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
	if n.SessionInActingVersion(actingVersion) {
		for idx := 0; idx < 3; idx++ {
			if n.Session[idx] < n.SessionMinValue() || n.Session[idx] > n.SessionMaxValue() {
				return fmt.Errorf("Range check failed on n.Session[%d] (%v < %v > %v)", idx, n.SessionMinValue(), n.Session[idx], n.SessionMaxValue())
			}
		}
	}
	if n.FirmInActingVersion(actingVersion) {
		for idx := 0; idx < 5; idx++ {
			if n.Firm[idx] < n.FirmMinValue() || n.Firm[idx] > n.FirmMaxValue() {
				return fmt.Errorf("Range check failed on n.Firm[%d] (%v < %v > %v)", idx, n.FirmMinValue(), n.Firm[idx], n.FirmMaxValue())
			}
		}
	}
	return nil
}

func Negotiate500Init(n *Negotiate500) {
	copy(n.CustomerFlow[:], "IDEMPOTENT")
	copy(n.HMACVersion[:], "CME-1-SHA-256")
	return
}

func (*Negotiate500) SbeBlockLength() (blockLength uint16) {
	return 76
}

func (*Negotiate500) SbeTemplateId() (templateId uint16) {
	return 500
}

func (*Negotiate500) SbeSchemaId() (schemaId uint16) {
	return 8
}

func (*Negotiate500) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}

func (*Negotiate500) SbeSemanticType() (semanticType []byte) {
	return []byte("Negotiate")
}

func (*Negotiate500) CustomerFlowId() uint16 {
	return 39000
}

func (*Negotiate500) CustomerFlowSinceVersion() uint16 {
	return 0
}

func (n *Negotiate500) CustomerFlowInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.CustomerFlowSinceVersion()
}

func (*Negotiate500) CustomerFlowDeprecated() uint16 {
	return 0
}

func (*Negotiate500) CustomerFlowMetaAttribute(meta int) string {
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

func (*Negotiate500) CustomerFlowMinValue() byte {
	return byte(32)
}

func (*Negotiate500) CustomerFlowMaxValue() byte {
	return byte(126)
}

func (*Negotiate500) CustomerFlowNullValue() byte {
	return 0
}

func (*Negotiate500) HMACVersionId() uint16 {
	return 39003
}

func (*Negotiate500) HMACVersionSinceVersion() uint16 {
	return 0
}

func (n *Negotiate500) HMACVersionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.HMACVersionSinceVersion()
}

func (*Negotiate500) HMACVersionDeprecated() uint16 {
	return 0
}

func (*Negotiate500) HMACVersionMetaAttribute(meta int) string {
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

func (*Negotiate500) HMACVersionMinValue() byte {
	return byte(32)
}

func (*Negotiate500) HMACVersionMaxValue() byte {
	return byte(126)
}

func (*Negotiate500) HMACVersionNullValue() byte {
	return 0
}

func (*Negotiate500) HMACSignatureId() uint16 {
	return 39005
}

func (*Negotiate500) HMACSignatureSinceVersion() uint16 {
	return 0
}

func (n *Negotiate500) HMACSignatureInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.HMACSignatureSinceVersion()
}

func (*Negotiate500) HMACSignatureDeprecated() uint16 {
	return 0
}

func (*Negotiate500) HMACSignatureMetaAttribute(meta int) string {
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

func (*Negotiate500) HMACSignatureMinValue() byte {
	return byte(32)
}

func (*Negotiate500) HMACSignatureMaxValue() byte {
	return byte(126)
}

func (*Negotiate500) HMACSignatureNullValue() byte {
	return 0
}

func (n *Negotiate500) HMACSignatureCharacterEncoding() string {
	return "US-ASCII"
}

func (*Negotiate500) AccessKeyIDId() uint16 {
	return 39004
}

func (*Negotiate500) AccessKeyIDSinceVersion() uint16 {
	return 0
}

func (n *Negotiate500) AccessKeyIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.AccessKeyIDSinceVersion()
}

func (*Negotiate500) AccessKeyIDDeprecated() uint16 {
	return 0
}

func (*Negotiate500) AccessKeyIDMetaAttribute(meta int) string {
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

func (*Negotiate500) AccessKeyIDMinValue() byte {
	return byte(32)
}

func (*Negotiate500) AccessKeyIDMaxValue() byte {
	return byte(126)
}

func (*Negotiate500) AccessKeyIDNullValue() byte {
	return 0
}

func (n *Negotiate500) AccessKeyIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*Negotiate500) UUIDId() uint16 {
	return 39001
}

func (*Negotiate500) UUIDSinceVersion() uint16 {
	return 0
}

func (n *Negotiate500) UUIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.UUIDSinceVersion()
}

func (*Negotiate500) UUIDDeprecated() uint16 {
	return 0
}

func (*Negotiate500) UUIDMetaAttribute(meta int) string {
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

func (*Negotiate500) UUIDMinValue() uint64 {
	return 0
}

func (*Negotiate500) UUIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*Negotiate500) UUIDNullValue() uint64 {
	return math.MaxUint64
}

func (*Negotiate500) RequestTimestampId() uint16 {
	return 39002
}

func (*Negotiate500) RequestTimestampSinceVersion() uint16 {
	return 0
}

func (n *Negotiate500) RequestTimestampInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.RequestTimestampSinceVersion()
}

func (*Negotiate500) RequestTimestampDeprecated() uint16 {
	return 0
}

func (*Negotiate500) RequestTimestampMetaAttribute(meta int) string {
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

func (*Negotiate500) RequestTimestampMinValue() uint64 {
	return 0
}

func (*Negotiate500) RequestTimestampMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*Negotiate500) RequestTimestampNullValue() uint64 {
	return math.MaxUint64
}

func (*Negotiate500) SessionId() uint16 {
	return 39006
}

func (*Negotiate500) SessionSinceVersion() uint16 {
	return 0
}

func (n *Negotiate500) SessionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.SessionSinceVersion()
}

func (*Negotiate500) SessionDeprecated() uint16 {
	return 0
}

func (*Negotiate500) SessionMetaAttribute(meta int) string {
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

func (*Negotiate500) SessionMinValue() byte {
	return byte(32)
}

func (*Negotiate500) SessionMaxValue() byte {
	return byte(126)
}

func (*Negotiate500) SessionNullValue() byte {
	return 0
}

func (n *Negotiate500) SessionCharacterEncoding() string {
	return "US-ASCII"
}

func (*Negotiate500) FirmId() uint16 {
	return 39007
}

func (*Negotiate500) FirmSinceVersion() uint16 {
	return 0
}

func (n *Negotiate500) FirmInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.FirmSinceVersion()
}

func (*Negotiate500) FirmDeprecated() uint16 {
	return 0
}

func (*Negotiate500) FirmMetaAttribute(meta int) string {
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

func (*Negotiate500) FirmMinValue() byte {
	return byte(32)
}

func (*Negotiate500) FirmMaxValue() byte {
	return byte(126)
}

func (*Negotiate500) FirmNullValue() byte {
	return 0
}

func (n *Negotiate500) FirmCharacterEncoding() string {
	return "US-ASCII"
}

func (*Negotiate500) CredentialsMetaAttribute(meta int) string {
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

func (*Negotiate500) CredentialsSinceVersion() uint16 {
	return 0
}

func (n *Negotiate500) CredentialsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.CredentialsSinceVersion()
}

func (*Negotiate500) CredentialsDeprecated() uint16 {
	return 0
}

func (Negotiate500) CredentialsCharacterEncoding() string {
	return "US-ASCII"
}

func (Negotiate500) CredentialsHeaderLength() uint64 {
	return 2
}
