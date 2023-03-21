package ast

type Element interface {
	ElementType() ElementType
}

type ElementType int

const (
	_ ElementType = iota
	ElementTypeAtom
	ElementTypeLiteral
	ElementTypeList
	ElementTypeSetq
	ElementTypeFunc
	ElementTypeLambda
	ElementTypeProg
	ElementTypeCond
	ElementTypeWhile
	ElementTypeReturn
	ElementTypeBreak
)

type LiteralType int

const (
	_ LiteralType = iota
	LiteralTypeInteger
	LiteralTypeReal
	LiteralTypeBoolean
)

type Atom struct {
	Name string
}

type Literal struct {
	Value string
	Type  LiteralType
}

type List struct {
	Elements []Element
}

func (a Atom) ElementType() ElementType {
	return ElementTypeAtom
}

func (l Literal) ElementType() ElementType {
	return ElementTypeLiteral
}

func (l List) ElementType() ElementType {
	return ElementTypeList
}

type Setq struct {
	Atom    Atom
	Element Element
}

type Func struct {
	Atom    Atom
	List    List
	Element Element
}

type Lambda struct {
	List    List
	Element Element
}

type Prog struct {
	List    List
	Element Element
}

type Cond struct {
	List     List
	Element1 Element
	Element2 Element
}

type While struct {
	Element1 Element
	Element2 Element
}

type Return struct {
	Element Element
}

type Break struct {
}

func (s Setq) ElementType() ElementType   { return ElementTypeSetq }
func (f Func) ElementType() ElementType   { return ElementTypeFunc }
func (l Lambda) ElementType() ElementType { return ElementTypeLambda }
func (p Prog) ElementType() ElementType   { return ElementTypeProg }
func (c Cond) ElementType() ElementType   { return ElementTypeCond }
func (w While) ElementType() ElementType  { return ElementTypeWhile }
func (r Return) ElementType() ElementType { return ElementTypeReturn }
func (b Break) ElementType() ElementType  { return ElementTypeBreak }

type Program struct {
	Elements []Element
}
