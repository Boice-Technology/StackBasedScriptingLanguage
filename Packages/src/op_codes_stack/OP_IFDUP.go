// Duplicate top element, if top element is not 0.

package op_codes_stack

import "stack"

func OP_IFDUP(mainStack *stack.Stack) bool {
	var verdict bool
	var dataItem stack.Item
	dataItem, verdict = mainStack.Peek()
	if verdict == true && dataItem != 0 {
		mainStack.Push(dataItem)
	} else {
		verdict = false
	}
	return false
}