package request

import (
	"errors"
	"net/http"

	"github.com/asaskevich/govalidator"
)

/**
 * Created by Muhammad Muflih Kholidin
 * at 2020-09-27 10:04:46
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 **/

func RequestValidator(r *http.Request, rr Reader, req interface{}) error {
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
