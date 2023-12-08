package house

import (
	"ProjectGuryanov/projectGuryanov/animals"
	"ProjectGuryanov/projectGuryanov/family"
	"ProjectGuryanov/projectGuryanov/furniture"
	"ProjectGuryanov/projectGuryanov/relatives"
	"ProjectGuryanov/projectGuryanov/toys"
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
	for _, object := range house.HouseFurniture {
		fmt.Printf("- %s, кол-во: %d, %.2f кв. м\n", object.Name, object.Count, object.Size)
	}

	fmt.Println("Семья:")
	for _, object := range house.HouseFamily {
		fmt.Printf("- %s, %d\n", object.Member, object.Age)
	}

	fmt.Println("Родственники:")
	for _, object := range house.HouseRelatives {
		fmt.Printf("- %s, %d\n", object.Member, object.Age)
	}

	fmt.Println("Животные:")
	for _, object := range house.HouseAnimals {
		fmt.Printf("- %s, кол-во: %d,\n", object.Animal, object.Count)
	}

	fmt.Println("Игрушки:")
	for _, object := range house.HouseToys {
		fmt.Printf("- %s, кол-во: %d,\n", object.Toy, object.Count)
	}
}
