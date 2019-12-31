package lib

import (
	"strings"

	uuid "github.com/satori/go.uuid"
)

/*
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 1/1/20 6:31 AM
 */

func GenerateUUID() string {
	out := uuid.NewV4()
	strOut := out.String()
	return strings.Replace(strOut, "\n", "", -1)
}
