package env

type Environment struct {
	values map[string]interface{}
	parent *Environment
}

func New() *Environment {
	return newWithParent(nil)
}

func (env *Environment) Child() *Environment {
	return newWithParent(env)
}

func (env *Environment) Get(name string) interface{} {
	value, ok := env.values[name]
	if !ok && env.parent != nil {
		return env.parent.Get(name)
	}
	return value
}

func (env *Environment) Set(name string, value interface{}) {
	env.values[name] = value
}

func newWithParent(parent *Environment) *Environment {
	return &Environment{
		values: map[string]interface{}{},
		parent: parent,
	}
}
