// Peek stack.

package stack

func (s *Stack) Peek() (Item, bool) {
	return s.AtNthToTop(1)
}