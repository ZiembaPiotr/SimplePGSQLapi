package entities

type Club struct {
	ClubName        string `json:"club_name"`
	StadiumName     string `json:"stadium_name"`
	StadiumCapacity int    `json:"stadium_capacity"`
}
