// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type PartyDetailsDefinitionRequest518 struct {
	PartyDetailsListReqID          uint64
	SendingTimeEpoch               uint64
	ListUpdateAction               ListUpdActEnum
	SeqNum                         uint32
	Memo                           [75]byte
	AvgPxGroupID                   [20]byte
	SelfMatchPreventionID          uint64
	CmtaGiveupCD                   CmtaGiveUpCDEnum
	CustOrderCapacity              CustOrderCapacityEnum
	ClearingAccountType            ClearingAcctTypeEnum
	SelfMatchPreventionInstruction SMPIEnum
	AvgPxIndicator                 AvgPxIndEnum
	ClearingTradePriceType         SLEDSEnum
	CustOrderHandlingInst          CustOrdHandlInstEnum
	Executor                       uint64
	IDMShortCode                   uint64
	NoPartyUpdates                 uint8
	NoPartyDetails                 []PartyDetailsDefinitionRequest518NoPartyDetails
	NoTrdRegPublications           []PartyDetailsDefinitionRequest518NoTrdRegPublications
}
type PartyDetailsDefinitionRequest518NoPartyDetails struct {
	PartyDetailID       [20]byte
	PartyDetailIDSource [1]byte
	PartyDetailRole     PartyDetailRoleEnum
}
type PartyDetailsDefinitionRequest518NoTrdRegPublications struct {
	TrdRegPublicationType   uint8
	TrdRegPublicationReason uint8
}

func (p *PartyDetailsDefinitionRequest518) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
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
	if err := p.ListUpdateAction.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, p.SeqNum); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, p.Memo[:]); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, p.AvgPxGroupID[:]); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, p.SelfMatchPreventionID); err != nil {
		return err
	}
	if err := p.CmtaGiveupCD.Encode(_m, _w); err != nil {
		return err
	}
	if err := p.CustOrderCapacity.Encode(_m, _w); err != nil {
		return err
	}
	if err := p.ClearingAccountType.Encode(_m, _w); err != nil {
		return err
	}
	if err := p.SelfMatchPreventionInstruction.Encode(_m, _w); err != nil {
		return err
	}
	if err := p.AvgPxIndicator.Encode(_m, _w); err != nil {
		return err
	}
	if err := p.ClearingTradePriceType.Encode(_m, _w); err != nil {
		return err
	}
	if err := p.CustOrderHandlingInst.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, p.Executor); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, p.IDMShortCode); err != nil {
		return err
	}
	var NoPartyDetailsBlockLength uint16 = 22
	if err := _m.WriteUint16(_w, NoPartyDetailsBlockLength); err != nil {
		return err
	}
	var NoPartyDetailsNumInGroup uint8 = uint8(len(p.NoPartyDetails))
	if err := _m.WriteUint8(_w, NoPartyDetailsNumInGroup); err != nil {
		return err
	}
	for _, prop := range p.NoPartyDetails {
		if err := prop.Encode(_m, _w); err != nil {
			return err
		}
	}
	var NoTrdRegPublicationsBlockLength uint16 = 2
	if err := _m.WriteUint16(_w, NoTrdRegPublicationsBlockLength); err != nil {
		return err
	}
	var NoTrdRegPublicationsNumInGroup uint8 = uint8(len(p.NoTrdRegPublications))
	if err := _m.WriteUint8(_w, NoTrdRegPublicationsNumInGroup); err != nil {
		return err
	}
	for _, prop := range p.NoTrdRegPublications {
		if err := prop.Encode(_m, _w); err != nil {
			return err
		}
	}
	return nil
}

