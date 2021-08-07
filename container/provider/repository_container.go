package provider

import (
	"github.com/mmuflih/go-arch/domain/repository/mysql"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-02-06 19:47
**/

func Repositories() []interface{} {
	var r []interface{}

	r = append(r, mysql.NewUserRepo)
	r = append(r, mysql.NewUserPasswordRepo)
	r = append(r, mysql.NewUserEmailRepo)

	return r
}
