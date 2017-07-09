package streamlimit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStreamLimit(t *testing.T) {
	limiter := New(8, 4)

	limiter.Write([]byte("TESTTEST"))
	limiter.Start()

	var count = 0

	firstRead := make([]byte, 8)
	count, _ = limiter.Read(firstRead)
	assert.Equal(t, 2, count)
	assert.Equal(t, []byte("TE"), firstRead[:count])

	secondRead := make([]byte, 8)
	count, _ = limiter.Read(secondRead)
	assert.Equal(t, 2, count)
	assert.Equal(t, []byte("ST"), secondRead[:count])
}

func TestStreamLimitWithHalfData(t *testing.T) {
	limiter := New(16, 1)

	limiter.Write([]byte("TESTTEST"))
	limiter.Start()

	var count = 0

	firstRead := make([]byte, 8)
	count, _ = limiter.Read(firstRead)
	assert.Equal(t, 8, count)
	assert.Equal(t, []byte("TESTTEST"), firstRead[:count])
}

func TestTimeRemaining(t *testing.T) {
	limiter := New(1, 1)

	limiter.Write([]byte("TESTTEST"))

	assert.Equal(t, 8, limiter.RemainingTime())
}
