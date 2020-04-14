// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type PartyDetailsDefinitionRequestAck519 struct {
	SeqNum                         uint32
	UUID                           uint64
	Memo                           [75]byte
	AvgPxGroupID                   [20]byte
	PartyDetailsListReqID          uint64
	SendingTimeEpoch               uint64
	SelfMatchPreventionID          uint64
	PartyDetailRequestStatus       uint8
	CustOrderCapacity              CustOrderCapacityEnum
	ClearingAccountType            ClearingAcctTypeEnum
	SelfMatchPreventionInstruction SMPIEnum
	AvgPxIndicator                 AvgPxIndEnum
	ClearingTradePriceType         SLEDSEnum
	CmtaGiveupCD                   CmtaGiveUpCDEnum
	CustOrderHandlingInst          CustOrdHandlInstEnum
	NoPartyUpdates                 uint8
	ListUpdateAction               ListUpdActEnum
	PartyDetailDefinitionStatus    uint8
	Executor                       uint64
	IDMShortCode                   uint64
	PossRetransFlag                BooleanFlagEnum
	SplitMsg                       SplitMsgEnum
	NoPartyDetails                 []PartyDetailsDefinitionRequestAck519NoPartyDetails
	NoTrdRegPublications           []PartyDetailsDefinitionRequestAck519NoTrdRegPublications
}
type PartyDetailsDefinitionRequestAck519NoPartyDetails struct {
	PartyDetailID       [20]byte
	PartyDetailIDSource [1]byte
	PartyDetailRole     PartyDetailRoleEnum
}
type PartyDetailsDefinitionRequestAck519NoTrdRegPublications struct {
	TrdRegPublicationType   uint8
	TrdRegPublicationReason uint8
}

