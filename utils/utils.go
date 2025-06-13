package utils

import (
	"fmt"
	"bytes"
	"io/ioutil"
	"log"
	"strconv"
	"time"
	"strings"
	"net/http"

	"github.com/datauniverse-lab/earth-asd/formats"
	"git.datau.co.kr/ferrari/ferrari-common/tcrsapi/tcrsformats/ktformats"
	"git.datau.co.kr/ferrari/ferrari-common/tcrsapi/tcrsformats/sktformats"
	"github.com/sirupsen/logrus"

	"git.datau.co.kr/ferrari/ferrari-common/commonutils"
	"git.datau.co.kr/ferrari/ferrari-common/tcrsapi/tcrsformats"
)

// 각 통신사 SKT, KT, LGUP TCRS 별 나이를 반환하는 값의 포멧이 달라 만 나이로 통일 시키기 위한 메서드
func ExtractAge(data map[string]interface{}, tel int) int {
	switch tel {
	case 0: // SKT
		body, ok := data["BodyInfo"].(sktformats.UserInfoRsp)
		if !ok {
			logrus.Error("SKT: error: 응답 바디 에러, ", body.SSN_BIRTH_DT)
			return -1
		}

		if len(body.SSN_BIRTH_DT) < 8 {
			logrus.Error("SKT: TCRS error: SSN_BIRTH_DT")
			return -1
		}

		bd := body.SSN_BIRTH_DT

		result, err := sktAgeConverter(bd)

		fmt.Println("SKT:", body.SSN_BIRTH_DT, "나이:", result)

		if err != nil {
			return -1
		}

		return result

	case 1: // KT
		body, ok := data["Body"].(ktformats.RSPUserInfoAndKways)
		if !ok {
			logrus.Error("KT: TCRS error: 응답 바디 에러")
			return -1
		}

		if len(body.USER_SSN_FRONT) < 2 {
			logrus.Error("KT: TCRS error: USER_SSN_FRONT too short")
			return -1
		}

		result := ktAgeConverter(body.USER_SSN_FRONT)
		fmt.Println("KT:", body.USER_SSN_FRONT, ":", result)

		return result

	case 2: // LGUP
		body, ok := data["Body"].(formats.LGUPRSPUserInfo)
		if !ok {
			logrus.Error("LGUP: TCRS error: 응답 바디 에러")
			return -1
		}

		fmt.Println("LGUP:", body.Age, "나이:", body.Age)

		result, err := strconv.Atoi(body.Age)
		if err != nil {
			logrus.Error("LGUP: TCRS error: 나이가 숫자 형식이 아님: ", body.Age)
			return -1
		}

		return result
	}

	logrus.Error("ExtractAge() error")
	return -1
}

// SSN_BIRTH_DT 포멧 YYYYMMDD에서 만 나이로 변환.
func sktAgeConverter(ssnBirthDt string) (int, error) {
	if len(ssnBirthDt) != 8 {
		log.Printf("Invalid SSN_BIRTH_DT format: %s", ssnBirthDt)
		return -1, fmt.Errorf("SSN_BIRTH_DT must be in YYYYMMDD format")
	}

	birthYear, err := strconv.Atoi(ssnBirthDt[:4])
	if err != nil {
		log.Printf("Error parsing birth year: %v", err)
		return -1, fmt.Errorf("invalid birth year")
	}

	birthMonth, err := strconv.Atoi(ssnBirthDt[4:6])
	if err != nil {
		log.Printf("Error parsing birth month: %v", err)
		return -1, fmt.Errorf("invalid birth month")
	}

	birthDay, err := strconv.Atoi(ssnBirthDt[6:8])
	if err != nil {
		log.Printf("Error parsing birth day: %v", err)
		return -1, fmt.Errorf("invalid birth day")
	}

	birthDate := time.Date(birthYear, time.Month(birthMonth), birthDay, 0, 0, 0, 0, time.UTC)
	currentDate := time.Now().UTC()

	age := currentDate.Year() - birthDate.Year()
	if currentDate.Month() < birthDate.Month() ||
		(currentDate.Month() == birthDate.Month() && currentDate.Day() < birthDate.Day()) {
		age--
	}

	return age, nil
}

