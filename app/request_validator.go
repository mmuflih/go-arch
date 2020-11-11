package app

import (
	"errors"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/mmuflih/go-httplib/httplib"
)

/**
 * Created by Muhammad Muflih Kholidin
 * at 2020-09-27 10:04:46
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 **/

func RequestValidator(r *http.Request, rr httplib.RequestReader, req interface{}) error {
	err := rr.GetJsonData(r, &req)
	if err != nil {
		return err
	}
	result, err := govalidator.ValidateStruct(req)
	if err != nil {
		return err
	}
	if !result {
		return errors.New("Validate error")
	}
	return nil
}
