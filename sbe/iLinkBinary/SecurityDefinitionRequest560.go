// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type SecurityDefinitionRequest560 struct {
	PartyDetailsListReqID uint64
	SecurityReqID         uint64
	ManualOrderIndicator  ManualOrdIndReqEnum
	SeqNum                uint32
	SenderID              [20]byte
	SecurityReqType       [1]byte
	SendingTimeEpoch      uint64
	SecuritySubType       [8]byte
	Location              [5]byte
	StartDate             uint16
	EndDate               uint16
	MaxNoOfSubstitutions  uint8
	SourceRepoID          int32
	NoLegs                []SecurityDefinitionRequest560NoLegs
}
type SecurityDefinitionRequest560NoLegs struct {
	LegSecurityIDSource [1]byte
	LegPrice            PRICENULL9
	LegSecurityID       int32
	LegOptionDelta      Decimal32NULL
	LegSide             SideReqEnum
	LegRatioQty         uint8
}

func (s *SecurityDefinitionRequest560) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := s.RangeCheck(s.SbeSchemaVersion(), s.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteUint64(_w, s.PartyDetailsListReqID); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, s.SecurityReqID); err != nil {
		return err
	}
	if err := s.ManualOrderIndicator.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, s.SeqNum); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, s.SenderID[:]); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, s.SendingTimeEpoch); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, s.SecuritySubType[:]); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, s.Location[:]); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, s.StartDate); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, s.EndDate); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, s.MaxNoOfSubstitutions); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, s.SourceRepoID); err != nil {
		return err
	}
	var NoLegsBlockLength uint16 = 19
	if err := _m.WriteUint16(_w, NoLegsBlockLength); err != nil {
		return err
	}
	var NoLegsNumInGroup uint8 = uint8(len(s.NoLegs))
	if err := _m.WriteUint8(_w, NoLegsNumInGroup); err != nil {
		return err
	}
	for _, prop := range s.NoLegs {
		if err := prop.Encode(_m, _w); err != nil {
			return err
		}
	}
	return nil
}

