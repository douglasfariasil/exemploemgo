package entity

import "errors"

type Order struct {
	ID         string
	Price      float32
	Tax        float32
	FinalPrice float32
}

func NewOrder(id string, price float32, tax float32) (*Order, error) {
	order := &Order{
		ID:    id,
		Price: price,
		Tax:   tax,
	}
	err := order.Validade()
	if err != nil {
		return nil, err
	}
	return order, nil	
}

func(o *Order) Validade() error {
	if o.ID == "" {
		return errors.New("id is required")
	}
	if o.Price <= 0 {
		return errors.New("price must be greater than zero")
	}
	if o.Tax <= 0 {
		return errors.New("invalid tax")
	}
	return nil
}

func (o *Order) CalcuculateFinalPrice() error {
	o.FinalPrice = o.Price + o.Tax
	err := o.Validade()
	if err != nil {
		return err
	}
	return nil
}
