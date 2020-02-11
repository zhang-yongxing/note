package sort

type AddSortForm struct {
	SortName string `json:"sort_name" binding:"required"`
}

type ALterSortForm struct {
	SortName string `json:"sort_name" binding:"required"`
}