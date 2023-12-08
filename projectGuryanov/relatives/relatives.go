package relatives

type Relatives struct {
	Member string
	Age    int
}

func AddRelatives() []Relatives {
	return []Relatives{
		{Member: "Прабабушка", Age: 100},
		{Member: "Прадедушка", Age: 95},
		{Member: "Бабушка", Age: 80},
		{Member: "Дедушка", Age: 80},
		{Member: "Тетя", Age: 40},
	}
}