func (s *SecurityDefinitionRequest560) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !s.PartyDetailsListReqIDInActingVersion(actingVersion) {
		s.PartyDetailsListReqID = s.PartyDetailsListReqIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &s.PartyDetailsListReqID); err != nil {
			return err
		}
	}
	if !s.SecurityReqIDInActingVersion(actingVersion) {
		s.SecurityReqID = s.SecurityReqIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &s.SecurityReqID); err != nil {
			return err
		}
	}
	if s.ManualOrderIndicatorInActingVersion(actingVersion) {
		if err := s.ManualOrderIndicator.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !s.SeqNumInActingVersion(actingVersion) {
		s.SeqNum = s.SeqNumNullValue()
	} else {
		if err := _m.ReadUint32(_r, &s.SeqNum); err != nil {
			return err
		}
	}
	if !s.SenderIDInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			s.SenderID[idx] = s.SenderIDNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, s.SenderID[:]); err != nil {
			return err
		}
	}
	s.SecurityReqType[0] = 49
	if !s.SendingTimeEpochInActingVersion(actingVersion) {
		s.SendingTimeEpoch = s.SendingTimeEpochNullValue()
	} else {
		if err := _m.ReadUint64(_r, &s.SendingTimeEpoch); err != nil {
			return err
		}
	}
	if !s.SecuritySubTypeInActingVersion(actingVersion) {
		for idx := 0; idx < 8; idx++ {
			s.SecuritySubType[idx] = s.SecuritySubTypeNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, s.SecuritySubType[:]); err != nil {
			return err
		}
	}
	if !s.LocationInActingVersion(actingVersion) {
		for idx := 0; idx < 5; idx++ {
			s.Location[idx] = s.LocationNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, s.Location[:]); err != nil {
			return err
		}
	}
	if !s.StartDateInActingVersion(actingVersion) {
		s.StartDate = s.StartDateNullValue()
	} else {
		if err := _m.ReadUint16(_r, &s.StartDate); err != nil {
			return err
		}
	}
	if !s.EndDateInActingVersion(actingVersion) {
		s.EndDate = s.EndDateNullValue()
	} else {
		if err := _m.ReadUint16(_r, &s.EndDate); err != nil {
			return err
		}
	}
	if !s.MaxNoOfSubstitutionsInActingVersion(actingVersion) {
		s.MaxNoOfSubstitutions = s.MaxNoOfSubstitutionsNullValue()
	} else {
		if err := _m.ReadUint8(_r, &s.MaxNoOfSubstitutions); err != nil {
			return err
		}
	}
	if !s.SourceRepoIDInActingVersion(actingVersion) {
		s.SourceRepoID = s.SourceRepoIDNullValue()
	} else {
		if err := _m.ReadInt32(_r, &s.SourceRepoID); err != nil {
			return err
		}
	}
	if actingVersion > s.SbeSchemaVersion() && blockLength > s.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-s.SbeBlockLength()))
	}

	if s.NoLegsInActingVersion(actingVersion) {
		var NoLegsBlockLength uint16
		if err := _m.ReadUint16(_r, &NoLegsBlockLength); err != nil {
			return err
		}
		var NoLegsNumInGroup uint8
		if err := _m.ReadUint8(_r, &NoLegsNumInGroup); err != nil {
			return err
		}
		if cap(s.NoLegs) < int(NoLegsNumInGroup) {
			s.NoLegs = make([]SecurityDefinitionRequest560NoLegs, NoLegsNumInGroup)
		}
		for i, _ := range s.NoLegs {
			if err := s.NoLegs[i].Decode(_m, _r, actingVersion, uint(NoLegsBlockLength)); err != nil {
				return err
			}
		}
	}
	if doRangeCheck {
		if err := s.RangeCheck(actingVersion, s.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (s *SecurityDefinitionRequest560) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if s.PartyDetailsListReqIDInActingVersion(actingVersion) {
		if s.PartyDetailsListReqID < s.PartyDetailsListReqIDMinValue() || s.PartyDetailsListReqID > s.PartyDetailsListReqIDMaxValue() {
			return fmt.Errorf("Range check failed on s.PartyDetailsListReqID (%v < %v > %v)", s.PartyDetailsListReqIDMinValue(), s.PartyDetailsListReqID, s.PartyDetailsListReqIDMaxValue())
		}
	}
	if s.SecurityReqIDInActingVersion(actingVersion) {
		if s.SecurityReqID < s.SecurityReqIDMinValue() || s.SecurityReqID > s.SecurityReqIDMaxValue() {
			return fmt.Errorf("Range check failed on s.SecurityReqID (%v < %v > %v)", s.SecurityReqIDMinValue(), s.SecurityReqID, s.SecurityReqIDMaxValue())
		}
	}
	if err := s.ManualOrderIndicator.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if s.SeqNumInActingVersion(actingVersion) {
		if s.SeqNum < s.SeqNumMinValue() || s.SeqNum > s.SeqNumMaxValue() {
			return fmt.Errorf("Range check failed on s.SeqNum (%v < %v > %v)", s.SeqNumMinValue(), s.SeqNum, s.SeqNumMaxValue())
		}
	}
	if s.SenderIDInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			if s.SenderID[idx] < s.SenderIDMinValue() || s.SenderID[idx] > s.SenderIDMaxValue() {
				return fmt.Errorf("Range check failed on s.SenderID[%d] (%v < %v > %v)", idx, s.SenderIDMinValue(), s.SenderID[idx], s.SenderIDMaxValue())
			}
		}
	}
	if s.SendingTimeEpochInActingVersion(actingVersion) {
		if s.SendingTimeEpoch < s.SendingTimeEpochMinValue() || s.SendingTimeEpoch > s.SendingTimeEpochMaxValue() {
			return fmt.Errorf("Range check failed on s.SendingTimeEpoch (%v < %v > %v)", s.SendingTimeEpochMinValue(), s.SendingTimeEpoch, s.SendingTimeEpochMaxValue())
		}
	}
	if s.SecuritySubTypeInActingVersion(actingVersion) {
		for idx := 0; idx < 8; idx++ {
			if s.SecuritySubType[idx] < s.SecuritySubTypeMinValue() || s.SecuritySubType[idx] > s.SecuritySubTypeMaxValue() {
				return fmt.Errorf("Range check failed on s.SecuritySubType[%d] (%v < %v > %v)", idx, s.SecuritySubTypeMinValue(), s.SecuritySubType[idx], s.SecuritySubTypeMaxValue())
			}
		}
	}
	if s.LocationInActingVersion(actingVersion) {
		for idx := 0; idx < 5; idx++ {
			if s.Location[idx] < s.LocationMinValue() || s.Location[idx] > s.LocationMaxValue() {
				return fmt.Errorf("Range check failed on s.Location[%d] (%v < %v > %v)", idx, s.LocationMinValue(), s.Location[idx], s.LocationMaxValue())
			}
		}
	}
	if s.StartDateInActingVersion(actingVersion) {
		if s.StartDate != s.StartDateNullValue() && (s.StartDate < s.StartDateMinValue() || s.StartDate > s.StartDateMaxValue()) {
			return fmt.Errorf("Range check failed on s.StartDate (%v < %v > %v)", s.StartDateMinValue(), s.StartDate, s.StartDateMaxValue())
		}
	}
	if s.EndDateInActingVersion(actingVersion) {
		if s.EndDate != s.EndDateNullValue() && (s.EndDate < s.EndDateMinValue() || s.EndDate > s.EndDateMaxValue()) {
			return fmt.Errorf("Range check failed on s.EndDate (%v < %v > %v)", s.EndDateMinValue(), s.EndDate, s.EndDateMaxValue())
		}
	}
	if s.MaxNoOfSubstitutionsInActingVersion(actingVersion) {
		if s.MaxNoOfSubstitutions != s.MaxNoOfSubstitutionsNullValue() && (s.MaxNoOfSubstitutions < s.MaxNoOfSubstitutionsMinValue() || s.MaxNoOfSubstitutions > s.MaxNoOfSubstitutionsMaxValue()) {
			return fmt.Errorf("Range check failed on s.MaxNoOfSubstitutions (%v < %v > %v)", s.MaxNoOfSubstitutionsMinValue(), s.MaxNoOfSubstitutions, s.MaxNoOfSubstitutionsMaxValue())
		}
	}
	if s.SourceRepoIDInActingVersion(actingVersion) {
		if s.SourceRepoID != s.SourceRepoIDNullValue() && (s.SourceRepoID < s.SourceRepoIDMinValue() || s.SourceRepoID > s.SourceRepoIDMaxValue()) {
			return fmt.Errorf("Range check failed on s.SourceRepoID (%v < %v > %v)", s.SourceRepoIDMinValue(), s.SourceRepoID, s.SourceRepoIDMaxValue())
		}
	}
	for _, prop := range s.NoLegs {
		if err := prop.RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	return nil
}

