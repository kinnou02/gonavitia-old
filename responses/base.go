package responses

import (
	"fmt"
	"time"
)

type Place struct {
	Id           *string `json:"id"`
	Name         *string `json:"name"`
	EmbeddedType *string `json:"embedded_type"`
	Quality      *int32  `json:"quality,omitempty"`
}

type NavitiaDatetime time.Time

func (t NavitiaDatetime) MarshalJSON() ([]byte, error) {
	//do your serializing here
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("20060102T150405"))
	return []byte(stamp), nil
}