func (p *PartyDetailsDefinitionRequest518) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
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
	if p.ListUpdateActionInActingVersion(actingVersion) {
		if err := p.ListUpdateAction.Decode(_m, _r, actingVersion); err != nil {
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
	if !p.MemoInActingVersion(actingVersion) {
		for idx := 0; idx < 75; idx++ {
			p.Memo[idx] = p.MemoNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, p.Memo[:]); err != nil {
			return err
		}
	}
	if !p.AvgPxGroupIDInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			p.AvgPxGroupID[idx] = p.AvgPxGroupIDNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, p.AvgPxGroupID[:]); err != nil {
			return err
		}
	}
	if !p.SelfMatchPreventionIDInActingVersion(actingVersion) {
		p.SelfMatchPreventionID = p.SelfMatchPreventionIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &p.SelfMatchPreventionID); err != nil {
			return err
		}
	}
	if p.CmtaGiveupCDInActingVersion(actingVersion) {
		if err := p.CmtaGiveupCD.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if p.CustOrderCapacityInActingVersion(actingVersion) {
		if err := p.CustOrderCapacity.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if p.ClearingAccountTypeInActingVersion(actingVersion) {
		if err := p.ClearingAccountType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if p.SelfMatchPreventionInstructionInActingVersion(actingVersion) {
		if err := p.SelfMatchPreventionInstruction.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if p.AvgPxIndicatorInActingVersion(actingVersion) {
		if err := p.AvgPxIndicator.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if p.ClearingTradePriceTypeInActingVersion(actingVersion) {
		if err := p.ClearingTradePriceType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if p.CustOrderHandlingInstInActingVersion(actingVersion) {
		if err := p.CustOrderHandlingInst.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !p.ExecutorInActingVersion(actingVersion) {
		p.Executor = p.ExecutorNullValue()
	} else {
		if err := _m.ReadUint64(_r, &p.Executor); err != nil {
			return err
		}
	}
	if !p.IDMShortCodeInActingVersion(actingVersion) {
		p.IDMShortCode = p.IDMShortCodeNullValue()
	} else {
		if err := _m.ReadUint64(_r, &p.IDMShortCode); err != nil {
			return err
		}
	}
	p.NoPartyUpdates = 1
	if actingVersion > p.SbeSchemaVersion() && blockLength > p.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-p.SbeBlockLength()))
	}

	if p.NoPartyDetailsInActingVersion(actingVersion) {
		var NoPartyDetailsBlockLength uint16
		if err := _m.ReadUint16(_r, &NoPartyDetailsBlockLength); err != nil {
			return err
		}
		var NoPartyDetailsNumInGroup uint8
		if err := _m.ReadUint8(_r, &NoPartyDetailsNumInGroup); err != nil {
			return err
		}
		if cap(p.NoPartyDetails) < int(NoPartyDetailsNumInGroup) {
			p.NoPartyDetails = make([]PartyDetailsDefinitionRequest518NoPartyDetails, NoPartyDetailsNumInGroup)
		}
		for i, _ := range p.NoPartyDetails {
			if err := p.NoPartyDetails[i].Decode(_m, _r, actingVersion, uint(NoPartyDetailsBlockLength)); err != nil {
				return err
			}
		}
	}

	if p.NoTrdRegPublicationsInActingVersion(actingVersion) {
		var NoTrdRegPublicationsBlockLength uint16
		if err := _m.ReadUint16(_r, &NoTrdRegPublicationsBlockLength); err != nil {
			return err
		}
		var NoTrdRegPublicationsNumInGroup uint8
		if err := _m.ReadUint8(_r, &NoTrdRegPublicationsNumInGroup); err != nil {
			return err
		}
		if cap(p.NoTrdRegPublications) < int(NoTrdRegPublicationsNumInGroup) {
			p.NoTrdRegPublications = make([]PartyDetailsDefinitionRequest518NoTrdRegPublications, NoTrdRegPublicationsNumInGroup)
		}
		for i, _ := range p.NoTrdRegPublications {
			if err := p.NoTrdRegPublications[i].Decode(_m, _r, actingVersion, uint(NoTrdRegPublicationsBlockLength)); err != nil {
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

func (p *PartyDetailsDefinitionRequest518) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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
	if err := p.ListUpdateAction.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if p.SeqNumInActingVersion(actingVersion) {
		if p.SeqNum < p.SeqNumMinValue() || p.SeqNum > p.SeqNumMaxValue() {
			return fmt.Errorf("Range check failed on p.SeqNum (%v < %v > %v)", p.SeqNumMinValue(), p.SeqNum, p.SeqNumMaxValue())
		}
	}
	if p.MemoInActingVersion(actingVersion) {
		for idx := 0; idx < 75; idx++ {
			if p.Memo[idx] < p.MemoMinValue() || p.Memo[idx] > p.MemoMaxValue() {
				return fmt.Errorf("Range check failed on p.Memo[%d] (%v < %v > %v)", idx, p.MemoMinValue(), p.Memo[idx], p.MemoMaxValue())
			}
		}
	}
	if p.AvgPxGroupIDInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			if p.AvgPxGroupID[idx] < p.AvgPxGroupIDMinValue() || p.AvgPxGroupID[idx] > p.AvgPxGroupIDMaxValue() {
				return fmt.Errorf("Range check failed on p.AvgPxGroupID[%d] (%v < %v > %v)", idx, p.AvgPxGroupIDMinValue(), p.AvgPxGroupID[idx], p.AvgPxGroupIDMaxValue())
			}
		}
	}
	if p.SelfMatchPreventionIDInActingVersion(actingVersion) {
		if p.SelfMatchPreventionID != p.SelfMatchPreventionIDNullValue() && (p.SelfMatchPreventionID < p.SelfMatchPreventionIDMinValue() || p.SelfMatchPreventionID > p.SelfMatchPreventionIDMaxValue()) {
			return fmt.Errorf("Range check failed on p.SelfMatchPreventionID (%v < %v > %v)", p.SelfMatchPreventionIDMinValue(), p.SelfMatchPreventionID, p.SelfMatchPreventionIDMaxValue())
		}
	}
	if err := p.CmtaGiveupCD.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := p.CustOrderCapacity.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := p.ClearingAccountType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := p.SelfMatchPreventionInstruction.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := p.AvgPxIndicator.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := p.ClearingTradePriceType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := p.CustOrderHandlingInst.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if p.ExecutorInActingVersion(actingVersion) {
		if p.Executor != p.ExecutorNullValue() && (p.Executor < p.ExecutorMinValue() || p.Executor > p.ExecutorMaxValue()) {
			return fmt.Errorf("Range check failed on p.Executor (%v < %v > %v)", p.ExecutorMinValue(), p.Executor, p.ExecutorMaxValue())
		}
	}
	if p.IDMShortCodeInActingVersion(actingVersion) {
		if p.IDMShortCode != p.IDMShortCodeNullValue() && (p.IDMShortCode < p.IDMShortCodeMinValue() || p.IDMShortCode > p.IDMShortCodeMaxValue()) {
			return fmt.Errorf("Range check failed on p.IDMShortCode (%v < %v > %v)", p.IDMShortCodeMinValue(), p.IDMShortCode, p.IDMShortCodeMaxValue())
		}
	}
	for _, prop := range p.NoPartyDetails {
		if err := prop.RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	for _, prop := range p.NoTrdRegPublications {
		if err := prop.RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	return nil
}

func PartyDetailsDefinitionRequest518Init(p *PartyDetailsDefinitionRequest518) {
	for idx := 0; idx < 75; idx++ {
		p.Memo[idx] = 0
	}
	for idx := 0; idx < 20; idx++ {
		p.AvgPxGroupID[idx] = 0
	}
	p.SelfMatchPreventionID = 18446744073709551615
	p.Executor = 18446744073709551615
	p.IDMShortCode = 18446744073709551615
	p.NoPartyUpdates = 1
	return
}

func (p *PartyDetailsDefinitionRequest518NoPartyDetails) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteBytes(_w, p.PartyDetailID[:]); err != nil {
		return err
	}
	if err := p.PartyDetailRole.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (p *PartyDetailsDefinitionRequest518NoPartyDetails) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	if !p.PartyDetailIDInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			p.PartyDetailID[idx] = p.PartyDetailIDNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, p.PartyDetailID[:]); err != nil {
			return err
		}
	}
	p.PartyDetailIDSource[0] = 67
	if p.PartyDetailRoleInActingVersion(actingVersion) {
		if err := p.PartyDetailRole.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > p.SbeSchemaVersion() && blockLength > p.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-p.SbeBlockLength()))
	}
	return nil
}

func (p *PartyDetailsDefinitionRequest518NoPartyDetails) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if p.PartyDetailIDInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			if p.PartyDetailID[idx] < p.PartyDetailIDMinValue() || p.PartyDetailID[idx] > p.PartyDetailIDMaxValue() {
				return fmt.Errorf("Range check failed on p.PartyDetailID[%d] (%v < %v > %v)", idx, p.PartyDetailIDMinValue(), p.PartyDetailID[idx], p.PartyDetailIDMaxValue())
			}
		}
	}
	if err := p.PartyDetailRole.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	return nil
}

