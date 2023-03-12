package gross

var units = map[string]int{
	"quarter_of_a_dozen": 3,
	"half_of_a_dozen":    6,
	"dozen":              12,
	"small_gross":        120,
	"gross":              144,
	"great_gross":        1728,
}

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	if qty, ok := units[unit]; ok {
		bill[item] += qty
		return ok
	}
	return false
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) (ok bool) {
	var cur, qty int

	if cur, ok = bill[item]; !ok {
		return
	}

	if qty, ok = units[unit]; !ok {
		return
	}

	switch v := cur - qty; {
	case v == 0:
		delete(bill, item)
	case v < 0:
		return false
	default:
		bill[item] -= qty
	}
	return ok
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (n int, ok bool) {
	n, ok = bill[item]
	return
}
