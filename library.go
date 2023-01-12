package library

import "sync"

type Book interface {
	Name() string
	ChangeName(i string)
}

type DefaultBook struct {
	name string
}

func (b *DefaultBook) Name() string {
	return b.name
}

func (b *DefaultBook) ChangeName(i string) {
	b.name = i
}

type MutexBook struct {
	sync.RWMutex
	original DefaultBook
}

func (b *MutexBook) Name() string {
	b.RLock()
	defer b.RUnlock()

	return b.original.Name()
}

func (b *MutexBook) ChangeName(i string) {
	b.Lock()
	defer b.Unlock()

	b.original.ChangeName(i)
}

func CreateBook() Book {
	return &MutexBook{original: DefaultBook{}}
	//return &DefaultBook{}
}
