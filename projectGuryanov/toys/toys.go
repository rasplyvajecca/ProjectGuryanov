package toys

type Toys struct {
	Toy   string
	Count int
}

func AddToys() []Toys {
	return []Toys{
		{Toy: "Конструктор", Count: 1},
		{Toy: "Мягкие игрушки", Count: 6},
		{Toy: "Пазлы", Count: 20},
		{Toy: "Настольные игры", Count: 18},
		{Toy: "Машина", Count: 1},
	}
}
