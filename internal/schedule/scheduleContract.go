package schedule

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

type Handler interface {
	ScheduledFeed(ctx *gin.Context)
}

type Service interface {
	AddSchedule(ctx context.Context, request *ScheduledFeedRequest) (*ScheduledFeedResponse, error)
}

type Repository interface {
	InsertSchedule(ctx context.Context, s *Schedule) (*Schedule, error)
	InsertFeedingSchedule(ctx context.Context, s *FeedingSchedule) (*FeedingSchedule, error)
	GetScheduleByDayAndTime(ctx context.Context, day string, time time.Time) (*Schedule, error)
}