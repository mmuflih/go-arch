package user

import (
	"github.com/mmuflih/go-di-arch/app"
	"github.com/mmuflih/go-di-arch/domain/model"
	"github.com/mmuflih/go-di-arch/http/requests"
)

/**
 * Created by Muhammad Muflih Kholidin
 * at 2020-09-26 17:46:02
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 **/

func (ru handle) Register(req requests.RegisterRequest) (error, interface{}) {
	tx := ru.uRepo.DBConn().Begin()

	var errs error
	u := model.NewUser(req.FullName)
	err := ru.uRepo.Save(u, tx)
	if err != nil {
		tx.Rollback()
		return errs, nil
	}

	if req.Email != "" {
		ue := model.NewUserEmail(u.ID, req.Email)
		err = ru.ueRepo.Save(ue, tx)
		if err != nil {
			tx.Rollback()
			return errs, nil
		}
	}

	if req.Pin != "" {
		pin := app.GeneratePassword(req.Pin)
		up := model.NewUserPassword(u.ID, pin)
		err = ru.upRepo.Save(up, tx)
		if err != nil {
			tx.Rollback()
			return errs, nil
		}
	}
	tx.Commit()
	return nil, NewMeResponse(u)
}
