// Puts number of items in the stack onto the stack.

package op_codes_stack

import "stack"

func OP_DEPTH(mainStack *stack.Stack) {
	height := mainStack.Height()
	mainStack.Push(height)
}