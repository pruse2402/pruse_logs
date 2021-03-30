# Log wrapper

## Setting up the dev environment

### 1. Install postgres database.
Steps Installing postgres in ubuntu.

```
sudo apt-get update
sudo apt-get install postgresql postgresql-contrib
```

You can set the following alias in **~/.profile** or **~/.bashrc** if you do not want to access **psql** without switching user to **postgres** or creating a new user.

```
alias psql='sudo -u postgres psql'
alias createdb='sudo -u postgres createdb'
alias dropdb='sudo -u postgres dropdb'
```
Setting up dev postgres environment without password. Replace version with the version of postgres installed.
```
sudo vi /etc/postgresql/9.5/main/pg_hba.conf
```

Make sure you change the privilege to **trust** in the file and save it.
```
local   all             postgres                                trust

# TYPE  DATABASE        USER            ADDRESS                 METHOD

# "local" is for Unix domain socket connections only
local   all             all                                     ident
# IPv4 local connections:
host    all             all             127.0.0.1/32            trust
# IPv6 local connections:
host    all             all             ::1/128                 trust
```
Restart postgres service
```
sudo service postgresql restart
```

### 2. Install go tools needed
Install dependent golang tools such as godep, goimports, etc. Clone the repository into **$GOPATH/src/lynk** folder.
```
go get github.com/tools/godep
go get golang.org/x/tools/cmd/goimports
go get ./...
```
### 3. Running the web application
Build and run the web application by using the following command. Basic exports are present in build.sh file. All confidential credentials can be added to dev.env file.
```
./build.sh
```