package dto

type (
	GetAntreanFarmasiRequestV2 struct {
		Kodebooking string `json:"kodebooking" validate:"required"`
	}
)
