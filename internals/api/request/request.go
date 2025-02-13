package request

type CountryQuery struct {
	Name string `form:"name" binding:"required"`
}
