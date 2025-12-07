package gross

var unitsMap = map[string]int{
    "quarter_of_a_dozen": 3,
    "half_of_a_dozen":    6,
    "dozen":              12,
    "small_gross":        120,
    "gross":              144,
    "great_gross":        1728,
}
// Units stores the Gross Store unit measurements.
func Units() map[string]int {
    return unitsMap
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	q, ok := units[unit]
	if !ok {
		return false
	}

	bill[item] += q
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	qty, ok := units[unit]
	if !ok {
		return false
	}
	current, ok := bill[item]
	if !ok {
		return false
	}

	newQty := current - qty
	switch {
	case newQty < 0:
		return false
	case newQty == 0:
		delete(bill, item)
	default: // newQty > 0
		bill[item] = newQty
	}
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	qty, ok := bill[item]
	return qty, ok
}
