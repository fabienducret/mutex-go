# mutex-go

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
