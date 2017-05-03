package main

import "sync"

/*
I know it looks weird to have a mock function and we replace it at runtime then call it to do real job, I did this simply because I can.
Please:
1. don't remove the mock function, use it.
2. make sure the function does what it appears to do correctly
3. run go test -race
*/

var mock = func() {
}

func compareOddAndEven(values []int) bool {
	evenCount := 0
	oddCount := 0
	var wg sync.WaitGroup
	var lock sync.Mutex
	for i := 0; i < 100; i++ {
		for _, val := range values {
			wg.Add(1)
			go func(v int) {
				lock.Lock()
				defer lock.Unlock()
				defer wg.Done()
				if v%2 == 0 {
					mock = func() {
						evenCount++
					}
				} else {
					mock = func() {
						oddCount++
					}
				}
				mock()
			}(val)
		}
	}
	wg.Wait()
	return evenCount == oddCount
}
