// Removes the top from the main stack and puts onto the top of the alt stack.

package op_codes_stack

import "stack"

func OP_TOALTSTACK(mainStack *stack.Stack, altStack *stack.Stack) bool {
	var verdict bool
	if(mainStack.IsEmpty()){
		verdict = false
	} else {
		verdict = true
		var top stack.Item
		top, _ = mainStack.Pop()
		altStack.Push(top)
	}
	return verdict
}