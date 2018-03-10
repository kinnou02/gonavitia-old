package serializer

import (
	"github.com/golang/protobuf/proto"
	"github.com/kinnou02/gonavitia/pbnavitia"
	"github.com/kinnou02/gonavitia/responses"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func make_ptoject(id, name string) *pbnavitia.PtObject {
	return &pbnavitia.PtObject{
		Uri:          &id,
		Name:         &name,
		EmbeddedType: pbnavitia.NavitiaType_STOP_AREA.Enum(),
	}
}

func make_section() *pbnavitia.Section {
	from := make_ptoject("uri:from", "from")
	to := make_ptoject("uri:to", "to")
	return &pbnavitia.Section{Origin: from, Destination: to}
}

func make_journey() *pbnavitia.Journey {
	journey := pbnavitia.Journey{
		Origin:            nil,
		Destination:       nil,
		Duration:          proto.Int32(60),
		NbTransfers:       proto.Int32(0),
		DepartureDateTime: proto.Uint64(1000),
		ArrivalDateTime:   proto.Uint64(1060),
		RequestedDateTime: proto.Uint64(1000),
		Sections:          []*pbnavitia.Section{make_section()},
	}
	return &journey
}

func TestNewSectionNil(t *testing.T) {
	assert.Nil(t, NewSection(nil))
}

func TestNewSection(t *testing.T) {
	section := NewSection(make_section())
	assert.Equal(t, *section.From.Id, "uri:from")
	assert.Equal(t, *section.To.Id, "uri:to")
}

func TestNewJourneyNil(t *testing.T) {
	assert.Nil(t, NewJourney(nil))
}

func TestNewJourneyDirectPath(t *testing.T) {
	pb_journey := make_journey()
	journey := NewJourney(pb_journey)
	assert.Nil(t, journey.From)
	assert.Nil(t, journey.To)
	assert.Equal(t, journey.Duration, *pb_journey.Duration)
	assert.Equal(t, journey.DepartureDateTime, responses.NavitiaDatetime(time.Unix(1000, 0)))
	assert.Equal(t, journey.ArrivalDateTime, responses.NavitiaDatetime(time.Unix(1060, 0)))
	assert.Equal(t, journey.NbTransfers, *pb_journey.NbTransfers)
	assert.Equal(t, len(journey.Sections), len(pb_journey.Sections))
}

func TestNewJourneyResponseNil(t *testing.T) {
	assert.Nil(t, NewJourneysReponse(nil))
}

func TestNewJourneyResponseEmpty(t *testing.T) {
	pb_response := pbnavitia.Response{}
	response := NewJourneysReponse(&pb_response)
	assert.Equal(t, len(response.Journeys), 0)
}

func TestNewJourneyResponseOne(t *testing.T) {
	pb_response := pbnavitia.Response{
		Journeys: []*pbnavitia.Journey{make_journey()},
	}
	response := NewJourneysReponse(&pb_response)
	assert.Equal(t, len(response.Journeys), 1)
}

func TestNewJourneyResponseTwo(t *testing.T) {
	pb_response := pbnavitia.Response{
		Journeys: []*pbnavitia.Journey{
			make_journey(),
			make_journey(),
		},
	}
	response := NewJourneysReponse(&pb_response)
	assert.Equal(t, len(response.Journeys), 2)
}
