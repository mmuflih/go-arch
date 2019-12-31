package lib

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

/*
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 1/1/20 6:31 AM
 */

func GeneratePassword(password string) string {
	bpass, err := bcrypt.GenerateFromPassword([]byte(password), 13)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(bpass)
}
