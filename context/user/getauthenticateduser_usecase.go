package user

import (
	"fmt"
	"net/http"

	"github.com/mmuflih/golib/middleware"
)

type GetAuthUserUsecase interface {
	GetUserID(*http.Request) uint64
}

type getAuthUserUsecase struct {
}

func NewGetAuthUserUsecase() GetAuthUserUsecase {
	return &getAuthUserUsecase{}
}

func (this getAuthUserUsecase) GetUserID(r *http.Request) uint64 {
	userID, err := middleware.ExtractClaim(r, "user_id")
	if err != nil {
		fmt.Println("<-> +> Get user id from token", err)
		return 0
	}
	val, ok := userID.(float64)
	if ok {
		return uint64(val)
	}
	xType := fmt.Sprintf("%T", userID)
	fmt.Println("error casting", "Value is ", xType)
	return 0
}
