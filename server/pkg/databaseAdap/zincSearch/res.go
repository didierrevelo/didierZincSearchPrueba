package zincsearch

// DocumentCreatorRes
type DocumentCreatorRes struct {
	Message     string `json:"message"`
	RecordCount int    `json:"record_count"`
}

// DocumentFinderRes
type DocumentFinderRes struct {
	Hits struct {
		Hits []struct {
			ID        string                 `json:"_id"`
			Timestamp string                 `json:"@timestamp"`
			Score     float64                `json:"_score"`
			Source    map[string]interface{} `json:"_source"`
		} `json:"hits"`
		Total struct {
			Value int `json:"value"`
		} `json:"total"`
		MaxScore float64 `json:"max_score"`
	} `json:"hits"`
	TimedOut bool    `json:"timed_out"`
	Took     float64 `json:"took"`
}

// ErrorRes
type ErrorRes struct {
	ErrorMessage string `json:"error"`
}
