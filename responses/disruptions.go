package responses

type Disruption struct {
	ID                 *string              `json:"id"`
	Uri                *string              `json:"id,omitempty"`
	DisruptionUri      *string              `json:"disruption_uri,omitempty"`
	ImpactID           *string              `json:"impact_id,omitempty"`
	DisruptionID       *string              `json:"disruption_id,omitempty"`
	Contributor        *string              `json:"contributor,omitempty"`
	Category           *string              `json:"category,omitempty"`
	Cause              *string              `json:"cause,omitempty"`
	Status             *string              `json:"status,omitempty"`
	Tags               []string             `json:"tags,omitempty"`
	UpdatedAt          NavitiaDatetime      `json:"updated_at,omitempty"`
	Severity           *Severity            `json:"severity"`
	ApplicationPeriods []*Period            `json:"application_periods"`
	Properties         []DisruptionProperty `json:"properties,omitempty"`
	Messages           []*Message           `json:"messages"`
}

type Severity struct {
	Name     *string `json:"name"`
	Effect   *string `json:"effect"`
	Color    *string `json:"color"`
	Priority int32   `json:"priority"`
}

type Message struct {
	Text    *string  `json:"text"`
	Channel *Channel `json:"channel"`
}

type Channel struct {
	ID          *string  `json:"id"`
	Name        *string  `json:"name"`
	ContentType *string  `json:"content_type"`
	Types       []string `json:"types"`
}

type Period struct {
	Begin NavitiaDatetime `json:"begin"`
	End   NavitiaDatetime `json:"end"`
}

type DisruptionProperty struct {
	Type  *string `json:"type"`
	Key   *string `json:"key"`
	Value *string `json:"value"`
}

type ImpactedObject struct {
	PtObject *Place `json:"pt_object,omitempty"` //TODO: should be a ptobject
	//	ImpactedSection *ImpactedSection `json:"impacted_section,omitempty"`
	//	ImpactedStops   []ImpactedStops  `json:"impacted_stops,omitempty"`
}
