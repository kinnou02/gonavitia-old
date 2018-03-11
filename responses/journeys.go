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
	GeoJson           *GeoJson        `json:"geojson,omitempty"`
	Mode              *string         `json:"mode,omitempty"`
	TransferType      *string         `json:"transfer_type,omitempty"`
	DisplayInfo       *PtDisplayInfo  `json:"display_informations,omitempty"`
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

type PtDisplayInfo struct {
	Direction      *string `json:"direction,omitempty"`
	Code           *string `json:"code,omitempty"`
	Network        *string `json:"network,omitempty"`
	Color          *string `json:"color,omitempty"`
	Name           *string `json:"name,omitempty"`
	PhysicalMode   *string `json:"physical_mode,omitempty"`
	Headsign       *string `json:"headsign,omitempty"`
	Label          *string `json:"label,omitempty"`
	TextColor      *string `json:"text_color,omitempty"`
	CommercialMode *string `json:"commercial_mode,omitempty"`
	Description    *string `json:"description,omitempty"`
}