func (p *PartyDetailsDefinitionRequestAck519) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := p.RangeCheck(p.SbeSchemaVersion(), p.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteUint32(_w, p.SeqNum); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, p.UUID); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, p.Memo[:]); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, p.AvgPxGroupID[:]); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, p.PartyDetailsListReqID); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, p.SendingTimeEpoch); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, p.SelfMatchPreventionID); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, p.PartyDetailRequestStatus); err != nil {
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
	if err := p.CmtaGiveupCD.Encode(_m, _w); err != nil {
		return err
	}
	if err := p.CustOrderHandlingInst.Encode(_m, _w); err != nil {
		return err
	}
	if err := p.ListUpdateAction.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, p.PartyDetailDefinitionStatus); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, p.Executor); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, p.IDMShortCode); err != nil {
		return err
	}
	if err := p.PossRetransFlag.Encode(_m, _w); err != nil {
		return err
	}
	if err := p.SplitMsg.Encode(_m, _w); err != nil {
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

func (p *PartyDetailsDefinitionRequestAck519) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !p.SeqNumInActingVersion(actingVersion) {
		p.SeqNum = p.SeqNumNullValue()
	} else {
		if err := _m.ReadUint32(_r, &p.SeqNum); err != nil {
			return err
		}
	}
	if !p.UUIDInActingVersion(actingVersion) {
		p.UUID = p.UUIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &p.UUID); err != nil {
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
	if !p.SelfMatchPreventionIDInActingVersion(actingVersion) {
		p.SelfMatchPreventionID = p.SelfMatchPreventionIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &p.SelfMatchPreventionID); err != nil {
			return err
		}
	}
	if !p.PartyDetailRequestStatusInActingVersion(actingVersion) {
		p.PartyDetailRequestStatus = p.PartyDetailRequestStatusNullValue()
	} else {
		if err := _m.ReadUint8(_r, &p.PartyDetailRequestStatus); err != nil {
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
	if p.CmtaGiveupCDInActingVersion(actingVersion) {
		if err := p.CmtaGiveupCD.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if p.CustOrderHandlingInstInActingVersion(actingVersion) {
		if err := p.CustOrderHandlingInst.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	p.NoPartyUpdates = 1
	if p.ListUpdateActionInActingVersion(actingVersion) {
		if err := p.ListUpdateAction.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !p.PartyDetailDefinitionStatusInActingVersion(actingVersion) {
		p.PartyDetailDefinitionStatus = p.PartyDetailDefinitionStatusNullValue()
	} else {
		if err := _m.ReadUint8(_r, &p.PartyDetailDefinitionStatus); err != nil {
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
	if p.PossRetransFlagInActingVersion(actingVersion) {
		if err := p.PossRetransFlag.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if p.SplitMsgInActingVersion(actingVersion) {
		if err := p.SplitMsg.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
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
			p.NoPartyDetails = make([]PartyDetailsDefinitionRequestAck519NoPartyDetails, NoPartyDetailsNumInGroup)
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
			p.NoTrdRegPublications = make([]PartyDetailsDefinitionRequestAck519NoTrdRegPublications, NoTrdRegPublicationsNumInGroup)
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

func (p *PartyDetailsDefinitionRequestAck519) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if p.SeqNumInActingVersion(actingVersion) {
		if p.SeqNum < p.SeqNumMinValue() || p.SeqNum > p.SeqNumMaxValue() {
			return fmt.Errorf("Range check failed on p.SeqNum (%v < %v > %v)", p.SeqNumMinValue(), p.SeqNum, p.SeqNumMaxValue())
		}
	}
	if p.UUIDInActingVersion(actingVersion) {
		if p.UUID < p.UUIDMinValue() || p.UUID > p.UUIDMaxValue() {
			return fmt.Errorf("Range check failed on p.UUID (%v < %v > %v)", p.UUIDMinValue(), p.UUID, p.UUIDMaxValue())
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
	if p.SelfMatchPreventionIDInActingVersion(actingVersion) {
		if p.SelfMatchPreventionID != p.SelfMatchPreventionIDNullValue() && (p.SelfMatchPreventionID < p.SelfMatchPreventionIDMinValue() || p.SelfMatchPreventionID > p.SelfMatchPreventionIDMaxValue()) {
			return fmt.Errorf("Range check failed on p.SelfMatchPreventionID (%v < %v > %v)", p.SelfMatchPreventionIDMinValue(), p.SelfMatchPreventionID, p.SelfMatchPreventionIDMaxValue())
		}
	}
	if p.PartyDetailRequestStatusInActingVersion(actingVersion) {
		if p.PartyDetailRequestStatus < p.PartyDetailRequestStatusMinValue() || p.PartyDetailRequestStatus > p.PartyDetailRequestStatusMaxValue() {
			return fmt.Errorf("Range check failed on p.PartyDetailRequestStatus (%v < %v > %v)", p.PartyDetailRequestStatusMinValue(), p.PartyDetailRequestStatus, p.PartyDetailRequestStatusMaxValue())
		}
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
	if err := p.CmtaGiveupCD.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := p.CustOrderHandlingInst.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := p.ListUpdateAction.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if p.PartyDetailDefinitionStatusInActingVersion(actingVersion) {
		if p.PartyDetailDefinitionStatus < p.PartyDetailDefinitionStatusMinValue() || p.PartyDetailDefinitionStatus > p.PartyDetailDefinitionStatusMaxValue() {
			return fmt.Errorf("Range check failed on p.PartyDetailDefinitionStatus (%v < %v > %v)", p.PartyDetailDefinitionStatusMinValue(), p.PartyDetailDefinitionStatus, p.PartyDetailDefinitionStatusMaxValue())
		}
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
	if err := p.PossRetransFlag.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := p.SplitMsg.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
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

func PartyDetailsDefinitionRequestAck519Init(p *PartyDetailsDefinitionRequestAck519) {
	for idx := 0; idx < 75; idx++ {
		p.Memo[idx] = 0
	}
	for idx := 0; idx < 20; idx++ {
		p.AvgPxGroupID[idx] = 0
	}
	p.SelfMatchPreventionID = 18446744073709551615
	p.NoPartyUpdates = 1
	p.Executor = 18446744073709551615
	p.IDMShortCode = 18446744073709551615
	return
}

func (p *PartyDetailsDefinitionRequestAck519NoPartyDetails) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteBytes(_w, p.PartyDetailID[:]); err != nil {
		return err
	}
	if err := p.PartyDetailRole.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (p *PartyDetailsDefinitionRequestAck519NoPartyDetails) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
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

func (p *PartyDetailsDefinitionRequestAck519NoPartyDetails) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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

func PartyDetailsDefinitionRequestAck519NoPartyDetailsInit(p *PartyDetailsDefinitionRequestAck519NoPartyDetails) {
	p.PartyDetailIDSource[0] = 67
	return
}

func (p *PartyDetailsDefinitionRequestAck519NoTrdRegPublications) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, p.TrdRegPublicationType); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, p.TrdRegPublicationReason); err != nil {
		return err
	}
	return nil
}

func (p *PartyDetailsDefinitionRequestAck519NoTrdRegPublications) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
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

func (p *PartyDetailsDefinitionRequestAck519NoTrdRegPublications) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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

func PartyDetailsDefinitionRequestAck519NoTrdRegPublicationsInit(p *PartyDetailsDefinitionRequestAck519NoTrdRegPublications) {
	return
}

func (*PartyDetailsDefinitionRequestAck519) SbeBlockLength() (blockLength uint16) {
	return 159
}

func (*PartyDetailsDefinitionRequestAck519) SbeTemplateId() (templateId uint16) {
	return 519
}

func (*PartyDetailsDefinitionRequestAck519) SbeSchemaId() (schemaId uint16) {
	return 8
}

func (*PartyDetailsDefinitionRequestAck519) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}

func (*PartyDetailsDefinitionRequestAck519) SbeSemanticType() (semanticType []byte) {
	return []byte("CY")
}

func (*PartyDetailsDefinitionRequestAck519) SeqNumId() uint16 {
	return 9726
}

func (*PartyDetailsDefinitionRequestAck519) SeqNumSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsDefinitionRequestAck519) SeqNumInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.SeqNumSinceVersion()
}

func (*PartyDetailsDefinitionRequestAck519) SeqNumDeprecated() uint16 {
	return 0
}

func (*PartyDetailsDefinitionRequestAck519) SeqNumMetaAttribute(meta int) string {
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

func (*PartyDetailsDefinitionRequestAck519) SeqNumMinValue() uint32 {
	return 0
}

func (*PartyDetailsDefinitionRequestAck519) SeqNumMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*PartyDetailsDefinitionRequestAck519) SeqNumNullValue() uint32 {
	return math.MaxUint32
}

func (*PartyDetailsDefinitionRequestAck519) UUIDId() uint16 {
	return 39001
}

func (*PartyDetailsDefinitionRequestAck519) UUIDSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsDefinitionRequestAck519) UUIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.UUIDSinceVersion()
}

func (*PartyDetailsDefinitionRequestAck519) UUIDDeprecated() uint16 {
	return 0
}

func (*PartyDetailsDefinitionRequestAck519) UUIDMetaAttribute(meta int) string {
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

func (*PartyDetailsDefinitionRequestAck519) UUIDMinValue() uint64 {
	return 0
}

func (*PartyDetailsDefinitionRequestAck519) UUIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*PartyDetailsDefinitionRequestAck519) UUIDNullValue() uint64 {
	return math.MaxUint64
}

func (*PartyDetailsDefinitionRequestAck519) MemoId() uint16 {
	return 5149
}

func (*PartyDetailsDefinitionRequestAck519) MemoSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsDefinitionRequestAck519) MemoInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.MemoSinceVersion()
}

func (*PartyDetailsDefinitionRequestAck519) MemoDeprecated() uint16 {
	return 0
}

func (*PartyDetailsDefinitionRequestAck519) MemoMetaAttribute(meta int) string {
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

func (*PartyDetailsDefinitionRequestAck519) MemoMinValue() byte {
	return byte(32)
}

func (*PartyDetailsDefinitionRequestAck519) MemoMaxValue() byte {
	return byte(126)
}

func (*PartyDetailsDefinitionRequestAck519) MemoNullValue() byte {
	return 0
}

func (p *PartyDetailsDefinitionRequestAck519) MemoCharacterEncoding() string {
	return "US-ASCII"
}

func (*PartyDetailsDefinitionRequestAck519) AvgPxGroupIDId() uint16 {
	return 1731
}

func (*PartyDetailsDefinitionRequestAck519) AvgPxGroupIDSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsDefinitionRequestAck519) AvgPxGroupIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.AvgPxGroupIDSinceVersion()
}

func (*PartyDetailsDefinitionRequestAck519) AvgPxGroupIDDeprecated() uint16 {
	return 0
}

func (*PartyDetailsDefinitionRequestAck519) AvgPxGroupIDMetaAttribute(meta int) string {
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

func (*PartyDetailsDefinitionRequestAck519) AvgPxGroupIDMinValue() byte {
	return byte(32)
}

func (*PartyDetailsDefinitionRequestAck519) AvgPxGroupIDMaxValue() byte {
	return byte(126)
}

func (*PartyDetailsDefinitionRequestAck519) AvgPxGroupIDNullValue() byte {
	return 0
}

func (p *PartyDetailsDefinitionRequestAck519) AvgPxGroupIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*PartyDetailsDefinitionRequestAck519) PartyDetailsListReqIDId() uint16 {
	return 1505
}

func (*PartyDetailsDefinitionRequestAck519) PartyDetailsListReqIDSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsDefinitionRequestAck519) PartyDetailsListReqIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.PartyDetailsListReqIDSinceVersion()
}

func (*PartyDetailsDefinitionRequestAck519) PartyDetailsListReqIDDeprecated() uint16 {
	return 0
}

func (*PartyDetailsDefinitionRequestAck519) PartyDetailsListReqIDMetaAttribute(meta int) string {
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

func (*PartyDetailsDefinitionRequestAck519) PartyDetailsListReqIDMinValue() uint64 {
	return 0
}

func (*PartyDetailsDefinitionRequestAck519) PartyDetailsListReqIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*PartyDetailsDefinitionRequestAck519) PartyDetailsListReqIDNullValue() uint64 {
	return math.MaxUint64
}

func (*PartyDetailsDefinitionRequestAck519) SendingTimeEpochId() uint16 {
	return 5297
}

func (*PartyDetailsDefinitionRequestAck519) SendingTimeEpochSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsDefinitionRequestAck519) SendingTimeEpochInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.SendingTimeEpochSinceVersion()
}

