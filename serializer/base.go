package serializer

import (
	"strings"

	"github.com/kinnou02/gonavitia/pbnavitia"
	"github.com/kinnou02/gonavitia/responses"
)

func NewError(pb *pbnavitia.Error) *responses.Error {
	if pb == nil {
		return nil
	}
	id := pb.Id.Enum().String()
	return &responses.Error{
		Id:      &id,
		Message: pb.Message,
	}
}

func NewCode(pb *pbnavitia.Code) *responses.Code {
	if pb == nil {
		return nil
	}
	return &responses.Code{
		Type:  pb.Type,
		Value: pb.Value,
	}
}

func NewPlace(pb *pbnavitia.PtObject) *responses.Place {
	if pb == nil {
		return nil
	}
	t := strings.ToLower(pb.EmbeddedType.String())
	place := responses.Place{
		Id:           pb.Uri,
		Name:         pb.Name,
		EmbeddedType: &t,
		Quality:      pb.Quality,
		StopPoint:    NewStopPoint(pb.StopPoint),
		StopArea:     NewStopArea(pb.StopArea),
		Admin:        NewAdmin(pb.AdministrativeRegion),
		Address:      NewAddress(pb.Address),
	}
	return &place
}

func NewAdmin(pb *pbnavitia.AdministrativeRegion) *responses.Admin {
	if pb == nil {
		return nil
	}
	admin := responses.Admin{
		Id:      pb.Uri,
		Name:    pb.Name,
		Label:   pb.Label,
		Coord:   NewCoord(pb.Coord),
		Insee:   pb.Insee,
		ZipCode: pb.ZipCode,
	}
	return &admin
}

func NewCoord(pb *pbnavitia.GeographicalCoord) *responses.Coord {
	if pb == nil {
		return nil
	}
	coord := responses.Coord{
		Lat: pb.GetLat(),
		Lon: pb.GetLon(),
	}
	return &coord
}

func NewStopPoint(pb *pbnavitia.StopPoint) *responses.StopPoint {
	if pb == nil {
		return nil
	}
	sp := responses.StopPoint{
		Id:       pb.Uri,
		Name:     pb.Name,
		Label:    pb.Label,
		Coord:    NewCoord(pb.Coord),
		Admins:   make([]*responses.Admin, 0),
		StopArea: NewStopArea(pb.StopArea),
	}
	for _, pb_admin := range pb.AdministrativeRegions {
		sp.Admins = append(sp.Admins, NewAdmin(pb_admin))
	}
	for _, code := range pb.Codes {
		sp.Codes = append(sp.Codes, NewCode(code))
	}
	return &sp
}

func NewStopArea(pb *pbnavitia.StopArea) *responses.StopArea {
	if pb == nil {
		return nil
	}
	sa := responses.StopArea{
		Id:       pb.Uri,
		Name:     pb.Name,
		Label:    pb.Label,
		Timezone: pb.Timezone,
		Coord:    NewCoord(pb.Coord),
		Admins:   make([]*responses.Admin, 0),
	}
	for _, pb_admin := range pb.AdministrativeRegions {
		sa.Admins = append(sa.Admins, NewAdmin(pb_admin))
	}
	for _, code := range pb.Codes {
		sa.Codes = append(sa.Codes, NewCode(code))
	}
	for _, sp := range pb.StopPoints {
		sa.StopPoints = append(sa.StopPoints, NewStopPoint(sp))
	}
	return &sa
}

func NewAddress(pb *pbnavitia.Address) *responses.Address {
	if pb == nil {
		return nil
	}
	address := responses.Address{
		Id:          pb.Uri,
		Name:        pb.Name,
		Label:       pb.Label,
		Coord:       NewCoord(pb.Coord),
		HouseNumber: pb.HouseNumber,
		Admins:      make([]*responses.Admin, 0),
	}
	for _, pb_admin := range pb.AdministrativeRegions {
		address.Admins = append(address.Admins, NewAdmin(pb_admin))
	}
	return &address
}
