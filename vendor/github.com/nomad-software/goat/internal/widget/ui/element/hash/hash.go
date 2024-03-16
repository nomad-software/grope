package hash

import (
	"fmt"
	"hash/fnv"
	"math/rand"
	"strings"
)

// Generate generates an element hash.
// If no arguments are given a random hash is returned. Otherwise the arguments
// are hashed.
func Generate(args ...string) string {
	var text string

	if len(args) > 0 {
		text = strings.Join(args, "")
	} else {
		text = fmt.Sprint(rand.Int63())
	}

	hash := fnv.New32a()
	hash.Write([]byte(text))
	return fmt.Sprintf("%X", hash.Sum32())
}
