package response

type LeadResponse struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	VehicleId int    `json:"vehicleId"`
}