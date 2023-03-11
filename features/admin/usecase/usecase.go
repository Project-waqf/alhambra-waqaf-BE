package services

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"time"
	"wakaf/config"
	"wakaf/features/admin/domain"
	"wakaf/middlewares"
	"wakaf/pkg/helper"
	"wakaf/utils/email"

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

	// Save To Redis
	redis := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: "",
		DB:       0,
	})

	token := encrypt([]byte("ini adalah kunci ya gest"), input.Email)

	if err := saveToRedis(redis, input.Email, token); err != nil {
		logger.Error("Failed to save data redis",  zap.Error(err))
		return domain.Admin{}, err
	}

	if err := email.SendOtpGmail(input.Email, token); err != nil {
		logger.Error("Failed to send email", zap.Error(err))
		return domain.Admin{}, err
	}

	return res, nil
}

func encrypt(key []byte, email string) string {
	// key := []byte(keyText)
	plaintext := []byte(email)

	block, err := aes.NewCipher(key)
	if err != nil {
		logger.Error("error new chipper", zap.Error(err))
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		logger.Error("error chippertext", zap.Error(err))
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// convert to base64
	return base64.URLEncoding.EncodeToString(ciphertext)
}

func decrypt(key []byte, cryptoText string) string {
	ciphertext, _ := base64.URLEncoding.DecodeString(cryptoText)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)

	return fmt.Sprintf("%s", ciphertext)
}

func saveToRedis(c *redis.Client, email, token string) error {
	logger.Info("Redis Connection", zap.Any("PING", c.Ping(context.Background())))
	cmd := c.Set(context.Background(), token, email, time.Duration(5)*time.Minute)
	if cmd.Err() != nil {
		return cmd.Err()
	}
	return nil
}
