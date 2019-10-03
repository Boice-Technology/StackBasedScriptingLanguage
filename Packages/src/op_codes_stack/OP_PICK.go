// Copy Nth item to the top of the stack, into the stack.

package op_codes_stack

import "stack"

func OP_PICK(mainStack *stack.Stack, N int) bool {
	var verdict bool
	var dataItem stack.Item
	dataItem, verdict = mainStack.AtNthToTop(N)
	if verdict==true {
		mainStack.Push(dataItem)
	}
	return verdict
}