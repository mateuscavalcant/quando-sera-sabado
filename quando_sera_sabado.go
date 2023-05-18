package main

import (
	"fmt"
	"time"
)

func proximoSabado() {
	fmt.Println("- Em qual dia do mês?")
	hoje := time.Now()
	diaAtual := hoje.Weekday()
	diasRestantes := time.Saturday - diaAtual

	if diasRestantes <= 0 {
		diasRestantes += 7
	}

	dataProximoSabado := hoje.AddDate(0, 0, int(diasRestantes))
	dataDia := dataProximoSabado.Format("02")
	nomeMes := obterNomeMesEmPortugues(dataProximoSabado.Month())

	fmt.Println("- O próximo sábado será no dia", dataDia, "de", nomeMes, ".")

}

func obterNomeMesEmPortugues(m time.Month) string {
	mesesEmPortugues := map[time.Month]string{
		time.January:   "Janeiro",
		time.February:  "Fevereiro",
		time.March:     "Março",
		time.April:     "Abril",
		time.May:       "Maio",
		time.June:      "Junho",
		time.July:      "Julho",
		time.August:    "Agosto",
		time.September: "Setembro",
		time.October:   "Outubro",
		time.November:  "Novembro",
		time.December:  "Dezembro",
	}

	return mesesEmPortugues[m]
}

func main() {
	fmt.Println("- Quando será sabádo?")
	hoje := time.Now().Weekday()
	switch time.Saturday {
	case hoje + 0:
		fmt.Println("- Hoje.")
	case hoje + 1:
		fmt.Println("- Será amanhã.")
	case hoje + 2:
		fmt.Println("- Em dois dias.")
	default:
		fmt.Println("- Ainda está muito longe...")
		proximoSabado()
	}

}
