package main

import (
	"context"
	"lowlevelserver/logger"
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

	ctx := context.Background()
	ctx = logger.SetLogger(ctx)
	defer logger.Fetch(ctx).Sync()

	logger.Fetch(ctx).Infow("Start")
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(PORT))
	if err != nil {
		logger.Fetch(ctx).Fatalw(
			"Can't create a listener",
			"Error", err,
		)
		return
	}
	defer listener.Close()

	wg := sync.WaitGroup{}
	connPool := make(chan net.Conn, CONNECTION_POOL_SIZE)
	defer close(connPool)

	go worker.FillConnectionPool(ctx, listener, connPool)
	worker.WorkerPool(ctx, WORKERS_COUNT, connPool, &wg)

	wg.Wait()

}
