// Remove Nth item to top of the stack and push to top of the stack.

package op_codes_stack

import "stack"

func OP_ROLL(mainStack *stack.Stack, N int) bool {
	var verdict bool
	var dataItem stack.Item
	dataItem, verdict = mainStack.RemoveNthToTop(N)
	if verdict == true {
		mainStack.Push(dataItem)
	}
	return verdict
}