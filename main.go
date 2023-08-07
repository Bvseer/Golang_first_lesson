package main

import (
	"fmt"
	"github.com/Shemistan/Lesson_2/auth"
	"github.com/Shemistan/Lesson_2/models"
)

const (
	EXIT         = "exit"
	AUTH         = "auth"
	REGISTRATION = "registration"
	ADD_PRODUCT  = "add_product"
	ORDERS_LIST  = "orders_list"
)

func main() {

	var command string

	var usersList []models.User
	d
	productsList := make([]string, 0, 10)

	_ = productsList

	for command != EXIT {
		fmt.Println("Введите команду") // Сделать красивый вывод, вывести список команд на этом шаге
		fmt.Printf("1. %s \n2. %s \n3. %s \n", EXIT, AUTH, REGISTRATION)

		fmt.Scan(&command)

		switch command {
		case EXIT:
			break
		case REGISTRATION:
			auth.Registrate(usersList)
		case AUTH:
			auth.Authenticate()
		}
	}
}

// Реализовать следующие АПИ
// add_product - добавить продукт который вводиться с консоли в productList
// order - выводит сообщение что вы купили и очищает корзину
