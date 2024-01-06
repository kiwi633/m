package object

func NewEnclosedEnvironment(outer *Envir) *Envir {
	env := NewEnvironment()
	env.outer = outer
	return env
}

func NewEnvironment() *Envir {
	s := make(map[string]Object)
	return &Envir{store: s, outer: nil}
}

type Envir struct {
	store map[string]Object
	outer *Envir
}

func (e *Envir) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

func (e *Envir) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
