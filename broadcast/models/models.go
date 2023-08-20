package models

type Broadcast struct {
	BroadcastId string `json:"broadcast_id"`
	UserId      string `json:"user_id"`
	// request object from common schemas
	BroadcastRadius float64 `json:"broadcast_radius"`
}
