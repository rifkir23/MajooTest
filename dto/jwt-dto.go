package dto

type UserValidateDTO struct {
	AccessUuid string `json:"access_uuid"`
	UserId     uint64 `json:"user_id"`
	Username   string `json:"username"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
	Exp        string `json:"exp"`
}
