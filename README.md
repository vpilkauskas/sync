# sync
[![Go Report Card](https://goreportcard.com/badge/github.com/vpilkauskas/sync)](https://goreportcard.com/report/github.com/vpilkauskas/sync)

This package is inspired by `sync.WaitGroup` idea and `time.Ticker` api

It allows you to add timeout for waiting response from your processes

```go
import (
	"time"

	"github.com/vpilkauskas/sync"
)

func main(){
	m := sync.New()
	m.Add(1)

	timer := time.NewTimer(time.Millisecond * 300)
	defer timer.Stop()

	m.Done()

	select {
	case <-timer.C:
		log.Println("Jobs didn't finished in time")
	case <-m.D:
		log.Println("Finished")
	}
}
```
