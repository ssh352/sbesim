// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type PartyDetailsListRequest537 struct {
	PartyDetailsListReqID uint64
	SendingTimeEpoch      uint64
	SeqNum                uint32
	NoRequestingPartyIDs  []PartyDetailsListRequest537NoRequestingPartyIDs
	NoPartyIDs            []PartyDetailsListRequest537NoPartyIDs
}
type PartyDetailsListRequest537NoRequestingPartyIDs struct {
	RequestingPartyID       [5]byte
	RequestingPartyIDSource byte
	RequestingPartyRole     byte
}
type PartyDetailsListRequest537NoPartyIDs struct {
	PartyID       uint64
	PartyIDSource byte
	PartyRole     uint16
}

func (p *PartyDetailsListRequest537) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := p.RangeCheck(p.SbeSchemaVersion(), p.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteUint64(_w, p.PartyDetailsListReqID); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, p.SendingTimeEpoch); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, p.SeqNum); err != nil {
		return err
	}
	var NoRequestingPartyIDsBlockLength uint16 = 7
	if err := _m.WriteUint16(_w, NoRequestingPartyIDsBlockLength); err != nil {
		return err
	}
	var NoRequestingPartyIDsNumInGroup uint8 = uint8(len(p.NoRequestingPartyIDs))
	if err := _m.WriteUint8(_w, NoRequestingPartyIDsNumInGroup); err != nil {
		return err
	}
	for _, prop := range p.NoRequestingPartyIDs {
		if err := prop.Encode(_m, _w); err != nil {
			return err
		}
	}
	var NoPartyIDsBlockLength uint16 = 11
	if err := _m.WriteUint16(_w, NoPartyIDsBlockLength); err != nil {
		return err
	}
	var NoPartyIDsNumInGroup uint8 = uint8(len(p.NoPartyIDs))
	if err := _m.WriteUint8(_w, NoPartyIDsNumInGroup); err != nil {
		return err
	}
	for _, prop := range p.NoPartyIDs {
		if err := prop.Encode(_m, _w); err != nil {
			return err
		}
	}
	return nil
}

