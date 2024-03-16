package store

import (
	"testing"

	"github.com/nomad-software/goat/example/image"
	"github.com/stretchr/testify/assert"
)

func TestStoreGetImage(t *testing.T) {
	store := New(image.FS)

	img := store.GetImage("png/thumbnail.png")
	assert.Regexp(t, `^\.png-[A-Z0-9]{1,8}$`, img.GetID())
	assert.Equal(t, "png", img.GetType())

	img = store.GetImage("gif/thumbnail.gif")
	assert.Regexp(t, `^\.gif-[A-Z0-9]{1,8}$`, img.GetID())
	assert.Equal(t, "gif", img.GetType())
}

func TestStoreGetImages(t *testing.T) {
	store := New(image.FS)
	imgs := store.GetImages("png/thumbnail.png", "gif/thumbnail.gif")

	for _, img := range imgs {
		assert.Regexp(t, `^\.(gif|png)-[A-Z0-9]{1,8}$`, img.GetID())
		assert.Contains(t, []string{"gif", "png"}, img.GetType())
	}
}
