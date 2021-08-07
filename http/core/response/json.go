package response

import (
	"log"
	"math"
	"net/http"
	"reflect"
	"runtime"
	"strconv"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
)

/**
 * Created by Muhammad Muflih Kholidin
 * at 2021-04-23 08:12:19
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 **/

type SuccessData struct {
	Data interface{} `json:"data"`
	Code int         `json:"code"`
}

type ErrorData struct {
	DeveloperMessage string `json:"developer_message"`
	ErrorCode        int    `json:"error_code"`
	MoreInfo         string `json:"more_info"`
	Status           int    `json:"status"`
	UserMessage      string `json:"user_message"`
}

type PaginateData struct {
	Data       interface{} `json:"data"`
	Additional interface{} `json:"additional,omitempty"`
	Paginate   struct {
		Count     int `json:"total"`
		Page      int `json:"page"`
		Size      int `json:"size"`
		PageCount int `json:"page_count"`
	} `json:"paginate"`
	Code int `json:"code"`
}

func NewPaginate(data interface{}, count, page, size int) *PaginateData {
	dp := PaginateData{
		Data: data,
		Paginate: struct {
			Count     int `json:"total"`
			Page      int `json:"page"`
			Size      int `json:"size"`
			PageCount int `json:"page_count"`
		}{
			count, page, size, int(math.Ceil(float64(count) / float64(size))),
		},
		Code: 0,
	}
	dp.Code = http.StatusOK
	return &dp
}

func Exception(c *gin.Context, err error, code int) {
	/** sentry */
	go sendSentry(err)

	pc, fn, line, _ := runtime.Caller(1)
	log.Printf("[error] %s:%d %v on %s", fn, line, err, pc)
	exception := ErrorData{
		err.Error() + " on " + fn + ":" + strconv.Itoa(line),
		code,
		"Contact developer or administrator",
		code,
		err.Error(),
	}
	c.JSON(code, exception)
}

func Success(c *gin.Context, data interface{}) {
	exception := SuccessData{
		data,
		http.StatusOK,
	}
	c.JSON(http.StatusOK, exception)
}

func Json(c *gin.Context, resp interface{}, err error) {
	if err != nil {
		Exception(c, err, 422)
		return
	}
	r := parseStruct(resp)
	if r == "*Paginator" {
		Paginate(c, resp)
		return
	}
	if r == "*PaginateData" {
		Paginate(c, resp)
		return
	}
	if r == "*PaginatorResponse" {
		Paginate(c, resp)
		return
	}
	Success(c, resp)
}

func Paginate(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

/** local func */
func sendSentry(err error) {
	defer sentry.Flush(time.Second * 2)
	sentry.CaptureException(err)
}

func parseStruct(myvar interface{}) string {
	if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	} else {
		return t.Name()
	}
}
