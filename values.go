package goweb

import ()

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

func (v Values) Has(key string) bool {
	_, ok := v[key]
	return ok
}

func (v Values) HasId() bool  { return v.Has("id") }
func (v Values) HasId1() bool { return v.Has("id1") }
func (v Values) HasId2() bool { return v.Has("id2") }

func (v Values) GetVal(key string) string {
	val, _ := v.Get(key)
	return val
}

func (v Values) GetId() string  { return v.GetVal("id") }
func (v Values) GetId1() string { return v.GetVal("id1") }
func (v Values) GetId2() string { return v.GetVal("id2") }
