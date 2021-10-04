package main

import (
	"fmt"
	"strconv"
	"strings"
)

//структура для хранения координат точки, x - по горизонтальной оси, y - по вертикальной
type Point struct {
	x, y int
}

//метод для преобразования координат точки в строковое представление
func (p Point) toStr() string {
	return strconv.Itoa(p.x) + ";" + strconv.Itoa(p.y)
}

//функция для считывания точки из строки
func pFromStr(s string) Point {
	spl := strings.Split(s, ";")
	x, _ := strconv.Atoi(spl[0])
	y, _ := strconv.Atoi(spl[1])
	return Point{x, y}
}

//функция, принимает n - размер поля, k - стартовая позиция, начиная с 0
func CalculatePaths(n, k int) uint64 {
	if k > n-1 || k < 0 {
		panic("Start position is out of range")
	}

	//тут можно пофилософствовать
	if n == 0 || n == 1 {
		return 0
	}
	var sum uint64
	var p, p1, p2 Point
	startPoint := Point{k, 0}
	//создаем map(отображение?) - линию, где ключ - координата точки на текущей линии(в строковом отображении)
	//значение - количество возможных маршрутов до данной точки (1-инициирующее значение для стартовой точки)
	thisLine := map[string]uint64{startPoint.toStr(): 1}
	for i := 1; i < n; i++ {
		//создаем новую линию и заполняем ее
		nextLine := make(map[string]uint64)
		for c, v := range thisLine {
			p = pFromStr(c)
			p1.x, p1.y = p.x-1, p.y+1
			p2.x, p2.y = p.x+1, p.y+1
			//проверяем, есть ли такая точка, если да, то
			//для каждой новой точки находим количество маршрутов до нее,
			//складывая количество маршрутов до тех точек с предыдущей линии,
			//из которых можно попасть в новую точку
			if p1.x >= 0 {
				nextLine[p1.toStr()] += v
			}
			if p2.x <= n-1 {
				nextLine[p2.toStr()] += v
			}
		}
		thisLine = nextLine
	}
	//проходимся по всем точкам на последней линии, складывая количество маршрутов до каждой из них
	for _, v := range thisLine {
		sum += v
	}
	return sum
}

func main() {
	var n, k int
	fmt.Print("Введите размер стороны поля (n > 0): ")
	for {
		fmt.Scan(&n)
		if n < 1 {
			fmt.Print("Поля такого размера не существует, попробуйте еще раз: ")
		} else {
			break
		}
	}
	fmt.Print("Введите стартову позицию (начиная с 0):")
	for {
		fmt.Scan(&k)
		if k < 0 || k > n-1 {
			fmt.Print("Некорректная стартовая позиция, попробуйте еще раз: ")
		} else {
			break
		}
	}
	fmt.Printf("Количество возможных маршрутов: %v ", CalculatePaths(n, k))
}
