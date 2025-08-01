package models

import (
    "time"
)

type Session struct {
    ID           string    `bson:"_id" json:"id"`
    Provider     string    `bson:"provider" json:"provider"`
    PhoneNumber  string    `bson:"phone_number" json:"phone_number"`
    Cookies      []byte    `bson:"cookies" json:"cookies"`
    DomSnapshot  []byte    `bson:"dom_snapshot" json:"dom_snapshot"`
    CurrentURL   string    `bson:"current_url" json:"current_url"`
    CreatedAt    time.Time `bson:"created_at" json:"created_at"`
    Status       string    `bson:"status" json:"status"` // e.g., "otp_sent", "logged_in"
}
