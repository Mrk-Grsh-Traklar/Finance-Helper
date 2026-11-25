package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Account struct {
	NameAc    string
	BalanceAc float64
}

type User struct {
	Name          string
	Balance       int
	ProcentStavka float32
	ccounts       []Account
}

func SaveUser(User User) error {
	data, err := json.MarshalIndent(User, "", "    ")
	if err != nil {
		return err
	}
	err = os.WriteFile("user.json", data, 0644)
	if err != nil {
		return err
	}
	return nil

}

func addMoney(user *User, amount int) {
	if amount > 0 {
		user.Balance += amount
		fmt.Println("успешное пополнение баланса")

	} else {
		fmt.Println("Ошибка, нельзя пополнить на отрицательную сумму")

	}
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

	monthProcent := float64(user.Balance) * (float64(user.ProcentStavka) / 100 / 12)
	fmt.Println("через месяц вы получите: ", float64(user.Balance)+monthProcent, "(", "+", monthProcent, ")")

	yearProcent := float64(user.Balance) * float64(user.ProcentStavka) / 100
	fmt.Println("через год вы получите: ", float64(user.Balance)+yearProcent, "(", "+", yearProcent, ")")
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
			fmt.Scan(&user.ProcentStavka)
			procent(&user)
		case 5:
			err := SaveUser(user)

			if err != nil {
				fmt.Println("Ошибка сохранения:", err)
			} else {
				fmt.Println("Данные сохранены!")
			}

			break MENU

		default:
			fmt.Println("Неверный ввод, попробуйте снова.")

		}

	}

}
