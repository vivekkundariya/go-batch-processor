package rate_limit

import (
	"fmt"
	"go.uber.org/ratelimit"
	"log"
	"sync"
	"time"
)

type RateLimitExecutor struct {
	Limit int
}

func (r RateLimitExecutor) Execute(worker RateLimitWorker) {
	workSize := (worker).WorkSize()

	log.Printf("Running ....")
	log.Printf("RateLimit %d", r.Limit)
	log.Printf("Work Size = %d", workSize)

	var index int
	wg := sync.WaitGroup{}
	wg.Add(workSize)

	rl := ratelimit.New(r.Limit)

	for index = 0; index < workSize; index++ {
		rl.Take()
		log.Println(fmt.Sprintf("starting %d", index))
		go singleTask(worker, index, &wg)
	}

	go (worker).HandleOutput()
	wg.Wait()
	worker.Close()
	time.Sleep(2 * time.Second)
}

func singleTask(worker RateLimitWorker, index int, wg *sync.WaitGroup) {
	defer func(g *sync.WaitGroup) {
		//log.Println()
		fmt.Println("Completed %d", index)
		g.Done()
	}(wg)

	input, err := (worker).GetInput(index)
	if err != nil || input == nil {
		log.Printf("Index %d, error in input %s or input is nil ", index, err)
		return
	}
	response, err := (worker).Work(input)
	if err != nil || response == nil {
		log.Printf("Index %d, error in work %s or input is nil ", index, err)
		return
	}
	(worker).ProcessResp(response, input)
}
