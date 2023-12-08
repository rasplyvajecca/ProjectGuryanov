package house

import (
	"ProjectGuryanov/ProjectGuryanov/animals"
	"ProjectGuryanov/ProjectGuryanov/family"
	"ProjectGuryanov/ProjectGuryanov/furniture"
	"ProjectGuryanov/ProjectGuryanov/relatives"
	"ProjectGuryanov/ProjectGuryanov/toys"
	"fmt"
)

type House struct {
	Rooms          int
	Area           float64
	HouseFurniture []furniture.Furniture
	HouseFamily    []family.Family
	HouseRelatives []relatives.Relatives
	HouseAnimals   []animals.Animals
	HouseToys      []toys.Toys
}

func NewHouse() House {
	return House{
		Rooms:          3,
		Area:           50.55,
		HouseFurniture: furniture.AddFurniture(),
		HouseFamily:    family.AddFamily(),
		HouseRelatives: relatives.AddRelatives(),
		HouseAnimals:   animals.AddAnimals(),
		HouseToys:      toys.AddToys(),
	}
}

func HouseObjects(house House) {
	fmt.Printf("Новый дом:\n")
	fmt.Printf("Количество комнат: %d\n", house.Rooms)
	fmt.Printf("Площадь: %.2f кв. м\n", house.Area)

	fmt.Println("Мебель:")
	for _, Object := range house.HouseFurniture {
		fmt.Printf("- %s, кол-во: %d, %.2f кв. м\n", Object.Name, Object.Count, Object.Size)
	}

	fmt.Println("Семья:")
	for _, Object := range house.HouseFamily {
		fmt.Printf("- %s, %d\n", Object.Member, Object.Age)
	}

	fmt.Println("Родственники:")
	for _, Object := range house.HouseRelatives {
		fmt.Printf("- %s, %d\n", Object.Member, Object.Age)
	}

	fmt.Println("Животные:")
	for _, Object := range house.HouseAnimals {
		fmt.Printf("- %s, кол-во: %d,\n", Object.Animal, Object.Count)
	}

	fmt.Println("Игрушки:")
	for _, Object := range house.HouseToys {
		fmt.Printf("- %s, кол-во: %d,\n", Object.Toy, Object.Count)
	}
}
