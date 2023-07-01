# ifsc 🚀

A simple tool to check & search IFSC codes of all Indian banks from the comfort of your terminal 🤓💪

Also includes a REST API server 💥

⚡ Single binary | Works Offline ⚡

# Demo

get bank details for an IFSC Code

```sh
curl http://insomnia247.nl:5100/YESB0DNB002

{"BANK":"Delhi Nagrik Sehkari Bank","IFSC":"YESB0DNB002","BRANCH":"Delhi Nagrik Sehkari Bank IMPS","CENTRE":"DELHI","DISTRICT":"DELHI","STATE":"MAHARASHTRA","ADDRESS":"720, NEAR GHANTAGHAR, SUBZI MANDI, DELHI - 110007","CONTACT":"+919560344685","IMPS":true,"RTGS":true,"CITY":"MUMBAI","ISO3166":"IN-MH","NEFT":true,"MICR":"110196002","UPI":true,"SWIFT":null}
```

search for banks in an area

```sh
curl http://insomnia247.nl:5100/search/hitech city

[{"BANK":"Bandhan Bank","IFSC":"BDBL0002291","BRANCH":"KAVURI HILLS BRANCH HYDERABAD","CENTRE":"HYDERABAD","DISTRICT":"HYDERABAD","STATE":"TELANGANA","ADDRESS":"2-44 2,MADHAPUR PRIDE,GUTTALA BEGUMPET,MADHAPUR,HITECH CITY MAIN ROAD,GROUND FLOOR,MADHAPUR POLICE STATION -500081,TELANGANA","CONTACT":"+913366090909","IMPS":true,"RTGS":true,"CITY":"HYDERABAD","ISO3166":"IN-TG","NEFT":true,"MICR":"500750012","UPI":true,"SWIFT":null},{"BANK":"Central Bank of India","IFSC":"CBIN0283164","BRANCH":"HITECH AGRICULTURAL FINANCE BRANCH","CENTRE":"BHOPAL","DISTRICT":"BHOPAL","STATE":"MADHYA PRADESH","ADDRESS":"9, ARERA HILL, JAIL ROAD, BHOPAL, DIST- BHOPAL, MADHYA PRADESH-462011","CONTACT":"+912222612008","IMPS":true,"RTGS":true,"CITY":"BHOPAL","ISO3166":"IN-MP","NEFT":true,"MICR":"462016022","UPI":true,"SWIFT":null},
```

**NOTE:** Please do not use this api endpoint for production, It's for testing purposes only.

# Download

There is no release ATM. All binaries are built on gitlab ci on each commit. You can download from the links below

## Desktop/Server

- 🐧 [Linux (x64)](https://kskarthik.gitlab.io/go-ifsc/linux/ifsc)
- 🪟 [Windows (x64)](https://kskarthik.gitlab.io/go-ifsc/win/ifsc.exe)
- 🍎 [Mac (darwin x64)](https://kskarthik.gitlab.io/go-ifsc/darwin/ifsc)

## Usage

```bash
# on the first start, you need to index the IFSC data locally, This is not required for subsequent runs.
ifsc index
```

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

# Build From Source 🛠️

### Linux

Tested with Go `>=1.19` & might work with Go versions `>=1.16` too.

```bash
# after cloning  this repo, cd into it
cd go-ifsc/

# build the binary, public/ will contain the built binaries
bash build.sh
```

# CLI Examples 😍

### Validate a IFSC code

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
SWIFT : N/A
```

### Search for banks

```bash
$ ifsc search tidel park
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
SWIFT : N/A
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
MICR : N/A
UPI : yes
SWIFT : N/A
----------------------
BANK : Karur Vysya Bank
IFSC : KVBL0001260
BRANCH : Karur Vysya Bank IMPS
CENTRE : CHENNAI
DISTRICT : CHENNAI
STATE : MAHARASHTRA
ADDRESS : D-6th FLOOR, D NORTH BLOCK, TIDEL PARK, TARAMANI
CONTACT : N/A
IMPS : yes
RTGS : yes
CITY : MUMBAI
ISO3166 : IN-MH
NEFT : no
MICR : 600053056
UPI : yes
SWIFT : N/A
----------------------
3 results
```

### REST API Examples

`ifsc server` command will launch the web server at `localhost:9000`. The port number can be customized with `--port` flag

#### GET `/:ifsc`

This validation endpoint is same as [Razorpay's API](https://github.com/razorpay/ifsc/wiki/API)

Below is the response of `/YESB0DNB002`

If the code is invalid or no result was found, Server responds with 404

```json
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

#### GET `/search/:search_term`

This API searches the data set for given `search_term` & returns an array of objects which match the search term

Example: `/search/king+koti`

```json
[
  {
    "BANK": "Cosmos Co-operative Bank",
    "IFSC": "COSB0000030",
    "BRANCH": "HYDERABAD",
    "CENTRE": "HYDERABAD",
    "DISTRICT": "HYDERABAD",
    "STATE": "ANDHRA PRADESH",
    "ADDRESS": "3-5-798, PRATHIMA SCHALASS, NEW NO.248, STREET NO.8, BASHEERBAUG HYDERGUDA, KING KOTI ROAD, HYDERABAD- 500 029",
    "CONTACT": "+914023231705",
    "IMPS": true,
    "RTGS": true,
    "CITY": "HYDERABAD",
    "ISO3166": "IN-AP",
    "NEFT": true,
    "MICR": "500164001",
    "UPI": true,
    "SWIFT": null
  }
]
```

# License ⚖️

All the code, except the `IFSC.csv` file dump, is licensed under `GPLv3` & server code is `AGPLv3` license

# TODO 📝

- [x] Provide a REST API for search & validation
- [x] handle the search command param properly
- [ ] Improve the check & search logic

# Credits 🤝

- Thanks to [Razorpay IFSC ](https://github.com/razorpay/ifsc/releases) repository for the csv dump
