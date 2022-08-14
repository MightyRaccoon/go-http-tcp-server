package main

import (
	"log" //zap
	"lowlevelserver/worker"
	"net"
	"strconv"
	"sync"
)

const (
	PORT                 = 80
	WORKERS_COUNT        = 200
	CONNECTION_POOL_SIZE = 10
)

func main() {

	log.Println("Start")
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(PORT))
	if err != nil {
		log.Fatal("Can't create a listener: ", err)
		return
	}
	defer listener.Close()

	wg := sync.WaitGroup{}
	connPool := make(chan net.Conn, CONNECTION_POOL_SIZE)
	defer close(connPool)

	go worker.FillConnectionPool(listener, connPool)
	worker.WorkerPool(WORKERS_COUNT, connPool, &wg)

	wg.Wait()

}
