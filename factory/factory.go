package factory

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
     dms "github.com/datauniverse-lab/tesla-common/proto/dms/dms-gen"
	dmsformats "github.com/datauniverse-lab/tesla-common/dmsapi/dmsformats"
	"github.com/datauniverse-lab/earth-common/dmrsapi/dmrsformats"
	"github.com/datauniverse-lab/earth-common/utils"
	"github.com/sirupsen/logrus"
	cronowriter "github.com/utahta/go-cronowriter"
		"google.golang.org/grpc"
)

type Saturn struct {
	DMRSINFO         dmrsformats.DMRSInfo   `json:"MIDDLECONF"`
		TcrsURL               string               `json:"TcrsURL"`
}
type Bentley struct {
	DMRSINFO         dmrsformats.DMRSInfo   `json:"MIDDLECONF"`
		TcrsURL               string               `json:"TcrsURL"`
}
type Tesla struct {
	DMSINFO         dmsformats.DMSInfo   `json:"MIDDLECONF"`
		TcrsURL               string               `json:"TcrsURL"`
}
type Benz struct {
	DMRSINFO         dmrsformats.DMRSInfo   `json:"MIDDLECONF"`
		TcrsURL               string               `json:"TcrsURL"`
}
type Ferrari struct {
	DMRSINFO         dmrsformats.DMRSInfo   `json:"MIDDLECONF"`
		TcrsURL               string               `json:"TcrsURL"`
}

type Config struct {
	SATURN Saturn  `json:"SATURN"`
	BENTLEY          Bentley `json:"BENTLEY"`
	TESLA            Tesla `json:"TESLA"`
	BENZ             Benz `json:"BENZ"`
	FERRARI          Ferrari `json:"FERRARI"`
	DelaySecSKT      int           `json:"DelaySecSKT"`
	DelaySecKT       int           `json:"DelaySecKT"`
	DelaySecLGUP     int           `json:"DelaySecLGUP"`
	MaxMemberList    int           `json:"MaxMemberList"`
	SKTProcess       bool          `json:"SKTProcess"`
	KTProcess        bool          `json:"KTProcess"`
	LGUPProcess      bool          `json:"LGUPProcess"`
	SaturnProcess    bool          `json:"SaturnProcess"`
	BentleyProcess   bool          `json:"BentleyProcess"`
	BenzProcess      bool          `json:"BenzProcess"`
	TeslaProcess     bool          `json:"TeslaProcess"`
	FerrariProcess   bool          `json:"FerrariProcess"`
		LogerfilePath    string                 `json:"LogerfilePath"`

}

type Factory struct {
	Logger         *logrus.Logger
	JSONConfigPath string
	JSONConfigURL  string
	Property       Config
	ConfigSet      string
	ConfigMap      map[string]interface{}
	HostName       string
	Dmscall      dms.ServicesClient
	GrpcClient     *grpc.ClientConn
}

func (_self *Factory) loadConfiguration(file string) {

	if _self.ConfigSet == "LIVE" {
		err := os.Mkdir(_self.JSONConfigPath, os.ModeDir)

		if err != nil {
			fmt.Print("Config 경로 생성 오류", err)
		} else {
			fmt.Print("Config 경로 생성 완료")
		}

		utils.DownloadFile(_self.JSONConfigURL, file)
	}

	dat, _ := ioutil.ReadFile(file)
	fmt.Print(string(dat))

	json.Unmarshal(dat, &_self.Property)

	_self.ConfigMap = make(map[string]interface{})
	json.Unmarshal(dat, &_self.ConfigMap)

}

func (_self *Factory) Initialize() {

	fmt.Print("", "", _self.JSONConfigPath+"config.json")

	_self.loadConfiguration(_self.JSONConfigPath + "config.json")

	fmt.Print("", "", _self.Property)

	writer := cronowriter.MustNew(_self.Propertys().LogerfilePath+"_"+_self.HostName+".log", cronowriter.WithMutex())
	mw := io.MultiWriter(writer, os.Stdout)
	logrus.SetOutput(mw)

	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"

	_self.Logger = logrus.New()
	_self.Logger.Formatter = new(logrus.TextFormatter)
	_self.Logger.SetFormatter(customFormatter)
	_self.Logger.Level = logrus.DebugLevel
	_self.Logger.Out = mw

	if _self.Propertys().TeslaProcess {
		// dms 연결
		conn, err := grpc.Dial(_self.Propertys().TESLA.DMSINFO.DMSURL, grpc.WithInsecure())
		if err != nil {
			_self.Print("Connection Error to DMS", err)
			panic(err)
		} else {
			_self.Print("Connection to DMS")
		}


		_self.GrpcClient = conn
		dmsclient := dms.NewServicesClient(conn)
		_self.Dmscall = dmsclient
	}

}

func (_self *Factory) ReloadConfig() {
	appEnv := flag.String("app-env", os.Getenv("APP_HOME"), "app env")
	flag.Parse()

	if *appEnv == "" {
		*appEnv = `./`
	}

	fmt.Println(*appEnv + "config.json")

	_self.loadConfiguration(*appEnv + "config.json")
}

func (_self *Factory) Propertys() Config {
	return _self.Property
}


func (_self *Factory) Print(v ...interface{}) {
	header := v[0]
	v = append(v[:0], v[0+1:]...)
	_self.Logger.Print("[REQUESTID][", header, "]", v)
}