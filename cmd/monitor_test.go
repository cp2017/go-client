package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cp2017/go-client/util"
)

func TestValidateStatus(t *testing.T) {
	res := map[string]int{}
	res["200"] = 5
	assert.True(t, util.ValidateStatus(res, 5, "200:100"), "5x 200, should get us a 100% 200")
	assert.False(t, util.ValidateStatus(res, 10, "200:100"), "10x 200, should get us only 50% 200")
	res["500"] = 5
	assert.True(t, util.ValidateStatus(res, 10, "200:50,500:50"), "10x 200, should get us only 50% 200")
}
