// Return 0 if input is 0 otherwise 1

package op_codes_stack

import ("stack";
		)
		
func OP_0NOTEQUAL(x stack.Item) (stack.Item) {
	value := x.(int64)
	if value!=0 {
		value = 1
	}
	return stack.Item(value)
}