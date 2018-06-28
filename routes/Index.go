package routes

import (
	"net/http"
	"net"
	"encoding/json"
	"github.com/mileusna/useragent"
	"github.com/bsm/openrtb"
	"github.com/oschwald/geoip2-golang"
	"fmt"
	"io/ioutil"
	"log"
)

const (
	DeviceTypeDesktop    = "Desktop"
	DeviceTypeMobile     = "Mobile"
	DeviceTypeTablet     = "Tablet"
	DeviceTypeBot        = "Bot"
	MaxMindCountryDbPath = "./maxmind/GeoLite2-Country.mmdb"
)

type BidParser struct {
	DeviceType string `json:"device_type"`
	OS         string `json:"os"`
	Browser    string `json:"browser"`
	Country    string `json:"country"`
	Url        string `json:"url,omitempty"`
}

func SetDeviceTypeByUserAgent(agent *ua.UserAgent, bp *BidParser) {

	switch {
	case agent.Mobile:
		bp.DeviceType = DeviceTypeMobile
	case agent.Desktop:
		bp.DeviceType = DeviceTypeDesktop
	case agent.Tablet:
		bp.DeviceType = DeviceTypeTablet
	case agent.Bot:
		bp.DeviceType = DeviceTypeBot
	}
	fmt.Println(agent)
	bp.OS = agent.OS
	bp.Browser = agent.Name
}

func SetUrl(br *openrtb.BidRequest, bp *BidParser) {
	if br.App != nil {
		bp.Url = br.App.Domain
	}
}

func SetCountryByIp(path string, bp *BidParser, ipClient string) {
	fmt.Println("IP===>", ipClient)
	db, err := geoip2.Open(path)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ip := net.ParseIP(ipClient)

	if ip == nil {
		bp.Country = "localhost"
	} else {

		record, err := db.Country(ip)

		if err != nil {
			log.Fatal(err)
		}
		bp.Country = record.Country.IsoCode
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	var bidParser BidParser

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
	}

	if body != nil {
		var req openrtb.BidRequest

		json.Unmarshal([]byte(body), &req)

		SetUrl(&req,&bidParser)

	}

	w.Header().Set("Content-Type", "application/json")
	ua := ua.Parse(r.Header.Get("User-Agent"))

	SetDeviceTypeByUserAgent(&ua, &bidParser)
	SetCountryByIp(MaxMindCountryDbPath,&bidParser, r.RemoteAddr)
	json, _ := json.Marshal(bidParser)

	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