func PartyDetailsDefinitionRequest518NoPartyDetailsInit(p *PartyDetailsDefinitionRequest518NoPartyDetails) {
	p.PartyDetailIDSource[0] = 67
	return
}

func (p *PartyDetailsDefinitionRequest518NoTrdRegPublications) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, p.TrdRegPublicationType); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, p.TrdRegPublicationReason); err != nil {
		return err
	}
	return nil
}

func (p *PartyDetailsDefinitionRequest518NoTrdRegPublications) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	if !p.TrdRegPublicationTypeInActingVersion(actingVersion) {
		p.TrdRegPublicationType = p.TrdRegPublicationTypeNullValue()
	} else {
		if err := _m.ReadUint8(_r, &p.TrdRegPublicationType); err != nil {
			return err
		}
	}
	if !p.TrdRegPublicationReasonInActingVersion(actingVersion) {
		p.TrdRegPublicationReason = p.TrdRegPublicationReasonNullValue()
	} else {
		if err := _m.ReadUint8(_r, &p.TrdRegPublicationReason); err != nil {
			return err
		}
	}
	if actingVersion > p.SbeSchemaVersion() && blockLength > p.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-p.SbeBlockLength()))
	}
	return nil
}

func (p *PartyDetailsDefinitionRequest518NoTrdRegPublications) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if p.TrdRegPublicationTypeInActingVersion(actingVersion) {
		if p.TrdRegPublicationType < p.TrdRegPublicationTypeMinValue() || p.TrdRegPublicationType > p.TrdRegPublicationTypeMaxValue() {
			return fmt.Errorf("Range check failed on p.TrdRegPublicationType (%v < %v > %v)", p.TrdRegPublicationTypeMinValue(), p.TrdRegPublicationType, p.TrdRegPublicationTypeMaxValue())
		}
	}
	if p.TrdRegPublicationReasonInActingVersion(actingVersion) {
		if p.TrdRegPublicationReason < p.TrdRegPublicationReasonMinValue() || p.TrdRegPublicationReason > p.TrdRegPublicationReasonMaxValue() {
			return fmt.Errorf("Range check failed on p.TrdRegPublicationReason (%v < %v > %v)", p.TrdRegPublicationReasonMinValue(), p.TrdRegPublicationReason, p.TrdRegPublicationReasonMaxValue())
		}
	}
	return nil
}

