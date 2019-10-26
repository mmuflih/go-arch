package lib

import (
	"strings"

	uuid "github.com/satori/go.uuid"
)

/**
 * Created by M. Muflih Kholidin
 * mmuflic@gmail.com
 * https://github.com/mmuflih
 **/

func GenerateUUID() string {
	out := uuid.NewV4()
	strOut := out.String()
	return strings.Replace(strOut, "\n", "", -1)
}
