package request

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

/**
 * Created by Muhammad Muflih Kholidin
 * at 2021-04-23 08:39:24
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 **/

type reader struct{}

type Reader interface {
	GetRouteParam(c *gin.Context, name string) string
	GetRouteParamInt(c *gin.Context, name string) int
	GetJsonData(r *http.Request, data interface{}) (err error)
	GetQuery(r *http.Request, query string) string
	GetQueryInt(r *http.Request, query string) int
}

func NewRequestReader() Reader {
	return &reader{}
}

func (rr *reader) GetRouteParam(c *gin.Context, name string) string {
	return c.Param(name)
}

func (rr *reader) GetRouteParamInt(c *gin.Context, name string) int {
	param := rr.GetRouteParam(c, name)
	i, err := strconv.Atoi(param)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return i
}

func (rr *reader) GetJsonData(r *http.Request, data interface{}) (err error) {
	err = json.NewDecoder(r.Body).Decode(data)
	return
}

func (rr *reader) GetQuery(r *http.Request, key string) string {
	qs := r.URL.Query()
	return qs.Get(key)
}

func (rr *reader) GetQueryInt(r *http.Request, key string) int {
	q := rr.GetQuery(r, key)
	qi, err := strconv.Atoi(q)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return qi
}
