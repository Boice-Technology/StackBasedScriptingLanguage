// Add 1 to the input

package op_codes_stack

import ("stack";
		)
		
func OP_1ADD(x stack.Item) (stack.Item) {
	value := x.(int64) + 1
	return stack.Item(value)
}