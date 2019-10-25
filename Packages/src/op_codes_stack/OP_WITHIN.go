// Return 1 if input lies between MIN and MAX (MIN inclusive) otherwise 0.

package op_codes_stack

import ("stack";
		)
		
func OP_WITHIN(x stack.Item, Min stack.Item, Max stack.Item) (stack.Item){
	value1 := x.(int64)
	MinValue := Min.(int64)
	MaxValue := Max.(int64)
	value := 0
	if value1>=MinValue && value1<MaxValue {
		value = 1
	}
	return stack.Item(value)
}