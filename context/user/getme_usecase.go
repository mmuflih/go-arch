package user

import (
	"github.com/mmuflih/go-di-arch/domain/repository/mysql"
)

/**
 * Created by Muhammad Muflih Kholidin
 * at 2020-09-29 01:54:09
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 **/

type GetMeUsecase interface {
	Me(id uint64) (error, interface{})
}

type getMeUc struct {
	uRepo mysql.UserRepository
}

func NewGetMeUsecase(
	uRepo mysql.UserRepository,
) GetMeUsecase {
	return &getMeUc{uRepo}
}

func (gm *getMeUc) Me(id uint64) (error, interface{}) {
	err, u := gm.uRepo.Find(id)
	if err != nil {
		return err, nil
	}

	return nil, NewMeResponse(u)
}
