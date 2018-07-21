package main

import (
	"log"
	"os"
	"text/template"
	"./rand"
	"github.com/satori/go.uuid"
)

type Project struct {
	Names []string
	Name  string
}

type Stack struct {
	Project
}

type Config struct {
	Stack
}

type NetworkInterface struct {
	PrivateIp     string
	PrivateIPType string
	PublicIP      string
	PublicType    string
}

type VirtualMachine struct {
	Id             string
	Size           string
	Name           string
	NetworkInterface
	Location       string
	Subscription   string
	ResourceGroup  string
	Tags           string
	ImagePublisher string
	ProvisionState string
	PowerState     string
}

type VMSSKU struct {
	Name     string
	Tier     string
	Capacity int
}

type VMSS struct {
	Id     int
	Name   string
	OSName string
	IP     string
}

type ScaleSet struct {
	Id                string
	Name              string
	Sku               VMSSKU
	Location          string
	Subscription      string
	ResourceGroup     string
	Tags              string
	ProvisioningState string
	Overprovision     bool
	VirtualMachines []VMSS
}


func main() {
	generateVM(500, false)
	generateVMSS(500, false)
}



func generateVMSS(count int , console bool) {

	templates := []string{
		"templates/scaleset.template",
	}

	output := "scalesets.json"

	capacity := 2

	var sss = make([]ScaleSet, count)

	for i := 0; i < count; i++ {

		sku := VMSSKU {
			Name: rand.VMSize(),
			Capacity: capacity,
			Tier: "Standard",
		}

		ssName := rand.VMName()

		var vmss = make([]VMSS, capacity)

		for j:=0; j <capacity; j++ {

			vmss[j] = VMSS {
				Name: rand.VMName(),
				Id: j,
				IP: rand.IpV4Address(),
				OSName: rand.ImagePublisher(),
			}
		}

		ss := ScaleSet{
			Id:   ssName,
			Name: ssName,
			Sku: sku,
			Location:       rand.Location(),
			Subscription:   rand.Subscription(),
			ResourceGroup:  rand.ResourceGroupName(),
			Tags:           "CostCenter=A123; Environment=test; DataClassification: Confidential; BusinessUnit: Cloud;",
			ProvisioningState : "Succeeded",
			Overprovision: false,
			VirtualMachines: vmss,
		}

		sss[i] = ss

	}

	generate(sss, templates, output, console)
}

func generateVM(count int , console bool) {

	templates := []string{
		"templates/virtualmachine.template",
	}

	output := "virtualmachines.json"

	var vms = make([]VirtualMachine, count)

	for i := 0; i < len(vms); i++ {

		prov, power := rand.VMStatus()

		vm := VirtualMachine{
			Id:   uuid.Must(uuid.NewV4()).String(),
			Size: rand.VMSize(),
			Name: rand.VMName(),
			NetworkInterface: NetworkInterface{
				PrivateIp:     rand.IpV4Address(),
				PrivateIPType: rand.IPAllocation(),
				PublicIP:      rand.IpV4Address(),
				PublicType:    rand.IPAllocation(),
			},
			Location:       rand.Location(),
			Subscription:   rand.Subscription(),
			ResourceGroup:  rand.ResourceGroupName(),
			Tags:           "CostCenter=A123; Environment=test; DataRestriction: Confidential;",
			ImagePublisher: rand.ImagePublisher(),
			ProvisionState: prov,
			PowerState:     power,
		}

		vms[i] = vm

	}

	generate(vms, templates, output, console)
}


func generate(list interface{}, templates []string, output string, console bool) {
	tmpl, err := template.ParseFiles(templates...)

	if err != nil {
		panic(err)
	}

	if !console {
		fo, err := os.Create(output)
		defer fo.Close()

		if err != nil {
			log.Println(err)
			return
		}

		err = tmpl.Execute(fo, list)
		if err != nil {
			log.Println(err)
			return
		}

	} else {
		err = tmpl.Execute(os.Stdout, list)
		if err != nil {
			log.Println(err)
			return
		}
	}
}