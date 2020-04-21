package handler

import (
	"context"
	mockentity "sbe/mock/entity"
	fix "sbe/sbe/iLinkBinary"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	r := New()
	assert.NotNil(t, r)
}

func TestOnOrderNew(t *testing.T) {
	tests := []struct {
		retErrs []error
		ret     error
	}{
		{
			retErrs: []error{assert.AnError, nil},
			ret:     assert.AnError,
		},
		{
			retErrs: []error{nil, assert.AnError},
			ret:     assert.AnError,
		},
		{
			retErrs: []error{nil, nil},
			ret:     nil,
		},
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	for _, tt := range tests {
		mockSession := mockentity.NewMockSBESession(ctrl)
		gomock.InOrder(
			mockSession.EXPECT().Send(gomock.Any()).Return(tt.retErrs[0]).MinTimes(1).MaxTimes(1),
			mockSession.EXPECT().Send(gomock.Any()).Return(tt.retErrs[1]).MinTimes(0).MaxTimes(1),
		)
		mockSession.EXPECT().GetUUID().Return(uint64(1)).AnyTimes()
		mockSession.EXPECT().GetSeqNo().Return(uint32(1)).AnyTimes()
		mockSession.EXPECT().SetSeqNo(gomock.Any()).AnyTimes()
		h := New()
		order := fix.NewOrderSingle514{}
		ret := h.OnOrderNew(context.Background(), &order, mockSession)
		assert := assert.New(t)
		assert.Equal(tt.ret, ret)
	}
}
