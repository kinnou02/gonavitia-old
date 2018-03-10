package serializer

import "github.com/kinnou02/gonavitia/responses"
import "github.com/kinnou02/gonavitia/pbnavitia"
import "time"

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
	section := responses.Section{
		Id:                pb.GetId(),
		From:              NewPlace(pb.Origin),
		To:                NewPlace(pb.Destination),
		DepartureDateTime: responses.NavitiaDatetime(time.Unix(int64(pb.GetBeginDateTime()), 0)),
		ArrivalDateTime:   responses.NavitiaDatetime(time.Unix(int64(pb.GetEndDateTime()), 0)),
		Duration:          pb.GetDuration(),
		Type:              pb.GetType().String(),
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
	durations := responses.Distances{
		Walking:     pb.GetWalking(),
		Bike:        pb.GetBike(),
		Car:         pb.GetCar(),
		Ridesharing: pb.GetRidesharing(),
	}
	return &durations
}
