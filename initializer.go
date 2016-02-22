package main

func Resolve(fn func() error) *Initialier {
	return New().Resolve(fn)
}

func Do(fn func()) *Initialier {
	return New().Do(fn)
}

func New() *Initialier {
	return &Initialier{}
}

type Initialier struct {
	err error
}

func (i *Initialier) Resolve(fn func() error) *Initialier {
	if i.err != nil {
		return i
	}
	i.err = fn()
	return i
}

func (i *Initialier) Do(fn func()) *Initialier {
	return i.Resolve(func() error {
		fn()
		return nil
	})
}

func (i *Initialier) GetError() error {
	return i.err
}
