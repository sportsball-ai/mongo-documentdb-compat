package mongodocumentdbcompat

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
)

func TestAppendAllKeys(t *testing.T) {
	for name, tc := range map[string]struct {
		Input    interface{}
		Expected []string
	}{
		"D": {
			Input:    bson.D{{"foo", "bar"}, {"hello", "world"}, {"pi", 3.14159}},
			Expected: []string{"foo", "hello", "pi"},
		},
		"M": {
			Input:    bson.M{"foo": "bar", "hello": "world", "pi": 3.14159},
			Expected: []string{"foo", "hello", "pi"},
		},
		"A": {
			Input:    bson.M{"a": bson.A{"bar", "world", 3.14159, bson.D{{"qux", 12345}}}},
			Expected: []string{"a", "qux"},
		},
	} {
		t.Run(name, func(t *testing.T) {
			raw, err := bson.Marshal(tc.Input)
			require.NoError(t, err)
			actual := appendAllKeys(nil, raw)
			assert.ElementsMatch(t, tc.Expected, actual)
		})
	}
}

func TestCheckKeys(t *testing.T) {
	for name, tc := range map[string]struct {
		Input interface{}
		Okay  bool
	}{
		"Okay": {
			Input: bson.D{{"$foo", "bar"}, {"$hello", "world"}},
			Okay:  true,
		},
		"NotOkay": {
			Input: bson.A{bson.M{"$project": bson.M{"country": 1.0, "city": 1.0}}, bson.M{"$sortByCount": "$city"}},
			Okay:  false,
		},
	} {
		t.Run(name, func(t *testing.T) {
			_, raw, err := bson.MarshalValue(tc.Input)
			require.NoError(t, err)
			if tc.Okay {
				assert.NoError(t, CheckKeys(raw, "4.0"))
			} else {
				assert.Error(t, CheckKeys(raw, "4.0"))
			}
		})
	}
}
