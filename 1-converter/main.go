package main

import "fmt"

func main(){
	// usd := 100.0
	const usd_eur = 0.9
	const usd_rub = 80.0
	//git remote add origin https://github.com/ТВОЕ_ИМЯ/ИМЯ_РЕПОЗИТОРИЯ.git
	const eur_rub = usd_rub / usd_eur
	fmt.Println(eur_rub * 100.0)
	
}