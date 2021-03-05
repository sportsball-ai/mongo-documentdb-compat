package mongodocumentdbcompat

import (
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

func appendAllKeys(keys []string, b bson.Raw) []string {
	elements, _ := b.Elements()
	for _, e := range elements {
		keys = append(keys, e.Key())
		keys = appendAllValueKeys(keys, e.Value())
	}
	return keys
}

func appendAllValueKeys(keys []string, b bson.RawValue) []string {
	if v, ok := b.ArrayOK(); ok {
		elements, _ := v.Elements()
		for _, e := range elements {
			keys = appendAllValueKeys(keys, e.Value())
		}
	} else if v, ok := b.DocumentOK(); ok {
		return appendAllKeys(keys, v)
	}
	return keys
}

func CheckKeys(b bson.Raw, version string) error {
	dollar, ok := dollar[version]
	if !ok {
		return fmt.Errorf("unsupported version")
	}
	for _, key := range appendAllKeys(nil, b) {
		if strings.HasPrefix(key, "$") {
			if supported, ok := dollar[key]; ok && !supported {
				return fmt.Errorf("unsupported key: " + key)
			}
		}
	}
	return nil
}
