package routers

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var registeredUser = []User{
	{
		Email:    "john@gmail.com",
		Password: "12345",
	},
	{
		Email:    "doe@gmail.com",
		Password: "54321",
	},
}
