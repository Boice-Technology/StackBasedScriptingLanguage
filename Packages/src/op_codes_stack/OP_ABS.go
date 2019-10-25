// To return absolute value of a integer

package op_codes_stack

import ("stack";
		)
		
func OP_ABS(x stack.Item) (stack.Item) {
	value := x.(int64)
	if value<0 {
		value = -value
	}
	return stack.Item(value)
}