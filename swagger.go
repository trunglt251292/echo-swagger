package echoswagger

type (
	// DescriptionAPI ...
	DescriptionAPI struct {
		Summary     string `json:"summary,omitempty"`
		Description string `json:"description,omitempty"`
	}

	// Params ...
	Params struct {
		Param    interface{} `json:"param"`
		Name     string      `json:"name"`
		Required bool        `json:"required"`
		Desc     string      `json:"desc"`
		Type     interface{} `json:"type,omitempty"`
	}

	// Test ...
	Test struct {
		ID     string `json:"id"`
		Total  int32  `json:"total"`
		Status bool   `json:"status"`
	}
	// ResponseJSON ...
	ResponseJSON struct {
		Code   int         `json:"code"`
		Desc   string      `json:"desc"`
		Schema interface{} `json:"schema,omitempty"`
		Header interface{} `json:"header,omitempty"`
	}
)

// NewDescription ...
func NewDescription(summary string, desc string) *DescriptionAPI {
	return &DescriptionAPI{
		Summary:     summary,
		Description: desc,
	}
}

// PageTokenQuery ...
func PageTokenQuery() (p interface{}, name, desc string, required bool) {
	return "string", "pageToken", "Page token for pagination", false
}

// KeywordQuery ...
func KeywordQuery() (p interface{}, name, desc string, required bool) {
	return "string", "keyword", "Search keyword", false
}

// SortQuery ...
func SortQuery(description string) (p interface{}, name, desc string, required bool) {
	return "string", "sort", description, false
}

// PageQuery ...
func PageQuery() (p interface{}, name, desc string, required bool) {
	return "int", "page", "", false
}

// ActiveStatusQuery ...
func ActiveStatusQuery() (p interface{}, name, desc string, required bool) {
	return "bool", "active", "true, false", false
}
