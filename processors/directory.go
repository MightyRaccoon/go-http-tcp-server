package processors

import (
	"context"
	"lowlevelserver/logger"
	"lowlevelserver/utils"
	"net"
	"strconv"
	"time"
)

func ProcesssDirectory(ctx context.Context, conn net.Conn, path string, sendBody bool) {
	logger.Fetch(ctx).Infow(
		"Run Directory Processor",
		"Worker", ctx.Value("Worker"),
	)

	fullPath := path + "/index.html"

	//Если файла не существует, то отдает 403
	if !utils.CheckFileExists(fullPath) {
		logger.Fetch(ctx).Warnw(
			"Path not exists",
			"Worker", ctx.Value("Worker"),
			"Path", fullPath,
		)
		headers := map[string]string{
			"Server: ":     strconv.Itoa(ctx.Value("Worker").(int)),
			"Date: ":       time.Now().String(),
			"Connection: ": "close",
		}
		utils.Response403(conn, headers)
		return
	}

	// Если файл существует, то пытаемся обработать
	body, err := utils.ReadTextContent(fullPath)
	if err != nil {
		// Вообще, технически тут должна быть какая-нибудь 5xx
		logger.Fetch(ctx).Errorw(
			"Can't read text content",
			"Worker", ctx.Value("Worker"),
			"Error", err,
		)
		return
	}
	ProcessText(ctx, conn, body, "text/html", true)
}
