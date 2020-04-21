package entity

import (
	"io"
	fix "sbe/sbe/iLinkBinary"
)
//go:generate mockgen -destination=../mock/entity/mock_SBEMessage.go -package=mockentity  sbe/entity SBEMessage
type SBEMessage interface {
	Decode(_m *fix.SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error
	Encode(_m *fix.SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error
	SbeBlockLength()  uint16
	SbeTemplateId()  uint16
	SbeSchemaId() uint16
	SbeSchemaVersion() uint16
}
