package academic

import "time"

type Event struct {
	Event_id string `gorm:"type:uuid; default:uuid_generate_v4(); primaryKey" json:"event_id"`
	// student id
	// subject id
	Title       string    `gorm:"type:varchar(100); not null" json:"title"`
	Type        string    `gorm:"type:varchar(50); not null" json:"type"`
	Description string    `gorm:"type:text; not null" json:"description"`
	Duration    int       `gorm:"type:int; not null" json:"duration"`
	Event_date  time.Time `gorm:"type:timestamp; not null" json:"event_date"`
}
