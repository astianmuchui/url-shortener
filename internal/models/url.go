package models

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/astianmuchui/url-shortener/internal/db"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Url struct {
	gorm.Model
	Uid        uuid.UUID
	ShortPath  string `json:"short_path" gorm:"uniqueIndex;not null"`
	TargetPath string `json:"target_path"`
}

type UrlCreateRequest struct {
	TargetPath string `json:"target_path"`
}

type URLResponse struct {
	ShortPath  string `json:"short_path"`
	TargetPath string `json:"target_path"`
}

func GenerateShortPath() string {

	bytes := make([]byte, 8)

	_, err := rand.Read(bytes)

	if err != nil {
		log.Error("Unable to generate token:", err)
	}

	token := hex.EncodeToString(bytes)
	log.Info("Generated Token:", token)

	return token[:4]
}

func (u *Url) BeforeCreate(tx *gorm.DB) error {
	u.Uid = uuid.New()
	u.ShortPath = GenerateShortPath()
	return nil
}

func (u *Url) ToResponse(baseUrl string) *URLResponse {
	return &URLResponse{
		ShortPath:  baseUrl + u.ShortPath,
		TargetPath: u.TargetPath,
	}
}

func (u *Url) Create() error {
	return db.DB.Create(u).Error
}

func (u *Url) Retreive() error {
	return db.DB.Model(Url{}).Where("uid = ? or short_path = ?", u.Uid, u.ShortPath).Find(&u).Error
}

func (u *Url) Update() error {
	return db.DB.Save(u).Error
}

func (u *Url) Delete() error {
	return db.DB.Delete(u).Error
}
