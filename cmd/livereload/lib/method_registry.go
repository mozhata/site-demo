package lib

type MethodRegistry struct {
	methods map[string]func()
}

func NewMethodRegistry() *MethodRegistry {
	return &MethodRegistry{
		methods: map[string]func(){},
	}
}

func (r *MethodRegistry) Register(name string, f func()) {
	r.methods[name] = f
}

func (r *MethodRegistry) Run(name string) {
	f := r.methods[name]
	if f == nil {
		panic("Unknown method " + name)
	}
	f()
}
