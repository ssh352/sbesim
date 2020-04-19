package entity

import (
	"io"
	fix "sbe/sbe/iLinkBinary"
)

type SBEMessage interface {
	Decode(_m *fix.SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error
	Encode(_m *fix.SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error
}

