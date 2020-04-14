// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type SecurityDefinitionResponse561 struct {
	SeqNum                      uint32
	UUID                        uint64
	Text                        [256]byte
	FinancialInstrumentFullName [35]byte
	SenderID                    [20]byte
	Symbol                      [20]byte
	PartyDetailsListReqID       uint64
	SecurityReqID               uint64
	SecurityResponseID          uint64
	SendingTimeEpoch            uint64
	SecurityGroup               [6]byte
	SecurityType                [6]byte
	Location                    [5]byte
	SecurityID                  int32
	Currency                    [3]byte
	SecurityIDSource            [1]byte
	MaturityMonthYear           MaturityMonthYear
	DelayDuration               uint16
	StartDate                   uint16
	EndDate                     uint16
	MaxNoOfSubstitutions        uint8
	SourceRepoID                int32
	TerminationType             [8]byte
	SecurityResponseType        SecRspTypEnum
	UserDefinedInstrument       [1]byte
	ExpirationCycle             ExpCycleEnum
	ManualOrderIndicator        ManualOrdIndReqEnum
	SplitMsg                    SplitMsgEnum
	AutoQuoteRequest            BooleanFlagEnum
	PossRetransFlag             BooleanFlagEnum
	NoLegs                      []SecurityDefinitionResponse561NoLegs
}
type SecurityDefinitionResponse561NoLegs struct {
	LegPrice            PRICENULL9
	LegOptionDelta      Decimal32NULL
	LegSecurityIDSource [1]byte
	LegSecurityID       int32
	LegSide             SideReqEnum
	LegRatioQty         uint8
}

