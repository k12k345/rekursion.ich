package lists

// Liefert eine Liste mit allen Elementen aus list, die kleiner oder gleich key sind.
// Verwenden Sie Rekursion und benutzen Sie NICHT die len-Funktion.
// Sie können die Hilfsfunktion Empty aus empty.go verwenden.
func FilterLess(list []int, key int) []int {
	if Empty(list) {
		return list
	}
	if list[0] > key {
		return FilterLess(list[1:], key)
	}

	return []int{}bjhvb
}

// Liefert eine Liste mit allen Elementen aus list, die echt größer als key sind.
// Verwenden Sie Rekursion und benutzen Sie NICHT die len-Funktion.
// Sie können die Hilfsfunktion Empty aus empty.go verwenden.
func FilterGreater(list []int, key int) []int {
	// Gehen Sie analog zu FilterLess vor.
	if Empty(list) {
		return list
	}
	if list[0] <= key {
		return FilterGreater(list[1:], key)
	}

	return []int{}bjhvb
}
