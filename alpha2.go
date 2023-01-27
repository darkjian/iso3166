package iso316

import (
	"fmt"
	"sync"
)

// alpha2 used for query using ISO3166 Alpha2 code
type alpha2 struct {
	data map[string]*data
	mu   *sync.RWMutex
}

// Exist check if Alpha2 code is correct, pass UPPERCASE iso code
func (a alpha2) Exist(code string) bool {
	a.mu.RLock()
	defer a.mu.RUnlock()

	if _, found := a.data[code]; !found {
		return false
	}

	return true
}

// GetAlpha3 return Alpha3 code by Alpha2 code, pass UPPERCASE iso code
func (a alpha2) GetAlpha3(code string) (string, error) {
	a.mu.RLock()
	defer a.mu.RUnlock()

	val, found := a.data[code]
	if !found {
		return "", fmt.Errorf("%q not in the list", code)
	}

	return val.A3, nil
}
