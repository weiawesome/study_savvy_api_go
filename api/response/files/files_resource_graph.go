package files

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

const PureText = "PureText"

type GraphFile struct {
	FilePath string
}

func (f *GraphFile) Exist() error {
	_, err := os.Stat(f.FilePath)
	if os.IsNotExist(err) {
		return err
	}
	return nil
}
func (f *GraphFile) CanOpenAndSent(c *gin.Context) error {
	file, err := os.Open(f.FilePath)
	if err != nil {
		return err
	}
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	defer func() {
		err = file.Close()
		if err != nil {
			return
		}
	}()

	if err != nil {
		return err
	}

	c.Header("Content-Disposition", "attachment; filename="+f.FilePath)
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Length", fmt.Sprintf("%d", fileInfo.Size()))
	_, err = io.Copy(c.Writer, file)

	return err
}
func (f *GraphFile) IsPureText() bool {
	return f.FilePath == PureText
}
