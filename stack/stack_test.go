package stack_test

import (
	"learning_golang/stack"
	"testing"
)

func TestStack(t *testing.T)  {
	count := 1
	var aStack stack.Stack
	assertTrue(t,aStack.Len()==0,"expected empty Stack",count)
}

// assertTrue() calls testing.T.Error() with the given message if the
// condition is false.
func assertTrue(t *testing.T, condition bool, message string, id int) {
	if !condition {
		t.Errorf("#%d: %s", id, message)
	}
}