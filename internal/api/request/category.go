package request

type CategoryDTO struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
	Sort string `json:"sort"`
	Cate string `json:"type"`
}

type CategoryPageQueryDTO struct {
	Name     string `form:"name"`     // 分页查询的Name
	Page     int    `form:"page"`     // 分页查询的页数
	PageSize int    `form:"pageSize"` // 分页查询的页容量
	Cate     int    `form:"type"`     // 分页类型: 1 为菜平分类, 2 为套餐分类
}