func (*PartyDetailsDefinitionRequestAck519) SendingTimeEpochDeprecated() uint16 {
	return 0
}

func (*PartyDetailsDefinitionRequestAck519) SendingTimeEpochMetaAttribute(meta int) string {
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

func (*PartyDetailsDefinitionRequestAck519) SendingTimeEpochMinValue() uint64 {
	return 0
}

func (*PartyDetailsDefinitionRequestAck519) SendingTimeEpochMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*PartyDetailsDefinitionRequestAck519) SendingTimeEpochNullValue() uint64 {
	return math.MaxUint64
}

func (*PartyDetailsDefinitionRequestAck519) SelfMatchPreventionIDId() uint16 {
	return 2362
}

func (*PartyDetailsDefinitionRequestAck519) SelfMatchPreventionIDSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsDefinitionRequestAck519) SelfMatchPreventionIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.SelfMatchPreventionIDSinceVersion()
}

func (*PartyDetailsDefinitionRequestAck519) SelfMatchPreventionIDDeprecated() uint16 {
	return 0
}

func (*PartyDetailsDefinitionRequestAck519) SelfMatchPreventionIDMetaAttribute(meta int) string {
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

func (*PartyDetailsDefinitionRequestAck519) SelfMatchPreventionIDMinValue() uint64 {
	return 0
}

func (*PartyDetailsDefinitionRequestAck519) SelfMatchPreventionIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*PartyDetailsDefinitionRequestAck519) SelfMatchPreventionIDNullValue() uint64 {
	return 18446744073709551615
}

