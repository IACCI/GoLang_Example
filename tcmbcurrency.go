package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"time"
)

type currencyday struct {
	ID         string
	date       time.Time
	dayno      string
	currencies []Currency
}

type Currency struct {
	code            string
	crossorder      int
	unit            int
	currencyNAMETR  string
	currencyname    string
	forexbuying     float64
	banknotebuying  float64
	banknoteselling float64
	forexselling    float64
	crossrateUSD    float64
	crossrateother  float64
}

type tarih_date struct {
	xmlname   xml.Name `xml:"tarih_date"`
	tarih     string   `xml:"tarih,attr"`
	date      string   `xml:"date,attr"`
	bulten_no string   `xml:"bulten_no,attr"`
	Currency  []currency
}

type currency struct {
	kod             string `xml:"kod,attr"`
	crossorder      string `xml:"crossorder,attr"`
	currencycode    string `xml:"currencycode,attr"`
	unit            string `xml:"unit"`
	isim            string `xml:"isim"`
	currencyname    string `xml:"currencyname"`
	forexbuying     string `xml:"forexbuying"`
	forexselling    string `xml:"forexselling"`
	banknotebuying  string `xml:"banknotebuying"`
	banknoteselling string `xml:"banknoteselling"`
	crossrateUSD    string `xml:"crossrateUSD"`
	crossrateother  string `xml:"crossrateother"`
}

func (c *currencyday) getdata(currencydate time.Time) {
	xdate := currencydate
	t := new(tarih_date)
	currday := t.getdata(currencydate, xdate)
	for {
		if currday == nil {
			currencydate = currencydate.AddDate(0, 0, -1)
			currday := t.getdata(currencydate, xdate)
			if currday != nil {
				break
			}
		} else {
			break
		}
	}
}

func (c *tarih_date) getdata(currencydate time.Time, xdate time.Time) *currencyday {
	currday := new(currencyday)
	var resp *http.Response
	var err error
	var url string

	currday = new(currencyday)
	url = "https://www.tcmb.gov.tr/kurlar/kurlar_tr.html" + currencydate.Format("200601") + "/" + currencydate.Format("02012006") + ".html"
	resp, err = http.Get(url)

	//fmt.Println(url)

	if err != nil {
		fmt.Println(err)
	} else {

		defer resp.Body.Close()

		if resp.StatusCode != http.StatusNotFound {
			tarih := new(tarih_date)
			d := xml.NewDecoder(resp.Body)
			marshalerr := d.Decode(&tarih)

			if marshalerr != nil {
				log.Printf("error : %v", marshalerr)
			}

			c = &tarih_date{}
			currday.ID = xdate.Format("20060102")
			currday.date = xdate
			currday.dayno = tarih.bulten_no
			currday.currencies = make([]Currency, len(tarih.Currency))

			for i, curr := range tarih.Currency {
				currday.currencies[i].code = curr.currencycode
				currday.currencies[i].currencyname = curr.currencyname
				currday.currencies[i].currencyNAMETR = curr.isim
				currday.currencies[i].banknotebuying, err = strconv.ParseFloat(curr.banknotebuying, 64)
				currday.currencies[i].banknoteselling, err = strconv.ParseFloat(curr.banknoteselling, 64)
				currday.currencies[i].forexbuying, err = strconv.ParseFloat(curr.forexbuying, 64)
				currday.currencies[i].forexselling, err = strconv.ParseFloat(curr.forexselling, 64)
				currday.currencies[i].crossorder, err = strconv.Atoi(curr.crossorder)
				currday.currencies[i].crossrateother, err = strconv.ParseFloat(curr.forexselling, 64)
				currday.currencies[i].crossrateUSD, err = strconv.ParseFloat(curr.forexselling, 64)
				currday.currencies[i].unit, err = strconv.Atoi(curr.unit)

			}
			fmt.Print(currday)

		} else {
			currday = nil
		}
	}
	return currday
}

func savejson(filename string, key interface{}) {
	outfile, err := os.Create(filename)
	errorcontrole(err)
	encoder := json.NewEncoder(outfile)
	err = encoder.Encode(key)
	errorcontrole(err)
	outfile.Close()
}

func errorcontrole(err error) {
	if err != nil {
		fmt.Println("fatal error : ", err.Error())
		os.Exit(1)
	}
}

func main() {

	runtime.GOMAXPROCS(2)
	starttime := time.Now()
	currencyday := new(currencyday)
	currencydate := time.Now()
	currencyday.getdata(currencydate)

	elepsedtime := time.Since(starttime)
	fmt.Printf("execution time: %s", elepsedtime)

}
