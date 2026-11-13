package main

type Sale struct {
	ID       int
	Product  string
	Category string
	Price    float64
	Quantity int
}

func CalculateRevenue(sales []Sale) float64 {
	var revenue float64
	for _, sale := range sales {
		revenue += sale.Price * float64(sale.Quantity)
	}
	return revenue
}

func TotalQuantity(sales []Sale) int {
	var count int
	for _, sale := range sales {
		count += sale.Quantity
	}
	return count
}

func QuantityByProduct(sales []Sale) map[string]int {
	counts := make(map[string]int)
	for _, sale := range sales {
		counts[sale.Product] += sale.Quantity
	}
	return counts
}

func GetTopProduct(sales []Sale) string {
	counts := QuantityByProduct(sales)
	best, maxCount := "", 0
	// Обходим исходный срез: при равенстве выбираем первое встретившееся название.
	for _, sale := range sales {
		if best == "" || counts[sale.Product] > maxCount {
			best, maxCount = sale.Product, counts[sale.Product]
		}
	}
	return best
}

func RevenueByCategory(sales []Sale) map[string]float64 {
	revenue := make(map[string]float64)
	for _, sale := range sales {
		revenue[sale.Category] += sale.Price * float64(sale.Quantity)
	}
	return revenue
}

func AverageSale(sales []Sale) float64 {
	if len(sales) == 0 {
		return 0
	}
	return CalculateRevenue(sales) / float64(len(sales))
}

// Указатель позволяет изменить цену в исходном элементе среза.
func ApplyDiscount(sale *Sale, discount float64) {
	if sale == nil || !(discount >= 0 && discount <= 100) {
		return
	}
	sale.Price *= 1 - discount/100
}

func DemoSales() []Sale {
	return []Sale{
		{1, "Ноутбук", "Электроника", 350000, 2},
		{2, "Мышь", "Аксессуары", 5000, 10},
		{3, "Монитор", "Электроника", 85000, 3},
		{4, "Клавиатура", "Аксессуары", 12000, 5},
		{5, "Наушники", "Аксессуары", 25000, 4},
		{6, "Ноутбук", "Электроника", 350000, 1},
	}
}
