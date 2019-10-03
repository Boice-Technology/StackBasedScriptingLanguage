// Remove the top stack element.

package op_codes_stack

import "stack"

func OP_DROP(mainStack *stack.Stack) bool {
	var verdict bool
	_, verdict = mainStack.Pop()
	return verdict
}