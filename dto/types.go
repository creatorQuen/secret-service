package dto

type UserCreateReq struct {
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SecretPutReq struct {
	Id     string `json:"id"`
	Secret string `json:"secret"`
}
