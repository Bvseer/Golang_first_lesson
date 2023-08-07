package auth

import (
	"errors"
	"fmt"
	"github.com/Shemistan/Lesson_2/models"
	"strings"
)

func Registrate(usersList []models.User) (string, error) {

	var input string

	fmt.Println("Введите логин и пароль через underscore login_password")

	fmt.Scan(&input) // Сделать так, что бы выводил сообщение, если пользователь уже существует

	loginAndPassword := strings.Split(input, "_")

	if len(loginAndPassword) != 2 {
		return "", errors.New("Incorrect input value")
	}

	for k := range usersList {
		if loginAndPassword[0] == usersList[k].Login {
			return "", errors.New("User already exists")
		}
	}

	newUser := models.User{Login: loginAndPassword[0], Password: loginAndPassword[1]}
	usersList = append(usersList, newUser)

	return "Пользователь успешно добавлен", nil
}

func Authenticate(usersList []models.User) {

	var command string

	fmt.Println("Введите логин и пароль в таком виде login_password")
	fmt.Scan(&command)
	loginAndPassword := strings.Split(command, "_")
	if len(loginAndPassword) != 2 {
		fmt.Println("Введены неправильные данные")
	}

	success := false
	for _, v := range usersList {
		if v.Login == loginAndPassword[0] && v.Password == loginAndPassword[1] {
			success = true
		}
	}
	if !success {
		fmt.Println("Вы не зарегистрированны")
	}

	fmt.Println("Добро пожаловать")
	fmt.Printf("Введите %s, чтобы добавить продукт и %s чтобы увидеть список ваших продуктов", ADD_PRODUCT, ORDERS_LIST)
	var msg string
	for msg != EXIT {
		fmt.Scan(&msg)

		switch msg {
		case ADD_PRODUCT:
		case ORDERS_LIST:
		}
	}
}
