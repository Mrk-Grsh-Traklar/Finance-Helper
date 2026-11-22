package main

import "fmt"

type User struct {
	Name    string
	Balance int
}

func addMoney(user *User, amount int) {
	user.Balance += amount
}

func withDraw(user *User, Drawamount int) {
	if user.Balance >= Drawamount && Drawamount > 0 {
		user.Balance -= Drawamount
		fmt.Println("успешное снятие денег")

	} else if Drawamount <= 0 {
		fmt.Println("Ошибка, нельзя снять отрицательную сумму")
	} else {
		fmt.Println("Ошибка, не хватвет средств")
	}

}

func main() {
	user := User{}

	var sumAdd int
	var sumMin int

	fmt.Println("Введите имя:")
	fmt.Scan(&user.Name)

	fmt.Println("Введите начальный баланс:")
	fmt.Scan(&user.Balance)

	fmt.Println("Введите сумму пополнения:")
	fmt.Scan(&sumAdd)

	addMoney(&user, sumAdd)

	fmt.Println("Введите сумму для снятия")
	fmt.Scan(&sumMin)

	withDraw(&user, sumMin)

	fmt.Println("Пользователь:", user.Name)
	fmt.Println("Итоговый баланс:", user.Balance)
}
