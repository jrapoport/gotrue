package account

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinked_Valid(t *testing.T) {
	l := &Linked{}
	err := l.Valid()
	assert.Error(t, err)
	l.Type = -1
	err = l.Valid()
	assert.Error(t, err)
	l.Type = Auth
	err = l.Valid()
	assert.Error(t, err)
	l.Provider = "test provider"
	err = l.Valid()
	assert.Error(t, err)
	const accountID = "test-id"
	l.AccountID = accountID
	err = l.Valid()
	assert.NoError(t, err)
}