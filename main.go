/*
Реализуйте паттерн-конвейер:
Программа принимает числа из стандартного ввода в бесконечном цикле и передаёт число в горутину.
Квадрат: горутина высчитывает квадрат этого числа и передаёт в следующую горутину.
Произведение: следующая горутина умножает квадрат числа на 2.
При вводе «стоп» выполнение программы останавливается.
Советы и рекомендации: воспользуйтесь небуферизированными каналами и waitgroup.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	en := enterNumber(&wg)
	sn := squareNumber(&wg, en)
	doubleNumber(&wg, sn)
	wg.Wait()
}

func enterNumber(wg *sync.WaitGroup) chan int {
	outChan := make(chan int)
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
			close(outChan)
		}()
		fmt.Println("Введите цифру, либо `стоп` для выхода: ")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			if value, err := strconv.Atoi(scanner.Text()); err == nil {
				fmt.Println("Ввод: ", value)
				outChan <- value
			} else if scanner.Text() == "стоп" {
				fmt.Printf("Выход из программы")
				break
			} else {
				fmt.Println("Некорректный ввод, повторите попытку")
				continue
			}
		}
	}()
	return outChan
}

func squareNumber(wg *sync.WaitGroup, inChan chan int) chan int {
	outChan := make(chan int)
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
			close(outChan)
		}()
		for val := range inChan {
			fmt.Println("Квадрат: ", val*val)
			outChan <- val * val
		}
	}()
	return outChan
}

func doubleNumber(wg *sync.WaitGroup, inChan chan int) {
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
		}()
		for val := range inChan {
			fmt.Println("Произведение: ", val*2)
		}
	}()
}
