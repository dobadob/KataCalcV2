package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/brandenc40/romannumeral"
	"os"
	"strconv"
	"strings"
)

type expression struct {
	num     [2]int
	mathOp  byte
	isRoman bool
}

func (e *expression) Read() (err error) {
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n') //считываем входные данные
	input = strings.TrimSuffix(input, string(byte(10)))    //обрезаем переход на следующую строку
	input = strings.TrimSuffix(input, string(byte(13)))    //обрезаем "Возврат каретки"
	s := strings.Split(input, " ")                         //создаём срез входных данных

	//обрабатываем простые ошибки
	if len(s) < 3 {
		return errors.New("строка не является математической операцией")
	}
	if len(s) > 3 || !(s[1] == "+" || s[1] == "-" || s[1] == "*" || s[1] == "/") {
		return errors.New("формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
	}

	e.mathOp = []byte(s[1])[0] //кладём в mathOp знак мат.операции
	s = append(s[:1], s[2])    //оставляем в срезе только два операнда

	//переводим римские цифры в арабские, обрабатываем ошибку NaN
	var numsIsRoman = [2]bool{false, false}
	for i, v := range s {
		e.num[i], err = romannumeral.StringToInt(v)
		if err == nil {
			numsIsRoman[i] = true
		} else {
			e.num[i], err = strconv.Atoi(v)
			if err != nil {
				return errors.New("минимум один из операндов не является целым числом")
			}
		}
	}

	//разбираемся с системой счисления
	if numsIsRoman == [2]bool{true, true} && numsIsRoman != [2]bool{false, false} {
		e.isRoman = true
	} else if numsIsRoman != [2]bool{false, false} {
		return errors.New("используются одновременно разные системы счисления")
	}

	//проверяем на соответствие диапазону от 1 до 10
	for _, v := range e.num {
		if v < 1 || v > 10 {
			return errors.New("минимум один из операндов не соответствует диапазону от 1 до 10")
		}
	}

	//проверяем на отрицательные римские числа
	if e.isRoman && e.mathOp == '-' && e.num[0]-e.num[1] < 1 {
		return errors.New("в римской системе нет отрицательных чисел")
	}

	return nil
}

func (e *expression) Calculate() string {
	var temp int //временная переменная

	//вычисляем выражение
	switch e.mathOp {
	case '+':
		temp = e.num[0] + e.num[1]
	case '-':
		temp = e.num[0] - e.num[1]
	case '*':
		temp = e.num[0] * e.num[1]
	case '/':
		temp = e.num[0] / e.num[1]
	}

	//возвращаем результат в виде строки
	if e.isRoman {
		res, _ := romannumeral.IntToString(temp)
		return res
	} else {
		return strconv.Itoa(temp)
	}
}

func main() {
	e := new(expression)
	err := e.Read()
	if err != nil {
		fmt.Printf("Вывод ошибки, так как %v.\n", err)
		return
	}
	fmt.Println(e.Calculate())
}
