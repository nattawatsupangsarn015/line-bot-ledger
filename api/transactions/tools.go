package transactions

import (
	"example/line-bot-ledger/model"
	"example/line-bot-ledger/utils"
	"fmt"
	"strconv"
	"strings"
)

func stringToFloat(str string, f *float64) error {
	var err error
	*f, err = strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Printf("could not convert string to float64: %v", err)
		return err
	}

	return nil
}

func ConvertTransactionIncome(transaction model.Income) (model.IncomeDecrypt, error) {
	decrypted, err := utils.Decrypt(transaction.UserId.Hex(), transaction.Value)
	if err != nil {
		return model.IncomeDecrypt{}, err
	}

	convertValue := formatNumber(string(decrypted))
	if err != nil {
		return model.IncomeDecrypt{}, err
	}

	return model.IncomeDecrypt{
		ID:          transaction.ID,
		Value:       convertValue,
		Description: transaction.Description,
	}, nil
}

func ConvertTransactionExpense(transaction model.Expense) (model.ExpenseDecrypt, error) {
	decrypted, err := utils.Decrypt(transaction.UserId.Hex(), transaction.Value)
	if err != nil {
		return model.ExpenseDecrypt{}, err
	}

	convertValue := formatNumber(string(decrypted))
	if err != nil {
		return model.ExpenseDecrypt{}, err
	}

	fmt.Println(convertValue)

	return model.ExpenseDecrypt{
		ID:          transaction.ID,
		Value:       convertValue,
		Description: transaction.Description,
	}, nil
}

func formatNumber(number string) string {
	var parts []string
	for len(number) > 3 {
		parts = append(parts, number[len(number)-3:])
		number = number[:len(number)-3]
	}
	parts = append(parts, number)
	for i := 0; i < len(parts)/2; i++ {
		j := len(parts) - i - 1
		parts[i], parts[j] = parts[j], parts[i]
	}
	return strings.Join(parts, ",")
}
