package processors

import (
	"context"
	"lowlevelserver/logger"
	"lowlevelserver/utils"
	"net"
	"strconv"
	"time"
)

func ProcessText(ctx context.Context, conn net.Conn, path string, contentType string, sendBody bool) {
	logger.Fetch(ctx).Infow(
		"Run Text Processor",
		"Worker", ctx.Value("Worker"),
	)
	content, err := utils.ReadTextContent(path)
	if err != nil {
		logger.Fetch(ctx).Errorw(
			"Path not exists",
			"Worker", ctx.Value("Worker"),
			"Path", path,
		)
	}

	headers := map[string]string{
		"Content-Length: ": strconv.Itoa(len(content)),
		"Content-Type: ":   contentType,
		"Server: ":         strconv.Itoa(ctx.Value("Worker").(int)),
		"Date: ":           time.Now().String(),
		"Connection: ":     "close",
	}

	utils.Response200Text(conn, content, headers, sendBody)

}
