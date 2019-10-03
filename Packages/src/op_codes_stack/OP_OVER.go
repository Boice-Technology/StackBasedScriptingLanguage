// Copy second item to top in the stack.

package op_codes_stack

import "stack"

func OP_OVER(mainStack * stack.Stack) bool {
	var verdict bool
	var dataItem stack.Item
	dataItem, verdict = mainStack.AtNthToTop(2)
	if verdict == true {
		mainStack.Push(dataItem)
	}
	return verdict
}