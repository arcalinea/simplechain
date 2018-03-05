package main 

import (
    "fmt"
    "log"
    "os"
    "net/url"
    "net/http"
    "encoding/json"
    // "io/ioutil"
    
    "github.com/urfave/cli"
    "../types"
)

// sendtx you 100 -from=me -memo=hi
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
    
    var res types.SendTxResponse
    err := Call("sendtx", map[string]string{
        "to": to, 
        "amount": amount,
        "from": from,
        "memo": memo,
    }, &res)
    
    out, err := json.MarshalIndent(res, "","  ")
    if err != nil {
        return err
    }
    fmt.Println(string(out))
    return nil
}

func GetInfo(c *cli.Context) error {
    var res types.GetInfoResponse
    err := Call("getinfo", map[string]string{}, &res)
    if err != nil {
        return err
    }
    out, err := json.MarshalIndent(res, "","  ")
    if err != nil {
        return err
    }
    fmt.Println(string(out))
    return nil
}

func Call(cmd string, options map[string]string, out interface{}) error {
    vals := make(url.Values)
    for k, v := range options {
        vals.Set(k, v)
    }
    resp, err := http.PostForm("http://127.0.0.1:1234/" + cmd, vals)
    //fmt.Println("Response:", resp)
    if err != nil {
        return err
    }
    // buf, err := ioutil.ReadAll(resp.Body)
    err = json.NewDecoder(resp.Body).Decode(out)
    if err != nil {
        return err
    }
    // fmt.Printl(string(buf))
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
      {
        Name:   "getinfo",
        Usage:  "get blockchain information",
        Action: GetInfo,
      },
  }

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}
