package stack

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestStack(t *testing.T) {
	stk := NewStack[int]()
	stk.Push(1)
	stk.Print(os.Stdout)
	stk.Push(2)
	stk.Print(os.Stdout)
	stk.Push(3)
	stk.Print(os.Stdout)
	stk.Push(4)
	stk.Print(os.Stdout)

	assert.True(t, stk.Size() == 4)
	item, _ := stk.Pop()
	fmt.Println(item)
}
