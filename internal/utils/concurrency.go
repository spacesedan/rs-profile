package utils

import (
	"github.com/spacesedan/profile-tracker/internal/models"
)

type SingleTask interface {
	*models.TaskSingleAsset | *models.TaskSingleCollection
}

func Worker[S SingleTask](chA, chB chan S) {
	task := <-chA
	chB <- task
}
