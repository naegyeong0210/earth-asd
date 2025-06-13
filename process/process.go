package process

import (
	"sync"
	"time"

	"github.com/datauniverse-lab/earth-asd/factory"
	"github.com/datauniverse-lab/earth-asd/process/batch"
	"github.com/datauniverse-lab/earth-common/utils"
)

type ASDProcess struct {
	Fac         *factory.Factory
	sktProcess  batch.SKTProcess
	//ktProcess   batch.KTProcess
	//lgupProcess batch.LGUPProcess
}

func (_self *ASDProcess) Initialize(fac *factory.Factory) {
	_self.Fac = fac
	_self.sktProcess = batch.SKTProcess{Fac: fac}
	//_self.ktProcess = batch.KTProcess{Fac: fac}
	//_self.lgupProcess = batch.LGUPProcess{Fac: fac}
}

func (_self *ASDProcess) Processing() {
	requestID := utils.GenTransID()

	var service string

	switch {
	case _self.Fac.Property.BenzProcess:
		service = "오토콜"
	case _self.Fac.Property.BentleyProcess:
		service = "스마트피싱보호"
	case _self.Fac.Property.SaturnProcess:
		service = "휴대폰쿠폰지갑"
	case _self.Fac.Property.FerrariProcess:
		service = "휴대폰분실보호"
	case _self.Fac.Property.TeslaProcess:
		service = "휴대폰가족보호"
	default:
		_self.Fac.Print("모든 프로세스가 비활성화 상태")
		for {
			time.Sleep(10 * time.Second)
		} // 파드가 죽지않도록 대기
	}
	

	_self.Fac.Print(requestID, service, "나이 조회 프로세스 실행")

	var wg sync.WaitGroup

	if _self.Fac.Propertys().SKTProcess {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_self.sktProcess.Process(requestID)
		}()
	}

	if _self.Fac.Propertys().KTProcess {
		wg.Add(1)
		go func() {
			defer wg.Done()
		//	_self.ktProcess.Process(requestID)
		}()
	}

	if _self.Fac.Propertys().LGUPProcess {
		wg.Add(1)
		go func() {
			defer wg.Done()
		//	_self.lgupProcess.Process(requestID)
		}()
	}

	wg.Wait()
}
