# Сервис для бронирования отелей
## Запуск
```
 sh ./installpkgs.sh
 go run main.go
```
## Использование
```
/addRoom [price - required, discription - optional] - создаёт номер в отеле и отдаёт его ID
/delRoom [id - required] - удаляет номер и все брони связанные с ним
/getRooms - отдаёт все существующие номера в отеле

/addBooking [id, datestart, dateend - required] - бронирует номер в отеле, отдаёт id бронирования
/delBooking [id - required] - удаляет бронь по id
/getBooking [id - required] - отдаёт все существующие брони в номере, сортирует по дате начала
```