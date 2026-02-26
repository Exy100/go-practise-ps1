package main

import (
	"fmt"
	"strings"
	"slices"
)

func main(){
	
	exchangeMoney()
	
}

func removeByValue(slice []string, value string) []string {
    for i, v := range slice {
        if v == value {
            return append(slice[:i], slice[i+1:]...)
        }
    }
    return slice
}

func inputData() (string, float64, string){
	
	defer func(){
		r := recover()
		if r != nil{
			fmt.Println("Пользователь ввел команду 'стоп' для досрочного завершения программы.\n Обмен отменен" )
		}
	}()
	
	currencies := []string{"USD", "EUR", "RUB"}
	inputAsset := ""
	assetAmount := 0.0
	outputAsset := ""
	
	fmt.Println("Для завершения всей программы введите 0")
	fmt.Printf("Введите НАЗВАНИЕ валюты для обмена(%s): \n", strings.Join(currencies, ", "))
	
	for{
		fmt.Scan(&inputAsset)
		inputAsset = strings.ToUpper(strings.TrimSpace(inputAsset))
		if inputAsset == "0" || inputAsset == ""{
			panic("stop")
		}
			
		if !slices.Contains(currencies, inputAsset){
			fmt.Println("Неверные данные - попробуйте еще раз")
			continue
		}else{
			break
		}
		
	}
	
	
	
	fmt.Println("Введите сколько валюты готовы обменять: ")
	
	for{
		fmt.Scan(&assetAmount)
		
		if assetAmount <= 0 {
			panic("stop")
		}
			
		if !slices.Contains(currencies, inputAsset){
			fmt.Println("Неверные данные - попробуйте еще раз")
			continue
		}else{
			break
		}
	}
	
	currencies = removeByValue(currencies, inputAsset)
	fmt.Printf("Введите НАЗВАНИЕ валюты, которую хотите получить (%s): \n", strings.Join(currencies, ", "))
	
	for{
		
		fmt.Scan(&outputAsset)
		
		if outputAsset == "0" || outputAsset == ""{
			panic("stop")
		}
		
		outputAsset = strings.ToUpper(outputAsset)
		if !slices.Contains(currencies, outputAsset){
			fmt.Println("Неверные данные - попробуйте еще раз")
			continue
		}else{
			break
		}
	}
	
	return inputAsset, assetAmount, outputAsset
	
}

func exchangeMoney(){
	
		const usd_eur = 0.9
		const usd_rub = 80.0

		inAsset, amount, outAsset := inputData()

		var amountInUSD float64
		var result float64

		
		switch inAsset {
		case "USD":
			amountInUSD = amount
		case "EUR":
			amountInUSD = amount / usd_eur
		case "RUB":
			amountInUSD = amount / usd_rub
		}

		
		switch outAsset {
		case "USD":
			result = amountInUSD
		case "EUR":
			result = amountInUSD * usd_eur
		case "RUB":
			result = amountInUSD * usd_rub
		}

		fmt.Printf("Результат обмена: %.2f %s\n", result, outAsset)
	
}
