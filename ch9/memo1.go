package main

type result struct {
	value interface{}
	err   error
}

type Memo struct {
	f     Func
	cache map[string]result
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

type Func func(key string) interface{}
