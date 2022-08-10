package worker

import (
	"log"
	"lowlevelserver/processors" // импортируется директория, а не конкретные файлы
	"lowlevelserver/utils"
	"net"
	"strings"
	"sync"
)

func worker(wg *sync.WaitGroup, workerId int, pool chan net.Conn) {

	log.Println("Worker ", workerId, " start")
	defer wg.Done()

	for conn := range pool {

		request, err := utils.ReadRequest(conn)
		if err != nil {
			log.Println("Worker ", workerId, " got error: ", err)
			continue
		}

		log.Println("Worker", workerId, "Got request:", request)

		method, path := utils.ParseRequest(request)
		full_path := "./data" + path
		switch method {
		case "GET":

			log.Println("Worker ", workerId, " Got 'GET' method")
			if !utils.CheckFileExists(full_path) {
				log.Println("Worker ", workerId, " :file not exists: ")
				processors.ProcesssNotExistedContent(conn, workerId)
				break
			}

			contentType := utils.DefineContentType(full_path)
			contentType_first_type := strings.Split(contentType, "/")[0]
			log.Println("Worker ", workerId, " Content type: ", contentType)

			switch contentType_first_type {
			case "text":
				processors.ProcessText(conn, workerId, full_path, contentType, true)
			case "image":
				processors.ProcesssImage(conn, workerId, full_path, contentType, true)
			case "application": // Непонятно как отображать
				processors.ProcesssApplication(conn, workerId, full_path, contentType, true)
			case "directory":
				processors.ProcesssDirectory(conn, workerId, full_path, true)
			default:
				log.Println("Worker ", workerId, " Got unknkown type") // В идеале тут надо какую-нибудь 4хх отправить в ответ.
			}

		case "HEAD":

			log.Println("Worker ", workerId, "Got 'HEAD' method")

			if !utils.CheckFileExists(full_path) {
				log.Println("Worker ", workerId, " ", err)
				processors.ProcesssNotExistedContent(conn, workerId)
				break
			}

			contentType := utils.DefineContentType(full_path)
			contentType_first_type := strings.Split(contentType, "/")[0]
			log.Println("Worker ", workerId, " Content type: ", contentType)

			switch contentType_first_type {
			case "text":
				processors.ProcessText(conn, workerId, full_path, contentType, false)
			case "image":
				processors.ProcesssImage(conn, workerId, full_path, contentType, false)
			case "application":
				processors.ProcesssApplication(conn, workerId, full_path, contentType, false)
			case "directory":
				processors.ProcesssDirectory(conn, workerId, full_path, false)
			default:
				log.Println("Worker ", workerId, " Got unknkown type")
			}

		default:
			log.Println("Worker ", workerId, "Not supported")
			processors.ProcesssNotSupportedMethod(conn, workerId)
		}

		log.Println("Worker", workerId, "Processed request:", request)
		conn.Close()
	}

}

func WorkerPool(workersCount int, pool chan net.Conn, wg *sync.WaitGroup) {

	for i := 0; i < workersCount; i++ {
		wg.Add(1)
		go worker(wg, i, pool)
	}

}

func FillConnectionPool(listener net.Listener, pool chan net.Conn) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		pool <- conn
	}
}
