// Define Stack data structure

package stack 

type item interface{}

type Stack struct{
	datastore []item
	height int64
}