# Сокращатель ссылок

Необходимо реализовать сервис, который должен предоставлять API по созданию сокращенных ссылок следующего формата:
Ссылка должна быть уникальной и на один оригинальный URL должна ссылаться только одна сокращенная ссылка.
Ссылка должна быть длинной 10 символов
Ссылка должна состоять из символов латинского алфавита в нижнем и верхнем регистре, цифр и символа _ (подчеркивание)
Сервис должен быть написан на Go и принимать следующие запросы по gRPC:
1. Метод Create, который будет сохранять оригинальный URL в базе и возвращать сокращённый
2. Метод Get, который будет принимать сокращённый URL и возвращать оригинальный URL
