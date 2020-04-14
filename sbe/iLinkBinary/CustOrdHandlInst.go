// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"reflect"
)

type CustOrdHandlInstEnum byte
type CustOrdHandlInstValues struct {
	FCMprovidedscreen   CustOrdHandlInstEnum
	Otherprovidedscreen CustOrdHandlInstEnum
	FCMAPIorFIX         CustOrdHandlInstEnum
	AlgoEngine          CustOrdHandlInstEnum
	DeskElectronic      CustOrdHandlInstEnum
	ClientElectronic    CustOrdHandlInstEnum
	NullValue           CustOrdHandlInstEnum
}

var CustOrdHandlInst = CustOrdHandlInstValues{67, 68, 71, 72, 87, 89, 0}

func (c CustOrdHandlInstEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, byte(c)); err != nil {
		return err
	}
	return nil
}

func (c *CustOrdHandlInstEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*byte)(c)); err != nil {
		return err
	}
	return nil
}

func (c CustOrdHandlInstEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(CustOrdHandlInst)
	for idx := 0; idx < value.NumField(); idx++ {
		if c == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on CustOrdHandlInst, unknown enumeration value %d", c)
}

func (*CustOrdHandlInstEnum) EncodedLength() int64 {
	return 1
}

func (*CustOrdHandlInstEnum) FCMprovidedscreenSinceVersion() uint16 {
	return 0
}

func (c *CustOrdHandlInstEnum) FCMprovidedscreenInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.FCMprovidedscreenSinceVersion()
}

func (*CustOrdHandlInstEnum) FCMprovidedscreenDeprecated() uint16 {
	return 0
}

func (*CustOrdHandlInstEnum) OtherprovidedscreenSinceVersion() uint16 {
	return 0
}

func (c *CustOrdHandlInstEnum) OtherprovidedscreenInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.OtherprovidedscreenSinceVersion()
}

func (*CustOrdHandlInstEnum) OtherprovidedscreenDeprecated() uint16 {
	return 0
}

func (*CustOrdHandlInstEnum) FCMAPIorFIXSinceVersion() uint16 {
	return 0
}

func (c *CustOrdHandlInstEnum) FCMAPIorFIXInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.FCMAPIorFIXSinceVersion()
}

func (*CustOrdHandlInstEnum) FCMAPIorFIXDeprecated() uint16 {
	return 0
}

func (*CustOrdHandlInstEnum) AlgoEngineSinceVersion() uint16 {
	return 0
}

func (c *CustOrdHandlInstEnum) AlgoEngineInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.AlgoEngineSinceVersion()
}

func (*CustOrdHandlInstEnum) AlgoEngineDeprecated() uint16 {
	return 0
}

func (*CustOrdHandlInstEnum) DeskElectronicSinceVersion() uint16 {
	return 0
}

func (c *CustOrdHandlInstEnum) DeskElectronicInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.DeskElectronicSinceVersion()
}

func (*CustOrdHandlInstEnum) DeskElectronicDeprecated() uint16 {
	return 0
}

func (*CustOrdHandlInstEnum) ClientElectronicSinceVersion() uint16 {
	return 0
}

func (c *CustOrdHandlInstEnum) ClientElectronicInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.ClientElectronicSinceVersion()
}

func (*CustOrdHandlInstEnum) ClientElectronicDeprecated() uint16 {
	return 0
}
