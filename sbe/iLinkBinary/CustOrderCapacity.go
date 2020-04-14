// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"reflect"
)

type CustOrderCapacityEnum uint8
type CustOrderCapacityValues struct {
	Membertradingfortheirownaccount             CustOrderCapacityEnum
	Clearingfirmtradingforitsproprietaryaccount CustOrderCapacityEnum
	Membertradingforanothermember               CustOrderCapacityEnum
	Allother                                    CustOrderCapacityEnum
	NullValue                                   CustOrderCapacityEnum
}

var CustOrderCapacity = CustOrderCapacityValues{1, 2, 3, 4, 255}

func (c CustOrderCapacityEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(c)); err != nil {
		return err
	}
	return nil
}

func (c *CustOrderCapacityEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(c)); err != nil {
		return err
	}
	return nil
}

func (c CustOrderCapacityEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(CustOrderCapacity)
	for idx := 0; idx < value.NumField(); idx++ {
		if c == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on CustOrderCapacity, unknown enumeration value %d", c)
}

func (*CustOrderCapacityEnum) EncodedLength() int64 {
	return 1
}

func (*CustOrderCapacityEnum) MembertradingfortheirownaccountSinceVersion() uint16 {
	return 0
}

func (c *CustOrderCapacityEnum) MembertradingfortheirownaccountInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.MembertradingfortheirownaccountSinceVersion()
}

func (*CustOrderCapacityEnum) MembertradingfortheirownaccountDeprecated() uint16 {
	return 0
}

func (*CustOrderCapacityEnum) ClearingfirmtradingforitsproprietaryaccountSinceVersion() uint16 {
	return 0
}

func (c *CustOrderCapacityEnum) ClearingfirmtradingforitsproprietaryaccountInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.ClearingfirmtradingforitsproprietaryaccountSinceVersion()
}

func (*CustOrderCapacityEnum) ClearingfirmtradingforitsproprietaryaccountDeprecated() uint16 {
	return 0
}

func (*CustOrderCapacityEnum) MembertradingforanothermemberSinceVersion() uint16 {
	return 0
}

func (c *CustOrderCapacityEnum) MembertradingforanothermemberInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.MembertradingforanothermemberSinceVersion()
}

func (*CustOrderCapacityEnum) MembertradingforanothermemberDeprecated() uint16 {
	return 0
}

func (*CustOrderCapacityEnum) AllotherSinceVersion() uint16 {
	return 0
}

func (c *CustOrderCapacityEnum) AllotherInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.AllotherSinceVersion()
}

func (*CustOrderCapacityEnum) AllotherDeprecated() uint16 {
	return 0
}
