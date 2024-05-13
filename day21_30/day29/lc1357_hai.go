package main

type Cashier struct {
	ProductPriceMap      map[int]int
	Cur                  int
	DiscountTriggerPoint int
	Discount             float64
}

func Constructor(n int, discount int, products []int, prices []int) Cashier {
	data := make(map[int]int, len(products))
	for i := 0; i < len(products); i++ {
		data[products[i]] = prices[i]
	}
	return Cashier{
		ProductPriceMap:      data,
		Cur:                  0,
		DiscountTriggerPoint: n,
		Discount:             float64(discount),
	}
}

func (this *Cashier) GetBill(product []int, amount []int) float64 {
	this.Cur++
	discountRate := float64(1.0)
	if this.Cur == this.DiscountTriggerPoint {
		this.Cur = 0
		discountRate = 1 - this.Discount/100
	}
	var bill int
	for i := 0; i < len(product); i++ {
		bill += this.ProductPriceMap[product[i]] * amount[i]
	}
	return float64(bill) * discountRate
}

func main() {
	cashier := Constructor(3, 50, []int{1, 2, 3, 4, 5, 6, 7}, []int{100, 200, 300, 400, 300, 200, 100})
	product := []int{1, 2}
	amount := []int{1, 2}
	println(cashier.GetBill(product, amount)) // 500.0
	// product = []int{3, 7}
}
