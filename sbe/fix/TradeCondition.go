// Generated SBE (Simple Binary Encoding) message codec

package fix

import (
	"io"
)

type TradeCondition [8]bool
type TradeConditionChoiceValue uint8
type TradeConditionChoiceValues struct {
	OpeningTrade   TradeConditionChoiceValue
	CmeGlobexPrice TradeConditionChoiceValue
}

var TradeConditionChoice = TradeConditionChoiceValues{0, 1}

func (t *TradeCondition) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	var wireval uint8 = 0
	for k, v := range t {
		if v {
			wireval |= (1 << uint(k))
		}
	}
	return _m.WriteUint8(_w, wireval)
}

func (t *TradeCondition) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	var wireval uint8

	if err := _m.ReadUint8(_r, &wireval); err != nil {
		return err
	}

	var idx uint
	for idx = 0; idx < 8; idx++ {
		t[idx] = (wireval & (1 << idx)) > 0
	}
	return nil
}

func (TradeCondition) EncodedLength() int64 {
	return 1
}

func (*TradeCondition) OpeningTradeSinceVersion() uint16 {
	return 0
}

func (t *TradeCondition) OpeningTradeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.OpeningTradeSinceVersion()
}

func (*TradeCondition) OpeningTradeDeprecated() uint16 {
	return 0
}

func (*TradeCondition) CmeGlobexPriceSinceVersion() uint16 {
	return 0
}

func (t *TradeCondition) CmeGlobexPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.CmeGlobexPriceSinceVersion()
}

func (*TradeCondition) CmeGlobexPriceDeprecated() uint16 {
	return 0
}
