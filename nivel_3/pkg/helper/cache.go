package helper

import (
	"log"
	"time"

	. "github.com/brunosaldivar/mercado_pago/nivel_3/pkg/structs"
	"github.com/patrickmn/go-cache"
)

var Cache = cache.New(24*time.Hour, 24*time.Hour)

func SetCache(satellite Satellite) *Satellites {

	s, found := GetCache()
	newSatellite := true
	log.Println("SetCache() found:", found, s, satellite)

	if found {
		for i := 0; i < len(s.Satellites); i++ {

			if s.Satellites[i].Name == satellite.Name {
				//ver aca
				log.Println("SetCache: ", satellite.Name)
				s.Satellites[i].Set(satellite)
				log.Println("vuelta: ", s.Satellites[i].Name)
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

//aca tiene q estar la logica de que si existe el satellite reemplazarlo
//y se guarda la lista completa o si va a buscarlo y lo reemplaza

func GetCache() (Satellites, bool) {
	var satellites Satellites
	var found bool
	data, found := Cache.Get("satellites")
	if found {
		satellites = data.(Satellites)
		log.Println("GetCache()  ", satellites)

	}
	return satellites, found
}

func Clear() {
	Cache.Flush()
}
