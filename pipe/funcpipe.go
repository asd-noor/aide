package pipe

import "funcgo/result"

type action func(input any) (any, error)

type funcpipe struct {
	value result.Result[any]
	chain []action
}

func NewFuncPipe(value any) *funcpipe {
	fp := &funcpipe{}

	switch v := value.(type) {
	case result.Result[any]:
		fp.value = v
	default:
		fp.value = *(result.New(v))
	}

	return fp
}

// Next simply stores the chain of action steps
func (p *funcpipe) Next(f action) *funcpipe {
	p.chain = append(p.chain, f)
	return p
}

// Exec executes the chain, or cuts it early in case of an error
func (p *funcpipe) Exec() result.Result[any] {
	for _, fn := range p.chain {
		v, e := p.value.Unwrap()
		if e != nil {
			break
		}

		x, e := fn(v)
		p.value = *(result.New(x).SetErr(e))
	}

	return p.value
}

// ExecUnwrap executes the Exec and unwraps the Result
func (p *funcpipe) ExecUnwrap() (any, error) {
	r := p.Exec()
	return r.Unwrap()
}
