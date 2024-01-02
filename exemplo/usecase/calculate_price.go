package usecase

import "github.com/douglasfariasil/exemploemgo/exemplo/entity"

type OrderInput struct {
	ID    string  `json:"id"`
	Price float64 `json:"xxxx"`
	Tax   float64 `json:"tax"`
}

// {"id": "1", "price": 10.0, "tax": 0.1}

type OrderOutput struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

// SOLID - "D" - Dependency Inversion Principle
type CalcuculateFinalPrice struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewCalculateFinalPrice(orderRepository entity.OrderRepositoryInterface) *CalcuculateFinalPrice {
	return &CalcuculateFinalPrice{
		OrderRepository: orderRepository,
	}
}

func (c *CalcuculateFinalPrice) Execute(input OrderInput) (*OrderOutput, error) {
	order, err := func() (*entity.Order, error) {
		order := &entity.Order{ID: input.ID, Price: float32(input.Price), Tax: float32(input.Tax)}
		err := order.Validade()
		if err != nil {
			return nil, err
		}
		return order, nil
	}()
	if err != nil {
		return nil, err
	}
	err = order.CalcuculateFinalPrice()
	if err != nil {
		return nil, err
	}
	err = c.OrderRepository.Save(order)
	if err != nil {
		return nil, err
	}
	return &OrderOutput{
		ID:         order.ID,
		Price:      float64(order.Price),
		Tax:        float64(order.Tax),
		FinalPrice: float64(order.FinalPrice),
	}, nil
}	