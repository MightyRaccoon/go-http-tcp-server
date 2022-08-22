package processors

import (
	"context"
	"lowlevelserver/logger"
	"lowlevelserver/utils"
	"net"
	"strconv"
	"time"
)

func ProcesssImage(ctx context.Context, conn net.Conn, path string, contentType string, sendBody bool) {
	logger.Fetch(ctx).Infow(
		"Run Image Processor",
		"Worker", ctx.Value("Worker"),
	)

	content, err := utils.ReadByteContent(path)
	if err != nil {
		logger.Fetch(ctx).Errorw(
			"Can't read byte content",
			"Worker", ctx.Value("Worker"),
			"Error", err,
		)
	}

	headers := map[string]string{
		"Content-Length: ": strconv.Itoa(len(content)),
		"Content-Type: ":   contentType,
		"Server: ":         strconv.Itoa(ctx.Value("Worker").(int)),
		"Date: ":           time.Now().String(),
		"Connection:":      "close",
	}

	utils.Response200Bytes(conn, content, headers, sendBody)
}
