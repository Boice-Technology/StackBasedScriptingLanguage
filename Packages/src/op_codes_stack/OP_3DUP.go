// Duplicate the top three stack items.

package op_codes_stack

import "stack"

func OP_3DUP(mainStack *stack.Stack) (bool){
	dataItem1, present := mainStack.AtNthToTop(3)
	if !present {
		return false
	}
	dataItem2,_ := mainStack.AtNthToTop(2)
	dataItem3,_ := mainStack.AtNthToTop(1)
	mainStack.Push(dataItem1)
	mainStack.Push(dataItem2)
	mainStack.Push(dataItem3)
	return true
}