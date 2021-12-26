package validations

import (
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ConvertToObjectIds(input string) ([]primitive.ObjectID, error) {
	var ids []string
	if len(input) > 0 {
		ids = strings.Split(input, ",")
	}
	oids := []primitive.ObjectID{}
	for _, id := range ids {
		tmp, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		oids = append(oids, tmp)
	}
	return oids, nil
}