func SecurityDefinitionRequest560Init(s *SecurityDefinitionRequest560) {
	for idx := 0; idx < 20; idx++ {
		s.SenderID[idx] = 0
	}
	s.SecurityReqType[0] = 49
	s.StartDate = 65535
	s.EndDate = 65535
	s.MaxNoOfSubstitutions = 255
	s.SourceRepoID = 2147483647
	return
}

func (s *SecurityDefinitionRequest560NoLegs) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := s.LegPrice.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, s.LegSecurityID); err != nil {
		return err
	}
	if err := s.LegOptionDelta.Encode(_m, _w); err != nil {
		return err
	}
	if err := s.LegSide.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, s.LegRatioQty); err != nil {
		return err
	}
	return nil
}

func (s *SecurityDefinitionRequest560NoLegs) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	s.LegSecurityIDSource[0] = 56
	if s.LegPriceInActingVersion(actingVersion) {
		if err := s.LegPrice.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !s.LegSecurityIDInActingVersion(actingVersion) {
		s.LegSecurityID = s.LegSecurityIDNullValue()
	} else {
		if err := _m.ReadInt32(_r, &s.LegSecurityID); err != nil {
			return err
		}
	}
	if s.LegOptionDeltaInActingVersion(actingVersion) {
		if err := s.LegOptionDelta.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if s.LegSideInActingVersion(actingVersion) {
		if err := s.LegSide.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !s.LegRatioQtyInActingVersion(actingVersion) {
		s.LegRatioQty = s.LegRatioQtyNullValue()
	} else {
		if err := _m.ReadUint8(_r, &s.LegRatioQty); err != nil {
			return err
		}
	}
	if actingVersion > s.SbeSchemaVersion() && blockLength > s.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-s.SbeBlockLength()))
	}
	return nil
}

