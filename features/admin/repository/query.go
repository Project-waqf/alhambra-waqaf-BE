package repository

import (
	"context"
	"errors"
	"time"
	"wakaf/features/admin/domain"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type AdminRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

func New(db *gorm.DB, redis *redis.Client) domain.RepoInterface {
	return &AdminRepository{
		db:    db,
		redis: redis,
	}
}

func (repo *AdminRepository) Login(data domain.Admin) (domain.Admin, error) {
	var input Admin

	if err := repo.db.Model(&Admin{}).Where("email = ?", data.Email).First(&input).Error; err != nil {
		return domain.Admin{}, err
	}

	return ToDomainLogin(input), nil
}

func (repo *AdminRepository) Register(data domain.Admin) error {
	input := FromDomainRegister(data)

	if err := repo.db.Create(&input).Error; err != nil {
		return err
	}
	return nil
}

func (repo *AdminRepository) GetUser(data domain.Admin) error {
	var res Admin

	if err := repo.db.Where("email", data.Email).First(&res).Error; err == nil {
		return errors.New("email has taken")
	}
	return nil
}

func (repo *AdminRepository) UpdatePassword(data domain.Admin) error {
	if res := repo.db.Model(Admin{}).Where("id = ?", data.ID).Update("password", data.Password).RowsAffected; res == 0 {
		return errors.New("not row affected")
	}
	return nil
}

func (repo *AdminRepository) GetFromRedis(token string) (string, error) {

	res, err := repo.redis.Get(context.Background(), token).Result()
	if err != nil {
		return "", err
	}
	return res, nil
}

func (repo *AdminRepository) SaveRedis(email, token string) error {

	err := repo.redis.Set(context.Background(), token, email, time.Duration(1*time.Hour)).Err()
	if err != nil {
		return err
	}
	return nil
}

func (repo *AdminRepository) UpdatePasswordByEmail(input domain.Admin) error {
	if res := repo.db.Model(Admin{}).Where("email = ?", input.Email).Update("password", input.Password).RowsAffected; res == 0 {
		return errors.New("not row affected")
	}
	return nil
}

func (repo *AdminRepository) DeleteToken(token string) error {

	if err := repo.redis.Del(context.Background(), token).Err(); err != nil {
		return err
	}
	return nil
}

func (repo *AdminRepository) UpdateProfile(data domain.Admin) (domain.Admin, error) {
	if err := repo.db.Model(Admin{}).Where("id = ?", data.ID).Updates(map[string]interface{}{"name": data.Name, "email": data.Email, "password": data.Password}).Error; err != nil {
		return domain.Admin{}, err
	}
	return data, nil
}

func (repo *AdminRepository) GetUserById(id uint) (domain.Admin, error) {
	var res Admin

	if err := repo.db.Where("id = ?", id).First(&res).Error; err != nil {
		return domain.Admin{}, err
	}
	return ToDomainLogin(res), nil
}

func (repo *AdminRepository) UpdateImage(input domain.Admin) error {

	if err := repo.db.Model(Admin{}).Where("id = ?", input.ID).Updates(map[string]interface{}{
		"image":   input.Image,
		"file_id": input.FileId,
	}).Error; err != nil {
		return err
	}
	return nil
}
