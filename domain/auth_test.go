package domain

import "testing"

func TestValidate(t *testing.T) {
	credentials := &Credentials{
		Username: "ASDF",
		Password: "asd&f3$4324fdfsf",
		Email:    "asdf@gmail.com",
	}

	if credentials.Validate() == false {
		t.Error("Expected the user to be valid but it returned invalid")
	}

	credentials = &Credentials{
		Password: "asd&f3$4324fdfsf",
		Email:    "asdf@gmail.com",
	}
	if credentials.Validate() == true {
		t.Error("Expected the user to be invalid but it returned valid")
	}

	credentials = &Credentials{
		Username: "ASDF",
		Email:    "asdf@gmail.com",
	}
	if credentials.Validate() == true {
		t.Error("Expected the user to be invalid but it returned valid")
	}

	credentials = &Credentials{
		Username: "ASDF",
		Password: "asdsadf$/&dsad",
	}
	if credentials.Validate() == true {
		t.Error("Expected the user to be invalid but it returned valid")
	}

	credentials = &Credentials{}
	if credentials.Validate() == true {
		t.Error("Expected the user to be invalid but it returned valid")
	}
}
