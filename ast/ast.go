package ast

type Element interface {
	ElementType() ElementType
	Eval(*Context) Element
}

type List interface {
	Element
	GetElements() []Element
}

type Callable interface {
	Call(*Context, []Element) Element
}

type ElementType int

const (
	_ ElementType = iota
	ElementTypeAtom
	ElementTypeLiteral
	ElementTypeList
	ElementTypeProgram

	keywords
	ElementTypeQuote
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
	LiteralTypeList
)

type Atom struct {
	Name string
}

type Literal interface {
	Element
	Type() LiteralType
}

type LiteralInteger struct {
	Value int64
}

type LiteralReal struct {
	Value float64
}

type LiteralBoolean struct {
	Value bool
}

type LiteralNull struct {
	Value interface{}
}

type LiteralList struct {
	Value []Literal
}

type ListElement struct {
	Elements []Element
}

type Program struct {
	Elements []Element
}

func (a Atom) ElementType() ElementType           { return ElementTypeAtom }
func (l LiteralInteger) ElementType() ElementType { return ElementTypeLiteral }
func (l LiteralReal) ElementType() ElementType    { return ElementTypeLiteral }
func (l LiteralBoolean) ElementType() ElementType { return ElementTypeLiteral }
func (l LiteralNull) ElementType() ElementType    { return ElementTypeLiteral }
func (l LiteralList) ElementType() ElementType    { return ElementTypeList }
func (l ListElement) ElementType() ElementType    { return ElementTypeList }
func (l ListElement) GetElements() []Element      { return l.Elements }
func (p Program) ElementType() ElementType        { return ElementTypeProgram }

func (l LiteralInteger) Type() LiteralType { return LiteralTypeInteger }
func (l LiteralReal) Type() LiteralType    { return LiteralTypeReal }
func (l LiteralBoolean) Type() LiteralType { return LiteralTypeBoolean }
func (l LiteralNull) Type() LiteralType    { return LiteralTypeNull }
func (l LiteralList) Type() LiteralType    { return LiteralTypeList }

type Quote struct {
	Element Element
}

type Setq struct {
	Atom    Atom
	Element Element
}

type Func struct {
	Atom    Atom
	List    List
	SubProg Program
}

type Lambda struct {
	List    List
	SubProg Program
}

type Prog struct {
	List    List
	SubProg Program
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

func (q Quote) ElementType() ElementType  { return ElementTypeQuote }
func (s Setq) ElementType() ElementType   { return ElementTypeSetq }
func (f Func) ElementType() ElementType   { return ElementTypeFunc }
func (l Lambda) ElementType() ElementType { return ElementTypeLambda }
func (p Prog) ElementType() ElementType   { return ElementTypeProg }
func (c Cond) ElementType() ElementType   { return ElementTypeCond }
func (w While) ElementType() ElementType  { return ElementTypeWhile }
func (r Return) ElementType() ElementType { return ElementTypeReturn }
func (b Break) ElementType() ElementType  { return ElementTypeBreak }

func (q Quote) GetElements() []Element  { return []Element{q.Element} }
func (s Setq) GetElements() []Element   { return []Element{s.Element} }
func (f Func) GetElements() []Element   { return []Element{f.List} }
func (l Lambda) GetElements() []Element { return []Element{l.List} }
func (p Prog) GetElements() []Element   { return []Element{p.List} }
func (c Cond) GetElements() []Element   { return []Element{c.List, c.Element1, c.Element2} }
func (w While) GetElements() []Element  { return []Element{w.Element1, w.Element2} }
func (r Return) GetElements() []Element { return []Element{r.Element} }
func (b Break) GetElements() []Element  { return []Element{} }

//
//func IsKeyword(e ElementType) bool {
//	return e >= keywords && e <= endkeywords
//}