func (s *SecurityDefinitionRequest560NoLegs) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if s.LegSecurityIDInActingVersion(actingVersion) {
		if s.LegSecurityID < s.LegSecurityIDMinValue() || s.LegSecurityID > s.LegSecurityIDMaxValue() {
			return fmt.Errorf("Range check failed on s.LegSecurityID (%v < %v > %v)", s.LegSecurityIDMinValue(), s.LegSecurityID, s.LegSecurityIDMaxValue())
		}
	}
	if err := s.LegSide.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if s.LegRatioQtyInActingVersion(actingVersion) {
		if s.LegRatioQty != s.LegRatioQtyNullValue() && (s.LegRatioQty < s.LegRatioQtyMinValue() || s.LegRatioQty > s.LegRatioQtyMaxValue()) {
			return fmt.Errorf("Range check failed on s.LegRatioQty (%v < %v > %v)", s.LegRatioQtyMinValue(), s.LegRatioQty, s.LegRatioQtyMaxValue())
		}
	}
	return nil
}

func SecurityDefinitionRequest560NoLegsInit(s *SecurityDefinitionRequest560NoLegs) {
	s.LegSecurityIDSource[0] = 56
	s.LegRatioQty = 255
	return
}

func (*SecurityDefinitionRequest560) SbeBlockLength() (blockLength uint16) {
	return 71
}

func (*SecurityDefinitionRequest560) SbeTemplateId() (templateId uint16) {
	return 560
}

func (*SecurityDefinitionRequest560) SbeSchemaId() (schemaId uint16) {
	return 8
}

func (*SecurityDefinitionRequest560) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}

func (*SecurityDefinitionRequest560) SbeSemanticType() (semanticType []byte) {
	return []byte("c")
}

func (*SecurityDefinitionRequest560) PartyDetailsListReqIDId() uint16 {
	return 1505
}

func (*SecurityDefinitionRequest560) PartyDetailsListReqIDSinceVersion() uint16 {
	return 0
}

func (s *SecurityDefinitionRequest560) PartyDetailsListReqIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.PartyDetailsListReqIDSinceVersion()
}

func (*SecurityDefinitionRequest560) PartyDetailsListReqIDDeprecated() uint16 {
	return 0
}

func (*SecurityDefinitionRequest560) PartyDetailsListReqIDMetaAttribute(meta int) string {
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

func (*SecurityDefinitionRequest560) PartyDetailsListReqIDMinValue() uint64 {
	return 0
}

func (*SecurityDefinitionRequest560) PartyDetailsListReqIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*SecurityDefinitionRequest560) PartyDetailsListReqIDNullValue() uint64 {
	return math.MaxUint64
}

