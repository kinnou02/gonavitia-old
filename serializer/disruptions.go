package serializer

import (
	"github.com/kinnou02/gonavitia/pbnavitia"
	"github.com/kinnou02/gonavitia/responses"
	"strings"
	"time"
)

func NewDisruption(pb *pbnavitia.Impact) *responses.Disruption {
	if pb == nil {
		return nil
	}
	status := strings.ToLower(pb.Status.String())
	d := responses.Disruption{
		ID:            pb.Uri,
		Uri:           pb.Uri,
		DisruptionUri: pb.DisruptionUri,
		ImpactID:      pb.Uri,
		DisruptionID:  pb.Uri,
		Cause:         pb.Cause,
		Contributor:   pb.Contributor,
		Category:      pb.Category,
		UpdatedAt:     responses.NavitiaDatetime(time.Unix(int64(pb.GetUpdatedAt()), 0)),
		Status:        &status,
		Severity:      NewSeverity(pb.Severity),
	}
	for _, message := range pb.Messages {
		d.Messages = append(d.Messages, NewMessage(message))
	}
	for _, period := range pb.ApplicationPeriods {
		d.ApplicationPeriods = append(d.ApplicationPeriods, NewPeriod(period))
	}
	//TODO add properties
	return &d
}

func NewPeriod(pb *pbnavitia.Period) *responses.Period {
	if pb == nil {
		return nil
	}
	p := responses.Period{
		Begin: responses.NavitiaDatetime(time.Unix(int64(pb.GetBegin()), 0)),
		End:   responses.NavitiaDatetime(time.Unix(int64(pb.GetEnd()), 0)),
	}
	return &p
}

func NewSeverity(pb *pbnavitia.Severity) *responses.Severity {
	if pb == nil {
		return nil
	}
	s := responses.Severity{
		Name:     pb.Name,
		Priority: pb.GetPriority(),
		Color:    pb.Color,
		Effect:   pb.Effect,
	}
	return &s
}

func NewMessage(pb *pbnavitia.MessageContent) *responses.Message {
	if pb == nil {
		return nil
	}
	message := responses.Message{
		Text:    pb.Text,
		Channel: NewChannel(pb.Channel),
	}
	return &message
}

func NewChannel(pb *pbnavitia.Channel) *responses.Channel {
	if pb == nil {
		return nil
	}
	channel := responses.Channel{
		ID:          pb.Id,
		Name:        pb.Name,
		ContentType: pb.ContentType,
	}
	for _, types := range pb.ChannelTypes {
		channel.Types = append(channel.Types, types.String())
	}
	return &channel
}
