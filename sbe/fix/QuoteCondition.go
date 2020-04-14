// Generated SBE (Simple Binary Encoding) message codec

package fix

import (
	"io"
)

type QuoteCondition [8]bool
type QuoteConditionChoiceValue uint8
type QuoteConditionChoiceValues struct {
	Implied      QuoteConditionChoiceValue
	ExchangeBest QuoteConditionChoiceValue
}

var QuoteConditionChoice = QuoteConditionChoiceValues{0, 1}

func (q *QuoteCondition) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	var wireval uint8 = 0
	for k, v := range q {
		if v {
			wireval |= (1 << uint(k))
		}
	}
	return _m.WriteUint8(_w, wireval)
}

func (q *QuoteCondition) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	var wireval uint8

	if err := _m.ReadUint8(_r, &wireval); err != nil {
		return err
	}

	var idx uint
	for idx = 0; idx < 8; idx++ {
		q[idx] = (wireval & (1 << idx)) > 0
	}
	return nil
}

func (QuoteCondition) EncodedLength() int64 {
	return 1
}

func (*QuoteCondition) ImpliedSinceVersion() uint16 {
	return 0
}

func (q *QuoteCondition) ImpliedInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.ImpliedSinceVersion()
}

func (*QuoteCondition) ImpliedDeprecated() uint16 {
	return 0
}

func (*QuoteCondition) ExchangeBestSinceVersion() uint16 {
	return 0
}

func (q *QuoteCondition) ExchangeBestInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.ExchangeBestSinceVersion()
}

func (*QuoteCondition) ExchangeBestDeprecated() uint16 {
	return 0
}
