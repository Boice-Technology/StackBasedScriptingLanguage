// Swap the top two pairs of items.

package op_codes_stack

import "stack"

func OP_2SWAP(mainStack * stack.Stack) bool {
	dataItem1, ok := mainStack.RemoveNthToTop(4)
	if !ok {
		return false
	}
	dataItem2 := mainStack.RemoveNthToTop(3)
	mainStack.Push(dataItem1)
	mainStack.Push(dataItem2)
	return true
}