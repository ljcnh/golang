package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64
type Inch float64
type Meter float64
type Ib float64
type Kilogram float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

//func CToF(c Celsius) Fahrenheit {
//	return Fahrenheit(c*9/5+32)
//}
//
//func FToC(f Fahrenheit) Celsius {
//	return Celsius((f-32)*5/9)
//}

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%g°K", k)
}

func (i Inch) String() string {
	return fmt.Sprintf("%g in", i)
}

func (m Meter) String() string {
	return fmt.Sprintf("%g m", m)
}

func (i Ib) String() string {
	return fmt.Sprintf("%g ib", i)
}

func (k Kilogram) String() string {
	return fmt.Sprintf("%g kg", k)
}

//func main() {
//	fmt.Printf( "%g\n" , BoilingC-FreezingC) // "100" °C
//	boilingF := CToF(BoilingC)
//	fmt.Printf( "%g\n" , boilingF-CToF(FreezingC)) // "180" °F
//	fmt.Printf( "%g\n" , boilingF-Fahrenheit(FreezingC))
//	c := FToC(212.0)
//	fmt.Println(c.String())
//	fmt.Printf("%v\n",c)
//	fmt.Printf("%s\n",c)
//	fmt.Println(c)
//	fmt.Printf("%g\n",c)
//	fmt.Println(float64(c))
//}
