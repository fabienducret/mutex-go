package library

import "sync"

type Book interface {
	Name() string
	ChangeName(i string)
}

// DefaultBook, pure without notion of mutex
type DefaultBook struct {
	name string
}

func (b *DefaultBook) Name() string {
	return b.name
}

func (b *DefaultBook) ChangeName(n string) {
	b.name = n
}

// MutexBook, it goal is to manage mutex
type MutexBook struct {
	sync.RWMutex
	original DefaultBook
}

func (b *MutexBook) Name() string {
	b.RLock()
	defer b.RUnlock()

	return b.original.Name()
}

func (b *MutexBook) ChangeName(n string) {
	b.Lock()
	defer b.Unlock()

	b.original.ChangeName(n)
}

func CreateBook() Book {
	//return &DefaultBook{}
	return &MutexBook{original: DefaultBook{}}
}
