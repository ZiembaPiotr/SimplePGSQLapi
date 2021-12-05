package entities

type Player struct {
	Name         string `json:"name"`
	Age          int    `json:"age"`
	JerseyNumber int    `json:"jerseyNumber"`
	Club         string `json:"club"`
}
