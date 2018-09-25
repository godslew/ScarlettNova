# ScarlettNova
Twitter DM chat bot

###環境構築
```
# Goがインストールされている前提です。

# 必要なライブラリの取得

go get -u github.com/golang/dep/cmd/dep
go get -u github.com/direnv/direnv

# 環境変数を.envrc (direnv) で管理している
echo 'eval "$(direnv hook bash)"' >> ~/.bashrc
cp .envrc.sample .envrc
direnv allow
```

###起動
```
go build main.go  
go run main.go 
```