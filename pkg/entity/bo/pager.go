package bo

// Pager defines page structure
type Pager struct {
	PageNo     *int        `json:"pageNo"`
	PageSize   *int        `json:"pageSize"`
	TotalCount *int        `json:"totalCount"`
	PageCount  *int        `json:"pageCount"`
	Data       interface{} `json:"data"`
}
