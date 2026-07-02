package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	quantity, exists := units[unit]
    if !exists {
        return false
    }

    itemQuantity, exists := bill[item]
    if !exists {
        bill[item] = quantity
    } else {
        bill[item] = quantity + itemQuantity
    }
    return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
    quantity, exists := units[unit]
    if !exists {
        return false
    }

    itemQuantity, exists := bill[item]
    if !exists {
        return false
    }

    newQuantity := itemQuantity - quantity
    switch {
        case newQuantity < 0:
        	return false

        case newQuantity == 0:
        	delete(bill, item)
        	return true

        default:
        	bill[item] = newQuantity
        	return true
    }
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (quantity int, exists bool) {
	quantity, exists = bill[item]
    return
}
