package main

import "fmt"

func main() {

	tasks := make(chan int, 45)   // store the numbers we will calculate
	results := make(chan int, 45) // store the results we are going to calculate

	//You can call the worker even before you have coupled the task channel
	go worker(tasks, results) // each additional worker gets faster using competition
	go worker(tasks, results) // we have 4 processes doing the same thing at the same time
	go worker(tasks, results)
	go worker(tasks, results) //I can have two because of the goroutine
	//now we have two processes that will be pulling numbers from the queue and executing the executions

	for k := 0; k < 45; k++ {
		tasks <- k
	}

	close(tasks) //closing task channel after for is finished

	for j := 0; j < 45; j++ {
		result := <-results
		fmt.Println(result)
	}

}

// we can specify what this channel does in the function
// can be a channel that only sends data or only receives data
// it doesn't have to be a channel that does both

// ## Channel that only receives data using the "arrow" in that position
// var <-chan type channel that only receives data
// var chan<- type channel that only sends data
func worker(tasks <-chan int, results chan<- int) { //func default worker
	for task := range tasks { //each task is a number
		results <- fibonacci(task)
	}
}

func fibonacci(position int) int {
	if position <= 1 {
		return position
	}

	return fibonacci(position-2) + fibonacci(position-1)
}
