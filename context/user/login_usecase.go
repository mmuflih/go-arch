package user

import (
	"errors"
	"fmt"

	"github.com/mmuflih/go-di-arch/app"
	"github.com/mmuflih/go-di-arch/domain/model"
	"github.com/mmuflih/go-di-arch/domain/repository/mysql"
	"golang.org/x/crypto/bcrypt"
)

/**
 * Created by Muhammad Muflih Kholidin
 * at 2020-09-27 16:28:59
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 **/

type LoginUsecase interface {
	Handle(req LoginRequest) (error, interface{})
}

type LoginRequest interface {
	GetEmail() string
	GetPin() string
}

type loginUsecase struct {
	uRepo      mysql.UserRepository
	uEmailRepo mysql.UserEmailRepository
	uPassRepo  mysql.UserPasswordRepository
	tokenUC    GenerateUsecase
}

func NewLoginUsecase(uEmailRepo mysql.UserEmailRepository,
	uPassRepo mysql.UserPasswordRepository,
	uRepo mysql.UserRepository,
	tokenUC GenerateUsecase,
) LoginUsecase {
	return &loginUsecase{uRepo, uEmailRepo, uPassRepo, tokenUC}
}

func (lu *loginUsecase) Handle(req LoginRequest) (error, interface{}) {
	var userID uint64
	/** check email */
	err, email := lu.validateEmail(req.GetEmail())
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

	err = lu.validatePassword(userID, req.GetPin())
	if err != nil {
		/** email atau phone number and pass not match */
		return err, nil
	}
	_, usr := lu.uRepo.Find(userID)
	_ = lu.uRepo.SetLastLogin(usr)

	return nil, lu.tokenUC.ClaimToken(usr)
}

func (lu *loginUsecase) validateEmail(email string) (error, *model.UserEmail) {
	err, ue := lu.uEmailRepo.GetByEmail(email)
	if ue != nil && ue.Email != "" && !ue.Active {
		return errors.New("Please verify email before"), ue
	}
	if err != nil {
		return err, nil
	}
	return nil, ue
}

func (lu *loginUsecase) validatePassword(userID uint64, pass string) error {
	err, upass := lu.uPassRepo.GetActiveByUser(userID)
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
