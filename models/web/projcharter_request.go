package web

type ProjectCharterRequest struct {
	ID          uint64 `json:"id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	User_id     uint64 `json:"user_id"`
	UpdatedBy   string `json:"updated_by"`
	DeletedBy   string `json:"deleted_by"`
}
