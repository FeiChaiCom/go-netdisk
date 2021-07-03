package form

import (
	"github.com/gaomugong/go-netdisk/models/db"
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
	*db.Matter
	Parent *db.Matter `json:"parent"`
}

type RootDirDetailMatter struct {
	*db.Matter
	Parent *string `json:"parent"`
}
