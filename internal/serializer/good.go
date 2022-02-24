package serializer

import "Opendulum/internal/model"

type Good struct {
	ID            uint   `json:"id"`
	Brand         string `json:"brand"`
	MapProduction string
	MapIngredient string
	MapTrends     string
	Energy        int32
	Protein       int32
	Fat           int32
	Carbohydrates int32
	Minerals      int32
	Other         int32
}

func BuildGood(it model.Good) Good {
	return Good{
		ID:            it.ID,
		Brand:         it.Brand,
		MapProduction: it.MapProduction,
		MapIngredient: it.MapIngredient,
		MapTrends:     it.MapTrends,
		Energy:        it.Energy,
		Protein:       it.Protein,
		Fat:           it.Fat,
		Carbohydrates: it.Carbohydrates,
		Minerals:      it.Minerals,
		Other:         it.Other,
	}
}
