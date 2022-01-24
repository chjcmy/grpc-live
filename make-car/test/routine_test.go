package test

import (
	"fmt"
	"testing"
	"time"
)

var startTime = time.Now()

type Task struct {
	// task payload
}

func (t *Task) Finish() {
	// what actually can finish the task
}

func AssignTasks(ch chan *Task) {
	// we have 100 tasks to assigned.
	// all we need to do is throw them all to the channel.
	// when we finish, close the channel, the read end of channel will stop  reading.
	for i := 0; i < 100; i++ {
		task := new(Task)
		ch <- task
	}
	close(ch)
}

func WorkerFinishTasks(ch chan *Task, done chan struct{}) {
	// Worker doesn't care what and how many tasks he do.
	// He just finish all the tasks as possible as he can.
	// when the channel closed, he will know all the task done.
	for task := range ch {
		task.Finish()
	}

	// send a signal the indicate that his work is done.
	done <- struct{}{}
}

func TestRoutine(t *testing.T) {
	workerNum := 4
	ch := make(chan *Task, workerNum)

	// use channel as an signal here, the struct{} type take zero byte in memory.
	done := make(chan struct{}, workerNum)

	go AssignTasks(ch)
	for i := 0; i < workerNum; i++ {
		go WorkerFinishTasks(ch, done)
	}

	// we hire 4 worker to work, so we need to receive 4 signal when they all done.
	for i := 0; i < workerNum; i++ {
		<-done
	}

	duration := time.Now().Sub(startTime)

	fmt.Println(duration)
}