func (*PartyDetailsDefinitionRequestAck519) PartyDetailRequestStatusId() uint16 {
	return 1878
}

func (*PartyDetailsDefinitionRequestAck519) PartyDetailRequestStatusSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsDefinitionRequestAck519) PartyDetailRequestStatusInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.PartyDetailRequestStatusSinceVersion()
}

func (*PartyDetailsDefinitionRequestAck519) PartyDetailRequestStatusDeprecated() uint16 {
	return 0
}

func (*PartyDetailsDefinitionRequestAck519) PartyDetailRequestStatusMetaAttribute(meta int) string {
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

func (*PartyDetailsDefinitionRequestAck519) PartyDetailRequestStatusMinValue() uint8 {
	return 0
}

func (*PartyDetailsDefinitionRequestAck519) PartyDetailRequestStatusMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*PartyDetailsDefinitionRequestAck519) PartyDetailRequestStatusNullValue() uint8 {
	return math.MaxUint8
}

func (*PartyDetailsDefinitionRequestAck519) CustOrderCapacityId() uint16 {
	return 582
}

func (*PartyDetailsDefinitionRequestAck519) CustOrderCapacitySinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsDefinitionRequestAck519) CustOrderCapacityInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.CustOrderCapacitySinceVersion()
}

func (*PartyDetailsDefinitionRequestAck519) CustOrderCapacityDeprecated() uint16 {
	return 0
}

