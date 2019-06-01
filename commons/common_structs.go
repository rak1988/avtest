package commons

// MarksStudent struct
type MarksStudent struct {
	UID     int64  `json:"uid"`
	Name    string `json:"name"`
	Maths   int    `json:"maths"`
	Physics int    `json:"physics"`
	Chem    int    `json:"chemistry"`
}
