Comando curl:

	curl -X POST \
	http://localhost:3001/topsecret \
	-H 'cache-control: no-cache' \
	-H 'content-type: application/json' \
	-H 'postman-token: fc139aff-8f9b-de88-495b-165d3fb2422e' \
	-d '{
		"satellites": [
		{
			"name": "kenobi",
			"distance": 100.0,
			"message": ["este", "", "", "mensaje", ""]
		},
		{
			"name": "skywalker",
			"distance": 115.5,
			"message": ["", "es", "", "", "secreto"]
		},
		{
			"name": "sato",
			"distance": 142.7,
			"message": ["este", "", "un", "", ""]
		}
		]
	}'

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

