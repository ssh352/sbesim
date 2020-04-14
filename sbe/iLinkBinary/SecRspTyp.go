// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"reflect"
)

type SecRspTypEnum uint8
type SecRspTypValues struct {
	AcceptSecurityProposalasis                                 SecRspTypEnum
	AcceptSecurityproposalwithrevisionsasindicatedinthemessage SecRspTypEnum
	RejectSecurityProposal                                     SecRspTypEnum
	NullValue                                                  SecRspTypEnum
}

var SecRspTyp = SecRspTypValues{1, 2, 5, 255}

func (s SecRspTypEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(s)); err != nil {
		return err
	}
	return nil
}

func (s *SecRspTypEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(s)); err != nil {
		return err
	}
	return nil
}

func (s SecRspTypEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(SecRspTyp)
	for idx := 0; idx < value.NumField(); idx++ {
		if s == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on SecRspTyp, unknown enumeration value %d", s)
}

func (*SecRspTypEnum) EncodedLength() int64 {
	return 1
}

func (*SecRspTypEnum) AcceptSecurityProposalasisSinceVersion() uint16 {
	return 0
}

func (s *SecRspTypEnum) AcceptSecurityProposalasisInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.AcceptSecurityProposalasisSinceVersion()
}

func (*SecRspTypEnum) AcceptSecurityProposalasisDeprecated() uint16 {
	return 0
}

func (*SecRspTypEnum) AcceptSecurityproposalwithrevisionsasindicatedinthemessageSinceVersion() uint16 {
	return 0
}

func (s *SecRspTypEnum) AcceptSecurityproposalwithrevisionsasindicatedinthemessageInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.AcceptSecurityproposalwithrevisionsasindicatedinthemessageSinceVersion()
}

func (*SecRspTypEnum) AcceptSecurityproposalwithrevisionsasindicatedinthemessageDeprecated() uint16 {
	return 0
}

func (*SecRspTypEnum) RejectSecurityProposalSinceVersion() uint16 {
	return 0
}

func (s *SecRspTypEnum) RejectSecurityProposalInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.RejectSecurityProposalSinceVersion()
}

func (*SecRspTypEnum) RejectSecurityProposalDeprecated() uint16 {
	return 0
}
