Nivel 1:



Luego de descargarse el proyecto, ingresar a la carpeta nivel_1 y ejecutar:

go run main.go

Dependencias:
github.com/chewxy/math32
(para utilizar el tipo indicado float32 y no float64 que utiliza la librería math de go por default)


Funcionamiento: 

Se obtiene archivo data.json de la carpeta data con la siguiente estructura:

{
    "satellites": [
    {
        "name": "kenobi",
        "distance": 100.0,
        "message": ["este", "", "", "", ""],
        "x" : -500,
        "y" : -200
    },
    {
        "name": "skywalker",
        "distance": 115.5,
        "message": ["", "mensaje", "es", "", ""],
        "x" : 100,
        "y" : -100
    },
    {
        "name": "sato",
        "distance": 142.7,
        "message": ["", "", "", "un", "secreto"],
        "x" : 500,
        "y" : 100
    }
    ]}
Con los datos de los satélites, se obtiene el punto desconocido:

Demostración:

Fórmula distancia:
	(x - x₁)² + (y - y₁)² = d²

Ecuaciones (E1,E2...):

	(x - x₁)² + (y - y₁)² = d₁² 	(E1)
	(x - x₂)² + (y - y₂)² = d₂² 	(E2)
	(x - x₃)² + (y - y₃)² = d₃² 	(E3)

Sea (a + b) = a² + 2ab + b², simplificando cuadráticas y sustituyendo ecuación 3 en E1 y E2

	2(x₁ - x₂)x + 2(y₁ - y₂)y =  (x₁² + y₁² - d₁²) - (x₂² + y₂² - d₂²) 	(E4)
	2(x₂ - x₃)x + 2(y₂ - y₃)y =  (x₂² + y₂² - d₂²) - (x₃² + y₃² - d₃²) 	 (E5)


X e Y se encuentran resolviendo E4 y E5 por regla de Cramer:


	 		| (x₁² + y₁² - d₁²) - (x₂² + y₂² - d₂²)		2(y₁ - y₂) |
			| (x₂² + y₂² - d₂²) - (x₃² + y₃² - d₃²)		2(y₂ - y₃) |
	X =		----------------------------------------------------------------------						|	2(x₁ - x₂)	2(y₁ - y₂)	|
			|	2(x₂ - x₃)	2(y₂ - y₃)	|


	 		| 2(x₁ - x₂)		(x₁² + y₁² - d₁²) - (x₂² + y₂² - d₂²) |
			| 2(x₂ - x₃)		(x₂² + y₂² - d₂²) - (x₃² + y₃² - d₃²) |
	Y =		-----------------------------------------------------------------------
				|	2(x₁ - x₂)	2(y₁ - y₂)	|
				|	2(x₂ - x₃)	2(y₂ - y₃)	|



En caso de datos faltantes en el archivo json la aplicaciòn devuelve el siguiente error:
Datos insuficientes, para obtener las coordenadas

Nota:
	Las funciones siguen la firma pedida en el requerimiento.

Nivel 2:



Ingresar a la carpeta nivel_2 y ejecutar:

go run main.go

Dependencias:
github.com/chewxy/math32
(para utilizar el tipo indicado float32 y no float64 que utiliza la librería math de go por default)
github.com/gorilla/mux
Nivel 3:



Ingresar a la carpeta nivel_3 y ejecutar:

go run main.go

Dependencias:
github.com/chewxy/math32
(para utilizar el tipo indicado float32 y no float64 que utiliza la librería math de go por default)
github.com/gorilla/mux
github.com/patrickmn/go-cache
