// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type PartyDetailsListReport538 struct {
	SeqNum                         uint32
	UUID                           uint64
	AvgPxGroupID                   [20]byte
	PartyDetailsListReqID          uint64
	PartyDetailsListReportID       uint64
	SendingTimeEpoch               uint64
	SelfMatchPreventionID          uint64
	TotNumParties                  uint16
	RequestResult                  ReqResultEnum
	LastFragment                   BooleanFlagEnum
	CustOrderCapacity              CustOrderCapacityEnum
	ClearingAccountType            ClearingAcctTypeEnum
	SelfMatchPreventionInstruction SMPIEnum
	AvgPxIndicator                 AvgPxIndEnum
	ClearingTradePriceType         SLEDSEnum
	CmtaGiveupCD                   CmtaGiveUpCDEnum
	CustOrderHandlingInst          CustOrdHandlInstEnum
	Executor                       uint64
	IDMShortCode                   uint64
	PossRetransFlag                BooleanFlagEnum
	SplitMsg                       SplitMsgEnum
	NoPartyDetails                 []PartyDetailsListReport538NoPartyDetails
	NoTrdRegPublications           []PartyDetailsListReport538NoTrdRegPublications
}
type PartyDetailsListReport538NoPartyDetails struct {
	PartyDetailID       [20]byte
	PartyDetailIDSource [1]byte
	PartyDetailRole     PartyDetailRoleEnum
}
type PartyDetailsListReport538NoTrdRegPublications struct {
	TrdRegPublicationType   uint8
	TrdRegPublicationReason uint8
}

