package handler

import (
	"fmt"
	"seed/errors"
	"testing"
)

func TestNewReply(t *testing.T) {
	r := NewReply(errors.New("ni"))
	fmt.Println(r)
}
