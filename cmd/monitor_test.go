package cmd

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestValidateStatus(t *testing.T) {
    res := map[string]int{}
    res["200"] = 5
    assert.True(t, ValidateStatus(res, 5, "200:100"), "5x 200, should get us a 100% 200")
    assert.False(t, ValidateStatus(res, 10, "200:100"), "10x 200, should get us only 50% 200")
    res["500"] = 5
    assert.True(t, ValidateStatus(res, 10, "200:50,500:50"), "10x 200, should get us only 50% 200")
}
