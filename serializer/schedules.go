package serializer

import (
	"github.com/golang/protobuf/proto"
	"github.com/kinnou02/gonavitia/pbnavitia"
	"github.com/kinnou02/gonavitia/responses"
	"strings"
	"time"
)

func NewRouteSchedulesResponse(pb *pbnavitia.Response) *responses.RouteScheduleResponse {
	if pb == nil {
		return nil
	}
	response := responses.RouteScheduleResponse{
		Error:          NewError(pb.Error),
		RouteSchedules: make([]*responses.RouteSchedule, 0),
	}
	for _, r := range pb.RouteSchedules {
		response.RouteSchedules = append(response.RouteSchedules, NewRouteSchedule(r))
	}
	return &response
}

func NewRouteSchedule(pb *pbnavitia.RouteSchedule) *responses.RouteSchedule {
	if pb == nil {
		return nil
	}
	var additionalInfo *string
	info := pb.ResponseStatus
	if info != nil {
		tmp := strings.ToLower(info.Enum().String())
		additionalInfo = &tmp
	}
	return &responses.RouteSchedule{
		DisplayInfo:    NewPtDisplayInfo(pb.PtDisplayInformations),
		Table:          NewTable(pb.Table),
		AdditionalInfo: additionalInfo,
		GeoJson:        NewGeoJsonMultistring(pb.Geojson),
		Links:          NewLinksFromUris(pb.PtDisplayInformations),
	}
}

func NewTable(pb *pbnavitia.Table) *responses.Table {
	if pb == nil {
		return nil
	}
	t := responses.Table{
		Headers: make([]*responses.Header, 0),
		Rows:    make([]*responses.Row, 0),
	}
	for _, h := range pb.Headers {
		t.Headers = append(t.Headers, NewHeader(h))
	}
	for _, r := range pb.Rows {
		t.Rows = append(t.Rows, NewRow(r))
	}
	return &t
}

func NewHeader(pb *pbnavitia.Header) *responses.Header {
	if pb == nil {
		return nil
	}
	return &responses.Header{
		DisplayInfo: NewPtDisplayInfo(pb.PtDisplayInformations),
		Links:       NewLinksFromUris(pb.PtDisplayInformations),
		//AdditionalInfo []string       `json:"additional_informations"`
	}
}

func NewRow(pb *pbnavitia.RouteScheduleRow) *responses.Row {
	if pb == nil {
		return nil
	}
	r := responses.Row{
		StopPoint: NewStopPoint(pb.StopPoint),
		DateTimes: make([]*responses.DateTime, 0, len(pb.DateTimes)),
	}
	for _, d := range pb.DateTimes {
		r.DateTimes = append(r.DateTimes, NewDatetime(d))
	}
	return &r
}

func NewDatetime(pb *pbnavitia.ScheduleStopTime) *responses.DateTime {
	if pb == nil {
		return nil
	}
	rtLevel := strings.ToLower(pb.GetRealtimeLevel().Enum().String())
	return &responses.DateTime{
		DateTime:       responses.NavitiaDatetime(time.Unix(int64(pb.GetDate()+pb.GetTime()), 0)),
		BaseDateTime:   responses.NavitiaDatetime(time.Unix(int64(pb.GetBaseDateTime()), 0)),
		AdditionalInfo: make([]string, 0),
		DataFreshness:  rtLevel,
		Links:          NewLinksFromProperties(pb.Properties),
	}
}

func NewLinksFromProperties(pb *pbnavitia.Properties) []responses.Link {
	result := []responses.Link{}
	if pb == nil {
		return result
	}
	if pb.VehicleJourneyId != nil {
		result = append(result, responses.Link{
			Id:    pb.VehicleJourneyId,
			Value: pb.VehicleJourneyId,
			Type:  proto.String("vehicle_journey"),
			Rel:   proto.String("vehicle_journeys"),
		})
	}
	return result
}
