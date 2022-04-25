package utils

import (
	"association/global"
	"github.com/ip2location/ip2location-go"
)

func Ip2LocationCity(ip string) string {
	ip2location.Open(global.ASS_CONFIG.System.Ip2locationPath)
	results := ip2location.Get_all(ip)
	defer ip2location.Close()
	return results.City

	//fmt.Printf("country_short: %s\n", results.Country_short)
	//fmt.Printf("country_long: %s\n", results.Country_long)
	//fmt.Printf("region: %v\n", results.Region)
	//fmt.Printf("city: %v\n", results.City)
	//fmt.Printf("isp: %s\n", results.Isp)
	//fmt.Printf("latitude: %f\n", results.Latitude)
	//fmt.Printf("longitude: %f\n", results.Longitude)
	//fmt.Printf("domain: %s\n", results.Domain)
	//fmt.Printf("zipcode: %s\n", results.Zipcode)
	//fmt.Printf("timezone: %s\n", results.Timezone)
	//fmt.Printf("netspeed: %s\n", results.Netspeed)
	//fmt.Printf("iddcode: %s\n", results.Iddcode)
	//fmt.Printf("areacode: %s\n", results.Areacode)
	//fmt.Printf("weatherstationcode: %s\n", results.Weatherstationcode)
	//fmt.Printf("weatherstationname: %s\n", results.Weatherstationname)
	//fmt.Printf("mcc: %s\n", results.Mcc)
	//fmt.Printf("mnc: %s\n", results.Mnc)
	//fmt.Printf("mobilebrand: %s\n", results.Mobilebrand)
	//fmt.Printf("elevation: %f\n", results.Elevation)
	//fmt.Printf("usagetype: %s\n", results.Usagetype)
	//fmt.Printf("api version: %s\n", ip2location.Api_version())

}
