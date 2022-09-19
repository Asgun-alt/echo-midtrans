package common

import "errors"

var (
	ErrRecordNotFound = errors.New(RecordNotFound)

	ErrUserAlreadyCreated = errors.New(UserAlreadyCreated)
	ErrPasswordNotMatch   = errors.New(PasswordNotMatch)
	ErrPasswordNotSame    = errors.New(PasswordNotSame)
	ErrInvalidPassword    = errors.New(InvalidPassword)
	ErrInvalidUserID      = errors.New(InvalidUserID)

	ErrCampaignAlreadyCreated = errors.New(CampaignAlreadyCreated)
	ErrFileImageError         = errors.New(FileImageError)
)
