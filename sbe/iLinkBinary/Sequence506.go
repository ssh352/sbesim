// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type Sequence506 struct {
	UUID                    uint64
	NextSeqNo               uint32
	FaultToleranceIndicator FTIEnum
	KeepAliveIntervalLapsed KeepAliveLapsedEnum
}

func (s *Sequence506) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := s.RangeCheck(s.SbeSchemaVersion(), s.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteUint64(_w, s.UUID); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, s.NextSeqNo); err != nil {
		return err
	}
	if err := s.FaultToleranceIndicator.Encode(_m, _w); err != nil {
		return err
	}
	if err := s.KeepAliveIntervalLapsed.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (s *Sequence506) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !s.UUIDInActingVersion(actingVersion) {
		s.UUID = s.UUIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &s.UUID); err != nil {
			return err
		}
	}
	if !s.NextSeqNoInActingVersion(actingVersion) {
		s.NextSeqNo = s.NextSeqNoNullValue()
	} else {
		if err := _m.ReadUint32(_r, &s.NextSeqNo); err != nil {
			return err
		}
	}
	if s.FaultToleranceIndicatorInActingVersion(actingVersion) {
		if err := s.FaultToleranceIndicator.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if s.KeepAliveIntervalLapsedInActingVersion(actingVersion) {
		if err := s.KeepAliveIntervalLapsed.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > s.SbeSchemaVersion() && blockLength > s.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-s.SbeBlockLength()))
	}
	if doRangeCheck {
		if err := s.RangeCheck(actingVersion, s.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (s *Sequence506) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if s.UUIDInActingVersion(actingVersion) {
		if s.UUID < s.UUIDMinValue() || s.UUID > s.UUIDMaxValue() {
			return fmt.Errorf("Range check failed on s.UUID (%v < %v > %v)", s.UUIDMinValue(), s.UUID, s.UUIDMaxValue())
		}
	}
	if s.NextSeqNoInActingVersion(actingVersion) {
		if s.NextSeqNo < s.NextSeqNoMinValue() || s.NextSeqNo > s.NextSeqNoMaxValue() {
			return fmt.Errorf("Range check failed on s.NextSeqNo (%v < %v > %v)", s.NextSeqNoMinValue(), s.NextSeqNo, s.NextSeqNoMaxValue())
		}
	}
	if err := s.FaultToleranceIndicator.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := s.KeepAliveIntervalLapsed.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	return nil
}

func Sequence506Init(s *Sequence506) {
	return
}

func (*Sequence506) SbeBlockLength() (blockLength uint16) {
	return 14
}

func (*Sequence506) SbeTemplateId() (templateId uint16) {
	return 506
}

func (*Sequence506) SbeSchemaId() (schemaId uint16) {
	return 8
}

func (*Sequence506) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}

func (*Sequence506) SbeSemanticType() (semanticType []byte) {
	return []byte("Sequence")
}

func (*Sequence506) UUIDId() uint16 {
	return 39001
}

func (*Sequence506) UUIDSinceVersion() uint16 {
	return 0
}

func (s *Sequence506) UUIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.UUIDSinceVersion()
}

func (*Sequence506) UUIDDeprecated() uint16 {
	return 0
}

func (*Sequence506) UUIDMetaAttribute(meta int) string {
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

func (*Sequence506) UUIDMinValue() uint64 {
	return 0
}

func (*Sequence506) UUIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*Sequence506) UUIDNullValue() uint64 {
	return math.MaxUint64
}

func (*Sequence506) NextSeqNoId() uint16 {
	return 39013
}

func (*Sequence506) NextSeqNoSinceVersion() uint16 {
	return 0
}

func (s *Sequence506) NextSeqNoInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.NextSeqNoSinceVersion()
}

func (*Sequence506) NextSeqNoDeprecated() uint16 {
	return 0
}

func (*Sequence506) NextSeqNoMetaAttribute(meta int) string {
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

func (*Sequence506) NextSeqNoMinValue() uint32 {
	return 0
}

func (*Sequence506) NextSeqNoMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*Sequence506) NextSeqNoNullValue() uint32 {
	return math.MaxUint32
}

func (*Sequence506) FaultToleranceIndicatorId() uint16 {
	return 39010
}

func (*Sequence506) FaultToleranceIndicatorSinceVersion() uint16 {
	return 0
}

func (s *Sequence506) FaultToleranceIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.FaultToleranceIndicatorSinceVersion()
}

func (*Sequence506) FaultToleranceIndicatorDeprecated() uint16 {
	return 0
}

func (*Sequence506) FaultToleranceIndicatorMetaAttribute(meta int) string {
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

func (*Sequence506) KeepAliveIntervalLapsedId() uint16 {
	return 39016
}

func (*Sequence506) KeepAliveIntervalLapsedSinceVersion() uint16 {
	return 0
}

func (s *Sequence506) KeepAliveIntervalLapsedInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.KeepAliveIntervalLapsedSinceVersion()
}

func (*Sequence506) KeepAliveIntervalLapsedDeprecated() uint16 {
	return 0
}

func (*Sequence506) KeepAliveIntervalLapsedMetaAttribute(meta int) string {
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
