package v1

type UserCreateParams struct {
	FirstName  string `form:"first_name" binding:"required"`
	MiddleName string `form:"middle_name" binding:"required"`
	LastName   string `form:"last_name" binding:"required"`
	Email      string `form:"email" binding:"required"`
}
