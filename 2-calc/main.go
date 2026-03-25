package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"slices"
)

func main() {
	r, err := calc()
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(r)	
	}
	
}

func calc() (float64, error) {
	fmt.Println("Выберите тип операции(укажите только название): 1) AVG, 2) SUM, 3) MED")
	choice := ""
	
	for{
		fmt.Scan(&choice)
		strings.ToUpper(choice)
		
		if slices.Contains([]string{"AVG", "SUM", "MED", "STOP"}, choice){
			break
		}else{
			fmt.Println("нет такого вариант - попробуйте снова")
		}

	}
	

	fmt.Println("Теперь укажите через запятую и пробел числа:")
	input_numbers_str := ""
	fmt.Scanln(&input_numbers_str)

	parts := strings.Split(input_numbers_str, ",")
	numbers := make([]int, 0, len(parts))

	for i, p := range parts {
		numStr := strings.TrimSpace(p)

		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Printf("Ошибка в элементе %d (%q): пропускаем\n", i, numStr)
			continue
		}

		numbers = append(numbers, num)
	}

	if len(numbers) == 0 {
		fmt.Println("Не удалось получить ни одного числа")
		return 0, fmt.Errorf("could not parse any numbers")
	}

	if choice == "AVG" {
		sum := 0
		for _, v := range numbers {
			sum += v
		}
		return float64(sum) / float64(len(numbers)), nil
	} else if choice == "SUM" {
		sum := 0
		for _, v := range numbers {
			sum += v
		}
		return float64(sum), nil

	} else if choice == "MED" {

		sort.Ints(numbers)
		mid := len(numbers) / 2

		var med float64
		if len(numbers)%2 == 0 {
			med = (float64(numbers[mid-1]) + float64(numbers[mid])) / 2.0
		} else if len(numbers)%2 == 1 {
			med = float64(numbers[mid])
		}

		return float64(med), nil
	}

	return 0.0, fmt.Errorf("Error")
}