func (*PartyDetailsDefinitionRequestAck519) CustOrderCapacityMetaAttribute(meta int) string {
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

func (*PartyDetailsDefinitionRequestAck519) ClearingAccountTypeId() uint16 {
	return 1816
}

func (*PartyDetailsDefinitionRequestAck519) ClearingAccountTypeSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsDefinitionRequestAck519) ClearingAccountTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.ClearingAccountTypeSinceVersion()
}

func (*PartyDetailsDefinitionRequestAck519) ClearingAccountTypeDeprecated() uint16 {
	return 0
}

func (*PartyDetailsDefinitionRequestAck519) ClearingAccountTypeMetaAttribute(meta int) string {
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

func (*PartyDetailsDefinitionRequestAck519) SelfMatchPreventionInstructionId() uint16 {
	return 8000
}

func (*PartyDetailsDefinitionRequestAck519) SelfMatchPreventionInstructionSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsDefinitionRequestAck519) SelfMatchPreventionInstructionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.SelfMatchPreventionInstructionSinceVersion()
}

func (*PartyDetailsDefinitionRequestAck519) SelfMatchPreventionInstructionDeprecated() uint16 {
	return 0
}

func (*PartyDetailsDefinitionRequestAck519) SelfMatchPreventionInstructionMetaAttribute(meta int) string {
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

func (*PartyDetailsDefinitionRequestAck519) AvgPxIndicatorId() uint16 {
	return 819
}

func (*PartyDetailsDefinitionRequestAck519) AvgPxIndicatorSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsDefinitionRequestAck519) AvgPxIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.AvgPxIndicatorSinceVersion()
}

func (*PartyDetailsDefinitionRequestAck519) AvgPxIndicatorDeprecated() uint16 {
	return 0
}

func (*PartyDetailsDefinitionRequestAck519) AvgPxIndicatorMetaAttribute(meta int) string {
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

func (*PartyDetailsDefinitionRequestAck519) ClearingTradePriceTypeId() uint16 {
	return 1598
}

func (*PartyDetailsDefinitionRequestAck519) ClearingTradePriceTypeSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsDefinitionRequestAck519) ClearingTradePriceTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.ClearingTradePriceTypeSinceVersion()
}

func (*PartyDetailsDefinitionRequestAck519) ClearingTradePriceTypeDeprecated() uint16 {
	return 0
}

func (*PartyDetailsDefinitionRequestAck519) ClearingTradePriceTypeMetaAttribute(meta int) string {
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

func (*PartyDetailsDefinitionRequestAck519) CmtaGiveupCDId() uint16 {
	return 9708
}

func (*PartyDetailsDefinitionRequestAck519) CmtaGiveupCDSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsDefinitionRequestAck519) CmtaGiveupCDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.CmtaGiveupCDSinceVersion()
}

func (*PartyDetailsDefinitionRequestAck519) CmtaGiveupCDDeprecated() uint16 {
	return 0
}

