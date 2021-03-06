package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"Asgard/constants"
	"Asgard/providers"
	"Asgard/web/utils"
)

type ArchiveController struct {
}

func NewArchiveController() *ArchiveController {
	return &ArchiveController{}
}

func (c *ArchiveController) App(ctx *gin.Context) {
	page := utils.DefaultInt(ctx, "page", 1)
	app := utils.GetApp(ctx)
	where := map[string]interface{}{
		"type":       constants.TYPE_APP,
		"related_id": app.ID,
	}
	archiveList, total := providers.ArchiveService.GetArchivePageList(where, page, constants.WEB_LIST_PAGE_SIZE)
	if archiveList == nil {
		utils.APIError(ctx, "获取归档列表失败")
	}
	list := []gin.H{}
	for _, archive := range archiveList {
		list = append(list, utils.ArchiveFormat(&archive))
	}
	mpurl := fmt.Sprintf("/archive/app?id=%d", app.ID)
	utils.Render(ctx, "archive/list", gin.H{
		"Subtitle":   "应用归档列表——" + app.Name,
		"BackUrl":    utils.GetReferer(ctx),
		"List":       list,
		"Total":      total,
		"Pagination": utils.PagerHtml(total, page, mpurl),
	})
}

func (c *ArchiveController) Job(ctx *gin.Context) {
	page := utils.DefaultInt(ctx, "page", 1)
	job := utils.GetJob(ctx)
	where := map[string]interface{}{
		"type":       constants.TYPE_JOB,
		"related_id": job.ID,
	}
	archiveList, total := providers.ArchiveService.GetArchivePageList(where, page, constants.WEB_LIST_PAGE_SIZE)
	if archiveList == nil {
		utils.APIError(ctx, "获取归档列表失败")
	}
	list := []gin.H{}
	for _, archive := range archiveList {
		list = append(list, utils.ArchiveFormat(&archive))
	}
	mpurl := fmt.Sprintf("/archive/job?id=%d", job.ID)
	utils.Render(ctx, "archive/list", gin.H{
		"Subtitle":   "计划任务归档列表——" + job.Name,
		"BackUrl":    utils.GetReferer(ctx),
		"List":       list,
		"Total":      total,
		"Pagination": utils.PagerHtml(total, page, mpurl),
	})
}

func (c *ArchiveController) Timing(ctx *gin.Context) {
	page := utils.DefaultInt(ctx, "page", 1)
	timing := utils.GetTiming(ctx)
	where := map[string]interface{}{
		"type":       constants.TYPE_TIMING,
		"related_id": timing.ID,
	}
	archiveList, total := providers.ArchiveService.GetArchivePageList(where, page, constants.WEB_LIST_PAGE_SIZE)
	if archiveList == nil {
		utils.APIError(ctx, "获取归档列表失败")
	}
	list := []gin.H{}
	for _, archive := range archiveList {
		list = append(list, utils.ArchiveFormat(&archive))
	}
	mpurl := fmt.Sprintf("/archive/timing?id=%d", timing.ID)
	utils.Render(ctx, "archive/list", gin.H{
		"Subtitle":   "定时任务归档列表——" + timing.Name,
		"BackUrl":    utils.GetReferer(ctx),
		"List":       list,
		"Total":      total,
		"Pagination": utils.PagerHtml(total, page, mpurl),
	})
}
