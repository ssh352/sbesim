// Generated SBE (Simple Binary Encoding) message codec

package fix

import (
	"fmt"
	"io"
	"reflect"
)

type CustOrderHandlingInstEnum byte
type CustOrderHandlingInstValues struct {
	PHONE_SIMPLE                                CustOrderHandlingInstEnum
	PHONE_COMPLEX                               CustOrderHandlingInstEnum
	FCM_PROVIDED_SCREEN                         CustOrderHandlingInstEnum
	OTHER_PROVIDED_SCREEN                       CustOrderHandlingInstEnum
	CLIENT_PROVIDED_PLATFORM_CONTROLLED_BY_FCM  CustOrderHandlingInstEnum
	CLIENT_PROVIDED_PLATFORM_DIRECT_TO_EXCHANGE CustOrderHandlingInstEnum
	FCM_API_OR_FIX                              CustOrderHandlingInstEnum
	ALGO_ENGINE                                 CustOrderHandlingInstEnum
	PRICE_AT_EXECUTION                          CustOrderHandlingInstEnum
	DESK_ELECTRONIC                             CustOrderHandlingInstEnum
	DESK_PIT                                    CustOrderHandlingInstEnum
	CLIENT_ELECTRONIC                           CustOrderHandlingInstEnum
	CLIENT_PIT                                  CustOrderHandlingInstEnum
	NullValue                                   CustOrderHandlingInstEnum
}

var CustOrderHandlingInst = CustOrderHandlingInstValues{65, 66, 67, 68, 69, 70, 71, 72, 74, 87, 88, 89, 90, 0}

func (c CustOrderHandlingInstEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, byte(c)); err != nil {
		return err
	}
	return nil
}

func (c *CustOrderHandlingInstEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*byte)(c)); err != nil {
		return err
	}
	return nil
}

func (c CustOrderHandlingInstEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(CustOrderHandlingInst)
	for idx := 0; idx < value.NumField(); idx++ {
		if c == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on CustOrderHandlingInst, unknown enumeration value %d", c)
}

func (*CustOrderHandlingInstEnum) EncodedLength() int64 {
	return 1
}

func (*CustOrderHandlingInstEnum) PHONE_SIMPLESinceVersion() uint16 {
	return 0
}

func (c *CustOrderHandlingInstEnum) PHONE_SIMPLEInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.PHONE_SIMPLESinceVersion()
}

func (*CustOrderHandlingInstEnum) PHONE_SIMPLEDeprecated() uint16 {
	return 0
}

func (*CustOrderHandlingInstEnum) PHONE_COMPLEXSinceVersion() uint16 {
	return 0
}

func (c *CustOrderHandlingInstEnum) PHONE_COMPLEXInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.PHONE_COMPLEXSinceVersion()
}

func (*CustOrderHandlingInstEnum) PHONE_COMPLEXDeprecated() uint16 {
	return 0
}

func (*CustOrderHandlingInstEnum) FCM_PROVIDED_SCREENSinceVersion() uint16 {
	return 0
}

func (c *CustOrderHandlingInstEnum) FCM_PROVIDED_SCREENInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.FCM_PROVIDED_SCREENSinceVersion()
}

func (*CustOrderHandlingInstEnum) FCM_PROVIDED_SCREENDeprecated() uint16 {
	return 0
}

func (*CustOrderHandlingInstEnum) OTHER_PROVIDED_SCREENSinceVersion() uint16 {
	return 0
}

func (c *CustOrderHandlingInstEnum) OTHER_PROVIDED_SCREENInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.OTHER_PROVIDED_SCREENSinceVersion()
}

func (*CustOrderHandlingInstEnum) OTHER_PROVIDED_SCREENDeprecated() uint16 {
	return 0
}

func (*CustOrderHandlingInstEnum) CLIENT_PROVIDED_PLATFORM_CONTROLLED_BY_FCMSinceVersion() uint16 {
	return 0
}

func (c *CustOrderHandlingInstEnum) CLIENT_PROVIDED_PLATFORM_CONTROLLED_BY_FCMInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.CLIENT_PROVIDED_PLATFORM_CONTROLLED_BY_FCMSinceVersion()
}

func (*CustOrderHandlingInstEnum) CLIENT_PROVIDED_PLATFORM_CONTROLLED_BY_FCMDeprecated() uint16 {
	return 0
}

func (*CustOrderHandlingInstEnum) CLIENT_PROVIDED_PLATFORM_DIRECT_TO_EXCHANGESinceVersion() uint16 {
	return 0
}

func (c *CustOrderHandlingInstEnum) CLIENT_PROVIDED_PLATFORM_DIRECT_TO_EXCHANGEInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.CLIENT_PROVIDED_PLATFORM_DIRECT_TO_EXCHANGESinceVersion()
}

func (*CustOrderHandlingInstEnum) CLIENT_PROVIDED_PLATFORM_DIRECT_TO_EXCHANGEDeprecated() uint16 {
	return 0
}

func (*CustOrderHandlingInstEnum) FCM_API_OR_FIXSinceVersion() uint16 {
	return 0
}

func (c *CustOrderHandlingInstEnum) FCM_API_OR_FIXInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.FCM_API_OR_FIXSinceVersion()
}

func (*CustOrderHandlingInstEnum) FCM_API_OR_FIXDeprecated() uint16 {
	return 0
}

func (*CustOrderHandlingInstEnum) ALGO_ENGINESinceVersion() uint16 {
	return 0
}

func (c *CustOrderHandlingInstEnum) ALGO_ENGINEInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.ALGO_ENGINESinceVersion()
}

func (*CustOrderHandlingInstEnum) ALGO_ENGINEDeprecated() uint16 {
	return 0
}

func (*CustOrderHandlingInstEnum) PRICE_AT_EXECUTIONSinceVersion() uint16 {
	return 0
}

func (c *CustOrderHandlingInstEnum) PRICE_AT_EXECUTIONInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.PRICE_AT_EXECUTIONSinceVersion()
}

func (*CustOrderHandlingInstEnum) PRICE_AT_EXECUTIONDeprecated() uint16 {
	return 0
}

func (*CustOrderHandlingInstEnum) DESK_ELECTRONICSinceVersion() uint16 {
	return 0
}

func (c *CustOrderHandlingInstEnum) DESK_ELECTRONICInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.DESK_ELECTRONICSinceVersion()
}

func (*CustOrderHandlingInstEnum) DESK_ELECTRONICDeprecated() uint16 {
	return 0
}

func (*CustOrderHandlingInstEnum) DESK_PITSinceVersion() uint16 {
	return 0
}

func (c *CustOrderHandlingInstEnum) DESK_PITInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.DESK_PITSinceVersion()
}

func (*CustOrderHandlingInstEnum) DESK_PITDeprecated() uint16 {
	return 0
}

func (*CustOrderHandlingInstEnum) CLIENT_ELECTRONICSinceVersion() uint16 {
	return 0
}

func (c *CustOrderHandlingInstEnum) CLIENT_ELECTRONICInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.CLIENT_ELECTRONICSinceVersion()
}

func (*CustOrderHandlingInstEnum) CLIENT_ELECTRONICDeprecated() uint16 {
	return 0
}

func (*CustOrderHandlingInstEnum) CLIENT_PITSinceVersion() uint16 {
	return 0
}

func (c *CustOrderHandlingInstEnum) CLIENT_PITInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.CLIENT_PITSinceVersion()
}

func (*CustOrderHandlingInstEnum) CLIENT_PITDeprecated() uint16 {
	return 0
}
