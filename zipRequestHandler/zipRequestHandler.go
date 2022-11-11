package zipRequestHandler

import (
	"archive/zip"
	"encoding/binary"
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
	timebytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(timebytes, uint64(time))

	header := zip.FileHeader{
		Name:  "Message" + fmt.Sprintf("%d", messageNum),
		Extra: timebytes,
	}
	zipFile, err := reqHandler.zipWriter.CreateHeader(&header)
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
