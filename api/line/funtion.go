package line

import (
	"example/line-bot-ledger/api/public"
	"example/line-bot-ledger/controller"
	"example/line-bot-ledger/model"
	"example/line-bot-ledger/request"
	"example/line-bot-ledger/utils"
	"os"
	"strings"
)

func ReplyUser(line request.LineMessage) (string, error) {
	utils.LogWithTypeStruct(line)
	lineId := line.Events[0].Source.UserID
	findUser, err := controller.GetUserByLineId(lineId)
	if err != nil {
		return "", err
	}

	// rawFileLogin, err := os.Open("stateUserLogin.json")
	// if err != nil {
	// 	return "", err
	// }

	// fileLogin, err := utils.ConvertFileToJson(rawFileLogin)
	// if err != nil {
	// 	return "", err
	// }

	rawFileNoneLogin, err := os.Open("stateUserNoneLogin.json")
	if err != nil {
		return "", err
	}

	fileNoneLogin, err := utils.ConvertFileToJson(rawFileNoneLogin)
	if err != nil {
		return "", err
	}

	var replyText string
	message := line.Events[0].Message
	stateLogin := CheckStateLogin(findUser)

	if message.Type != "text" {
		replyText = "ตอนนี้ผมยังรองรับแค่การพิมปกติ 🙇‍♂️"
	} else if line.Events[0].Type == "follow" {
		replyText = "สวัสดีครับ ยินดีต้อนรับสู่โปรแกรม Rai rub Rai jia (รายรับ รายจ่าย)\nวันนี้มีอะไรให้ผมรับใช้พิมได้เลยครับ 😊"
	} else {
		if stateLogin {
			// replyText, isLogout, err := StateUserLogin(message.Text, fileLogin)
		} else {
			replyText, err = StateUserNoneLogin(message.Text, fileNoneLogin, lineId)
			if err != nil {
				return "", err
			}
		}
	}

	text := model.Text{
		Type: "text",
		Text: replyText,
	}

	messageReply := model.ReplyMessage{
		ReplyToken: line.Events[0].ReplyToken,
		Messages: []model.Text{
			text,
		},
	}

	err = utils.ReplyMessageLine(messageReply)
	if err != nil {
		return "", err
	}

	return "OK", nil
}

func StateUserLogin(text string, state interface{}) (string, bool, error) {
	// switch text {
	// case "login":
	// 	return "กรุณากรอกข้อมูลการลงชื่อเข้าใช้ด้วยวิธีการดังนี้ \n", true, nil
	// }

	return "", false, nil
}

func StateUserNoneLogin(text string, state model.StateUser, lineId string) (string, error) {
	splitText := strings.Split(text, " ")
	messageText := splitText[0]
	var replyText string

	switch messageText {
	case "login":
		userLogin := request.Login{
			Email:    splitText[1],
			Password: splitText[2],
		}

		rawUser, err := public.LoginUser(userLogin, lineId)
		if err != nil {
			return "", err
		}

		if rawUser == nil {
			replyText = "ท่านยังไม่ได้ลงทะเบียน\n" + "สามารถดูวิธีการลงเบียนง่ายๆด้วยคำสั่ง \"how-to-register\" ได้เลยครับ 😎"
		} else {
			var user model.User
			err = utils.ConvertInterfaceToStruct(rawUser, &user)
			replyText = "ยินดีต้อนรับคุณ " + user.Name + "\n" + "สามารถพิมคำสั่งเพื่อใช้งานโปรแกรมรายรับ-รายจ่ายได้เลยครับ 😊\n" + "ท่านสามารถพิม \"help\" เพื่อตรวจสอบคำสั่งทั้งหมดที่มี"
		}
		break
	case "register":
		userRegister := request.Register{
			Email:    splitText[1],
			Password: splitText[2],
			Name:     splitText[3],
			Phone:    splitText[4],
		}

		_, err := public.CreateUser(userRegister)
		if err != nil {
			return "", err
		}

		replyText = "ลงทะเบียนเรียบร้อยแล้ว 🎉\n" + "ท่านสามารถดูวิธีการเข้าใช้งานระบบได้ผ่านทางการพิม \"how-to-login\""
		break
	case "help":
		var allState []string
		for _, s := range state {
			allState = append(allState, "- "+"\""+s.Type+"\" "+s.Description+"\n")
		}
		joinArr := strings.Join(allState[:], "")
		replyText = "ตอนนี้คำสั่งที่สามารถใช้ได้ก่อนล็อคอินคือ\n" + joinArr
		break
	default:
		findResponse := FindState(state, text)
		if findResponse == "" {
			findResponse = "ผมยังไม่เข้าใจที่คุณพิม กรุณาลองใหม่ภายหลังครับ 🙇‍♂️\nพิม \"help\" เพื่อตรวจสอบคำสั่งทั้งหมดที่มี"
		}
		replyText = findResponse
	}

	return replyText, nil
}

func CheckStateLogin(user interface{}) bool {
	if user != nil {
		return true
	} else {
		return false
	}
}

func FindState(data model.StateUser, text string) string {
	for i := range data {
		if data[i].Type == text {
			return data[i].Response
		}
	}
	return ""
}
