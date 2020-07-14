package domain

// Passenger Package
type Passenger struct {
	ID       uint64 `json: "passenger_id"`
	FullName string `json:"full_name"`
	Age      uint8  `json:age`
	Gender   string `json:age`
	TicketNo string `json:ticket_no`
}
