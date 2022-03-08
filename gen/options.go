package gen

//go:generate genopts --outfile=gen/options.go "force"

type Option func(*optionImpl)

type Options interface {
	Force() bool
}

func Force(force bool) Option {
	return func(opts *optionImpl) {
		opts.force = force
	}
}

type optionImpl struct {
	force bool
}

func (o *optionImpl) Force() bool { return o.force }

func makeOptionImpl(opts ...Option) *optionImpl {
	res := &optionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeOptions(opts ...Option) Options {
	return makeOptionImpl(opts...)
}