func PartyDetailsDefinitionRequest518NoTrdRegPublicationsInit(p *PartyDetailsDefinitionRequest518NoTrdRegPublications) {
	return
}

func (*PartyDetailsDefinitionRequest518) SbeBlockLength() (blockLength uint16) {
	return 147
}

func (*PartyDetailsDefinitionRequest518) SbeTemplateId() (templateId uint16) {
	return 518
}

func (*PartyDetailsDefinitionRequest518) SbeSchemaId() (schemaId uint16) {
	return 8
}

func (*PartyDetailsDefinitionRequest518) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}

func (*PartyDetailsDefinitionRequest518) SbeSemanticType() (semanticType []byte) {
	return []byte("CX")
}

func (*PartyDetailsDefinitionRequest518) PartyDetailsListReqIDId() uint16 {
	return 1505
}

func (*PartyDetailsDefinitionRequest518) PartyDetailsListReqIDSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsDefinitionRequest518) PartyDetailsListReqIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.PartyDetailsListReqIDSinceVersion()
}

func (*PartyDetailsDefinitionRequest518) PartyDetailsListReqIDDeprecated() uint16 {
	return 0
}

func (*PartyDetailsDefinitionRequest518) PartyDetailsListReqIDMetaAttribute(meta int) string {
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

func (*PartyDetailsDefinitionRequest518) PartyDetailsListReqIDMinValue() uint64 {
	return 0
}

func (*PartyDetailsDefinitionRequest518) PartyDetailsListReqIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*PartyDetailsDefinitionRequest518) PartyDetailsListReqIDNullValue() uint64 {
	return math.MaxUint64
}

func (*PartyDetailsDefinitionRequest518) SendingTimeEpochId() uint16 {
	return 5297
}

func (*PartyDetailsDefinitionRequest518) SendingTimeEpochSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsDefinitionRequest518) SendingTimeEpochInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.SendingTimeEpochSinceVersion()
}

func (*PartyDetailsDefinitionRequest518) SendingTimeEpochDeprecated() uint16 {
	return 0
}

