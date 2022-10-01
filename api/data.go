package api

import "time"

var magicTable *Pocket

type Pocket struct {
}

type Coffee struct {
	Name              string
	ManufacturingTime time.Time
}
