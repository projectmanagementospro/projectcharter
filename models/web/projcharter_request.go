package web

type ProjectCharterRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	User_id     uint64 `json:"user_id"`
	UpdatedBy   string `json:"updated_by"`
	DeletedBy   string `json:"deleted_by"`
}

type ProjectCharterUpdateRequest struct {
	ID          uint   `json:"id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	User_id     uint64 `json:"user_id"`
	UpdatedBy   string `json:"updated_by" binding:"required"`
	DeletedBy   string `json:"deleted_by"`
}
