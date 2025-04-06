package domain





type IAnimal interface {
	GetID() int

	GetName() string

	GetType() string

	GetPrice() float64

	GetOwnerID() int

	SetName(name string)

	SetType(animalType string)

	SetPrice(price float64)

	SetOwnerID(ownerID int)

	Speak() string

	SetNickname(nickname string)

	GetNickname() string
}
