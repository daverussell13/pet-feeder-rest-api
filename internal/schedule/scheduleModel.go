package schedule

type ScheduledFeedRequest struct {
	DeviceID   string `json:"device_id" binding:"required,uuid4"`
	DayOfWeek  string `json:"day_of_week" binding:"required,oneof=Sunday Monday Tuesday Wednesday Thursday Friday Saturday"`
	FeedTime   string `json:"feed_time" binding:"required,timeFormat"`
	FeedAmount int    `json:"feed_amount" binding:"required,min=1,max=7"`
}

type ScheduledFeedResponse struct {
	ScheduleID string `json:"schedule_id" binding:"required"`
	CreatedAt  string `json:"created_at" binding:"required"`
}