package languages_test

import (
	"testing"

	"github.com/ondbyte/proman/protocmanager/languages"
	"github.com/stretchr/testify/assert"
)

func TestFindProtocGenGo(t *testing.T) {
	p, err := languages.FindProtocGenGo()
	if !assert.NoError(t, err) {
		return
	}
	assert.NotEmpty(t, p)
}
