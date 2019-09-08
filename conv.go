package tempconv

func CToF(c Celsius) Fanrenheit {
	return Fanrenheit(c*9/5 + 32)
}

func FToC(f Fanrenheit) Celsius  {
	return Celsius(f - 32) * 5 / 9
}