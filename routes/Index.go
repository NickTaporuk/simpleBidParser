package routes

import (
	"net/http"
	"encoding/json"
	"github.com/mileusna/useragent"
)

const (
	DeviceTypeDesktop = "Desktop"
	DeviceTypeMobile  = "Mobile"
	DeviceTypeTablet  = "Tablet"
	DeviceTypeBot     = "Bot"
)

type BidParser struct {
	DeviceType string `json:"device_type"`
	Os         string `json:"os"`
	OsVersion  string `json:"os_version"`
	Browser    string `json:"browser"`
	Country    string `json:"country"`
	Url        string `json:"url,omitempty"`
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	ua := ua.Parse(r.Header.Get("User-Agent"))

	BidParser := &BidParser{
		Os:        ua.OS,
		OsVersion: ua.OSVersion,
		Browser:   r.Host,
		Country:   r.URL.Path,
	}

	if ua.Mobile {
		BidParser.DeviceType = DeviceTypeMobile
	} else if ua.Tablet {
		BidParser.DeviceType = DeviceTypeTablet
	} else if ua.Desktop {
		BidParser.DeviceType = DeviceTypeDesktop
	} else if ua.Bot {
		BidParser.DeviceType = DeviceTypeBot
	}

	json, _ := json.Marshal(BidParser)

	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
