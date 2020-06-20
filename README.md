# Тестовая задача для стажировки в my.com

## Формулировка задачи:
Процессу на stdin приходят строки, содержащие URL или названия файлов. Каждый такой URL нужно запросить, каждый файл нужно прочитать, и посчитать кол-во вхождений строки "Go" в ответе. В конце работы приложение выводит на экран общее кол-во найденных строк "Go" во всех источниках данных, например:
```
$ echo -e 'https://golang.org\n/etc/passwd\nhttps://golang.org\nhttps://golang.org' | go run 1.go

Count for https://golang.org: 9
Count for /etc/passwd: 0
Count for https://golang.org: 9
Count for https://golang.org: 9
Total: 27
```
Каждый источник данных должен начать обрабатываться сразу после вычитывания и параллельно с вычитыванием следующего. Источники должны обрабатываться параллельно,но не более k=5 одновременно. Обработчики данных не должны порождать лишних горутин, т.е. если k=1000 а обрабатываемых источников нет, не должно создаваться 1000 горутин. Нужно обойтись без глобальных переменных и использовать только стандартные библиотеки.

## Пояснения к решению:
* Данное решение считает колличество вхождений именно подстроки, а не слова. 
Если где-то в тегах HTML-кода будет использована подстрока Go, то такое вхождение тоже посчитается
* В данном решении k — колличество одновременно исполняемых горутин с логикой, также есть
дополнительные горутины, которые создаются, чтобы писать задания в канал. Другим вариантом
были буферизированные каналы, но тогда бы они были слишком длинными.
* Также было написано немного тестов для запуска в CI
* Настроен CI, который запускает тесты и линтер
