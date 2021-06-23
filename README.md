![License](https://img.shields.io/github/license/p12s/avito-advertising-http-api?style=plastic)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/p12s/avito-advertising-http-api?style=plastic)

**Внимание:** *Тестовое задание найдено на просторах github-а. Для обучения и тренировки, попробовал решить ее в меру своего понимания. На ревью не отправлял, за оптимальность не ручаюсь.*

# Тестовое задание для backend-стажёра в команду Advertising

## Задача
Необходимо создать сервис для хранения и подачи объявлений.  
Объявления должны храниться в базе данных.  
Сервис должен предоставлять API, работающее поверх HTTP в формате JSON.  
Подробнее [здесь](task.md)

## Порядок обработки запроса:
1. по HTTP от клиента приходит запрос
2. передаем его в handler
3. передаем в service
4. передаем в repository. Достаем данные из БД, и отправляем клиенту в обратном порядке

## Выполненные требования
- ✅ Язык программирования Go
- ✅ Простая инструкция для запуска (в идеале — с возможностью запустить через `docker-compose up`, но это необязательно) - **make run**;
* ✅ 3 метода: получение списка объявлений, получение одного объявления, создание объявления;
* ✅ Валидация полей: не больше 3 ссылок на фото, описание не больше 1000 символов, название не больше 200 символов;

## Усложнения
Не обязательно, но задание может быть выполнено с любым числом усложнений:
* ❌ Юнит тесты: постарайтесь достичь покрытия в 70% и больше;
* ❌ Контейнеризация: есть возможность поднять проект с помощью команды `docker-compose up`;
* ✅ Архитектура сервиса описана в виде текста и/или диаграмм (Текст: Все просто - есть **Сервис**, есть **БД**. Запросы приходят от **Клиента** и через **Сервис** отправляются в **БД**)
* ✅ Документация: есть структурированное описание методов сервиса.