// ktAgeConverter  YYMMDD format to age.
func ktAgeConverter(userSSNFront string) int {
	if len(userSSNFront) != 6 {
		log.Printf("Error: TCRS USER_SSN_FRONT format invalid: %s", userSSNFront)
		return -1
	}

	yearStr := userSSNFront[0:2]
	monthStr := userSSNFront[2:4]
	dayStr := userSSNFront[4:6]

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		log.Printf("TCRS Error: Invalid year format: %s", yearStr)
		return -1
	}

	if year <= 25 {
		year += 2000
	} else {
		year += 1900
	}

	month, err := strconv.Atoi(monthStr)
	if err != nil {
		log.Printf("TCRS Error: Invalid month format: %s", monthStr)
		return -1
	}

	day, err := strconv.Atoi(dayStr)
	if err != nil {
		log.Printf("TCRS Error: Invalid day format: %s", dayStr)
		return -1
	}

	birthDate := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
	currentTime := time.Now()

	age := currentTime.Year() - birthDate.Year()
	if currentTime.Before(time.Date(year+age, time.Month(month), day, 0, 0, 0, 0, time.Local)) {
		age--
	}

	return age
}

// GetMemberInfoTCRS AGE_OUT을 포함하는 LGUPRSPUserInfo 사용이 불가피함에 따라 ferrari-common에서 현재 프로젝트로 가져와 수정하였습니다.
func GetMemberInfoTCRS(TCRSURL string, telecom string, pnumber string) map[string]interface{} {

	var retData map[string]interface{}
	var teleName string

	if telecom == "0" || telecom == "1" || telecom == "2" {
		teleName = commonutils.TeleTypeNumToTelecomName(telecom)
	} else {
		teleName = telecom

	}

	cmdType := "USERINFO"

	if telecom == "1" {
		cmdType = "USERINFOANDKWAYS"

	}

	tcrsHeader := tcrsformats.ReqHeader{CmdType: cmdType}
	tcrsBody := tcrsformats.ReqBodyPNumber{PNumber: pnumber}

	tcrsRSP := RestfulSendData(TCRSURL+teleName, commonutils.MakeJsonData(tcrsHeader, tcrsBody))

	retData = make(map[string]interface{})

	if strings.ToUpper(teleName) == "SKT" {

		var tcrsRspHeader tcrsformats.RspHeader
		var tcrsRspBody sktformats.RspMain
		var tcrsRspBodyDetail sktformats.UserInfoRsp
		tcrsRspBody.Body = &tcrsRspBodyDetail

		commonutils.JsonToHaderBody([]byte(tcrsRSP), &tcrsRspHeader, &tcrsRspBody)
		retData["Header"] = tcrsRspHeader
		retData["Body"] = tcrsRspBody
		retData["BodyInfo"] = tcrsRspBodyDetail

	} else if strings.ToUpper(teleName) == "KT" {

		var tcrsRspHeader tcrsformats.RspHeader
		var tcrsRspBody ktformats.RSPUserInfoAndKways

		commonutils.JsonToHaderBody([]byte(tcrsRSP), &tcrsRspHeader, &tcrsRspBody)

		retData["Header"] = tcrsRspHeader
		retData["Body"] = tcrsRspBody

	} else if strings.ToUpper(teleName) == "LGUP" {
		var tcrsRspHeader tcrsformats.RspHeader

		//원래 매핑하던 포멧에는 AGE가 포함하지 않는 이슈로 LGUPRSPUserInfo로 변경
		var tcrsRspBody formats.LGUPRSPUserInfo

		commonutils.JsonToHaderBody([]byte(tcrsRSP), &tcrsRspHeader, &tcrsRspBody)

		retData["Header"] = tcrsRspHeader
		retData["Body"] = tcrsRspBody

	}

	return retData

}

func RestfulSendData(url string, inData []byte) []byte {
	// TCRS의 KT API 안정성 이슈로 응답이 돌아오지 않을 때가 있어 5초 타임아웃이 있는 HTTP 클라이언트 사용
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	reqData := bytes.NewBuffer(inData)
	// http.Post 대신 client.Post 사용
	resp, err := client.Post(url, "application/json", reqData)
	if err != nil {
		fmt.Println("RestfulSendData error: ", err.Error())
	}
	var f []byte
	if resp != nil {

		f, _ = ioutil.ReadAll(resp.Body)
		resp.Body.Close()
	}

	return f

}
