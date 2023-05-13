# go-ifsc 🚀

A simple tool to check & search IFSC codes of all Indian banks from your terminal ! 🤓💪

⚡ Single binary, No dependencies, Works Offline ⚡

```bash
This utility helps to search, validate IFSC codes of Indian banks

Usage:
  ifsc [command]

Available Commands:
  check       Check whether a given IFSC code is valid
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  search      Fuzzy search for banks / IFSC codes

Flags:
  -h, --help     help for ifsc
  -t, --toggle   Help message for toggle

Use "ifsc [command] --help" for more information about a command.
```

# Download

All binaries are built on gitlab ci for each commit.

Links:

- 🐧 [Linux (x64)](https://kskarthik.gitlab.io/go-ifsc/linux/ifsc)
- 🪟 [Windows (x64)](https://kskarthik.gitlab.io/go-ifsc/win/ifsc.exe)
- 🍎 [Mac (darwin x64)](https://kskarthik.gitlab.io/go-ifsc/darwin/ifsc)

# Build From Source 🛠️

Tested with Go `>=1.19` & might work with Go versions `>=1.16` too.

```bash
# after cloning  this repo, cd into it
cd go-ifsc/

# download the IFSC.csv to cmd/
wget https://github.com/razorpay/ifsc/releases/download/v2.0.12/IFSC.csv -P cmd/

# build the binary, public/ will contain the built binaries
bash build.sh
```

# Examples 😍

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
SWIFT : ?
```

### Search for banks

```bash
$ ifsc search tidel park
kar@earth:~/my/projects/go-ifsc$ go run main.go search tidel park
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
3 results
```

# License ⚖️

All the code, except the `IFSC.csv`, file is licensed under `GPLv3`

# TODO 📝

- [ ] Provide a REST API for search & validation
- [ ] Improve the search logic
- [ ] handle the search command param properly

# Credits

- Thanks to [Razorpay IFSC](https://github.com/razorpay/ifsc/releases) for the csv dump
