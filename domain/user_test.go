package domain

import "testing"

func TestValidate(t *testing.T) {
	user := &User{
		UserName: "ASDF",
		Password: "asd&f3$4324fdfsf",
		Email:    "asdf@gmail.com",
	}

	if user.Validate() == false {
		t.Error("Expected the user to be valid but it returned invalid")
	}

	user = &User{
		Password: "asd&f3$4324fdfsf",
		Email:    "asdf@gmail.com",
	}
	if user.Validate() == true {
		t.Error("Expected the user to be invalid but it returned valid")
	}

	user = &User{
		UserName: "ASDF",
		Email:    "asdf@gmail.com",
	}
	if user.Validate() == true {
		t.Error("Expected the user to be invalid but it returned valid")
	}

	user = &User{
		UserName: "ASDF",
		Password: "asdsadf$/&dsad",
	}
	if user.Validate() == true {
		t.Error("Expected the user to be invalid but it returned valid")
	}

	user = &User{}
	if user.Validate() == true {
		t.Error("Expected the user to be invalid but it returned valid")
	}
}
