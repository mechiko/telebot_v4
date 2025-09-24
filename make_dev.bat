setlocal
  @rem SET PATH=E:\bin\TDM-GCC-64\bin;%PATH%
  SET PATH=E:\bin\mingw\mingw64\bin;%PATH%

  SET GOARCH=amd64
  SET GOOS=windows
  
  PUSHD .
  cd E:\src\goproj\telegram_bot_telebot\telebot_js
  call npx vite build -m development
  POPD
  rmdir E:\src\goproj\telegram_bot_telebot\telebot\pkg\vueapp /s /q

  @REM https://ss64.com/nt/robocopy.html
  @REM /E : Copy Subfolders, including Empty Subfolders.
  robocopy E:\src\goproj\telegram_bot_telebot\telebot_js\dist\ E:\src\goproj\telegram_bot_telebot\telebot\pkg\vueapp /e
  
  go build -ldflags "-s -w" -o telebot.exe ./cmd/bot

endlocal
