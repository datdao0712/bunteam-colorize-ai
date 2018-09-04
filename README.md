# BUNTEAM COLORIZE AI
Colorizing black and white photos
## Getting Started
### Prerequisites
> Golang version go1.10.1

> Npm version 5.6.0
### Installing
In your folder `/src` of $GOPATH:

```
go get -u github.com/astaxie/beego
```
```
go get -u github.com/beego/bee
```
Bunteam app use [Beego](https://beego.me/) framework for backend

Next, go to folder `/frontend/bunteam-colorize-ai/`:
```
npm install
```

### Run

In your current project directory:
```
bee run
```
*Expose port `:3000` (Backend)*

Start frontend in `/frontend/bunteam-colorize-ai/`:
```
npm start
```
*Expose port `:8080` (Frontend)*