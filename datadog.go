package golang_datadog

import (
	"net/http"
	"sync"
	"fmt"
	"time"
)

type DataDog struct {
	sync.Once

	timestamp time.Time
	ttl       uint64
	status    bool
}

func NewDataDog(ttl uint64) *DataDog {
	return &DataDog{
		timestamp: time.Now(),
		ttl:       ttl,
		status:    false}
}

func (d *DataDog) Handle(host string, port uint64, router string) {
	d.Once.Do(func() {
		addr := fmt.Sprintf("%s:%d", host, port)

		http.HandleFunc(router, func(w http.ResponseWriter, r *http.Request) {
			if d.isOk() {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("true"))
			} else {
				w.WriteHeader(http.StatusNoContent)
				w.Write([]byte("false"))
			}
		})

		go http.ListenAndServe(addr, nil)
	})
}

func (d *DataDog) isOk() bool {
	return d.status &&
		d.timestamp.Add(time.Duration(d.ttl) * time.Second).After(time.Now())
}

func (d *DataDog) TTL(ttl uint64) {
	d.ttl = ttl
}

func (d *DataDog) SetStatus(status bool) {
	d.status = status
	d.timestamp = time.Now()
}
