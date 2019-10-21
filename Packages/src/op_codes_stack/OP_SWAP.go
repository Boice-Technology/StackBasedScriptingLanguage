// Swap the top two items on the stack.

package op_codes_stack

import "stack"

func OP_SWAP(mainStack * stack.Stack) bool {
	dataItem, ok := mainStack.RemoveNthToTop(2)
	if !ok {
		return false
	}
	mainStack.Push(dataItem)
	return true
}