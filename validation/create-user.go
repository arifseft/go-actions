package validation

type CreateUserSchema struct {
	Name string `validate:"required"`
}
