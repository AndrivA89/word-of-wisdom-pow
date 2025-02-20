# Task description

Design and implement “Word of Wisdom” tcp server.
- TCP server should be protected from DDOS attacks with the Prof of Work
(https://en.wikipedia.org/wiki/Proof_of_work), the challenge-response protocol should
be used.
- The choice of the POW algorithm should be explained.
- After Prof Of Work verification, server should send one of the quotes from “word of
wisdom” book or any other collection of the quotes.
- Docker file should be provided both for the server and for the client that solves the
POW challenge.

## Описание
Проект реализует TCP-сервер, защищенный от DDoS-атак с помощью алгоритма Proof of Work (PoW). После решения задачи PoW сервер отправляет случайную цитату.

## Структура проекта
- **cmd** - точки входа для сервера и клиента.
- **internal** - логика приложения, включая бизнес-логику, сервисы и репозитории.
- **citations.csv** - файл с цитатами для сервера.

## Как запустить
1. Склонируйте репозиторий.
2. Сборка и запуск:
    ```bash
    make build-server
    make run-server
    ```

## Зависимости
- Go 1.20+
- Docker (для контейнеризации)
