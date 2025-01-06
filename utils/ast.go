package utils

type NodeKind int

const (
    NK_X NodeKind = iota
    NK_Y
    NK_ADD
    NK_MUL
	NK_TRIPLE
)

type Node struct {
	kind NodeKind
	as NodeAs
}

type NodeTriple struct {
	first *Node
	second *Node
	third *Node
}


type NodeBinop struct {
	lhs *Node
	rhs *Node
}

type NodeAs struct {
	number float32
	binop NodeBinop
	triple NodeTriple
}
