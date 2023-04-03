package ast

type Element interface {
	ElementType() ElementType
}

type List interface {
	Element
	GetElements() []Element
}

type ElementType int

const (
	_ ElementType = iota
	ElementTypeAtom
	ElementTypeLiteral
	ElementTypeList

	keywords
	ElementTypeSetq
	ElementTypeFunc
	ElementTypeLambda
	ElementTypeProg
	ElementTypeCond
	ElementTypeWhile
	ElementTypeReturn
	ElementTypeBreak
	endkeywords
)

type LiteralType int

const (
	_ LiteralType = iota
	LiteralTypeInteger
	LiteralTypeReal
	LiteralTypeBoolean
	LiteralTypeNull
)

type Atom struct {
	Name string
}

type Literal struct {
	Value string
	Type  LiteralType
}

type ListElement struct {
	Elements []Element
}

func (a Atom) ElementType() ElementType        { return ElementTypeAtom }
func (l Literal) ElementType() ElementType     { return ElementTypeLiteral }
func (l ListElement) ElementType() ElementType { return ElementTypeList }
func (l ListElement) GetElements() []Element   { return l.Elements }

type Setq struct {
	Atom    Atom
	Element Element
}

type Func struct {
	Atom    Atom
	List    ListElement
	Element Element
}

type Lambda struct {
	List    ListElement
	Element Element
}

type Prog struct {
	List    ListElement
	Element Element
}

type Cond struct {
	List     ListElement
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

func (s Setq) GetElements() []Element   { return []Element{s.Element} }
func (f Func) GetElements() []Element   { return []Element{f.List, f.Element} }
func (l Lambda) GetElements() []Element { return []Element{l.List, l.Element} }
func (p Prog) GetElements() []Element   { return []Element{p.List, p.Element} }
func (c Cond) GetElements() []Element   { return []Element{c.List, c.Element1, c.Element2} }
func (w While) GetElements() []Element  { return []Element{w.Element1, w.Element2} }
func (r Return) GetElements() []Element { return []Element{r.Element} }
func (b Break) GetElements() []Element  { return []Element{} }

type Program struct {
	Elements []Element
}

func IsKeyword(e ElementType) bool {
	return e >= keywords && e <= endkeywords
}
