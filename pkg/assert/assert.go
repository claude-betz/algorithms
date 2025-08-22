package assert

import "log"

func True(p bool, msg string) {
	if !p {
		/* invariants */
		log.Panic(msg)
	}
}
