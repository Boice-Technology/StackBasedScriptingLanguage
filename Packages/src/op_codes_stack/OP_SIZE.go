// Push the length of the string on the top of the stack.

package op_codes_stack

import ("stack";
		"fmt";
		)
		
func OP_SIZE(s *stack.Stack) bool {
	var verdict bool
	var value stack.Item
	value, verdict = s.Peek()
	if verdict {
		var valueString string = fmt.Sprint(value)
		s.Push(len(valueString))
	}
	return verdict
}