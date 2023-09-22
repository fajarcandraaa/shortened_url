package entity

type Error string

// Declare error messege
const (
	ErrPermissionNotAllowed = Error("permission.not_allowed")
	ErrExpiredTime          = Error("Expired Link URL")
)

func (e Error) Error() string {
	return string(e)
}
