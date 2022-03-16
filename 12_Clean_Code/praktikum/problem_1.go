package main

type user struct { // struct awali Uppercase
	id       int
	username int // username int ??
	password int
}

type userservice struct { // menggunakan Camel case
	t []user // tidak jelas t itu apa
}

func (u userservice) getallusera() []user { //penamaan fungsi lebih baik Camel case
	return u.t
}

func (u userservice) getuserbyid(id int) user { // penamaan fungsi lebih baik Camel case
	for _, r := range u.t { // variable r tidak jelas
		if id == r.id {
			return r
		}
	}
	return user{}
}
