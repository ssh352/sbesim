// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type RetransmitRequest508 struct {
	UUID             uint64
	LastUUID         uint64
	RequestTimestamp uint64
	FromSeqNo        uint32
	MsgCount         uint16
}

func (r *RetransmitRequest508) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := r.RangeCheck(r.SbeSchemaVersion(), r.SbeSchemaVersion()); err != nil {
			return err
		}
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
	if err := _m.WriteUint32(_w, r.FromSeqNo); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, r.MsgCount); err != nil {
		return err
	}
	return nil
}

func (r *RetransmitRequest508) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
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
	if !r.FromSeqNoInActingVersion(actingVersion) {
		r.FromSeqNo = r.FromSeqNoNullValue()
	} else {
		if err := _m.ReadUint32(_r, &r.FromSeqNo); err != nil {
			return err
		}
	}
	if !r.MsgCountInActingVersion(actingVersion) {
		r.MsgCount = r.MsgCountNullValue()
	} else {
		if err := _m.ReadUint16(_r, &r.MsgCount); err != nil {
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

func (r *RetransmitRequest508) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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
	if r.FromSeqNoInActingVersion(actingVersion) {
		if r.FromSeqNo < r.FromSeqNoMinValue() || r.FromSeqNo > r.FromSeqNoMaxValue() {
			return fmt.Errorf("Range check failed on r.FromSeqNo (%v < %v > %v)", r.FromSeqNoMinValue(), r.FromSeqNo, r.FromSeqNoMaxValue())
		}
	}
	if r.MsgCountInActingVersion(actingVersion) {
		if r.MsgCount < r.MsgCountMinValue() || r.MsgCount > r.MsgCountMaxValue() {
			return fmt.Errorf("Range check failed on r.MsgCount (%v < %v > %v)", r.MsgCountMinValue(), r.MsgCount, r.MsgCountMaxValue())
		}
	}
	return nil
}

func RetransmitRequest508Init(r *RetransmitRequest508) {
	r.LastUUID = 18446744073709551615
	return
}

func (*RetransmitRequest508) SbeBlockLength() (blockLength uint16) {
	return 30
}

func (*RetransmitRequest508) SbeTemplateId() (templateId uint16) {
	return 508
}

func (*RetransmitRequest508) SbeSchemaId() (schemaId uint16) {
	return 8
}

func (*RetransmitRequest508) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}

func (*RetransmitRequest508) SbeSemanticType() (semanticType []byte) {
	return []byte("RetransmitRequest")
}

func (*RetransmitRequest508) UUIDId() uint16 {
	return 39001
}

func (*RetransmitRequest508) UUIDSinceVersion() uint16 {
	return 0
}

func (r *RetransmitRequest508) UUIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.UUIDSinceVersion()
}

func (*RetransmitRequest508) UUIDDeprecated() uint16 {
	return 0
}

func (*RetransmitRequest508) UUIDMetaAttribute(meta int) string {
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

func (*RetransmitRequest508) UUIDMinValue() uint64 {
	return 0
}

func (*RetransmitRequest508) UUIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*RetransmitRequest508) UUIDNullValue() uint64 {
	return math.MaxUint64
}

func (*RetransmitRequest508) LastUUIDId() uint16 {
	return 39017
}

func (*RetransmitRequest508) LastUUIDSinceVersion() uint16 {
	return 0
}

func (r *RetransmitRequest508) LastUUIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.LastUUIDSinceVersion()
}

func (*RetransmitRequest508) LastUUIDDeprecated() uint16 {
	return 0
}

func (*RetransmitRequest508) LastUUIDMetaAttribute(meta int) string {
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

func (*RetransmitRequest508) LastUUIDMinValue() uint64 {
	return 0
}

func (*RetransmitRequest508) LastUUIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*RetransmitRequest508) LastUUIDNullValue() uint64 {
	return 18446744073709551615
}

func (*RetransmitRequest508) RequestTimestampId() uint16 {
	return 39002
}

func (*RetransmitRequest508) RequestTimestampSinceVersion() uint16 {
	return 0
}

func (r *RetransmitRequest508) RequestTimestampInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.RequestTimestampSinceVersion()
}

func (*RetransmitRequest508) RequestTimestampDeprecated() uint16 {
	return 0
}

func (*RetransmitRequest508) RequestTimestampMetaAttribute(meta int) string {
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

func (*RetransmitRequest508) RequestTimestampMinValue() uint64 {
	return 0
}

func (*RetransmitRequest508) RequestTimestampMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*RetransmitRequest508) RequestTimestampNullValue() uint64 {
	return math.MaxUint64
}

func (*RetransmitRequest508) FromSeqNoId() uint16 {
	return 39018
}

func (*RetransmitRequest508) FromSeqNoSinceVersion() uint16 {
	return 0
}

func (r *RetransmitRequest508) FromSeqNoInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.FromSeqNoSinceVersion()
}

func (*RetransmitRequest508) FromSeqNoDeprecated() uint16 {
	return 0
}

func (*RetransmitRequest508) FromSeqNoMetaAttribute(meta int) string {
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

func (*RetransmitRequest508) FromSeqNoMinValue() uint32 {
	return 0
}

func (*RetransmitRequest508) FromSeqNoMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*RetransmitRequest508) FromSeqNoNullValue() uint32 {
	return math.MaxUint32
}

func (*RetransmitRequest508) MsgCountId() uint16 {
	return 39019
}

func (*RetransmitRequest508) MsgCountSinceVersion() uint16 {
	return 0
}

func (r *RetransmitRequest508) MsgCountInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.MsgCountSinceVersion()
}

func (*RetransmitRequest508) MsgCountDeprecated() uint16 {
	return 0
}

func (*RetransmitRequest508) MsgCountMetaAttribute(meta int) string {
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

func (*RetransmitRequest508) MsgCountMinValue() uint16 {
	return 0
}

func (*RetransmitRequest508) MsgCountMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*RetransmitRequest508) MsgCountNullValue() uint16 {
	return math.MaxUint16
}
