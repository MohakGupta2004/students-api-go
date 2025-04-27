package types

type Student struct {
	Id         string
	Name       string `validate:"required"`
	Profession string `validate:"required"`
}
