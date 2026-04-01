package models

type Item struct {
	ItemID        int
	ItemName      string
	ItemSellPrice int
	ItemBuyPrice  int
	ItemType      uint // 0 - electrionics, 1 - food, 2 - medicine, 3 - seed and plants
}

type Planet struct {
	PlanetID    int
	PlanetName  string
	PlanetStore []Item
	PlanetType  int // 0 - commons planets, 1 - rare planets
}

type User struct {
	Money     int
	Position  int
	Fuel      int
	Inventory []Item
	INV_SIZE  int
}