func (s *SecurityDefinitionResponse561) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := s.RangeCheck(s.SbeSchemaVersion(), s.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteUint32(_w, s.SeqNum); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, s.UUID); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, s.Text[:]); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, s.FinancialInstrumentFullName[:]); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, s.SenderID[:]); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, s.Symbol[:]); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, s.PartyDetailsListReqID); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, s.SecurityReqID); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, s.SecurityResponseID); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, s.SendingTimeEpoch); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, s.SecurityGroup[:]); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, s.SecurityType[:]); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, s.Location[:]); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, s.SecurityID); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, s.Currency[:]); err != nil {
		return err
	}
	if err := s.MaturityMonthYear.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, s.DelayDuration); err != nil {
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
	if err := _m.WriteBytes(_w, s.TerminationType[:]); err != nil {
		return err
	}
	if err := s.SecurityResponseType.Encode(_m, _w); err != nil {
		return err
	}
	if err := s.ExpirationCycle.Encode(_m, _w); err != nil {
		return err
	}
	if err := s.ManualOrderIndicator.Encode(_m, _w); err != nil {
		return err
	}
	if err := s.SplitMsg.Encode(_m, _w); err != nil {
		return err
	}
	if err := s.AutoQuoteRequest.Encode(_m, _w); err != nil {
		return err
	}
	if err := s.PossRetransFlag.Encode(_m, _w); err != nil {
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

func (s *SecurityDefinitionResponse561) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !s.SeqNumInActingVersion(actingVersion) {
		s.SeqNum = s.SeqNumNullValue()
	} else {
		if err := _m.ReadUint32(_r, &s.SeqNum); err != nil {
			return err
		}
	}
	if !s.UUIDInActingVersion(actingVersion) {
		s.UUID = s.UUIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &s.UUID); err != nil {
			return err
		}
	}
	if !s.TextInActingVersion(actingVersion) {
		for idx := 0; idx < 256; idx++ {
			s.Text[idx] = s.TextNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, s.Text[:]); err != nil {
			return err
		}
	}
	if !s.FinancialInstrumentFullNameInActingVersion(actingVersion) {
		for idx := 0; idx < 35; idx++ {
			s.FinancialInstrumentFullName[idx] = s.FinancialInstrumentFullNameNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, s.FinancialInstrumentFullName[:]); err != nil {
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
	if !s.SymbolInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			s.Symbol[idx] = s.SymbolNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, s.Symbol[:]); err != nil {
			return err
		}
	}
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
	if !s.SecurityResponseIDInActingVersion(actingVersion) {
		s.SecurityResponseID = s.SecurityResponseIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &s.SecurityResponseID); err != nil {
			return err
		}
	}
	if !s.SendingTimeEpochInActingVersion(actingVersion) {
		s.SendingTimeEpoch = s.SendingTimeEpochNullValue()
	} else {
		if err := _m.ReadUint64(_r, &s.SendingTimeEpoch); err != nil {
			return err
		}
	}
	if !s.SecurityGroupInActingVersion(actingVersion) {
		for idx := 0; idx < 6; idx++ {
			s.SecurityGroup[idx] = s.SecurityGroupNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, s.SecurityGroup[:]); err != nil {
			return err
		}
	}
	if !s.SecurityTypeInActingVersion(actingVersion) {
		for idx := 0; idx < 6; idx++ {
			s.SecurityType[idx] = s.SecurityTypeNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, s.SecurityType[:]); err != nil {
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
	if !s.SecurityIDInActingVersion(actingVersion) {
		s.SecurityID = s.SecurityIDNullValue()
	} else {
		if err := _m.ReadInt32(_r, &s.SecurityID); err != nil {
			return err
		}
	}
	if !s.CurrencyInActingVersion(actingVersion) {
		for idx := 0; idx < 3; idx++ {
			s.Currency[idx] = s.CurrencyNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, s.Currency[:]); err != nil {
			return err
		}
	}
	s.SecurityIDSource[0] = 56
	if s.MaturityMonthYearInActingVersion(actingVersion) {
		if err := s.MaturityMonthYear.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !s.DelayDurationInActingVersion(actingVersion) {
		s.DelayDuration = s.DelayDurationNullValue()
	} else {
		if err := _m.ReadUint16(_r, &s.DelayDuration); err != nil {
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
	if !s.TerminationTypeInActingVersion(actingVersion) {
		for idx := 0; idx < 8; idx++ {
			s.TerminationType[idx] = s.TerminationTypeNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, s.TerminationType[:]); err != nil {
			return err
		}
	}
	if s.SecurityResponseTypeInActingVersion(actingVersion) {
		if err := s.SecurityResponseType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	s.UserDefinedInstrument[0] = 89
	if s.ExpirationCycleInActingVersion(actingVersion) {
		if err := s.ExpirationCycle.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if s.ManualOrderIndicatorInActingVersion(actingVersion) {
		if err := s.ManualOrderIndicator.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if s.SplitMsgInActingVersion(actingVersion) {
		if err := s.SplitMsg.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if s.AutoQuoteRequestInActingVersion(actingVersion) {
		if err := s.AutoQuoteRequest.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if s.PossRetransFlagInActingVersion(actingVersion) {
		if err := s.PossRetransFlag.Decode(_m, _r, actingVersion); err != nil {
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
			s.NoLegs = make([]SecurityDefinitionResponse561NoLegs, NoLegsNumInGroup)
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

func (s *SecurityDefinitionResponse561) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if s.SeqNumInActingVersion(actingVersion) {
		if s.SeqNum < s.SeqNumMinValue() || s.SeqNum > s.SeqNumMaxValue() {
			return fmt.Errorf("Range check failed on s.SeqNum (%v < %v > %v)", s.SeqNumMinValue(), s.SeqNum, s.SeqNumMaxValue())
		}
	}
	if s.UUIDInActingVersion(actingVersion) {
		if s.UUID < s.UUIDMinValue() || s.UUID > s.UUIDMaxValue() {
			return fmt.Errorf("Range check failed on s.UUID (%v < %v > %v)", s.UUIDMinValue(), s.UUID, s.UUIDMaxValue())
		}
	}
	if s.TextInActingVersion(actingVersion) {
		for idx := 0; idx < 256; idx++ {
			if s.Text[idx] < s.TextMinValue() || s.Text[idx] > s.TextMaxValue() {
				return fmt.Errorf("Range check failed on s.Text[%d] (%v < %v > %v)", idx, s.TextMinValue(), s.Text[idx], s.TextMaxValue())
			}
		}
	}
	if s.FinancialInstrumentFullNameInActingVersion(actingVersion) {
		for idx := 0; idx < 35; idx++ {
			if s.FinancialInstrumentFullName[idx] < s.FinancialInstrumentFullNameMinValue() || s.FinancialInstrumentFullName[idx] > s.FinancialInstrumentFullNameMaxValue() {
				return fmt.Errorf("Range check failed on s.FinancialInstrumentFullName[%d] (%v < %v > %v)", idx, s.FinancialInstrumentFullNameMinValue(), s.FinancialInstrumentFullName[idx], s.FinancialInstrumentFullNameMaxValue())
			}
		}
	}
	if s.SenderIDInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			if s.SenderID[idx] < s.SenderIDMinValue() || s.SenderID[idx] > s.SenderIDMaxValue() {
				return fmt.Errorf("Range check failed on s.SenderID[%d] (%v < %v > %v)", idx, s.SenderIDMinValue(), s.SenderID[idx], s.SenderIDMaxValue())
			}
		}
	}
	if s.SymbolInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			if s.Symbol[idx] < s.SymbolMinValue() || s.Symbol[idx] > s.SymbolMaxValue() {
				return fmt.Errorf("Range check failed on s.Symbol[%d] (%v < %v > %v)", idx, s.SymbolMinValue(), s.Symbol[idx], s.SymbolMaxValue())
			}
		}
	}
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
	if s.SecurityResponseIDInActingVersion(actingVersion) {
		if s.SecurityResponseID < s.SecurityResponseIDMinValue() || s.SecurityResponseID > s.SecurityResponseIDMaxValue() {
			return fmt.Errorf("Range check failed on s.SecurityResponseID (%v < %v > %v)", s.SecurityResponseIDMinValue(), s.SecurityResponseID, s.SecurityResponseIDMaxValue())
		}
	}
	if s.SendingTimeEpochInActingVersion(actingVersion) {
		if s.SendingTimeEpoch < s.SendingTimeEpochMinValue() || s.SendingTimeEpoch > s.SendingTimeEpochMaxValue() {
			return fmt.Errorf("Range check failed on s.SendingTimeEpoch (%v < %v > %v)", s.SendingTimeEpochMinValue(), s.SendingTimeEpoch, s.SendingTimeEpochMaxValue())
		}
	}
	if s.SecurityGroupInActingVersion(actingVersion) {
		for idx := 0; idx < 6; idx++ {
			if s.SecurityGroup[idx] < s.SecurityGroupMinValue() || s.SecurityGroup[idx] > s.SecurityGroupMaxValue() {
				return fmt.Errorf("Range check failed on s.SecurityGroup[%d] (%v < %v > %v)", idx, s.SecurityGroupMinValue(), s.SecurityGroup[idx], s.SecurityGroupMaxValue())
			}
		}
	}
	if s.SecurityTypeInActingVersion(actingVersion) {
		for idx := 0; idx < 6; idx++ {
			if s.SecurityType[idx] < s.SecurityTypeMinValue() || s.SecurityType[idx] > s.SecurityTypeMaxValue() {
				return fmt.Errorf("Range check failed on s.SecurityType[%d] (%v < %v > %v)", idx, s.SecurityTypeMinValue(), s.SecurityType[idx], s.SecurityTypeMaxValue())
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
	if s.SecurityIDInActingVersion(actingVersion) {
		if s.SecurityID != s.SecurityIDNullValue() && (s.SecurityID < s.SecurityIDMinValue() || s.SecurityID > s.SecurityIDMaxValue()) {
			return fmt.Errorf("Range check failed on s.SecurityID (%v < %v > %v)", s.SecurityIDMinValue(), s.SecurityID, s.SecurityIDMaxValue())
		}
	}
	if s.CurrencyInActingVersion(actingVersion) {
		for idx := 0; idx < 3; idx++ {
			if s.Currency[idx] < s.CurrencyMinValue() || s.Currency[idx] > s.CurrencyMaxValue() {
				return fmt.Errorf("Range check failed on s.Currency[%d] (%v < %v > %v)", idx, s.CurrencyMinValue(), s.Currency[idx], s.CurrencyMaxValue())
			}
		}
	}
	if s.DelayDurationInActingVersion(actingVersion) {
		if s.DelayDuration != s.DelayDurationNullValue() && (s.DelayDuration < s.DelayDurationMinValue() || s.DelayDuration > s.DelayDurationMaxValue()) {
			return fmt.Errorf("Range check failed on s.DelayDuration (%v < %v > %v)", s.DelayDurationMinValue(), s.DelayDuration, s.DelayDurationMaxValue())
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
	if s.TerminationTypeInActingVersion(actingVersion) {
		for idx := 0; idx < 8; idx++ {
			if s.TerminationType[idx] < s.TerminationTypeMinValue() || s.TerminationType[idx] > s.TerminationTypeMaxValue() {
				return fmt.Errorf("Range check failed on s.TerminationType[%d] (%v < %v > %v)", idx, s.TerminationTypeMinValue(), s.TerminationType[idx], s.TerminationTypeMaxValue())
			}
		}
	}
	if err := s.SecurityResponseType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := s.ExpirationCycle.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := s.ManualOrderIndicator.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := s.SplitMsg.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := s.AutoQuoteRequest.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := s.PossRetransFlag.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	for _, prop := range s.NoLegs {
		if err := prop.RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	return nil
}

func SecurityDefinitionResponse561Init(s *SecurityDefinitionResponse561) {
	for idx := 0; idx < 256; idx++ {
		s.Text[idx] = 0
	}
	for idx := 0; idx < 35; idx++ {
		s.FinancialInstrumentFullName[idx] = 0
	}
	for idx := 0; idx < 20; idx++ {
		s.Symbol[idx] = 0
	}
	for idx := 0; idx < 6; idx++ {
		s.SecurityGroup[idx] = 0
	}
	for idx := 0; idx < 6; idx++ {
		s.SecurityType[idx] = 0
	}
	s.SecurityID = 2147483647
	for idx := 0; idx < 3; idx++ {
		s.Currency[idx] = 0
	}
	s.SecurityIDSource[0] = 56
	s.DelayDuration = 65535
	s.StartDate = 65535
	s.EndDate = 65535
	s.MaxNoOfSubstitutions = 255
	s.SourceRepoID = 2147483647
	for idx := 0; idx < 8; idx++ {
		s.TerminationType[idx] = 0
	}
	s.UserDefinedInstrument[0] = 89
	return
}

func (s *SecurityDefinitionResponse561NoLegs) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := s.LegPrice.Encode(_m, _w); err != nil {
		return err
	}
	if err := s.LegOptionDelta.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, s.LegSecurityID); err != nil {
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

func (s *SecurityDefinitionResponse561NoLegs) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	if s.LegPriceInActingVersion(actingVersion) {
		if err := s.LegPrice.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if s.LegOptionDeltaInActingVersion(actingVersion) {
		if err := s.LegOptionDelta.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	s.LegSecurityIDSource[0] = 56
	if !s.LegSecurityIDInActingVersion(actingVersion) {
		s.LegSecurityID = s.LegSecurityIDNullValue()
	} else {
		if err := _m.ReadInt32(_r, &s.LegSecurityID); err != nil {
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

func (s *SecurityDefinitionResponse561NoLegs) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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

func SecurityDefinitionResponse561NoLegsInit(s *SecurityDefinitionResponse561NoLegs) {
	s.LegSecurityIDSource[0] = 56
	s.LegRatioQty = 255
	return
}

func (*SecurityDefinitionResponse561) SbeBlockLength() (blockLength uint16) {
	return 429
}

func (*SecurityDefinitionResponse561) SbeTemplateId() (templateId uint16) {
	return 561
}

func (*SecurityDefinitionResponse561) SbeSchemaId() (schemaId uint16) {
	return 8
}

func (*SecurityDefinitionResponse561) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}

func (*SecurityDefinitionResponse561) SbeSemanticType() (semanticType []byte) {
	return []byte("d")
}

func (*SecurityDefinitionResponse561) SeqNumId() uint16 {
	return 9726
}

func (*SecurityDefinitionResponse561) SeqNumSinceVersion() uint16 {
	return 0
}

func (s *SecurityDefinitionResponse561) SeqNumInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.SeqNumSinceVersion()
}

func (*SecurityDefinitionResponse561) SeqNumDeprecated() uint16 {
	return 0
}

func (*SecurityDefinitionResponse561) SeqNumMetaAttribute(meta int) string {
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

func (*SecurityDefinitionResponse561) SeqNumMinValue() uint32 {
	return 0
}

func (*SecurityDefinitionResponse561) SeqNumMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*SecurityDefinitionResponse561) SeqNumNullValue() uint32 {
	return math.MaxUint32
}

func (*SecurityDefinitionResponse561) UUIDId() uint16 {
	return 39001
}

func (*SecurityDefinitionResponse561) UUIDSinceVersion() uint16 {
	return 0
}

func (s *SecurityDefinitionResponse561) UUIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.UUIDSinceVersion()
}

func (*SecurityDefinitionResponse561) UUIDDeprecated() uint16 {
	return 0
}

func (*SecurityDefinitionResponse561) UUIDMetaAttribute(meta int) string {
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

func (*SecurityDefinitionResponse561) UUIDMinValue() uint64 {
	return 0
}

func (*SecurityDefinitionResponse561) UUIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*SecurityDefinitionResponse561) UUIDNullValue() uint64 {
	return math.MaxUint64
}

func (*SecurityDefinitionResponse561) TextId() uint16 {
	return 58
}

func (*SecurityDefinitionResponse561) TextSinceVersion() uint16 {
	return 0
}

func (s *SecurityDefinitionResponse561) TextInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.TextSinceVersion()
}

func (*SecurityDefinitionResponse561) TextDeprecated() uint16 {
	return 0
}

func (*SecurityDefinitionResponse561) TextMetaAttribute(meta int) string {
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

func (*SecurityDefinitionResponse561) TextMinValue() byte {
	return byte(32)
}

func (*SecurityDefinitionResponse561) TextMaxValue() byte {
	return byte(126)
}

func (*SecurityDefinitionResponse561) TextNullValue() byte {
	return 0
}

func (s *SecurityDefinitionResponse561) TextCharacterEncoding() string {
	return "US-ASCII"
}

func (*SecurityDefinitionResponse561) FinancialInstrumentFullNameId() uint16 {
	return 2714
}

func (*SecurityDefinitionResponse561) FinancialInstrumentFullNameSinceVersion() uint16 {
	return 0
}

func (s *SecurityDefinitionResponse561) FinancialInstrumentFullNameInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.FinancialInstrumentFullNameSinceVersion()
}

func (*SecurityDefinitionResponse561) FinancialInstrumentFullNameDeprecated() uint16 {
	return 0
}

func (*SecurityDefinitionResponse561) FinancialInstrumentFullNameMetaAttribute(meta int) string {
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

func (*SecurityDefinitionResponse561) FinancialInstrumentFullNameMinValue() byte {
	return byte(32)
}

func (*SecurityDefinitionResponse561) FinancialInstrumentFullNameMaxValue() byte {
	return byte(126)
}

func (*SecurityDefinitionResponse561) FinancialInstrumentFullNameNullValue() byte {
	return 0
}

func (s *SecurityDefinitionResponse561) FinancialInstrumentFullNameCharacterEncoding() string {
	return "US-ASCII"
}

func (*SecurityDefinitionResponse561) SenderIDId() uint16 {
	return 5392
}

func (*SecurityDefinitionResponse561) SenderIDSinceVersion() uint16 {
	return 0
}

func (s *SecurityDefinitionResponse561) SenderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.SenderIDSinceVersion()
}

func (*SecurityDefinitionResponse561) SenderIDDeprecated() uint16 {
	return 0
}

func (*SecurityDefinitionResponse561) SenderIDMetaAttribute(meta int) string {
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

func (*SecurityDefinitionResponse561) SenderIDMinValue() byte {
	return byte(32)
}

func (*SecurityDefinitionResponse561) SenderIDMaxValue() byte {
	return byte(126)
}

func (*SecurityDefinitionResponse561) SenderIDNullValue() byte {
	return 0
}

func (s *SecurityDefinitionResponse561) SenderIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*SecurityDefinitionResponse561) SymbolId() uint16 {
	return 55
}

func (*SecurityDefinitionResponse561) SymbolSinceVersion() uint16 {
	return 0
}

func (s *SecurityDefinitionResponse561) SymbolInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.SymbolSinceVersion()
}

func (*SecurityDefinitionResponse561) SymbolDeprecated() uint16 {
	return 0
}

func (*SecurityDefinitionResponse561) SymbolMetaAttribute(meta int) string {
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

func (*SecurityDefinitionResponse561) SymbolMinValue() byte {
	return byte(32)
}

func (*SecurityDefinitionResponse561) SymbolMaxValue() byte {
	return byte(126)
}

func (*SecurityDefinitionResponse561) SymbolNullValue() byte {
	return 0
}

func (s *SecurityDefinitionResponse561) SymbolCharacterEncoding() string {
	return "US-ASCII"
}

func (*SecurityDefinitionResponse561) PartyDetailsListReqIDId() uint16 {
	return 1505
}

func (*SecurityDefinitionResponse561) PartyDetailsListReqIDSinceVersion() uint16 {
	return 0
}

func (s *SecurityDefinitionResponse561) PartyDetailsListReqIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.PartyDetailsListReqIDSinceVersion()
}

func (*SecurityDefinitionResponse561) PartyDetailsListReqIDDeprecated() uint16 {
	return 0
}

func (*SecurityDefinitionResponse561) PartyDetailsListReqIDMetaAttribute(meta int) string {
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

func (*SecurityDefinitionResponse561) PartyDetailsListReqIDMinValue() uint64 {
	return 0
}

func (*SecurityDefinitionResponse561) PartyDetailsListReqIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*SecurityDefinitionResponse561) PartyDetailsListReqIDNullValue() uint64 {
	return math.MaxUint64
}

func (*SecurityDefinitionResponse561) SecurityReqIDId() uint16 {
	return 320
}

func (*SecurityDefinitionResponse561) SecurityReqIDSinceVersion() uint16 {
	return 0
}

func (s *SecurityDefinitionResponse561) SecurityReqIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.SecurityReqIDSinceVersion()
}

func (*SecurityDefinitionResponse561) SecurityReqIDDeprecated() uint16 {
	return 0
}

func (*SecurityDefinitionResponse561) SecurityReqIDMetaAttribute(meta int) string {
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

func (*SecurityDefinitionResponse561) SecurityReqIDMinValue() uint64 {
	return 0
}

func (*SecurityDefinitionResponse561) SecurityReqIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*SecurityDefinitionResponse561) SecurityReqIDNullValue() uint64 {
	return math.MaxUint64
}

func (*SecurityDefinitionResponse561) SecurityResponseIDId() uint16 {
	return 322
}

func (*SecurityDefinitionResponse561) SecurityResponseIDSinceVersion() uint16 {
	return 0
}

func (s *SecurityDefinitionResponse561) SecurityResponseIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.SecurityResponseIDSinceVersion()
}

func (*SecurityDefinitionResponse561) SecurityResponseIDDeprecated() uint16 {
	return 0
}

func (*SecurityDefinitionResponse561) SecurityResponseIDMetaAttribute(meta int) string {
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

func (*SecurityDefinitionResponse561) SecurityResponseIDMinValue() uint64 {
	return 0
}

func (*SecurityDefinitionResponse561) SecurityResponseIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*SecurityDefinitionResponse561) SecurityResponseIDNullValue() uint64 {
	return math.MaxUint64
}

func (*SecurityDefinitionResponse561) SendingTimeEpochId() uint16 {
	return 5297
}

func (*SecurityDefinitionResponse561) SendingTimeEpochSinceVersion() uint16 {
	return 0
}

func (s *SecurityDefinitionResponse561) SendingTimeEpochInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.SendingTimeEpochSinceVersion()
}

func (*SecurityDefinitionResponse561) SendingTimeEpochDeprecated() uint16 {
	return 0
}

func (*SecurityDefinitionResponse561) SendingTimeEpochMetaAttribute(meta int) string {
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

func (*SecurityDefinitionResponse561) SendingTimeEpochMinValue() uint64 {
	return 0
}

func (*SecurityDefinitionResponse561) SendingTimeEpochMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*SecurityDefinitionResponse561) SendingTimeEpochNullValue() uint64 {
	return math.MaxUint64
}

func (*SecurityDefinitionResponse561) SecurityGroupId() uint16 {
	return 1151
}

func (*SecurityDefinitionResponse561) SecurityGroupSinceVersion() uint16 {
	return 0
}

func (s *SecurityDefinitionResponse561) SecurityGroupInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.SecurityGroupSinceVersion()
}

func (*SecurityDefinitionResponse561) SecurityGroupDeprecated() uint16 {
	return 0
}

func (*SecurityDefinitionResponse561) SecurityGroupMetaAttribute(meta int) string {
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

func (*SecurityDefinitionResponse561) SecurityGroupMinValue() byte {
	return byte(32)
}

func (*SecurityDefinitionResponse561) SecurityGroupMaxValue() byte {
	return byte(126)
}

func (*SecurityDefinitionResponse561) SecurityGroupNullValue() byte {
	return 0
}

func (s *SecurityDefinitionResponse561) SecurityGroupCharacterEncoding() string {
	return "US-ASCII"
}

func (*SecurityDefinitionResponse561) SecurityTypeId() uint16 {
	return 167
}

func (*SecurityDefinitionResponse561) SecurityTypeSinceVersion() uint16 {
	return 0
}

func (s *SecurityDefinitionResponse561) SecurityTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.SecurityTypeSinceVersion()
}

func (*SecurityDefinitionResponse561) SecurityTypeDeprecated() uint16 {
	return 0
}

func (*SecurityDefinitionResponse561) SecurityTypeMetaAttribute(meta int) string {
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

func (*SecurityDefinitionResponse561) SecurityTypeMinValue() byte {
	return byte(32)
}

func (*SecurityDefinitionResponse561) SecurityTypeMaxValue() byte {
	return byte(126)
}

func (*SecurityDefinitionResponse561) SecurityTypeNullValue() byte {
	return 0
}

func (s *SecurityDefinitionResponse561) SecurityTypeCharacterEncoding() string {
	return "US-ASCII"
}

func (*SecurityDefinitionResponse561) LocationId() uint16 {
	return 9537
}

func (*SecurityDefinitionResponse561) LocationSinceVersion() uint16 {
	return 0
}

func (s *SecurityDefinitionResponse561) LocationInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.LocationSinceVersion()
}

func (*SecurityDefinitionResponse561) LocationDeprecated() uint16 {
	return 0
}

func (*SecurityDefinitionResponse561) LocationMetaAttribute(meta int) string {
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

func (*SecurityDefinitionResponse561) LocationMinValue() byte {
	return byte(32)
}

func (*SecurityDefinitionResponse561) LocationMaxValue() byte {
	return byte(126)
}

func (*SecurityDefinitionResponse561) LocationNullValue() byte {
	return 0
}

func (s *SecurityDefinitionResponse561) LocationCharacterEncoding() string {
	return "US-ASCII"
}

func (*SecurityDefinitionResponse561) SecurityIDId() uint16 {
	return 48
}

func (*SecurityDefinitionResponse561) SecurityIDSinceVersion() uint16 {
	return 0
}

func (s *SecurityDefinitionResponse561) SecurityIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.SecurityIDSinceVersion()
}

func (*SecurityDefinitionResponse561) SecurityIDDeprecated() uint16 {
	return 0
}

func (*SecurityDefinitionResponse561) SecurityIDMetaAttribute(meta int) string {
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

func (*SecurityDefinitionResponse561) SecurityIDMinValue() int32 {
	return math.MinInt32 + 1
}

func (*SecurityDefinitionResponse561) SecurityIDMaxValue() int32 {
	return math.MaxInt32
}

func (*SecurityDefinitionResponse561) SecurityIDNullValue() int32 {
	return 2147483647
}

func (*SecurityDefinitionResponse561) CurrencyId() uint16 {
	return 15
}

func (*SecurityDefinitionResponse561) CurrencySinceVersion() uint16 {
	return 0
}

func (s *SecurityDefinitionResponse561) CurrencyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.CurrencySinceVersion()
}

func (*SecurityDefinitionResponse561) CurrencyDeprecated() uint16 {
	return 0
}

func (*SecurityDefinitionResponse561) CurrencyMetaAttribute(meta int) string {
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

func (*SecurityDefinitionResponse561) CurrencyMinValue() byte {
	return byte(32)
}

func (*SecurityDefinitionResponse561) CurrencyMaxValue() byte {
	return byte(126)
}

func (*SecurityDefinitionResponse561) CurrencyNullValue() byte {
	return 0
}

func (s *SecurityDefinitionResponse561) CurrencyCharacterEncoding() string {
	return "US-ASCII"
}

func (*SecurityDefinitionResponse561) SecurityIDSourceId() uint16 {
	return 22
}

func (*SecurityDefinitionResponse561) SecurityIDSourceSinceVersion() uint16 {
	return 0
}

func (s *SecurityDefinitionResponse561) SecurityIDSourceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.SecurityIDSourceSinceVersion()
}

func (*SecurityDefinitionResponse561) SecurityIDSourceDeprecated() uint16 {
	return 0
}

func (*SecurityDefinitionResponse561) SecurityIDSourceMetaAttribute(meta int) string {
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

func (*SecurityDefinitionResponse561) SecurityIDSourceMinValue() byte {
	return byte(32)
}

func (*SecurityDefinitionResponse561) SecurityIDSourceMaxValue() byte {
	return byte(126)
}

func (*SecurityDefinitionResponse561) SecurityIDSourceNullValue() byte {
	return 0
}

func (*SecurityDefinitionResponse561) MaturityMonthYearId() uint16 {
	return 200
}

func (*SecurityDefinitionResponse561) MaturityMonthYearSinceVersion() uint16 {
	return 0
}

func (s *SecurityDefinitionResponse561) MaturityMonthYearInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.MaturityMonthYearSinceVersion()
}

func (*SecurityDefinitionResponse561) MaturityMonthYearDeprecated() uint16 {
	return 0
}

func (*SecurityDefinitionResponse561) MaturityMonthYearMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "MonthYear"
	case 4:
		return "required"
	}
	return ""
}

func (*SecurityDefinitionResponse561) DelayDurationId() uint16 {
	return 5904
}

func (*SecurityDefinitionResponse561) DelayDurationSinceVersion() uint16 {
	return 0
}

func (s *SecurityDefinitionResponse561) DelayDurationInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.DelayDurationSinceVersion()
}

func (*SecurityDefinitionResponse561) DelayDurationDeprecated() uint16 {
	return 0
}

func (*SecurityDefinitionResponse561) DelayDurationMetaAttribute(meta int) string {
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

func (*SecurityDefinitionResponse561) DelayDurationMinValue() uint16 {
	return 0
}

func (*SecurityDefinitionResponse561) DelayDurationMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*SecurityDefinitionResponse561) DelayDurationNullValue() uint16 {
	return 65535
}

func (*SecurityDefinitionResponse561) StartDateId() uint16 {
	return 916
}

func (*SecurityDefinitionResponse561) StartDateSinceVersion() uint16 {
	return 0
}

func (s *SecurityDefinitionResponse561) StartDateInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.StartDateSinceVersion()
}

func (*SecurityDefinitionResponse561) StartDateDeprecated() uint16 {
	return 0
}

func (*SecurityDefinitionResponse561) StartDateMetaAttribute(meta int) string {
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

func (*SecurityDefinitionResponse561) StartDateMinValue() uint16 {
	return 0
}

func (*SecurityDefinitionResponse561) StartDateMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*SecurityDefinitionResponse561) StartDateNullValue() uint16 {
	return 65535
}

func (*SecurityDefinitionResponse561) EndDateId() uint16 {
	return 917
}

func (*SecurityDefinitionResponse561) EndDateSinceVersion() uint16 {
	return 0
}

func (s *SecurityDefinitionResponse561) EndDateInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.EndDateSinceVersion()
}

func (*SecurityDefinitionResponse561) EndDateDeprecated() uint16 {
	return 0
}

func (*SecurityDefinitionResponse561) EndDateMetaAttribute(meta int) string {
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

func (*SecurityDefinitionResponse561) EndDateMinValue() uint16 {
	return 0
}

func (*SecurityDefinitionResponse561) EndDateMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*SecurityDefinitionResponse561) EndDateNullValue() uint16 {
	return 65535
}

func (*SecurityDefinitionResponse561) MaxNoOfSubstitutionsId() uint16 {
	return 37715
}

func (*SecurityDefinitionResponse561) MaxNoOfSubstitutionsSinceVersion() uint16 {
	return 0
}

func (s *SecurityDefinitionResponse561) MaxNoOfSubstitutionsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.MaxNoOfSubstitutionsSinceVersion()
}

func (*SecurityDefinitionResponse561) MaxNoOfSubstitutionsDeprecated() uint16 {
	return 0
}

func (*SecurityDefinitionResponse561) MaxNoOfSubstitutionsMetaAttribute(meta int) string {
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

func (*SecurityDefinitionResponse561) MaxNoOfSubstitutionsMinValue() uint8 {
	return 0
}

func (*SecurityDefinitionResponse561) MaxNoOfSubstitutionsMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*SecurityDefinitionResponse561) MaxNoOfSubstitutionsNullValue() uint8 {
	return 255
}

func (*SecurityDefinitionResponse561) SourceRepoIDId() uint16 {
	return 5677
}

func (*SecurityDefinitionResponse561) SourceRepoIDSinceVersion() uint16 {
	return 0
}

func (s *SecurityDefinitionResponse561) SourceRepoIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.SourceRepoIDSinceVersion()
}

func (*SecurityDefinitionResponse561) SourceRepoIDDeprecated() uint16 {
	return 0
}

func (*SecurityDefinitionResponse561) SourceRepoIDMetaAttribute(meta int) string {
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

func (*SecurityDefinitionResponse561) SourceRepoIDMinValue() int32 {
	return math.MinInt32 + 1
}

func (*SecurityDefinitionResponse561) SourceRepoIDMaxValue() int32 {
	return math.MaxInt32
}

func (*SecurityDefinitionResponse561) SourceRepoIDNullValue() int32 {
	return 2147483647
}

func (*SecurityDefinitionResponse561) TerminationTypeId() uint16 {
	return 788
}

func (*SecurityDefinitionResponse561) TerminationTypeSinceVersion() uint16 {
	return 3
}

func (s *SecurityDefinitionResponse561) TerminationTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.TerminationTypeSinceVersion()
}

func (*SecurityDefinitionResponse561) TerminationTypeDeprecated() uint16 {
	return 0
}

func (*SecurityDefinitionResponse561) TerminationTypeMetaAttribute(meta int) string {
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

func (*SecurityDefinitionResponse561) TerminationTypeMinValue() byte {
	return byte(32)
}

func (*SecurityDefinitionResponse561) TerminationTypeMaxValue() byte {
	return byte(126)
}

func (*SecurityDefinitionResponse561) TerminationTypeNullValue() byte {
	return 0
}

func (s *SecurityDefinitionResponse561) TerminationTypeCharacterEncoding() string {
	return "US-ASCII"
}

func (*SecurityDefinitionResponse561) SecurityResponseTypeId() uint16 {
	return 323
}

func (*SecurityDefinitionResponse561) SecurityResponseTypeSinceVersion() uint16 {
	return 0
}

func (s *SecurityDefinitionResponse561) SecurityResponseTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.SecurityResponseTypeSinceVersion()
}

func (*SecurityDefinitionResponse561) SecurityResponseTypeDeprecated() uint16 {
	return 0
}

func (*SecurityDefinitionResponse561) SecurityResponseTypeMetaAttribute(meta int) string {
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

func (*SecurityDefinitionResponse561) UserDefinedInstrumentId() uint16 {
	return 9779
}

func (*SecurityDefinitionResponse561) UserDefinedInstrumentSinceVersion() uint16 {
	return 0
}

func (s *SecurityDefinitionResponse561) UserDefinedInstrumentInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.UserDefinedInstrumentSinceVersion()
}

func (*SecurityDefinitionResponse561) UserDefinedInstrumentDeprecated() uint16 {
	return 0
}

func (*SecurityDefinitionResponse561) UserDefinedInstrumentMetaAttribute(meta int) string {
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

func (*SecurityDefinitionResponse561) UserDefinedInstrumentMinValue() byte {
	return byte(32)
}

func (*SecurityDefinitionResponse561) UserDefinedInstrumentMaxValue() byte {
	return byte(126)
}

func (*SecurityDefinitionResponse561) UserDefinedInstrumentNullValue() byte {
	return 0
}

func (*SecurityDefinitionResponse561) ExpirationCycleId() uint16 {
	return 827
}

func (*SecurityDefinitionResponse561) ExpirationCycleSinceVersion() uint16 {
	return 0
}

func (s *SecurityDefinitionResponse561) ExpirationCycleInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.ExpirationCycleSinceVersion()
}

func (*SecurityDefinitionResponse561) ExpirationCycleDeprecated() uint16 {
	return 0
}

func (*SecurityDefinitionResponse561) ExpirationCycleMetaAttribute(meta int) string {
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

func (*SecurityDefinitionResponse561) ManualOrderIndicatorId() uint16 {
	return 1028
}

func (*SecurityDefinitionResponse561) ManualOrderIndicatorSinceVersion() uint16 {
	return 0
}

func (s *SecurityDefinitionResponse561) ManualOrderIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.ManualOrderIndicatorSinceVersion()
}

func (*SecurityDefinitionResponse561) ManualOrderIndicatorDeprecated() uint16 {
	return 0
}

func (*SecurityDefinitionResponse561) ManualOrderIndicatorMetaAttribute(meta int) string {
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

func (*SecurityDefinitionResponse561) SplitMsgId() uint16 {
	return 9553
}

func (*SecurityDefinitionResponse561) SplitMsgSinceVersion() uint16 {
	return 0
}

func (s *SecurityDefinitionResponse561) SplitMsgInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.SplitMsgSinceVersion()
}

func (*SecurityDefinitionResponse561) SplitMsgDeprecated() uint16 {
	return 0
}

func (*SecurityDefinitionResponse561) SplitMsgMetaAttribute(meta int) string {
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

func (*SecurityDefinitionResponse561) AutoQuoteRequestId() uint16 {
	return 9776
}

func (*SecurityDefinitionResponse561) AutoQuoteRequestSinceVersion() uint16 {
	return 0
}

func (s *SecurityDefinitionResponse561) AutoQuoteRequestInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.AutoQuoteRequestSinceVersion()
}

func (*SecurityDefinitionResponse561) AutoQuoteRequestDeprecated() uint16 {
	return 0
}

func (*SecurityDefinitionResponse561) AutoQuoteRequestMetaAttribute(meta int) string {
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

func (*SecurityDefinitionResponse561) PossRetransFlagId() uint16 {
	return 9765
}

func (*SecurityDefinitionResponse561) PossRetransFlagSinceVersion() uint16 {
	return 0
}

func (s *SecurityDefinitionResponse561) PossRetransFlagInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.PossRetransFlagSinceVersion()
}

func (*SecurityDefinitionResponse561) PossRetransFlagDeprecated() uint16 {
	return 0
}

func (*SecurityDefinitionResponse561) PossRetransFlagMetaAttribute(meta int) string {
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

func (*SecurityDefinitionResponse561NoLegs) LegPriceId() uint16 {
	return 566
}

func (*SecurityDefinitionResponse561NoLegs) LegPriceSinceVersion() uint16 {
	return 0
}

func (s *SecurityDefinitionResponse561NoLegs) LegPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.LegPriceSinceVersion()
}

func (*SecurityDefinitionResponse561NoLegs) LegPriceDeprecated() uint16 {
	return 0
}

func (*SecurityDefinitionResponse561NoLegs) LegPriceMetaAttribute(meta int) string {
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

func (*SecurityDefinitionResponse561NoLegs) LegOptionDeltaId() uint16 {
	return 1017
}

func (*SecurityDefinitionResponse561NoLegs) LegOptionDeltaSinceVersion() uint16 {
	return 0
}

func (s *SecurityDefinitionResponse561NoLegs) LegOptionDeltaInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.LegOptionDeltaSinceVersion()
}

func (*SecurityDefinitionResponse561NoLegs) LegOptionDeltaDeprecated() uint16 {
	return 0
}

func (*SecurityDefinitionResponse561NoLegs) LegOptionDeltaMetaAttribute(meta int) string {
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

func (*SecurityDefinitionResponse561NoLegs) LegSecurityIDSourceId() uint16 {
	return 603
}

func (*SecurityDefinitionResponse561NoLegs) LegSecurityIDSourceSinceVersion() uint16 {
	return 0
}

func (s *SecurityDefinitionResponse561NoLegs) LegSecurityIDSourceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.LegSecurityIDSourceSinceVersion()
}

func (*SecurityDefinitionResponse561NoLegs) LegSecurityIDSourceDeprecated() uint16 {
	return 0
}

func (*SecurityDefinitionResponse561NoLegs) LegSecurityIDSourceMetaAttribute(meta int) string {
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

func (*SecurityDefinitionResponse561NoLegs) LegSecurityIDSourceMinValue() byte {
	return byte(32)
}

func (*SecurityDefinitionResponse561NoLegs) LegSecurityIDSourceMaxValue() byte {
	return byte(126)
}

func (*SecurityDefinitionResponse561NoLegs) LegSecurityIDSourceNullValue() byte {
	return 0
}

func (*SecurityDefinitionResponse561NoLegs) LegSecurityIDId() uint16 {
	return 602
}

func (*SecurityDefinitionResponse561NoLegs) LegSecurityIDSinceVersion() uint16 {
	return 0
}

func (s *SecurityDefinitionResponse561NoLegs) LegSecurityIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.LegSecurityIDSinceVersion()
}

func (*SecurityDefinitionResponse561NoLegs) LegSecurityIDDeprecated() uint16 {
	return 0
}

func (*SecurityDefinitionResponse561NoLegs) LegSecurityIDMetaAttribute(meta int) string {
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

func (*SecurityDefinitionResponse561NoLegs) LegSecurityIDMinValue() int32 {
	return math.MinInt32 + 1
}

func (*SecurityDefinitionResponse561NoLegs) LegSecurityIDMaxValue() int32 {
	return math.MaxInt32
}

func (*SecurityDefinitionResponse561NoLegs) LegSecurityIDNullValue() int32 {
	return math.MinInt32
}

func (*SecurityDefinitionResponse561NoLegs) LegSideId() uint16 {
	return 624
}

func (*SecurityDefinitionResponse561NoLegs) LegSideSinceVersion() uint16 {
	return 0
}

func (s *SecurityDefinitionResponse561NoLegs) LegSideInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.LegSideSinceVersion()
}

func (*SecurityDefinitionResponse561NoLegs) LegSideDeprecated() uint16 {
	return 0
}

func (*SecurityDefinitionResponse561NoLegs) LegSideMetaAttribute(meta int) string {
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

func (*SecurityDefinitionResponse561NoLegs) LegRatioQtyId() uint16 {
	return 623
}

func (*SecurityDefinitionResponse561NoLegs) LegRatioQtySinceVersion() uint16 {
	return 0
}

func (s *SecurityDefinitionResponse561NoLegs) LegRatioQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.LegRatioQtySinceVersion()
}

func (*SecurityDefinitionResponse561NoLegs) LegRatioQtyDeprecated() uint16 {
	return 0
}

func (*SecurityDefinitionResponse561NoLegs) LegRatioQtyMetaAttribute(meta int) string {
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

func (*SecurityDefinitionResponse561NoLegs) LegRatioQtyMinValue() uint8 {
	return 0
}

func (*SecurityDefinitionResponse561NoLegs) LegRatioQtyMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*SecurityDefinitionResponse561NoLegs) LegRatioQtyNullValue() uint8 {
	return 255
}

func (*SecurityDefinitionResponse561) NoLegsId() uint16 {
	return 555
}

func (*SecurityDefinitionResponse561) NoLegsSinceVersion() uint16 {
	return 0
}

func (s *SecurityDefinitionResponse561) NoLegsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.NoLegsSinceVersion()
}

func (*SecurityDefinitionResponse561) NoLegsDeprecated() uint16 {
	return 0
}

func (*SecurityDefinitionResponse561NoLegs) SbeBlockLength() (blockLength uint) {
	return 19
}

func (*SecurityDefinitionResponse561NoLegs) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}
