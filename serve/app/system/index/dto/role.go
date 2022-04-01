package dto

type RoleQuery struct {
	Page   int   `p:"page"`
	Limit  int   `p:"limit"`
	Type   []int `p:"type"`
	Status int   `p:"status"`
}
