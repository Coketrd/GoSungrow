package psHourPointsValue

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/powerStationService/psHourPointsValue"
const Disabled = false

type RequestData struct {
	PsId valueTypes.PsId `json:"ps_id" required:"true"`
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	P24001List []valueTypes.Float `json:"p24001List" PointId:"p24001" PointUnitFrom:"P24001Unit" PointArrayFlatten:"false" DataTable:"true" DataTableMerge:"true" DataTableShowIndex:"true"`
	P24001Unit valueTypes.String  `json:"p24001_unit" PointIgnore:"true"`
	P24004List []valueTypes.Float `json:"p24004List" PointId:"p24004" PointUnitFrom:"P24004Unit" PointArrayFlatten:"false" DataTable:"true" DataTableMerge:"true" DataTableShowIndex:"true"`
	P24004Unit valueTypes.String  `json:"p24004_unit" PointIgnore:"true"`
	P83002List []valueTypes.Float `json:"p83002List" PointId:"p83002" PointUnitFrom:"P83002Unit" PointArrayFlatten:"false" DataTable:"true" DataTableMerge:"true" DataTableShowIndex:"true"`
	P83002Unit valueTypes.String  `json:"p83002_unit" PointIgnore:"true"`
	P83012List []valueTypes.Float `json:"p83012List" PointId:"p83012" PointUnitFrom:"P83012Unit" PointArrayFlatten:"false" DataTable:"true" DataTableMerge:"true" DataTableShowIndex:"true"`
	P83012Unit valueTypes.String  `json:"p83012_unit" PointIgnore:"true"`
	P83022List []valueTypes.Float `json:"p83022List" PointId:"p83022" PointUnitFrom:"P83022Unit" PointArrayFlatten:"false" DataTable:"true" DataTableMerge:"true" DataTableShowIndex:"true"`
	P83022Unit valueTypes.String  `json:"p83022_unit" PointIgnore:"true"`
	P83033List []valueTypes.Float `json:"p83033List" PointId:"p83033" PointUnitFrom:"P83033Unit" PointArrayFlatten:"false" DataTable:"true" DataTableMerge:"true" DataTableShowIndex:"true"`
	P83033Unit valueTypes.String  `json:"p83033_unit" PointIgnore:"true"`
	P83039List []valueTypes.Float `json:"p83039List" PointId:"p83039" PointUnitFrom:"P83039Unit" PointArrayFlatten:"false" DataTable:"true" DataTableMerge:"true" DataTableShowIndex:"true"`
	P83039Unit valueTypes.String  `json:"p83039_unit" PointIgnore:"true"`
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		// pkg := apiReflect.GetName("", *e)
		// dt := valueTypes.NewDateTime(valueTypes.Now)
		// name := pkg + "." + e.Request.PsId.String()
		entries.StructToDataMap(*e, e.Request.PsId.String(), apiReflect.NewEndPointPath(e.Request.PsId.String()))
	}

	return entries
}
