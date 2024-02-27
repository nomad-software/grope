package image

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	// 1x1 black png image data.
	data = "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mP8z8BQDwAEhQGAhKmMIQAAAABJRU5ErkJggg=="
	typ  = "png"
)

func TestImage(t *testing.T) {
	img := New(data, typ)

	assert.Regexp(t, `^\.png-[A-Z0-9]{1,8}$`, img.GetID())
}

func TestImageGamma(t *testing.T) {
	img := New(data, typ)

	img.SetGamma(0.5)
	assert.Equal(t, 0.5, img.GetGamma())

	img.Destroy()
}

func TestImageDimensions(t *testing.T) {
	img := New(data, typ)

	img.SetWidth(1)
	assert.Equal(t, 1, img.GetWidth())

	img.SetHeight(1)
	assert.Equal(t, 1, img.GetHeight())

	img.Destroy()
}