func (p *PartyDetailsListRequest537) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !p.PartyDetailsListReqIDInActingVersion(actingVersion) {
		p.PartyDetailsListReqID = p.PartyDetailsListReqIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &p.PartyDetailsListReqID); err != nil {
			return err
		}
	}
	if !p.SendingTimeEpochInActingVersion(actingVersion) {
		p.SendingTimeEpoch = p.SendingTimeEpochNullValue()
	} else {
		if err := _m.ReadUint64(_r, &p.SendingTimeEpoch); err != nil {
			return err
		}
	}
	if !p.SeqNumInActingVersion(actingVersion) {
		p.SeqNum = p.SeqNumNullValue()
	} else {
		if err := _m.ReadUint32(_r, &p.SeqNum); err != nil {
			return err
		}
	}
	if actingVersion > p.SbeSchemaVersion() && blockLength > p.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-p.SbeBlockLength()))
	}

	if p.NoRequestingPartyIDsInActingVersion(actingVersion) {
		var NoRequestingPartyIDsBlockLength uint16
		if err := _m.ReadUint16(_r, &NoRequestingPartyIDsBlockLength); err != nil {
			return err
		}
		var NoRequestingPartyIDsNumInGroup uint8
		if err := _m.ReadUint8(_r, &NoRequestingPartyIDsNumInGroup); err != nil {
			return err
		}
		if cap(p.NoRequestingPartyIDs) < int(NoRequestingPartyIDsNumInGroup) {
			p.NoRequestingPartyIDs = make([]PartyDetailsListRequest537NoRequestingPartyIDs, NoRequestingPartyIDsNumInGroup)
		}
		for i, _ := range p.NoRequestingPartyIDs {
			if err := p.NoRequestingPartyIDs[i].Decode(_m, _r, actingVersion, uint(NoRequestingPartyIDsBlockLength)); err != nil {
				return err
			}
		}
	}

	if p.NoPartyIDsInActingVersion(actingVersion) {
		var NoPartyIDsBlockLength uint16
		if err := _m.ReadUint16(_r, &NoPartyIDsBlockLength); err != nil {
			return err
		}
		var NoPartyIDsNumInGroup uint8
		if err := _m.ReadUint8(_r, &NoPartyIDsNumInGroup); err != nil {
			return err
		}
		if cap(p.NoPartyIDs) < int(NoPartyIDsNumInGroup) {
			p.NoPartyIDs = make([]PartyDetailsListRequest537NoPartyIDs, NoPartyIDsNumInGroup)
		}
		for i, _ := range p.NoPartyIDs {
			if err := p.NoPartyIDs[i].Decode(_m, _r, actingVersion, uint(NoPartyIDsBlockLength)); err != nil {
				return err
			}
		}
	}
	if doRangeCheck {
		if err := p.RangeCheck(actingVersion, p.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (p *PartyDetailsListRequest537) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if p.PartyDetailsListReqIDInActingVersion(actingVersion) {
		if p.PartyDetailsListReqID < p.PartyDetailsListReqIDMinValue() || p.PartyDetailsListReqID > p.PartyDetailsListReqIDMaxValue() {
			return fmt.Errorf("Range check failed on p.PartyDetailsListReqID (%v < %v > %v)", p.PartyDetailsListReqIDMinValue(), p.PartyDetailsListReqID, p.PartyDetailsListReqIDMaxValue())
		}
	}
	if p.SendingTimeEpochInActingVersion(actingVersion) {
		if p.SendingTimeEpoch < p.SendingTimeEpochMinValue() || p.SendingTimeEpoch > p.SendingTimeEpochMaxValue() {
			return fmt.Errorf("Range check failed on p.SendingTimeEpoch (%v < %v > %v)", p.SendingTimeEpochMinValue(), p.SendingTimeEpoch, p.SendingTimeEpochMaxValue())
		}
	}
	if p.SeqNumInActingVersion(actingVersion) {
		if p.SeqNum < p.SeqNumMinValue() || p.SeqNum > p.SeqNumMaxValue() {
			return fmt.Errorf("Range check failed on p.SeqNum (%v < %v > %v)", p.SeqNumMinValue(), p.SeqNum, p.SeqNumMaxValue())
		}
	}
	for _, prop := range p.NoRequestingPartyIDs {
		if err := prop.RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	for _, prop := range p.NoPartyIDs {
		if err := prop.RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	return nil
}

func PartyDetailsListRequest537Init(p *PartyDetailsListRequest537) {
	return
}

func (p *PartyDetailsListRequest537NoRequestingPartyIDs) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteBytes(_w, p.RequestingPartyID[:]); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, p.RequestingPartyIDSource); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, p.RequestingPartyRole); err != nil {
		return err
	}
	return nil
}

func (p *PartyDetailsListRequest537NoRequestingPartyIDs) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	if !p.RequestingPartyIDInActingVersion(actingVersion) {
		for idx := 0; idx < 5; idx++ {
			p.RequestingPartyID[idx] = p.RequestingPartyIDNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, p.RequestingPartyID[:]); err != nil {
			return err
		}
	}
	if !p.RequestingPartyIDSourceInActingVersion(actingVersion) {
		p.RequestingPartyIDSource = p.RequestingPartyIDSourceNullValue()
	} else {
		if err := _m.ReadUint8(_r, &p.RequestingPartyIDSource); err != nil {
			return err
		}
	}
	if !p.RequestingPartyRoleInActingVersion(actingVersion) {
		p.RequestingPartyRole = p.RequestingPartyRoleNullValue()
	} else {
		if err := _m.ReadUint8(_r, &p.RequestingPartyRole); err != nil {
			return err
		}
	}
	if actingVersion > p.SbeSchemaVersion() && blockLength > p.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-p.SbeBlockLength()))
	}
	return nil
}

func (p *PartyDetailsListRequest537NoRequestingPartyIDs) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if p.RequestingPartyIDInActingVersion(actingVersion) {
		for idx := 0; idx < 5; idx++ {
			if p.RequestingPartyID[idx] < p.RequestingPartyIDMinValue() || p.RequestingPartyID[idx] > p.RequestingPartyIDMaxValue() {
				return fmt.Errorf("Range check failed on p.RequestingPartyID[%d] (%v < %v > %v)", idx, p.RequestingPartyIDMinValue(), p.RequestingPartyID[idx], p.RequestingPartyIDMaxValue())
			}
		}
	}
	if p.RequestingPartyIDSourceInActingVersion(actingVersion) {
		if p.RequestingPartyIDSource < p.RequestingPartyIDSourceMinValue() || p.RequestingPartyIDSource > p.RequestingPartyIDSourceMaxValue() {
			return fmt.Errorf("Range check failed on p.RequestingPartyIDSource (%v < %v > %v)", p.RequestingPartyIDSourceMinValue(), p.RequestingPartyIDSource, p.RequestingPartyIDSourceMaxValue())
		}
	}
	if p.RequestingPartyRoleInActingVersion(actingVersion) {
		if p.RequestingPartyRole < p.RequestingPartyRoleMinValue() || p.RequestingPartyRole > p.RequestingPartyRoleMaxValue() {
			return fmt.Errorf("Range check failed on p.RequestingPartyRole (%v < %v > %v)", p.RequestingPartyRoleMinValue(), p.RequestingPartyRole, p.RequestingPartyRoleMaxValue())
		}
	}
	return nil
}

