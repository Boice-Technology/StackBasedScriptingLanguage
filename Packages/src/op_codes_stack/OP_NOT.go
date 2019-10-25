// If number is 0 flip it to 1 otherwise 0

package op_codes_stack

import ("stack";
		)
		
func OP_NOT(x stack.Item) (stack.Item) {
	value := x.(int64)
	if value==0 {
		value = 1
	} else {
		value = 0
	}
	return stack.Item(value)
}