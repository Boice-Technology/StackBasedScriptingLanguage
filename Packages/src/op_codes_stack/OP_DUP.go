// Duplicates top stack item.

package op_codes_stack

import "stack"

func OP_DUP(mainStack *stack.Stack) (bool){
	var verdict bool
	var dataItem stack.Item
	dataItem, verdict = mainStack.Peek()
	if verdict {
		mainStack.Push(dataItem)
	}
	return verdict
}