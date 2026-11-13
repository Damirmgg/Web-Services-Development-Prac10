package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func money(amount float64) string {
	whole, fraction, _ := strings.Cut(strconv.FormatFloat(amount, 'f', 2, 64), ".")
	for i := len(whole) - 3; i > 0; i -= 3 {
		whole = whole[:i] + " " + whole[i:]
	}
	if fraction != "00" {
		whole += "," + fraction
	}
	return whole + " тг"
}

func main() {
	sales := DemoSales()
	discounts := [3]float64{5, 10, 15}

	fmt.Println("============ SALES ANALYZER ============")
	fmt.Println("ID | Товар | Категория | Цена | Количество | Сумма")
	for _, sale := range sales {
		fmt.Printf("%d | %s | %s | %s | %d шт. | %s\n",
			sale.ID, sale.Product, sale.Category, money(sale.Price), sale.Quantity,
			money(sale.Price*float64(sale.Quantity)))
	}
	fmt.Printf("\nКоличество записей: %d\n", len(sales))
	fmt.Printf("Продано единиц товара: %d\n", TotalQuantity(sales))
	fmt.Println("Общая выручка:", money(CalculateRevenue(sales)))
	fmt.Println("Средняя сумма продажи:", money(AverageSale(sales)))

	top := GetTopProduct(sales)
	counts := QuantityByProduct(sales)
	fmt.Printf("\nСамый продаваемый товар:\n%s — %d шт.\n", top, counts[top])

	fmt.Println("\nВыручка по категориям:")
	revenue := RevenueByCategory(sales)
	var categories []string
	for category := range revenue {
		categories = append(categories, category)
	}
	sort.Strings(categories)
	for _, category := range categories {
		fmt.Printf("%s: %s\n", category, money(revenue[category]))
	}

	discount := discounts[1]
	oldPrice := sales[0].Price
	oldRevenue := CalculateRevenue(sales)
	ApplyDiscount(&sales[0], discounts[1])
	fmt.Printf("\nПрименение скидки %.0f%% к записи ID %d:\n", discount, sales[0].ID)
	fmt.Println("Товар:", sales[0].Product)
	fmt.Println("Старая цена:", money(oldPrice))
	fmt.Println("Новая цена:", money(sales[0].Price))
	fmt.Println("Выручка после скидки:", money(CalculateRevenue(sales)))
	fmt.Println("Снижение выручки:", money(oldRevenue-CalculateRevenue(sales)))
}
