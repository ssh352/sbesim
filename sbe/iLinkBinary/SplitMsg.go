// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"reflect"
)

type SplitMsgEnum uint8
type SplitMsgValues struct {
	SplitMessageDelayed      SplitMsgEnum
	OutofOrderMessageDelayed SplitMsgEnum
	CompleteMessageDelayed   SplitMsgEnum
	NullValue                SplitMsgEnum
}

var SplitMsg = SplitMsgValues{0, 1, 2, 255}

func (s SplitMsgEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(s)); err != nil {
		return err
	}
	return nil
}

func (s *SplitMsgEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(s)); err != nil {
		return err
	}
	return nil
}

func (s SplitMsgEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(SplitMsg)
	for idx := 0; idx < value.NumField(); idx++ {
		if s == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on SplitMsg, unknown enumeration value %d", s)
}

func (*SplitMsgEnum) EncodedLength() int64 {
	return 1
}

func (*SplitMsgEnum) SplitMessageDelayedSinceVersion() uint16 {
	return 0
}

func (s *SplitMsgEnum) SplitMessageDelayedInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.SplitMessageDelayedSinceVersion()
}

func (*SplitMsgEnum) SplitMessageDelayedDeprecated() uint16 {
	return 0
}

func (*SplitMsgEnum) OutofOrderMessageDelayedSinceVersion() uint16 {
	return 0
}

func (s *SplitMsgEnum) OutofOrderMessageDelayedInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.OutofOrderMessageDelayedSinceVersion()
}

func (*SplitMsgEnum) OutofOrderMessageDelayedDeprecated() uint16 {
	return 0
}

func (*SplitMsgEnum) CompleteMessageDelayedSinceVersion() uint16 {
	return 0
}

func (s *SplitMsgEnum) CompleteMessageDelayedInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.CompleteMessageDelayedSinceVersion()
}

func (*SplitMsgEnum) CompleteMessageDelayedDeprecated() uint16 {
	return 0
}
