package main

type Simpler interface {
	Get() int
	Set(u int)
}

type Simple struct {
	integer int
}

func (p *Simple) Get() int {
	return p.integer
}

func (p *Simple) Set(u int) {
	p.integer = u
}

type RSimple struct {
	integer int
}

func (p *RSimple) Get() int {
	return p.integer
}

func (p *RSimple) Set(u int) {
	p.integer = u
}

func fI(it Simpler) int {
	switch it.(type) {
	case *Simple:
		it.Set(5)
		return it.Get()
	case *RSimple:
		it.Set(10)
		return it.Get()
	}
	return 0
}

func main() {
	sinple := Simpler(new(Simple))
	rSinple := Simpler(new(RSimple))
	println(fI(sinple))
	println(fI(rSinple))
}
