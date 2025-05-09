package application

import (
	"context"
	pb "github.com/SteeperMold/Calculator-go/orchestrator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"
)

func processTask(task *pb.GetTaskResponse) *pb.PostTaskResult {
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

	return &pb.PostTaskResult{
		ExpressionId: task.ExpressionId,
		NodeId:       task.NodeId,
		Result:       result,
	}
}

func worker(ctx context.Context, client pb.OrchestratorClient, workerID int) {
	log.Printf("Worker %d started\n", workerID)
	for {
		select {
		case <-ctx.Done():
			log.Printf("Worker %d finished\n", workerID)
			return
		default:
		}

		task, err := client.FetchTask(ctx, &pb.Empty{})
		if err != nil {
			log.Printf("Worker %d failed to fetch task: %v\n", workerID, err)
			time.Sleep(1 * time.Second)
			continue
		}

		processedTask := processTask(task)

		_, err = client.SendResult(ctx, processedTask)
		if err != nil {
			log.Printf("Worker %d failed to send result: %v\n", workerID, err)
		} else {
			log.Printf("Worker %d completed task %d", workerID, task.ExpressionId)
		}
	}
}

func (a *Application) RunDaemon() {
	conn, err := grpc.NewClient(
		a.Config.OrchestratorAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to connect orchestrator: %v", err)
	}
	defer conn.Close()

	client := pb.NewOrchestratorClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	for i := 0; i < a.Config.ComputingPower; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			worker(ctx, client, id)
		}(i)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan

	log.Println("Interrupt received, shutting down...")
	cancel()
	wg.Wait()
	log.Println("Turned off successfully")
}
