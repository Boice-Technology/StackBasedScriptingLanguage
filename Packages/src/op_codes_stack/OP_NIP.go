// Removes second to top stack item.

package op_codes_stack

import "stack"

func OP_NIP(mainStack *stack.Stack) bool {
	var verdict bool
	_, verdict = mainStack.RemoveNthToTop(2)
	return verdict
}