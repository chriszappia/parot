package zipRequestHandler

import (
	"archive/zip"
	"fmt"
	"os"
)

type ZipReqHandler struct {
	file      *os.File
	zipWriter *zip.Writer
}

func NewZipRequestHandler(filename string) ZipReqHandler {

	archive, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	zipWriter := zip.NewWriter(archive)

	zrh := ZipReqHandler{
		file:      archive,
		zipWriter: zipWriter,
	}
	return zrh
}

func (reqHandler ZipReqHandler) HandleRequest(messageNum int, time int64, req []byte) {
	zipFile, err := reqHandler.zipWriter.Create("Message" + fmt.Sprintf("%d", messageNum))
	if err != nil {
		fmt.Println(err)
	}
	_, err = zipFile.Write(req)
	if err != nil {
		fmt.Println(err)
	}
}

func (reqHandler ZipReqHandler) Close() {
	reqHandler.zipWriter.Close()
	reqHandler.file.Close()
}
