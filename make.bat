setlocal
  @rem SET PATH=E:\bin\TDM-GCC-64\bin;%PATH%
  SET PATH=E:\bin\mingw\mingw64\bin;%PATH%

  @rem SET GOARCH=386
  SET GOARCH=amd64
  SET GOOS=windows
  
  go build -ldflags "-s -w" -o telebot.exe ./cmd/bot

endlocal