func (p *PartyDetailsListReport538) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
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
	if err := _m.WriteBytes(_w, p.AvgPxGroupID[:]); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, p.PartyDetailsListReqID); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, p.PartyDetailsListReportID); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, p.SendingTimeEpoch); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, p.SelfMatchPreventionID); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, p.TotNumParties); err != nil {
		return err
	}
	if err := p.RequestResult.Encode(_m, _w); err != nil {
		return err
	}
	if err := p.LastFragment.Encode(_m, _w); err != nil {
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

func (p *PartyDetailsListReport538) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
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
	if !p.PartyDetailsListReportIDInActingVersion(actingVersion) {
		p.PartyDetailsListReportID = p.PartyDetailsListReportIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &p.PartyDetailsListReportID); err != nil {
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
	if !p.TotNumPartiesInActingVersion(actingVersion) {
		p.TotNumParties = p.TotNumPartiesNullValue()
	} else {
		if err := _m.ReadUint16(_r, &p.TotNumParties); err != nil {
			return err
		}
	}
	if p.RequestResultInActingVersion(actingVersion) {
		if err := p.RequestResult.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if p.LastFragmentInActingVersion(actingVersion) {
		if err := p.LastFragment.Decode(_m, _r, actingVersion); err != nil {
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
			p.NoPartyDetails = make([]PartyDetailsListReport538NoPartyDetails, NoPartyDetailsNumInGroup)
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
			p.NoTrdRegPublications = make([]PartyDetailsListReport538NoTrdRegPublications, NoTrdRegPublicationsNumInGroup)
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

func (p *PartyDetailsListReport538) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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
	if p.PartyDetailsListReportIDInActingVersion(actingVersion) {
		if p.PartyDetailsListReportID < p.PartyDetailsListReportIDMinValue() || p.PartyDetailsListReportID > p.PartyDetailsListReportIDMaxValue() {
			return fmt.Errorf("Range check failed on p.PartyDetailsListReportID (%v < %v > %v)", p.PartyDetailsListReportIDMinValue(), p.PartyDetailsListReportID, p.PartyDetailsListReportIDMaxValue())
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
	if p.TotNumPartiesInActingVersion(actingVersion) {
		if p.TotNumParties < p.TotNumPartiesMinValue() || p.TotNumParties > p.TotNumPartiesMaxValue() {
			return fmt.Errorf("Range check failed on p.TotNumParties (%v < %v > %v)", p.TotNumPartiesMinValue(), p.TotNumParties, p.TotNumPartiesMaxValue())
		}
	}
	if err := p.RequestResult.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := p.LastFragment.RangeCheck(actingVersion, schemaVersion); err != nil {
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
	if err := p.CmtaGiveupCD.RangeCheck(actingVersion, schemaVersion); err != nil {
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

func PartyDetailsListReport538Init(p *PartyDetailsListReport538) {
	for idx := 0; idx < 20; idx++ {
		p.AvgPxGroupID[idx] = 0
	}
	p.SelfMatchPreventionID = 18446744073709551615
	p.Executor = 18446744073709551615
	p.IDMShortCode = 18446744073709551615
	return
}

func (p *PartyDetailsListReport538NoPartyDetails) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteBytes(_w, p.PartyDetailID[:]); err != nil {
		return err
	}
	if err := p.PartyDetailRole.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (p *PartyDetailsListReport538NoPartyDetails) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
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

func (p *PartyDetailsListReport538NoPartyDetails) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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

func PartyDetailsListReport538NoPartyDetailsInit(p *PartyDetailsListReport538NoPartyDetails) {
	p.PartyDetailIDSource[0] = 67
	return
}

func (p *PartyDetailsListReport538NoTrdRegPublications) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, p.TrdRegPublicationType); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, p.TrdRegPublicationReason); err != nil {
		return err
	}
	return nil
}

func (p *PartyDetailsListReport538NoTrdRegPublications) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
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

func (p *PartyDetailsListReport538NoTrdRegPublications) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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

func PartyDetailsListReport538NoTrdRegPublicationsInit(p *PartyDetailsListReport538NoTrdRegPublications) {
	return
}

func (*PartyDetailsListReport538) SbeBlockLength() (blockLength uint16) {
	return 93
}

func (*PartyDetailsListReport538) SbeTemplateId() (templateId uint16) {
	return 538
}

func (*PartyDetailsListReport538) SbeSchemaId() (schemaId uint16) {
	return 8
}

func (*PartyDetailsListReport538) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}

func (*PartyDetailsListReport538) SbeSemanticType() (semanticType []byte) {
	return []byte("CG")
}

func (*PartyDetailsListReport538) SeqNumId() uint16 {
	return 9726
}

func (*PartyDetailsListReport538) SeqNumSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsListReport538) SeqNumInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.SeqNumSinceVersion()
}

func (*PartyDetailsListReport538) SeqNumDeprecated() uint16 {
	return 0
}

func (*PartyDetailsListReport538) SeqNumMetaAttribute(meta int) string {
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

func (*PartyDetailsListReport538) SeqNumMinValue() uint32 {
	return 0
}

func (*PartyDetailsListReport538) SeqNumMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*PartyDetailsListReport538) SeqNumNullValue() uint32 {
	return math.MaxUint32
}

func (*PartyDetailsListReport538) UUIDId() uint16 {
	return 39001
}

func (*PartyDetailsListReport538) UUIDSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsListReport538) UUIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.UUIDSinceVersion()
}

func (*PartyDetailsListReport538) UUIDDeprecated() uint16 {
	return 0
}

func (*PartyDetailsListReport538) UUIDMetaAttribute(meta int) string {
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

func (*PartyDetailsListReport538) UUIDMinValue() uint64 {
	return 0
}

func (*PartyDetailsListReport538) UUIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*PartyDetailsListReport538) UUIDNullValue() uint64 {
	return math.MaxUint64
}

func (*PartyDetailsListReport538) AvgPxGroupIDId() uint16 {
	return 1731
}

func (*PartyDetailsListReport538) AvgPxGroupIDSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsListReport538) AvgPxGroupIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.AvgPxGroupIDSinceVersion()
}

func (*PartyDetailsListReport538) AvgPxGroupIDDeprecated() uint16 {
	return 0
}

func (*PartyDetailsListReport538) AvgPxGroupIDMetaAttribute(meta int) string {
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

func (*PartyDetailsListReport538) AvgPxGroupIDMinValue() byte {
	return byte(32)
}

func (*PartyDetailsListReport538) AvgPxGroupIDMaxValue() byte {
	return byte(126)
}

func (*PartyDetailsListReport538) AvgPxGroupIDNullValue() byte {
	return 0
}

func (p *PartyDetailsListReport538) AvgPxGroupIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*PartyDetailsListReport538) PartyDetailsListReqIDId() uint16 {
	return 1505
}

