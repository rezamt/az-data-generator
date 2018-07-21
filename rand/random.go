package rand

import (
	"strconv"
	"strings"
	"net"
	"math/rand"
	"time"
)

var randomizer *rand.Rand

func init() {
	randomizer = rand.New(rand.NewSource(time.Now().UnixNano()))
}

var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[randomizer.Intn(len(letterRunes))]
	}
	return string(b)
}

func ResourceGroupName() string {
	return "rg-"+RandStringRunes(10)
}

type MachineStatus struct {
	ProvisioningState string
	PowerStatus       string
}

func Subscription() string {
	subs := []string{"production", "development", "sandbox", "services"}

	return randomFromList(subs)
}

func Location() string {
	locations := []string{"australiaeast", "australiasoutheast", "uksouth", "ukwest", "uswest", "australiacentra"}

	return randomFromList(locations)
}

func VMSize() string {
	vmSizes := []string{"Basic_A0",	"Basic_A1",	"Basic_A2",	"Basic_A3",	"Basic_A4",	"Standard_A0",	"Standard_A1",	"Standard_A10",	"Standard_A11",	"Standard_A1_v2",	"Standard_A2",	"Standard_A2_v2",	"Standard_A2m_v2",	"Standard_A3",	"Standard_A4",	"Standard_A4_v2",	"Standard_A4m_v2",	"Standard_A5",	"Standard_A6",	"Standard_A7",	"Standard_A8",	"Standard_A8_v2",	"Standard_A8m_v2",	"Standard_A9",	"Standard_B1ms",	"Standard_B1s",	"Standard_B2ms",	"Standard_B2s",	"Standard_B4ms",	"Standard_B8ms",	"Standard_D1",	"Standard_D11",	"Standard_D11_v2",	"Standard_D12",	"Standard_D12_v2",	"Standard_D13",	"Standard_D13_v2",	"Standard_D14",	"Standard_D14_v2",	"Standard_D15_v2",	"Standard_D16_v3",	"Standard_D16s_v3",	"Standard_D1_v2",	"Standard_D2",	"Standard_D2_v2",	"Standard_D2_v3",	"Standard_D2s_v3",	"Standard_D3",	"Standard_D32_v3",	"Standard_D32s_v3",	"Standard_D3_v2",	"Standard_D4",	"Standard_D4_v2",	"Standard_D4_v3",	"Standard_D4s_v3",	"Standard_D5_v2",	"Standard_D64_v3",	"Standard_D64s_v3",	"Standard_D8_v3",	"Standard_D8s_v3",	"Standard_DS1",	"Standard_DS11",	"Standard_DS11_v2",	"Standard_DS12",	"Standard_DS12_v2",	"Standard_DS13",	"Standard_DS13-2_v2",	"Standard_DS13-4_v2",	"Standard_DS13_v2",	"Standard_DS14",	"Standard_DS14-4_v2",	"Standard_DS14-8_v2",	"Standard_DS14_v2",	"Standard_DS15_v2",	"Standard_DS1_v2",	"Standard_DS2",	"Standard_DS2_v2",	"Standard_DS3",	"Standard_DS3_v2",	"Standard_DS4",	"Standard_DS4_v2",	"Standard_DS5_v2",	"Standard_E16_v3",	"Standard_E16s_v3",	"Standard_E2_v3",	"Standard_E2s_v3",	"Standard_E32-16_v3",	"Standard_E32-8s_v3",	"Standard_E32_v3",	"Standard_E32s_v3",	"Standard_E4_v3",	"Standard_E4s_v3",	"Standard_E64-16s_v3",	"Standard_E64-32s_v3",	"Standard_E64_v3",	"Standard_E64s_v3",	"Standard_E8_v3",	"Standard_E8s_v3",	"Standard_F1",	"Standard_F16",	"Standard_F16s",	"Standard_F16s_v2",	"Standard_F1s",	"Standard_F2",	"Standard_F2s",	"Standard_F2s_v2",	"Standard_F32s_v2",	"Standard_F4",	"Standard_F4s",	"Standard_F4s_v2",	"Standard_F64s_v2",	"Standard_F72s_v2",	"Standard_F8",	"Standard_F8s",	"Standard_F8s_v2",	"Standard_G1",	"Standard_G2",	"Standard_G3",	"Standard_G4",	"Standard_G5",	"Standard_GS1",	"Standard_GS2",	"Standard_GS3",	"Standard_GS4",	"Standard_GS4-4",	"Standard_GS4-8",	"Standard_GS5",	"Standard_GS5-16",	"Standard_GS5-8",	"Standard_H16",	"Standard_H16m",	"Standard_H16mr",	"Standard_H16r",	"Standard_H8",	"Standard_H8m",	"Standard_L16s",	"Standard_L32s",	"Standard_L4s",	"Standard_L8s",	"Standard_M128-32ms",	"Standard_M128-64ms",	"Standard_M128ms",	"Standard_M128s",	"Standard_M64-16ms",	"Standard_M64-32ms",	"Standard_M64ms",	"Standard_M64s",	"Standard_NC12",	"Standard_NC12s_v2",	"Standard_NC12s_v3",	"Standard_NC24",	"Standard_NC24r",	"Standard_NC24rs_v2",	"Standard_NC24rs_v3",	"Standard_NC24s_v2",	"Standard_NC24s_v3",	"Standard_NC6",	"Standard_NC6s_v2",	"Standard_NC6s_v3",	"Standard_ND12s",	"Standard_ND24rs",	"Standard_ND24s",	"Standard_ND6s",	"Standard_NV12",	"Standard_NV24",	"Standard_NV6",}

	size := randomFromList(vmSizes)

	return size
}

func VMStatus() (provisioningState string, powerStatus string) {

	status := []MachineStatus{
		{ProvisioningState: "succeeded", PowerStatus: "running"},
		{ProvisioningState: "succeeded", PowerStatus: "stopped"},
		{ProvisioningState: "failed", PowerStatus: "stopped"},
	}

	stat := randomFromVMStatus(status)

	provisioningState = stat.ProvisioningState
	powerStatus = stat.PowerStatus

	return
}

func VMName() string {
	return "VM-"+RandStringRunes(10)
}

// IpV4Address returns a valid IPv4 address as string
func IpV4Address() string {
	blocks := []string{}
	for i := 0; i < 4; i++ {
		number := randomizer.Intn(255)
		blocks = append(blocks, strconv.Itoa(number))
	}

	return strings.Join(blocks, ".")
}

// IpV6Address returns a valid IPv6 address as net.IP
func IpV6Address() string {
	var ip net.IP
	for i := 0; i < net.IPv6len; i++ {
		number := uint8(randomizer.Intn(255))
		ip = append(ip, number)
	}
	return ip.String()
}

func randomFromList(source []string) string {
	return source[randomizer.Intn(len(source))]
}

func randomFromVMStatus(source []MachineStatus) MachineStatus {
	return source[randomizer.Intn(len(source))]
}

func ImagePublisher() string {
	publishers := []string {"Microsoft Windows", "Linux"}

	return publishers[randomizer.Intn(len(publishers))]
}

func IPAllocation() string {
	alloc := []string {"static", "dynamic"}
	return alloc[randomizer.Intn(len(alloc))]
}