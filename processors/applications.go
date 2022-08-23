package processors

import (
	"context"
	"lowlevelserver/logger"
	"lowlevelserver/utils"
	"net"
	"strconv"
	"time"
)

func ProcesssApplication(ctx context.Context, conn net.Conn, path string, contentType string, sendBody bool) {
	logger.Fetch(ctx).Infow(
		"Run Application Processor",
		"Worker", ctx.Value(utils.WorkerId),
	)
	content, err := utils.ReadByteContent(path)
	if err != nil {
		// Тут в идеале надо 5хх
		logger.Fetch(ctx).Errorw(
			"Can't read byte content",
			"Worker", ctx.Value(utils.WorkerId),
			"Error", err,
		)
	}

	headers := map[string]string{
		"Content-Length: ": strconv.Itoa(len(content)),
		"Content-Type: ":   contentType,
		"Server: ":         strconv.Itoa(ctx.Value(utils.WorkerId).(int)),
		"Date: ":           time.Now().String(),
		"Connection:":      "close",
	}

	utils.Response200Bytes(conn, content, headers, sendBody)
}
