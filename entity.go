package main

import (
	"image/color"
)

type Position struct {
	x float64
	y float64
}

type ConnectionDataObject struct {
	entityId int
	circuitId int
}

type ConnectionPoint struct {
	red []ConnectionDataObject
	green []ConnectionDataObject
}

type Connection struct {
	one ConnectionPoint
	two ConnectionPoint
}

type ItemRequestObject struct {
	key string
	value int
}

type ItemFilterObject struct {
	name string
	index int
}

type Inventory struct {
	filters []ItemFilterObject
	bar int
}

type LogisticsFilterObject struct {
	name string
	index int
	count int
}

type SpeakerParameterObject struct {
	playbackVolume float64
	playbackGlobally bool
	allowPolyphony bool
}

type SignalId struct {
	name string
	t string
}

type SpeakerAlertParameterObject struct {
	showAlert bool
	showOnMap bool
	iconSignalId SignalId
	alertMessage string
}

type Entity struct {
	number int
	name string
	position Position
	direction int
	orientation float64
	connections Connection
	//controlBehaviour
	items ItemRequestObject
	recipe string
	bar int
	inventory Inventory
	t string
	inputPriority string
	outputPriority string
	filter string
	filters []ItemFilterObject
	filterMode string
	overrideStackSize uint8
	dropPosition Position
	pickupPosition Position
	requestFilters []LogisticsFilterObject
	requestFromBuffers bool
	parameters SpeakerParameterObject
	alertParameters SpeakerAlertParameterObject
	autoLaunch bool
	color color.RGBA
	station string
}

type Tile struct {
	name string
	position Position
}

type Icon struct {
	index int
	signal SignalId
}

type WaitCondition struct {
	t string
	compareType string
	ticks uint
	//condition CircuitCondition
}

type ScheduleRecord struct {
	station string
	waitConditions []WaitCondition
}

type ScheduleObject struct {
	schedule []ScheduleRecord
	locomotives int
}

type Blueprint struct {
	item string
	label string
	labelColor color.RGBA
	entities []Entity
	tiles []Tile
	icons []Icon
	schedules []ScheduleObject
	version int64
}
