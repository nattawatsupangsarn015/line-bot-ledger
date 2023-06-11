package transactions

import (
	"example/line-bot-ledger/controller"
	"example/line-bot-ledger/model"
	"example/line-bot-ledger/utils"
	"fmt"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetLastestTransactions(lineId string) (string, error) {
	findUser, err := controller.GetUserByLineId(lineId)
	if err != nil {
		return "", err
	}

	if findUser == nil {
		return "ขออภัย ระบบไม่สามารถค้นหาข้อมูลของท่านได้\n กรุณาลองออกจากระบบแล้วเข้าสู่ระบบใหม่อีกครั้งครับ 🙇‍♂️", nil
	}

	var user model.User
	err = utils.ConvertInterfaceToStruct(findUser, &user)
	if err != nil {
		return "", err
	}

	var incomeTransactions []model.Income
	rawIncomeTransactions, err := controller.GetLastestIncome(user.ID)
	if err != nil {
		return "", err
	}

	err = utils.ConvertInterfaceToStruct(rawIncomeTransactions, &incomeTransactions)
	if err != nil {
		return "", err
	}

	var expenseTransactions []model.Expense
	rawExpenseTransactions, err := controller.GetLastestExpense(user.ID)
	if err != nil {
		return "", err
	}

	err = utils.ConvertInterfaceToStruct(rawExpenseTransactions, &expenseTransactions)
	if err != nil {
		return "", err
	}

	if len(incomeTransactions) == 0 && len(expenseTransactions) == 0 {
		return "ท่านยังไม่มีการสร้างรายการ 🙇‍♂️\nท่านสามารถพิม \"how-to-use\" เพื่อดูวิธีการใช้งาน", nil
	}

	replyText := ""
	var decryptedIncomeTransactions []model.IncomeDecrypt
	var decryptedExpenseTransactions []model.ExpenseDecrypt

	for _, s := range incomeTransactions {
		decrypted, err := ConvertTransactionIncome(s)
		if err != nil {
			return "", err
		}
		decryptedIncomeTransactions = append(decryptedIncomeTransactions, decrypted)
	}

	for _, s := range expenseTransactions {
		decrypted, err := ConvertTransactionExpense(s)
		if err != nil {
			return "", err
		}
		decryptedExpenseTransactions = append(decryptedExpenseTransactions, decrypted)
	}

	replyText = "----- รายรับ -----\n"
	for _, s := range decryptedIncomeTransactions {
		replyText = replyText + "\n"
		replyText = replyText + "ID: " + s.ID.Hex() + " จำนวน: " + s.Value + " หัวข้อ: " + s.Description + "\n"
	}

	replyText = replyText + "\n" + "----- รายจ่าย -----"
	for _, s := range decryptedExpenseTransactions {
		replyText = replyText + "\n"
		replyText = replyText + "ID: " + s.ID.Hex() + " จำนวน: " + s.Value + " หัวข้อ: " + s.Description + "\n"
	}

	return replyText, nil
}

func CreateTransactions(lineId string, transaction model.RequestTransactions) (string, error) {
	rawFindUser, err := controller.GetUserByLineId(lineId)
	if err != nil {
		return "", err
	}

	if rawFindUser == nil {
		return "ขออภัย ระบบไม่สามารถค้นหาข้อมูลของท่านได้\n กรุณาลองออกจากระบบแล้วเข้าสู่ระบบใหม่อีกครั้งครับ 🙇‍♂️", nil
	}

	var findUser model.User
	err = utils.ConvertInterfaceToStruct(rawFindUser, &findUser)
	if err != nil {
		return "", err
	}

	splitData := strings.Split(transaction.Data, "")
	if len(splitData) <= 1 {
		return "ท่านได้กรอกข้อมูลไม่ถูกต้อง \n กรุณาพิม \"how-to-use\" เพื่อตรวจสอบวิธีการกรอกข้อมูลอีกครั้งครับ 🙇‍♂️", nil
	}

	if splitData[0] != "+" && splitData[0] != "-" {
		return "ท่านได้กรอกข้อมูลไม่ถูกต้อง \n กรุณาพิม \"how-to-use\" เพื่อตรวจสอบวิธีการกรอกข้อมูลอีกครั้งครับ 🙇‍♂️", nil
	}

	var valueTransaction string
	if splitData[0] == "+" {
		valueTransaction = strings.Split(transaction.Data, "+")[1]
	} else {
		valueTransaction = strings.Split(transaction.Data, "-")[1]
	}

	var value float64
	err = stringToFloat(valueTransaction, &value)
	if err != nil {
		return "ท่านได้กรอกข้อมูลไม่ถูกต้อง \n กรุณาพิม \"how-to-use\" เพื่อตรวจสอบวิธีการกรอกข้อมูลอีกครั้งครับ 🙇‍♂️", nil
	}

	convertValueToString := strconv.FormatFloat(value, 'f', -1, 64)
	ciphertext, err := utils.Encrypt(findUser.ID.Hex(), convertValueToString)
	if err != nil {
		return "", err
	}

	if splitData[0] == "+" {
		newTransaction := model.Income{
			ID:          primitive.NewObjectID(),
			UserId:      findUser.ID,
			Value:       ciphertext,
			Description: transaction.Description,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		err = controller.CreateIncome(newTransaction)
		if err != nil {
			return "", err
		}

		return "สร้างรายการรายรับเรียบร้อยแล้ว 🎉", nil
	} else {
		newTransaction := model.Expense{
			ID:          primitive.NewObjectID(),
			UserId:      findUser.ID,
			Value:       ciphertext,
			Description: transaction.Description,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		err = controller.CreateExpense(newTransaction)
		if err != nil {
			return "", err
		}

		return "สร้างรายการรายจ่ายเรียบร้อยแล้ว 🎉", nil
	}

	// decrypted, err := utils.Decrypt(findUser.ID.Hex(), ciphertext)
	// if err != nil {
	// 	return "", err
	// }

	// decryptedString := string(decrypted)
	// fmt.Println(decryptedString)

}

func GetBalance(lineId string) (string, error) {
	findUser, err := controller.GetUserByLineId(lineId)
	if err != nil {
		return "", err
	}

	if findUser == nil {
		return "ขออภัย ระบบไม่สามารถค้นหาข้อมูลของท่านได้\n กรุณาลองออกจากระบบแล้วเข้าสู่ระบบใหม่อีกครั้งครับ 🙇‍♂️", nil
	}

	var user model.User
	err = utils.ConvertInterfaceToStruct(findUser, &user)
	if err != nil {
		return "", err
	}

	var incomeTransactions []model.Income
	rawIncomeTransactions, err := controller.GetAllIncome(user.ID)
	if err != nil {
		return "", err
	}

	err = utils.ConvertInterfaceToStruct(rawIncomeTransactions, &incomeTransactions)
	if err != nil {
		return "", err
	}

	var expenseTransactions []model.Expense
	rawExpenseTransactions, err := controller.GetAllExpense(user.ID)
	if err != nil {
		return "", err
	}

	err = utils.ConvertInterfaceToStruct(rawExpenseTransactions, &expenseTransactions)
	if err != nil {
		return "", err
	}

	if len(incomeTransactions) == 0 && len(expenseTransactions) == 0 {
		return "ท่านยังไม่มีการสร้างรายการ 🙇‍♂️\nท่านสามารถพิม \"how-to-use\" เพื่อดูวิธีการใช้งาน", nil
	}

	var decryptedIncomeTransactions []model.IncomeDecrypt
	var decryptedExpenseTransactions []model.ExpenseDecrypt

	for _, s := range incomeTransactions {
		decrypted, err := ConvertTransactionIncome(s)
		if err != nil {
			return "", err
		}
		decryptedIncomeTransactions = append(decryptedIncomeTransactions, decrypted)
	}

	for _, s := range expenseTransactions {
		decrypted, err := ConvertTransactionExpense(s)
		if err != nil {
			return "", err
		}
		decryptedExpenseTransactions = append(decryptedExpenseTransactions, decrypted)
	}

	fmt.Println(decryptedIncomeTransactions)
	fmt.Println(decryptedExpenseTransactions)

	return "OK", nil
}