func (*SecurityDefinitionRequest560) SecurityReqIDId() uint16 {
	return 320
}

func (*SecurityDefinitionRequest560) SecurityReqIDSinceVersion() uint16 {
	return 0
}

func (s *SecurityDefinitionRequest560) SecurityReqIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.SecurityReqIDSinceVersion()
}

func (*SecurityDefinitionRequest560) SecurityReqIDDeprecated() uint16 {
	return 0
}

func (*SecurityDefinitionRequest560) SecurityReqIDMetaAttribute(meta int) string {
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

func (*SecurityDefinitionRequest560) SecurityReqIDMinValue() uint64 {
	return 0
}

func (*SecurityDefinitionRequest560) SecurityReqIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*SecurityDefinitionRequest560) SecurityReqIDNullValue() uint64 {
	return math.MaxUint64
}

func (*SecurityDefinitionRequest560) ManualOrderIndicatorId() uint16 {
	return 1028
}

func (*SecurityDefinitionRequest560) ManualOrderIndicatorSinceVersion() uint16 {
	return 0
}

func (s *SecurityDefinitionRequest560) ManualOrderIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.ManualOrderIndicatorSinceVersion()
}

func (*SecurityDefinitionRequest560) ManualOrderIndicatorDeprecated() uint16 {
	return 0
}

func (*SecurityDefinitionRequest560) ManualOrderIndicatorMetaAttribute(meta int) string {
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

func (*SecurityDefinitionRequest560) SeqNumId() uint16 {
	return 9726
}

func (*SecurityDefinitionRequest560) SeqNumSinceVersion() uint16 {
	return 0
}

func (s *SecurityDefinitionRequest560) SeqNumInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.SeqNumSinceVersion()
}

func (*SecurityDefinitionRequest560) SeqNumDeprecated() uint16 {
	return 0
}

func (*SecurityDefinitionRequest560) SeqNumMetaAttribute(meta int) string {
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

func (*SecurityDefinitionRequest560) SeqNumMinValue() uint32 {
	return 0
}

func (*SecurityDefinitionRequest560) SeqNumMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*SecurityDefinitionRequest560) SeqNumNullValue() uint32 {
	return math.MaxUint32
}

func (*SecurityDefinitionRequest560) SenderIDId() uint16 {
	return 5392
}

func (*SecurityDefinitionRequest560) SenderIDSinceVersion() uint16 {
	return 0
}

func (s *SecurityDefinitionRequest560) SenderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.SenderIDSinceVersion()
}

func (*SecurityDefinitionRequest560) SenderIDDeprecated() uint16 {
	return 0
}

func (*SecurityDefinitionRequest560) SenderIDMetaAttribute(meta int) string {
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

func (*SecurityDefinitionRequest560) SenderIDMinValue() byte {
	return byte(32)
}

func (*SecurityDefinitionRequest560) SenderIDMaxValue() byte {
	return byte(126)
}

func (*SecurityDefinitionRequest560) SenderIDNullValue() byte {
	return 0
}

func (s *SecurityDefinitionRequest560) SenderIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*SecurityDefinitionRequest560) SecurityReqTypeId() uint16 {
	return 321
}

func (*SecurityDefinitionRequest560) SecurityReqTypeSinceVersion() uint16 {
	return 0
}

func (s *SecurityDefinitionRequest560) SecurityReqTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.SecurityReqTypeSinceVersion()
}

func (*SecurityDefinitionRequest560) SecurityReqTypeDeprecated() uint16 {
	return 0
}

