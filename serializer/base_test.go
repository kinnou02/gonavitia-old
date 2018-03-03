package serializer

import "github.com/kinnou02/pbnavitia"
import "testing"
import "github.com/stretchr/testify/assert"
import "github.com/golang/protobuf/proto"


func TestNewPlaceNil(t *testing.T){
    assert.Nil(t, NewPlace(nil))
}

func TestNewPlace(t *testing.T){
    pb := pbnavitia.PtObject{Uri: proto.String("foo"), Name: proto.String("bar")}
    place := NewPlace(&pb)
    assert.Equal(t, *place.Id, "foo")
    assert.Equal(t, *place.Name, "bar")
}