func PartyDetailsListRequest537NoRequestingPartyIDsInit(p *PartyDetailsListRequest537NoRequestingPartyIDs) {
	for idx := 0; idx < 5; idx++ {
		p.RequestingPartyID[idx] = 0
	}
	return
}

func (p *PartyDetailsListRequest537NoPartyIDs) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint64(_w, p.PartyID); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, p.PartyIDSource); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, p.PartyRole); err != nil {
		return err
	}
	return nil
}

func (p *PartyDetailsListRequest537NoPartyIDs) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	if !p.PartyIDInActingVersion(actingVersion) {
		p.PartyID = p.PartyIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &p.PartyID); err != nil {
			return err
		}
	}
	if !p.PartyIDSourceInActingVersion(actingVersion) {
		p.PartyIDSource = p.PartyIDSourceNullValue()
	} else {
		if err := _m.ReadUint8(_r, &p.PartyIDSource); err != nil {
			return err
		}
	}
	if !p.PartyRoleInActingVersion(actingVersion) {
		p.PartyRole = p.PartyRoleNullValue()
	} else {
		if err := _m.ReadUint16(_r, &p.PartyRole); err != nil {
			return err
		}
	}
	if actingVersion > p.SbeSchemaVersion() && blockLength > p.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-p.SbeBlockLength()))
	}
	return nil
}

func (p *PartyDetailsListRequest537NoPartyIDs) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if p.PartyIDInActingVersion(actingVersion) {
		if p.PartyID < p.PartyIDMinValue() || p.PartyID > p.PartyIDMaxValue() {
			return fmt.Errorf("Range check failed on p.PartyID (%v < %v > %v)", p.PartyIDMinValue(), p.PartyID, p.PartyIDMaxValue())
		}
	}
	if p.PartyIDSourceInActingVersion(actingVersion) {
		if p.PartyIDSource < p.PartyIDSourceMinValue() || p.PartyIDSource > p.PartyIDSourceMaxValue() {
			return fmt.Errorf("Range check failed on p.PartyIDSource (%v < %v > %v)", p.PartyIDSourceMinValue(), p.PartyIDSource, p.PartyIDSourceMaxValue())
		}
	}
	if p.PartyRoleInActingVersion(actingVersion) {
		if p.PartyRole < p.PartyRoleMinValue() || p.PartyRole > p.PartyRoleMaxValue() {
			return fmt.Errorf("Range check failed on p.PartyRole (%v < %v > %v)", p.PartyRoleMinValue(), p.PartyRole, p.PartyRoleMaxValue())
		}
	}
	return nil
}

func PartyDetailsListRequest537NoPartyIDsInit(p *PartyDetailsListRequest537NoPartyIDs) {
	return
}

func (*PartyDetailsListRequest537) SbeBlockLength() (blockLength uint16) {
	return 20
}

func (*PartyDetailsListRequest537) SbeTemplateId() (templateId uint16) {
	return 537
}

func (*PartyDetailsListRequest537) SbeSchemaId() (schemaId uint16) {
	return 8
}

func (*PartyDetailsListRequest537) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}

func (*PartyDetailsListRequest537) SbeSemanticType() (semanticType []byte) {
	return []byte("CF")
}

func (*PartyDetailsListRequest537) PartyDetailsListReqIDId() uint16 {
	return 1505
}

func (*PartyDetailsListRequest537) PartyDetailsListReqIDSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsListRequest537) PartyDetailsListReqIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.PartyDetailsListReqIDSinceVersion()
}

func (*PartyDetailsListRequest537) PartyDetailsListReqIDDeprecated() uint16 {
	return 0
}

