package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateMsg(t *testing.T) {
	valid := []int {
		500, 501, 502, 503, 504, 505, 506, 507, 508,
		509, 510, 513, 514, 515, 516, 517, 518, 519, 521,
		522, 523, 524, 525, 526, 527, 528, 529, 530, 531,
		532, 533, 534, 535, 536, 537, 538, 539, 543, 544,
		545, 546, 548, 549, 550, 560, 561, 562, 563,
	}
	assert := assert.New(t)
	for _, v := range valid {
		r, err := createMsg(v)
		assert.NotNil(r)
		assert.NoError(err)
	}
	r, err := createMsg(199)
	assert.Nil(r)
	assert.Error(err)
}