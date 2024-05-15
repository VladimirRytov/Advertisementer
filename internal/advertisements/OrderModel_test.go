package advertisements

import (
	"errors"
	"testing"
	"time"
)

func TestNewAdvertisementOrder(t *testing.T) {
	want := AdvertisementOrder{clientId: "154"}
	got, err := NewAdvertisementOrder("154")
	if want != got && err != nil {
		t.Errorf("want %v, got %v", want, got)
	}
	_, err = NewAdvertisementOrder("1")
	if !errors.Is(err, errClientName) {
		t.Errorf("want error, got %v", err)
	}

}

func TestSetOrderId(t *testing.T) {
	want := AdvertisementOrder{id: 10}
	got, _ := NewAdvertisementOrder("123")
	got.SetOrderId(10)
	if want.id != got.id {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestSetOrderClientId(t *testing.T) {
	want := AdvertisementOrder{clientId: "100"}
	got, _ := NewAdvertisementOrder("200")
	got.SetClientId("100")
	if want.clientId != got.clientId {
		t.Errorf("want %v, got %v", want.clientId, got.clientId)
	}
}
func TestSetOrderCost(t *testing.T) {
	want := []int{30, 45, 20, 90.00}
	ao := AdvertisementOrder{}
	for i := range want {
		ao.SetOrderCost(want[i])
		if want[i] != ao.cost {
			t.Errorf("want %v, got %v", want[i], ao.cost)
		}
	}
}
func TestOrderId(t *testing.T) {
	want := 20
	testAdv, _ := NewAdvertisementOrder("100")
	testAdv.SetOrderId(want)
	got := testAdv.OrderId()
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestOrderClientId(t *testing.T) {
	want := "Вася"
	testAdv, _ := NewAdvertisementOrder("Вяся")
	testAdv.SetClientId(want)
	got := testAdv.ClientId()
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}
func TestCost(t *testing.T) {
	want := 142
	testAdv, _ := NewAdvertisementOrder("100")
	testAdv.SetOrderCost(want)
	got := testAdv.Cost()
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}
func TestPaymentStatus(t *testing.T) {
	order, _ := NewAdvertisementOrder("100")
	order.SetPaymentStatus(true)
	if !order.PaymentStatus() {
		t.Errorf("want true,get false")
	}
}
func TestSetGetCreatedDate(t *testing.T) {
	want := time.Date(2023, 12, 11, 0, 0, 0, 0, time.UTC)
	order, _ := NewAdvertisementOrder("100")
	order.SetCreaatedDate(want)
	if got := order.CreaatedDate(); got != want {
		t.Errorf("want %v,get %v", want, got)
	}
}
func TestSetGetPaymentType(t *testing.T) {
	want := "Чек"
	order, _ := NewAdvertisementOrder("100")
	order.SetPaymentType(want)
	if got := order.PaymentType(); got != want {
		t.Errorf("want %v,get %v", want, got)
	}
}
