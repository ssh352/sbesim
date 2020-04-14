// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type RetransmitReject510 struct {
	Reason           [48]byte
	UUID             uint64
	LastUUID         uint64
	RequestTimestamp uint64
	ErrorCodes       uint16
	SplitMsg         SplitMsgEnum
}

func (r *RetransmitReject510) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := r.RangeCheck(r.SbeSchemaVersion(), r.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteBytes(_w, r.Reason[:]); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, r.UUID); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, r.LastUUID); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, r.RequestTimestamp); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, r.ErrorCodes); err != nil {
		return err
	}
	if err := r.SplitMsg.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (r *RetransmitReject510) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !r.ReasonInActingVersion(actingVersion) {
		for idx := 0; idx < 48; idx++ {
			r.Reason[idx] = r.ReasonNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, r.Reason[:]); err != nil {
			return err
		}
	}
	if !r.UUIDInActingVersion(actingVersion) {
		r.UUID = r.UUIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &r.UUID); err != nil {
			return err
		}
	}
	if !r.LastUUIDInActingVersion(actingVersion) {
		r.LastUUID = r.LastUUIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &r.LastUUID); err != nil {
			return err
		}
	}
	if !r.RequestTimestampInActingVersion(actingVersion) {
		r.RequestTimestamp = r.RequestTimestampNullValue()
	} else {
		if err := _m.ReadUint64(_r, &r.RequestTimestamp); err != nil {
			return err
		}
	}
	if !r.ErrorCodesInActingVersion(actingVersion) {
		r.ErrorCodes = r.ErrorCodesNullValue()
	} else {
		if err := _m.ReadUint16(_r, &r.ErrorCodes); err != nil {
			return err
		}
	}
	if r.SplitMsgInActingVersion(actingVersion) {
		if err := r.SplitMsg.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > r.SbeSchemaVersion() && blockLength > r.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-r.SbeBlockLength()))
	}
	if doRangeCheck {
		if err := r.RangeCheck(actingVersion, r.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (r *RetransmitReject510) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if r.ReasonInActingVersion(actingVersion) {
		for idx := 0; idx < 48; idx++ {
			if r.Reason[idx] < r.ReasonMinValue() || r.Reason[idx] > r.ReasonMaxValue() {
				return fmt.Errorf("Range check failed on r.Reason[%d] (%v < %v > %v)", idx, r.ReasonMinValue(), r.Reason[idx], r.ReasonMaxValue())
			}
		}
	}
	if r.UUIDInActingVersion(actingVersion) {
		if r.UUID < r.UUIDMinValue() || r.UUID > r.UUIDMaxValue() {
			return fmt.Errorf("Range check failed on r.UUID (%v < %v > %v)", r.UUIDMinValue(), r.UUID, r.UUIDMaxValue())
		}
	}
	if r.LastUUIDInActingVersion(actingVersion) {
		if r.LastUUID != r.LastUUIDNullValue() && (r.LastUUID < r.LastUUIDMinValue() || r.LastUUID > r.LastUUIDMaxValue()) {
			return fmt.Errorf("Range check failed on r.LastUUID (%v < %v > %v)", r.LastUUIDMinValue(), r.LastUUID, r.LastUUIDMaxValue())
		}
	}
	if r.RequestTimestampInActingVersion(actingVersion) {
		if r.RequestTimestamp < r.RequestTimestampMinValue() || r.RequestTimestamp > r.RequestTimestampMaxValue() {
			return fmt.Errorf("Range check failed on r.RequestTimestamp (%v < %v > %v)", r.RequestTimestampMinValue(), r.RequestTimestamp, r.RequestTimestampMaxValue())
		}
	}
	if r.ErrorCodesInActingVersion(actingVersion) {
		if r.ErrorCodes < r.ErrorCodesMinValue() || r.ErrorCodes > r.ErrorCodesMaxValue() {
			return fmt.Errorf("Range check failed on r.ErrorCodes (%v < %v > %v)", r.ErrorCodesMinValue(), r.ErrorCodes, r.ErrorCodesMaxValue())
		}
	}
	if err := r.SplitMsg.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	return nil
}

func RetransmitReject510Init(r *RetransmitReject510) {
	for idx := 0; idx < 48; idx++ {
		r.Reason[idx] = 0
	}
	r.LastUUID = 18446744073709551615
	return
}

func (*RetransmitReject510) SbeBlockLength() (blockLength uint16) {
	return 75
}

func (*RetransmitReject510) SbeTemplateId() (templateId uint16) {
	return 510
}

func (*RetransmitReject510) SbeSchemaId() (schemaId uint16) {
	return 8
}

func (*RetransmitReject510) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}