func (*PartyDetailsListRequest537) PartyDetailsListReqIDMetaAttribute(meta int) string {
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

func (*PartyDetailsListRequest537) PartyDetailsListReqIDMinValue() uint64 {
	return 0
}

func (*PartyDetailsListRequest537) PartyDetailsListReqIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*PartyDetailsListRequest537) PartyDetailsListReqIDNullValue() uint64 {
	return math.MaxUint64
}

func (*PartyDetailsListRequest537) SendingTimeEpochId() uint16 {
	return 5297
}

func (*PartyDetailsListRequest537) SendingTimeEpochSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsListRequest537) SendingTimeEpochInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.SendingTimeEpochSinceVersion()
}

func (*PartyDetailsListRequest537) SendingTimeEpochDeprecated() uint16 {
	return 0
}

func (*PartyDetailsListRequest537) SendingTimeEpochMetaAttribute(meta int) string {
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

func (*PartyDetailsListRequest537) SendingTimeEpochMinValue() uint64 {
	return 0
}

func (*PartyDetailsListRequest537) SendingTimeEpochMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*PartyDetailsListRequest537) SendingTimeEpochNullValue() uint64 {
	return math.MaxUint64
}

func (*PartyDetailsListRequest537) SeqNumId() uint16 {
	return 9726
}

func (*PartyDetailsListRequest537) SeqNumSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsListRequest537) SeqNumInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.SeqNumSinceVersion()
}

func (*PartyDetailsListRequest537) SeqNumDeprecated() uint16 {
	return 0
}

func (*PartyDetailsListRequest537) SeqNumMetaAttribute(meta int) string {
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

func (*PartyDetailsListRequest537) SeqNumMinValue() uint32 {
	return 0
}

func (*PartyDetailsListRequest537) SeqNumMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*PartyDetailsListRequest537) SeqNumNullValue() uint32 {
	return math.MaxUint32
}

func (*PartyDetailsListRequest537NoRequestingPartyIDs) RequestingPartyIDId() uint16 {
	return 1658
}

func (*PartyDetailsListRequest537NoRequestingPartyIDs) RequestingPartyIDSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsListRequest537NoRequestingPartyIDs) RequestingPartyIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.RequestingPartyIDSinceVersion()
}

func (*PartyDetailsListRequest537NoRequestingPartyIDs) RequestingPartyIDDeprecated() uint16 {
	return 0
}

func (*PartyDetailsListRequest537NoRequestingPartyIDs) RequestingPartyIDMetaAttribute(meta int) string {
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

func (*PartyDetailsListRequest537NoRequestingPartyIDs) RequestingPartyIDMinValue() byte {
	return byte(32)
}

func (*PartyDetailsListRequest537NoRequestingPartyIDs) RequestingPartyIDMaxValue() byte {
	return byte(126)
}

func (*PartyDetailsListRequest537NoRequestingPartyIDs) RequestingPartyIDNullValue() byte {
	return 0
}

func (p *PartyDetailsListRequest537NoRequestingPartyIDs) RequestingPartyIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*PartyDetailsListRequest537NoRequestingPartyIDs) RequestingPartyIDSourceId() uint16 {
	return 1659
}

func (*PartyDetailsListRequest537NoRequestingPartyIDs) RequestingPartyIDSourceSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsListRequest537NoRequestingPartyIDs) RequestingPartyIDSourceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.RequestingPartyIDSourceSinceVersion()
}

func (*PartyDetailsListRequest537NoRequestingPartyIDs) RequestingPartyIDSourceDeprecated() uint16 {
	return 0
}

func (*PartyDetailsListRequest537NoRequestingPartyIDs) RequestingPartyIDSourceMetaAttribute(meta int) string {
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

func (*PartyDetailsListRequest537NoRequestingPartyIDs) RequestingPartyIDSourceMinValue() byte {
	return byte(32)
}

func (*PartyDetailsListRequest537NoRequestingPartyIDs) RequestingPartyIDSourceMaxValue() byte {
	return byte(126)
}

func (*PartyDetailsListRequest537NoRequestingPartyIDs) RequestingPartyIDSourceNullValue() byte {
	return 0
}

func (*PartyDetailsListRequest537NoRequestingPartyIDs) RequestingPartyRoleId() uint16 {
	return 1660
}

func (*PartyDetailsListRequest537NoRequestingPartyIDs) RequestingPartyRoleSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsListRequest537NoRequestingPartyIDs) RequestingPartyRoleInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.RequestingPartyRoleSinceVersion()
}

func (*PartyDetailsListRequest537NoRequestingPartyIDs) RequestingPartyRoleDeprecated() uint16 {
	return 0
}

