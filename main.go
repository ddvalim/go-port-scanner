package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"net"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"
)

type PortScanner struct {
	ip   string              // host IP address
	lock *semaphore.Weighted // semaphore to control the amount of goroutines executing at the same time
}

func NewPortScanner(ip string, lock *semaphore.Weighted) *PortScanner {
	return &PortScanner{
		ip:   ip,
		lock: lock,
	}
}

func (p PortScanner) Scan(port int, timeout time.Duration) {
	target := fmt.Sprintf("%s:%d", p.ip, port)

	conn, err := net.DialTimeout("tcp", target, timeout)
	if err != nil {
		if strings.Contains(err.Error(), "connection refused") {
			fmt.Println(fmt.Sprintf("port %d closed", port))

			return
		}

		panic(err)
	}

	conn.Close()

	fmt.Println(fmt.Sprintf("port %d open", port))

	return
}

func GetLockBasedOnOperationalSystem() int64 {
	out, err := exec.Command("ulimit", "-n").Output()
	if err != nil {
		panic(err)
	}

	s := strings.TrimSpace(string(out))

	max, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}

	return max
}

func main() {
	lock := GetLockBasedOnOperationalSystem()

	// creates a new port scanner with an ip address and a semaphore to control the number of threads executing at the same time
	ps := NewPortScanner("127.0.0.1", semaphore.NewWeighted(lock))

	wg := sync.WaitGroup{}

	defer wg.Done()

	for port := 1; port <= 8080; port++ {
		wg.Add(1)

		// If there is no available unity to be acquired, the goroutine will be blocked by the semaphore until a new unity become available
		ps.lock.Acquire(context.Background(), 1)

		go func(port int) {
			// When a goroutine gets permission to execute, it will liberate a unity for another goroutine to use
			defer ps.lock.Release(1)

			defer wg.Done()

			ps.Scan(port, 800*time.Millisecond)
		}(port)
	}
}
