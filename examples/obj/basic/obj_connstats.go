package basic

import (
	. "github.com/zubairhamed/go-lwm2m/api"
	"github.com/zubairhamed/go-lwm2m/core"
	"github.com/zubairhamed/go-lwm2m/objects/oma"
	"github.com/zubairhamed/go-lwm2m/core/response"
)

type ConnectivityStatistics struct {
	Model ObjectModel
	Data  *core.ObjectsData
}

func (o *ConnectivityStatistics) OnExecute(instanceId int, resourceId int, req Request) Response {
	return response.Unauthorized()
}

func (o *ConnectivityStatistics) OnCreate(instanceId int, resourceId int, req Request) Response {
	return response.Unauthorized()
}

func (o *ConnectivityStatistics) OnDelete(instanceId int, req Request) Response {
	return response.Unauthorized()
}

func (o *ConnectivityStatistics) OnRead(instanceId int, resourceId int, req Request) Response {
	return response.Unauthorized()
}

func (o *ConnectivityStatistics) OnWrite(instanceId int, resourceId int, req Request) Response {
	return response.Unauthorized()
}

func NewExampleConnectivityStatisticsObject(reg Registry) *ConnectivityStatistics {
	data := &core.ObjectsData{
		Data: make(map[string]interface{}),
	}

	return &ConnectivityStatistics{
		Model: reg.GetModel(oma.OBJECT_LWM2M_CONNECTIVITY_STATISTICS),
		Data:  data,
	}
}
