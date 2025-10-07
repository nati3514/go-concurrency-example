package main

import (
    "fmt"
    "io"
    "net/http"
    "sync"
    "time"
)

func fetchURL(url string, wg *sync.WaitGroup) {
    defer wg.Done()
    start := time.Now()
    
    // Simulate network latency (happens in parallel)
    time.Sleep(100 * time.Millisecond)
    
    resp, err := http.Get(url)
    if err != nil {
        fmt.Printf("Error fetching %s: %v\n", url, err)
        return
    }
    defer resp.Body.Close()
    
    // Simulate processing time (happens in parallel)
    time.Sleep(100 * time.Millisecond)
    
    _, err = io.Copy(io.Discard, resp.Body)
    if err != nil {
        fmt.Printf("Error reading %s: %v\n", url, err)
        return
    }
    
    // Simulate additional processing (happens in parallel)
    time.Sleep(100 * time.Millisecond)
    
    elapsed := time.Since(start)
    fmt.Printf("Fetched %s in %s\n", url, elapsed)
}

func main() {
    start := time.Now()
    var wg sync.WaitGroup
    
    urls := []string{
        "https://jsonplaceholder.typicode.com/posts/1",
        "https://jsonplaceholder.typicode.com/comments/1",
        "https://jsonplaceholder.typicode.com/albums/1",
        "https://jsonplaceholder.typicode.com/photos/1",
    }

    // Concurrent fetch
    wg.Add(len(urls))
    for _, url := range urls {
        go fetchURL(url, &wg)
    }
    wg.Wait()

    fmt.Printf("Total time taken: %s\n", time.Since(start))
}