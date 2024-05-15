package advertisements

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	want := Client{name: "Bob"}
	got, _ := NewClient("Bob")
	if got != want {
		t.Errorf("want %v,got %v", want, got)
	}
}
func TestSetName(t *testing.T) {
	want := Client{name: "Bo"}
	got, _ := NewClient("Bib")
	got.SetName("Bo")
	if got != want {
		t.Errorf("want %v,got %v", want, got)
	}
	err := got.SetName("b")
	if err == nil {
		t.Errorf("want error, got nil")
	}
}
func TestSetContactNumber(t *testing.T) {
	numbers := []string{"88005553535", "8 800 555 35 35", "8(800)5553535", "8 (800) 555 35 35", "8-800-555-35-35"}
	want := Client{name: "Bob", contactNumbers: "88005553535"}
	got, _ := NewClient("Bob")
	got.setContactNumber("88005553535")
	if got != want {
		t.Errorf("want %v,got %v", want, got)
	}
	for _, v := range numbers {
		err := got.setContactNumber(v)
		if err != nil {
			t.Errorf("error %v,number %s", err, v)
		}
		if got.contactNumbers != v {
			t.Errorf("want contact number %s,got contact number %s", v, got.contactNumbers)
		}
	}
}

func TestSetEmail(t *testing.T) {
	want := []Client{{name: "Bob", email: "asdzxc@gmail.com"}, {name: "Bob", email: "xcvxcv@gmail.com"}}
	mailCases := []string{"asdzxc@gmail.com", "xcvxcv@gmail.com"}
	for i, v := range want {
		got, _ := NewClient("Bob")
		err := got.setEmail(mailCases[i])
		if err != nil {
			t.Fatal(err)
		}
		if v.email != got.email {
			t.Errorf("want %v, got %v", v.email, got.email)
		}
	}
}

func TestSetWrongEmailEmail(t *testing.T) {
	want := []Client{{name: "Bob", email: "asdzxc@gmail.com"}, {name: "Bob", email: ""}}
	mailCases := []string{"asdzxc@.com", "@gmail.com", "asdzxcgmail.com"}
	for i := range want {
		got, _ := NewClient("Bob")
		err := got.setEmail(mailCases[i])
		if err == nil {
			t.Fatalf("want error, got nil")
		}
	}
}

func TestName(t *testing.T) {
	want := Client{name: "Bob"}
	got := want.Name()
	if got != want.name {
		t.Errorf("want %v,got %v", want, got)
	}

}

func TestContactNumber(t *testing.T) {
	want := Client{name: "Bob", contactNumbers: "88005553535"}
	got := want.ContactNumber()
	if got != want.contactNumbers {
		t.Errorf("want %v,got %v", want, got)
	}
}
func TestEmail(t *testing.T) {
	want := Client{name: "Bob", email: "asdzxc@gmail.com"}
	got := want.Email()
	if got != want.email {
		t.Errorf("want %v,got %v", want, got)
	}
}

func TestSetAdditionalInformation(t *testing.T) {
	want := "Иван город Тверь"
	testClient, _ := NewClient("Иван")
	testClient.SetAdditionalInformation(want)
	if testClient.additionalInformation != want {
		t.Errorf("want %s, got %s", want, testClient.additionalInformation)
	}
}

func TestAdditionalInformation(t *testing.T) {
	want := "Иван город Тверь"
	testClient, _ := NewClient("Иван")
	testClient.SetAdditionalInformation(want)
	if got := testClient.AdditionalInformation(); got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}
