package models

type Query struct {
	Query_ID    int32  `json:"query_id" form:"query_id"`
	Query_Name  string `json:"query_name" form:"query_name"`
	State       int32  `json:"state" form:"state"`
	Remark      string `json:"remark" form:"remark"`
	Start_Time  int64  `json:"start_time" form:"start_time"`
	End_Time    int64  `json:"end_time" form:"end_time"`
	End_Method  int32  `json:"end_method" form:"end_method"` // 0：问卷到期 1: 手动终止 ...
	Background  string `json:"background" form:"background"`
	Creator     int32  `json:"creator" form:"creator"`
	Create_Time int64  `json:"create_time" form:"create_time"`
}
