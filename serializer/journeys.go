package serializer

import "github.com/kinnou02/gonavitia/responses"
import "github.com/kinnou02/gonavitia/pbnavitia"
import "time"

func NewJourneysReponse(pb *pbnavitia.Response) responses.JourneysResponse{
    r := responses.JourneysResponse{}
    for _, pb_journey := range pb.Journeys {
        r.Journeys = append(r.Journeys, NewJourney(pb_journey))
    }
    return r
}

func NewJourney(pb *pbnavitia.Journey) *responses.Journey{
    journey := responses.Journey{
        From: NewPlace(pb.Origin),
        To: NewPlace(pb.Destination),
        Duration: pb.GetDuration(),
        NbTransfers: pb.GetNbTransfers(),
        DepartureDateTime: time.Unix(int64(pb.GetDepartureDateTime()), 0),
        ArrivalDateTime: time.Unix(int64(pb.GetArrivalDateTime()), 0),
        RequestedDateTime: time.Unix(int64(pb.GetRequestedDateTime()), 0),
    }
    for _, pb_section := range pb.Sections {
        journey.Sections = append(journey.Sections, NewSection(pb_section))
    }
    return &journey
}

func NewSection(pb *pbnavitia.Section) *responses.Section{
    section := responses.Section{
        From: NewPlace(pb.Origin),
        To: NewPlace(pb.Destination),
    }
    return &section
}
