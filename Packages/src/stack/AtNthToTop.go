// Return item at Nth to top in the stack.

package stack

func (s *Stack) AtNthToTop(N int) (Item, bool){
	var dataItem Item
	var verdict bool
	if N > int(s.height) {
		dataItem = nil
		verdict = false
	} else {
		dataItem = s.datastore[int(s.height) - N]
		verdict = true
	}
	return dataItem, verdict
}