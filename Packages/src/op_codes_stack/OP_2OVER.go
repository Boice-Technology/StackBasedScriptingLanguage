// Copy the pair of items two spaces back in the stack to the front.

package op_codes_stack

import "stack"

func OP_2OVER(mainStack * stack.Stack) bool {
	dataItem1, present := mainStack.AtNthToTop(4)
	if !present {
		return false
	}
	dataItem2 := mainStack.AtNthToTop(3)
	mainStack.Push(dataItem1)
	mainStack.Push(dataItem2)
	return true
}