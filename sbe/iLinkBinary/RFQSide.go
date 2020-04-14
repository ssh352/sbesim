// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"reflect"
)

type RFQSideEnum uint8
type RFQSideValues struct {
	Buy       RFQSideEnum
	Sell      RFQSideEnum
	Cross     RFQSideEnum
	NullValue RFQSideEnum
}

var RFQSide = RFQSideValues{1, 2, 8, 255}

func (r RFQSideEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(r)); err != nil {
		return err
	}
	return nil
}

func (r *RFQSideEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(r)); err != nil {
		return err
	}
	return nil
}

func (r RFQSideEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(RFQSide)
	for idx := 0; idx < value.NumField(); idx++ {
		if r == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on RFQSide, unknown enumeration value %d", r)
}

func (*RFQSideEnum) EncodedLength() int64 {
	return 1
}

func (*RFQSideEnum) BuySinceVersion() uint16 {
	return 0
}

func (r *RFQSideEnum) BuyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.BuySinceVersion()
}

func (*RFQSideEnum) BuyDeprecated() uint16 {
	return 0
}

func (*RFQSideEnum) SellSinceVersion() uint16 {
	return 0
}

func (r *RFQSideEnum) SellInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.SellSinceVersion()
}

func (*RFQSideEnum) SellDeprecated() uint16 {
	return 0
}

func (*RFQSideEnum) CrossSinceVersion() uint16 {
	return 0
}

func (r *RFQSideEnum) CrossInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.CrossSinceVersion()
}

func (*RFQSideEnum) CrossDeprecated() uint16 {
	return 0
}