func (*PartyDetailsDefinitionRequestAck519) CmtaGiveupCDMetaAttribute(meta int) string {
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

func (*PartyDetailsDefinitionRequestAck519) CustOrderHandlingInstId() uint16 {
	return 1031
}

func (*PartyDetailsDefinitionRequestAck519) CustOrderHandlingInstSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsDefinitionRequestAck519) CustOrderHandlingInstInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.CustOrderHandlingInstSinceVersion()
}

func (*PartyDetailsDefinitionRequestAck519) CustOrderHandlingInstDeprecated() uint16 {
	return 0
}

func (*PartyDetailsDefinitionRequestAck519) CustOrderHandlingInstMetaAttribute(meta int) string {
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

func (*PartyDetailsDefinitionRequestAck519) NoPartyUpdatesId() uint16 {
	return 1676
}

func (*PartyDetailsDefinitionRequestAck519) NoPartyUpdatesSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsDefinitionRequestAck519) NoPartyUpdatesInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.NoPartyUpdatesSinceVersion()
}

func (*PartyDetailsDefinitionRequestAck519) NoPartyUpdatesDeprecated() uint16 {
	return 0
}

func (*PartyDetailsDefinitionRequestAck519) NoPartyUpdatesMetaAttribute(meta int) string {
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

func (*PartyDetailsDefinitionRequestAck519) NoPartyUpdatesMinValue() uint8 {
	return 0
}

func (*PartyDetailsDefinitionRequestAck519) NoPartyUpdatesMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*PartyDetailsDefinitionRequestAck519) NoPartyUpdatesNullValue() uint8 {
	return math.MaxUint8
}

func (*PartyDetailsDefinitionRequestAck519) ListUpdateActionId() uint16 {
	return 1324
}

func (*PartyDetailsDefinitionRequestAck519) ListUpdateActionSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsDefinitionRequestAck519) ListUpdateActionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.ListUpdateActionSinceVersion()
}

func (*PartyDetailsDefinitionRequestAck519) ListUpdateActionDeprecated() uint16 {
	return 0
}

func (*PartyDetailsDefinitionRequestAck519) ListUpdateActionMetaAttribute(meta int) string {
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

func (*PartyDetailsDefinitionRequestAck519) PartyDetailDefinitionStatusId() uint16 {
	return 1879
}

func (*PartyDetailsDefinitionRequestAck519) PartyDetailDefinitionStatusSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsDefinitionRequestAck519) PartyDetailDefinitionStatusInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.PartyDetailDefinitionStatusSinceVersion()
}

func (*PartyDetailsDefinitionRequestAck519) PartyDetailDefinitionStatusDeprecated() uint16 {
	return 0
}

func (*PartyDetailsDefinitionRequestAck519) PartyDetailDefinitionStatusMetaAttribute(meta int) string {
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

func (*PartyDetailsDefinitionRequestAck519) PartyDetailDefinitionStatusMinValue() uint8 {
	return 0
}

func (*PartyDetailsDefinitionRequestAck519) PartyDetailDefinitionStatusMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*PartyDetailsDefinitionRequestAck519) PartyDetailDefinitionStatusNullValue() uint8 {
	return math.MaxUint8
}

func (*PartyDetailsDefinitionRequestAck519) ExecutorId() uint16 {
	return 5290
}

func (*PartyDetailsDefinitionRequestAck519) ExecutorSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsDefinitionRequestAck519) ExecutorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.ExecutorSinceVersion()
}

func (*PartyDetailsDefinitionRequestAck519) ExecutorDeprecated() uint16 {
	return 0
}

func (*PartyDetailsDefinitionRequestAck519) ExecutorMetaAttribute(meta int) string {
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

func (*PartyDetailsDefinitionRequestAck519) ExecutorMinValue() uint64 {
	return 0
}

func (*PartyDetailsDefinitionRequestAck519) ExecutorMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*PartyDetailsDefinitionRequestAck519) ExecutorNullValue() uint64 {
	return 18446744073709551615
}

func (*PartyDetailsDefinitionRequestAck519) IDMShortCodeId() uint16 {
	return 36023
}

func (*PartyDetailsDefinitionRequestAck519) IDMShortCodeSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsDefinitionRequestAck519) IDMShortCodeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.IDMShortCodeSinceVersion()
}

func (*PartyDetailsDefinitionRequestAck519) IDMShortCodeDeprecated() uint16 {
	return 0
}

