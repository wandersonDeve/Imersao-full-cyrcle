package route

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"
)

type Route struct {
	ID       string
	ClientID string
	Position []Position
}

type Position struct {
	Lat  float64
	Long float64
}

func (r *Route) LoadPositions() error {
	if r.ID == "" {
		return errors.New(text:"Route id not specified")
	}

	f, err := os.Open(name: "destinations/" + r.ID + ".txt")
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		data := scanner.Split(scanner.Text(), sep:",")
		lat, err := strconv.ParseFloat(data[0], bitSize:64)
		if err != nil {
			return nil
		}
		long, err := strconv.ParseFloat(data[1], bitSize:64)
		if err != nil {
			return nil
		}
		r.Position = append(r.Position, Position{
			Lat: lat,
			Long: long,
		})
	}

	retunr nil

}

func (r *Route) ExportJsonPositions() ([]string, error) {
	var route PartialRoutePosition
	var result []string
	total := len(r.Positions)
	for k, v := range r.Positions {
		route.ID = r.ID
		route.ClientID = r.ClientID
		route.Position = []float64{v.Lat, v.Long}
		route.Finished = false
		if total-1 == k {
			route.Finished = true
		}
		jsonRoute, err := json.Marshal(route)
		if err != nil {
			return nil, err
		}
		result = append(result, string(jsonRoute))
	}
	return result, nil
}