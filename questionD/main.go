package main

import log "github.com/sirupsen/logrus"

var (
	end = make(chan bool)
)

func main() {
	sampleNums := []int{1, 2, 3, 4, 5}

	inputChan := make(chan int)
	responseChan := make(chan int, len(sampleNums))

	CreateWorkerInstances(inputChan, responseChan, 2)

	go listenToResponse(responseChan, len(sampleNums))
	sum(inputChan, sampleNums)

	<-end
	close(inputChan)
	close(responseChan)
	close(end)
}

func CreateWorkerInstances(inputChan, responseChan chan int, instanceCount int) {
	for i := 0; i < instanceCount; i++ {
		w := NewWorker(inputChan, responseChan)
		go w.Start()
	}
}

func listenToResponse(responseChan chan int, numCount int) {
	total := 0
LOOP:
	for {
		select {
		case num := <-responseChan:
			total += num
			numCount--
			if numCount == 0 {
				break LOOP
			}
		}
	}

	printOutput(total)
}

func sum(inputChan chan int, nums []int) {
	for _, num := range nums {
		inputChan <- num
	}
}

func printOutput(total int) {
	log.Infof("sum of numbers: %v", total)
	end <- true
}
