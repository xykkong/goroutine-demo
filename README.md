# Go routine demo

This is a simple demo of goroutine. This application spawns ten workers (goroutines), where each worker simultaneously searches a stream of pseudo-random data for the string 'Demo', then send a message to the parent using channels.

# How to build
```
  go build
```

# How to run
- Running with default configuration
```
  go run main.go
```

- Running with custom timeout in ms
```
  go run main.go --timeout 30000
```
