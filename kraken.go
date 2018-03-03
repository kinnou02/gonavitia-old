package gonavitia

import (
	"github.com/golang/protobuf/proto"
	"github.com/kinnou02/pbnavitia"
	zmq "github.com/pebbe/zmq2"
	"github.com/pkg/errors"
	"time"
)

var (
	KrakenTimeout = errors.New("kraken timeout")
)

type Kraken struct {
	Name    string
	Addr    string
	Timeout time.Duration
}

func (k *Kraken) Call(request *pbnavitia.Request) (*pbnavitia.Response, error) {
	requester, _ := zmq.NewSocket(zmq.REQ)
	requester.Connect(k.Addr)
	defer requester.Close()
	data, _ := proto.Marshal(request)
	requester.Send(string(data), 0)
	poller := zmq.NewPoller()
	poller.Add(requester, zmq.POLLIN)
	p, err := poller.Poll(k.Timeout)
	if err != nil {
		return nil, err
	}
	if len(p) < 1 {
		return nil, errors.Wrapf(KrakenTimeout, "calling kraken %s", k.Name)
	}
	raw_resp, _ := p[0].Socket.Recv(0)
	resp := &pbnavitia.Response{}
	_ = proto.Unmarshal([]byte(raw_resp), resp)
	return resp, nil
}