func (*PartyDetailsDefinitionRequestAck519) IDMShortCodeMetaAttribute(meta int) string {
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

func (*PartyDetailsDefinitionRequestAck519) IDMShortCodeMinValue() uint64 {
	return 0
}

func (*PartyDetailsDefinitionRequestAck519) IDMShortCodeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*PartyDetailsDefinitionRequestAck519) IDMShortCodeNullValue() uint64 {
	return 18446744073709551615
}

func (*PartyDetailsDefinitionRequestAck519) PossRetransFlagId() uint16 {
	return 9765
}

func (*PartyDetailsDefinitionRequestAck519) PossRetransFlagSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsDefinitionRequestAck519) PossRetransFlagInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.PossRetransFlagSinceVersion()
}

func (*PartyDetailsDefinitionRequestAck519) PossRetransFlagDeprecated() uint16 {
	return 0
}

func (*PartyDetailsDefinitionRequestAck519) PossRetransFlagMetaAttribute(meta int) string {
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

func (*PartyDetailsDefinitionRequestAck519) SplitMsgId() uint16 {
	return 9553
}

func (*PartyDetailsDefinitionRequestAck519) SplitMsgSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsDefinitionRequestAck519) SplitMsgInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.SplitMsgSinceVersion()
}

func (*PartyDetailsDefinitionRequestAck519) SplitMsgDeprecated() uint16 {
	return 0
}

func (*PartyDetailsDefinitionRequestAck519) SplitMsgMetaAttribute(meta int) string {
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

func (*PartyDetailsDefinitionRequestAck519NoPartyDetails) PartyDetailIDId() uint16 {
	return 1691
}

func (*PartyDetailsDefinitionRequestAck519NoPartyDetails) PartyDetailIDSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsDefinitionRequestAck519NoPartyDetails) PartyDetailIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.PartyDetailIDSinceVersion()
}

func (*PartyDetailsDefinitionRequestAck519NoPartyDetails) PartyDetailIDDeprecated() uint16 {
	return 0
}

func (*PartyDetailsDefinitionRequestAck519NoPartyDetails) PartyDetailIDMetaAttribute(meta int) string {
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

func (*PartyDetailsDefinitionRequestAck519NoPartyDetails) PartyDetailIDMinValue() byte {
	return byte(32)
}

func (*PartyDetailsDefinitionRequestAck519NoPartyDetails) PartyDetailIDMaxValue() byte {
	return byte(126)
}

func (*PartyDetailsDefinitionRequestAck519NoPartyDetails) PartyDetailIDNullValue() byte {
	return 0
}

func (p *PartyDetailsDefinitionRequestAck519NoPartyDetails) PartyDetailIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*PartyDetailsDefinitionRequestAck519NoPartyDetails) PartyDetailIDSourceId() uint16 {
	return 1692
}

func (*PartyDetailsDefinitionRequestAck519NoPartyDetails) PartyDetailIDSourceSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsDefinitionRequestAck519NoPartyDetails) PartyDetailIDSourceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.PartyDetailIDSourceSinceVersion()
}

func (*PartyDetailsDefinitionRequestAck519NoPartyDetails) PartyDetailIDSourceDeprecated() uint16 {
	return 0
}

func (*PartyDetailsDefinitionRequestAck519NoPartyDetails) PartyDetailIDSourceMetaAttribute(meta int) string {
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

func (*PartyDetailsDefinitionRequestAck519NoPartyDetails) PartyDetailIDSourceMinValue() byte {
	return byte(32)
}

func (*PartyDetailsDefinitionRequestAck519NoPartyDetails) PartyDetailIDSourceMaxValue() byte {
	return byte(126)
}

func (*PartyDetailsDefinitionRequestAck519NoPartyDetails) PartyDetailIDSourceNullValue() byte {
	return 0
}

func (*PartyDetailsDefinitionRequestAck519NoPartyDetails) PartyDetailRoleId() uint16 {
	return 1693
}

func (*PartyDetailsDefinitionRequestAck519NoPartyDetails) PartyDetailRoleSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsDefinitionRequestAck519NoPartyDetails) PartyDetailRoleInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.PartyDetailRoleSinceVersion()
}

