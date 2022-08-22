package processors

import (
	"context"
	"lowlevelserver/logger"
	"lowlevelserver/utils"
	"net"
	"strconv"
	"time"
)

func ProcesssNotSupportedMethod(ctx context.Context, conn net.Conn) {
	logger.Fetch(ctx).Infow(
		"Run Not Supported Method Processor",
		"Worker", ctx.Value("Worker"),
	)
	headers := map[string]string{
		"Server":     strconv.Itoa(ctx.Value("Worker").(int)),
		"Date":       time.Now().String(),
		"Connection": "close",
	}
	utils.Response405(conn, headers)
}
