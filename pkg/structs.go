package structs

// Point Estructura x,y
type Point struct {
	X float64
	Y float64
}
type Satellite struct {
	name     string
	coord    Point   //coordenadas del satelite, punto conocido
	distance float64 //distancia entre el satelite y la nave enemiga
	//TODO: recordar pasar a 32
}
