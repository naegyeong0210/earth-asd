package formats

// 추후 common에 통합
type AsdMember struct {
	PNumber  string `json:"PNumber"`
	Telecom  int    `json:"Telecom"`
	PCode    string `json:"PCCode"`
	Age      int    `json:"Age"`
	RegDT    string `json:"RegDT"`
	Complete int    `json:"Complete"`
}

// Ferrari-common의 LGUP TCRS의 AGE_OUT이 주석처리 되어 있어 ASD에 자체 포멧이 필요해 추가
type LGUPRSPUserInfo struct {
	RESPCODE      int    `json:"RESPCODE"`
	RESPMSG       string `json:"RESPMSG"`
	Age           string `json:"AGE_OUT"`
	CTN_STUS_CODE string `json:"CTN_STUS_CODE"` //CTN 상태 코드(A:정상/S:일시 중지)
	//PRE_PAY_CODE  string `json:"PRE_PAY_CODE"`
	//REF_TYPE_CODE string `json:"REF_TYPE_CODE"`
	MDL_VALUE          string `json:"MDL_VALUE"`    //단말속성정보
	UNIT_MDL           string `json:"UNIT_MDL"`     //단말기명
	YOUNG_FEE_YN       string `json:"YOUNG_FEE_YN"` //청소년 정보료 상한요금제
	SVC_AUTH_DT        string `json:"SVC_AUTH_DT"`
	UNIT_LOSS_YN_CODE  string `json:"UNIT_LOSS_YN_CODE"` //분실여부
	REAL_BIRTH_PERS_ID string `json:"REAL_BIRTH_PERS_ID,omitempty"`
	SUB_BIRTH_PERS_ID  string `json:"SUB_BIRTH_PERS_ID,omitempty"`
}