func (*PartyDetailsListReport538) PartyDetailsListReqIDSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsListReport538) PartyDetailsListReqIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.PartyDetailsListReqIDSinceVersion()
}

func (*PartyDetailsListReport538) PartyDetailsListReqIDDeprecated() uint16 {
	return 0
}

func (*PartyDetailsListReport538) PartyDetailsListReqIDMetaAttribute(meta int) string {
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

func (*PartyDetailsListReport538) PartyDetailsListReqIDMinValue() uint64 {
	return 0
}

func (*PartyDetailsListReport538) PartyDetailsListReqIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*PartyDetailsListReport538) PartyDetailsListReqIDNullValue() uint64 {
	return math.MaxUint64
}

func (*PartyDetailsListReport538) PartyDetailsListReportIDId() uint16 {
	return 1510
}

func (*PartyDetailsListReport538) PartyDetailsListReportIDSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsListReport538) PartyDetailsListReportIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.PartyDetailsListReportIDSinceVersion()
}

func (*PartyDetailsListReport538) PartyDetailsListReportIDDeprecated() uint16 {
	return 0
}

func (*PartyDetailsListReport538) PartyDetailsListReportIDMetaAttribute(meta int) string {
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

func (*PartyDetailsListReport538) PartyDetailsListReportIDMinValue() uint64 {
	return 0
}

func (*PartyDetailsListReport538) PartyDetailsListReportIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*PartyDetailsListReport538) PartyDetailsListReportIDNullValue() uint64 {
	return math.MaxUint64
}

func (*PartyDetailsListReport538) SendingTimeEpochId() uint16 {
	return 5297
}

func (*PartyDetailsListReport538) SendingTimeEpochSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsListReport538) SendingTimeEpochInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.SendingTimeEpochSinceVersion()
}

func (*PartyDetailsListReport538) SendingTimeEpochDeprecated() uint16 {
	return 0
}

func (*PartyDetailsListReport538) SendingTimeEpochMetaAttribute(meta int) string {
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

func (*PartyDetailsListReport538) SendingTimeEpochMinValue() uint64 {
	return 0
}

func (*PartyDetailsListReport538) SendingTimeEpochMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*PartyDetailsListReport538) SendingTimeEpochNullValue() uint64 {
	return math.MaxUint64
}

func (*PartyDetailsListReport538) SelfMatchPreventionIDId() uint16 {
	return 2362
}

func (*PartyDetailsListReport538) SelfMatchPreventionIDSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsListReport538) SelfMatchPreventionIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.SelfMatchPreventionIDSinceVersion()
}

func (*PartyDetailsListReport538) SelfMatchPreventionIDDeprecated() uint16 {
	return 0
}

func (*PartyDetailsListReport538) SelfMatchPreventionIDMetaAttribute(meta int) string {
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

func (*PartyDetailsListReport538) SelfMatchPreventionIDMinValue() uint64 {
	return 0
}

func (*PartyDetailsListReport538) SelfMatchPreventionIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*PartyDetailsListReport538) SelfMatchPreventionIDNullValue() uint64 {
	return 18446744073709551615
}

func (*PartyDetailsListReport538) TotNumPartiesId() uint16 {
	return 1512
}

func (*PartyDetailsListReport538) TotNumPartiesSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsListReport538) TotNumPartiesInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.TotNumPartiesSinceVersion()
}

func (*PartyDetailsListReport538) TotNumPartiesDeprecated() uint16 {
	return 0
}

func (*PartyDetailsListReport538) TotNumPartiesMetaAttribute(meta int) string {
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

func (*PartyDetailsListReport538) TotNumPartiesMinValue() uint16 {
	return 0
}

func (*PartyDetailsListReport538) TotNumPartiesMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*PartyDetailsListReport538) TotNumPartiesNullValue() uint16 {
	return math.MaxUint16
}

func (*PartyDetailsListReport538) RequestResultId() uint16 {
	return 1511
}

