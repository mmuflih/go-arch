package mysql

import (
	"github.com/mmuflih/go-arch/config"
	"github.com/mmuflih/go-arch/domain/model"
	"gorm.io/gorm"
)

/**
 * Created by Muhammad Muflih Kholidin
 * at 2020-09-26 18:14:57
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 **/

type UserEmailRepository interface {
	Save(ue *model.UserEmail, tx *gorm.DB) error
	Update(ue *model.UserEmail, tx *gorm.DB) error
	GetByEmail(email string) (error, *model.UserEmail)
	GetByUser(userID uint64) (error, []*model.UserEmail)
}

type userEmailRepo struct {
	db *gorm.DB
}

func NewUserEmailRepo(myConn *config.MySQL) UserEmailRepository {
	return &userEmailRepo{
		myConn.MainDB,
	}
}

func (uer *userEmailRepo) Save(ue *model.UserEmail, tx *gorm.DB) error {
	return tx.Save(ue).Error
}

func (uer *userEmailRepo) Update(ue *model.UserEmail, tx *gorm.DB) error {
	return tx.Model(&model.UserEmail{}).Updates(ue).Error
}

func (uer *userEmailRepo) GetByEmail(email string) (error, *model.UserEmail) {
	ue := new(model.UserEmail)
	err := uer.db.
		Where("email = ?", email).
		First(&ue).
		Error
	return err, ue
}

func (uer *userEmailRepo) GetByUser(userID uint64) (error, []*model.UserEmail) {
	var ues []*model.UserEmail
	err := uer.db.Find(&ues).Where("user_id", userID).Error
	return err, ues
}