func (*SecurityDefinitionRequest560) SecurityReqTypeMetaAttribute(meta int) string {
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

func (*SecurityDefinitionRequest560) SecurityReqTypeMinValue() byte {
	return byte(32)
}

func (*SecurityDefinitionRequest560) SecurityReqTypeMaxValue() byte {
	return byte(126)
}

func (*SecurityDefinitionRequest560) SecurityReqTypeNullValue() byte {
	return 0
}

func (*SecurityDefinitionRequest560) SendingTimeEpochId() uint16 {
	return 5297
}

func (*SecurityDefinitionRequest560) SendingTimeEpochSinceVersion() uint16 {
	return 0
}

func (s *SecurityDefinitionRequest560) SendingTimeEpochInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.SendingTimeEpochSinceVersion()
}

func (*SecurityDefinitionRequest560) SendingTimeEpochDeprecated() uint16 {
	return 0
}

func (*SecurityDefinitionRequest560) SendingTimeEpochMetaAttribute(meta int) string {
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

func (*SecurityDefinitionRequest560) SendingTimeEpochMinValue() uint64 {
	return 0
}

func (*SecurityDefinitionRequest560) SendingTimeEpochMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*SecurityDefinitionRequest560) SendingTimeEpochNullValue() uint64 {
	return math.MaxUint64
}

func (*SecurityDefinitionRequest560) SecuritySubTypeId() uint16 {
	return 762
}

func (*SecurityDefinitionRequest560) SecuritySubTypeSinceVersion() uint16 {
	return 0
}

func (s *SecurityDefinitionRequest560) SecuritySubTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.SecuritySubTypeSinceVersion()
}

func (*SecurityDefinitionRequest560) SecuritySubTypeDeprecated() uint16 {
	return 0
}

func (*SecurityDefinitionRequest560) SecuritySubTypeMetaAttribute(meta int) string {
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

func (*SecurityDefinitionRequest560) SecuritySubTypeMinValue() byte {
	return byte(32)
}

func (*SecurityDefinitionRequest560) SecuritySubTypeMaxValue() byte {
	return byte(126)
}

func (*SecurityDefinitionRequest560) SecuritySubTypeNullValue() byte {
	return 0
}

func (s *SecurityDefinitionRequest560) SecuritySubTypeCharacterEncoding() string {
	return "US-ASCII"
}

func (*SecurityDefinitionRequest560) LocationId() uint16 {
	return 9537
}

func (*SecurityDefinitionRequest560) LocationSinceVersion() uint16 {
	return 0
}

func (s *SecurityDefinitionRequest560) LocationInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.LocationSinceVersion()
}

func (*SecurityDefinitionRequest560) LocationDeprecated() uint16 {
	return 0
}

func (*SecurityDefinitionRequest560) LocationMetaAttribute(meta int) string {
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

func (*SecurityDefinitionRequest560) LocationMinValue() byte {
	return byte(32)
}

func (*SecurityDefinitionRequest560) LocationMaxValue() byte {
	return byte(126)
}

func (*SecurityDefinitionRequest560) LocationNullValue() byte {
	return 0
}

func (s *SecurityDefinitionRequest560) LocationCharacterEncoding() string {
	return "US-ASCII"
}

func (*SecurityDefinitionRequest560) StartDateId() uint16 {
	return 916
}

func (*SecurityDefinitionRequest560) StartDateSinceVersion() uint16 {
	return 0
}

func (s *SecurityDefinitionRequest560) StartDateInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.StartDateSinceVersion()
}

func (*SecurityDefinitionRequest560) StartDateDeprecated() uint16 {
	return 0
}

func (*SecurityDefinitionRequest560) StartDateMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "LocalMktDate"
	case 4:
		return "optional"
	}
	return ""
}

func (*SecurityDefinitionRequest560) StartDateMinValue() uint16 {
	return 0
}

func (*SecurityDefinitionRequest560) StartDateMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*SecurityDefinitionRequest560) StartDateNullValue() uint16 {
	return 65535
}

func (*SecurityDefinitionRequest560) EndDateId() uint16 {
	return 917
}

func (*SecurityDefinitionRequest560) EndDateSinceVersion() uint16 {
	return 0
}

func (s *SecurityDefinitionRequest560) EndDateInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.EndDateSinceVersion()
}

