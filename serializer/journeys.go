package serializer

import "github.com/kinnou02/gonavitia/responses"
import "github.com/kinnou02/gonavitia/pbnavitia"
import "time"
import "strings"
import "github.com/golang/protobuf/proto"

func NewJourneysReponse(pb *pbnavitia.Response) *responses.JourneysResponse {
	if pb == nil {
		return nil
	}
	r := responses.JourneysResponse{}
	for _, pb_journey := range pb.Journeys {
		r.Journeys = append(r.Journeys, NewJourney(pb_journey))
	}
	return &r
}

func NewJourney(pb *pbnavitia.Journey) *responses.Journey {
	if pb == nil {
		return nil
	}
	journey := responses.Journey{
		From:              NewPlace(pb.Origin),
		To:                NewPlace(pb.Destination),
		Duration:          pb.GetDuration(),
		NbTransfers:       pb.GetNbTransfers(),
		DepartureDateTime: responses.NavitiaDatetime(time.Unix(int64(pb.GetDepartureDateTime()), 0)),
		ArrivalDateTime:   responses.NavitiaDatetime(time.Unix(int64(pb.GetArrivalDateTime()), 0)),
		RequestedDateTime: responses.NavitiaDatetime(time.Unix(int64(pb.GetRequestedDateTime()), 0)),
		Status:            pb.GetMostSeriousDisruptionEffect(),
		Durations:         NewDurations(pb.Durations),
		Distances:         NewDistances(pb.Distances),
		Tags:              make([]string, 0),
	}
	for _, pb_section := range pb.Sections {
		journey.Sections = append(journey.Sections, NewSection(pb_section))
	}
	return &journey
}

func NewSection(pb *pbnavitia.Section) *responses.Section {
	if pb == nil {
		return nil
	}
	var mode *string
	if sn := pb.StreetNetwork; sn != nil {
		m := strings.ToLower(sn.Mode.String())
		mode = &m
	}
	var transferType *string
	if pb.TransferType != nil {
		t := strings.ToLower(pb.TransferType.String())
		transferType = &t
	}
	section := responses.Section{
		Id:                pb.GetId(),
		From:              NewPlace(pb.Origin),
		To:                NewPlace(pb.Destination),
		DepartureDateTime: responses.NavitiaDatetime(time.Unix(int64(pb.GetBeginDateTime()), 0)),
		ArrivalDateTime:   responses.NavitiaDatetime(time.Unix(int64(pb.GetEndDateTime()), 0)),
		Duration:          pb.GetDuration(),
		Type:              strings.ToLower(pb.GetType().String()),
		GeoJson:           NewGeoJson(pb),
		Mode:              mode,
		TransferType:      transferType,
		DisplayInfo:       NewPtDisplayInfo(pb.PtDisplayInformations),
		Co2Emission:       NewCo2Emission(pb.Co2Emission),
		AdditionalInfo:    make([]string, 0),
		Links:             NewLinksFromUris(pb.PtDisplayInformations),
	}
	for _, info := range pb.GetAdditionalInformations() {
		section.AdditionalInfo = append(section.AdditionalInfo, strings.ToLower(info.String()))
	}

	return &section
}

func NewDurations(pb *pbnavitia.Durations) *responses.Durations {
	if pb == nil {
		return nil
	}
	durations := responses.Durations{
		Total:       pb.GetTotal(),
		Walking:     pb.GetWalking(),
		Bike:        pb.GetBike(),
		Car:         pb.GetCar(),
		Ridesharing: pb.GetRidesharing(),
	}
	return &durations
}

func NewDistances(pb *pbnavitia.Distances) *responses.Distances {
	if pb == nil {
		return nil
	}
	distances := responses.Distances{
		Walking:     pb.GetWalking(),
		Bike:        pb.GetBike(),
		Car:         pb.GetCar(),
		Ridesharing: pb.GetRidesharing(),
	}
	return &distances
}

func NewPtDisplayInfo(pb *pbnavitia.PtDisplayInfo) *responses.PtDisplayInfo {
	if pb == nil {
		return nil
	}
	info := responses.PtDisplayInfo{
		Direction:      pb.Direction,
		Code:           pb.Code,
		Network:        pb.Network,
		Color:          pb.Color,
		Name:           pb.Name,
		PhysicalMode:   pb.PhysicalMode,
		Headsign:       pb.Headsign,
		TextColor:      pb.TextColor,
		CommercialMode: pb.CommercialMode,
		Description:    pb.Description,
		Links:          make([]responses.Link, 0),
	}
	return &info

}

func NewCo2Emission(pb *pbnavitia.Co2Emission) *responses.Amount {
	if pb == nil {
		return nil
	}
	co2 := responses.Amount{
		Value: pb.GetValue(),
		Unit:  pb.GetUnit(),
	}
	return &co2

}

func NewLinksFromUris(pb *pbnavitia.PtDisplayInfo) []responses.Link {
	if pb == nil || pb.Uris == nil {
		return nil
	}
	uris := pb.Uris
	res := make([]responses.Link, 0)
	res = appendLinksFromUri(uris.Company, "company", &res)
	res = appendLinksFromUri(uris.VehicleJourney, "vehicle_journey", &res)
	res = appendLinksFromUri(uris.Line, "line", &res)
	res = appendLinksFromUri(uris.Route, "route", &res)
	res = appendLinksFromUri(uris.CommercialMode, "commercial_mode", &res)
	res = appendLinksFromUri(uris.PhysicalMode, "physical_mode", &res)
	res = appendLinksFromUri(uris.Network, "Network", &res)
	res = appendLinksFromUri(uris.Note, "note", &res)
	res = appendLinksFromUri(uris.JourneyPattern, "journey_pattern", &res)
	return res
}

func appendLinksFromUri(pb *string, typ string, links *[]responses.Link) []responses.Link {
	if pb == nil {
		return *links
	}
	return append(*links, responses.Link{Id: pb, Type: proto.String(typ)})
}
