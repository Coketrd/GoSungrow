package getPointInfo

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/devService/getPointInfo"
const Disabled = false

type RequestData struct {
	// DeviceType valueTypes.String `json:"device_type" required:"true"`
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	DeviceTypeList []struct {
		DeviceName valueTypes.String  `json:"device_name"`
		DeviceType valueTypes.Integer `json:"device_type"`
	} `json:"deviceTypeList"`
	DisplayModeList []struct {
		CodeName  valueTypes.String `json:"code_name"`
		PointType valueTypes.String `json:"point_type"`
	} `json:"displayModeList"`
	PointCalTypeList []struct {
		CodeName   valueTypes.String `json:"code_name"`
		CodeValue  valueTypes.String `json:"code_value"`
		CodeValue2 interface{}       `json:"code_value2"`
	} `json:"pointCalTypeList"`
	PointTypeList []struct {
		CodeName  valueTypes.String `json:"code_name"`
		PointType valueTypes.String `json:"point_type"`
	} `json:"pointTypeList"`
	PolymerizationModeList []struct {
		CodeName  valueTypes.String `json:"code_name"`
		PointType valueTypes.String `json:"point_type"`
	} `json:"polymerizationModeList"`
	PowerPointManage interface{} `json:"powerPointManage"`
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		pkg := apiReflect.GetName("", *e)
		dt := valueTypes.NewDateTime(valueTypes.Now)
		name := pkg	// + "." + e.Request.PsId.String()
		entries.StructToPoints(e.Response.ResultData, name, "system", dt)
	}

	return entries
}
