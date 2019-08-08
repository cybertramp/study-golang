# Go 언어 마스터

## 기초

### 1. Go 설치

Go 버전은 1.12 버전을 설치한다.

#### 설치)

##### 1. Go 1.12 설치

32bit 혹은 이외의 버전은 [공식 사이트](https://golang.org/dl/)에서 다운받을수 있다.

```bash
wget https://dl.google.com/go/go1.12.7.linux-amd64.tar.gz
sudo tar -xzf go /usr/local
```

##### 2. Go 환경설정

GOROOT는 Go가 설치된 경로이며, GOPATH 작업할 디렉토리의 경로이다.

여기서는 work-go를 작업 디렉토리로 지정한다.

```bash
# 작업할 디렉토리 생성
mkdir ~/work-go
# 환경 변수 설정
export GOROOT=/usr/local/go
export GOPATH=$HOME/work-go
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
```

위의 설정은 현재 세션에서만 유효하기 때문에 지속적으로 적용하기 위해 `~/.bashrc`에 기록하면된다.

기록한 이후에는 `source ~/.bashrc`를 통해 적용시켜준다.

##### 3. 설치확인

```bash
# go 버전확인
go version
# go 환경설정 확인
go env
```

### 2. Helloworld 그리고 실행, 빌드

언어를 들어갈때 제일 처음 접하는 Helloworld를 출력해본다.

``` bash
mkdir helloworld
vi main.go
```

``` go
package main
import "fmt"

func main(){
    fmt.Println("Helloworld!")
}
```

#### 실행

``` bash
go run main.go
```

#### 빌드 후 바이너리 실행

``` bash
go build -o helloworld main.go

# 32bit 빌드
GOARCH=386 go build -o helloworldx86 main.go
# 64bit 빌드
GOARCH=amd64 go build -o helloworldx64 main.go
```

> 재미있는 것은 -o 플래그가 [소스파일.go] 전에 있어야 빌드가 된다는 것

원하는 아키텍쳐로 빌드하고 싶은 경우 [GOARCH](https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63) 를 참고할 것