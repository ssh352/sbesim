// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type NotApplied513 struct {
	UUID      uint64
	FromSeqNo uint32
	MsgCount  uint32
	SplitMsg  SplitMsgEnum
}

func (n *NotApplied513) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := n.RangeCheck(n.SbeSchemaVersion(), n.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteUint64(_w, n.UUID); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, n.FromSeqNo); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, n.MsgCount); err != nil {
		return err
	}
	if err := n.SplitMsg.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (n *NotApplied513) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !n.UUIDInActingVersion(actingVersion) {
		n.UUID = n.UUIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &n.UUID); err != nil {
			return err
		}
	}
	if !n.FromSeqNoInActingVersion(actingVersion) {
		n.FromSeqNo = n.FromSeqNoNullValue()
	} else {
		if err := _m.ReadUint32(_r, &n.FromSeqNo); err != nil {
			return err
		}
	}
	if !n.MsgCountInActingVersion(actingVersion) {
		n.MsgCount = n.MsgCountNullValue()
	} else {
		if err := _m.ReadUint32(_r, &n.MsgCount); err != nil {
			return err
		}
	}
	if n.SplitMsgInActingVersion(actingVersion) {
		if err := n.SplitMsg.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > n.SbeSchemaVersion() && blockLength > n.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-n.SbeBlockLength()))
	}
	if doRangeCheck {
		if err := n.RangeCheck(actingVersion, n.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (n *NotApplied513) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if n.UUIDInActingVersion(actingVersion) {
		if n.UUID < n.UUIDMinValue() || n.UUID > n.UUIDMaxValue() {
			return fmt.Errorf("Range check failed on n.UUID (%v < %v > %v)", n.UUIDMinValue(), n.UUID, n.UUIDMaxValue())
		}
	}
	if n.FromSeqNoInActingVersion(actingVersion) {
		if n.FromSeqNo < n.FromSeqNoMinValue() || n.FromSeqNo > n.FromSeqNoMaxValue() {
			return fmt.Errorf("Range check failed on n.FromSeqNo (%v < %v > %v)", n.FromSeqNoMinValue(), n.FromSeqNo, n.FromSeqNoMaxValue())
		}
	}
	if n.MsgCountInActingVersion(actingVersion) {
		if n.MsgCount < n.MsgCountMinValue() || n.MsgCount > n.MsgCountMaxValue() {
			return fmt.Errorf("Range check failed on n.MsgCount (%v < %v > %v)", n.MsgCountMinValue(), n.MsgCount, n.MsgCountMaxValue())
		}
	}
	if err := n.SplitMsg.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	return nil
}

func NotApplied513Init(n *NotApplied513) {
	return
}

func (*NotApplied513) SbeBlockLength() (blockLength uint16) {
	return 17
}

func (*NotApplied513) SbeTemplateId() (templateId uint16) {
	return 513
}

func (*NotApplied513) SbeSchemaId() (schemaId uint16) {
	return 8
}

func (*NotApplied513) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}

func (*NotApplied513) SbeSemanticType() (semanticType []byte) {
	return []byte("NotApplied")
}

func (*NotApplied513) UUIDId() uint16 {
	return 39001
}

func (*NotApplied513) UUIDSinceVersion() uint16 {
	return 0
}

func (n *NotApplied513) UUIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.UUIDSinceVersion()
}

func (*NotApplied513) UUIDDeprecated() uint16 {
	return 0
}

func (*NotApplied513) UUIDMetaAttribute(meta int) string {
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

func (*NotApplied513) UUIDMinValue() uint64 {
	return 0
}

func (*NotApplied513) UUIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*NotApplied513) UUIDNullValue() uint64 {
	return math.MaxUint64
}

func (*NotApplied513) FromSeqNoId() uint16 {
	return 39018
}

func (*NotApplied513) FromSeqNoSinceVersion() uint16 {
	return 0
}

func (n *NotApplied513) FromSeqNoInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.FromSeqNoSinceVersion()
}

func (*NotApplied513) FromSeqNoDeprecated() uint16 {
	return 0
}

func (*NotApplied513) FromSeqNoMetaAttribute(meta int) string {
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

func (*NotApplied513) FromSeqNoMinValue() uint32 {
	return 0
}

func (*NotApplied513) FromSeqNoMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*NotApplied513) FromSeqNoNullValue() uint32 {
	return math.MaxUint32
}

func (*NotApplied513) MsgCountId() uint16 {
	return 39019
}

func (*NotApplied513) MsgCountSinceVersion() uint16 {
	return 0
}

func (n *NotApplied513) MsgCountInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.MsgCountSinceVersion()
}

func (*NotApplied513) MsgCountDeprecated() uint16 {
	return 0
}

func (*NotApplied513) MsgCountMetaAttribute(meta int) string {
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

func (*NotApplied513) MsgCountMinValue() uint32 {
	return 0
}

func (*NotApplied513) MsgCountMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*NotApplied513) MsgCountNullValue() uint32 {
	return math.MaxUint32
}

func (*NotApplied513) SplitMsgId() uint16 {
	return 9553
}

func (*NotApplied513) SplitMsgSinceVersion() uint16 {
	return 0
}

func (n *NotApplied513) SplitMsgInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.SplitMsgSinceVersion()
}

func (*NotApplied513) SplitMsgDeprecated() uint16 {
	return 0
}

func (*NotApplied513) SplitMsgMetaAttribute(meta int) string {
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
