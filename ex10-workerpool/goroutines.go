package goroutines

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

var waitg sync.WaitGroup

func Run(workers int) {
	scan := bufio.NewScanner(os.Stdin)
	id := 1
	scan.Split(bufio.ScanLines)
	task := make(chan string, workers)
	for scan.Scan() {
		if id <= workers {
			go Employee(task, id)
			waitg.Add(1)
			id++
		}
		tmp := string(scan.Bytes())
		task <- tmp
	}
	waitg.Wait()
}

func Employee(task chan string, id int) {
	defer waitg.Done()
	var sleepstr string
	flag := 0
	for {
		if flag == 0 {
			fmt.Printf("worker:%d spawning\n", id)
			flag = 1
		}
		select {
		case sleepstr = <-task:
			fmt.Printf("worker:%d sleep:%s\n", id, sleepstr)
			sleep, _ := time.ParseDuration(sleepstr + "s")
			time.Sleep(sleep)
		default:
			fmt.Printf("worker:%d stopping\n", id)
			return
		}
	}
}
