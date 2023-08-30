package request

type LeadCreateRequest struct {
	Name      string `validate:"required min=3,max=100" json:"name"`
	Email     string `validate:"required email" json:"email"`
	Phone     string `validate:"required min=11" json:"phone"`
	VehicleId int    `validate:"required" json:"vehicleid"`
}