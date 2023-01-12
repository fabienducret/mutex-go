# mutex-go

## Example of mutex management with decorator pattern.

First we need to create two tests that can run in paralel and access to the same resource (Book's name):

```go
func TestMutex(t *testing.T) {
	b := library.CreateBook()

	t.Run("Change books's name", func(t *testing.T) {
		t.Parallel()

		for i := 0; i < 1000000; i++ {
			b.ChangeName("test")
		}
	})

	t.Run("Get Book's name", func(t *testing.T) {
		t.Parallel()

		for i := 0; i < 1000000; i++ {
			b.Name()
		}
	})
}
```

Method `CreateBook` simply return an instance of `DefaultBook` struct :

```go
func CreateBook() Book {
	return &DefaultBook{}
}
```

We run them in paralel with the command : `go test --race ./...`
And here it the error :

```
WARNING: DATA RACE
Write at 0x00c00011a530 by goroutine 8:
  library.(*DefaultBook).ChangeName()
      /Users/fab/Projects/mutex-go/library.go:19 +0x34
  library_test.TestMutex.func1()
      /Users/fab/Projects/mutex-go/library_test.go:15 +0x60
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1446 +0x188
  testing.(*T).Run.func1()
      /usr/local/go/src/testing/testing.go:1493 +0x40

Previous read at 0x00c00011a530 by goroutine 9:
  library.(*DefaultBook).Name()
      /Users/fab/Projects/mutex-go/library.go:15 +0x2c
  library_test.TestMutex.func2()
      /Users/fab/Projects/mutex-go/library_test.go:23 +0x54
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1446 +0x188
  testing.(*T).Run.func1()
      /usr/local/go/src/testing/testing.go:1493 +0x40

Goroutine 8 (running) created at:
  testing.(*T).Run()
      /usr/local/go/src/testing/testing.go:1493 +0x55c
  library_test.TestMutex()
      /Users/fab/Projects/mutex-go/library_test.go:11 +0xe4
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1446 +0x188
  testing.(*T).Run.func1()
      /usr/local/go/src/testing/testing.go:1493 +0x40

Goroutine 9 (running) created at:
  testing.(*T).Run()
      /usr/local/go/src/testing/testing.go:1493 +0x55c
  library_test.TestMutex()
      /Users/fab/Projects/mutex-go/library_test.go:19 +0x164
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1446 +0x188
  testing.(*T).Run.func1()
      /usr/local/go/src/testing/testing.go:1493 +0x40
==================
--- FAIL: TestMutex (0.00s)
    --- FAIL: TestMutex/Set_identity (0.04s)
        testing.go:1319: race detected during execution of test
    --- FAIL: TestMutex/Get_identity (0.06s)
        testing.go:1319: race detected during execution of test
FAIL
FAIL    library 0.452s
FAIL
```

One goroutine changes the book's name while another reads it.
It fails !

Change `CreateBook` to return an instance of MutexBook :

```go
func CreateBook() Book {
	return &MutexBook{original: DefaultBook{}}
}
```

The test doesn't fail anymore :

```
ok      library 1.430s
```
