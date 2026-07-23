package models

import "time"

type Role struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type User struct {
	ID            uint      `json:"id"`
	IDRole        uint      `json:"id_role"`
	Password      string    `json:"-"` // tidak muncul di JSON
	Email         string    `json:"email"`
	HpNumber      *string    `json:"hp_number"`
	StatusAccount string    `json:"status_account"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Role          Role      `json:"role,omitempty"`
	Profile       Profile   `json:"profile,omitempty"`
}

type Profile struct {
	ID         uint       `json:"id"`
	IDUser     uint       `json:"id_user"`
	Name       string     `json:"name"`
	Gender     *string    `json:"gender,omitempty"`
	Picture    *string    `json:"picture,omitempty"`
	PlaceBirth *string    `json:"place_birth,omitempty"`
	DateBirth  *time.Time `json:"date_birth,omitempty"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

type Session struct {
	ID        uint      `json:"id"`
	IDUser    uint      `json:"id_user"`
	Token     string    `json:"token"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UsersLog struct {
	IDUser         uint      `json:"id_user"`
	IDSession      *uint     `json:"id_session,omitempty"`
	ActivityDetail string    `json:"activity_detail"`
	IPAddress      *string   `json:"ip_address,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
}
