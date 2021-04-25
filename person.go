package service

type person struct {
	name string
	age  int
}

func initPerson() person {
	m := person{
		name: "tuyentg",
		age:  32,
	}
	return m
}