func (*PartyDetailsListReport538) RequestResultSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsListReport538) RequestResultInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.RequestResultSinceVersion()
}

func (*PartyDetailsListReport538) RequestResultDeprecated() uint16 {
	return 0
}

func (*PartyDetailsListReport538) RequestResultMetaAttribute(meta int) string {
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

func (*PartyDetailsListReport538) LastFragmentId() uint16 {
	return 893
}

func (*PartyDetailsListReport538) LastFragmentSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsListReport538) LastFragmentInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.LastFragmentSinceVersion()
}

func (*PartyDetailsListReport538) LastFragmentDeprecated() uint16 {
	return 0
}

func (*PartyDetailsListReport538) LastFragmentMetaAttribute(meta int) string {
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

func (*PartyDetailsListReport538) CustOrderCapacityId() uint16 {
	return 582
}

func (*PartyDetailsListReport538) CustOrderCapacitySinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsListReport538) CustOrderCapacityInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.CustOrderCapacitySinceVersion()
}

func (*PartyDetailsListReport538) CustOrderCapacityDeprecated() uint16 {
	return 0
}

func (*PartyDetailsListReport538) CustOrderCapacityMetaAttribute(meta int) string {
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

func (*PartyDetailsListReport538) ClearingAccountTypeId() uint16 {
	return 1816
}

func (*PartyDetailsListReport538) ClearingAccountTypeSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsListReport538) ClearingAccountTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.ClearingAccountTypeSinceVersion()
}

func (*PartyDetailsListReport538) ClearingAccountTypeDeprecated() uint16 {
	return 0
}

func (*PartyDetailsListReport538) ClearingAccountTypeMetaAttribute(meta int) string {
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

func (*PartyDetailsListReport538) SelfMatchPreventionInstructionId() uint16 {
	return 8000
}

func (*PartyDetailsListReport538) SelfMatchPreventionInstructionSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsListReport538) SelfMatchPreventionInstructionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.SelfMatchPreventionInstructionSinceVersion()
}

func (*PartyDetailsListReport538) SelfMatchPreventionInstructionDeprecated() uint16 {
	return 0
}

func (*PartyDetailsListReport538) SelfMatchPreventionInstructionMetaAttribute(meta int) string {
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

func (*PartyDetailsListReport538) AvgPxIndicatorId() uint16 {
	return 819
}

func (*PartyDetailsListReport538) AvgPxIndicatorSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsListReport538) AvgPxIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.AvgPxIndicatorSinceVersion()
}

func (*PartyDetailsListReport538) AvgPxIndicatorDeprecated() uint16 {
	return 0
}

func (*PartyDetailsListReport538) AvgPxIndicatorMetaAttribute(meta int) string {
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

func (*PartyDetailsListReport538) ClearingTradePriceTypeId() uint16 {
	return 1598
}

func (*PartyDetailsListReport538) ClearingTradePriceTypeSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsListReport538) ClearingTradePriceTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.ClearingTradePriceTypeSinceVersion()
}

func (*PartyDetailsListReport538) ClearingTradePriceTypeDeprecated() uint16 {
	return 0
}

func (*PartyDetailsListReport538) ClearingTradePriceTypeMetaAttribute(meta int) string {
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

func (*PartyDetailsListReport538) CmtaGiveupCDId() uint16 {
	return 9708
}

func (*PartyDetailsListReport538) CmtaGiveupCDSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsListReport538) CmtaGiveupCDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.CmtaGiveupCDSinceVersion()
}

func (*PartyDetailsListReport538) CmtaGiveupCDDeprecated() uint16 {
	return 0
}

func (*PartyDetailsListReport538) CmtaGiveupCDMetaAttribute(meta int) string {
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

func (*PartyDetailsListReport538) CustOrderHandlingInstId() uint16 {
	return 1031
}

func (*PartyDetailsListReport538) CustOrderHandlingInstSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsListReport538) CustOrderHandlingInstInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.CustOrderHandlingInstSinceVersion()
}

func (*PartyDetailsListReport538) CustOrderHandlingInstDeprecated() uint16 {
	return 0
}

