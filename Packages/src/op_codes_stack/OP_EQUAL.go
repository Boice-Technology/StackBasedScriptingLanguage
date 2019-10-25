// Check whether two inputs are equal.

package op_codes_stack
		
import ("stack";
		)
		
func OP_EQUAL(x1 stack.Item, x2 stack.Item) bool {
	return x1==x2
}