package iron_payload

import (
  "os"
  "io/ioutil"
  "fmt"
  "encoding/json"
)

// Get's payload from IronWorker config
func GetPayload() interface{} {
  payloadIndex := 0
  for index, arg := range(os.Args) {
    if arg == "-payload" {
      payloadIndex = index + 1
    }
  }
  fmt.Println(os.Args)
  if payloadIndex >= len(os.Args) {
    panic("No payload specified")
  }
  payload := os.Args[payloadIndex]
  var data interface{}
  raw, err := ioutil.ReadFile(payload)
  if err != nil {
    panic(err.Error())
  }

  err = json.Unmarshal(raw, &data)
  if err != nil {
    panic(err.Error())
  }
  return data
}
