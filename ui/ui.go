package ui

import (
	"fmt"
	"space_trade/game_funcs"
	"space_trade/models"
	"strconv"
	"time"
)

func ShowPlanetList(plantes []models.Planet, user_position int) {
	fmt.Println("Planets list")
	for i := range plantes {
		if i == user_position {
			fmt.Println(strconv.Itoa(i+1) + ") Name: " + plantes[i].PlanetName + "   Type: " + strconv.Itoa(plantes[i].PlanetType) + " <--- you here")
			continue
		}
		fmt.Println(strconv.Itoa(i+1) + ") Name: " + plantes[i].PlanetName + "   Type: " + strconv.Itoa(plantes[i].PlanetType))
	}
}
func WaitInput(timeDelay int64) {
	time.Sleep(time.Duration(timeDelay) * time.Millisecond)
	fmt.Println("Please enter any button...")
	fmt.Scanln()

}
func ShowPlanetMenu(planet models.Planet, user models.User) int {
	var option int
	fmt.Println("-------------------------")
	fmt.Println("-------------------------")
	fmt.Println("You on - " + planet.PlanetName + " planet")
	fmt.Println("-------------------------")
	fmt.Println("Please choose option...")
	fmt.Println("1) Show planet shop")
	fmt.Println("2) Show spaceship storage")
	fmt.Println("3) Escape from planet")
	fmt.Println("4) Exit game")

	for {
		fmt.Println("Enter option number")
		_, err := fmt.Scanln(&option)
		if err == nil {
			if option > 0 && option < 5 {
				return option
			} else {
				fmt.Println("Wrong number, please repeat")
			}

		}
	}
}
func ShowPlanetShop(user models.User, planet models.Planet) int {
	option := 0
	fmt.Println("-------------------------------------")
	fmt.Println("-----------Planet-Shop---------------")
	fmt.Println("-------------------------------------")
	fmt.Println("You balance --- " + strconv.Itoa(user.Money))

	for i := range planet.PlanetStore {
		fmt.Println("#" + strconv.Itoa(i+1) + " - '" + planet.PlanetStore[i].ItemName + "'  // Buy price - <" + strconv.Itoa(planet.PlanetStore[i].ItemBuyPrice) + "> // Sell price - <" + strconv.Itoa(planet.PlanetStore[i].ItemSellPrice) + ">")
	}
	fmt.Println("-------------------------")
	fmt.Println("Please choose option...")
	fmt.Println("1) Buy items")
	fmt.Println("2) Sell items")
	fmt.Println("3) Back")

	for {
		fmt.Println("Enter option number")
		_, err := fmt.Scanln(&option)
		if err == nil {
			if option > 0 && option < 4 {
				return option
			} else {
				fmt.Println("Wrong number, please repeat")
			}

		}
	}

}
func ShowBuyMenu(store []models.Item) (int, int) {
	var ind, count int
	for {
		fmt.Println("Enter item number")
		_, err := fmt.Scanln(&ind)
		if err == nil {
			if ind > 0 && ind < (len(store)+1) {
				break
			} else {
				fmt.Println("Wrong number, please repeat")
			}

		}
	}
	fmt.Println("You choosed - " + store[ind-1].ItemName + " buy price - " + strconv.Itoa(store[ind-1].ItemBuyPrice))
	for {
		fmt.Println("Enter count number")
		_, err := fmt.Scanln(&count)
		if err == nil {
			if count > 0 {
				break
			} else {
				fmt.Println("Wrong number, please repeat")
			}

		}
	}
	return ind, count
}
func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}
func ShowStorage(user models.User, planet models.Planet) {
	fmt.Println("-------------------------------------")
	fmt.Println("---------Spaceship-Storage-----------")
	fmt.Println("-------------------------------------")

	for i := range 10 {
		if i < len(user.Inventory) {
			fmt.Println("Slot #" + strconv.Itoa((i + 1)) + "  --- " + user.Inventory[i].ItemName + " --- Sell Price on this Planet <" + strconv.Itoa(game_funcs.FoundSellPrice(user.Inventory[i].ItemID, planet)) + ">")

		} else {
			fmt.Println("Slot #" + strconv.Itoa((i + 1)) + "  --- EMPTY SLOT  ---")

		}
	}
}
func ShowSellMenu(user models.User, planet models.Planet) int {
	option := 0
	ShowStorage(user, planet)
	fmt.Println("-------------------------")
	fmt.Println("Please choose item...")
	fmt.Println("0) Back")

	for {
		fmt.Println("Enter option number")
		_, err := fmt.Scanln(&option)
		if err == nil {
			if option > 0 && option < 11 {
				fmt.Println("You choosed - " + user.Inventory[option-1].ItemName + " sell price - " + strconv.Itoa(game_funcs.FoundSellPrice(option-1, planet)))
				return option
			}
			if option == 0 {
				return 0

			} else {
				fmt.Println("Wrong number, please repeat")
			}

		}
	}
}
func ShowChangePlanet(planets []models.Planet) int {
	option := 0
	fmt.Println("-------------------------")
	fmt.Println("Please choose planet...")
	fmt.Println("0) Back")
	for {
		fmt.Println("Enter planet ID")
		_, err := fmt.Scanln(&option)
		if err == nil {
			if option > 0 && option <= len(planets) {
				fmt.Println("You choosed - " + planets[option-1].PlanetName)
				return option
			} else {
				fmt.Println("Wrong number, please repeat")
			}
			if option == 0 {
				return 0

			}

		}
	}
}
