package application

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

type Task struct {
	ExpressionID  int     `json:"expression_id"`
	NodeID        int     `json:"node_id"`
	Arg1          float64 `json:"arg1"`
	Arg2          float64 `json:"arg2"`
	Operation     string  `json:"operation"`
	OperationTime int     `json:"operation_time"`
}

type TaskPayload struct {
	Task Task `json:"task"`
}

func fetchTask(orchestratorAddress string) (*Task, error) {
	apiEndpoint := fmt.Sprintf("%s/internal/task", orchestratorAddress)
	resp, err := http.Get(apiEndpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, nil
	}

	var payload TaskPayload
	err = json.NewDecoder(resp.Body).Decode(&payload)
	if err != nil {
		return nil, err
	}

	return &payload.Task, err
}

type Result struct {
	ExpressionID int     `json:"expression_id"`
	NodeID       int     `json:"node_id"`
	Result       float64 `json:"result"`
}

func sendResult(result *Result, orchestratorAddress string) error {
	data, err := json.Marshal(result)
	if err != nil {
		return err
	}

	apiEndpoint := fmt.Sprintf("%s/internal/task", orchestratorAddress)
	resp, err := http.Post(apiEndpoint, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func processTask(task *Task) *Result {
	var result float64

	switch task.Operation {
	case "+":
		result = task.Arg1 + task.Arg2
	case "-":
		result = task.Arg1 - task.Arg2
	case "*":
		result = task.Arg1 * task.Arg2
	case "/":
		if task.Arg2 != 0 {
			result = task.Arg1 / task.Arg2
		} else {
			result = 0
		}
	default:
		result = 0
	}

	time.Sleep(time.Duration(task.OperationTime) * time.Millisecond)

	return &Result{
		ExpressionID: task.ExpressionID,
		NodeID:       task.NodeID,
		Result:       result,
	}
}

func worker(ctx context.Context, config *Config, workerID int) {
	log.Printf("Worker %d started\n", workerID)
	for {
		select {
		case <-ctx.Done():
			log.Printf("Worker %d finished\n", workerID)
			return
		default:
			task, err := fetchTask(config.OrchestratorAddress)
			if err != nil {
				log.Printf("Worker %d failed to fetch task: %v\n", workerID, err)
				time.Sleep(1 * time.Second)
				continue
			}
			if task == nil {
				log.Printf("Worker %d didn't get any tasks\n", workerID)
				time.Sleep(1 * time.Second)
				continue
			}

			result := processTask(task)
			err = sendResult(result, config.OrchestratorAddress)
			if err != nil {
				log.Printf("Worker %d completed task %d", workerID, task.ExpressionID)
			}
		}
	}
}

func (a *Application) RunDaemon() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	for i := 0; i < a.Config.ComputingPower; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			worker(ctx, a.Config, id)
		}(i)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan

	log.Println("Interrupt received, shutting down...")
	cancel()
	wg.Wait()
	log.Println("Turned off down successfully")
}
