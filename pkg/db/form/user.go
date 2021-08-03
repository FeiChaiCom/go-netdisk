package form

type UserParam struct {
	Page            int    `form:"page"`
	PageSize        int    `form:"pageSize"`
	OrderCreateTime string `form:"orderCreateTime"`
}
