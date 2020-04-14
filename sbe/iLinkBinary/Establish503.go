// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type Establish503 struct {
	HMACVersion          [13]byte
	HMACSignature        [32]byte
	AccessKeyID          [20]byte
	TradingSystemName    [30]byte
	TradingSystemVersion [10]byte
	TradingSystemVendor  [10]byte
	UUID                 uint64
	RequestTimestamp     uint64
	NextSeqNo            uint32
	Session              [3]byte
	Firm                 [5]byte
	KeepAliveInterval    uint16
	Credentials          []byte
}

func (e *Establish503) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := e.RangeCheck(e.SbeSchemaVersion(), e.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteBytes(_w, e.HMACSignature[:]); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, e.AccessKeyID[:]); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, e.TradingSystemName[:]); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, e.TradingSystemVersion[:]); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, e.TradingSystemVendor[:]); err != nil {
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
	if err := _m.WriteBytes(_w, e.Session[:]); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, e.Firm[:]); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, e.KeepAliveInterval); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, uint16(len(e.Credentials))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, e.Credentials); err != nil {
		return err
	}
	return nil
}

func (e *Establish503) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	copy(e.HMACVersion[:], "CME-1-SHA-256")
	if !e.HMACSignatureInActingVersion(actingVersion) {
		for idx := 0; idx < 32; idx++ {
			e.HMACSignature[idx] = e.HMACSignatureNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, e.HMACSignature[:]); err != nil {
			return err
		}
	}
	if !e.AccessKeyIDInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			e.AccessKeyID[idx] = e.AccessKeyIDNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, e.AccessKeyID[:]); err != nil {
			return err
		}
	}
	if !e.TradingSystemNameInActingVersion(actingVersion) {
		for idx := 0; idx < 30; idx++ {
			e.TradingSystemName[idx] = e.TradingSystemNameNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, e.TradingSystemName[:]); err != nil {
			return err
		}
	}
	if !e.TradingSystemVersionInActingVersion(actingVersion) {
		for idx := 0; idx < 10; idx++ {
			e.TradingSystemVersion[idx] = e.TradingSystemVersionNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, e.TradingSystemVersion[:]); err != nil {
			return err
		}
	}
	if !e.TradingSystemVendorInActingVersion(actingVersion) {
		for idx := 0; idx < 10; idx++ {
			e.TradingSystemVendor[idx] = e.TradingSystemVendorNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, e.TradingSystemVendor[:]); err != nil {
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
	if !e.SessionInActingVersion(actingVersion) {
		for idx := 0; idx < 3; idx++ {
			e.Session[idx] = e.SessionNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, e.Session[:]); err != nil {
			return err
		}
	}
	if !e.FirmInActingVersion(actingVersion) {
		for idx := 0; idx < 5; idx++ {
			e.Firm[idx] = e.FirmNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, e.Firm[:]); err != nil {
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
	if actingVersion > e.SbeSchemaVersion() && blockLength > e.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-e.SbeBlockLength()))
	}

	if e.CredentialsInActingVersion(actingVersion) {
		var CredentialsLength uint16
		if err := _m.ReadUint16(_r, &CredentialsLength); err != nil {
			return err
		}
		if cap(e.Credentials) < int(CredentialsLength) {
			e.Credentials = make([]byte, CredentialsLength)
		}
		if err := _m.ReadBytes(_r, e.Credentials); err != nil {
			return err
		}
	}
	if doRangeCheck {
		if err := e.RangeCheck(actingVersion, e.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (e *Establish503) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if e.HMACSignatureInActingVersion(actingVersion) {
		for idx := 0; idx < 32; idx++ {
			if e.HMACSignature[idx] < e.HMACSignatureMinValue() || e.HMACSignature[idx] > e.HMACSignatureMaxValue() {
				return fmt.Errorf("Range check failed on e.HMACSignature[%d] (%v < %v > %v)", idx, e.HMACSignatureMinValue(), e.HMACSignature[idx], e.HMACSignatureMaxValue())
			}
		}
	}
	if e.AccessKeyIDInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			if e.AccessKeyID[idx] < e.AccessKeyIDMinValue() || e.AccessKeyID[idx] > e.AccessKeyIDMaxValue() {
				return fmt.Errorf("Range check failed on e.AccessKeyID[%d] (%v < %v > %v)", idx, e.AccessKeyIDMinValue(), e.AccessKeyID[idx], e.AccessKeyIDMaxValue())
			}
		}
	}
	if e.TradingSystemNameInActingVersion(actingVersion) {
		for idx := 0; idx < 30; idx++ {
			if e.TradingSystemName[idx] < e.TradingSystemNameMinValue() || e.TradingSystemName[idx] > e.TradingSystemNameMaxValue() {
				return fmt.Errorf("Range check failed on e.TradingSystemName[%d] (%v < %v > %v)", idx, e.TradingSystemNameMinValue(), e.TradingSystemName[idx], e.TradingSystemNameMaxValue())
			}
		}
	}
	if e.TradingSystemVersionInActingVersion(actingVersion) {
		for idx := 0; idx < 10; idx++ {
			if e.TradingSystemVersion[idx] < e.TradingSystemVersionMinValue() || e.TradingSystemVersion[idx] > e.TradingSystemVersionMaxValue() {
				return fmt.Errorf("Range check failed on e.TradingSystemVersion[%d] (%v < %v > %v)", idx, e.TradingSystemVersionMinValue(), e.TradingSystemVersion[idx], e.TradingSystemVersionMaxValue())
			}
		}
	}
	if e.TradingSystemVendorInActingVersion(actingVersion) {
		for idx := 0; idx < 10; idx++ {
			if e.TradingSystemVendor[idx] < e.TradingSystemVendorMinValue() || e.TradingSystemVendor[idx] > e.TradingSystemVendorMaxValue() {
				return fmt.Errorf("Range check failed on e.TradingSystemVendor[%d] (%v < %v > %v)", idx, e.TradingSystemVendorMinValue(), e.TradingSystemVendor[idx], e.TradingSystemVendorMaxValue())
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
	if e.SessionInActingVersion(actingVersion) {
		for idx := 0; idx < 3; idx++ {
			if e.Session[idx] < e.SessionMinValue() || e.Session[idx] > e.SessionMaxValue() {
				return fmt.Errorf("Range check failed on e.Session[%d] (%v < %v > %v)", idx, e.SessionMinValue(), e.Session[idx], e.SessionMaxValue())
			}
		}
	}
	if e.FirmInActingVersion(actingVersion) {
		for idx := 0; idx < 5; idx++ {
			if e.Firm[idx] < e.FirmMinValue() || e.Firm[idx] > e.FirmMaxValue() {
				return fmt.Errorf("Range check failed on e.Firm[%d] (%v < %v > %v)", idx, e.FirmMinValue(), e.Firm[idx], e.FirmMaxValue())
			}
		}
	}
	if e.KeepAliveIntervalInActingVersion(actingVersion) {
		if e.KeepAliveInterval < e.KeepAliveIntervalMinValue() || e.KeepAliveInterval > e.KeepAliveIntervalMaxValue() {
			return fmt.Errorf("Range check failed on e.KeepAliveInterval (%v < %v > %v)", e.KeepAliveIntervalMinValue(), e.KeepAliveInterval, e.KeepAliveIntervalMaxValue())
		}
	}
	return nil
}

func Establish503Init(e *Establish503) {
	copy(e.HMACVersion[:], "CME-1-SHA-256")
	return
}

func (*Establish503) SbeBlockLength() (blockLength uint16) {
	return 132
}

func (*Establish503) SbeTemplateId() (templateId uint16) {
	return 503
}

func (*Establish503) SbeSchemaId() (schemaId uint16) {
	return 8
}

func (*Establish503) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}

func (*Establish503) SbeSemanticType() (semanticType []byte) {
	return []byte("Establish")
}

func (*Establish503) HMACVersionId() uint16 {
	return 39003
}

func (*Establish503) HMACVersionSinceVersion() uint16 {
	return 0
}

func (e *Establish503) HMACVersionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.HMACVersionSinceVersion()
}

func (*Establish503) HMACVersionDeprecated() uint16 {
	return 0
}

func (*Establish503) HMACVersionMetaAttribute(meta int) string {
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

func (*Establish503) HMACVersionMinValue() byte {
	return byte(32)
}

func (*Establish503) HMACVersionMaxValue() byte {
	return byte(126)
}

func (*Establish503) HMACVersionNullValue() byte {
	return 0
}

func (*Establish503) HMACSignatureId() uint16 {
	return 39005
}

func (*Establish503) HMACSignatureSinceVersion() uint16 {
	return 0
}

func (e *Establish503) HMACSignatureInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.HMACSignatureSinceVersion()
}

func (*Establish503) HMACSignatureDeprecated() uint16 {
	return 0
}

func (*Establish503) HMACSignatureMetaAttribute(meta int) string {
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

func (*Establish503) HMACSignatureMinValue() byte {
	return byte(32)
}

func (*Establish503) HMACSignatureMaxValue() byte {
	return byte(126)
}

func (*Establish503) HMACSignatureNullValue() byte {
	return 0
}

func (e *Establish503) HMACSignatureCharacterEncoding() string {
	return "US-ASCII"
}

func (*Establish503) AccessKeyIDId() uint16 {
	return 39004
}

func (*Establish503) AccessKeyIDSinceVersion() uint16 {
	return 0
}

func (e *Establish503) AccessKeyIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.AccessKeyIDSinceVersion()
}

func (*Establish503) AccessKeyIDDeprecated() uint16 {
	return 0
}

func (*Establish503) AccessKeyIDMetaAttribute(meta int) string {
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

func (*Establish503) AccessKeyIDMinValue() byte {
	return byte(32)
}

func (*Establish503) AccessKeyIDMaxValue() byte {
	return byte(126)
}

func (*Establish503) AccessKeyIDNullValue() byte {
	return 0
}

func (e *Establish503) AccessKeyIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*Establish503) TradingSystemNameId() uint16 {
	return 1603
}

func (*Establish503) TradingSystemNameSinceVersion() uint16 {
	return 0
}

func (e *Establish503) TradingSystemNameInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TradingSystemNameSinceVersion()
}

