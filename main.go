package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type TestResult struct {
	TotalRequests      int
	SuccessfulRequests int
	StatusCodes        map[int]int
	TotalTime          time.Duration
}

func main() {
	url := flag.String("url", "", "URL to test")
	requests := flag.Int("requests", 0, "Total number of requests")
	concurrency := flag.Int("concurrency", 0, "Number of concurrent requests")
	flag.Parse()

	if *url == "" || *requests <= 0 || *concurrency <= 0 {
		fmt.Println("Error: All parameters (url, requests, concurrency) are required and must be positive")
		return
	}

	workChan := make(chan struct{}, *requests)
	resultChan := make(chan int, *requests)

	startTime := time.Now()

	var wg sync.WaitGroup
	for i := 0; i < *concurrency; i++ {
		wg.Add(1)
		go worker(*url, workChan, resultChan, &wg)
	}

	for i := 0; i < *requests; i++ {
		workChan <- struct{}{}
	}
	close(workChan)

	wg.Wait()
	close(resultChan)

	duration := time.Since(startTime)
	results := TestResult{
		TotalRequests: *requests,
		StatusCodes:   make(map[int]int),
		TotalTime:     duration,
	}

	for statusCode := range resultChan {
		if statusCode == http.StatusOK {
			results.SuccessfulRequests++
		}
		results.StatusCodes[statusCode]++
	}

	printReport(results)
}

func worker(url string, workChan <-chan struct{}, resultChan chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	for range workChan {
		resp, err := client.Get(url)
		if err != nil {
			resultChan <- 0
			continue
		}
		resultChan <- resp.StatusCode
		resp.Body.Close()
	}
}

func printReport(results TestResult) {
	fmt.Println("\n=== Load Test Report ===")
	fmt.Printf("Total Time: %v\n", results.TotalTime)
	fmt.Printf("Total Requests: %d\n", results.TotalRequests)
	fmt.Printf("Successful Requests (200): %d\n", results.SuccessfulRequests)
	fmt.Println("\nStatus Code Distribution:")
	for code, count := range results.StatusCodes {
		if code == 0 {
			fmt.Printf("Failed Requests (Error): %d\n", count)
		} else {
			fmt.Printf("Status %d: %d\n", code, count)
		}
	}
	fmt.Printf("\nRequests per second: %.2f\n", float64(results.TotalRequests)/results.TotalTime.Seconds())
}
