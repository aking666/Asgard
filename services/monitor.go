package services

import (
	"github.com/dalonghahaha/avenger/components/db"
	"github.com/dalonghahaha/avenger/components/logger"

	"Asgard/models"
)

type MonitorService struct {
}

func NewMonitorService() *MonitorService {
	return &MonitorService{}
}

func (s *MonitorService) GetMonitorPageList(where map[string]interface{}, page int, pageSize int) (list []models.Monitor, count int) {
	err := models.PageList(&models.Monitor{}, where, page, pageSize, "created_at desc", &list, &count)
	if err != nil {
		logger.Error("GetMonitorPageList Error:", err)
		return nil, 0
	}
	return
}

func (s *MonitorService) GetAgentMonitor(id int, size int) (list []models.Monitor) {
	where := map[string]interface{}{
		"type":       models.TYPE_AGENT,
		"related_id": id,
	}
	err := db.Get(models.DB_NAME).Where(where).Limit(size).Order("created_at desc").Find(&list).Error
	if err != nil {
		logger.Error("GetAppMonitor Error:", err)
		return nil
	}
	return
}

func (s *MonitorService) GetAppMonitor(id int, size int) (list []models.Monitor) {
	where := map[string]interface{}{
		"type":       models.TYPE_APP,
		"related_id": id,
	}
	err := db.Get(models.DB_NAME).Where(where).Limit(size).Order("created_at desc").Find(&list).Error
	if err != nil {
		logger.Error("GetAppMonitor Error:", err)
		return nil
	}
	return
}

func (s *MonitorService) GetJobMonitor(id int, size int) (list []models.Monitor) {
	where := map[string]interface{}{
		"type":       models.TYPE_JOB,
		"related_id": id,
	}
	err := db.Get(models.DB_NAME).Where(where).Limit(size).Order("created_at desc").Find(&list).Error
	if err != nil {
		logger.Error("GetJobMonitor Error:", err)
		return nil
	}
	return
}

func (s *MonitorService) CreateMonitor(monitor *models.Monitor) bool {
	err := models.Create(monitor)
	if err != nil {
		logger.Error("CreateMonitor Error:", err)
		return false
	}
	return true
}