setlocal
  @rem SET PATH=E:\bin\TDM-GCC-64\bin;%PATH%
  SET PATH=E:\bin\mingw\mingw64\bin;%PATH%

  PUSHD .
  cd E:\src\goproj\!telegram_bot_telebot\telebot_js
  call npx vite build
  POPD
  rmdir E:\src\goproj\!telegram_bot_telebot\telebot\pkg\vueapp /s /q
  robocopy E:\src\goproj\!telegram_bot_telebot\telebot_js\dist\ E:\src\goproj\!telegram_bot_telebot\telebot\pkg\vueapp /e
  
endlocal
