package netaic

import (
	"errors"
	"log"
	"net"
	"strings"
)

func GetNetInterfaceAddr(interfaceName string) (string, error) {
	netIF, getNetIFErr := net.InterfaceByName(interfaceName)
	if getNetIFErr != nil {
		return "", getNetIFErr
	}

	addrList, getAddrListErr := netIF.Addrs()
	if getAddrListErr != nil {
		return "", getAddrListErr
	}

	if len(addrList) <= 0 {
		return "", errors.New("len(addrList) <= 0")
	}

	lAddr := strings.Split(addrList[0].String(), "/")[0]
	return lAddr, nil
}

type ScanNetInterfaceFunc func(net.Interface)

func ScanAllNetInterfaces(f ScanNetInterfaceFunc) error {
	netIFList, err := net.Interfaces()
	if err != nil {
		return err
	}

	for _, netIF := range netIFList {
		f(netIF)
	}

	return nil
}

func PrintfAllNetInterfaceInfo() error {
	return ScanAllNetInterfaces(func(netIF net.Interface) {
		addrList, err := netIF.Addrs()
		if err != nil {
			log.Printf("Get the %v net interface addrs error, %v\n", netIF.Name, err)
			return
		}

		for _, addr := range addrList {
			log.Printf("Print net interface info, name: %s, addr: %v, network: %v, mtu: %v, hardware addr: %v, flags: %v\n",
				netIF.Name, addr.String(), addr.Network(), netIF.MTU, netIF.HardwareAddr.String(), netIF.Flags.String())
		}
	})
}
