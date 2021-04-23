package main

import (
	"fmt"
	"net"
	"os"

	"github.com/urfave/cli"
)

// func server(c) error {
// 	ns, err := net.LookupNS(c.String("host"))
// 	errorLog(err)
// 	for i := 0; i < len(ns); i++ {
// 		fmt.Println(ns[i].Host)
// 	}
// 	return nil
// }

// func errorLog(a) {
// 	if a != nil {
// 		return a
// 	}
// }

func main() {
	app := cli.NewApp()
	app.Name = "website lookup cli"
	app.Usage = "lets you query ip"

	myFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "anand-mohanan.vercel.app",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "ns",
			Usage: "used to lookup the NAME SERVERS for particular host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				ns, err := net.LookupNS(c.String("host"))
				if err != nil {
					return err
				}
				for i := 0; i < len(ns); i++ {
					fmt.Println(ns[i].Host)
				}
				return nil
			},
		},
		{
			Name:  "ip",
			Usage: "used to lookup the IP ADDRESS for particular host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				ip, err := net.LookupIP(c.String("host"))
				if err != nil {
					return err
				}
				for i := 0; i < len(ip); i++ {
					fmt.Println(ip[i])
				}
				return nil
			},
		},
		{
			Name:  "cname",
			Usage: "used to lookup the CNAME for particular host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				cname, err := net.LookupCNAME(c.String("host"))
				if err != nil {
					return err
				}

				fmt.Println(cname)

				return nil
			},
		},
		{
			Name:  "mx",
			Usage: "used to lookup the MX RECORDS for particular host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				mx, err := net.LookupMX(c.String("host"))
				if err != nil {
					return err
				}
				for i := 0; i < len(mx); i++ {
					fmt.Println(mx[i].Host, mx[i].Pref)
				}
				return nil
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		return
	}

}

// var errorLog = func(a) {
// 	if a != nil {
// 		return a
// 	}
// }
