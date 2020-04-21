package configure

import (
	assert "github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	got := New()
	assert := assert.New(t)
	assert.NotNil(got)
	assert.Equal(got.APPName, "sbesim")
}
