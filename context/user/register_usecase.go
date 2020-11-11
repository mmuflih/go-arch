package user

import (
	"github.com/mmuflih/go-di-arch/app"
	"github.com/mmuflih/go-di-arch/domain/model"
	"github.com/mmuflih/go-di-arch/domain/repository/mysql"
)

/**
 * Created by Muhammad Muflih Kholidin
 * at 2020-09-26 17:46:02
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 **/

type RegisterUsecase interface {
	Handle(RegisterRequest) (error, interface{})
}

type RegisterRequest interface {
	GetName() string
	GetPin() string
	GetEmail() string
}

type registerUsecase struct {
	userRepo         mysql.UserRepository
	userEmailRepo    mysql.UserEmailRepository
	userPasswordRepo mysql.UserPasswordRepository
}

func NewRegisterUsecase(userRepo mysql.UserRepository,
	userEmailRepo mysql.UserEmailRepository,
	userPasswordRepo mysql.UserPasswordRepository,
) RegisterUsecase {
	return &registerUsecase{userRepo, userEmailRepo, userPasswordRepo}
}

func (ru *registerUsecase) Handle(req RegisterRequest) (error, interface{}) {
	tx := ru.userRepo.DBConn().Begin()

	u := model.NewUser(req.GetName())
	err := ru.userRepo.Save(u, tx)
	if err != nil {
		tx.Rollback()
		return err, nil
	}

	if req.GetEmail() != "" {
		ue := model.NewUserEmail(u.ID, req.GetEmail())
		err = ru.userEmailRepo.Save(ue, tx)
		if err != nil {
			tx.Rollback()
			return err, nil
		}
	}

	if req.GetPin() != "" {
		pin := app.GeneratePassword(req.GetPin())
		up := model.NewUserPassword(u.ID, pin)
		err = ru.userPasswordRepo.Save(up, tx)
		if err != nil {
			tx.Rollback()
			return err, nil
		}
	}
	tx.Commit()
	return nil, NewMeResponse(u)
}
