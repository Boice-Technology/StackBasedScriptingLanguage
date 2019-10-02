// Peek stack.

package stack

func (s *Stack) Peek() (Item, bool) {
	var top Item
	var verdict bool
	if(s.IsEmpty()){
		top, verdict = nil, false
	} else {
		top, verdict = s.datastore[s.height - 1], true
	}
	return top, verdict
}