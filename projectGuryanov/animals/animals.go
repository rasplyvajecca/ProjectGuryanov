package animals

type Animals struct {
	Animal string
	Count  int
}

func AddAnimals() []Animals {
	return []Animals{
		{Animal: "Хомяки", Count: 5},
		{Animal: "Кошка", Count: 1},
		{Animal: "Собаки", Count: 2},
		{Animal: "Морская свинка", Count: 1},
		{Animal: "Змеи", Count: 3},
	}
}
