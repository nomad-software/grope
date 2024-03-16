package variable

import (
	"fmt"

	"github.com/nomad-software/goat/internal/widget/ui/element/hash"
)

// GenerateName generates a custom variable name.
func GenerateName(args ...string) string {
	args = append(args, "variable")
	hash := hash.Generate(args...)

	return fmt.Sprintf("var-%s", hash)
}
