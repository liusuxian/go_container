package stack

type Stack interface {
	Push(v interface{}) interface{} // add v as element Len()
	Pop() (interface{}, error)      // remove and return element Len() - 1.
	Peek() (interface{}, error)     // return element Len() - 1.
	Empty() bool                    // return true if the stack contains no items, false otherwise
	Search(v interface{}) int       // return The 1 based depth of the element, or -1 if the element is not on the stack
}
