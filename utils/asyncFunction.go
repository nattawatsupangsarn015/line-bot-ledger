package utils

import "sync"

type AsyncFunc func() (interface{}, error)

func executeAsyncFunc(fn AsyncFunc, index int, results chan<- interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	res, err := fn()
	if err != nil {
		// handle the error appropriately
		return
	}
	// send the result and index to the results channel
	results <- struct {
		Index int
		Value interface{}
	}{Index: index, Value: res}
}

func functionWaitGroup(fns []AsyncFunc) ([]interface{}, error) {
	// Create a slice to store the results
	results := make([]interface{}, len(fns))

	// Create a channel to receive results
	resultCh := make(chan interface{}, len(fns))

	// Create a WaitGroup to wait for all operations to complete
	var wg sync.WaitGroup
	wg.Add(len(fns))

	// Perform each asynchronous operation concurrently
	for i, fn := range fns {
		go executeAsyncFunc(fn, i, resultCh, &wg)
	}

	// Wait for all operations to complete and collect the results
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	for res := range resultCh {
		// parse the result and store it in the results slice
		r := res.(struct {
			Index int
			Value interface{}
		})
		results[r.Index] = r.Value
	}

	return results, nil
}
