package game_funcs

import (
	"fmt"
	"space_trade/models"
	"strconv"
)

func BuyItem(user models.User, store []models.Item, count int, ind int) (models.User, bool) {
	succsess := false
	if store[ind-1].ItemBuyPrice*count <= user.Money && count <= user.INV_SIZE-len(user.Inventory) {
		for range count {
			user.Inventory = append(user.Inventory, store[ind-1])
		}
		user.Money -= store[ind-1].ItemBuyPrice * count
		succsess = true
		return user, succsess
	} else {
		return user, succsess
	}

}
func SellItem(user models.User, ind int, planet models.Planet) (models.User, bool) {
	status := false
	var u = user
	for i := range user.Inventory {
		if ind-1 == i {
			price := FoundSellPrice(user.Inventory[i].ItemID, planet)
			if price != 0 {
				fmt.Println("Price " + strconv.Itoa(price) + " balance " + strconv.Itoa(u.Money) + " summary " + strconv.Itoa((u.Money + price)))
				u.Money += price
				u.Inventory = DeleteItem(user.Inventory, i)
				status = true

			}
		}
	}
	return u, status
}
func DeleteItem(inv []models.Item, ind int) []models.Item {
	var new_inv []models.Item
	for i := range inv {
		if i != ind {
			new_inv = append(new_inv, inv[i])
		}
	}
	return new_inv
}
func FoundSellPrice(idItem int, planet models.Planet) int {
	for i := range planet.PlanetStore {
		if planet.PlanetStore[i].ItemID == idItem {
			return planet.PlanetStore[i].ItemSellPrice
		}
	}
	return 0
}
func ChangePlanet(user models.User, planets []models.Planet, ind_planet int) (models.User, bool) {
	status := false

	fmt.Println("You fly to --- " + planets[ind_planet-1].PlanetName + " planet")
	user.Position = ind_planet - 1

	return user, status
}
