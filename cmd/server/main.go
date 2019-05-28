package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/urfave/cli"
	"gopkg.in/macaron.v1"

	"github.com/liangchenye/oss-deps/service"
)

var webCommand = cli.Command{
	Name:        "web",
	Usage:       "Hub Server",
	Description: "Hub Server stores the data/info.",
	Action:      runHubServer,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "address",
			Value: "0.0.0.0",
			Usage: "web service listen ip, default is 0.0.0.0; if listen with Unix Socket, the value is sock file path.",
		},
		cli.StringFlag{
			Name:  "listen-mode",
			Value: "http",
			Usage: "web service listen mode, default is http.",
		},
		cli.IntFlag{
			Name:  "port",
			Value: 1234,
			Usage: "web service listen at port 80; if run with https will be 443.",
		},
	},
}

func LoadSetting() error {
	localConfig := make(map[string]string)
	localConfig["train-meta-url"] = "../../test/local/trainmeta"
	localConfig["meta-url"] = "../../test/local/meta"
	localConfig["data-dir"] = "../../test/local"
	if err := service.SetDefault("local"); err != nil {
		return err
	}

	h, err := service.GetDefault()
	if err != nil {
		fmt.Println(err)
		return err
	}

	h.Init(localConfig)
	return nil
}

func runHubServer(c *cli.Context) error {
	m := macaron.New()

	SetRouters(m)
	if err := LoadSetting(); err != nil {
		return nil
	}

	switch c.String("listen-mode") {
	case "http":
		listenaddr := fmt.Sprintf("%s:%d", c.String("address"), c.Int("port"))
		fmt.Printf("Start listen to :%s\n", listenaddr)
		if err := http.ListenAndServe(listenaddr, m); err != nil {
			fmt.Printf("Start Hub Server http mode error: %v\n", err.Error())
			return err
		}
		break
	default:
		break
	}

	return nil
}

func main() {
	app := cli.NewApp()

	app.Name = "hub server"
	app.Usage = "Hub Server"
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		webCommand,
	}

	app.Run(os.Args)
}
