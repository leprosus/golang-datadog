# Golang package to response service status for datadog service

## Usage

```go
d := golang_datadog.NewDataDog(300) // Set status time-to-live - 300 seconds
// If set -1 then ttl will be switched off

d.Handle("localhost", 8080, "/status") // Handle interface at localhost:8080 by /status http path (can use curl -X POST localhost:8080/status)

d.SetStatus(status) // To re-new status
// Before 300 seconds interface returns true & 200 status
// After 300 seconds interface returns false & 204 status
```