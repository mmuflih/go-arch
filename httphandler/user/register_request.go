package user

/**
 * Created by Muhammad Muflih Kholidin
 * at 2020-09-26 17:58:55
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 **/

type registerRequest struct {
	FullName string `json:"full_name" valid:"required"`
	Email    string `json:"email" valid:"email"`
	Pin      string `json:"pin" valid:"required"`
}

func (rr registerRequest) GetName() string {
	return rr.FullName
}

func (rr registerRequest) GetEmail() string {
	return rr.Email
}

func (rr registerRequest) GetPin() string {
	return rr.Pin
}
