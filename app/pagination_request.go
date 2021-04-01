package app

import (
	"net/http"

	"github.com/mmuflih/golib/request"
)

/**
 * Created by Muhammad Muflih Kholidin
 * at 2020-09-24 00:28:54
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 **/

type PaginationRequest interface {
	GetPage() int
	GetSize() int
	GetQuery() string
}

type PaginationReq struct {
	page  int
	size  int
	query string
}

func (pr PaginationReq) GetPage() int {
	return pr.page
}

func (pr PaginationReq) GetSize() int {
	return pr.size
}

func (pr PaginationReq) GetQuery() string {
	return pr.query
}

func (pr PaginationReq) FromRequest(rr request.Reader, r *http.Request) PaginationRequest {
	page := rr.GetQueryInt(r, "page")
	if page == 0 {
		page = 1
	}
	size := rr.GetQueryInt(r, "size")
	if size == 0 {
		size = 100
	}
	return PaginationReq{
		page:  page,
		size:  size,
		query: rr.GetQuery(r, "q"),
	}
}
