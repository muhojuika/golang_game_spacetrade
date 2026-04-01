package main

import (
	"fmt"
	"os"
	"space_trade/game_funcs"
	"space_trade/init_g"
	"space_trade/ui"
)

func main() {
	user := init_g.InitUser()
	gamePlanets := init_g.InitPlanets()
	//ui.ShowPlanetList(gamePlanets, user.Position)
	for {
		ui.ClearScreen()
		option := ui.ShowPlanetMenu(gamePlanets[user.Position], user)
		switch option {
		case 1: //shop
			ui.ClearScreen()
			shp_o := ui.ShowPlanetShop(user, gamePlanets[user.Position])
			switch shp_o {
			case 1: // buy items
				fmt.Println("buy")
				ind, count := ui.ShowBuyMenu(gamePlanets[user.Position].PlanetStore)
				u, st := game_funcs.BuyItem(user, gamePlanets[user.Position].PlanetStore, count, ind)
				if st == true {
					user = u
					fmt.Println("*** Succsess ***")
					ui.WaitInput(1000)

				} else {
					fmt.Println("*** Faild ***")
					ui.WaitInput(1000)
				}
			case 2: // sell items
				fmt.Println("selly")
				ind := ui.ShowSellMenu(user, gamePlanets[user.Position])
				if ind != 0 {
					u, st := game_funcs.SellItem(user, ind, gamePlanets[user.Position])
					if st == true {
						user = u
						fmt.Println("*** Succsess ***")
						ui.WaitInput(1000)

					} else {
						fmt.Println("*** Faild ***")
						ui.WaitInput(1000)
					}

				} else {
					fmt.Println("Back to menu")
					ui.WaitInput(1000)
					continue
				}
			case 3: // back to main menu
				fmt.Println("Back to menu")
				ui.WaitInput(1000)
				continue
			}
		case 2: // inventory
			ui.ClearScreen()
			fmt.Println("Shop spaceship storage")
			ui.ShowStorage(user, gamePlanets[user.Position])
			ui.WaitInput(1000)
		case 3: // list planet
			ui.ShowPlanetList(gamePlanets, user.Position)
			op := ui.ShowChangePlanet(gamePlanets)
			if op != 0 {
				user, _ = game_funcs.ChangePlanet(user, gamePlanets, op)
			} else {
				fmt.Println("Back to menu")
				ui.WaitInput(1000)
				continue
			}
			ui.WaitInput(1000)
		case 4: // close game
			fmt.Println("Exit")
			os.Exit(0)
		}
	}

}
