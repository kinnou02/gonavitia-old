package responses

import "time"

type JourneysResponse struct {
	Journeys []*Journey `json:"journeys"`
}

type Journey struct {
	From              *Place     `json:"from,omitempty"`
	To                *Place     `json:"to,omitempty"`
	Duration          int32      `json:"duration"`
	NbTransfers       int32      `json:"nb_transfers"`
	DepartureDateTime time.Time  `json:"departure_date_time"`
	ArrivalDateTime   time.Time  `json:"arrival_date_time"`
	RequestedDateTime time.Time  `json:"requested_date_time"`
	Type              *string    `json:"type"`
	Tags              []string   `json:"tags"`
	Sections          []*Section `json:"sections"`
}

type Section struct {
	From *Place `json:"from,omitempty"`
	To   *Place `json:"to,omitempty"`
}
