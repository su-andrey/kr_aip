# Курсовая работа по АИПУ 1 курс

## Документация и иструкция для запуска этого чуда на своих печках

Для начала установим GOLANG на свой ПК глобально по [этой инструкции](https://go.dev/doc/install)
```ВАЖНО иметь версию не старше 1.23```

## Далее нужно завести сервер PostgreSQL

### Запуск на MacOS

[Инструкция для MacOS](https://ploshadka.net/ustanovka-i-podkljuchenie-postgresql-na-mac-os/)  
`Исполнять инструкицю стоит до момента "Как подключиться к PostgreSQL на Mac OS". Остальное нам не интересно.`
```ВАЖНО```
!Запомнить имя пользователя, пароль и название базы данных для доступа на сервер. Они пригодятся нам позже для подключения БД к приложению!  
Далее необходимо выдать права на редактирование таблиц созданному пользователю.  
Заходим от лица суперпользователя: (замените `<SELECTED NAME>` на выбранное для работы имя)  

```bash
GRANT CREATE ON SCHEMA public TO <SELECTED NAME>;
GRANT USAGE ON SCHEMA public TO <SELECTED NAME>;
```

!!ТАКЖЕ НЕ ЗАБЫВАЙТЕ СТАВИТЬ ТОЧКУ С ЗАПЯТОЙ В КОНЦЕ КАЖДОЙ КОМАНДЫ!!
Для работы с БД также понадобится сброс таблиц, команда: (исполняется внутри psql <db_name>)

```bash
DROP TABLE categories, comments, posts, users;
```

### Запуск на Windows

Алгоритм действий схож с запуском на MacOS. Важным отличием является другой алгоритм скачивания, наиболее простым вариантом является скачивания PGAdmin4 и дальнейшая работа с помощью psql (запуск в bin директории установки Postgres).
Также необходимо установить node.js с [официального сайта](https://nodejs.org/en/download). После установки необходимо перезапустить компьютер. Все остальные пункты инструкции такие же как и работе на Mac/Linux. Важно, что по умолчанию в OS Windows из-за настроек безопасности отключена возможность запуска в оболочке PowerShell, эту настройку требуется изменить (открыть терминал PowerShell от имени администратора, выполнить Set-ExecutionPolicy RemoteSigned, разрешить для всех) Set-ExecutionPolicy RemoteSigned

## Создание .env файлов

Далее необходимо создать в проекте собственный файл .env по образцу .env_example, руководствуясь инструкциями внутри него.  
В .env необходимо сформировать собственную ссылку для доступа к БД и указать локальный порт localhost, на котором будет запущено !!go!! -приложение  
Аналогично нужно поступить и с .env- файлом в дирректории `/client`

## Настройка air

Далее для стабильной работы с проектом на GO без перезапуска сервера при малейшем изменении нужно будет накатить расширение Air, которое будет автоматически обновлять сервер при каждом сохранении документа.  
Кратко:

```bash
curl -sSfL https://raw.githubusercontent.com/air-verse/air/master/install.sh | sh -s
sudo nano ~/.zshrc

Дописать в конец файла:
alias air="~/.air"

Далее:
^X              #CTRL+X на маке
Y
ENTER
```

(Работает на MacOS)  
[Полная инструкция по установке](https://github.com/air-verse/air)  
!!Ставить нужно ГЛОБАЛЬНО при помощи второго пункта (Via install.sh) и назначить alias как во второй строчке краткой инструкции!!  
  
!!alias нужно назначать после каждого перезапуска терминала!!  

## Установка зависимостей React

Далее нужно установить для работы с проектом- Node.JS и NPM (должны ставиться одновременно при установке ноды)  
[Установка Node.JS](https://nodejs.org/en/download)  
После установки нужно зайти в дирректорию ./client и запустить комманду `npm install`  

## Запуск

Для работы над проектом потребуется открыть 2 терминала:  
В первом из корневой дирректории проекта вызвать (Для запуска localhost для работы с API)  
```air```
Во втором в дирректории, в которой лежит React (\client) (Для запуска localhost для работы с React)  
```npm run dev```
Перемещение по дирректориям:  
```cd <directoryName>``` для входа в дочернюю дирректорию
```cd ../``` для возврата в родительскую дирректорию
```ls -a``` просмотр дочерних дирректорий текущей, в том числе скрытых

## Структура продукта

```structure
.
├── README.md                  # Документация по проекту
├── air.toml                   # Конфигурация для Air (live reload для Go)
├── client                     # Клиентская часть (React-приложение)
│   ├── public                 # Публичные файлы (иконки, статические ресурсы)
│   ├── src                    # Исходный код клиентского приложения
│   │   ├── assets             # Медиафайлы и иконки
│   │   ├── pages              # Компоненты страниц (React Router)
│   ├── package.json           # Зависимости клиентского приложения
│   ├── vite.config.js         # Конфигурация Vite
├── config                     # Конфигурационные файлы
├── creators                   # Логика создания таблиц в БД
├── database                   # Подключение и управление базой данных
├── handlers                   # Обработчики HTTP-запросов
├── models                     # Определения моделей данных
├── routes                     # Определение маршрутов API
├── go.mod                     # Определение Go-модуля и зависимостей
├── go.sum                     # Контрольные суммы зависимостей
├── main.go                    # Точка входа в серверное приложение
```

## Работа с API

Для работы с API применимы следующие инструкции:

1. Все запросы отправляются через `http://localhost:<PORT>/api + СУФФИКС`  
1. Методы:  
🔹 GET – получение всех объектов и одного конкретного:

```requests
📌 /resource – получить все объекты ресурса
📌 /resource/{id} – получить объект по id
```

🔹 POST – создание нового объекта:

```requests
📌 /resource – создать объект
```

🔹 PUT – обновление существующего объекта:

```requests
📌 /resource/{id} – обновить объект по id
```

🔹 DELETE – удаление объекта:

```requests
📌 /resource/{id} – удалить объект по id
```

1. Суффиксы для получения нужных объектов:  
🔹 Пользователи:

```requests
/users – получить всех пользователей
/users/{id} – получить пользователя по id
```

🔹 Посты:

```requests
/posts – получить все посты
/posts/{id} – получить пост по id
```

🔹 Комментарии:

```requests
/comments – получить все комментарии
/comments/{id} – получить комментарий по id
```

🔹 Категории:

```requests
/categories – получить все категории
/categories/{id} – получить категорию по id
```

1. Структура тел запросов:

1️⃣ Пользователи (users)

```db
{
    "name": "string",
    "password": "string",
    "is_admin": true
}

```db
2️⃣ Категории (categories)

```db
{
    "name": "string"
}

```db
3️⃣ Посты (posts)

```db
{
    "category_id": 1,
    "author_id": 1,
    "name": "string",
    "body": "string"
}
```

4️⃣ Комментарии (comments)

```db
{
    "post_id": 1,
    "author_id": 1,
    "body": "string"
}
```

## Список используемых библиотек

[Fiber V3](https://docs.gofiber.io/next/)  
[GoDotEnv](https://github.com/joho/godotenv)  
[PGX Pool]("github.com/jackc/pgx/v5/pgxpool")  
[React](https://react.dev/)  
[TanStack Query](https://tanstack.com/query/latest)
