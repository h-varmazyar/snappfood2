package main

type Worker struct {
	responseChan chan int
	inputChan    chan int
}

func NewWorker(inputChan, responseChan chan int) *Worker {
	return &Worker{
		responseChan: responseChan,
		inputChan:    inputChan,
	}
}

func (w *Worker) Start() {
	for {
		select {
		case <-end:
			return
		case num := <-w.inputChan:
			w.calculator(num)
		}
	}
}

func (w *Worker) calculator(num int) {
	resp := num * num

	w.responseChan <- resp
}
