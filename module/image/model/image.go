package imageModel

import (
	"errors"
	"go_service_food_organic/common"
)

const (
	EntityName = "Image"

	ErrFileUploadTooLarge = "ErrFileTooLarge"
	MsgFileTooLarge       = "file too large"

	ErrFileUploadIsNotImage = "ErrFileIsNotImage"
	MsgFileIsNotImage       = "file is not image"

	ErrCanNotSaveFile = "ErrCanNotSaveFile"
	MsgCanNotSaveFile = "can not save file"

	ErrInvalidImageFormat = "ErrInvalidImageFormat"
	MsgInvalidImageFormat = "unknown format image"

	MsgErrorFileExists = "file exists"
	ErrFileExists      = "ErrFileExists"

	MsgCanNotDeleteFileUpload = "can not delete file upload"
	ErrCanNotDeleteFileUpload = "ErrCanNotDeleteFileUpload"
)

type Image struct {
	common.SQLModel `json:",inline"`
	Url             string `json:"url" gorm:"column:url;"`
	Width           int    `json:"width" gorm:"column:width;"`
	Height          int    `json:"height" gorm:"column:height;"`
	HashValue       string `json:"hash_value" gorm:"column:hash_value;"`
	Type            string `json:"type" gorm:"column:type;"`
	CloudName       string `json:"cloud_name,omitempty" gorm:"-"`
	Extension       string `json:"extension,omitempty" gorm:"-"`
}

func (Image) TableName() string {
	return "images"
}

func (img *Image) Mark(isAdminOrOwner bool) {
	img.GetUID(common.OjbTypeImage)
}

type ImageProfile struct {
	common.SQLModel `json:",inline"`
	Url             string `json:"url" gorm:"column:url;"`
}

func (ImageProfile) TableName() string {
	return Image{}.TableName()
}

func (img *ImageProfile) Mark(isAdminOrOwner bool) {
	img.GetUID(common.OjbTypeImage)
}

func ErrorFileExists() *common.AppError {
	return common.NewCustomError(
		errors.New(MsgErrorFileExists),
		MsgErrorFileExists,
		ErrFileExists,
	)
}

func ErrorInvalidImageFormat(err error) *common.AppError {
	return common.NewCustomError(err, MsgInvalidImageFormat, ErrInvalidImageFormat)
}

func ErrFileTooLarge() *common.AppError {
	return common.NewCustomError(
		errors.New(MsgFileTooLarge),
		MsgFileTooLarge,
		ErrFileUploadTooLarge,
	)
}

func ErrFileIsNotImage(err error) *common.AppError {
	return common.NewCustomError(err, MsgFileIsNotImage, ErrFileUploadIsNotImage)
}

func CanNotServerSave(err error) *common.AppError {
	return common.NewCustomError(err, MsgCanNotSaveFile, ErrCanNotSaveFile)
}
func CanNotDeleteFileUpload(err error) *common.AppError {
	return common.NewCustomError(err, MsgCanNotDeleteFileUpload, ErrCanNotDeleteFileUpload)
}

type ErrorInfo struct {
	FileName string
	ImgInfo  *Image
	ErrInfo  error
}

//
//func (img *Image) Scan(value interface{}) error {
//	bytes, ok := value.([]byte)
//	if !ok {
//		return NewCustomError(nil, "Failed to unmarshal  JSON value", "ErrInternal")
//	}
//
//	var newImg Image
//	if err := json.Unmarshal(bytes, &newImg); err != nil {
//		return NewCustomError(nil, "Failed to decode  JSON value", "ErrInternal")
//	}
//
//	*img = newImg
//
//	return nil
//}
//
//func (img *Image) Value() (driver.Value, error) {
//	if img == nil {
//		return nil, nil
//	}
//	return json.Marshal(img)
//}
//
//type Images []Image
//
//func (imgs *Images) Scan(value interface{}) error {
//	bytes, ok := value.([]byte)
//	if !ok {
//		return NewCustomError(
//			nil,
//			fmt.Sprintf("Failed to unmarshal  JSON value: %s", value),
//			"ErrInternal")
//	}
//
//	var newImgs Images
//	if err := json.Unmarshal(bytes, &newImgs); err != nil {
//		return NewCustomError(
//			nil,
//			fmt.Sprintf("Failed to decode  JSON value: %s", value),
//			"ErrInternal")
//	}
//	*imgs = newImgs
//
//	return nil
//}
//
//func (imgs *Images) Value() (driver.Value, error) {
//	if imgs == nil {
//		return nil, nil
//	}
//	return json.Marshal(imgs)
//}
