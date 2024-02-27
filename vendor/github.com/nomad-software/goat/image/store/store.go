package store

import (
	"embed"
	"encoding/base64"
	"path/filepath"

	"github.com/nomad-software/goat/image"
	"github.com/nomad-software/goat/internal/log"
)

// Store is an image store of embedded images.
type Store struct {
	fs embed.FS
}

// New creates a new image store.
func New(fs embed.FS) *Store {
	st := &Store{
		fs: fs,
	}

	return st
}

// GetImage gets an image from the store.
// The passed name must be a valid path in the embedded store.
// Only Png and Gif formats are supported.
func (s *Store) GetImage(path string) *image.Image {
	b, err := s.fs.ReadFile(path)
	if err != nil {
		log.Panic(err, "cannot read file")
	}

	data := base64.StdEncoding.EncodeToString(b)
	ext := filepath.Ext(path)

	switch ext {
	case ".gif":
		return image.New(data, "gif")
	case ".png":
		return image.New(data, "png")
	default:
		panic("image extension not recognised")
	}
}

// GetImages gets multiple images at once.
func (s *Store) GetImages(paths ...string) []*image.Image {
	imgs := make([]*image.Image, 0)

	for _, p := range paths {
		img := s.GetImage(p)
		imgs = append(imgs, img)
	}

	return imgs
}
