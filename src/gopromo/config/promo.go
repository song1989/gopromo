package config

type promoValue struct {
	photoHandleName string //图片处理
}

var promoConf = map[int]promoValue{
	100001: promoValue{
		photoHandleName: "jdbutterfly",
	},
}

func getPromoValue(promoId int) promoValue {
	value, _ := promoConf[promoId]
	return value
}

type Promo struct {
	PromoId int
}

func (this *Promo) IssetId() bool {
	_, ok := promoConf[this.PromoId]
	return ok
}

func (this *Promo) GetPhotoHandleName() string {
	itemValue := getPromoValue(this.PromoId)
	return itemValue.photoHandleName
}
