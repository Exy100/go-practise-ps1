package main

import "fmt"

func main(){

	
}

func inputData() (float64){
	number := 0.0
	fmt.Println("Введите сколько валюты готовы обменять: ")
	fmt.Scan(&number)
	return number
}

func exchangeMoney(inputAsset string, inputAssetAmount float64, outputAsset string){
	// usd := 100.0
	// const usd_eur = 0.9
	// const usd_rub = 80.0
	// //git remote add origin https://github.com/ТВОЕ_ИМЯ/ИМЯ_РЕПОЗИТОРИЯ.git
	// const eur_rub = usd_rub / usd_eur
	// fmt.Println(eur_rub * 100.0)
}