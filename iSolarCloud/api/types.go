package api

import (
	"GoSungrow/iSolarCloud/api/output"
)


func (e *Web) ApiGetEndPointResultTable(endpoint EndPoint) output.Table {
	var table output.Table
	// for range Only.Once {
	// 	// table = output.NewTable()
	// 	data := endpoint.GetEndPointData()
	// 	table = data.CreateEndPointResultTable()
	// 	table.SetTitle(fmt.Sprintf("Data Request %s.%s", endpoint.GetArea(), endpoint.GetName()))
	// 	table.SetFilePrefix(fmt.Sprintf("%s_%s", endpoint.GetArea(), endpoint.GetName()))
	// 	table.SetJson([]byte(endpoint.GetJsonData(false)))
	// 	table.SetRaw([]byte(endpoint.GetJsonData(true)))
	//
	// 	// table.InitGraph(output.GraphRequest {
	// 	// 	Title:        "",
	// 	// 	TimeColumn:   output.SetInteger(1),
	// 	// 	SearchColumn: output.SetInteger(2),
	// 	// 	NameColumn:   output.SetInteger(4),
	// 	// 	ValueColumn:  output.SetInteger(5),
	// 	// 	UnitsColumn:  output.SetInteger(6),
	// 	// 	SearchString: output.SetString(""),
	// 	// 	MinLeftAxis:  output.SetFloat(0),
	// 	// 	MaxLeftAxis:  output.SetFloat(0),
	// 	// })
	// }
	return table
}

func (e *Web) ApiGetEndPointDataTables(endpoint EndPoint) output.Tables {
	table := output.NewTables()
	// for range Only.Once {
	//
	// 	var tp apiReflect.DataStructures
	// 	tp.Debug = true
	// 	tp.ShowEmpty = true
	// 	var Ref apiReflect.Reflect
	// 	Ref.SetByFieldName(endpoint.ResponseRef(), endpoint.ResponseRef(), "")
	// 	// tp.FindDataTables(Ref, Ref, apiReflect.EndPointPath{}, false)
	//
	// 	// data := endpoint.GetEndPointData()
	// 	//
	// 	// table = data.CreateTable()
	// 	// table.SetTitle(fmt.Sprintf("Data Request %s.%s", endpoint.GetArea(), endpoint.GetName()))
	// 	// table.SetFilePrefix(fmt.Sprintf("%s_%s", endpoint.GetArea(), endpoint.GetName()))
	// 	// table.SetJson([]byte(endpoint.GetJsonData(false)))
	// 	// table.SetRaw([]byte(endpoint.GetJsonData(true)))
	// 	//
	// 	// table.InitGraph(output.GraphRequest {
	// 	// 	Title:        "",
	// 	// 	TimeColumn:   output.SetInteger(1),
	// 	// 	SearchColumn: output.SetInteger(2),
	// 	// 	NameColumn:   output.SetInteger(4),
	// 	// 	ValueColumn:  output.SetInteger(5),
	// 	// 	UnitsColumn:  output.SetInteger(6),
	// 	// 	SearchString: output.SetString(""),
	// 	// 	MinLeftAxis:  output.SetFloat(0),
	// 	// 	MaxLeftAxis:  output.SetFloat(0),
	// 	// })
	// }
	return table
}