package goroutines

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

func Scheduler(id int, wg *sync.WaitGroup, jobs <-chan string) {
	defer wg.Done()
	if len(jobs) > 0 {
		fmt.Printf("worker:%d spawning\n", id)
		for i := range jobs {
			task_time, _ := time.ParseDuration(i + "s")
			fmt.Printf("worker:%d sleep:%s\n", id, i)
			time.Sleep(task_time)
		}
		fmt.Printf("worker:%d stopping\n", id)
	}

}

func Run(poolSize int) {
	var finish_all sync.WaitGroup
	workers := make(chan string, poolSize)
	id := 1
	sscan := bufio.NewScanner(os.Stdin)
	for sscan.Scan() {
		sleep_time := sscan.Text()
		workers <- sleep_time
		if id <= poolSize {
			finish_all.Add(1)
			go Scheduler(id, &finish_all, workers)
			id++
		}
	}
	close(workers)
	finish_all.Wait()
}
