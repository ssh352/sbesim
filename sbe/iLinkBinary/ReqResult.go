// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"reflect"
)

type ReqResultEnum uint8
type ReqResultValues struct {
	ValidRequest                            ReqResultEnum
	NoDataFoundThatMatchedSelectionCriteria ReqResultEnum
	NotAuthorizedtoRetrieveData             ReqResultEnum
	DataTemporarilyUnavailable              ReqResultEnum
	NullValue                               ReqResultEnum
}

var ReqResult = ReqResultValues{0, 2, 3, 4, 255}

func (r ReqResultEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(r)); err != nil {
		return err
	}
	return nil
}

func (r *ReqResultEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(r)); err != nil {
		return err
	}
	return nil
}

func (r ReqResultEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(ReqResult)
	for idx := 0; idx < value.NumField(); idx++ {
		if r == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on ReqResult, unknown enumeration value %d", r)
}

func (*ReqResultEnum) EncodedLength() int64 {
	return 1
}

func (*ReqResultEnum) ValidRequestSinceVersion() uint16 {
	return 0
}

func (r *ReqResultEnum) ValidRequestInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.ValidRequestSinceVersion()
}

func (*ReqResultEnum) ValidRequestDeprecated() uint16 {
	return 0
}

func (*ReqResultEnum) NoDataFoundThatMatchedSelectionCriteriaSinceVersion() uint16 {
	return 0
}

func (r *ReqResultEnum) NoDataFoundThatMatchedSelectionCriteriaInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.NoDataFoundThatMatchedSelectionCriteriaSinceVersion()
}

func (*ReqResultEnum) NoDataFoundThatMatchedSelectionCriteriaDeprecated() uint16 {
	return 0
}

func (*ReqResultEnum) NotAuthorizedtoRetrieveDataSinceVersion() uint16 {
	return 0
}

func (r *ReqResultEnum) NotAuthorizedtoRetrieveDataInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.NotAuthorizedtoRetrieveDataSinceVersion()
}

func (*ReqResultEnum) NotAuthorizedtoRetrieveDataDeprecated() uint16 {
	return 0
}

func (*ReqResultEnum) DataTemporarilyUnavailableSinceVersion() uint16 {
	return 0
}

func (r *ReqResultEnum) DataTemporarilyUnavailableInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.DataTemporarilyUnavailableSinceVersion()
}

func (*ReqResultEnum) DataTemporarilyUnavailableDeprecated() uint16 {
	return 0
}
