package user

import (
	"errors"
	"fmt"

	"github.com/mmuflih/go-di-arch/app"
	"github.com/mmuflih/go-di-arch/domain/model"
	"github.com/mmuflih/go-di-arch/http/requests"
	"golang.org/x/crypto/bcrypt"
)

/**
 * Created by Muhammad Muflih Kholidin
 * at 2020-09-27 16:28:59
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 **/

func (lu handle) Login(req requests.LoginRequest) (error, interface{}) {
	var userID uint64
	/** check email */
	err, email := lu.validateEmail(req.Email)
	if err != nil && email != nil {
		return err, nil
	}

	/** phone and email not found */
	if err != nil {
		return errors.New("Check email and password"), nil
	}

	if email != nil {
		/** email found */
		userID = email.UserID
	}

	err = lu.validatePassword(userID, req.Pin)
	if err != nil {
		/** email atau phone number and pass not match */
		return err, nil
	}
	_, usr := lu.uRepo.Find(userID)
	_ = lu.uRepo.SetLastLogin(usr)

	return nil, lu.ClaimToken(usr)
}

func (lu handle) validateEmail(email string) (error, *model.UserEmail) {
	err, ue := lu.ueRepo.GetByEmail(email)
	if ue != nil && ue.Email != "" && !ue.Active {
		return errors.New("Please verify email before"), ue
	}
	if err != nil {
		return err, nil
	}
	return nil, ue
}

func (lu handle) validatePassword(userID uint64, pass string) error {
	err, upass := lu.upRepo.GetActiveByUser(userID)
	if err != nil {
		return errors.New("Check email and password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(upass.Password), []byte(pass))
	if err != nil {
		fmt.Println(err)
		return app.NewError("Check email and password")
	}
	return nil
}
