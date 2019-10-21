// Remove the top two stack items.

package op_codes_stack

import "stack"

func OP_2DROP(mainStack *stack.Stack) bool {
	d, ok := mainStack.Pop()
	if !ok {
		return false
	}
	_, ok := mainStack.Pop()
	if !ok {
		mainStack.Push(d)
		return false
	}
	return true
}