package serializer

import "Opendulum/internal/model"

type Good struct {
	ID            uint   `json:"id"`
	Brand         string `json:"brand"`
	MapProduction string `json:"map_production"`
	MapIngredient string `json:"map_ingredient"`
	MapTrends     string `json:"map_trends"`
	Energy        int32  `json:"energy"`
	Protein       int32  `json:"protein"`
	Fat           int32  `json:"fat"`
	Carbohydrates int32  `json:"carbohydrates"`
	Minerals      int32  `json:"minerals"`
	Other         int32  `json:"other"`
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
