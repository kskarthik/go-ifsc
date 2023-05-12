# go-ifsc 🚀

A simple tool to check & search IFSC codes of all Indian banks from your terminal ! 🤓💪

⚡ Single binary, No dependencies, Works Offline ⚡

```bash
This utility shows the bank details of given IFSC code

 USAGE: ./ifsc [COMMAND] [INPUT]

 COMMANDS:
	check - checks the given IFSC code & return the bank details if valid
	search - return results of banks based on keyword
	serve - starts the REST API server [TODO]
```

# Download

All binaries are built in gitlab ci for each commit.

Links: 🐧 [Linux (x64)](https://kskarthik.gitlab.io/go-ifsc/linux/ifsc) | 🪟 [Windows (x64)](https://kskarthik.gitlab.io/go-ifsc/win/ifsc.exe) | 🍎 [Mac (darwin x64)](https://kskarthik.gitlab.io/go-ifsc/darwin/ifsc)

# Build From Source 🛠️

Tested with Go `>=1.19` & might work with Go versions `>=1.16` too.

1. Download the latest `IFSC.csv` file from https://github.com/razorpay/ifsc/releases into the cloned repository

2. Run `go build -o ifsc main.go`

# Examples 😍

## Validate a IFSC code

```bash
$ ifsc check ICIC0004530

BANK : ICICI Bank
IFSC : ICIC0004530
BRANCH : RAJAHMUNDRY AVA ROAD
CENTRE : RAJAHMUNDRY
DISTRICT : RAJAHMUNDRY
STATE : ANDHRA PRADESH
ADDRESS : ICICI Bank Ltd., D.No 80.1.13, Jayasree Buildings, Jayasree Garden, AVA Road, Rajahmundry, Dist. East Godavari, Andhra Pradesh.533103
CONTACT : +918008104316
IMPS : yes
RTGS : yes
CITY : EAST GODAVARI
ISO3166 : IN-AP
NEFT : yes
MICR : 533229007
UPI : yes
SWIFT : ?
```

## Search for banks

```bash
$ ./ifsc search "tidel park"
BANK : Canara Bank
IFSC : CNRB0002715
BRANCH : TIDEL PARK, CHENNAI
CENTRE : CHENNAI
DISTRICT : CHENNAI
STATE : TAMIL NADU
ADDRESS : MODULE NO.1,4,CANAL BK ROAD, TARAMANI, CHENNAI 600 113.
CONTACT : +914422540371
IMPS : yes
RTGS : yes
CITY : CHENNAI
ISO3166 : IN-TN
NEFT : yes
MICR : 600015084
UPI : yes
SWIFT : ?
----------------------
BANK : HDFC Bank
IFSC : HDFC0007033
BRANCH : TIDEL PARK COIMBATORE
CENTRE : COIMBATORE
DISTRICT : COIMBATORE
STATE : TAMIL NADU
ADDRESS : GR FLR ELCOT SEZ IT ITES ILLANKURICHI RD AERODROME COIMBATORE COIMBATORE TAMIL NADU 641014
CONTACT : +9118602676161
IMPS : yes
RTGS : yes
CITY : COIMBATORE
ISO3166 : IN-TN
NEFT : yes
MICR : 641240033
UPI : yes
SWIFT : HDFCINBB
----------------------
BANK : ICICI Bank
IFSC : ICIC0007729
BRANCH : CHENNAITIDEL PARK
CENTRE : CHENNAI
DISTRICT : CHENNAI
STATE : TAMIL NADU
ADDRESS : ICICI Bank Ltd., 1st Floor, Tidel Park No. 4, Canal Bank Road, Taramani, Chennai.600113, Tamil Nadu
CONTACT : +917397482952
IMPS : yes
RTGS : yes
CITY : CHENNAI
ISO3166 : IN-TN
NEFT : yes
MICR : 600229167
UPI : yes
SWIFT : ?
----------------------
BANK : ICICI Bank
IFSC : ICIC0007729
BRANCH : CHENNAITIDEL PARK
CENTRE : CHENNAI
DISTRICT : CHENNAI
STATE : TAMIL NADU
ADDRESS : ICICI Bank Ltd., 1st Floor, Tidel Park No. 4, Canal Bank Road, Taramani, Chennai.600113, Tamil Nadu
CONTACT : +917397482952
IMPS : yes
RTGS : yes
CITY : CHENNAI
ISO3166 : IN-TN
NEFT : yes
MICR : 600229167
UPI : yes
SWIFT : ?
----------------------
BANK : Karur Vysya Bank
IFSC : KVBL0001901
BRANCH : CALL CENTRE CHENNAI
CENTRE : CHENNAI
DISTRICT : CHENNAI
STATE : TAMIL NADU
ADDRESS : FIRST FLOOR A NORTH BLOCK TIDEL PARK TARAMANI CHENNAI 600113
CONTACT : +912224398197
IMPS : yes
RTGS : yes
CITY : CHENNAI
ISO3166 : IN-TN
NEFT : yes
MICR : ?
UPI : yes
SWIFT : ?
----------------------
BANK : State Bank of India
IFSC : SBIN0004285
BRANCH : TIDEL PARK CHENNAI
CENTRE : CHENNAI
DISTRICT : CHENNAI
STATE : TAMIL NADU
ADDRESS : NO 4 BANAL BANKTARAMANI ,CHENNAI ,PIN - 600113
CONTACT : ?
IMPS : yes
RTGS : yes
CITY : CHENNAI
ISO3166 : IN-TN
NEFT : yes
MICR : 600002105
UPI : yes
SWIFT : SBININBB475
----------------------
BANK : State Bank of India
IFSC : SBIN0014361
BRANCH : TIDEL PARK, COIMBATORE
CENTRE : COIMBATORE
DISTRICT : COIMBATORE
STATE : TAMIL NADU
ADDRESS : ELCOTS IT SEZ, GOVERNMENT MEDICAL COLLEGE CAMPUS, VILANKURICHY VILLAGE, COIMBATORE NORTH TALUK, COIMBATORE641004
CONTACT : ?
IMPS : yes
RTGS : yes
CITY : COIMBATORE
ISO3166 : IN-TN
NEFT : yes
MICR : 641002068
UPI : yes
SWIFT : ?
----------------------
BANK : State Bank of India
IFSC : SBIN0070867
BRANCH : TIDEL PARK, COIMBATORE
CENTRE : COIMBATORE
DISTRICT : COIMBATORE
STATE : TAMIL NADU
ADDRESS : TIDEL PARK LTD, TPCL BUILDING, GROUND FLOOR, CIVIL AERODROME POST,COIMBATORE 641014 TIDELPARKCBEATSBT
CONTACT : ?
IMPS : yes
RTGS : yes
CITY : COIMBATORE
ISO3166 : IN-TN
NEFT : yes
MICR : 641002309
UPI : yes
SWIFT : ?
----------------------
BANK : State Bank of India
IFSC : SBIN0070867
BRANCH : TIDEL PARK, COIMBATORE
CENTRE : COIMBATORE
DISTRICT : COIMBATORE
STATE : TAMIL NADU
ADDRESS : TIDEL PARK LTD, TPCL BUILDING, GROUND FLOOR, CIVIL AERODROME POST,COIMBATORE 641014 TIDELPARKCBEATSBT
CONTACT : ?
IMPS : yes
RTGS : yes
CITY : COIMBATORE
ISO3166 : IN-TN
NEFT : yes
MICR : 641002309
UPI : yes
SWIFT : ?
----------------------
BANK : Karur Vysya Bank
IFSC : KVBL0001260
BRANCH : Karur Vysya Bank IMPS
CENTRE : CHENNAI
DISTRICT : CHENNAI
STATE : MAHARASHTRA
ADDRESS : D-6th FLOOR, D NORTH BLOCK, TIDEL PARK, TARAMANI
CONTACT : ?
IMPS : yes
RTGS : yes
CITY : MUMBAI
ISO3166 : IN-MH
NEFT : no
MICR : 600053056
UPI : yes
SWIFT : ?
----------------------
10 results
```

# License ⚖️

All the code, except the `IFSC.csv`, file is licensed under `GPLv3`

# TODO 📝

- [ ] Provide a REST API for search & validation
- [ ] Improve the search logic
- [ ] handle the search command param properly

# Credits

- Thanks to [Razorpay IFSC](https://github.com/razorpay/ifsc/releases) for the csv dump
