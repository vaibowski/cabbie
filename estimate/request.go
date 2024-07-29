package estimate

import "cabbie/models"

type Request struct {
	Origin      models.Location `json:"origin"`
	Destination models.Location `json:"destination"`
}
