package utils

func RemoveDuplicates[T comparable](arr []T) []T {
	keys := make(map[T]struct{}) // Используем struct{}, так как он не занимает памяти
	var list []T                 // Используем var для более идиоматичного кода
	for _, entry := range arr {
		if _, exists := keys[entry]; !exists { // Более явное имя переменной
			keys[entry] = struct{}{} // Используем struct{} для хранения факта присутствия
			list = append(list, entry)
		}
	}
	return list
}
