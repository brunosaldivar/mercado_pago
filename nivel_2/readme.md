
1- definir estructuras de tipo:
	- points
	- el json que llega
	- tipos:
		satellites []
		satellite : {
			"name"		: string , //o enum o tipo con point
			"distance"	: float64,
			"message"	: string[] // ver esto con listas enlazadas y como manejarlo ["este", "", "", "mensaje", ""]
		}
		"position": {
			"x": -100.0,
			"y": 75.5
		},

2- interfaces
3- fc' que reciba distancia
4- donde guardamos las 3 coordenadas? con los nombres de satelites?

 TODO:  RECORDAR SACAR ESTO EN PAPEL PARA EXPLICARLO
₀₁₂₃₄₅₆₇₈₉

Demostración:

	Fórmula distancia:
		(x - x₁)² + (y - y₁)² = d²

	Ecuaciones (E1,E2...):

		(x - x₁)² + (y - y₁)² = d₁² (E1)
		(x - x₂)² + (y - y₂)² = d₂² (E2)
		(x - x₃)² + (y - y₃)² = d₃² (E3)

	Sea (a + b) = a² + 2ab + b², simplificando cuadraticas y sustituyendo ecuacion 3 en E1 y E2

		2(x₁ - x₂)x + 2(y₁ - y₂)y =  (x₁² + y₁² - d₁²) - (x₂² + y₂² - d₂²) 	(E4)
		2(x₂ - x₃)x + 2(y₂ - y₃)y =  (x₂² + y₂² - d₂²) - (x₃² + y₃² - d₃²)  (E5)


	X e Y se encuentran resolviendo E4 y E5 por regla de Cramer:

	 		| (x₁² + y₁² - d₁²) - (x₂² + y₂² - d₂²)		2(y₁ - y₂) |
			| (x₂² + y₂² - d₂²) - (x₃² + y₃² - d₃²)		2(y₂ - y₃) |
	X =		----------------------------------------------------
					|	2(x₁ - x₂)	2(y₁ - y₂)	|
					|	2(x₂ - x₃)	2(y₂ - y₃)	|


	 		| 2(x₁ - x₂)		(x₁² + y₁² - d₁²) - (x₂² + y₂² - d₂²) |
			| 2(x₂ - x₃)		(x₂² + y₂² - d₂²) - (x₃² + y₃² - d₃²) |
	Y =		---------------------------------------------------
					|	2(x₁ - x₂)	2(y₁ - y₂)	|
					|	2(x₂ - x₃)	2(y₂ - y₃)	|

	//. "github.com/brunosaldivar/mercado_pago/pkg"

02-03
1 - rellenar con lo que viene del payload las estructuras y los puntos OK

2- cumplir con la firma :
	interfaz OK
	heredarla

	if len(satellites) > 0 {
		if satellites[0].Name == "Kenobi" {
			satellites[0].Coord = Point{X: -500, Y: -200}
		}
		if satellites[1].Name == "Skywalker" {
			satellites[1].Coord = Point{X: 100, Y: -100}
		}
		if satellites[2].Name == "Sato" {
			satellites[2].Coord = Point{X: 500, Y: 100}
		}
		log.Println(satellites)
	}

3- el programa que lea de un json igual al del payload