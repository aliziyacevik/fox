package main

import "fmt"

type Err struct {
	info string
	line int
}

func NewError(info string, line int) *Err {
	return &Err{
		info: info,
		line: line,
	}
}

type Errs []Err

func (e *Err) Error() string {
	return fmt.Sprintf("Error at line: %d. Error: %s", e.line, e.info)
}

type Reporter interface {
	Report(Err)
	Show()
}

type reporter struct {
	errs Errs
}

func (r *reporter) Report(err Err) {
	r.errs = append(r.errs, err)
}

func (r *reporter) Show() {
	for _, err := range r.errs {
		fmt.Println(err.Error())
	}
}

func NewReporter() Reporter {
	return &reporter{}
}
