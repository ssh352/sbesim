// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"reflect"
)

type PartyDetailRoleEnum uint16
type PartyDetailRoleValues struct {
	ExecutingFirm   PartyDetailRoleEnum
	CustomerAccount PartyDetailRoleEnum
	TakeUpFirm      PartyDetailRoleEnum
	Operator        PartyDetailRoleEnum
	TakeUpAccount   PartyDetailRoleEnum
	NullValue       PartyDetailRoleEnum
}

var PartyDetailRole = PartyDetailRoleValues{1, 24, 96, 118, 1000, 65535}

func (p PartyDetailRoleEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint16(_w, uint16(p)); err != nil {
		return err
	}
	return nil
}

func (p *PartyDetailRoleEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint16(_r, (*uint16)(p)); err != nil {
		return err
	}
	return nil
}

func (p PartyDetailRoleEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(PartyDetailRole)
	for idx := 0; idx < value.NumField(); idx++ {
		if p == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on PartyDetailRole, unknown enumeration value %d", p)
}

func (*PartyDetailRoleEnum) EncodedLength() int64 {
	return 2
}

func (*PartyDetailRoleEnum) ExecutingFirmSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailRoleEnum) ExecutingFirmInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.ExecutingFirmSinceVersion()
}

func (*PartyDetailRoleEnum) ExecutingFirmDeprecated() uint16 {
	return 0
}

func (*PartyDetailRoleEnum) CustomerAccountSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailRoleEnum) CustomerAccountInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.CustomerAccountSinceVersion()
}

func (*PartyDetailRoleEnum) CustomerAccountDeprecated() uint16 {
	return 0
}

func (*PartyDetailRoleEnum) TakeUpFirmSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailRoleEnum) TakeUpFirmInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.TakeUpFirmSinceVersion()
}

func (*PartyDetailRoleEnum) TakeUpFirmDeprecated() uint16 {
	return 0
}

func (*PartyDetailRoleEnum) OperatorSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailRoleEnum) OperatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.OperatorSinceVersion()
}

func (*PartyDetailRoleEnum) OperatorDeprecated() uint16 {
	return 0
}

func (*PartyDetailRoleEnum) TakeUpAccountSinceVersion() uint16 {
	return 0
}

func (p *PartyDetailRoleEnum) TakeUpAccountInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.TakeUpAccountSinceVersion()
}

func (*PartyDetailRoleEnum) TakeUpAccountDeprecated() uint16 {
	return 0
}
