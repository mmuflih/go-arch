package request

import (
	"errors"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/mmuflih/golib/request"
)

/**
 * Created by Muhammad Muflih Kholidin
 * at 2020-09-27 10:04:46
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 **/

func RequestValidator(r *http.Request, rr request.Reader, req interface{}) error {
	err := rr.GetJsonData(r, &req)
	if err != nil {
		return err
	}
	result, err := govalidator.ValidateStruct(req)
	if err != nil {
		return err
	}
	if !result {
		return errors.New("validate error")
	}
	return nil
}
