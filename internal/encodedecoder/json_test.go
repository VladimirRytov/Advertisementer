package encodedecoder

import (
	"bytes"
	"log/slog"
	"os"
	"testing"
	"time"

	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
	"github.com/VladimirRytov/advertisementer/internal/logging"
)

var testModel = &datatransferobjects.JsonStr{
	Clients: []datatransferobjects.ClientDTO{{
		Name:                  "Вася",
		Phones:                "88005553535",
		Email:                 "asdad@asdd.ru",
		AdditionalInformation: "Норм чел",

		Orders: []datatransferobjects.OrderDTO{{
			ID:            1,
			ClientName:    "Вася",
			Cost:          11,
			PaymentType:   "Картой",
			CreatedDate:   time.Now(),
			PaymentStatus: true,

			LineAdvertisements: []datatransferobjects.LineAdvertisementDTO{{
				Advertisement: datatransferobjects.Advertisement{
					ID:           1,
					OrderID:      1,
					ReleaseCount: 1,
					Cost:         11,
					Text:         "asd",
					Tags:         []string{"tag A", "tag B", "tag C"},
					ExtraCharges: []string{"charge A", "charge B", "charge C"},
					ReleaseDates: []time.Time{time.Now()},
				},
			}},

			BlockAdvertisements: []datatransferobjects.BlockAdvertisementDTO{{
				Advertisement: datatransferobjects.Advertisement{
					ID:           1,
					OrderID:      1,
					ReleaseCount: 1,
					Cost:         11,
					Text:         "asd",
					Tags:         []string{"tag A", "tag B", "tag C"},
					ExtraCharges: []string{"charge A", "charge B", "charge C"},
					ReleaseDates: []time.Time{time.Now()},
				},
				Size: 1,
			}},
		}},
	}},

	Tags: []datatransferobjects.TagDTO{
		{TagName: "tag A", TagCost: 1}, {TagName: "tag B", TagCost: 2}, {TagName: "tag C", TagCost: 3}},

	ExtraCharges: []datatransferobjects.ExtraChargeDTO{
		{ChargeName: "charge 1", Multiplier: 1}, {ChargeName: "charge 2", Multiplier: 2}, {ChargeName: "charge 3", Multiplier: 3}},
}

func CreateLogger() {
	logging.CreateLogger(".", 0, &slog.HandlerOptions{Level: slog.LevelInfo}, false, os.Stderr)
}

func TestToJSON(t *testing.T) {
	CreateLogger()
	var f bytes.Buffer
	err := ToJSON(&f, testModel, true)
	if err != nil {
		t.Fatal(err)
	}
}

func BenchmarkToJSONReadable(b *testing.B) {
	CreateLogger()

	var f bytes.Buffer
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ToJSON(&f, testModel, true)
		f.Reset()
	}
}

func BenchmarkToJSONUnReadable(b *testing.B) {
	CreateLogger()

	var f bytes.Buffer
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ToJSON(&f, testModel, false)
		f.Reset()
	}
}

func TestFromJSON(t *testing.T) {
	CreateLogger()
	var b bytes.Buffer
	ToJSON(&b, testModel, false)
	err := FromJSON(&datatransferobjects.JsonStr{}, &b)
	if err != nil {
		t.Fatal(err)
	}
}

func TestFromJSONArray(t *testing.T) {
	CreateLogger()
	var b bytes.Buffer
	var test []datatransferobjects.JsonStr
	for i := 0; i < 10; i++ {
		test = append(test, *testModel)
	}
	ToJSON(&b, test, false)
	var got []datatransferobjects.JsonStr
	err := FromJSON(&got, &b)
	if err != nil {
		t.Fatal(err)
	}
	if len(test) != len(got) {
		t.Fatal("different length")
	}
}

func BenchmarkFromJSONUnreadable(b *testing.B) {
	CreateLogger()

	var f bytes.Buffer
	ToJSON(&f, testModel, false)

	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FromJSON(&datatransferobjects.JsonStr{}, &f)
	}
}

func BenchmarkFromJSONReadable(b *testing.B) {
	CreateLogger()

	var f bytes.Buffer
	ToJSON(&f, testModel, true)

	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FromJSON(&datatransferobjects.JsonStr{}, &f)
	}
}
