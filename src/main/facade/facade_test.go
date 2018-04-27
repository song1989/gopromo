package facade

import (
	"testing"
)

func TestFacade(t *testing.T) {
	f := NewFacade("3.4", 1, 2.3)
	f.OutOne()
	f.OutTwo()
}
