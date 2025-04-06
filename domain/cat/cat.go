package domain

type Cat struct {
	ID         int
	Name       string
	AnimalType string
	Price      float64
	OwnerID    int
	Nickname   string
}

func (c *Cat) GetID() int {
	return c.ID
}

func (c *Cat) GetName() string {
	return c.Name
}

func (c *Cat) GetType() string {
	return c.AnimalType
}

func (c *Cat) GetPrice() float64 {
	return c.Price
}

func (c *Cat) GetOwnerID() int {
	return c.OwnerID
}

func (c *Cat) SetName(name string) {
	c.Name = name
}

func (c *Cat) SetType(animalType string) {
	c.AnimalType = animalType
}

func (c *Cat) SetPrice(price float64) {
	c.Price = price
}

func (c *Cat) SetOwnerID(ownerID int) {
	c.OwnerID = ownerID
}
