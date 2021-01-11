package main

import (
	"encoding/json"
	"fmt"
)

type messageStat struct {
	Name           string `json:"name"`
	Origin         string `json:"origin"`
	CalledRecvmmsg int64  `json:"called.recvmmsg"`
	CalledRecvmsg  int64  `json:"called.recvmsg"`
	MsgsReceived   int64  `json:"msgs.received"`
}

func newMessageFromJSON(b []byte) (*messageStat, error) {
	var pstat messageStat
	err := json.Unmarshal(b, &pstat)
	if err != nil {
		return nil, fmt.Errorf("error message stat `%v`: %v", string(b), err)
	}
	return &pstat, nil
}

func (m *messageStat) toPoints() []*point {
	points := make([]*point, 3)

	points[0] = &point{
		Name:        "imudp_called_recvmmsg",
		Type:        counter,
		Value:       m.CalledRecvmmsg,
		Description: "number of recvmmsg() OS calls done",
		LabelName:   "name",
		LabelValue:  m.Name,
	}
	points[1] = &point{
		Name:        "imudp_called_recvmsg",
		Type:        counter,
		Value:       m.CalledRecvmsg,
		Description: "number of recvmsg() OS calls done",
		LabelName:   "name",
		LabelValue:  m.Name,
	}
	points[2] = &point{
		Name:        "imudp_msgs_received",
		Type:        counter,
		Value:       m.MsgsReceived,
		Description: "number of actual messages received",
		LabelName:   "name",
		LabelValue:  m.Name,
	}

	return points
}
