// Copy the item at the top of the stack and insert it before the second-to-top item.

package op_codes_stack

import "stack"

func OP_TUCK(mainStack * stack.Stack) bool {
	_, present := mainStack.AtNthToTop(2)
	if !present {
		return false
	}
	dataItem2,_ := mainStack.Pop()
	dataItem1,_ := mainStack.Pop()
	mainStack.Push(dataItem2)
	mainStack.Push(dataItem1)
	mainStack.Push(dataItem2)
	return true
}