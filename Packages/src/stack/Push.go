// Push data-items to the stack.

package stack

func (s *Stack) Push(dataItems ...Item) int64 {
	s.datastore = append(s.datastore, dataItems...)
	s.height += int64(len(dataItems))
	return s.height
}