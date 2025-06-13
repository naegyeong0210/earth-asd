package utils

import (
	"fmt"

	"git.datau.co.kr/benz/benz-common/commonformats/dmrs"
	"git.datau.co.kr/benz/benz-common/requests"
	"github.com/datauniverse-lab/earth-asd/formats"
)

const (
	SelectQuery  = "DBMW_00010"
	InsertQuery  = "DBMW_00020"
	ExecuteQuery = "DBMW_00030"
)

func setDmrsHeader(requestID string, query, cmdType string) dmrs.ReqDmrsHeader {
	return dmrs.ReqDmrsHeader{
		TransactionID: requestID,
		CallApp:       "ASD",
		XMLName:       "ASD",
		CmdType:       cmdType,
		Query:         query,
	}
}

func ReturnBenzAsdMembers(requestID string, dmrsURL string, telecom int, asdMember *[]formats.AsdMember, maxMemberList int) dmrs.RspDmrsHeader {
	body := asdMember

	dmrsHeader := setDmrsHeader(requestID, "SelectAsdMember", SelectQuery)

	d := []interface{}{
		telecom,
		maxMemberList,
	}

	header := requests.DmrsCall(dmrsURL, dmrsHeader, &body, d...)
	fmt.Printf("RspDmrsHeader: %+v\n", header)

	return header
}

func UpdateAge(requestID string, dmrsURL string, pNumber string, age int) {
	dmrsHeader := setDmrsHeader(requestID, "UpdateAgeCheck", ExecuteQuery)

	d := []interface{}{
		age,
		pNumber,
	}

	requests.DmrsCall(dmrsURL, dmrsHeader, nil, d...)
}
