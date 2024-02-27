package hash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashRandomHash(t *testing.T) {
	hash := Generate()

	assert.Regexp(t, `^[A-Z0-9]{1,8}$`, hash)
}

func TestHashDeterministicHash(t *testing.T) {
	hash := Generate("foo", "bar", "baz")

	assert.Equal(t, "606D6255", hash)
}
