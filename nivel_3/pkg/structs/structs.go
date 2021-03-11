package structs

import (
	"errors"
	"strings"

	math "github.com/chewxy/math32"
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
	Coord Point
}

type Satellites struct {
	Satellites []Satellite `json:"satellites"`
}

type ResponseTopSecret struct {
	Position Point  `json:"position"`
	Message  string `json:"message"`
}

func (s *Satellites) CalculateCoordinates() (*Point, error) {

	var a, b, c, d, e, f float32
	var point = new(Point)
	if len(s.Satellites) != 3 {
		return nil, errors.New("Datos insuficientes, para obtener las coordenadas")
	}
	//seteo coordenadas
	for i := 0; i < len(s.Satellites); i++ {
		s.Satellites[i].SetPoint()
	}
	s1 := s.Satellites[0]
	s2 := s.Satellites[1]
	s3 := s.Satellites[2]
	//2(x₁ - x₂)
	a = 2 * (s1.Coord.X - s2.Coord.X)
	//2(y₁ - y₂)
	b = 2 * (s1.Coord.Y - s2.Coord.Y)
	//(x₁² + y₁² - d₁²) - (x₂² + y₂² - d₂²)
	c = (math.Pow(s1.Coord.X, 2) + math.Pow(s1.Coord.Y, 2) + math.Pow(s1.Distance, 2)) - (math.Pow(s2.Coord.X, 2) + math.Pow(s2.Coord.Y, 2) + math.Pow(s2.Distance, 2))
	//2(x₂ - x₃)
	d = 2 * (s2.Coord.X - s3.Coord.X)
	//2(y₂ - y₃)
	e = 2 * (s2.Coord.Y - s3.Coord.Y)
	//x₂² + y₂² - d₂²) - (x₃² + y₃² - d₃²)
	f = (math.Pow(s2.Coord.X, 2) + math.Pow(s2.Coord.Y, 2) + math.Pow(s2.Distance, 2)) - (math.Pow(s3.Coord.X, 2) + math.Pow(s3.Coord.Y, 2) + math.Pow(s3.Distance, 2))

	//determinantes:
	detS := (a * e) - (d * b)
	detX := (c * e) - (f * b)
	detY := (a * f) - (d * c)

	x := detX / detS
	y := detY / detS

	point.X = x
	point.Y = y

	return point, nil
}

func (s *Satellites) GetMessage() (*string, error) {
	var parts []string

	if len(s.Satellites) != 3 {
		return nil, errors.New("Datos insuficientes, para obtener el mensaje")
	}
	for i := 0; i < len(s.Satellites); i++ {
		for j := 0; j < len(s.Satellites[i].Message); j++ {
			if strings.TrimSpace(s.Satellites[i].Message[j]) != "" {
				parts = append(parts, s.Satellites[i].Message[j])
			}
		}
	}
	rtn := strings.Join(parts, " ")
	return &rtn, nil
}

//Setea coordenadas. Datos conocidos
func (s *Satellite) SetPoint() error {

	switch strings.ToLower(s.Name) {
	case "kenobi":
		s.Coord = Point{X: -500, Y: -200}
		break
	case "skywalker":
		s.Coord = Point{X: 100, Y: -100}
		break
	case "sato":
		s.Coord = Point{X: 500, Y: 100}
		break
	}
	return nil
}

func (s *Satellite) Set(satelite Satellite) {
	s.Name = satelite.Name
	s.Distance = satelite.Distance
	s.Message = satelite.Message

}