func (*PartyDetailsListReport538) CustOrderHandlingInstMetaAttribute(meta int) string {
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

func (*PartyDetailsListReport538) ExecutorId() uint16 {
	return 5290
}

func (*PartyDetailsListReport538) ExecutorSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsListReport538) ExecutorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.ExecutorSinceVersion()
}

func (*PartyDetailsListReport538) ExecutorDeprecated() uint16 {
	return 0
}

func (*PartyDetailsListReport538) ExecutorMetaAttribute(meta int) string {
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

func (*PartyDetailsListReport538) ExecutorMinValue() uint64 {
	return 0
}

func (*PartyDetailsListReport538) ExecutorMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*PartyDetailsListReport538) ExecutorNullValue() uint64 {
	return 18446744073709551615
}

func (*PartyDetailsListReport538) IDMShortCodeId() uint16 {
	return 36023
}

func (*PartyDetailsListReport538) IDMShortCodeSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsListReport538) IDMShortCodeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.IDMShortCodeSinceVersion()
}

func (*PartyDetailsListReport538) IDMShortCodeDeprecated() uint16 {
	return 0
}

func (*PartyDetailsListReport538) IDMShortCodeMetaAttribute(meta int) string {
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

func (*PartyDetailsListReport538) IDMShortCodeMinValue() uint64 {
	return 0
}

func (*PartyDetailsListReport538) IDMShortCodeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*PartyDetailsListReport538) IDMShortCodeNullValue() uint64 {
	return 18446744073709551615
}

func (*PartyDetailsListReport538) PossRetransFlagId() uint16 {
	return 9765
}

func (*PartyDetailsListReport538) PossRetransFlagSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsListReport538) PossRetransFlagInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.PossRetransFlagSinceVersion()
}

func (*PartyDetailsListReport538) PossRetransFlagDeprecated() uint16 {
	return 0
}

func (*PartyDetailsListReport538) PossRetransFlagMetaAttribute(meta int) string {
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

func (*PartyDetailsListReport538) SplitMsgId() uint16 {
	return 9553
}

func (*PartyDetailsListReport538) SplitMsgSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsListReport538) SplitMsgInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.SplitMsgSinceVersion()
}

func (*PartyDetailsListReport538) SplitMsgDeprecated() uint16 {
	return 0
}

func (*PartyDetailsListReport538) SplitMsgMetaAttribute(meta int) string {
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

func (*PartyDetailsListReport538NoPartyDetails) PartyDetailIDId() uint16 {
	return 1691
}

func (*PartyDetailsListReport538NoPartyDetails) PartyDetailIDSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsListReport538NoPartyDetails) PartyDetailIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.PartyDetailIDSinceVersion()
}

func (*PartyDetailsListReport538NoPartyDetails) PartyDetailIDDeprecated() uint16 {
	return 0
}

func (*PartyDetailsListReport538NoPartyDetails) PartyDetailIDMetaAttribute(meta int) string {
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

func (*PartyDetailsListReport538NoPartyDetails) PartyDetailIDMinValue() byte {
	return byte(32)
}

func (*PartyDetailsListReport538NoPartyDetails) PartyDetailIDMaxValue() byte {
	return byte(126)
}

func (*PartyDetailsListReport538NoPartyDetails) PartyDetailIDNullValue() byte {
	return 0
}

func (p *PartyDetailsListReport538NoPartyDetails) PartyDetailIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*PartyDetailsListReport538NoPartyDetails) PartyDetailIDSourceId() uint16 {
	return 1692
}

func (*PartyDetailsListReport538NoPartyDetails) PartyDetailIDSourceSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsListReport538NoPartyDetails) PartyDetailIDSourceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.PartyDetailIDSourceSinceVersion()
}

func (*PartyDetailsListReport538NoPartyDetails) PartyDetailIDSourceDeprecated() uint16 {
	return 0
}

