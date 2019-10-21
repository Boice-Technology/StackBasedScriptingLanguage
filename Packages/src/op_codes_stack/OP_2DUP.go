// Duplicate the top two stack items.

package op_codes_stack

import "stack"

func OP_2DUP(mainStack *stack.Stack) (bool){
	dataItem1, present := mainStack.AtNthToTop(2)
	if !present {
		return false
	}
	dataItem2 := mainStack.AtNthToTop(1)
	mainStack.Push(dataItem1)
	mainStack.Push(dataItem2)
	return true
}