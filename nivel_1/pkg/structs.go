package structs

import (
	"strings"
)

// Point Estructura x,y

type Point struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

type Satellite struct {
	Name     string   `json:"name"`
	Distance float32  `json:"distance"`
	Message  []string `json:"message"`
	//coordenadas del satelite, punto conocido
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

type Satellites struct {
	Satellites []Satellite `json:"satellites"`
}

type ResponseTopSecret struct {
	Position Point  `json:"position"`
	Message  string `json:"message"`
}

func (s *Satellites) GetMessage() string {
	var parts []string
	for i := 0; i < len(s.Satellites); i++ {
		for j := 0; j < len(s.Satellites[i].Message); j++ {
			if strings.TrimSpace(s.Satellites[i].Message[j]) != "" {
				parts = append(parts, s.Satellites[i].Message[j])
			}
		}
	}
	return strings.Join(parts, " ")
}

func (s *Satellite) SetDistance(distance float32) error {

	s.Distance = distance
	return nil
}
