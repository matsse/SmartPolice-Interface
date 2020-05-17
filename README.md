# SmartPolice-Interface
Smart Police Interface is an application that lets you create and manage format for 
Internet-of-Things devices. 



# Currently  this repository is set to private for the time being. The following guide is not applicable for now. 

## Requirements
The main requirement needed to install this program is to have go installed. 

But other installed packages that you have to install are: 
- wget
- git

1. On Ubuntu and Debian golang can be installed with the following command. 

```bash
# sudo apt update && sudo apt upgrade
```
Download golang .tar from their official website. 
```bash 
# wget https://dl.google.com/go/go1.14.3.linux-amd64.tar.gz
```

Unpack the archive and move it to desired folder. 
```bash 
tar -xvf go1.14.3.linux-amd64.tar.gz
sudo mv go /usr/local
```

Golang requires that you set a few environment variables for BASH and the system to 
know where to store yours and others packages. 

To set the GOROOT environment variable use the following command. This is where you  
unpacked the tarball. 
```bash 
export GOROOT=/usr/local/go
```


Next, set the directory of your go worskpace (GOPATH)
```bash 
export GOPATH=$HOME/go
```

Lastly set the variable to access the binary of the go code and other packages system-wide. 
```bash 
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
```





## Installation

```bash
go get -u github.com/matsse/SmartPolice-Interface 
```