func (*SecurityDefinitionRequest560) EndDateDeprecated() uint16 {
	return 0
}

func (*SecurityDefinitionRequest560) EndDateMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "LocalMktDate"
	case 4:
		return "optional"
	}
	return ""
}

func (*SecurityDefinitionRequest560) EndDateMinValue() uint16 {
	return 0
}

func (*SecurityDefinitionRequest560) EndDateMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*SecurityDefinitionRequest560) EndDateNullValue() uint16 {
	return 65535
}

func (*SecurityDefinitionRequest560) MaxNoOfSubstitutionsId() uint16 {
	return 37715
}

func (*SecurityDefinitionRequest560) MaxNoOfSubstitutionsSinceVersion() uint16 {
	return 0
}

func (s *SecurityDefinitionRequest560) MaxNoOfSubstitutionsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.MaxNoOfSubstitutionsSinceVersion()
}

func (*SecurityDefinitionRequest560) MaxNoOfSubstitutionsDeprecated() uint16 {
	return 0
}

func (*SecurityDefinitionRequest560) MaxNoOfSubstitutionsMetaAttribute(meta int) string {
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

func (*SecurityDefinitionRequest560) MaxNoOfSubstitutionsMinValue() uint8 {
	return 0
}

func (*SecurityDefinitionRequest560) MaxNoOfSubstitutionsMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*SecurityDefinitionRequest560) MaxNoOfSubstitutionsNullValue() uint8 {
	return 255
}

func (*SecurityDefinitionRequest560) SourceRepoIDId() uint16 {
	return 5677
}

func (*SecurityDefinitionRequest560) SourceRepoIDSinceVersion() uint16 {
	return 0
}

func (s *SecurityDefinitionRequest560) SourceRepoIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.SourceRepoIDSinceVersion()
}

func (*SecurityDefinitionRequest560) SourceRepoIDDeprecated() uint16 {
	return 0
}

func (*SecurityDefinitionRequest560) SourceRepoIDMetaAttribute(meta int) string {
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

func (*SecurityDefinitionRequest560) SourceRepoIDMinValue() int32 {
	return math.MinInt32 + 1
}

func (*SecurityDefinitionRequest560) SourceRepoIDMaxValue() int32 {
	return math.MaxInt32
}

func (*SecurityDefinitionRequest560) SourceRepoIDNullValue() int32 {
	return 2147483647
}

func (*SecurityDefinitionRequest560NoLegs) LegSecurityIDSourceId() uint16 {
	return 603
}

func (*SecurityDefinitionRequest560NoLegs) LegSecurityIDSourceSinceVersion() uint16 {
	return 0
}

func (s *SecurityDefinitionRequest560NoLegs) LegSecurityIDSourceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.LegSecurityIDSourceSinceVersion()
}

func (*SecurityDefinitionRequest560NoLegs) LegSecurityIDSourceDeprecated() uint16 {
	return 0
}

func (*SecurityDefinitionRequest560NoLegs) LegSecurityIDSourceMetaAttribute(meta int) string {
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

func (*SecurityDefinitionRequest560NoLegs) LegSecurityIDSourceMinValue() byte {
	return byte(32)
}

func (*SecurityDefinitionRequest560NoLegs) LegSecurityIDSourceMaxValue() byte {
	return byte(126)
}

func (*SecurityDefinitionRequest560NoLegs) LegSecurityIDSourceNullValue() byte {
	return 0
}

func (*SecurityDefinitionRequest560NoLegs) LegPriceId() uint16 {
	return 566
}

func (*SecurityDefinitionRequest560NoLegs) LegPriceSinceVersion() uint16 {
	return 0
}

func (s *SecurityDefinitionRequest560NoLegs) LegPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.LegPriceSinceVersion()
}

func (*SecurityDefinitionRequest560NoLegs) LegPriceDeprecated() uint16 {
	return 0
}

