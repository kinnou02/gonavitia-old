package serializer

import "github.com/kinnou02/gonavitia/responses"
import "github.com/kinnou02/pbnavitia"


func NewPlace(pb *pbnavitia.PtObject) *responses.Place{
    if pb == nil {
        return nil
    }
    place := responses.Place{pb.Uri, pb.Name}
    return &place
}
