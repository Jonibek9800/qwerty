package services

import (
	"bank/models"
	"bufio"
	"context"
	"fmt"
	"github.com/jackc/pgx/pgxpool"
	"log"
	"os"
)


const LoginOperation = `Введите логин и пароль:`
const AuthorizedOperation = `1. Показать Баланс
2.Перевод Денег
3.Оплата Услуг
4.История транзакций
5.Добавить Адресс банкомата
6.Выход`

func Authorized(database *pgxpool.Pool, id int64){
	for {
		fmt.Println(AuthorizedOperation)
		fmt.Println(`Выбери команду:`)
		var number int64
		fmt.Scan(&number)
		switch number {
		case 1:
			fmt.Println(`Показываю Баланс`)
			CheckBalance(database, id)
		case 2:
			//TODO: Сделать функцию для перевода денег
			fmt.Println(`Перевод денег`)
		case 3:
			//TODO: СДелать функцию для Оплаты услуги
			fmt.Println(`Оплачиваю услугу`)
		case 4:
			//TODO:Сделать функцию для истории транзакций
			fmt.Println(`Показываю историю транзакций`)
		case 5:
			AddAtm(database)
		case 6:
			return
		default:
			fmt.Println("Некорректный ввод попробуйте еще раз")
		}
	}
	return
}

func CheckBalance(database *pgxpool.Pool, id int64){
	var amount int
	_ = database.QueryRow(context.Background(), `select amount from accounts 
where user_id = ($1)`, id).Scan(
	&amount,
	)
	fmt.Println(amount)
}

func AddAtm(database *pgxpool.Pool) (ok bool) {
	fmt.Println("Enter ATMs address")
	var s string
	fmt.Scan(&s)
	reader := bufio.NewReader(os.Stdin)
	Address, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("Can't read command: %v", err)
		return false
	}
	fmt.Println(s)
	sprintf := fmt.Sprintf("%s %s", s, Address)
	fmt.Println(sprintf)
	_, err = models.AddATM(database, sprintf)
	if err != nil {
		fmt.Println("vse ploxo")
		return false
	}
	fmt.Printf("Был добавлен АТМ по адрессу: %s %s\n",s, Address)
	//_, err := models.AddATM(database, text.Text())
	return true
}

func Login(database *pgxpool.Pool) (ok bool, id int64){
	fmt.Println(LoginOperation)
	var login, password string
	fmt.Println("login:")
	fmt.Scan(&login)
	fmt.Println("password:")
	fmt.Scan(&password)

	var User models.User
	_ = database.QueryRow(context.Background(), `Select *from users where login = ($1) 
and password = ($2)`, login, password).Scan(
	&User.ID,
	&User.Name,
	&User.Surname,
	&User.Age,
	&User.Gender,
	&User.Login,
	&User.Password,
	&User.Remove)
	if User.ID > 0{
		return true, User.ID
	}
	return false, User.ID
}