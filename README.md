# Install required tools

Install PostGres 10

Install golang 1.11.1, nodejs 10.13.0, gobuffalo 0.13.1

Turn gomodules ON
```bash
export GO111MODULE=on
```

```bash
buffalo db create
buffalo db migrate
buffalo dev
```

Set GOOGLE_KEY and GOOGLE_SECRET from your google api account for auth. Use the .env file for development.

Go to [http://localhost:3000](http://localhost:3000)



## Get the cert (so that I don't forget - not used)

```bash
certbot --manual -d *.domainname.org certonly -m email --agree-tos --config-dir=~/hsm/letsencrypt --work-dir=~/hsm/letsencrypt --logs-dir=~/hsmletsencrypt --preferred-challenges=dns
```