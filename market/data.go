package market

type Quota []float64

type TickData struct {
	Bids      []Quota `json:"bids"`
	Asks      []Quota `json:"asks"`
	Version   int     `json:"version"`
	Timestamp int     `json:"ts"`
}