func (*PartyDetailsDefinitionRequest518) SendingTimeEpochMetaAttribute(meta int) string {
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

func (*PartyDetailsDefinitionRequest518) SendingTimeEpochMinValue() uint64 {
	return 0
}

func (*PartyDetailsDefinitionRequest518) SendingTimeEpochMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*PartyDetailsDefinitionRequest518) SendingTimeEpochNullValue() uint64 {
	return math.MaxUint64
}

func (*PartyDetailsDefinitionRequest518) ListUpdateActionId() uint16 {
	return 1324
}

func (*PartyDetailsDefinitionRequest518) ListUpdateActionSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsDefinitionRequest518) ListUpdateActionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.ListUpdateActionSinceVersion()
}

func (*PartyDetailsDefinitionRequest518) ListUpdateActionDeprecated() uint16 {
	return 0
}

func (*PartyDetailsDefinitionRequest518) ListUpdateActionMetaAttribute(meta int) string {
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

func (*PartyDetailsDefinitionRequest518) SeqNumId() uint16 {
	return 9726
}

func (*PartyDetailsDefinitionRequest518) SeqNumSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsDefinitionRequest518) SeqNumInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.SeqNumSinceVersion()
}

func (*PartyDetailsDefinitionRequest518) SeqNumDeprecated() uint16 {
	return 0
}

func (*PartyDetailsDefinitionRequest518) SeqNumMetaAttribute(meta int) string {
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

func (*PartyDetailsDefinitionRequest518) SeqNumMinValue() uint32 {
	return 0
}

func (*PartyDetailsDefinitionRequest518) SeqNumMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*PartyDetailsDefinitionRequest518) SeqNumNullValue() uint32 {
	return math.MaxUint32
}

func (*PartyDetailsDefinitionRequest518) MemoId() uint16 {
	return 5149
}

func (*PartyDetailsDefinitionRequest518) MemoSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsDefinitionRequest518) MemoInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.MemoSinceVersion()
}

func (*PartyDetailsDefinitionRequest518) MemoDeprecated() uint16 {
	return 0
}

func (*PartyDetailsDefinitionRequest518) MemoMetaAttribute(meta int) string {
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

func (*PartyDetailsDefinitionRequest518) MemoMinValue() byte {
	return byte(32)
}

func (*PartyDetailsDefinitionRequest518) MemoMaxValue() byte {
	return byte(126)
}

func (*PartyDetailsDefinitionRequest518) MemoNullValue() byte {
	return 0
}

func (p *PartyDetailsDefinitionRequest518) MemoCharacterEncoding() string {
	return "US-ASCII"
}

func (*PartyDetailsDefinitionRequest518) AvgPxGroupIDId() uint16 {
	return 1731
}

func (*PartyDetailsDefinitionRequest518) AvgPxGroupIDSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsDefinitionRequest518) AvgPxGroupIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.AvgPxGroupIDSinceVersion()
}

func (*PartyDetailsDefinitionRequest518) AvgPxGroupIDDeprecated() uint16 {
	return 0
}

func (*PartyDetailsDefinitionRequest518) AvgPxGroupIDMetaAttribute(meta int) string {
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

func (*PartyDetailsDefinitionRequest518) AvgPxGroupIDMinValue() byte {
	return byte(32)
}

func (*PartyDetailsDefinitionRequest518) AvgPxGroupIDMaxValue() byte {
	return byte(126)
}

func (*PartyDetailsDefinitionRequest518) AvgPxGroupIDNullValue() byte {
	return 0
}

func (p *PartyDetailsDefinitionRequest518) AvgPxGroupIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*PartyDetailsDefinitionRequest518) SelfMatchPreventionIDId() uint16 {
	return 2362
}

func (*PartyDetailsDefinitionRequest518) SelfMatchPreventionIDSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsDefinitionRequest518) SelfMatchPreventionIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.SelfMatchPreventionIDSinceVersion()
}

func (*PartyDetailsDefinitionRequest518) SelfMatchPreventionIDDeprecated() uint16 {
	return 0
}

func (*PartyDetailsDefinitionRequest518) SelfMatchPreventionIDMetaAttribute(meta int) string {
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

func (*PartyDetailsDefinitionRequest518) SelfMatchPreventionIDMinValue() uint64 {
	return 0
}

func (*PartyDetailsDefinitionRequest518) SelfMatchPreventionIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*PartyDetailsDefinitionRequest518) SelfMatchPreventionIDNullValue() uint64 {
	return 18446744073709551615
}

