package localization

import (
	"errors"
	"fmt"
	"math"

	"github.com/google/uuid"
)

var ErrInvalidCourseID = errors.New("invaild location uuid")

const roundConst = 5

type LocalizationUUID struct {
	value string
}

func (luuid LocalizationUUID) String() string {
	return luuid.value
}

func NewUuid(value string) (LocalizationUUID, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return LocalizationUUID{}, fmt.Errorf("%w: %s", ErrInvalidCourseID, value)
	}

	return LocalizationUUID{
		value: v.String(),
	}, nil
}

type Localization struct {
	Id   LocalizationUUID `json:"id"`
	Lat  float64          `json:"lat"`
	Long float64          `json:"long"`
}

func NewLocalization(id string, lat, long float64) (*Localization, error) {

	locUuid, err := NewUuid(id)
	if err != nil {
		return &Localization{}, err
	}

	lat = math.Round(lat*math.Pow10(roundConst)) / math.Pow10(roundConst)
	long = math.Round(long*math.Pow10(roundConst)) / math.Pow10(roundConst)

	return &Localization{
		Id:   locUuid,
		Lat:  lat,
		Long: long,
	}, nil

}

type GroupAvaerageLocation struct {
	IdLocation string  `json:"id_location"`
	Lat        float64 `json:"lat"`
	Long       float64 `json:"long"`
}

func NewGroupAvaerageLocation(idLocation string, lat float64, long float64) *GroupAvaerageLocation {
	lat = math.Round(lat*math.Pow10(roundConst)) / math.Pow10(roundConst)
	long = math.Round(long*math.Pow10(roundConst)) / math.Pow10(roundConst)

	return &GroupAvaerageLocation{
		IdLocation: idLocation,
		Lat:        lat,
		Long:       long,
	}
}

func ErrorGroupAvaerageLocation() *GroupAvaerageLocation {
	return &GroupAvaerageLocation{}
}

func (loc Localization) DistanceToLoc(otherLoc Localization) float64 {
	distance := math.Sqrt(math.Pow(loc.Lat-otherLoc.Lat, 2) + math.Pow(loc.Long-otherLoc.Long, 2))
	return distance
}

func (loc Localization) DistanceToLocList(LocList []Localization) map[Localization]float64 {
	distances := make(map[Localization]float64)

	if len(LocList) == 0 {
		return distances
	}

	for _, otherLoc := range LocList {
		distances[otherLoc] = loc.DistanceToLoc(otherLoc)
	}

	return distances
}

func (loc Localization) GetGroupAvaerageLocation(LocList []Localization) *GroupAvaerageLocation {
	if len(LocList) == 0 {
		return NewGroupAvaerageLocation(loc.Id.String(), loc.Lat, loc.Long)
	}

	distances := loc.DistanceToLocList(LocList)
	locsOfGroup := []Localization{loc}

	for otherLoc, distance := range distances {
		if distance <= 1 {
			locsOfGroup = append(locsOfGroup, otherLoc)
		}
	}

	latMedia := 0.0
	longMedia := 0.0
	nLocInGroup := len(locsOfGroup)

	for _, locOfGroup := range locsOfGroup {
		latMedia = latMedia + locOfGroup.Lat/float64(nLocInGroup)
		longMedia = longMedia + locOfGroup.Long/float64(nLocInGroup)

	}

	group := NewGroupAvaerageLocation(loc.Id.String(), latMedia, longMedia)
	return group

}
