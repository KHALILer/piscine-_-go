package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	preptime := 0
	if order != "burger" && order != "chips" && order != "nuggets" {
		preptime = 404
	}
	if order == "burger" {
		preptime = 15
	}
	if order == "chips" {
		preptime = 10
	}
	if order == "nuggets" {
		preptime = 12
	}

	return preptime
}
