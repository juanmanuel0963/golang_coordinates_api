package models

type UserCoordinates struct {
	Id        int32   `json:"user_id"`
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude,string"`
	Longitude float64 `json:"longitude,string"`
	Distance  float64 `json:"distance,string"`
}

type ByUserID []UserCoordinates

func (a ByUserID) Len() int           { return len(a) }
func (a ByUserID) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByUserID) Less(i, j int) bool { return a[i].Id < a[j].Id }
