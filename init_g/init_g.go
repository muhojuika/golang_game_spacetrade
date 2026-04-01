package init_g

import (
	"math/rand/v2"
	"space_trade/models"
)

func InitPlanets() []models.Planet {
	items := Init_Items()

	var planets []models.Planet
	for i := range 10 {
		var pl models.Planet
		pl.PlanetID = i
		pl.PlanetName = getPlanetName()
		pl.PlanetStore = nil
		pl.PlanetType = rand.IntN(4)

		for i := range items {
			if items[i].ItemType == uint(pl.PlanetType) {
				//c:= rand.IntN((120-80)+80)
				items[i].ItemSellPrice = int(float64(rand.IntN(120-80)+80) * 1.2)
				items[i].ItemBuyPrice = int(float64(items[i].ItemSellPrice) * 1.1)
			} else {
				items[i].ItemSellPrice = int(float64(rand.IntN(120-80)+80) * 1)
				items[i].ItemBuyPrice = int(float64(items[i].ItemSellPrice) * 1.1)
			}
			pl.PlanetStore = append(pl.PlanetStore, items[i])
		}
		planets = append(planets, pl)
	}
	return planets
}
func InitUser() models.User {
	return models.User{Money: 1000, Position: 0, Fuel: 100, Inventory: nil, INV_SIZE: 10}
}
func getPlanetName() string {
	var str1 [10]string
	str1[0] = "Mir"
	str1[1] = "Res"
	str1[2] = "Suk"
	str1[3] = "Kin"
	str1[4] = "Lun"
	str1[5] = "Arul"
	str1[6] = "Froz"
	str1[7] = "Olen"
	str1[8] = "Gres"
	str1[9] = "Zink"

	var str2 [10]string
	str2[0] = "lings X"
	str2[1] = "xinig"
	str2[2] = "umbra"
	str2[3] = "iviella"
	str2[4] = "ompus"
	str2[5] = "sillos"
	str2[6] = "olimpis"
	str2[7] = "jumbis"
	str2[8] = "urus"
	str2[9] = "firles"

	return str1[rand.IntN(len(str1))] + str2[rand.IntN(len(str2))]
}
func Init_Items() []models.Item {
	var item [5]models.Item
	item[0].ItemID = 0
	item[0].ItemName = "Can with bobs"
	item[0].ItemType = 0

	item[1].ItemID = 1
	item[1].ItemName = "Apple Macbook"
	item[1].ItemType = 1

	item[2].ItemID = 2
	item[2].ItemName = "Apple Iphone"
	item[2].ItemType = 1

	item[3].ItemID = 3
	item[3].ItemName = "TV"
	item[3].ItemType = 2

	item[4].ItemID = 4
	item[4].ItemName = "Pizza"
	item[4].ItemType = 3

	return item[:]
}
