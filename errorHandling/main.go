// package main

// import (
// 	"fmt"
// 	"sync"
// )

// type errorStatus struct {
// 	errorMessage string
// 	errorCode    int
// }

// func main() {
// 	errorch := make(chan errorStatus)
// 	es := errorStatus{}
// 	go es.makeError(true, errorch)
// 	var wg sync.WaitGroup
// 	wg.Add(1) // Add one goroutine to the WaitGroup
// 	go func() {
// 		defer wg.Done() // Decrement the WaitGroup when done
// 		es.makeError(true, errorch)
// 	}()
// 	for {
// 		select {
// 		case err := <-errorch:
// 			fmt.Printf("Error Message is: %s\nError Code is: %d", err.errorMessage, err.errorCode)
// 			return
// 		}
// 	}
// 	wg.Wait()
// }

// func (e *errorStatus) makeError(isError bool, errorch chan errorStatus) {
// 	if isError {
// 		errorpayload := errorStatus{
// 			errorMessage: "My error Type",
// 			errorCode:    123,
// 		}
// 		errorch <- errorpayload
// 	} else {
// 		fmt.Println("Success Message")
// 	}
// }

//Error Handling

package main

import (
	"fmt"
	"strconv"
)

type errorStatus struct {
	errorMessage string
	errorCode    int
}

func main() {
	fmt.Println("Error Handling")
	result, err := makeError(true)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

}

func (e errorStatus) Error() string {
	return e.errorMessage + " " + strconv.Itoa(e.errorCode)
}

func makeError(isError bool) (string, error) {
	if isError {
		return " ", errorStatus{errorMessage: "My error Type", errorCode: 123}
	} else {
		return "Success Message", nil
	}
}
