package domaintype

type Promotion struct {
}

func NewPromotion() *Promotion {
	return &Promotion{}
}

type PromotionType struct {
	Slug         string
	Name         string
	Description  string
	IsDiscount   bool
	IsOffer      bool
	IsPercentage bool
	IsQuantity   bool
	IsShipping   bool
}

func (f *Promotion) ListPromotionTypes() []PromotionType {
	return []PromotionType{
		{
			Slug:         `discount`,
			Name:         `Discount`,
			Description:  ``,
			IsDiscount:   true,
			IsOffer:      false,
			IsPercentage: false,
			IsQuantity:   false,
			IsShipping:   false,
		},
		{
			Slug:         `offer`,
			Name:         `Offer`,
			Description:  ``,
			IsDiscount:   false,
			IsOffer:      true,
			IsPercentage: false,
			IsQuantity:   false,
			IsShipping:   false,
		},
		{
			Slug:         `percentage`,
			Name:         `Percentage`,
			Description:  ``,
			IsDiscount:   false,
			IsOffer:      false,
			IsPercentage: true,
			IsQuantity:   false,
			IsShipping:   false,
		},
		{
			Slug:         `quantity`,
			Name:         `Quantity`,
			Description:  ``,
			IsDiscount:   false,
			IsOffer:      false,
			IsPercentage: false,
			IsQuantity:   true,
			IsShipping:   false,
		},
		{
			Slug:         `shipping`,
			Name:         `Shipping`,
			Description:  ``,
			IsDiscount:   false,
			IsOffer:      false,
			IsPercentage: false,
			IsQuantity:   false,
			IsShipping:   true,
		},
	}
}
