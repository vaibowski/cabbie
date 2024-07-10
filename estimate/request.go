package estimate

type Request struct {
	Pickup  Location `json:"pickup"`
	DropOff Location `json:"dropOff"`
}

type Location struct {
	XCoordinate float64 `json:"xCoordinate"`
}
