api-for-Makefile 🛠️
Описание проекта
api-for-Makefile — это проект, который предоставляет API для автоматизации задач с использованием Makefile. Он создан, чтобы упростить работу с Makefile, позволяя управлять задачами через API-интерфейс.Этот проект идеально подходит для тех, кто хочет интегрировать Makefile в свои приложения или автоматизировать процессы сборки через API.
Основные возможности

API для управления задачами Makefile.
Простая интеграция с любым проектом, использующим Makefile.
Лёгкий способ запускать команды сборки через HTTP-запросы.

Как установить и запустить
Следуй этим шагам, чтобы запустить проект на своём компьютере:

Склонируй репозиторий:git clone https://github.com/LutiDrist/api-for-Makefile.git


Перейди в папку проекта:cd api-for-Makefile


Установи зависимости (если есть, например, для Node.js):npm install


Запусти проект:node server.js

Примечание: Убедись, что у тебя установлен Node.js и Make (или другие инструменты, которые нужны для твоего проекта).

Как использовать

После запуска API будет доступен по адресу http://localhost:3000 (или другой порт, который ты настроил).
Пример запроса для выполнения команды из Makefile:GET /api/run?target=build


