package worker

import (
	"context"
	"lowlevelserver/logger"
	"lowlevelserver/processors" // импортируется директория, а не конкретные файлы
	"lowlevelserver/utils"
	"net"
	"sync"
)

func worker(ctx context.Context, wg *sync.WaitGroup, workerId int, pool chan net.Conn) {

	ctx = context.WithValue(ctx, utils.WorkerId, workerId)
	logger.Fetch(ctx).Infow(
		"Worker started",
		"Worker", ctx.Value(utils.WorkerId),
	)
	defer wg.Done()

	for conn := range pool {

		request, err := utils.ReadRequest(conn)
		if err != nil {
			logger.Fetch(ctx).Errorw(
				"Can't read data from connection",
				"Worker", ctx.Value(utils.WorkerId),
				"Error", err,
			)
			continue
		}

		logger.Fetch(ctx).Infow(
			"Got request",
			"Worker", ctx.Value(utils.WorkerId),
			"Request", request,
		)

		method, path := utils.ParseRequest(request)
		switch method {
		case "GET":

			logger.Fetch(ctx).Infow(
				"Got 'GET' method",
				"Worker", ctx.Value(utils.WorkerId),
			)
			if !utils.CheckFileExists(path) {
				logger.Fetch(ctx).Warnw(
					"Path not exists",
					"Worker", ctx.Value(utils.WorkerId),
				)
				processors.ProcesssNotExistedContent(ctx, conn)
				break
			}

			fileType, contentType := utils.DefineContentType(path)
			logger.Fetch(ctx).Infow(
				"Content type is defined",
				"Worker", ctx.Value(utils.WorkerId),
				"Request", request,
				"FileType", fileType,
				"ContentType", contentType,
			)

			switch fileType {
			case "text":
				processors.ProcessText(ctx, conn, path, contentType, true)
			case "image":
				processors.ProcesssImage(ctx, conn, path, contentType, true)
			case "application": // Непонятно как отображать
				processors.ProcesssApplication(ctx, conn, path, contentType, true)
			case "directory":
				processors.ProcesssDirectory(ctx, conn, path, true)
			default:
				// В идеале тут надо какую-нибудь 4хх отправить в ответ.
				logger.Fetch(ctx).Infow(
					"Got unknown type",
					"Worker", ctx.Value(utils.WorkerId),
					"Request", request,
				)
			}

		case "HEAD":

			logger.Fetch(ctx).Infow(
				"Got 'HEAD' method",
				"Worker", ctx.Value(utils.WorkerId),
			)

			if !utils.CheckFileExists(path) {
				logger.Fetch(ctx).Warnw(
					"Path not exists",
					"Worker", ctx.Value(utils.WorkerId),
				)
				processors.ProcesssNotExistedContent(ctx, conn)
				break
			}

			fileType, contentType := utils.DefineContentType(path)
			logger.Fetch(ctx).Infow(
				"Content type is defined",
				"Worker", ctx.Value(utils.WorkerId),
				"Request", request,
				"FileType", fileType,
				"ContentType", contentType,
			)

			switch fileType {
			case "text":
				processors.ProcessText(ctx, conn, path, contentType, false)
			case "image":
				processors.ProcesssImage(ctx, conn, path, contentType, false)
			case "application":
				processors.ProcesssApplication(ctx, conn, path, contentType, false)
			case "directory":
				processors.ProcesssDirectory(ctx, conn, path, false)
			default:
				// В идеале тут надо какую-нибудь 4хх отправить в ответ.
				logger.Fetch(ctx).Infow(
					"Got unknown type",
					"Worker", ctx.Value(utils.WorkerId),
					"Request", request,
				)
			}

		default:
			logger.Fetch(ctx).Infow(
				"Method not supported",
				"Worker", ctx.Value(utils.WorkerId),
				"Request", request,
				"Method", method,
			)
			processors.ProcesssNotSupportedMethod(ctx, conn)
		}

		logger.Fetch(ctx).Infow(
			"Processed",
			"Worker", ctx.Value(utils.WorkerId),
			"Request", request,
		)
		conn.Close()
	}

}

func WorkerPool(ctx context.Context, workersCount int, pool chan net.Conn, wg *sync.WaitGroup) {
	logger.Fetch(ctx).Infow(
		"WorkerPool started",
		"Workers count", workersCount,
	)
	for i := 0; i < workersCount; i++ {
		wg.Add(1)
		go worker(ctx, wg, i, pool)
	}

}

func FillConnectionPool(ctx context.Context, listener net.Listener, pool chan net.Conn) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.Fetch(ctx).Errorw(
				"Can't create connection",
				"Error", err,
			)
			continue
		}
		pool <- conn
	}
}
