package geo

import (
	"context"
	"fmt"
	"strings"

	"github.com/ip2location/ip2location-go/v9"
	"google.golang.org/grpc/metadata"
)

// package init will load bin file in memory
func init() {
	ip2location.Open("./IP2LOCATION-LITE-DB5.BIN")
}

// Ip2Location gets ipv4 address,
// and returns country_short, region, latitude and longitude
func Ip2Location(ip string) (string, string, float32, float32) {
	results := ip2location.Get_all(ip)
	return results.Country_short, results.Region, results.Latitude, results.Longitude
}

// ExtractIp extract ipv4 adress from context
// that passed down from grpc nginx ingress server
func ExtractIp(ctx context.Context) (string, error) {
	addr := ""
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", fmt.Errorf("failed to get metadata")
	}
	if len(md["x-real-ip"]) != 0 {
		addr = md["x-real-ip"][0]
	} else if len(md[":authority"]) != 0 {
		s := strings.Split(md[":authority"][0], ":")
		if len(s) == 0 {
			return "", fmt.Errorf("failed to get local ip")
		}
		addr = s[0]
	} else {
		return "", fmt.Errorf("failed to get real ip")
	}
	return addr, nil
}
