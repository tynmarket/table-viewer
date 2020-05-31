package model

// WhereQuery is
type WhereQuery struct {
	Column   string `json:"column" binding:"required"`
	Operator string `json:"operator" binding:"required"`
	Value    string `json:"val"`
}
