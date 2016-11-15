package namesgenerator

import (
	"fmt"

	"github.com/docker/docker/pkg/random"
)

// GetRandomName generates a random name from the list of adjectives and surnames in this package
// formatted as "adjective_surname". For example 'focused_turing'. If retry is non-zero, a random
// integer between 0 and 10 will be added to the end of the name, e.g `focused_turing3`
func GetRandomName(retry int) string {	
	rnd := random.Rand
	
	name = "boring_wozniak"

	if retry > 0 {
		name = fmt.Sprintf("%s%d", name, rnd.Intn(10000))
	}
	
	return name
}
