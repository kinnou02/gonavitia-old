package responses

type JourneysResponse struct {
	Journeys []*Journey `json:"journeys"`
}

type Journey struct {
	From              *Place          `json:"from,omitempty"`
	To                *Place          `json:"to,omitempty"`
	Duration          int32           `json:"duration"`
	NbTransfers       int32           `json:"nb_transfers"`
	DepartureDateTime NavitiaDatetime `json:"departure_date_time"`
	ArrivalDateTime   NavitiaDatetime `json:"arrival_date_time"`
	RequestedDateTime NavitiaDatetime `json:"requested_date_time"`
	Type              *string         `json:"type"`
	Tags              []string        `json:"tags"`
	Sections          []*Section      `json:"sections"`
	Status            string          `json:"status"`
	Durations         *Durations      `json:"durations,omitempty"`
	Distances         *Distances      `json:"distances,omitempty"`
}

type Section struct {
	Id                string          `json:"id"`
	From              *Place          `json:"from,omitempty"`
	To                *Place          `json:"to,omitempty"`
	DepartureDateTime NavitiaDatetime `json:"departure_date_time"`
	ArrivalDateTime   NavitiaDatetime `json:"arrival_date_time"`
	Duration          int32           `json:"duration"`
	Type              string          `json:"type"`
}

type Durations struct {
	Total       int32 `json:"total"`
	Walking     int32 `json:"walking"`
	Bike        int32 `json:"bike"`
	Car         int32 `json:"car"`
	Ridesharing int32 `json:"ridesharing"`
}

type Distances struct {
	Total       int32 `json:"total"`
	Walking     int32 `json:"walking"`
	Bike        int32 `json:"bike"`
	Car         int32 `json:"car"`
	Ridesharing int32 `json:"ridesharing"`
}
