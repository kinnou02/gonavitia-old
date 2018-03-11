package serializer

import "github.com/kinnou02/gonavitia/pbnavitia"
import "testing"
import "github.com/stretchr/testify/assert"
import "github.com/golang/protobuf/proto"

func TestNewPlaceNil(t *testing.T) {
	assert.Nil(t, NewPlace(nil))
}

func TestNewPlace(t *testing.T) {
	pb := pbnavitia.PtObject{Uri: proto.String("foo"), Name: proto.String("bar"), EmbeddedType: pbnavitia.NavitiaType_STOP_AREA.Enum()}
	place := NewPlace(&pb)
	assert.Equal(t, *place.Id, "foo")
	assert.Equal(t, *place.Name, "bar")
	assert.Equal(t, *place.EmbeddedType, "stop_area")
}
