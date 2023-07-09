# go-ifsc 🚀

A simple tool to check & search IFSC codes of all Indian banks from the comfort of your terminal 🤓💪

# Features

- ⚡CLI tool
- ⚡Includes a REST API server 💥
- ⚡Single binary
- ⚡Works Offline
- ⚡Search functionality

# API Demo

Get bank details for an IFSC Code (Similar to [Razorpay's API](https://ifsc.razorpay.com/))

```sh
curl -s http://insomnia247.nl:5100/YESB0DNB002 | jq
{
  "BANK": "Delhi Nagrik Sehkari Bank",
  "IFSC": "YESB0DNB002",
  "BRANCH": "Delhi Nagrik Sehkari Bank IMPS",
  "CENTRE": "DELHI",
  "DISTRICT": "DELHI",
  "STATE": "MAHARASHTRA",
  "ADDRESS": "720, NEAR GHANTAGHAR, SUBZI MANDI, DELHI - 110007",
  "CONTACT": "+919560344685",
  "IMPS": true,
  "RTGS": true,
  "CITY": "MUMBAI",
  "ISO3166": "IN-MH",
  "NEFT": true,
  "MICR": "110196002",
  "UPI": true,
  "SWIFT": null
}
```

Search for banks in an area

```sh
curl -s http://insomnia247.nl:5100/search?q=hitech city

[{"BANK":"Bandhan Bank","IFSC":"BDBL0002291","BRANCH":"KAVURI HILLS BRANCH HYDERABAD","CENTRE":"HYDERABAD","DISTRICT":"HYDERABAD","STATE":"TELANGANA","ADDRESS":"2-44 2,MADHAPUR PRIDE,GUTTALA BEGUMPET,MADHAPUR,HITECH CITY MAIN ROAD,GROUND FLOOR,MADHAPUR POLICE STATION -500081,TELANGANA","CONTACT":"+913366090909","IMPS":true,"RTGS":true,"CITY":"HYDERABAD","ISO3166":"IN-TG","NEFT":true,"MICR":"500750012","UPI":true,"SWIFT":null},{"BANK":"Central Bank of India","IFSC":"CBIN0283164","BRANCH":"HITECH AGRICULTURAL FINANCE BRANCH","CENTRE":"BHOPAL","DISTRICT":"BHOPAL","STATE":"MADHYA PRADESH","ADDRESS":"9, ARERA HILL, JAIL ROAD, BHOPAL, DIST- BHOPAL, MADHYA PRADESH-462011","CONTACT":"+912222612008","IMPS":true,"RTGS":true,"CITY":"BHOPAL","ISO3166":"IN-MP","NEFT":true,"MICR":"462016022","UPI":true,"SWIFT":null}]
```

Search for all axis banks in a state, Eg: Goa

```sh
curl -s http://insomnia247.nl:5100/search?q=axis in-ga

[{"BANK":"Axis Bank","IFSC":"UTIB0003418","BRANCH":"GOGOL","CENTRE":"SOUTH","DISTRICT":"SOUTH","STATE":"GOA","ADDRESS":"SHOP NO 12345 AR MANSION GOGOL","CONTACT":"+918326570622","IMPS":true,"RTGS":true,"CITY":"MARGAO","ISO3166":"IN-GA","NEFT":true,"MICR":"403211014","UPI":true,"SWIFT":null}]
```

> **NOTE:** Please do not use this api endpoint for production, It's for testing purposes only.

# Download

All binaries are built on gitlab ci for every release. You can download the latest release from the links below

## Desktop/Server

- 🐧 [Linux (x64)](https://kskarthik.gitlab.io/go-ifsc/linux/ifsc)
- 🪟 [Windows (x64)](https://kskarthik.gitlab.io/go-ifsc/win/ifsc.exe)
- 🍎 [Mac (darwin x64)](https://kskarthik.gitlab.io/go-ifsc/darwin/ifsc)

## Docker

The docker image is built for each commit & uploaded to docker hub

Docker hub: https://hub.docker.com/r/kskarthik/ifsc

```sh
$ docker pull kskarthik/ifsc:latest
```

Configure the docker compose file. You can override the default entrypoint & port

```yaml
ifsc-server:
  image: kskarthik/ifsc:latest
  entrypoint: ["ifsc", "server", "--port", "3000"]
  expose:
    - "3000"
```

# Usage

```bash
# give execute permissions
chmod +x ifsc
# on the first start, you need to index the IFSC data locally, This is not required for subsequent runs.
./ifsc index
# start a rest API server
./ifsc server
# Print usage info
./ifsc help
```

# CLI Examples 😍

### Validate a IFSC code

```bash
$ ifsc check ICIC0004530

    BANK	ICICI Bank
    IFSC	ICIC0004530
  BRANCH	RAJAHMUNDRY AVA ROAD
  CENTRE	RAJAHMUNDRY
DISTRICT	RAJAHMUNDRY
   STATE	ANDHRA PRADESH
 ADDRESS	ICICI Bank Ltd., D.No 80.1.13, Jayasree Buildings, Jayasree Garden, AVA Road, Rajahmundry, Dist. East Godavari, Andhra Pradesh.533103
 CONTACT	+918008104316
    IMPS	yes
    RTGS	yes
    CITY	EAST GODAVARI
 ISO3166	IN-AP
    NEFT	yes
    MICR	533229007
     UPI	yes
   SWIFT	N/A
```

### Search for banks

```bash
$ ifsc search -m all axis karol bagh
+-------------+-----------+-------+-------+---------------------------+
| IFSC        | BANK      | CITY  | STATE | ADDRESS                   |
+-------------+-----------+-------+-------+---------------------------+
| UTIB0000223 | Axis Bank | DELHI | DELHI | 6/83,PADAM SINGH RD,WESTE |
|             |           |       |       | RN EXTN AREA  KAROL BAGH, |
|             |           |       |       |  WEST DELHI               |
| UTIB0SIPSB2 | Axis Bank | DELHI | DELHI | 794 JOSHI ROAD KAROL  BAG |
|             |           |       |       | H NEW DELHI-110005        |
+-------------+-----------+-------+-------+---------------------------+
```

# Build From Source 🛠️

### Linux

Tested with Go `>=1.19` & might work with Go versions `>=1.16` too.

```bash
# after cloning  this repo, cd into it
cd go-ifsc/

# build the binary, public/ will contain the built binaries
bash build.sh
```

# License ⚖️

All the code is licensed under `GPLv3` & server code is `AGPLv3` license

# TODO 📝

- [x] Provide a REST API for search & validation
- [x] handle the search command param properly
- [x] Improve the check & search logic

# Credits 🤝

- [Razorpay IFSC](https://github.com/razorpay/ifsc/releases) repository for the csv dump
- [Bleve](https://pkg.go.dev/github.com/blevesearch/bleve/v2) search library
- [Cobra](https://github.com/spf13/cobra) CLI library
