# Сокращатель урлов.

Прошу прощение. Мне захотелось упростить структуру и нейминг. Например, у сократил количество лейеров всего до двух или называл общепринятый repository как model, имея ввиду модель взаимодействия с базой данных. Я знаю, как делают обычно. Еще не завез сваггер, хотя мог бы.

## Порядок запуска

1. Поднятие бд `docker-compose up -d`.
2. Загрузка структуры бд `cat bd_backups/only_structs.sql | docker exec -i postgres_url_shortener psql -U user -d url`.
- Получение структуры бд -> `docker exec -t postgres_url_shortener pg_dump -s -U user url > only_structs.sql`.
3. Запуск сервера `go run main.go` или `go run main.go -d`, чтобы отключить postgres и использовать временную внутреннюю хеш-таблицу для хранения данных.

*** Возможен запуск теста эндпоинтов при запущенно сервере, но там надо прямо в файле править входные данные эндпоинтов. ***
Путь к тестам: `routes/url_test.go`.

## Тесты с помощью kurl

Post: 
- Linux -> `curl -X POST -d "url=https://www.yandex.ru/search/?text=%D0%9C%D0%BE%D0%B3%D1%83+%D0%BB%D0%B8+%D1%8F+%D0%B8%D1%81%D0%BF%D0%BE%D0%BB%D1%8C%D0%B7%D0%BE%D0%B2%D0%B0%D1%82%D1%8C+%D1%87%D0%B0%D1%82+gpt+%D0%B2+%D0%BA%D0%BE%D0%BC%D0%BC%D0%B5%D1%80%D1%87%D0%B5%D1%81%D0%BA%D0%B8%D1%85+%D1%86%D0%B5%D0%BB%D1%8F%D1%85&lr=47" http://localhost:8080/`.
- Windows -> `Invoke-WebRequest -Uri http://localhost:8080/ -Method POST -Body "url=https://www.yandex.ru/search/?text=%D0%9C%D0%BE%D0%B3%D1%83+%D0%BB%D0%B8+%D1%8F+%D0%B8%D1%81%D0%BF%D0%BE%D0%BB%D1%8C%D0%B7%D0%BE%D0%B2%D0%B0%D1%82%D1%8C+%D1%87%D0%B0%D1%82+gpt+%D0%B2+%D0%BA%D0%BE%D0%BC%D0%BC%D0%B5%D1%80%D1%87%D0%B5%D1%81%D0%BA%D0%B8%D1%85+%D1%86%D0%B5%D0%BB%D1%8F%D1%85&lr=47"`.

Get: 
- Linux -> `curl <ваш ответ на post-запрос>`, например, `curl http://localhost:8080/bs6s39i`.
- Windows -> `Invoke-WebRequest -Uri http://localhost:8080/bs6s39i -Method Get`.
