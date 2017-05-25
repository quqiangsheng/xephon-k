package common

import (
	"encoding/json"
	"fmt"
	"time"
)

// check interface
var _ Series = (*SeriesMeta)(nil)
var _ Series = (*RawSeries)(nil)
var _ Series = (*IntSeries)(nil)
var _ Series = (*DoubleSeries)(nil)

const (
	_ = iota
	TypeIntSeries
	TypeDoubleSeries
	TypeBoolSeries
	TypeStringSeries
)

type Series interface {
	Hashable
	GetSeriesType() int
	// NOTE: series decoded from JSON has 0 as SeriesID, so the implementation would recalculate the Hash
	GetSeriesID() SeriesID
}

type SeriesMeta struct {
	id         SeriesID
	Name       string            `json:"name"`
	Tags       map[string]string `json:"tags"`
	SeriesType int               `json:"type"`
	Precision  time.Duration     `json:"precision"`
}

type RawSeries struct {
	id         SeriesID
	Name       string            `json:"name"`
	Tags       map[string]string `json:"tags"`
	SeriesType int               `json:"type"`
	Precision  time.Duration     `json:"precision"`
	Points     json.RawMessage   `json:"points"`
}

type IntSeries struct {
	id         SeriesID
	Name       string            `json:"name"`
	Tags       map[string]string `json:"tags"`
	SeriesType int               `json:"type,omitempty"`
	Precision  time.Duration     `json:"precision,omitempty"`
	Points     []IntPoint        `json:"points"`
}

type DoubleSeries struct {
	id         SeriesID
	Name       string            `json:"name"`
	Tags       map[string]string `json:"tags"`
	SeriesType int               `json:"type,omitempty"`
	Precision  time.Duration     `json:"precision,omitempty"`
	Points     []DoublePoint     `json:"points"`
}

// TODO: int series of other precision, maybe we should add millisecond to the default function as well
func NewIntSeries(name string) *IntSeries {
	return &IntSeries{
		Name:       name,
		Tags:       make(map[string]string, 1),
		SeriesType: TypeIntSeries,
		Precision:  time.Millisecond,
	}
}

func NewDoubleSeries(name string) *DoubleSeries {
	return &DoubleSeries{
		Name:       name,
		Tags:       make(map[string]string, 1),
		SeriesType: TypeDoubleSeries,
		Precision:  time.Millisecond,
	}
}

func (series *SeriesMeta) GetName() string {
	return series.Name
}

func (series *SeriesMeta) GetTags() map[string]string {
	return series.Tags
}

func (series *SeriesMeta) GetSeriesType() int {
	return series.SeriesType
}

func (series *SeriesMeta) GetSeriesID() SeriesID {
	if series.id == 0 {
		series.id = Hash(series)
	}
	return series.id
}

func (series *RawSeries) GetName() string {
	return series.Name
}

func (series *RawSeries) GetTags() map[string]string {
	return series.Tags
}

func (series *RawSeries) GetSeriesType() int {
	return series.SeriesType
}

func (series *RawSeries) GetSeriesID() SeriesID {
	if series.id == 0 {
		series.id = Hash(series)
	}
	return series.id
}

func (series *IntSeries) GetName() string {
	return series.Name
}

func (series *IntSeries) GetTags() map[string]string {
	return series.Tags
}

func (series *IntSeries) GetSeriesType() int {
	// TODO: do we still need the variable if we return constant, and should we check consistency for that
	return TypeIntSeries
}

func (series *IntSeries) GetSeriesID() SeriesID {
	if series.id == 0 {
		series.id = Hash(series)
	}
	return series.id
}

func (series *DoubleSeries) GetName() string {
	return series.Name
}

func (series *DoubleSeries) GetTags() map[string]string {
	return series.Tags
}

func (series *DoubleSeries) GetSeriesType() int {
	return TypeDoubleSeries
}

func (series *DoubleSeries) GetSeriesID() SeriesID {
	if series.id == 0 {
		series.id = Hash(series)
	}
	return series.id
}

func SeriesTypeString(seriesType int) string {
	switch seriesType {
	case TypeIntSeries:
		return "int"
	case TypeDoubleSeries:
		return "double"
	case TypeBoolSeries:
		return "bool"
	case TypeStringSeries:
		return "string"
	default:
		return fmt.Sprintf("unknown: %d", seriesType)
	}
}
