package countrydto

type CountryResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name" form:"name" validate:"required"`
	// Subcribe bool   `json:"subcribe" form:"subcribe"`
}
