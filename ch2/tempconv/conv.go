package tempconv

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func KToC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

func CToK(c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}

func IToM(i Inch) Meter {
	return Meter(i * 0.0254)
}

func MToI(m Meter) Inch {
	return Inch(m / 0.0254)
}

func IToK(i Ib) Kilogram {
	return Kilogram(i * 0.4535)
}

func KToI(k Kilogram) Ib {
	return Ib(k / 0.4535)
}
