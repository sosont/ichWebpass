package main

import (
	"fmt"
	"log"
	"time"
)

type user struct {
	username string
	password string
}

type NatInfo struct {
	id          int
	cliAddress  string
	serAddress  string
	cliHttpport int
	serHttpport int
	linkPort    int
	creattime   time
	lasttime    time
	status      int
	remark      string
	vKey        string
}

func NewNatInfo(CliHttpPort) *NatInfo {

}

func (ni *NatInfo) AddNat() (id int, err error) {
	rs, err := db.Exec("Insert into nat_info(cli_address,ser_address,cli_httpport,ser_httpport,link_port,creattime,lasttime,status,remark,vkey) values(?,?,?,?,?,?,?,?,?,?)", ni.cliAddress, ni.serAddress, ni.cliHttpport, ni.serHttpport, ni.linkPort, time.Now(), time.Now(), ni.status, ni.remark, ni.vKey)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	return
}

func (ni *NatInfo) GetTrpProt() _p int{
     c, _ := listerner.
	
	}