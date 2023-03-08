package services

import (
	"context"
	"crypto/rand"
	"io"
	"time"
	"wakaf/config"
	"wakaf/features/admin/domain"
	"wakaf/middlewares"
	"wakaf/pkg/helper"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type AdminServices struct {
	AdminRepository domain.RepoInterface
}

func New(data domain.RepoInterface) domain.UseCaseInterface {
	return &AdminServices{
		AdminRepository: data,
	}
}

var (
	logger = helper.Logger()
)

func (service *AdminServices) Login(input domain.Admin) (domain.Admin, error) {
	res, err := service.AdminRepository.Login(input)
	if err != nil {
		return domain.Admin{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(config.Getconfig().SALT1+input.Password+config.Getconfig().SALT2)); err != nil {
		return domain.Admin{}, err
	}

	token, err := middlewares.CreateToken(int(res.ID), res.Email)
	if err != nil {
		return domain.Admin{}, err
	}
	res.Token = token

	return res, nil
}

func (u *AdminServices) Register(input domain.Admin) error {

	if err := u.AdminRepository.GetUser(input); err != nil {
		return err
	}

	saltPw := config.Getconfig().SALT1 + input.Password + config.Getconfig().SALT2
	hash, err := bcrypt.GenerateFromPassword([]byte(saltPw), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	input.Password = string(hash)
	err = u.AdminRepository.Register(input)
	if err != nil {
		return err
	}

	return nil
}

func (u *AdminServices) UpdatePassword(input domain.Admin) error {

	saltPw := config.Getconfig().SALT1 + input.Password + config.Getconfig().SALT2
	hash, err := bcrypt.GenerateFromPassword([]byte(saltPw), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	input.Password = string(hash)
	if err := u.AdminRepository.UpdatePassword(input); err != nil {
		return err
	}
	return nil
}

func (u *AdminServices) ForgotSendEmail(input domain.Admin) (domain.Admin, error) {
	res, err := u.AdminRepository.Login(input)
	if err != nil {
		return domain.Admin{}, err
	}

	// Generate OTP
	otp := encodeToString(6)

	// Save To Redis
	redis := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "alhambra",
		DB:       0,
	})

	if err := saveToRedis(redis, input.Email, otp); err != nil {
		logger.Error("Failed to save data redis",  zap.Error(err))
		return domain.Admin{}, err
	}

	return res, nil
}

func encodeToString(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func saveToRedis(c *redis.Client, email, otp string) error {
	logger.Info("Redis Connection", zap.Any("PING", c.Ping(context.Background())))
	cmd := c.Set(context.Background(), email, otp, time.Duration(5)*time.Minute)
	if cmd.Err() != nil {
		return cmd.Err()
	}
	return nil
}
