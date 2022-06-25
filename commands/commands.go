package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ELTITO310/goapp/expenses"
	"github.com/fatih/color"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func GetInput() (string, error) {

	d := color.New(color.FgCyan, color.Bold)
	d.Printf("-> ")

	str, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	str = strings.Replace(str, "\r\n", "", 1)

	return str, nil

}

func ShowInConsole(expns []float32) {
	buildString := contentString(expns)

	fmt.Println(buildString)
}

func contentString(expns []float32) string {
	builder := strings.Builder{}

	total, min, max, avarage := expensesDetailt(expns)

	for i, v := range expns {
		builder.WriteString(fmt.Sprintf("Expense: %6.2f\n", v))
		if i == len(expns)-1 {
			fmt.Println("================================")
			builder.WriteString(fmt.Sprintf("Total: %6.2f\n", total))
			builder.WriteString(fmt.Sprintf("Min: %6.2f\n", min))
			builder.WriteString(fmt.Sprintf("Max: %6.2f\n", max))
			builder.WriteString(fmt.Sprintf("Avarage: %6.2f\n", avarage))
		}
	}

	return builder.String()
}

func expensesDetailt(expns []float32) (total, min, max, avarage float32) {
	if len(expns) == 0 {
		return
	}
	total = expenses.Sum(expns...)
	min = expenses.Min(expns...)
	max = expenses.Max(expns...)
	avarage = expenses.Avarage(expns...)
	return
}

func Export(filename string, list []float32) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer f.Close()

	w := bufio.NewWriter(f)

	_, err = w.WriteString(contentString(list))
	if err != nil {
		return err
	}

	return w.Flush()
}
