package api

type Dept struct {
	DeptID uint   `json:"dept_id"`
	Name   string `json:"name"`
	Desc   string `json:"desc"`
	Leader *User  `json:"leader"`
}