func (*PartyDetailsListReport538NoPartyDetails) PartyDetailIDSourceMetaAttribute(meta int) string {
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

func (*PartyDetailsListReport538NoPartyDetails) PartyDetailIDSourceMinValue() byte {
	return byte(32)
}

func (*PartyDetailsListReport538NoPartyDetails) PartyDetailIDSourceMaxValue() byte {
	return byte(126)
}

func (*PartyDetailsListReport538NoPartyDetails) PartyDetailIDSourceNullValue() byte {
	return 0
}

func (*PartyDetailsListReport538NoPartyDetails) PartyDetailRoleId() uint16 {
	return 1693
}

func (*PartyDetailsListReport538NoPartyDetails) PartyDetailRoleSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsListReport538NoPartyDetails) PartyDetailRoleInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.PartyDetailRoleSinceVersion()
}

func (*PartyDetailsListReport538NoPartyDetails) PartyDetailRoleDeprecated() uint16 {
	return 0
}

func (*PartyDetailsListReport538NoPartyDetails) PartyDetailRoleMetaAttribute(meta int) string {
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

func (*PartyDetailsListReport538NoTrdRegPublications) TrdRegPublicationTypeId() uint16 {
	return 2669
}

func (*PartyDetailsListReport538NoTrdRegPublications) TrdRegPublicationTypeSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsListReport538NoTrdRegPublications) TrdRegPublicationTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.TrdRegPublicationTypeSinceVersion()
}

func (*PartyDetailsListReport538NoTrdRegPublications) TrdRegPublicationTypeDeprecated() uint16 {
	return 0
}

func (*PartyDetailsListReport538NoTrdRegPublications) TrdRegPublicationTypeMetaAttribute(meta int) string {
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

func (*PartyDetailsListReport538NoTrdRegPublications) TrdRegPublicationTypeMinValue() uint8 {
	return 0
}

func (*PartyDetailsListReport538NoTrdRegPublications) TrdRegPublicationTypeMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*PartyDetailsListReport538NoTrdRegPublications) TrdRegPublicationTypeNullValue() uint8 {
	return math.MaxUint8
}

func (*PartyDetailsListReport538NoTrdRegPublications) TrdRegPublicationReasonId() uint16 {
	return 2670
}

func (*PartyDetailsListReport538NoTrdRegPublications) TrdRegPublicationReasonSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsListReport538NoTrdRegPublications) TrdRegPublicationReasonInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.TrdRegPublicationReasonSinceVersion()
}

func (*PartyDetailsListReport538NoTrdRegPublications) TrdRegPublicationReasonDeprecated() uint16 {
	return 0
}

func (*PartyDetailsListReport538NoTrdRegPublications) TrdRegPublicationReasonMetaAttribute(meta int) string {
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

func (*PartyDetailsListReport538NoTrdRegPublications) TrdRegPublicationReasonMinValue() uint8 {
	return 0
}

func (*PartyDetailsListReport538NoTrdRegPublications) TrdRegPublicationReasonMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*PartyDetailsListReport538NoTrdRegPublications) TrdRegPublicationReasonNullValue() uint8 {
	return math.MaxUint8
}

func (*PartyDetailsListReport538) NoPartyDetailsId() uint16 {
	return 1671
}

func (*PartyDetailsListReport538) NoPartyDetailsSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailsListReport538) NoPartyDetailsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.NoPartyDetailsSinceVersion()
}

func (*PartyDetailsListReport538) NoPartyDetailsDeprecated() uint16 {
	return 0
}

func (*PartyDetailsListReport538NoPartyDetails) SbeBlockLength() (blockLength uint) {
	return 22
}

func (*PartyDetailsListReport538NoPartyDetails) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}

func (*PartyDetailsListReport538) NoTrdRegPublicationsId() uint16 {
	return 2668
}

func (*PartyDetailsListReport538) NoTrdRegPublicationsSinceVersion() uint16 {
	return 2
}

func (p *PartyDetailsListReport538) NoTrdRegPublicationsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.NoTrdRegPublicationsSinceVersion()
}

func (*PartyDetailsListReport538) NoTrdRegPublicationsDeprecated() uint16 {
	return 0
}

func (*PartyDetailsListReport538NoTrdRegPublications) SbeBlockLength() (blockLength uint) {
	return 2
}

func (*PartyDetailsListReport538NoTrdRegPublications) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}
