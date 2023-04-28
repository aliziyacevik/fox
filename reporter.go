package main

import "fmt"

type Err struct {
	info string
}

func NewError(info string) *Err {
	return &Err{
		info: info,
	}
}

type Errs []*Err

func (e *Err) Error() string {
	return fmt.Sprintf("Error: %s", e.info)
}

type Reporter struct {
	errs Errs
}

func (r *Reporter) Report(err *Err) {
	r.errs = append(r.errs, err)
}

func (r *Reporter) ReportInfo(info string) {
	er := NewError(info)

	r.Report(er)
}

func (r *Reporter) ReportStream(info string, stream ...interface{}) {
	infoFormatted := fmt.Sprintf(info, stream...)
	er := NewError(infoFormatted)
	
	r.Report(er)
}

func (r *Reporter) CountErrors() int{
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