func (*Establish503) TradingSystemNameDeprecated() uint16 {
	return 0
}

func (*Establish503) TradingSystemNameMetaAttribute(meta int) string {
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

func (*Establish503) TradingSystemNameMinValue() byte {
	return byte(32)
}

func (*Establish503) TradingSystemNameMaxValue() byte {
	return byte(126)
}

func (*Establish503) TradingSystemNameNullValue() byte {
	return 0
}

func (e *Establish503) TradingSystemNameCharacterEncoding() string {
	return "US-ASCII"
}

func (*Establish503) TradingSystemVersionId() uint16 {
	return 1604
}

func (*Establish503) TradingSystemVersionSinceVersion() uint16 {
	return 0
}

func (e *Establish503) TradingSystemVersionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TradingSystemVersionSinceVersion()
}

func (*Establish503) TradingSystemVersionDeprecated() uint16 {
	return 0
}

func (*Establish503) TradingSystemVersionMetaAttribute(meta int) string {
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

func (*Establish503) TradingSystemVersionMinValue() byte {
	return byte(32)
}

func (*Establish503) TradingSystemVersionMaxValue() byte {
	return byte(126)
}

func (*Establish503) TradingSystemVersionNullValue() byte {
	return 0
}

func (e *Establish503) TradingSystemVersionCharacterEncoding() string {
	return "US-ASCII"
}

func (*Establish503) TradingSystemVendorId() uint16 {
	return 1605
}

func (*Establish503) TradingSystemVendorSinceVersion() uint16 {
	return 0
}

func (e *Establish503) TradingSystemVendorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TradingSystemVendorSinceVersion()
}

