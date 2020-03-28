package http

type nodeType uint8

const (
	static nodeType = iota
	root 
	param
	cart
)

type node struct {
	path string
	indices string
	children []*node
	handlers []HandlerFunc
	priority uint32
	nType nodeType
	maxParams uint8
	wildChild bool
}


type Param struct {
	Key   string
	Value string
}

type Params []Param

type methodTree struct {
	method string
	root *node
}
type methodTrees []methodTree