func (*PartyDetailsDefinitionRequestAck519NoPartyDetails) PartyDetailRoleDeprecated() uint16 {
	return 0
}

func (*PartyDetailsDefinitionRequestAck519NoPartyDetails) PartyDetailRoleMetaAttribute(meta int) string {
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

func (*PartyDetailsDefinitionRequestAck519NoTrdRegPublications) TrdRegPublicationTypeId() uint16 {
	return 2669
}

func (*PartyDetailsDefinitionRequestAck519NoTrdRegPublications) TrdRegPublicationTypeSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsDefinitionRequestAck519NoTrdRegPublications) TrdRegPublicationTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.TrdRegPublicationTypeSinceVersion()
}

func (*PartyDetailsDefinitionRequestAck519NoTrdRegPublications) TrdRegPublicationTypeDeprecated() uint16 {
	return 0
}

func (*PartyDetailsDefinitionRequestAck519NoTrdRegPublications) TrdRegPublicationTypeMetaAttribute(meta int) string {
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

func (*PartyDetailsDefinitionRequestAck519NoTrdRegPublications) TrdRegPublicationTypeMinValue() uint8 {
	return 0
}

func (*PartyDetailsDefinitionRequestAck519NoTrdRegPublications) TrdRegPublicationTypeMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*PartyDetailsDefinitionRequestAck519NoTrdRegPublications) TrdRegPublicationTypeNullValue() uint8 {
	return math.MaxUint8
}

func (*PartyDetailsDefinitionRequestAck519NoTrdRegPublications) TrdRegPublicationReasonId() uint16 {
	return 2670
}

func (*PartyDetailsDefinitionRequestAck519NoTrdRegPublications) TrdRegPublicationReasonSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsDefinitionRequestAck519NoTrdRegPublications) TrdRegPublicationReasonInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.TrdRegPublicationReasonSinceVersion()
}

func (*PartyDetailsDefinitionRequestAck519NoTrdRegPublications) TrdRegPublicationReasonDeprecated() uint16 {
	return 0
}

func (*PartyDetailsDefinitionRequestAck519NoTrdRegPublications) TrdRegPublicationReasonMetaAttribute(meta int) string {
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

func (*PartyDetailsDefinitionRequestAck519NoTrdRegPublications) TrdRegPublicationReasonMinValue() uint8 {
	return 0
}

func (*PartyDetailsDefinitionRequestAck519NoTrdRegPublications) TrdRegPublicationReasonMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*PartyDetailsDefinitionRequestAck519NoTrdRegPublications) TrdRegPublicationReasonNullValue() uint8 {
	return math.MaxUint8
}

func (*PartyDetailsDefinitionRequestAck519) NoPartyDetailsId() uint16 {
	return 1671
}

func (*PartyDetailsDefinitionRequestAck519) NoPartyDetailsSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsDefinitionRequestAck519) NoPartyDetailsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.NoPartyDetailsSinceVersion()
}

func (*PartyDetailsDefinitionRequestAck519) NoPartyDetailsDeprecated() uint16 {
	return 0
}

func (*PartyDetailsDefinitionRequestAck519NoPartyDetails) SbeBlockLength() (blockLength uint) {
	return 22
}

func (*PartyDetailsDefinitionRequestAck519NoPartyDetails) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}

func (*PartyDetailsDefinitionRequestAck519) NoTrdRegPublicationsId() uint16 {
	return 2668
}

func (*PartyDetailsDefinitionRequestAck519) NoTrdRegPublicationsSinceVersion() uint16 {
	return 2
}

func (p *PartyDetailsDefinitionRequestAck519) NoTrdRegPublicationsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.NoTrdRegPublicationsSinceVersion()
}

func (*PartyDetailsDefinitionRequestAck519) NoTrdRegPublicationsDeprecated() uint16 {
	return 0
}

func (*PartyDetailsDefinitionRequestAck519NoTrdRegPublications) SbeBlockLength() (blockLength uint) {
	return 2
}

func (*PartyDetailsDefinitionRequestAck519NoTrdRegPublications) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}
