// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type Terminate507 struct {
	Reason           [48]byte
	UUID             uint64
	RequestTimestamp uint64
	ErrorCodes       uint16
	SplitMsg         SplitMsgEnum
}

func (t *Terminate507) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := t.RangeCheck(t.SbeSchemaVersion(), t.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteBytes(_w, t.Reason[:]); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, t.UUID); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, t.RequestTimestamp); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, t.ErrorCodes); err != nil {
		return err
	}
	if err := t.SplitMsg.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (t *Terminate507) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !t.ReasonInActingVersion(actingVersion) {
		for idx := 0; idx < 48; idx++ {
			t.Reason[idx] = t.ReasonNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, t.Reason[:]); err != nil {
			return err
		}
	}
	if !t.UUIDInActingVersion(actingVersion) {
		t.UUID = t.UUIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &t.UUID); err != nil {
			return err
		}
	}
	if !t.RequestTimestampInActingVersion(actingVersion) {
		t.RequestTimestamp = t.RequestTimestampNullValue()
	} else {
		if err := _m.ReadUint64(_r, &t.RequestTimestamp); err != nil {
			return err
		}
	}
	if !t.ErrorCodesInActingVersion(actingVersion) {
		t.ErrorCodes = t.ErrorCodesNullValue()
	} else {
		if err := _m.ReadUint16(_r, &t.ErrorCodes); err != nil {
			return err
		}
	}
	if t.SplitMsgInActingVersion(actingVersion) {
		if err := t.SplitMsg.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > t.SbeSchemaVersion() && blockLength > t.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-t.SbeBlockLength()))
	}
	if doRangeCheck {
		if err := t.RangeCheck(actingVersion, t.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (t *Terminate507) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if t.ReasonInActingVersion(actingVersion) {
		for idx := 0; idx < 48; idx++ {
			if t.Reason[idx] < t.ReasonMinValue() || t.Reason[idx] > t.ReasonMaxValue() {
				return fmt.Errorf("Range check failed on t.Reason[%d] (%v < %v > %v)", idx, t.ReasonMinValue(), t.Reason[idx], t.ReasonMaxValue())
			}
		}
	}
	if t.UUIDInActingVersion(actingVersion) {
		if t.UUID < t.UUIDMinValue() || t.UUID > t.UUIDMaxValue() {
			return fmt.Errorf("Range check failed on t.UUID (%v < %v > %v)", t.UUIDMinValue(), t.UUID, t.UUIDMaxValue())
		}
	}
	if t.RequestTimestampInActingVersion(actingVersion) {
		if t.RequestTimestamp < t.RequestTimestampMinValue() || t.RequestTimestamp > t.RequestTimestampMaxValue() {
			return fmt.Errorf("Range check failed on t.RequestTimestamp (%v < %v > %v)", t.RequestTimestampMinValue(), t.RequestTimestamp, t.RequestTimestampMaxValue())
		}
	}
	if t.ErrorCodesInActingVersion(actingVersion) {
		if t.ErrorCodes < t.ErrorCodesMinValue() || t.ErrorCodes > t.ErrorCodesMaxValue() {
			return fmt.Errorf("Range check failed on t.ErrorCodes (%v < %v > %v)", t.ErrorCodesMinValue(), t.ErrorCodes, t.ErrorCodesMaxValue())
		}
	}
	if err := t.SplitMsg.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	return nil
}

func Terminate507Init(t *Terminate507) {
	for idx := 0; idx < 48; idx++ {
		t.Reason[idx] = 0
	}
	return
}

func (*Terminate507) SbeBlockLength() (blockLength uint16) {
	return 67
}

func (*Terminate507) SbeTemplateId() (templateId uint16) {
	return 507
}

func (*Terminate507) SbeSchemaId() (schemaId uint16) {
	return 8
}

func (*Terminate507) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}

func (*Terminate507) SbeSemanticType() (semanticType []byte) {
	return []byte("Terminate")
}

func (*Terminate507) ReasonId() uint16 {
	return 39011
}

func (*Terminate507) ReasonSinceVersion() uint16 {
	return 0
}

func (t *Terminate507) ReasonInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.ReasonSinceVersion()
}

func (*Terminate507) ReasonDeprecated() uint16 {
	return 0
}

func (*Terminate507) ReasonMetaAttribute(meta int) string {
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

func (*Terminate507) ReasonMinValue() byte {
	return byte(32)
}

func (*Terminate507) ReasonMaxValue() byte {
	return byte(126)
}

func (*Terminate507) ReasonNullValue() byte {
	return 0
}

func (t *Terminate507) ReasonCharacterEncoding() string {
	return "US-ASCII"
}

func (*Terminate507) UUIDId() uint16 {
	return 39001
}

func (*Terminate507) UUIDSinceVersion() uint16 {
	return 0
}

func (t *Terminate507) UUIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.UUIDSinceVersion()
}

func (*Terminate507) UUIDDeprecated() uint16 {
	return 0
}

func (*Terminate507) UUIDMetaAttribute(meta int) string {
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

func (*Terminate507) UUIDMinValue() uint64 {
	return 0
}

func (*Terminate507) UUIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*Terminate507) UUIDNullValue() uint64 {
	return math.MaxUint64
}

func (*Terminate507) RequestTimestampId() uint16 {
	return 39002
}

func (*Terminate507) RequestTimestampSinceVersion() uint16 {
	return 0
}

func (t *Terminate507) RequestTimestampInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.RequestTimestampSinceVersion()
}

func (*Terminate507) RequestTimestampDeprecated() uint16 {
	return 0
}

func (*Terminate507) RequestTimestampMetaAttribute(meta int) string {
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

func (*Terminate507) RequestTimestampMinValue() uint64 {
	return 0
}

func (*Terminate507) RequestTimestampMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*Terminate507) RequestTimestampNullValue() uint64 {
	return math.MaxUint64
}

func (*Terminate507) ErrorCodesId() uint16 {
	return 39012
}

func (*Terminate507) ErrorCodesSinceVersion() uint16 {
	return 0
}

func (t *Terminate507) ErrorCodesInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.ErrorCodesSinceVersion()
}

func (*Terminate507) ErrorCodesDeprecated() uint16 {
	return 0
}

func (*Terminate507) ErrorCodesMetaAttribute(meta int) string {
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

func (*Terminate507) ErrorCodesMinValue() uint16 {
	return 0
}

func (*Terminate507) ErrorCodesMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*Terminate507) ErrorCodesNullValue() uint16 {
	return math.MaxUint16
}

func (*Terminate507) SplitMsgId() uint16 {
	return 9553
}

func (*Terminate507) SplitMsgSinceVersion() uint16 {
	return 0
}

func (t *Terminate507) SplitMsgInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.SplitMsgSinceVersion()
}

func (*Terminate507) SplitMsgDeprecated() uint16 {
	return 0
}

func (*Terminate507) SplitMsgMetaAttribute(meta int) string {
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
