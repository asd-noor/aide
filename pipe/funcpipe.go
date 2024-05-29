package pipe

type action[T any] func(input T) (T, error)

type Pipe[T any] struct {
	chain []action[T]
}

// Next simply stores the chain of action steps
func (p *Pipe[T]) Next(f action[T]) *Pipe[T] {
	p.chain = append(p.chain, f)
	return p
}

// Do executes the chain, or cuts it early in case of an error
func (p *Pipe[T]) Do() (T, error) {
	var res T
	var err error
	for _, fn := range p.chain {
		res, err = fn(res)
		if err != nil {
			break
		}
	}

	return res, err
}
