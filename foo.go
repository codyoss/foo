package foo

import (
	"fmt"
	"strings"

	"github.com/codyoss/bar"
)

const (
	// Something is a random float.
	Something = 1e7
)

// Combine makes a string from ints.
func Combine(is ...int) string {
	var b strings.Builder
	for _, v := range is {
		fmt.Fprintf(&b, "%s ", bar.String(v))
	}
	return b.String()
}
