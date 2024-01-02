package gin_test_functions

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/textproto"
	"net/url"

	"github.com/gin-gonic/gin"
)

func BuildGinTestEngine(w *httptest.ResponseRecorder) (*gin.Engine, *gin.Context) {

	gin.SetMode(gin.TestMode)
	e := gin.Default()

	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}

	return e, ctx
}

func MockJsonGet(c *gin.Context, params gin.Params, u url.Values) {
	c.Request.Method = "GET"
	c.Request.Header.Set("Content-Type", "application/json")

	c.Params = params

	c.Request.URL.RawQuery = u.Encode()
}

func CreateFileHeaderFromBytes(fileContent []byte, contentType, fileName, fieldName string) (*multipart.FileHeader, error) {
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)

	mimeHeader := make(textproto.MIMEHeader)
	mimeHeader.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, fieldName, fileName))
	mimeHeader.Set("Content-Type", contentType)

	filePart, err := writer.CreatePart(mimeHeader)
	if err != nil {
		return nil, err
	}

	if _, err = io.Copy(filePart, bytes.NewReader(fileContent)); err != nil {
		return nil, err
	}

	if err = writer.Close(); err != nil {
		return nil, err
	}

	reader := multipart.NewReader(&buffer, writer.Boundary())
	form, err := reader.ReadForm(1 << 20)
	if err != nil {
		return nil, err
	}
	if fileHeader := form.File[fieldName]; len(fileHeader) > 0 {
		return fileHeader[0], nil
	}

	return nil, errors.New("failed")

}
