package di

import (
	"strings"
)

type seen map[string]int

func (s seen) add(info dependencyInfo) seen {
	newList := make(seen, len(s))

	for k, v := range s {
		newList[k] = v
	}
	newList[info.key] = len(newList)

	return newList
}

func (s seen) ordered() []string {
	keys := make([]string, len(s))

	for key, i := range s {
		keys[i] = key
	}

	return keys
}

func (s seen) String() string {
	return strings.Join(s.ordered(), ",")
}
