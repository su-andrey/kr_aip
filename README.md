# Курсовая работа по АИПУ 1 курс
# Документация и иструкция для запуска этого чуда на своих печках.
Для начала установим GOLANG на свой ПК глобально по [этой инструкции](https://go.dev/doc/install)
```ВАЖНО иметь версию не старше 1.23```

## Далее нужно завести сервер PostgreSQL.
[Инструкция для MacOS](https://ploshadka.net/ustanovka-i-podkljuchenie-postgresql-na-mac-os/)  
`Исполнять инструкицю стоит до момента "Как подключиться к PostgreSQL на Mac OS". Остальное нам не интересно.`
```ВАЖНО```
!Запомнить имя пользователя, пароль и название базы данных для доступа на сервер. Они пригодятся нам позже для подключения БД к приложению!  
Далее необходимо выдать права на редактирование таблиц созданному пользователю.  
Заходим от лица суперпользователя: (замените `<SELECTED NAME>` на выбранное для работы имя)  
```
GRANT CREATE ON SCHEMA public TO <SELECTED NAME>;
GRANT USAGE ON SCHEMA public TO <SELECTED NAME>;
```
!!ТАКЖЕ НЕ ЗАБЫВАЙТЕ СТАВИТЬ ТОЧКУ С ЗАПЯТОЙ В КОНЦЕ КАЖДОЙ КОМАНДЫ!!
Для работы с БД также понадобится сброс таблиц, команда: (исполняется внутри psql <db_name>)
```
DROP TABLE categories, comments, posts, users;
```

## Создание .env файлов
Далее необходимо создать в проекте собственный файл .env по образцу .env_example, руководствуясь инструкциями внутри него.  
В .env необходимо сформировать собственную ссылку для доступа к БД и указать локальный порт localhost, на котором будет запущено !!go!! -приложение  
Аналогично нужно поступить и с .env- файлом в дирректории `/client`

## Настройка air
Далее для стабильной работы с проектом на GO без перезапуска сервера при малейшем изменении нужно будет накатить расширение Air, которое будет автоматически обновлять сервер при каждом сохранении документа.  
Кратко:
```
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

## Работа с API
<b>Для работы с API применимы следующие инструкции:<b>  
1. Все запросы отправляются через `http://localhost:<PORT>/api + СУФФИКС`  
2. Методы:  
🔹 GET – получение всех объектов и одного конкретного:
```
📌 /resource – получить все объекты ресурса
📌 /resource/{id} – получить объект по id
```
🔹 POST – создание нового объекта:
```
📌 /resource – создать объект
```
🔹 PUT – обновление существующего объекта:
```
📌 /resource/{id} – обновить объект по id
```
🔹 DELETE – удаление объекта:
```
📌 /resource/{id} – удалить объект по id
```
3. Суффиксы для получения нужных объектов:  
🔹 Пользователи:
```
/users – получить всех пользователей
/users/{id} – получить пользователя по id
```
🔹 Посты:
```
/posts – получить все посты
/posts/{id} – получить пост по id
```
🔹 Комментарии:
```
/comments – получить все комментарии
/comments/{id} – получить комментарий по id
```
🔹 Категории:
```
/categories – получить все категории
/categories/{id} – получить категорию по id
```
4. Структура тел запросов:
1️⃣ Пользователи (users)
```
{
    "name": "string",
    "password": "string",
    "is_admin": true
}
```
2️⃣ Категории (categories)
```
{
    "name": "string"
}
```
3️⃣ Посты (posts)
```
{
    "category_id": 1,
    "author_id": 1,
    "name": "string",
    "body": "string"
}
```
4️⃣ Комментарии (comments)
```
{
    "post_id": 1,
    "author_id": 1,
    "body": "string"
}
```




## Список используемых библиотек:
[Fiber V3](https://docs.gofiber.io/next/)  
[GoDotEnv](https://github.com/joho/godotenv)  
[PGX Pool](	"github.com/jackc/pgx/v5/pgxpool")  
[React](https://react.dev/)  
[TanStack Query](https://tanstack.com/query/latest)  