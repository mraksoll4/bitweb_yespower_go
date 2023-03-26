install all build tools.

```
sudo apt-get update
sudo apt-get install autoconf automake libtool build-essential libssl-dev
```

install golang.

```
sudo apt-get install golang
```

check version.

```
go version
```
do if need.

```
export GOPATH=$HOME/go
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
```

next clone this repo.

```
git clone https://github.com/bitweb-project/bitweb_yespower_go.git
```

next build libs and run test.

```
cd bitweb_yespower_go
cd yespower
make
cd ..
go build -o test test.go
./test
```

if all correct it will give output.

```
6123f3f1c4e95d052dbe2f0238426f192a9592c55e55a3394054be399a000000
```
you can change rawblock and result at test.go.

Important!

Block hash is as is not reversed!

for comprare it with real block , you can do at console this

```
echo 6123f3f1c4e95d052dbe2f0238426f192a9592c55e55a3394054be399a000000 | grep -o .. | tac | paste -sd '' -

0000009a39be544039a3555ec592952a196f4238022fbe2d055de9c4f1f32361
```

https://explorer.bitwebcore.net/block/0000009a39be544039a3555ec592952a196f4238022fbe2d055de9c4f1f32361