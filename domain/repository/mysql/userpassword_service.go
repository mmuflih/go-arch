package mysql

import (
	"github.com/mmuflih/go-di-arch/app"
	"github.com/mmuflih/go-di-arch/config"
	"github.com/mmuflih/go-di-arch/domain/model"
	"gorm.io/gorm"
)

/**
 * Created by Muhammad Muflih Kholidin
 * at 2020-09-26 18:15:14
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 **/

type UserPasswordRepository interface {
	Save(ue *model.UserPassword, tx *gorm.DB) error
	Update(ue *model.UserPassword, tx *gorm.DB) error
	GetActiveByUser(userID uint64) (error, *model.UserPassword)
	GetByUser(userID uint64) (error, []*model.UserPassword)
	GetResetPass(userID uint64) (error, *model.UserPassword)
	FindByResetToken(tokenID string) (error, *model.UserPassword)
	ResetPassword(tokenID, password string, tx *gorm.DB) error
	DisableOld(userID uint64, tx *gorm.DB) error
	DBConn() *gorm.DB
}

type userPasswordRepo struct {
	db *gorm.DB
}

func NewUserPasswordRepo(myConn *config.MySQL) UserPasswordRepository {
	return &userPasswordRepo{
		myConn.MainDB,
	}
}

func (upr *userPasswordRepo) DBConn() *gorm.DB {
	return upr.db
}

func (upr *userPasswordRepo) Save(ue *model.UserPassword, tx *gorm.DB) error {
	return tx.Save(ue).Error
}

func (upr *userPasswordRepo) Update(ue *model.UserPassword, tx *gorm.DB) error {
	return tx.Model(&model.UserPassword{}).Updates(ue).Error
}

func (upr *userPasswordRepo) GetActiveByUser(userID uint64) (error, *model.UserPassword) {
	up := new(model.UserPassword)
	err := upr.db.
		Where("user_id = ? and active = ?", userID, true).
		First(&up).
		Error
	return err, up
}

func (upr *userPasswordRepo) GetByUser(userID uint64) (error, []*model.UserPassword) {
	var ues []*model.UserPassword
	err := upr.db.Find(&ues).Where("user_id", userID).Error
	return err, ues
}

func (upr *userPasswordRepo) GetResetPass(userID uint64) (error, *model.UserPassword) {
	up := new(model.UserPassword)
	err := upr.db.
		Where("user_id = ? and active = false", userID).
		First(&up).
		Error
	return err, up
}

func (upr *userPasswordRepo) FindByResetToken(token string) (error, *model.UserPassword) {
	up := new(model.UserPassword)
	err := upr.db.
		Where("reset_token = ? ", token).
		First(&up).
		Error
	return err, up
}

func (upr *userPasswordRepo) ResetPassword(token, newPassword string, tx *gorm.DB) error {
	return tx.Where("reset_token = ?", token).
		Updates(model.UserPassword{
			Password: app.GeneratePassword(newPassword), Active: true,
		}).Error
}

func (upr *userPasswordRepo) DisableOld(userID uint64, tx *gorm.DB) error {
	return tx.Model(&model.UserPassword{}).Where("user_id = ?", userID).Update("active", false).Error
}
