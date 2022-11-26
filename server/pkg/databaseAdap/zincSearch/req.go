package zincsearch

// DocumentcreatorReq
type DocumentcreatorReq struct {
	Index   string      `json:"index"`
	Records interface{} `json:"records"`
}

// DocumentFinderReq finds a document
type DocumentFinderReq struct {
	SearchType string                      `json:"search_type"`
	Query      SearchDocumentsRequestQuery `json:"query"`
	SortFields []string                    `json:"sort_fields"`
	From       int                         `json:"from"`
	MaxResults int                         `json:"max_results"`
	Source     map[string]interface{}      `json:"_source"`
}

// SearchDocumentsRequestQuery
type SearchDocumentsRequestQuery struct {
	Term      string `json:"term"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}
