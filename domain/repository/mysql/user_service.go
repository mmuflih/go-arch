package mysql

import (
	"time"

	"github.com/mmuflih/go-di-arch/config"
	"github.com/mmuflih/go-di-arch/domain/model"
	paginator "github.com/mmuflih/gorm-paginator"
	"gorm.io/gorm"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-03-09 21:15
**/

type UserRepository interface {
	DBConn() *gorm.DB
	Save(u *model.User, tx *gorm.DB) error
	Update(u *model.User, tx *gorm.DB) error
	SetLastLogin(u *model.User) error
	Find(id uint64) (error, *model.User)
	FindBy(page, size int) *paginator.Paginator
}

type userService struct {
	db *gorm.DB
}

func (us userService) DBConn() *gorm.DB {
	return us.db
}

func (us userService) Save(u *model.User, tx *gorm.DB) error {
	return tx.Save(u).Error
}

func (us userService) SetLastLogin(u *model.User) error {
	u.LastLogin = time.Now()
	return us.db.
		Updates(u).
		Error
}

func (us userService) Update(u *model.User, tx *gorm.DB) error {
	return tx.Model(&model.User{}).Updates(u).
		Where("id", u.ID).Error
}

func (us userService) Find(id uint64) (error, *model.User) {
	u := new(model.User)
	err := us.db.First(&u, id).Error
	return err, u
}

func (us userService) FindBy(page, size int) *paginator.Paginator {
	var users []model.User
	paginatior := paginator.Make(&paginator.Config{
		DB:      us.db,
		Page:    page,
		Size:    size,
		OrderBy: []string{"id desc"},
		ShowSQL: true,
	}, &users)
	return paginatior
}

func NewUserRepo(myConn *config.MyConn) UserRepository {
	return &userService{myConn.Conn1}
}
