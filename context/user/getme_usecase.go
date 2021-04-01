package user

/**
 * Created by Muhammad Muflih Kholidin
 * at 2020-09-29 01:54:09
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 **/

func (gm read) Me(id uint64) (error, interface{}) {
	err, u := gm.uRepo.Find(id)
	if err != nil {
		return err, nil
	}

	return nil, NewMeResponse(u)
}
