//translate celsius and Fanrenheit
package tempconv

import (
	"fmt"
)

type Celsius float64
type Fanrenheit float64
const (
	AbosluteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%gC", c)
}

func (f Fanrenheit) String() string {
	return fmt.Sprintf("%g F", f)
}


