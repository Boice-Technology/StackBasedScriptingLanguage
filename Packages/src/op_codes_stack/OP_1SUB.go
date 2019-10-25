// Subtract 1 from the input.

package op_codes_stack

import ("stack";
		)
		
func OP_1SUB(x stack.Item) (stack.Item) {
	value := x.(int64) - 1
	return stack.Item(value)
}