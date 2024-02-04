# Запустите Docker контейнер с помощью команды:
```docker run -e urls="https://go.dev/copyright, https://go.dev/tos" -e search="go" [app name]```
# В данной команде:
-e urls="https://go.dev/copyright, https://go.dev/tos": Указывает список URL-адресов для обработки приложением.

-e search="go": Указывает строку для поиска в указанных URL-адресах.

[app name]: Название Docker образа, который нужно запустить.
