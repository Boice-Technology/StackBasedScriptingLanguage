// Remove top item from the alt stack and put it on to main stack.

package op_codes_stack

import "stack"

func OP_FROMALTSTACK(mainStack *stack.Stack, altStack *stack.Stack) bool {
	var verdict bool
	if(altStack.IsEmpty()){
		verdict = false
	} else {
		verdict = true
		var dataItem stack.Item
		dataItem,_ = altStack.Pop()
		mainStack.Push(dataItem)
	}
	return verdict
}