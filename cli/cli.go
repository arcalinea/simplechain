package main 

import (
    "fmt"
    "log"
    "os"
    "net/url"
    "net/http"
    
    "github.com/urfave/cli"
)

// sendtx you 100 -from=me -memo=ily
func SendTx(c *cli.Context) error {
    if len(c.Args()) != 2 {
        return fmt.Errorf("To and amount must be specified")
    }
    to := c.Args()[0]
    amount := c.Args()[1]
    var from string 
    if c.String("from") == "" {
        from = "default"
    } else {
        from = c.String("from")
    }
    memo := c.String("memo")
    
    return Call("sendtx", map[string]string{
        "to": to, 
        "amount": amount,
        "from": from,
        "memo": memo,
    })
}

func Call(cmd string, options map[string]string) error {
    vals := make(url.Values)
    for k, v := range options {
        vals.Set(k, v)
    }
    resp, err := http.PostForm("http://127.0.0.1:1234/" + cmd, vals)
    fmt.Println("Response:", resp)
    if err != nil {
        return err
    }
    return nil
}

func main() {
  app := cli.NewApp()
  app.Name = "linkchain-cli"
  app.Usage = "rpc client for linkchain"
  
  app.Commands = []cli.Command{
      {
        Name:    "sendtx",
        Usage:   "send a transaction",
        Flags: []cli.Flag {
            cli.StringFlag{
                Name: "from",
                Value: "",
                Usage: "specify from address",
            },
            cli.StringFlag{
                Name: "memo",
                Value: "",
                Usage: "add memo to transaction",
            },
        },
        Action:  SendTx,
      },
  }

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}
