package main

type Parser interface {
	Parse()
}

type parser struct{}

func NewParser() Parser {
	return &parser{}
}

func (p *parser) Parse() {

}
