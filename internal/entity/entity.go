package entity

type Error string

// Declare error messege
const (
	ErrPermissionNotAllowed = Error("permission.not_allowed")
)

func (e Error) Error() string {
	return string(e)
}
