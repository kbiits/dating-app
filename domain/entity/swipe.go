package entity

import "time"

type SwipeDirection string

var (
	SwipeDirectionLeft  SwipeDirection = "dislike"
	SwipeDirectionRight SwipeDirection = "like"
)

type Swipe struct {
	ID        string         `db:"id"`
	SwiperID  string         `db:"swiper_id"`
	SwipedID  string         `db:"swiped_id"`
	SwipeType SwipeDirection `db:"swipe_type"`
	SwipeDate time.Time      `db:"swipe_date"`
	CreatedAt time.Time      `db:"created_at"`
}
