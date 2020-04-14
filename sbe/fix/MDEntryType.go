// Generated SBE (Simple Binary Encoding) message codec

package fix

import (
	"fmt"
	"io"
	"reflect"
)

type MDEntryTypeEnum byte
type MDEntryTypeValues struct {
	BID                        MDEntryTypeEnum
	OFFER                      MDEntryTypeEnum
	TRADE                      MDEntryTypeEnum
	OPENING_PRICE              MDEntryTypeEnum
	SETTLEMENT_PRICE           MDEntryTypeEnum
	TRADING_SESSION_HIGH_PRICE MDEntryTypeEnum
	TRADING_SESSION_LOW_PRICE  MDEntryTypeEnum
	TRADE_VOLUME               MDEntryTypeEnum
	OPEN_INTEREST              MDEntryTypeEnum
	SIMULATED_SELL             MDEntryTypeEnum
	SIMULATED_BUY              MDEntryTypeEnum
	EMPTY_THE_BOOK             MDEntryTypeEnum
	SESSION_HIGH_BID           MDEntryTypeEnum
	SESSION_LOW_OFFER          MDEntryTypeEnum
	FIXING_PRICE               MDEntryTypeEnum
	CASH_NOTE                  MDEntryTypeEnum
	NullValue                  MDEntryTypeEnum
}

var MDEntryType = MDEntryTypeValues{48, 49, 50, 52, 54, 55, 56, 66, 67, 69, 70, 74, 78, 79, 87, 88, 0}

func (m MDEntryTypeEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, byte(m)); err != nil {
		return err
	}
	return nil
}

func (m *MDEntryTypeEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*byte)(m)); err != nil {
		return err
	}
	return nil
}

func (m MDEntryTypeEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(MDEntryType)
	for idx := 0; idx < value.NumField(); idx++ {
		if m == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on MDEntryType, unknown enumeration value %d", m)
}

func (*MDEntryTypeEnum) EncodedLength() int64 {
	return 1
}

func (*MDEntryTypeEnum) BIDSinceVersion() uint16 {
	return 0
}

func (m *MDEntryTypeEnum) BIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.BIDSinceVersion()
}

func (*MDEntryTypeEnum) BIDDeprecated() uint16 {
	return 0
}

func (*MDEntryTypeEnum) OFFERSinceVersion() uint16 {
	return 0
}

func (m *MDEntryTypeEnum) OFFERInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.OFFERSinceVersion()
}

func (*MDEntryTypeEnum) OFFERDeprecated() uint16 {
	return 0
}

func (*MDEntryTypeEnum) TRADESinceVersion() uint16 {
	return 0
}

func (m *MDEntryTypeEnum) TRADEInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.TRADESinceVersion()
}

func (*MDEntryTypeEnum) TRADEDeprecated() uint16 {
	return 0
}

func (*MDEntryTypeEnum) OPENING_PRICESinceVersion() uint16 {
	return 0
}

func (m *MDEntryTypeEnum) OPENING_PRICEInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.OPENING_PRICESinceVersion()
}

func (*MDEntryTypeEnum) OPENING_PRICEDeprecated() uint16 {
	return 0
}

func (*MDEntryTypeEnum) SETTLEMENT_PRICESinceVersion() uint16 {
	return 0
}

func (m *MDEntryTypeEnum) SETTLEMENT_PRICEInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SETTLEMENT_PRICESinceVersion()
}

func (*MDEntryTypeEnum) SETTLEMENT_PRICEDeprecated() uint16 {
	return 0
}

func (*MDEntryTypeEnum) TRADING_SESSION_HIGH_PRICESinceVersion() uint16 {
	return 0
}

func (m *MDEntryTypeEnum) TRADING_SESSION_HIGH_PRICEInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.TRADING_SESSION_HIGH_PRICESinceVersion()
}

func (*MDEntryTypeEnum) TRADING_SESSION_HIGH_PRICEDeprecated() uint16 {
	return 0
}

func (*MDEntryTypeEnum) TRADING_SESSION_LOW_PRICESinceVersion() uint16 {
	return 0
}

func (m *MDEntryTypeEnum) TRADING_SESSION_LOW_PRICEInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.TRADING_SESSION_LOW_PRICESinceVersion()
}

func (*MDEntryTypeEnum) TRADING_SESSION_LOW_PRICEDeprecated() uint16 {
	return 0
}

func (*MDEntryTypeEnum) TRADE_VOLUMESinceVersion() uint16 {
	return 0
}

func (m *MDEntryTypeEnum) TRADE_VOLUMEInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.TRADE_VOLUMESinceVersion()
}

func (*MDEntryTypeEnum) TRADE_VOLUMEDeprecated() uint16 {
	return 0
}

func (*MDEntryTypeEnum) OPEN_INTERESTSinceVersion() uint16 {
	return 0
}

func (m *MDEntryTypeEnum) OPEN_INTERESTInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.OPEN_INTERESTSinceVersion()
}

func (*MDEntryTypeEnum) OPEN_INTERESTDeprecated() uint16 {
	return 0
}

func (*MDEntryTypeEnum) SIMULATED_SELLSinceVersion() uint16 {
	return 0
}

func (m *MDEntryTypeEnum) SIMULATED_SELLInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SIMULATED_SELLSinceVersion()
}

func (*MDEntryTypeEnum) SIMULATED_SELLDeprecated() uint16 {
	return 0
}

func (*MDEntryTypeEnum) SIMULATED_BUYSinceVersion() uint16 {
	return 0
}

func (m *MDEntryTypeEnum) SIMULATED_BUYInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SIMULATED_BUYSinceVersion()
}

func (*MDEntryTypeEnum) SIMULATED_BUYDeprecated() uint16 {
	return 0
}

func (*MDEntryTypeEnum) EMPTY_THE_BOOKSinceVersion() uint16 {
	return 0
}

func (m *MDEntryTypeEnum) EMPTY_THE_BOOKInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.EMPTY_THE_BOOKSinceVersion()
}

func (*MDEntryTypeEnum) EMPTY_THE_BOOKDeprecated() uint16 {
	return 0
}

func (*MDEntryTypeEnum) SESSION_HIGH_BIDSinceVersion() uint16 {
	return 0
}

func (m *MDEntryTypeEnum) SESSION_HIGH_BIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SESSION_HIGH_BIDSinceVersion()
}

func (*MDEntryTypeEnum) SESSION_HIGH_BIDDeprecated() uint16 {
	return 0
}

func (*MDEntryTypeEnum) SESSION_LOW_OFFERSinceVersion() uint16 {
	return 0
}

func (m *MDEntryTypeEnum) SESSION_LOW_OFFERInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SESSION_LOW_OFFERSinceVersion()
}

func (*MDEntryTypeEnum) SESSION_LOW_OFFERDeprecated() uint16 {
	return 0
}

func (*MDEntryTypeEnum) FIXING_PRICESinceVersion() uint16 {
	return 0
}

func (m *MDEntryTypeEnum) FIXING_PRICEInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.FIXING_PRICESinceVersion()
}

func (*MDEntryTypeEnum) FIXING_PRICEDeprecated() uint16 {
	return 0
}

func (*MDEntryTypeEnum) CASH_NOTESinceVersion() uint16 {
	return 0
}

func (m *MDEntryTypeEnum) CASH_NOTEInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.CASH_NOTESinceVersion()
}

func (*MDEntryTypeEnum) CASH_NOTEDeprecated() uint16 {
	return 0
}
