package form

import (
	"go-netdisk/pkg/db/models"
	"mime/multipart"
)

type PageParam struct {
	PUUID           string `form:"puuid"`
	Name            string `form:"name"`
	Page            int    `form:"page"`
	PageSize        int    `form:"pageSize"`
	OrderCreateTime string `form:"orderCreateTime"`
}

type BaseQueryParam struct {
	UUID string `form:"uuid" binding:"required"`
}

type BasePostParam struct {
	UserUUID string `form:"userUuid" binding:"required"`
	PUUID    string `form:"puuid" binding:"required"`
}

type UploadParam struct {
	BasePostParam
	File *multipart.FileHeader `form:"file" binding:"required"`
}

type CreateDirParam struct {
	BasePostParam
	Name string `form:"name" binding:"required"`
}

type SubDirDetailMatter struct {
	*models.Matter
	Parent *models.Matter `json:"parent"`
}

type RootDirDetailMatter struct {
	*models.Matter
	Parent *string `json:"parent"`
}
