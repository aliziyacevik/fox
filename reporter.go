package main

import "fmt"

type Err struct {
	info   string
	line   int
	offset int
}

func NewError(info string, line int, offset int) *Err {
	return &Err{
		info:   info,
		line:   line,
		offset: offset,
	}
}

type Errs []*Err

func (e *Err) Error() string {
	return fmt.Sprintf("Error: %s at: [ offset:line ] [ %d:%d ]", e.info, e.line, e.offset)
}

type Reporter struct {
	errs Errs
}

func (r *Reporter) Report(err *Err) {
	r.errs = append(r.errs, err)
}

func (r *Reporter) ReportInfo(info string, line, offset int) {
	er := NewError(info, line, offset)

	r.Report(er)
}

func (r *Reporter) ReportInfoStream(line, offset int, info string, args ...interface{}) {
	formattedInfo := fmt.Sprintf(info, args...)
	er := NewError(formattedInfo, line, offset)

	r.Report(er)
}

func (r *Reporter) CountErrors() int {
	return len(r.errs)
}

func (r *Reporter) Error() {
	if len(r.errs) == 0 {
		fmt.Println("There are no errors")
		return
	}

	for _, err := range r.errs {
		fmt.Println(err.Error())
	}
}

func NewReporter() *Reporter {
	return &Reporter{}
}
