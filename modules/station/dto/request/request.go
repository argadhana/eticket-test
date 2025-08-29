package request

type CreateStationRequest struct {
	Name     string `json:"name" validate:"required"`
	Location string `json:"location"`
}
