package main

import (
	"fmt"
	libvirt "libvirt.org/go/libvirt"
	"os"
	"time"
)

func processClient(c libvirt.AdmClient) error {
	id, err := c.GetID()
	if err != nil {
		return nil
	}

	timestamp, err := c.GetTimestamp()
	if err != nil {
		return nil
	}
	timestampStr := time.Unix(timestamp, 0).Format(time.UnixDate)

	transport, err := c.GetTransport()
	if err != nil {
		return nil
	}
	transportStr := map[libvirt.ClientTransport]string{
		libvirt.CLIENT_TRANS_TCP:  "TCP",
		libvirt.CLIENT_TRANS_TLS:  "TLS",
		libvirt.CLIENT_TRANS_UNIX: "unix",
	}[transport]

	info, err := c.GetInfo(0)
	if err != nil {
		return nil
	}

	fmt.Printf("\tClient ID: %d timestamp: %s transport: %s\n",
		id, timestampStr, transportStr)

	var infoStr string = ""
	if info.ReadonlySet {
		infoStr += fmt.Sprintf("RO: %t ", info.Readonly)
	}
	if info.SocketAddressSet {
		infoStr += fmt.Sprintf("Socket address: %s ", info.SocketAddress)
	}
	if info.SaslUsernameSet {
		infoStr += fmt.Sprintf("SASL username: %s ", info.SaslUsername)
	}
	if info.TlsX509DistinguishedNameSet {
		infoStr += fmt.Sprintf("DN: %s ", info.TlsX509DistinguishedName)
	}
	if info.UnixUsernameSet {
		infoStr += fmt.Sprintf("UNIX username: %s ", info.UnixUsername)
	}
	if info.UnixGroupnameSet {
		infoStr += fmt.Sprintf("UNIX groupname: %s ", info.UnixGroupname)
	}
	if info.UnixUserIdSet {
		infoStr += fmt.Sprintf("UID:GID %d:%d ", info.UnixUserId, info.UnixGroupId)
	}
	if info.UnixProcessIdSet {
		infoStr += fmt.Sprintf("PID: %d ", info.UnixProcessId)
	}
	if info.SelinuxContextSet {
		infoStr += fmt.Sprintf("SELinux: %s ", info.SelinuxContext)
	}
	fmt.Printf("\t\t%s\n", infoStr)

	return nil
}

func processServer(s libvirt.AdmServer) error {
	sName, err := s.GetName()
	if err != nil {
		return nil
	}
	fmt.Printf("Server name: %s\n", sName)

	threadParams, err := s.GetThreadPoolParameters(0)
	if err != nil {
		return nil
	}
	fmt.Printf("Workers (min/curr/max/free/prio) %d/%d/%d/%d/%d\n",
		threadParams.MinWorkers, threadParams.CurrentWorkers,
		threadParams.MaxWorkers, threadParams.FreeWorkers,
		threadParams.PrioWorkers)

	clientLimits, err := s.GetClientLimits(0)
	if err != nil {
		return nil
	}
	fmt.Printf("Client limits (curr/max/currUnauth/maxUnauth) %d/%d/%d/%d\n",
		clientLimits.CurrentClients, clientLimits.MaxClients,
		clientLimits.CurrentUnauthClients, clientLimits.MaxUnauthClients)

	clients, err := s.ListClients(0)
	if err != nil {
		return nil
	}

	for _, c := range clients {
		fmt.Println()
		err := processClient(c)
		if err != nil {
			return nil
		}

		c.Free()
	}

	return nil
}

func main() {
	var uri string = "libvirtd:///system"

	if len(os.Args) >= 2 {
		uri = os.Args[1]
	}

	conn, err := libvirt.NewAdmConnect(uri, 0)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	uriStr, err := conn.GetURI()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Connected to: %s\n", uriStr)

	filtersStr, err := conn.GetLoggingFilters(0)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Logging filters: %s\n", filtersStr)

	outputsStr, err := conn.GetLoggingOutputs(0)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Logging outputs: %s\n", outputsStr)

	servers, err := conn.ListServers(0)
	if err != nil {
		panic(err)
	}

	for _, s := range servers {
		fmt.Println()
		err := processServer(s)
		if err != nil {
			panic(err)
		}

		s.Free()
	}
}
