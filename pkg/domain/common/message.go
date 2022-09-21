package common

const (
	DataSuccess     = "processing data success"
	DataFailed      = "processing data failed"
	UnknownError    = "unknown error detected, please contact administrator"
	RecordNotFound  = "requested record is not found"
	URINotFound     = "your requested URI is not found"
	ValidationError = "please check the validation error"

	InvalidUserID      = "invalid user id"
	UserAlreadyCreated = "the user is already created"
	InvalidPassword    = "password must have at least 8 characters with minimum of 1 uppercase, 1 number, and 1 special character"
	PasswordNotSame    = "password 1 and password 2 must be the same"
	PasswordNotMatch   = "invalid username/password. please try again"
	PasswordEmpty      = "please fill the password"

	CampaignAlreadyCreated = "the campaign is already created"
	InvalidCampaignID      = "invalid campaign id"
	FileImageError         = "something is wrong when uploading image"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Errors  interface{} `json:"errors"`
	Code    int         `json:"code"`
}
