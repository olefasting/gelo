package gelo

type Word interface {
	Ser() Symbol
	Copy() Word     //these two should be in a MutableWord interface since
	DeepCopy() Word //a lot of the intrinsic types are immutable
	Equals(Word) bool
	Type() Symbol
}

type Symbol interface {
	Word
	Bytes() []byte
	Runes() []rune
	String() string
	interned() bool
}

type Quote interface {
	Word
	unprotect() *quote
}

type Port interface {
	Word
	Send(Word)
	Recv() Word
	Close()
	Closed() bool
}

type Error interface {
	Word
	error
	From() uint32
	_tag()
}

type Bool bool

type Number struct {
	num float64
	ser []byte
}

type List struct {
	Value Word
	Next  *List
}

type Dict struct {
	rep map[string]Word
	ser []byte
}

type Alien func(*VM, *List, uint) Word

func (a Alien) Ser() Symbol {
	return a.Type() //we hope some day that reflection can get the name of a
}

func (a Alien) Copy() Word {
	return a
}

func (a Alien) DeepCopy() Word {
	return a
}

func (a Alien) Equals(w Word) bool {
	return false
}

func (Alien) Type() Symbol {
	return interns("*ALIEN*")
}

//defined at the top of vm.go as it is a special internal tag
func (d *defert) Ser() Symbol {
	return d.Type()
}

func (d *defert) Copy() Word {
	return d
}

func (d *defert) DeepCopy() Word {
	return d
}

func (*defert) Equals(w Word) bool {
	_, ok := w.(*defert)
	return ok
}

func (*defert) Type() Symbol {
	return interns("*DEFER*")
}
