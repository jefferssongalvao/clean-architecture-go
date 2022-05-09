package recommendation

import (
	"clean-architecture/src/domain/entities/student"
	"errors"
)

type Recommendation struct {
	indicator student.Student
	indicated student.Student
}

func CreateRecommendation(indicator, indicated student.Student) (Recommendation, error) {
	var recommendation Recommendation
	recommendation.indicator = indicator
	recommendation.indicated = indicated

	if error := recommendation.validate(); error != nil {
		return Recommendation{}, error
	}

	return recommendation, nil
}

func (recommendation *Recommendation) GetIndicator() student.Student {
	return recommendation.indicator
}

func (recommendation *Recommendation) GetIndicated() student.Student {
	return recommendation.indicated
}

func (recommendation *Recommendation) validate() error {
	if recommendation.indicator.Cpf() == recommendation.indicated.Cpf() {
		return errors.New("invalid recommendation")
	}
	return nil
}