func (*PartyDetailsDefinitionRequest518) CmtaGiveupCDId() uint16 {
	return 9708
}

func (*PartyDetailsDefinitionRequest518) CmtaGiveupCDSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsDefinitionRequest518) CmtaGiveupCDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.CmtaGiveupCDSinceVersion()
}

func (*PartyDetailsDefinitionRequest518) CmtaGiveupCDDeprecated() uint16 {
	return 0
}

func (*PartyDetailsDefinitionRequest518) CmtaGiveupCDMetaAttribute(meta int) string {
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

func (*PartyDetailsDefinitionRequest518) CustOrderCapacityId() uint16 {
	return 582
}

func (*PartyDetailsDefinitionRequest518) CustOrderCapacitySinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsDefinitionRequest518) CustOrderCapacityInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.CustOrderCapacitySinceVersion()
}

func (*PartyDetailsDefinitionRequest518) CustOrderCapacityDeprecated() uint16 {
	return 0
}

func (*PartyDetailsDefinitionRequest518) CustOrderCapacityMetaAttribute(meta int) string {
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

func (*PartyDetailsDefinitionRequest518) ClearingAccountTypeId() uint16 {
	return 1816
}

func (*PartyDetailsDefinitionRequest518) ClearingAccountTypeSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsDefinitionRequest518) ClearingAccountTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.ClearingAccountTypeSinceVersion()
}

func (*PartyDetailsDefinitionRequest518) ClearingAccountTypeDeprecated() uint16 {
	return 0
}

func (*PartyDetailsDefinitionRequest518) ClearingAccountTypeMetaAttribute(meta int) string {
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

func (*PartyDetailsDefinitionRequest518) SelfMatchPreventionInstructionId() uint16 {
	return 8000
}

func (*PartyDetailsDefinitionRequest518) SelfMatchPreventionInstructionSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsDefinitionRequest518) SelfMatchPreventionInstructionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.SelfMatchPreventionInstructionSinceVersion()
}

func (*PartyDetailsDefinitionRequest518) SelfMatchPreventionInstructionDeprecated() uint16 {
	return 0
}

func (*PartyDetailsDefinitionRequest518) SelfMatchPreventionInstructionMetaAttribute(meta int) string {
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

func (*PartyDetailsDefinitionRequest518) AvgPxIndicatorId() uint16 {
	return 819
}

func (*PartyDetailsDefinitionRequest518) AvgPxIndicatorSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsDefinitionRequest518) AvgPxIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.AvgPxIndicatorSinceVersion()
}

func (*PartyDetailsDefinitionRequest518) AvgPxIndicatorDeprecated() uint16 {
	return 0
}

func (*PartyDetailsDefinitionRequest518) AvgPxIndicatorMetaAttribute(meta int) string {
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

func (*PartyDetailsDefinitionRequest518) ClearingTradePriceTypeId() uint16 {
	return 1598
}

func (*PartyDetailsDefinitionRequest518) ClearingTradePriceTypeSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsDefinitionRequest518) ClearingTradePriceTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.ClearingTradePriceTypeSinceVersion()
}

func (*PartyDetailsDefinitionRequest518) ClearingTradePriceTypeDeprecated() uint16 {
	return 0
}

func (*PartyDetailsDefinitionRequest518) ClearingTradePriceTypeMetaAttribute(meta int) string {
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

func (*PartyDetailsDefinitionRequest518) CustOrderHandlingInstId() uint16 {
	return 1031
}

func (*PartyDetailsDefinitionRequest518) CustOrderHandlingInstSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsDefinitionRequest518) CustOrderHandlingInstInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.CustOrderHandlingInstSinceVersion()
}

func (*PartyDetailsDefinitionRequest518) CustOrderHandlingInstDeprecated() uint16 {
	return 0
}

func (*PartyDetailsDefinitionRequest518) CustOrderHandlingInstMetaAttribute(meta int) string {
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

func (*PartyDetailsDefinitionRequest518) ExecutorId() uint16 {
	return 5290
}

func (*PartyDetailsDefinitionRequest518) ExecutorSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsDefinitionRequest518) ExecutorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.ExecutorSinceVersion()
}

func (*PartyDetailsDefinitionRequest518) ExecutorDeprecated() uint16 {
	return 0
}

