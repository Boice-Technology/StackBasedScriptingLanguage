// Rotate the top three items on the stack to the left.

package op_codes_stack

import "stack"

func OP_ROT(mainStack * stack.Stack) bool {
	dataItem, ok := mainStack.RemoveNthToTop(3)
	if !ok {
		return false
	}
	mainStack.Push(dataItem)
	return true
}