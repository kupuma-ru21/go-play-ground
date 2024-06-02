package main

type Simpler interface {
	Get() int
	Set(u int)
}

type Simple struct {
	i int
}

func (p *Simple) Get() int {
	return p.i
}

func (p *Simple) Set(u int) {
	p.i = u
}

func fI(it Simpler) int { // function calling both methods through interface
	it.Set(10)
	return it.Get()
}

func main() {
	simple := new(Simple)
	fI(simple)
	println(fI(simple))
}
