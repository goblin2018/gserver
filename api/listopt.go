package api

type ListOpt struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Order  string `json:"order"`
	All    bool   `json:"all"`
}

type ListRes struct {
	Total int64       `json:"total"`
	Items interface{} `json:"items"`
}
