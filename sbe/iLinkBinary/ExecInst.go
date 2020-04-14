// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"io"
)

type ExecInst [8]bool
type ExecInstChoiceValue uint8
type ExecInstChoiceValues struct {
	AON       ExecInstChoiceValue
	OB        ExecInstChoiceValue
	NH        ExecInstChoiceValue
	Reserved1 ExecInstChoiceValue
	Reserved2 ExecInstChoiceValue
	Reserved3 ExecInstChoiceValue
	Reserved4 ExecInstChoiceValue
	Reserved5 ExecInstChoiceValue
}

var ExecInstChoice = ExecInstChoiceValues{0, 1, 2, 3, 4, 5, 6, 7}

func (e *ExecInst) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	var wireval uint8 = 0
	for k, v := range e {
		if v {
			wireval |= (1 << uint(k))
		}
	}
	return _m.WriteUint8(_w, wireval)
}

func (e *ExecInst) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	var wireval uint8

	if err := _m.ReadUint8(_r, &wireval); err != nil {
		return err
	}

	var idx uint
	for idx = 0; idx < 8; idx++ {
		e[idx] = (wireval & (1 << idx)) > 0
	}
	return nil
}

func (ExecInst) EncodedLength() int64 {
	return 1
}

func (*ExecInst) AONSinceVersion() uint16 {
	return 0
}

func (e *ExecInst) AONInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.AONSinceVersion()
}

func (*ExecInst) AONDeprecated() uint16 {
	return 0
}

func (*ExecInst) OBSinceVersion() uint16 {
	return 0
}

func (e *ExecInst) OBInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OBSinceVersion()
}

func (*ExecInst) OBDeprecated() uint16 {
	return 0
}

func (*ExecInst) NHSinceVersion() uint16 {
	return 0
}

func (e *ExecInst) NHInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.NHSinceVersion()
}

func (*ExecInst) NHDeprecated() uint16 {
	return 0
}

func (*ExecInst) Reserved1SinceVersion() uint16 {
	return 0
}

func (e *ExecInst) Reserved1InActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.Reserved1SinceVersion()
}

func (*ExecInst) Reserved1Deprecated() uint16 {
	return 0
}

func (*ExecInst) Reserved2SinceVersion() uint16 {
	return 0
}

func (e *ExecInst) Reserved2InActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.Reserved2SinceVersion()
}

func (*ExecInst) Reserved2Deprecated() uint16 {
	return 0
}

func (*ExecInst) Reserved3SinceVersion() uint16 {
	return 0
}

func (e *ExecInst) Reserved3InActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.Reserved3SinceVersion()
}

func (*ExecInst) Reserved3Deprecated() uint16 {
	return 0
}

func (*ExecInst) Reserved4SinceVersion() uint16 {
	return 0
}

func (e *ExecInst) Reserved4InActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.Reserved4SinceVersion()
}

func (*ExecInst) Reserved4Deprecated() uint16 {
	return 0
}

func (*ExecInst) Reserved5SinceVersion() uint16 {
	return 0
}

func (e *ExecInst) Reserved5InActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.Reserved5SinceVersion()
}

func (*ExecInst) Reserved5Deprecated() uint16 {
	return 0
}
