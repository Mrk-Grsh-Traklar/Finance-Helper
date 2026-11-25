package main

import "fmt"

type User struct {
	Name          string
	Balance       int
	procentStavka float32
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

func procent(user *User) {
	fmt.Println(user.Balance)

	monthProcent := user.Balance * (int(user.procentStavka) / 12) / 100
	fmt.Println("через месяц вы получите: ", user.Balance+monthProcent, "(", "+", monthProcent, ")")

	yearProcent := user.Balance * int(user.procentStavka) / 100
	fmt.Println("через год вы получите: ", user.Balance+yearProcent, "(", "+", yearProcent, ")")
}

func main() {
	user := User{}

	var menu int
MENU:
	for {
		fmt.Println("Пользователь:", user.Name)
		fmt.Println("Итоговый баланс:", user.Balance)
		fmt.Println("===ВЫБЕРИТЕ ДЕЙСТВИЕ===\n1.Создать профиль\n2.Пополнить счёт\n3.Снять деньги\n4.Расчитать процент\n5.Выход")
		fmt.Scan(&menu)

		switch menu {
		case 1:

			fmt.Println("Введите имя:")
			fmt.Scan(&user.Name)

			fmt.Println("Введите начальный баланс:")
			fmt.Scan(&user.Balance)
		case 2:
			var sumAdd int
			fmt.Println("Введите сумму пополнения:")
			fmt.Scan(&sumAdd)

			addMoney(&user, sumAdd)
		case 3:
			var sumMin int

			fmt.Println("Введите сумму для снятия")
			fmt.Scan(&sumMin)

			withDraw(&user, sumMin)
		case 4:
			fmt.Println("Введите годовую ставку: ")
			fmt.Scan(&user.procentStavka)
			procent(&user)
		case 5:
			break MENU

		default:
			fmt.Println("Неверный ввод, попробуйте снова.")

		}

	}
}
