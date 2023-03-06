package config

import (
	"fmt"
)

func ExamplePrintScript() {
	config := Config{}
	config.Load("testdata/pixie.yaml")
	d, _ := config.Resolved["00:00:00:00:01:11"]
	s, _ := config.RenderScript(d, config.StaticRoot)
	fmt.Println(s)

	// Output:
	// #!ipxe
	//
	// set name h4
	// set ubuntu http://10.0.0.1:8880/static/ubuntu-22.04
	// set docker http://10.0.0.1:8880/static/docker-hosts-test
	//
	// kernel ${ubuntu}/vmlinuz cloud-config-url=/dev/null url=${ubuntu}/ubuntu-22.04.1-live-server-amd64.iso initrd=initrd ip=dhcp toram autoinstall ds=nocloud-net;s=testdata/cloud-init/
	//
	// initrd ${ubuntu}/initrd
	//
	// boot
}

func ExampleGetScriptPath() {
	// Test with a complex config
	config := Config{}
	config.Load("testdata/pixie.yaml")
	fmt.Println(config.GetScriptPath(""))
	fmt.Println(config.GetScriptPath("00:00:00:00:01:01"))
	fmt.Println(config.GetScriptPath("00:00:00:00:01:02"))
	fmt.Println(config.GetScriptPath("00:00:00:00:01:03"))
	fmt.Println(config.GetScriptPath("00:00:00:00:01:11"))

	// Test with minimal config
	m := Config{}
	m.Load("testdata/minimal.yaml")
	fmt.Println(m.GetScriptPath("11:22:33:44:55:66"))
	fmt.Println(m.GetScriptPath("66:55:44:33:22:11"))

	// Output:
	//
	// myscript/boot.ipxe
	// myscript/boot.ipxe
	// other/boot.ipxe
	// example.ipxe
	// testscript1
	// testscript2
}

func ExamplePrintResolvedVars() {
	config := Config{}
	config.Load("testdata/pixie.yaml")
	fmt.Println(config.Resolved["00:00:00:00:01:01"].Vars)
	fmt.Println(config.Resolved["00:00:00:00:01:02"].Vars)
	fmt.Println(config.Resolved["00:00:00:00:01:03"].Vars)
	fmt.Println(config.Resolved["00:00:00:00:01:11"].Vars)

	// Output:
	// &map[cloudinit:mygroup/cloud-init/ pixiehost:10.0.0.1:8880]
	// &map[cloudinit:mygroup/cloud-init/ pixiehost:10.0.0.1:8880]
	// &map[cloudinit:mygroup/cloud-init/ pixiehost:10.0.0.1:8880]
	// &map[cloudinit:testdata/cloud-init/ pixiehost:10.0.0.1:8880]
}