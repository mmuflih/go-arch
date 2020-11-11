package user

/**
 * Created by Muhammad Muflih Kholidin
 * at 2020-09-26 17:58:55
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 **/

type loginRequest struct {
	Email string `json:"email" valid:"required"`
	Pin   string `json:"pin" valid:"required"`
}

func (lr loginRequest) GetEmail() string {
	return lr.Email
}

func (lr loginRequest) GetPin() string {
	return lr.Pin
}