func (*RetransmitReject510) SbeSemanticType() (semanticType []byte) {
	return []byte("RetransmitReject")
}

func (*RetransmitReject510) ReasonId() uint16 {
	return 39011
}

func (*RetransmitReject510) ReasonSinceVersion() uint16 {
	return 0
}

func (r *RetransmitReject510) ReasonInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.ReasonSinceVersion()
}

func (*RetransmitReject510) ReasonDeprecated() uint16 {
	return 0
}

func (*RetransmitReject510) ReasonMetaAttribute(meta int) string {
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

func (*RetransmitReject510) ReasonMinValue() byte {
	return byte(32)
}

func (*RetransmitReject510) ReasonMaxValue() byte {
	return byte(126)
}

func (*RetransmitReject510) ReasonNullValue() byte {
	return 0
}

func (r *RetransmitReject510) ReasonCharacterEncoding() string {
	return "US-ASCII"
}

func (*RetransmitReject510) UUIDId() uint16 {
	return 39001
}

func (*RetransmitReject510) UUIDSinceVersion() uint16 {
	return 0
}

func (r *RetransmitReject510) UUIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.UUIDSinceVersion()
}

func (*RetransmitReject510) UUIDDeprecated() uint16 {
	return 0
}

func (*RetransmitReject510) UUIDMetaAttribute(meta int) string {
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

func (*RetransmitReject510) UUIDMinValue() uint64 {
	return 0
}

func (*RetransmitReject510) UUIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*RetransmitReject510) UUIDNullValue() uint64 {
	return math.MaxUint64
}

func (*RetransmitReject510) LastUUIDId() uint16 {
	return 39017
}

func (*RetransmitReject510) LastUUIDSinceVersion() uint16 {
	return 0
}

func (r *RetransmitReject510) LastUUIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.LastUUIDSinceVersion()
}

func (*RetransmitReject510) LastUUIDDeprecated() uint16 {
	return 0
}

func (*RetransmitReject510) LastUUIDMetaAttribute(meta int) string {
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

func (*RetransmitReject510) LastUUIDMinValue() uint64 {
	return 0
}

func (*RetransmitReject510) LastUUIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*RetransmitReject510) LastUUIDNullValue() uint64 {
	return 18446744073709551615
}

func (*RetransmitReject510) RequestTimestampId() uint16 {
	return 39002
}

func (*RetransmitReject510) RequestTimestampSinceVersion() uint16 {
	return 0
}

func (r *RetransmitReject510) RequestTimestampInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.RequestTimestampSinceVersion()
}

func (*RetransmitReject510) RequestTimestampDeprecated() uint16 {
	return 0
}

func (*RetransmitReject510) RequestTimestampMetaAttribute(meta int) string {
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

func (*RetransmitReject510) RequestTimestampMinValue() uint64 {
	return 0
}

func (*RetransmitReject510) RequestTimestampMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*RetransmitReject510) RequestTimestampNullValue() uint64 {
	return math.MaxUint64
}

func (*RetransmitReject510) ErrorCodesId() uint16 {
	return 39012
}

func (*RetransmitReject510) ErrorCodesSinceVersion() uint16 {
	return 0
}

func (r *RetransmitReject510) ErrorCodesInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.ErrorCodesSinceVersion()
}

func (*RetransmitReject510) ErrorCodesDeprecated() uint16 {
	return 0
}

func (*RetransmitReject510) ErrorCodesMetaAttribute(meta int) string {
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

func (*RetransmitReject510) ErrorCodesMinValue() uint16 {
	return 0
}

func (*RetransmitReject510) ErrorCodesMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*RetransmitReject510) ErrorCodesNullValue() uint16 {
	return math.MaxUint16
}

func (*RetransmitReject510) SplitMsgId() uint16 {
	return 9553
}

func (*RetransmitReject510) SplitMsgSinceVersion() uint16 {
	return 0
}

func (r *RetransmitReject510) SplitMsgInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.SplitMsgSinceVersion()
}

func (*RetransmitReject510) SplitMsgDeprecated() uint16 {
	return 0
}

func (*RetransmitReject510) SplitMsgMetaAttribute(meta int) string {
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
