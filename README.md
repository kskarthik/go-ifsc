# go-ifsc 🚀

Search all the Indian banks info from the comfort of your terminal 🤓💪

# Features

- ⚡CLI tool
- ⚡Includes a REST API server 💥
- ⚡Single binary
- ⚡Works Offline
- ⚡Search & validation

# Download

All binaries are built on gitlab ci for every release. You can download the latest release binary from the respective links below:

- 🐧 [Linux (x64)](https://kskarthik.gitlab.io/go-ifsc/linux/ifsc)
- 🪟 [Windows (x64)](https://kskarthik.gitlab.io/go-ifsc/win/ifsc.exe)
- 🍎 [Mac (darwin x64)](https://kskarthik.gitlab.io/go-ifsc/darwin/ifsc)

## Docker

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

# API Reference 📡

REST API is documented using swagger. See [Swagger UI](https://kskarthik.gitlab.io/go-ifsc)

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

Match documents which satisfy all search terms

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

### Customize Columns

```bash
go run . search axis karol --limit 2 --columns 'BANK,ADDRESS,CONTACT'
+-----------+---------------------------+---------------+
| BANK      | ADDRESS                   | CONTACT       |
+-----------+---------------------------+---------------+
| Axis Bank | 6/83,PADAM SINGH RD,WESTE | +919582802231 |
|           | RN EXTN AREA  KAROL BAGH, |               |
|           |  WEST DELHI               |               |
| Axis Bank | 794 JOSHI ROAD KAROL  BAG | +919810840772 |
|           | H NEW DELHI-110005        |               |
+-----------+---------------------------+---------------+

```

> For more search options refer `ifsc search --help`

# Build From Source 🛠️

### Linux

Tested with Go `>=1.19` & might work with Go versions `>=1.16` too.

```bash
# after cloning  this repo, cd into it
cd go-ifsc/

# build the binary, public/ will contain the built binaries
bash build.sh
```

# License ⚖️

All the code is licensed under `GPLv3` & server code is `AGPLv3` license

# Credits 🤝

- [Razorpay IFSC](https://github.com/razorpay/ifsc/releases) repository for the csv dump
- [Bleve](https://pkg.go.dev/github.com/blevesearch/bleve/v2) search library
- [Cobra](https://github.com/spf13/cobra) CLI library
- Many thanks to [Dhruvin](https://dhruvin.dev/) for his suggestions in the design aspects
