// Flip the sign of the input.

package op_codes_stack

import ("stack";
		)
		
func OP_NEGATE(x stack.Item) (stack.Item) {
	value := - x.(int64)
	return stack.Item(value)
}