// Remove Nth item to top from the stack.

package stack

func (s *Stack) RemoveNthToTop(N int) (Item, bool) {
	var dataItem Item
	var verdict bool
	if N > int(s.height) {
		dataItem = nil
		verdict = false
	} else {
		verdict = true
		dataItem = s.datastore[int(s.height) - N]
		temp := s.datastore[int(s.height) - N + 1:]
		s.datastore = s.datastore[:int(s.height) - N]
		s.datastore = append(s.datastore, temp...)
		s.height -= 1
	}
	return dataItem, verdict
}