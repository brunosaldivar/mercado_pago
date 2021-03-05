package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	. "github.com/brunosaldivar/mercado_pago/nivel_1/pkg"
	math "github.com/chewxy/math32"
)

var satellitesJSON Satellites

func getData() (Satellites, error) {

	satellitesJSON = Satellites{}
	data, err := ioutil.ReadFile("./data/data.json")
	if err != nil {
		fmt.Println(FileOpenError, err)
	}
	err = json.Unmarshal(data, &satellitesJSON)
	if err != nil {
		fmt.Println(ParseError, err)
	}
	return satellitesJSON, err
}
func main() {

	//data del json
	s, err := getData()
	if err != nil {
		os.Exit(0)
	}
	var distances []float32
	var messages [][]string

	//recorro los satellites para enviar solo las distancias a la fc' segùn requerimiento
	for i := 0; i < len(s.Satellites); i++ {
		distances = append(distances, s.Satellites[i].Distance)
	}
	//recorro los satellites para enviar solo los msg a la fc' segùn requerimiento
	for i := 0; i < len(s.Satellites); i++ {
		messages = append(messages, s.Satellites[i].Message)
	}
	msg := GetMessage(messages...)
	x, y := GetLocation(distances...)

	rtn := ResponseMain{
		Position: Point{X: x, Y: y},
		Message:  msg,
	}
	fmt.Println(EnemyCoordinates, "{X:", rtn.Position.X, "Y:", rtn.Position.Y, "}")
	fmt.Println(MessageReceived, rtn.Message)
}

// input: distancia al emisor tal cual se recibe en cada satélite
// output: las coordenadas ‘x’ e ‘y’ del emisor del mensaje
func GetLocation(distances ...float32) (x, y float32) {
	var a, b, c, d, e, f float32
	countData := len(satellitesJSON.Satellites)
	if len(distances) != countData || countData != 3 {
		fmt.Println(GetCoordinatesFail, VerifyJSONError)
		os.Exit(0)
	}

	for i := 0; i < len(satellitesJSON.Satellites); i++ {
		satellitesJSON.Satellites[0].SetDistance(distances[i])
	}

	s1 := satellitesJSON.Satellites[0]
	s2 := satellitesJSON.Satellites[1]
	s3 := satellitesJSON.Satellites[2]

	//2(x₁ - x₂)
	a = 2 * (s1.X - s2.X)
	//2(y₁ - y₂)
	b = 2 * (s1.Y - s2.Y)
	//(x₁² + y₁² - d₁²) - (x₂² + y₂² - d₂²)
	c = (math.Pow(s1.X, 2) + math.Pow(s1.Y, 2) + math.Pow(s1.Distance, 2)) - (math.Pow(s2.X, 2) + math.Pow(s2.Y, 2) + math.Pow(s2.Distance, 2))
	//2(x₂ - x₃)
	d = 2 * (s2.X - s3.X)
	//2(y₂ - y₃)
	e = 2 * (s2.Y - s3.Y)
	//x₂² + y₂² - d₂²) - (x₃² + y₃² - d₃²)
	f = (math.Pow(s2.X, 2) + math.Pow(s2.Y, 2) + math.Pow(s2.Distance, 2)) - (math.Pow(s3.X, 2) + math.Pow(s3.Y, 2) + math.Pow(s3.Distance, 2))

	//determinantes:
	detS := (a * e) - (d * b)
	detX := (c * e) - (f * b)
	detY := (a * f) - (d * c)

	x = detX / detS
	y = detY / detS

	return x, y

}

// input: el mensaje tal cual es recibido en cada satélite
// output: el mensaje tal cual lo genera el emisor del mensaje
func GetMessage(messages ...[]string) (msg string) {

	var parts []string
	for i := 0; i < len(messages); i++ {
		for j := 0; j < len(messages[i]); j++ {
			if strings.TrimSpace(messages[i][j]) != "" {
				parts = append(parts, messages[i][j])
			}
		}
	}
	return strings.Join(parts, " ")
}
