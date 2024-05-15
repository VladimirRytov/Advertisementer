package advertisements

import (
	"errors"
	"time"
)

type AdvertisementOrder struct {
	id            int
	clientId      string
	cost          int
	createdDate   time.Time
	paymentStatus bool
	paymentType   string
}

func NewAdvertisementOrder(clientID string) (AdvertisementOrder, error) {
	_, err := NewClient(clientID)
	if err != nil {
		return AdvertisementOrder{}, err
	}
	return AdvertisementOrder{clientId: clientID}, nil
}

func (ao *AdvertisementOrder) SetOrderId(n int) error {
	if n < 0 {
		return errors.New("номер заказа должен быть больше или равен 0")
	}
	ao.id = n
	return nil
}
func (ao *AdvertisementOrder) SetPaymentStatus(status bool) {
	ao.paymentStatus = status
}

func (ao *AdvertisementOrder) SetClientId(id string) error {
	_, err := NewClient(id)
	if err != nil {
		return err
	}
	ao.clientId = id
	return nil
}

func (ao *AdvertisementOrder) SetOrderCost(cost int) error {
	if cost < 0 {
		return errors.New("стоимость заказа болжен быть больше 0")
	}
	ao.cost = cost
	return nil
}

func (ao *AdvertisementOrder) SetCreaatedDate(date time.Time) {
	ao.createdDate = date
}
func (ao *AdvertisementOrder) SetPaymentType(pay string) {
	ao.paymentType = pay
}

func (ao *AdvertisementOrder) OrderId() int {
	return ao.id
}

func (ao *AdvertisementOrder) ClientId() string {
	return ao.clientId
}
func (ao *AdvertisementOrder) Cost() int {
	return ao.cost
}
func (ao *AdvertisementOrder) CreaatedDate() time.Time {
	return ao.createdDate
}

func (ao *AdvertisementOrder) PaymentStatus() bool {
	return ao.paymentStatus
}
func (ao *AdvertisementOrder) PaymentType() string {
	return ao.paymentType
}
