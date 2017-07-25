package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	wsdl := "http://10.129.7.88:6025"

	xmlString := "<soapenv:Envelope xmlns:soapenv=\"http://schemas.xmlsoap.org/soap/envelope/\" xmlns:tem=\"http://tempuri.org/\" xmlns:beis=\"http://schemas.datacontract.org/2004/07/Beisen.UserFramework.Models\"><soapenv:Header/><soapenv:Body><tem:CreatePosition><!--Optional:--><tem:args><!--Optional:--><beis:opid>123456</beis:opid><!--Optional:--><beis:tid>100001</beis:tid><!--Optional:--><beis:Code>\"t22sr2r1223\"</beis:Code><!--Optional:--><beis:Description>\"122223\"</beis:Description><!--Optional:--><beis:Name>\"VC322测2试\"</beis:Name></tem:args></tem:CreatePosition></soapenv:Body></soapenv:Envelope>"

	client := &http.Client{}
	req, err := http.NewRequest("POST", wsdl, bytes.NewBufferString(xmlString))
	req.Header.Add("Content-Type", "text/xml;charset=UTF-8")
	req.Header.Add("SOAPAction", "http://tempuri.org/IPositionProvider/CreatePosition")
	resp, err := client.Do(req)
	fmt.Println(resp)
	fmt.Println(err)
	body, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))
	defer resp.Body.Close()
}
