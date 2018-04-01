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
	Co2Emission       *Amount         `json:"co2_emission,omitempty"`
	AdditionalInfo    []string        `json:"additional_informations"`
	Links             []Link          `json:"links,omitempty"`
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
	Direction      *string `json:"direction"`
	Code           *string `json:"code"`
	Network        *string `json:"network"`
	Color          *string `json:"color"`
	Name           *string `json:"name"`
	PhysicalMode   *string `json:"physical_mode"`
	Headsign       *string `json:"headsign"`
	Label          *string `json:"label"`
	TextColor      *string `json:"text_color"`
	CommercialMode *string `json:"commercial_mode"`
	Description    *string `json:"description"`
	Links          []Link  `json:"links"`
}

type Amount struct {
	Value float64 `json:"value"`
	Unit  string  `json:"unit"`
}
type Link struct {
	Title     *string `json:"title,omitempty"`
	Id        *string `json:"id,omitempty"`
	Href      *string `json:"href,omitempty"`
	Rel       *string `json:"rel,omitempty"`
	Value     *string `json:"value,omitempty"`
	Type      *string `json:"type,omitempty"`
	Templated *bool   `json:"templated,omitempty"`
	Internal  *bool   `json:"internal,omitempty"`
}
