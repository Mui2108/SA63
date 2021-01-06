## Installation

ติดตั้ง swag  ด้วยคำสั่ง

```bash
go get -u github.com/swaggo/swag/cmd/swag@v1.6.7
swag init
```
สั่ง download dependency ด้วยคำสั่ง
```bash
go mod tidy
```
แล้วสั่ง compile โปรแกรม backend ด้วยคำสั่ง
```bash
go build -o main.exe main.go
```
ถ้า error ให้ลง GCC ** exec: "gcc": executable file not found in %PATH%
ติดตั้ง GCC compiler: https://github.com/jmeubank/tdm-gcc/releases/download/v9.2.0-tdm64-1/tdm64-gcc-9.2.0.exe

เมื่อโปรแกรม compile ได้อย่างถูกต้องแล้ว
รันโปรแกรมโดยการรันคำสั่ง main
```bash
.\main.exe
```
