// Return the smaller of two inputs.

package op_codes_stack

import ("stack";
		)
		
func OP_MIN(x1 stack.Item, x2 stack.Item) (stack.Item) {
	value1 := x1.(int64)
	value2 := x2.(int64)
	value := value2	
	if value1 < value2 {
		value = value1
	}
	return stack.Item(value)
}