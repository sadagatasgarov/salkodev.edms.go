package controller

// For simple requests, with only one uid field
type UIDRequest struct {
	UID string `json:"uid" binding:"required"`
}
