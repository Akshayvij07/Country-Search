package request

//request queryParams
type CountryQuery struct {
	Name string `form:"name" binding:"required"`
}
