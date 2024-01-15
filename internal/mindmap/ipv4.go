package mindmap

import "net"

// Masks for /8, /16, and /24
var mask8 = net.IPv4Mask(255, 0, 0, 0)
var mask16 = net.IPv4Mask(255, 255, 0, 0)
var mask24 = net.IPv4Mask(255, 255, 255, 0)

// isIPv4 checks if the given string is a valid IPv4 address.
func isIPv4(address string) bool {
	// net.ParseIP returns a valid IP address if the string is a valid IP.
	// However, it also accepts IPv6, so further checks are needed.
	ip := net.ParseIP(address)
	if ip == nil {
		return false
	}

	// Finally, ensure that the original string matches the format of the IP.
	// net.IP.String() returns the IP in normalized form, which helps in this comparison.
	return ip.To4().String() == address
}

// cidrBaseIP calculates the base IP of the given IP address for /8, /16, and /24 subnets.
func cidrBaseIP(ipStr string) []string {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return nil
	}

	ip = ip.To4()
	if ip == nil {
		return nil
	}

	// Calculating base IPs
	baseIP8 := ip.Mask(mask8).String()
	baseIP16 := ip.Mask(mask16).String()
	baseIP24 := ip.Mask(mask24).String()

	// Using a map to store unique base IPs
	uniqueBaseIPs := make(map[string]bool)

	// Adding base IPs in the order of /24, /16, /8
	baseIPsOrder := []string{baseIP24, baseIP16, baseIP8}
	var baseIPs []string

	for _, baseIP := range baseIPsOrder {
		if !uniqueBaseIPs[baseIP] {
			uniqueBaseIPs[baseIP] = true
			baseIPs = append(baseIPs, baseIP)
		}
	}

	return baseIPs
}

// ipv4Parsing parses IPv4 addresses
type ipv4Parsing struct{}

func (p *ipv4Parsing) Parse(input string) []string {
	if !isIPv4(input) {
		return nil
	}
	parts := cidrBaseIP(input)
	return append([]string{input}, parts...)
}
