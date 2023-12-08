package family

type Family struct {
	Member string
	Age    int
}

func AddFamily() []Family {
	return []Family{
		{Member: "Мать", Age: 45},
		{Member: "Отец", Age: 60},
		{Member: "Сын", Age: 1},
		{Member: "Дочь", Age: 2},
		{Member: "Дочь", Age: 3},
	}
}
