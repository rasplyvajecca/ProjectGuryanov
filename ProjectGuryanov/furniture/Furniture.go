package furniture

type Furniture struct {
	Name  string
	Count int
	Size  float64
}

func AddFurniture() []Furniture {
	return []Furniture{
		{Name: "Шкаф", Count: 4, Size: 2},
		{Name: "Стол", Count: 2, Size: 2},
		{Name: "Стул", Count: 15, Size: 2},
		{Name: "Кровать", Count: 5, Size: 2},
		{Name: "Ковер", Count: 5, Size: 2},
	}
}
