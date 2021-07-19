package jsonstrs

type Resource struct {
	Objects map[string]Object `json:"objects"`
}

type Object struct {
	Hash string `json:"hash"`
	Size int64  `json:"size"`
}