func (*PartyDetailsListRequest537NoRequestingPartyIDs) RequestingPartyRoleMetaAttribute(meta int) string {
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

func (*PartyDetailsListRequest537NoRequestingPartyIDs) RequestingPartyRoleMinValue() byte {
	return byte(32)
}

func (*PartyDetailsListRequest537NoRequestingPartyIDs) RequestingPartyRoleMaxValue() byte {
	return byte(126)
}

func (*PartyDetailsListRequest537NoRequestingPartyIDs) RequestingPartyRoleNullValue() byte {
	return 0
}

func (*PartyDetailsListRequest537NoPartyIDs) PartyIDId() uint16 {
	return 448
}

func (*PartyDetailsListRequest537NoPartyIDs) PartyIDSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsListRequest537NoPartyIDs) PartyIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.PartyIDSinceVersion()
}

func (*PartyDetailsListRequest537NoPartyIDs) PartyIDDeprecated() uint16 {
	return 0
}

func (*PartyDetailsListRequest537NoPartyIDs) PartyIDMetaAttribute(meta int) string {
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

func (*PartyDetailsListRequest537NoPartyIDs) PartyIDMinValue() uint64 {
	return 0
}

func (*PartyDetailsListRequest537NoPartyIDs) PartyIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*PartyDetailsListRequest537NoPartyIDs) PartyIDNullValue() uint64 {
	return math.MaxUint64
}

func (*PartyDetailsListRequest537NoPartyIDs) PartyIDSourceId() uint16 {
	return 447
}

func (*PartyDetailsListRequest537NoPartyIDs) PartyIDSourceSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsListRequest537NoPartyIDs) PartyIDSourceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.PartyIDSourceSinceVersion()
}

func (*PartyDetailsListRequest537NoPartyIDs) PartyIDSourceDeprecated() uint16 {
	return 0
}

func (*PartyDetailsListRequest537NoPartyIDs) PartyIDSourceMetaAttribute(meta int) string {
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

func (*PartyDetailsListRequest537NoPartyIDs) PartyIDSourceMinValue() byte {
	return byte(32)
}

func (*PartyDetailsListRequest537NoPartyIDs) PartyIDSourceMaxValue() byte {
	return byte(126)
}

func (*PartyDetailsListRequest537NoPartyIDs) PartyIDSourceNullValue() byte {
	return 0
}

func (*PartyDetailsListRequest537NoPartyIDs) PartyRoleId() uint16 {
	return 452
}

func (*PartyDetailsListRequest537NoPartyIDs) PartyRoleSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsListRequest537NoPartyIDs) PartyRoleInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.PartyRoleSinceVersion()
}

func (*PartyDetailsListRequest537NoPartyIDs) PartyRoleDeprecated() uint16 {
	return 0
}

func (*PartyDetailsListRequest537NoPartyIDs) PartyRoleMetaAttribute(meta int) string {
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

func (*PartyDetailsListRequest537NoPartyIDs) PartyRoleMinValue() uint16 {
	return 0
}

func (*PartyDetailsListRequest537NoPartyIDs) PartyRoleMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*PartyDetailsListRequest537NoPartyIDs) PartyRoleNullValue() uint16 {
	return math.MaxUint16
}

func (*PartyDetailsListRequest537) NoRequestingPartyIDsId() uint16 {
	return 1657
}

func (*PartyDetailsListRequest537) NoRequestingPartyIDsSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsListRequest537) NoRequestingPartyIDsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.NoRequestingPartyIDsSinceVersion()
}

func (*PartyDetailsListRequest537) NoRequestingPartyIDsDeprecated() uint16 {
	return 0
}

func (*PartyDetailsListRequest537NoRequestingPartyIDs) SbeBlockLength() (blockLength uint) {
	return 7
}

func (*PartyDetailsListRequest537NoRequestingPartyIDs) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}

func (*PartyDetailsListRequest537) NoPartyIDsId() uint16 {
	return 453
}

func (*PartyDetailsListRequest537) NoPartyIDsSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsListRequest537) NoPartyIDsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.NoPartyIDsSinceVersion()
}

func (*PartyDetailsListRequest537) NoPartyIDsDeprecated() uint16 {
	return 0
}

func (*PartyDetailsListRequest537NoPartyIDs) SbeBlockLength() (blockLength uint) {
	return 11
}

func (*PartyDetailsListRequest537NoPartyIDs) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}
