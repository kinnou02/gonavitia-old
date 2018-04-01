package responses

import (
	"fmt"
	"time"
)

type Place struct {
	Id           *string    `json:"id"`
	Name         *string    `json:"name"`
	EmbeddedType *string    `json:"embedded_type"`
	Quality      *int32     `json:"quality,omitempty"`
	StopPoint    *StopPoint `json:"stop_point,omitempty"`
	StopArea     *StopArea  `json:"stop_area,omitempty"`
	Admin        *Admin     `json:"administrative_region,omitempty"`
	Address      *Address   `json:"address,omitempty"`
}

type Pagination struct {
	StartPage    int32 `json:"start_page"`
	ItemsOnPage  int32 `json:"items_on_page"`
	ItemsPerPage int32 `json:"items_per_page"`
	TotalResult  int32 `json:"total_result"`
}

type Error struct {
	Id      *string `json:"id"`
	Message *string `json:"message"`
}

type NavitiaDatetime time.Time

func (t NavitiaDatetime) MarshalJSON() ([]byte, error) {
	//do your serializing here
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("20060102T150405"))
	return []byte(stamp), nil
}

type Coord struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type Code struct {
	Type  *string `json:"type"`
	Value *string `json:"value"`
}

type StopArea struct {
	Id         *string      `json:"id"`
	Name       *string      `json:"name"`
	Label      *string      `json:"label"`
	Timezone   *string      `json:"Timezone,omitempty"`
	Coord      *Coord       `json:"coord"`
	Admins     []*Admin     `json:"administrative_regions"`
	Codes      []*Code      `json:"codes,omitempty"`
	StopPoints []*StopPoint `json:"stop_points,omitempty"`
}

type StopPoint struct {
	Id       *string   `json:"id"`
	Name     *string   `json:"name"`
	Label    *string   `json:"label"`
	Coord    *Coord    `json:"coord"`
	Admins   []*Admin  `json:"administrative_regions"`
	Codes    []*Code   `json:"codes,omitempty"`
	StopArea *StopArea `json:"stop_area,omitempty"`
}

type Admin struct {
	Id      *string `json:"id"`
	Name    *string `json:"name"`
	Label   *string `json:"label"`
	Coord   *Coord  `json:"coord"`
	Insee   *string `json:"insee,omitempty"`
	ZipCode *string `json:"zip_code,omitempty"`
}

type Address struct {
	Id          *string  `json:"id"`
	Name        *string  `json:"name"`
	Label       *string  `json:"label"`
	Coord       *Coord   `json:"coord"`
	HouseNumber *int32   `json:"house_number,omitempty"`
	Admins      []*Admin `json:"administrative_regions"`
}
