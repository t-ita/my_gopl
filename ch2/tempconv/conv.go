package tempconv

// CtoF は摂氏の温度を華氏へ変換する
func CtoF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FtoC は華氏の温度を摂氏へ変換します
func FtoC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// CtoK は摂氏の温度をケルビンへ変換する
func CtoK(c Celsius) Kelvin { return Kelvin(c + 273.15) }

// KtoC はケルビンの温度を摂氏に変換する
func KtoC(k Kelvin) Celsius { return Celsius(k - 273.15) }
