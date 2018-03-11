package main

import (
  "log"
  "./mySubLib"
)

func main() {
  log.Println("Start...")

  dbConfig := &mySubLib.DbConfig{
    User:     "user",
    Pass:     "pass",
    Host:     "host",
    Port:     "port",
    Database: "dbName",
    Log:      true,
  }

  dbClient, err := mySubLib.DbConnect(dbConfig)
  if err != nil {
    panic(err)
  }

  defer dbClient.Db.Close()

  log.Println("Stop...")
}
