package api

type YourStruct struct {
	Name     string `json:"name"`
	Download string `json:"download"`
	Notes    string `json:"notes"`
}

type AdministratorInformation struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Token    string `json:"token"`
	Access   string `json:"access"`
}
