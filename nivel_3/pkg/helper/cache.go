package helper

import (
	"time"

	. "mercado_pago/nivel_3/pkg/structs"

	"github.com/patrickmn/go-cache"
)

//seteo cache por 24 hs
var Cache = cache.New(24*time.Hour, 24*time.Hour)

//Guarda en cache el satellite enviado por POST en caso de no existir
//crea la cache con struct satellites
func SetCache(satellite Satellite) *Satellites {

	s, found := GetCache()
	newSatellite := true
	if found {
		for i := 0; i < len(s.Satellites); i++ {

			if s.Satellites[i].Name == satellite.Name {
				s.Satellites[i].Set(satellite)
				newSatellite = false
			}
		}
		if newSatellite && len(s.Satellites) <= 3 {
			s.Satellites = append(s.Satellites,
				Satellite{
					Name:     satellite.Name,
					Distance: satellite.Distance})
		}
	} else {

		s = Satellites{}
		s.Satellites = append(s.Satellites,
			Satellite{
				Name:     satellite.Name,
				Distance: satellite.Distance})
	}
	Cache.Set("satellites", s, cache.NoExpiration)
	return &s
}

//Obtiene struct satelites cacheada. En caso de no encontrarla retorna found = false
func GetCache() (Satellites, bool) {
	var satellites Satellites
	var found bool
	data, found := Cache.Get("satellites")
	if found {
		satellites = data.(Satellites)
	}
	return satellites, found
}

//borra cache
func Clear() {
	Cache.Flush()
}