func (*PartyDetailsDefinitionRequest518) ExecutorMetaAttribute(meta int) string {
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

func (*PartyDetailsDefinitionRequest518) ExecutorMinValue() uint64 {
	return 0
}

func (*PartyDetailsDefinitionRequest518) ExecutorMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*PartyDetailsDefinitionRequest518) ExecutorNullValue() uint64 {
	return 18446744073709551615
}

func (*PartyDetailsDefinitionRequest518) IDMShortCodeId() uint16 {
	return 36023
}

func (*PartyDetailsDefinitionRequest518) IDMShortCodeSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsDefinitionRequest518) IDMShortCodeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.IDMShortCodeSinceVersion()
}

func (*PartyDetailsDefinitionRequest518) IDMShortCodeDeprecated() uint16 {
	return 0
}

func (*PartyDetailsDefinitionRequest518) IDMShortCodeMetaAttribute(meta int) string {
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

func (*PartyDetailsDefinitionRequest518) IDMShortCodeMinValue() uint64 {
	return 0
}

func (*PartyDetailsDefinitionRequest518) IDMShortCodeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*PartyDetailsDefinitionRequest518) IDMShortCodeNullValue() uint64 {
	return 18446744073709551615
}

func (*PartyDetailsDefinitionRequest518) NoPartyUpdatesId() uint16 {
	return 1676
}

func (*PartyDetailsDefinitionRequest518) NoPartyUpdatesSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsDefinitionRequest518) NoPartyUpdatesInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.NoPartyUpdatesSinceVersion()
}

func (*PartyDetailsDefinitionRequest518) NoPartyUpdatesDeprecated() uint16 {
	return 0
}

func (*PartyDetailsDefinitionRequest518) NoPartyUpdatesMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "int"
	case 4:
		return "constant"
	}
	return ""
}

func (*PartyDetailsDefinitionRequest518) NoPartyUpdatesMinValue() uint8 {
	return 0
}

func (*PartyDetailsDefinitionRequest518) NoPartyUpdatesMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*PartyDetailsDefinitionRequest518) NoPartyUpdatesNullValue() uint8 {
	return math.MaxUint8
}

func (*PartyDetailsDefinitionRequest518NoPartyDetails) PartyDetailIDId() uint16 {
	return 1691
}

func (*PartyDetailsDefinitionRequest518NoPartyDetails) PartyDetailIDSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsDefinitionRequest518NoPartyDetails) PartyDetailIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.PartyDetailIDSinceVersion()
}

func (*PartyDetailsDefinitionRequest518NoPartyDetails) PartyDetailIDDeprecated() uint16 {
	return 0
}

func (*PartyDetailsDefinitionRequest518NoPartyDetails) PartyDetailIDMetaAttribute(meta int) string {
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

func (*PartyDetailsDefinitionRequest518NoPartyDetails) PartyDetailIDMinValue() byte {
	return byte(32)
}

func (*PartyDetailsDefinitionRequest518NoPartyDetails) PartyDetailIDMaxValue() byte {
	return byte(126)
}

func (*PartyDetailsDefinitionRequest518NoPartyDetails) PartyDetailIDNullValue() byte {
	return 0
}

func (p *PartyDetailsDefinitionRequest518NoPartyDetails) PartyDetailIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*PartyDetailsDefinitionRequest518NoPartyDetails) PartyDetailIDSourceId() uint16 {
	return 1692
}

func (*PartyDetailsDefinitionRequest518NoPartyDetails) PartyDetailIDSourceSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsDefinitionRequest518NoPartyDetails) PartyDetailIDSourceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.PartyDetailIDSourceSinceVersion()
}

func (*PartyDetailsDefinitionRequest518NoPartyDetails) PartyDetailIDSourceDeprecated() uint16 {
	return 0
}

func (*PartyDetailsDefinitionRequest518NoPartyDetails) PartyDetailIDSourceMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "char"
	case 4:
		return "constant"
	}
	return ""
}

func (*PartyDetailsDefinitionRequest518NoPartyDetails) PartyDetailIDSourceMinValue() byte {
	return byte(32)
}

func (*PartyDetailsDefinitionRequest518NoPartyDetails) PartyDetailIDSourceMaxValue() byte {
	return byte(126)
}

