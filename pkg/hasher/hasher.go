// Package hasher предоставляет функции для хеширования строк
package hasher

import (
	"hash/fnv"
	"strconv"
)

// StringToInt64 преобразует строку в int64 с помощью хеш-функции FNV-1a
// Возвращает положительное число int64
func StringToInt64(s string) int64 {
	h := fnv.New64a()
	_, _ = h.Write([]byte(s))
	hash := int64(h.Sum64())

	// Гарантируем, что хеш будет положительным
	if hash < 0 {
		hash = -hash
	}

	return hash
}

// Int64ToString преобразует int64 в строковое представление
func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}
