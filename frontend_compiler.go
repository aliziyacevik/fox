package main

type FrontendCompiler interface {
	Scanner
	Parser
	Run()
}

type frontendCompiler struct {
	scanner  Scanner
	parser   Parser
	reporter Reporter
}

func NewFrontendCompiler(s Scanner, p Parser, r Reporter) FrontendCompiler {
	return &frontendCompiler{
		scanner:  s,
		parser:   p,
		reporter: r,
	}
}

func NewFrontendCompilerBySource(source string) FrontendCompiler{
	s := NewScanner(source)
	p := NewParser()
	r := NewReporter()

	return &frontendCompiler{
		scanner:  s,
		parser:   p,
		reporter: r,
	}
}

func (f *frontendCompiler) Run() {}
func (f *frontendCompiler) Scan() Tokens {return nil}
func (f *frontendCompiler) Parse(){}
