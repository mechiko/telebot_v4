setlocal
  @rem SET PATH=E:\bin\TDM-GCC-64\bin;%PATH%
  SET PATH=E:\bin\mingw\mingw64\bin;%PATH%

  SET GOARCH=amd64
  SET GOOS=linux
  SET CGO_ENABLED=0
  
  go build -ldflags "-s -w" -o ./linux/telebot ./cmd/bot

endlocal
