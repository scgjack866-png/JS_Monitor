package utils

import (
	"fmt"
	"testing"
)

func TestIsNull(t *testing.T) {
	var s uint64
	fmt.Print(IsNull(s))
}
