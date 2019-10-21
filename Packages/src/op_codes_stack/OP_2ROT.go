// Move the fifth and sixth items back, to the top of the stack.

package op_codes_stack

import "stack"

func OP_2ROT(mainStack * stack.Stack) bool {
	dataItem1, ok := mainStack.RemoveNthToTop(6)
	if !ok {
		return false
	}
	dataItem2 := mainStack.RemoveNthToTop(5)
	mainStack.Push(dataItem1)
	mainStack.Push(dataItem2)
	return true
}