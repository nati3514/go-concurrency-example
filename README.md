# Go Concurrency Example

This repository demonstrates the power of Go's concurrency model by comparing sequential and parallel execution of HTTP requests.

## Features

- Sequential HTTP request execution
- Concurrent HTTP request execution using goroutines and WaitGroup
- Performance comparison between both approaches

## Getting Started

### Prerequisites

- Go 1.16 or higher

### Installation

```bash
git clone https://github.com/nati3514/go-concurrency-example.git
cd go-concurrency-example
```

### Running the Examples

#### Sequential Version

```bash
go run sequential/main.go
```

#### Concurrent Version

```bash
go run concurrent/main.go
```

## Results

The concurrent version typically shows significant performance improvements over the sequential version, especially for I/O-bound operations like HTTP requests.

## Performance Comparison

| Version     | Execution Time (avg) | Speed Improvement |
|-------------|----------------------|-------------------|
| Sequential  | ~1.35s              | 1x                |
| Concurrent  | ~400ms              | ~3.4x faster      |

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.