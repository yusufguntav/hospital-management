package user

import (
	"context"
	"errors"
	"math/rand"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/yusufguntav/hospital-management/pkg/cache"
	"github.com/yusufguntav/hospital-management/pkg/dtos"
	"github.com/yusufguntav/hospital-management/pkg/entities"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Register(c context.Context, req dtos.DTOUserRegister) error
	Login(c context.Context, req dtos.DTOUserLogin) (string, error)
	ResetPasswordApprove(c context.Context, phoneNumber string, areaCode string) (int, error)
	ResetPassword(c context.Context, req dtos.DTOResetPassword) error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) Login(c context.Context, req dtos.DTOUserLogin) (string, error) {
	var user entities.User
	if err := ur.db.WithContext(c).Where("email = ? OR CONCAT(area_code,phone) = ?", req.MailOrPhone, req.MailOrPhone).First(&user).Error; err != nil {
		return "", errors.New("mail or password invalid")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return "", errors.New("mail or password invalid")
	}

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   user.Base.UUID.String(),
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
		"role": user.Role,
	})

	token, err := generateToken.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		return "", err
	}

	return token, nil
}
func (ur *UserRepository) Register(c context.Context, req dtos.DTOUserRegister) error {
	return nil
}

func (ur *UserRepository) ResetPasswordApprove(c context.Context, phoneNumber string, areaCode string) (int, error) {

	//Check if phone number is valid
	var count int64
	ur.db.WithContext(c).Model(entities.User{}).Where("phone = ? AND area_code = ?", phoneNumber, areaCode).Count(&count)
	if count == 0 {
		return 0, errors.New("phone number not found")
	}

	// Check if verification code already sent
	if cache.IsExist(c, phoneNumber) {
		return 0, errors.New("verification code already sent")
	}
	rand.Seed(time.Now().UnixNano())
	verificationCode := rand.Intn(9000) + 1000
	if err := cache.Set(c, areaCode+phoneNumber, verificationCode, 60); err != nil {
		return 0, err
	}
	return verificationCode, nil
}

func (ur *UserRepository) ResetPassword(c context.Context, req dtos.DTOResetPassword) error {
	if !cache.IsExist(c, req.AreaCode+req.PhoneNumber) {
		return errors.New("verification code invalid or expired")
	}
	var verificationCode int
	cache.Get(c, req.AreaCode+req.PhoneNumber, &verificationCode)

	if verificationCode != req.Code {
		return errors.New("verification code invalid or expired")
	}

	// Update password

	// Password hashing
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	ur.db.WithContext(c).Model(entities.User{}).Where("phone = ? AND area_code = ?", req.PhoneNumber, req.AreaCode).Update("password", string(passwordHash))

	return nil
}