func (*Establish503) TradingSystemVendorDeprecated() uint16 {
	return 0
}

func (*Establish503) TradingSystemVendorMetaAttribute(meta int) string {
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

func (*Establish503) TradingSystemVendorMinValue() byte {
	return byte(32)
}

func (*Establish503) TradingSystemVendorMaxValue() byte {
	return byte(126)
}

func (*Establish503) TradingSystemVendorNullValue() byte {
	return 0
}

func (e *Establish503) TradingSystemVendorCharacterEncoding() string {
	return "US-ASCII"
}

func (*Establish503) UUIDId() uint16 {
	return 39001
}

func (*Establish503) UUIDSinceVersion() uint16 {
	return 0
}

func (e *Establish503) UUIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.UUIDSinceVersion()
}

func (*Establish503) UUIDDeprecated() uint16 {
	return 0
}

func (*Establish503) UUIDMetaAttribute(meta int) string {
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

func (*Establish503) UUIDMinValue() uint64 {
	return 0
}

func (*Establish503) UUIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*Establish503) UUIDNullValue() uint64 {
	return math.MaxUint64
}

func (*Establish503) RequestTimestampId() uint16 {
	return 39002
}

func (*Establish503) RequestTimestampSinceVersion() uint16 {
	return 0
}

func (e *Establish503) RequestTimestampInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.RequestTimestampSinceVersion()
}

