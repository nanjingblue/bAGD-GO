package serializer

import "Opendulum/internal/model"

type GoodJingDong struct {
	ID    uint   `json:"id"`
	Brand string `json:"brand"`
	JDId  string `json:"jd_id"`
}

func BuildGoodJingDong(it model.GoodJingDong) GoodJingDong {
	return GoodJingDong{
		ID:    it.ID,
		Brand: it.Brand,
		JDId:  it.JDId,
	}
}
