package repositories

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"waizly.com/adriantoprasetyo/ticket_app/models"
	"waizly.com/adriantoprasetyo/ticket_app/repositories"
	"waizly.com/adriantoprasetyo/ticket_app/utility"
)

type psqlUserRepo struct {
	Conn *gorm.DB
}

func NewPSQLUser(conn *gorm.DB) repositories.User {
	return &psqlUserRepo{
		Conn: conn,
	}
}

func (p *psqlUserRepo) Register(m models.User) (models.User, error) {

	err := p.Conn.Create(&m).Error
	if err != nil {
		return models.User{}, err
	}
	return m, nil
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (p *psqlUserRepo) LoginCheck(username string, password string) (string, error) {

	var (
		err error
		m   models.User
	)

	err = p.Conn.Model(models.User{}).Where("username = ?", username).Take(&m).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, m.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := utility.GenerateToken(uint(m.Id))

	if err != nil {
		return "", err
	}

	return token, nil

}

func (p *psqlUserRepo) GetByUsername(username string) (int64, error) {
	var (
		ttl int64
		err error
	)

	err = p.Conn.Model(&models.User{}).Where("username=?", username).Count(&ttl).Error
	if err != nil {
		return 0, err
	}

	return ttl, nil
}

func (p *psqlUserRepo) Get(id int) (models.User, error) {
	var (
		user models.User
		err  error
	)
	err = p.Conn.Where("id=?", id).Find(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (p *psqlUserRepo) Delete(id int) error {
	res := p.Conn.Where("id=?", id).Delete(&models.User{})
	if res.Error != nil {
		return res.Error
	}

	return nil
}