func (*Establish503) RequestTimestampDeprecated() uint16 {
	return 0
}

func (*Establish503) RequestTimestampMetaAttribute(meta int) string {
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

func (*Establish503) RequestTimestampMinValue() uint64 {
	return 0
}

func (*Establish503) RequestTimestampMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*Establish503) RequestTimestampNullValue() uint64 {
	return math.MaxUint64
}

func (*Establish503) NextSeqNoId() uint16 {
	return 39013
}

func (*Establish503) NextSeqNoSinceVersion() uint16 {
	return 0
}

func (e *Establish503) NextSeqNoInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.NextSeqNoSinceVersion()
}

func (*Establish503) NextSeqNoDeprecated() uint16 {
	return 0
}

func (*Establish503) NextSeqNoMetaAttribute(meta int) string {
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

func (*Establish503) NextSeqNoMinValue() uint32 {
	return 0
}

func (*Establish503) NextSeqNoMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*Establish503) NextSeqNoNullValue() uint32 {
	return math.MaxUint32
}

func (*Establish503) SessionId() uint16 {
	return 39006
}

func (*Establish503) SessionSinceVersion() uint16 {
	return 0
}

func (e *Establish503) SessionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SessionSinceVersion()
}

func (*Establish503) SessionDeprecated() uint16 {
	return 0
}

func (*Establish503) SessionMetaAttribute(meta int) string {
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

func (*Establish503) SessionMinValue() byte {
	return byte(32)
}

func (*Establish503) SessionMaxValue() byte {
	return byte(126)
}

func (*Establish503) SessionNullValue() byte {
	return 0
}

func (e *Establish503) SessionCharacterEncoding() string {
	return "US-ASCII"
}

func (*Establish503) FirmId() uint16 {
	return 39007
}

func (*Establish503) FirmSinceVersion() uint16 {
	return 0
}

func (e *Establish503) FirmInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.FirmSinceVersion()
}

func (*Establish503) FirmDeprecated() uint16 {
	return 0
}

func (*Establish503) FirmMetaAttribute(meta int) string {
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

func (*Establish503) FirmMinValue() byte {
	return byte(32)
}

func (*Establish503) FirmMaxValue() byte {
	return byte(126)
}

func (*Establish503) FirmNullValue() byte {
	return 0
}

func (e *Establish503) FirmCharacterEncoding() string {
	return "US-ASCII"
}

func (*Establish503) KeepAliveIntervalId() uint16 {
	return 39014
}

func (*Establish503) KeepAliveIntervalSinceVersion() uint16 {
	return 0
}

func (e *Establish503) KeepAliveIntervalInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.KeepAliveIntervalSinceVersion()
}

func (*Establish503) KeepAliveIntervalDeprecated() uint16 {
	return 0
}

func (*Establish503) KeepAliveIntervalMetaAttribute(meta int) string {
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

func (*Establish503) KeepAliveIntervalMinValue() uint16 {
	return 0
}

func (*Establish503) KeepAliveIntervalMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*Establish503) KeepAliveIntervalNullValue() uint16 {
	return math.MaxUint16
}

func (*Establish503) CredentialsMetaAttribute(meta int) string {
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

func (*Establish503) CredentialsSinceVersion() uint16 {
	return 0
}

func (e *Establish503) CredentialsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.CredentialsSinceVersion()
}

func (*Establish503) CredentialsDeprecated() uint16 {
	return 0
}

func (Establish503) CredentialsCharacterEncoding() string {
	return "US-ASCII"
}

func (Establish503) CredentialsHeaderLength() uint64 {
	return 2
}
