package goweb

type Values map[string][]string

func (v Values) Get(key string) (string, bool) {
	if v == nil {
		return "", false
	}
	vs, ok := v[key]
	if !ok || len(vs) == 0 {
		return "", ok
	}
	return vs[0], ok
}

func (v Values) Set(key, value string) {
	v[key] = []string{value}
}

func (v Values) Add(key, value string) {
	v[key] = append(v[key], value)
}

func (v Values) Del(key string) {
	delete(v, key)
}
