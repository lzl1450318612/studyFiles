package gin

import (
	"fmt"
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

type Ctx struct {
	Req        *http.Request
	Writer     http.ResponseWriter
	Path       string
	Method     string
	StatusCode int
}

func NewCtx(writer http.ResponseWriter, req *http.Request) Ctx {
	return Ctx{
		Req:    req,
		Writer: writer,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

//------------------- get from req----------------------

func (c *Ctx) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c Ctx) PostForm(key string) {
	c.Req.FormValue(key)
}

//------------------- set response----------------------

func (c *Ctx) SetHeader(key string, val string) {
	c.Writer.Header().Set(key, val)
}

func (c *Ctx) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Ctx) Json(statusCode int, obj interface{}) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(statusCode)
	encoder := jsoniter.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
		c.Status(500)
		return
	}
}

func (c *Ctx) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}

func (c *Ctx) Data(code int, data string) {
	c.Status(code)
	c.Writer.Write([]byte(data))
}

func (c *Ctx) String(code int, format string, values ...interface{}) {
	c.Status(code)
	c.SetHeader("Content-Type", "text/plain")
	c.Writer.Write([]byte(fmt.Sprintf(format, values)))
}