func (*PartyDetailsDefinitionRequest518NoPartyDetails) PartyDetailIDSourceNullValue() byte {
	return 0
}

func (*PartyDetailsDefinitionRequest518NoPartyDetails) PartyDetailRoleId() uint16 {
	return 1693
}

func (*PartyDetailsDefinitionRequest518NoPartyDetails) PartyDetailRoleSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsDefinitionRequest518NoPartyDetails) PartyDetailRoleInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.PartyDetailRoleSinceVersion()
}

func (*PartyDetailsDefinitionRequest518NoPartyDetails) PartyDetailRoleDeprecated() uint16 {
	return 0
}

func (*PartyDetailsDefinitionRequest518NoPartyDetails) PartyDetailRoleMetaAttribute(meta int) string {
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

func (*PartyDetailsDefinitionRequest518NoTrdRegPublications) TrdRegPublicationTypeId() uint16 {
	return 2669
}

func (*PartyDetailsDefinitionRequest518NoTrdRegPublications) TrdRegPublicationTypeSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsDefinitionRequest518NoTrdRegPublications) TrdRegPublicationTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.TrdRegPublicationTypeSinceVersion()
}

func (*PartyDetailsDefinitionRequest518NoTrdRegPublications) TrdRegPublicationTypeDeprecated() uint16 {
	return 0
}

func (*PartyDetailsDefinitionRequest518NoTrdRegPublications) TrdRegPublicationTypeMetaAttribute(meta int) string {
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

func (*PartyDetailsDefinitionRequest518NoTrdRegPublications) TrdRegPublicationTypeMinValue() uint8 {
	return 0
}

func (*PartyDetailsDefinitionRequest518NoTrdRegPublications) TrdRegPublicationTypeMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*PartyDetailsDefinitionRequest518NoTrdRegPublications) TrdRegPublicationTypeNullValue() uint8 {
	return math.MaxUint8
}

func (*PartyDetailsDefinitionRequest518NoTrdRegPublications) TrdRegPublicationReasonId() uint16 {
	return 2670
}

func (*PartyDetailsDefinitionRequest518NoTrdRegPublications) TrdRegPublicationReasonSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsDefinitionRequest518NoTrdRegPublications) TrdRegPublicationReasonInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.TrdRegPublicationReasonSinceVersion()
}

func (*PartyDetailsDefinitionRequest518NoTrdRegPublications) TrdRegPublicationReasonDeprecated() uint16 {
	return 0
}

func (*PartyDetailsDefinitionRequest518NoTrdRegPublications) TrdRegPublicationReasonMetaAttribute(meta int) string {
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

func (*PartyDetailsDefinitionRequest518NoTrdRegPublications) TrdRegPublicationReasonMinValue() uint8 {
	return 0
}

func (*PartyDetailsDefinitionRequest518NoTrdRegPublications) TrdRegPublicationReasonMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*PartyDetailsDefinitionRequest518NoTrdRegPublications) TrdRegPublicationReasonNullValue() uint8 {
	return math.MaxUint8
}

func (*PartyDetailsDefinitionRequest518) NoPartyDetailsId() uint16 {
	return 1671
}

func (*PartyDetailsDefinitionRequest518) NoPartyDetailsSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsDefinitionRequest518) NoPartyDetailsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.NoPartyDetailsSinceVersion()
}

func (*PartyDetailsDefinitionRequest518) NoPartyDetailsDeprecated() uint16 {
	return 0
}

func (*PartyDetailsDefinitionRequest518NoPartyDetails) SbeBlockLength() (blockLength uint) {
	return 22
}

func (*PartyDetailsDefinitionRequest518NoPartyDetails) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}

func (*PartyDetailsDefinitionRequest518) NoTrdRegPublicationsId() uint16 {
	return 2668
}

func (*PartyDetailsDefinitionRequest518) NoTrdRegPublicationsSinceVersion() uint16 {
	return 2
}

func (p *PartyDetailsDefinitionRequest518) NoTrdRegPublicationsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.NoTrdRegPublicationsSinceVersion()
}

func (*PartyDetailsDefinitionRequest518) NoTrdRegPublicationsDeprecated() uint16 {
	return 0
}

func (*PartyDetailsDefinitionRequest518NoTrdRegPublications) SbeBlockLength() (blockLength uint) {
	return 2
}

func (*PartyDetailsDefinitionRequest518NoTrdRegPublications) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}
