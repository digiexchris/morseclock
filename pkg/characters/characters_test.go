package characters
import "testing"

func TestGetCodeStringWithValidString(t *testing.T){

	err,got := getCodeString("test")
	want := "- . ... -"
	
	if string(got) != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	if err != nil {
		t.Errorf("got error %q", err)
	}
}

func TestGetCodeStringWithInvalidStringReturnsSpaces(t *testing.T){

	err,got := getCodeString("%%####")
	want := "     "

	if string(got) != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	if err != nil {
		t.Errorf("got error %q", err)
	}
}

func TestGetCodeStringWithInvalidStringReturnsSpacesForSome(t *testing.T){

	err,got := getCodeString("df#t")
	want := "-.. ..-. -"

	if string(got) != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	if err != nil {
		t.Errorf("got error %q", err)
	}
}