func (*SecurityDefinitionRequest560NoLegs) LegPriceMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "Price"
	case 4:
		return "required"
	}
	return ""
}

func (*SecurityDefinitionRequest560NoLegs) LegSecurityIDId() uint16 {
	return 602
}

func (*SecurityDefinitionRequest560NoLegs) LegSecurityIDSinceVersion() uint16 {
	return 0
}

func (s *SecurityDefinitionRequest560NoLegs) LegSecurityIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.LegSecurityIDSinceVersion()
}

func (*SecurityDefinitionRequest560NoLegs) LegSecurityIDDeprecated() uint16 {
	return 0
}

func (*SecurityDefinitionRequest560NoLegs) LegSecurityIDMetaAttribute(meta int) string {
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

func (*SecurityDefinitionRequest560NoLegs) LegSecurityIDMinValue() int32 {
	return math.MinInt32 + 1
}

func (*SecurityDefinitionRequest560NoLegs) LegSecurityIDMaxValue() int32 {
	return math.MaxInt32
}

func (*SecurityDefinitionRequest560NoLegs) LegSecurityIDNullValue() int32 {
	return math.MinInt32
}

func (*SecurityDefinitionRequest560NoLegs) LegOptionDeltaId() uint16 {
	return 1017
}

func (*SecurityDefinitionRequest560NoLegs) LegOptionDeltaSinceVersion() uint16 {
	return 0
}

func (s *SecurityDefinitionRequest560NoLegs) LegOptionDeltaInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.LegOptionDeltaSinceVersion()
}

func (*SecurityDefinitionRequest560NoLegs) LegOptionDeltaDeprecated() uint16 {
	return 0
}

func (*SecurityDefinitionRequest560NoLegs) LegOptionDeltaMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "float"
	case 4:
		return "required"
	}
	return ""
}

func (*SecurityDefinitionRequest560NoLegs) LegSideId() uint16 {
	return 624
}

func (*SecurityDefinitionRequest560NoLegs) LegSideSinceVersion() uint16 {
	return 0
}

func (s *SecurityDefinitionRequest560NoLegs) LegSideInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.LegSideSinceVersion()
}

func (*SecurityDefinitionRequest560NoLegs) LegSideDeprecated() uint16 {
	return 0
}

func (*SecurityDefinitionRequest560NoLegs) LegSideMetaAttribute(meta int) string {
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

func (*SecurityDefinitionRequest560NoLegs) LegRatioQtyId() uint16 {
	return 623
}

func (*SecurityDefinitionRequest560NoLegs) LegRatioQtySinceVersion() uint16 {
	return 0
}

func (s *SecurityDefinitionRequest560NoLegs) LegRatioQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.LegRatioQtySinceVersion()
}

func (*SecurityDefinitionRequest560NoLegs) LegRatioQtyDeprecated() uint16 {
	return 0
}

func (*SecurityDefinitionRequest560NoLegs) LegRatioQtyMetaAttribute(meta int) string {
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

func (*SecurityDefinitionRequest560NoLegs) LegRatioQtyMinValue() uint8 {
	return 0
}

func (*SecurityDefinitionRequest560NoLegs) LegRatioQtyMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*SecurityDefinitionRequest560NoLegs) LegRatioQtyNullValue() uint8 {
	return 255
}

func (*SecurityDefinitionRequest560) NoLegsId() uint16 {
	return 555
}

func (*SecurityDefinitionRequest560) NoLegsSinceVersion() uint16 {
	return 0
}

func (s *SecurityDefinitionRequest560) NoLegsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.NoLegsSinceVersion()
}

func (*SecurityDefinitionRequest560) NoLegsDeprecated() uint16 {
	return 0
}

func (*SecurityDefinitionRequest560NoLegs) SbeBlockLength() (blockLength uint) {
	return 19
}

func (*SecurityDefinitionRequest560NoLegs) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}
