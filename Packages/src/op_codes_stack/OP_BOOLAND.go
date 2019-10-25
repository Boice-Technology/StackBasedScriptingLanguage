// If both the inputs are not 0 then return 1 otherwise 0.

package op_codes_stack

import ("stack";
		)
		
func OP_BOOLAND(x1 stack.Item, x2 stack.Item) (stack.Item) {
	value1 := x1.(int64)
	value2 := x2.(int64)
	var value int64 = 0
	if value1 != 0 && value2 != 0 {
		value = 1
	}
	return stack.Item(value)
}