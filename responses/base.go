package responses

type Place struct {
	Id           *string `json:"id"`
	Name         *string `json:"name"`
	EmbeddedType *string `json:"embedded_type"`
	Quality      *int32  `json:"quality,omitempty"`
}
