package response

type ResponseData struct {
	Data interface{} `json:"data"`
}

type ListRecords struct {
	Current int         `json:"current"`
	Size    int         `json:"size"`
	Total   int64       `json:"total"`
	Records interface{} `json:"records"`
}

type ListTreeRecords struct {
	Records interface{} `json:"records"`
	Total   int64       `json:"total"`
	Size    int         `json:"size"`
	Current int         `json:"current"`
}

type Dict struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

type DictTree struct {
	Value    string     `json:"value"`
	Label    string     `json:"label"`
	Children []DictTree `json:"children"`
}
