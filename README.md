# Платформа для обучающих курсов по безопасности
Ссылка на продакшен - https://zin-hack-25.antalkon.ru

---

## Локальный запуск бэкенда:
**Для запуска необходим установленный Docker или docker-compose**
1. Склонируйте репозиторий:
```
git clone https://github.com/zinHack-read-specification-team/backend.git
```
2. Сборка и запуск контейнеров
```
docker compose up -d --build       
```
3. **! Важно, после сборки перезапустите контейнер бэкенда (zinhack_app), перед этим убедитесь что все остальный контенерсы запущены**

API бдуте доступно на http://localhost:6611/api/v1

--- 
## Postman коллекция для тестирования API:
https://app.getpostman.com/join-team?invite_code=6b6788766242519117744c637f52dc709a2474e931ee704751d9b6cf12888a0d&target_code=8822919b2a99124213d889268c098dbd



## Спецификация сваггер
Спецификация в /docs/swagger.yaml