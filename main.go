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
	connection_pool := make(chan net.Conn, CONNECTION_POOL_SIZE)
	defer close(connection_pool)

	go worker.FillConnectionPool(listener, connection_pool)
	worker.WorkerPool(WORKERS_COUNT, connection_pool, &wg)

	wg.Wait()

}
