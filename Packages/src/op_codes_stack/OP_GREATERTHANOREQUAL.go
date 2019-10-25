// If first input number is greater than or equal to the other then return 1 otherwise 0.

package op_codes_stack

import ("stack";
		)
		
func OP_GREATERTHANOREQUAL(x1 stack.Item, x2 stack.Item) (stack.Item) {
	value1 := x1.(int64)
	value2 := x2.(int64)
	value := 0
	if value1 >= value2 {
		value = 1
	}
	return stack.Item(value)
}