package presentation

type Response struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type MetaPagination struct {
	SortBy  string `json:"sortby,omitempty"`
	OrderBy string `json:"orderby,omitempty"`
	PerPage int    `json:"perpage,omitempty"`
	Page    int    `json:"page,omitempty"`
}
