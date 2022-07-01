package explorer

type ResponseMove struct {
	Uci           string `json:"uci"`
	San           string `json:"san"`
	AverageRating int    `json:"averageRating"`
	White         int    `json:"white"`
	Draws         int    `json:"draws"`
	Black         int    `json:"black"`
}

type Response struct {
	White int            `json:"white"`
	Draws int            `json:"draws"`
	Black int            `json:"black"`
	Moves []ResponseMove `json:"moves"`
}
