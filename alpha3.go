package iso316

import (
	"fmt"
	"sync"
)

// alpha3 used for query using ISO3166 Alpha3 code
type alpha3 struct {
	data map[string]*data
	mu   *sync.RWMutex
}

// Exist check if Alpha3 code is correct, pass UPPERCASE iso code
func (a alpha3) Exist(code string) bool {
	a.mu.RLock()
	defer a.mu.RUnlock()

	if _, found := a.data[code]; !found {
		return false
	}

	return true
}

// GetAlpha2 return Alpha2 code by Alpha3 code, pass UPPERCASE iso code
func (a alpha3) GetAlpha2(code string) (string, error) {
	a.mu.RLock()
	defer a.mu.RUnlock()

	val, found := a.data[code]
	if !found {
		return "", fmt.Errorf("%q not in the list", code)
	}

	return val.A2, nil
}
