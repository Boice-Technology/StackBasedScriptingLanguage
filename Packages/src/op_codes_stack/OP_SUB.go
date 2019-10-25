// Subtract one input from other.

package op_codes_stack

import ("stack";
		)
		
func OP_SUB(x1 stack.Item, x2 stack.Item) (stack.Item) {
	value1 := x1.(int64)
	value2 := x2.(int64)
	value := value1 - value2
	return stack.Item(value)
}