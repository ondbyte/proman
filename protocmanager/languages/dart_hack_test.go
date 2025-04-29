package languages_test

import (
	"testing"

	"github.com/ondbyte/proman/protocmanager/languages"
	"github.com/stretchr/testify/assert"
)

func TestFindProtocGenDart(t *testing.T) {
	p, err := languages.FindProtocGenDart()
	if !assert.NoError(t, err) {
		return
	}
	assert.NotEmpty(t, p)
}
