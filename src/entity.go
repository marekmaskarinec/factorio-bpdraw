package bpdraw

import (
	"image/color"
)

// There structs implement the object types from the blueprint string
// https://wiki.factorio.com/Blueprint_string_format

type PositionObject struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type ConnectionDataObject struct {
	EntityId  int `json:"entity_id"`
	CircuitId int `json:"circuit_id"`
}

type ConnectionPoint struct {
	Red   []ConnectionDataObject `json:"red"`
	Green []ConnectionDataObject `json:"green"`
}

type Connection struct {
	One ConnectionPoint `json:"1"`
	Two ConnectionPoint `json:"2"`
}

type ItemRequestObject struct {
	Key   string `json:"key"`
	Value int    `json:"value"`
}

type ItemFilterObject struct {
	Name  string `json:"name"`
	Index int    `json:"index"`
}

type InventoryObject struct {
	Filters []ItemFilterObject `json:"filters"`
	Bar     int                `json:"bar"`
}

type LogisticsFilterObject struct {
	Name  string `json:"name"`
	Index int    `json:"index"`
	Count int    `json:"count"`
}

type SpeakerParameterObject struct {
	PlaybackVolume   float64 `json:"playback_volume"`
	PlaybackGlobally bool    `json:"playback_globally"`
	AllowPolyphony   bool    `json:"allow_polyphony"`
}

type SignalId struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type SpeakerAlertParameterObject struct {
	ShowAlert    bool     `json:"show_alert"`
	ShowOnMap    bool     `json:"show_on_map"`
	IconSignalId SignalId `json:"icon_signal_id"`
	AlertMessage string   `json:"alert_message"`
}

type Entity struct {
	Number      int            `json:"entity_number"`
	Name        string         `json:"name"`
	Position    PositionObject `json:"position"`
	Direction   int            `json:"direction"`
	Orientation float64        `json:"orientation"`
	Connections Connection     `json:"connections"`
	//TODO: controlBehaviour
	Items              ItemRequestObject           `json:"items"`
	Recipe             string                      `json:"recipe"`
	Bar                int                         `json:"bar"`
	Inventory          InventoryObject             `json:"inventory"`
	Type               string                      `json:"type"`
	InputPriority      string                      `json:"input_priority"`
	OutputPriority     string                      `json:"output_priority"`
	Filter             string                      `json:"filter"`
	Filters            []ItemFilterObject          `json:"filters"`
	FilterMode         string                      `json:"filter_mode"`
	OverrideStackSize  uint8                       `json:"override_stack_size"`
	DropPosition       PositionObject              `json:"drop_position"`
	PickupPosition     PositionObject              `json:"pickup_position"`
	RequestFilters     []LogisticsFilterObject     `json:"request_filters"`
	RequestFromBuffers bool                        `json:"request_from_buffers"`
	Parameters         SpeakerParameterObject      `json:"parameters"`
	AlertParameters    SpeakerAlertParameterObject `json:"alert_parameters"`
	AutoLaunch         bool                        `json:"auto_launch"`
	Color              color.RGBA                  `json:"color"`
	Station            string                      `json:"station"`
}

type Tile struct {
	Name     string         `json:"name"`
	Position PositionObject `json:"position"`
}

type Icon struct {
	Index  int      `json:"index"`
	Signal SignalId `json:"signal"`
}

type WaitCondition struct {
	Type        string `json:"type"`
	CompareType string `json:"compare_type"`
	Ticks       uint   `json:"ticks"`
	//TODO: condition CircuitCondition
}

type ScheduleRecord struct {
	Station        string          `json:"entity_number"`
	WaitConditions []WaitCondition `json:"entity_number"`
}

type ScheduleObject struct {
	Schedule    []ScheduleRecord `json:"entity_number"`
	Locomotives int              `json:"entity_number"`
}

type Blueprint struct {
	Item       string           `json:"item"`
	Label      string           `json:"label"`
	LabelColor color.RGBA       `json:"label_color"`
	Entities   []Entity         `json:"entities"`
	Tiles      []Tile           `json:"tiles"`
	Icons      []Icon           `json:"icons"`
	Schedules  []ScheduleObject `json:"schedules"`
	Version    int64            `json:"version"`
}
