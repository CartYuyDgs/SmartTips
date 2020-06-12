package logic

type SchoolInfo struct {
	SchoolId   int    `json:"school_id"`
	Provice    string `json:"provice"`
	City       string `json:"city"`
	SchoolType int    `json:"school_type"`
	SchoolName string `json:"school_name"`
}
