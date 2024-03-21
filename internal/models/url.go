package models

import (
	"RestApiGolang/internal/database/sqlite"
	log "RestApiGolang/internal/logger"
	"errors"
	"github.com/jinzhu/gorm"
	"time"
)

var (
	db                        = &sqlite.Db
	ErrAliasNotSpecified      = errors.New("alias for URL not specified")
	ErrUrlNotSpecified        = errors.New("URL not specified")
	ErrUrlWithAliasOrIdExists = errors.New("URL with specified alias or id already exists")
)

type Url struct {
	Id        int64     `json:"id"`
	Alias     string    `json:"alias"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *Url) Validate() error {
	switch {
	case u.Alias == "":
		return ErrAliasNotSpecified
	case u.Url == "":
		return ErrUrlNotSpecified
	}

	return nil
}

func SaveURL(u *Url) error {
	const op = "internal.models.SaveURL"
	if err := u.Validate(); err != nil {
		return err
	}

	var existingUrl Url
	if err := (*db).Where("alias=? OR id=?", u.Alias, u.Id).First(&existingUrl).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Errorf("%s: %s", op, err.Error())
			return err
		}
		err = (*db).Save(u).Error
		if err != nil {
			log.Errorf("%s: %s", op, err.Error())
			return err
		}
	} else {
		err = ErrUrlWithAliasOrIdExists
		log.Errorf("%s: %s", op, err.Error())
		return err
	}

	return nil
}

func GetURL(alias string) (string, error) {
	const op = "internal.models.GetURL"

	u := Url{}
	err := (*db).Where("alias=?", alias).Find(&u).Error
	if err != nil {
		log.Errorf("%s: %s", op, err.Error())
		return "", err
	}

	return u.Url, nil
}

func GetURLs() ([]Url, error) {
	const op = "internal.models.GetURLs"

	var us []Url
	err := (*db).Table("urls").Find(&us).Error
	if err != nil {
		log.Errorf("%s: %s", op, err.Error())
		return us, err
	}

	return us, nil
}

func DeleteURL(alias string) error {
	const op = "internal.models.DeleteURL"

	u := Url{}
	err := (*db).Where("alias=?", alias).Find(&u).Error
	if gorm.IsRecordNotFoundError(err) {
		log.Errorf("%s: %s", op, err.Error())
		return err
	}
	err = (*db).Delete(u).Error
	if err != nil {
		log.Errorf("%s: %s", op, err.Error())
		return err
	}
	return err
}

func PutURL(u *Url) error {
	const op = "internal.models.PutURL"

	if err := u.Validate(); err != nil {
		log.Errorf("%s: %s", op, err.Error())
		return err
	}

	err := (*db).Where("alias=?", u.Alias).Find(&u).Error
	if gorm.IsRecordNotFoundError(err) {
		log.Errorf("%s: %s", op, err.Error())
		return err
	}

	u.UpdatedAt = time.Now()
	err = (*db).Save(u).Error

	return err
}
