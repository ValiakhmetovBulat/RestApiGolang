package models

import (
	"RestApiGolang/internal/database/sqlite"
	log "RestApiGolang/internal/logger"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

var (
	db                   = &sqlite.Db
	ErrAliasNotSpecified = errors.New("alias for URL not specified")
	ErrUrlNotSpecified   = errors.New("URL not specified")
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
	const op = "internal.database.sqlite.SaveURL"
	if err := u.Validate(); err != nil {
		return err
	}

	err := (*db).Save(u).Error
	if err != nil {
		log.Errorf("%s: %s", op, err.Error())
		return err
	}

	return err
}

func GetURL(alias string) (string, error) {
	const op = "internal.database.sqlite.GetURL"

	u := Url{}
	err := (*db).Where("alias=?", alias).Find(&u).Error
	if err != nil {
		log.Errorf("%s: %s", op, err.Error())
		return "", err
	}

	return u.Url, nil
}

func GetURLs() ([]Url, error) {
	const op = "internal.database.sqlite.GetURLs"

	us := []Url{}
	err := (*db).Table("urls").Find(&us).Error
	if err != nil {
		log.Errorf("%s: %s", op, err.Error())
		return us, err
	}

	return us, nil
}

func DeleteURL(alias string) error {
	const op = "internal.database.sqlite.DeleteURL"

	u := Url{}
	err := (*db).Where("alias=?", alias).Find(&u).Error
	if gorm.IsRecordNotFoundError(err) {
		return fmt.Errorf("%s: %w", op, err)
	}
	err = (*db).Delete(u).Error
	if err != nil {
		log.Errorf("%s: %s", op, err.Error())
		return err
	}
	return err
}

func PutURL(u *Url) error {
	const op = "internal.database.sqlite.PutURL"

	if err := u.Validate(); err != nil {
		log.Errorf("%s: %s", op, err.Error())
		return err
	}
	err := (*db).Update(u).Error
	return err
}
