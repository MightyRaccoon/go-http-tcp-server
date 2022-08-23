package processors

import (
	"context"
	"lowlevelserver/logger"
	"lowlevelserver/utils"
	"net"
	"strconv"
	"time"
)

func ProcesssNotExistedContent(ctx context.Context, conn net.Conn) {
	logger.Fetch(ctx).Infow(
		"Run Not Existed Content Processor",
		"Worker", ctx.Value(utils.WorkerId),
	)
	headers := map[string]string{
		"Server: ":     strconv.Itoa(ctx.Value(utils.WorkerId).(int)),
		"Date: ":       time.Now().String(),
		"Connection: ": "close",
	}
	utils.Response404(conn, headers)
}
