# Основы алгоритмизации и программирования

## [Тема 1. Основные принципы алгоритмизации и программирования](topics/01-algorithm-basics/index.md)
+ [Лекция 1. Основные понятия алгоритмизации](topics/01-algorithm-basics/lecture-01-basics.md)<br/>
Понятие алгоритма. Свойства алгоритма. Формы записи алгоритмов. Общие принципы построения алгоритмов. Определение сложности работы алгоритмов.
+ [Лекция 2. Основные алгоритмические конструкции](topics/01-algorithm-basics/lecture-02-constructs.md)<br/>
Линейные, разветвляющиеся, циклические алгоритмы. Программы для графического отображения алгоритмов.
+ [Лекции 3 и 4. Логические основы алгоритмизации](topics/01-algorithm-basics/lecture-03-logic.md) (4 часа)<br/>
Основы алгебры логики. Логические операции с высказываниями: конъюнкция, дизъюнкция, инверсия. Законы логических операций. Таблицы истинности.

- Лабораторная работа 1  
Разработка алгоритмов для конкретных задач.
- Лабораторная работа 2  
Использование программ для графического отображения алгоритмов.
- Лабораторная работа 3  
Определение сложности работы алгоритмов.

## [Тема 2. Языки и методы программирования](topics/02-languages-and-methods/index.md)
+ [Лекция 1. Языки и системы программирования](topics/02-languages-and-methods/lecture-01-languages.md)<br/>
Эволюция и классификация языков программирования. Компиляция и интерпретация. Понятие системы программирования. Исходный, объектный и загрузочный модули. IDE (Visual Studio Code).
+ [Лекция 2. Методы программирования и SDLC](topics/02-languages-and-methods/lecture-02-methods.md)<br/>
Структурный, модульный, объектно-ориентированный методы. Общие принципы разработки ПО. Жизненный цикл программного обеспечения. Типы приложений.

## [Тема 3. Программирование на Python (Python Basics)](topics/03-python-basics/index.md)
+ [Лекция 1. Синтаксис, переменные и типы данных](topics/03-python-basics/lecture-01-syntax-types.md)<br/>Лексика, отступы, идентификаторы, ключевые слова; простые типы (`int`, `float`, `str`, `bool`, `None`); операторы и приоритет; параллель Go — статическая типизация, фиксированные размеры целых.
+ [Лекция 2. Управляющие конструкции и ввод-вывод](topics/03-python-basics/lecture-02-control-flow-io.md)<br/>`if`/`elif`/`else`, тернарный, `match`; `while`, `for` + `range`; `break`/`continue`/`else`; `input` / `print`; в Go — `if` с инициализацией, `switch`, `for` как единственный цикл.
+ [Лекция 3. Коллекции: списки, кортежи, словари, множества](topics/03-python-basics/lecture-03-collections.md)<br/>Списки и генераторы списков, кортежи и распаковка, словари (`get`/`setdefault`/`defaultdict`/`Counter`), множества и математические операции; в Go — slices, maps, structs, `map[T]struct{}` как set.
+ [Лекция 4. Машинное представление: числа, кодировки, байты](topics/03-python-basics/lecture-04-encodings-bytes.md)<br/>Системы счисления, дополнительный код, IEEE 754; ASCII/Unicode/UTF-8/UTF-16; `bytes`/`bytearray`/`memoryview` в Python; `string`/`[]byte`/`rune` в Go.

- Лабораторная работа 1<br/>Работа в среде программирования. Реализация построенных алгоритмов.
- Лабораторная работа 2 (4 часа)<br/>Составление программ линейной структуры. Составление программ разветвляющейся структуры.
- Лабораторная работа 3 (4 часа)<br/>Составление программ циклической структуры. Обработка одномерных и двумерных массивов.
- Лабораторная работа 4<br/>Работа со строковыми переменными. Работа с данными типа множество.

## [Тема 4. Подпрограммы и работа с файлами](topics/04-procedures-and-files/index.md)
+ [Лекция 1. Функции и область видимости](topics/04-procedures-and-files/lecture-01-functions.md)<br/>Подпрограммы; `def`; аргументы (позиционные, именованные, `*args`/`**kwargs`, default); LEGB; `global`/`nonlocal`; `lambda`, замыкания; передача по ссылке/значению; в Go — multiple return, variadic, functional options.
+ [Лекция 2. Строки и рекурсия](topics/04-procedures-and-files/lecture-02-strings-recursion.md)<br/>Методы строк, f-strings, кодировки; рекурсия (факториал, Фибоначчи, бинарный поиск, обход дерева), мемоизация, ограничения стека.
+ [Лекция 3. Работа с файлами и форматами данных](topics/04-procedures-and-files/lecture-03-files.md)<br/>Открытие/закрытие (`with`/`defer`), `pathlib` vs `os.path`, обход каталогов; форматы CSV/JSON/YAML/TOML/INI; `pickle` и его опасности.
+ [Лекция 4. Бинарные файлы и произвольный доступ](topics/04-procedures-and-files/lecture-04-binary-random-access.md)<br/>`seek`/`tell`, `struct` (Python) и `encoding/binary` (Go), чтение заголовка WAV, `mmap` для больших файлов.

- Лабораторная работа 1 (4 часа)<br/>Организация и использование процедур.
- Лабораторная работа 2 (4 часа)<br/>Организация и использование функций.
- Лабораторная работа 3 (4 часа)<br/>Работа с файлами последовательного и произвольного доступа.

## [Тема 5. Библиотеки и модули](topics/05-libraries-and-modules/index.md)
+ [Лекция 1. Модули, пакеты и управление зависимостями](topics/05-libraries-and-modules/lecture-01-modules-packages.md) (4 часа)<br/>Модули и пакеты Python (`__init__.py`, `__main__.py`, `if __name__ == "__main__"`); `sys.path`; `pip` и современный `uv`; `pyproject.toml`, lock-файлы; виртуальные окружения; в Go — `go mod`, `go.sum`, vendoring; FFI через `ctypes` и `cgo`.

- Лабораторная работа 1 (4 часа)<br/>Программирование модуля. Создание библиотеки подпрограмм.

## [Тема 6. Основные принципы объектно-ориентированного программирования](topics/06-oop-principles/index.md)
+ [Лекция 1. Базовые понятия ООП. Инкапсуляция, наследование, полиморфизм](topics/06-oop-principles/lecture-01-oop-basics.md)<br/>
История ООП. Класс, объект, интерфейс. Принципы инкапсуляции, наследования, полиморфизма. Реализация ООП в Python и Go.
+ [Лекция 2. Событийно-управляемая модель и компонентно-ориентированный подход](topics/06-oop-principles/lecture-02-event-driven.md)<br/>
Генераторы и yield в Python, `asyncio`. Горутины и каналы в Go. Компонентно-ориентированное программирование.

## [Тема 7. Среда разработки: Visual Studio Code](topics/07-ide-vscode/index.md)
+ [Лекция 1. VS Code: установка, расширения, отладка](topics/07-ide-vscode/lecture-01-setup.md)<br/>Переносной режим, русификация, расширение Python (Pylance, Ruff, debugpy), расширение Go (gopls, dlv), горячие клавиши.
+ [Лекция 2. Параметры командной строки](topics/07-ide-vscode/lecture-02-cli-args.md)<br/>`sys.argv` и `os.Args`, `argparse`/`typer`, `flag`/`cobra`, отладка с `args` в `launch.json`.
+ [Лекция 3. Ошибки, исключения и декораторы](topics/07-ide-vscode/lecture-03-errors-decorators.md)<br/>`try/except/finally`, `raise from`, собственные исключения, декораторы; `error` как значение в Go, `errors.Is/As`, `panic/recover`, middleware.

- Лабораторная работа 1 (10 часов)<br/>Создание простого проекта по индивидуальным заданиям.

## [Тема 8. Этапы разработки приложения](topics/08-app-development-stages/index.md)
+ [Лекция 1. GUI на Tkinter](topics/08-app-development-stages/lecture-01-tkinter.md)<br/>Встроенный модуль `tkinter`, виджеты, упаковщики (`pack`/`grid`/`place`), события, `ttk`. Параллель: Fyne в Go.
+ [Лекция 2. GUI на PyQt и Qt Designer](topics/08-app-development-stages/lecture-02-pyqt.md)<br/>PyQt6/PySide6, визуальный редактор Qt Designer, сигналы и слоты, `.ui`-файлы.
+ [Лекция 3. Web-обёртка как UI: pywebview, CEF Python, Wails](topics/08-app-development-stages/lecture-03-web-ui.md)<br/>Встроенный браузер как окно приложения, биндинги Python/Go ↔ JavaScript.

## [Тема 9. Иерархия классов](topics/09-class-hierarchies/index.md)
+ [Лекция 1. Стандартная библиотека Python и Go](topics/09-class-hierarchies/lecture-01-stdlib.md)<br/>Встроенные функции, основные модули `sys`, `os`, `datetime`, `collections`, `contextlib`, `re`, `json`, `sqlite3`; параллели в Go.
+ [Лекция 2. Шаблоны проектирования](topics/09-class-hierarchies/lecture-02-patterns.md) (4 часа)<br/>Порождающие, структурные, поведенческие шаблоны с примерами на Python и Go.

- Лабораторная работа 1 (6 часов)   
Создание объектно-ориентированного приложения по индивидуальным заданиям.

## [Тема 10. Стандартные модули Python](topics/10-standard-modules/index.md)
+ [Лекция 1. HTTP-клиент и HTTP-сервер](topics/10-standard-modules/lecture-01-http.md)<br/>`urllib`, `requests`, `net/http` в Go.
+ [Лекция 2. Регулярные выражения](topics/10-standard-modules/lecture-02-regex.md)<br/>Модуль `re` в Python, `regexp` в Go.
+ [Лекция 3. SQLite](topics/10-standard-modules/lecture-03-sqlite.md)<br/>Встраиваемая БД, защита от SQL-инъекций, миграции.
+ [Лекция 4. Конкурентность и потоки](topics/10-standard-modules/lecture-04-concurrency.md)<br/>GIL, `threading`, `asyncio`, горутины и каналы в Go.

- Лабораторная работа 1 (4 часа)   
Создание проекта с использованием компонентов для работы с текстом.
- Лабораторная работа 2 (6 часов)   
Создание проекта с использованием компонентов стандартных диалогов и системы меню.

## [Тема 11. Разработка приложений](topics/11-application-development/index.md)
+ [Лекция 1. Логирование](topics/11-application-development/lecture-01-logging.md)<br/>Модуль `logging` в Python, `log/slog` в Go.
+ [Лекция 2. Django — основы фреймворка](topics/11-application-development/lecture-02-django-basics.md)<br/>MVT, модели и миграции, ORM, маршруты, views, шаблоны, готовая админка.
+ [Лекция 3. Django — формы, сигналы и REST API](topics/11-application-development/lecture-03-django-events-api.md)<br/>`ModelForm`, аутентификация, сигналы, DRF, JWT.
+ [Лекция 4. Компиляция в standalone](topics/11-application-development/lecture-04-standalone.md)<br/>PyInstaller, Nuitka, Briefcase для Python; `go build` и кросс-компиляция; Docker.

- Лабораторная работа 1 (4 часа)   
Разработка оконного приложения. 
- Лабораторная работа 2 (4 часа)   
Разработка оконного приложения с несколькими формами.

## [Тема 12. Качество кода и тестирование](topics/12-code-quality-and-testing/index.md)
+ [Лекция 1. Качество кода и тестирование](topics/12-code-quality-and-testing/lecture-01-testing.md)<br/>
`unittest` и `pytest` в Python, встроенный `testing` в Go, параметризация, fixtures, табличные тесты, покрытие.

- Лабораторная работа 2 (4 часа)   
Оформление, отладка кода программы.

## [Тема 13. Основы Go (Golang Basics)](topics/13-golang-basics/index.md)
+ [Лекция 1. Знакомство с Go](topics/13-golang-basics/lecture-01-intro.md)<br/>История, установка, `go mod`, первая программа, инструментарий (`gofmt`, `golangci-lint`, VS Code).
+ [Лекция 2. Переменные, типы и управляющие конструкции](topics/13-golang-basics/lecture-02-types-control-flow.md)<br/>`var`/`const`/`iota`, базовые типы, `if`/`for`/`switch`, `defer`, указатели.
+ [Лекция 3. Функции, ошибки и `panic`/`recover`](topics/13-golang-basics/lecture-03-functions-errors.md)<br/>Multiple return, variadic, functional options, `error` как значение, `errors.Is`/`errors.As`.
+ [Лекция 4. Композитные типы](topics/13-golang-basics/lecture-04-composite-types.md)<br/>Массивы, срезы (slice header), карты, структуры, теги, embedding.
+ [Лекция 5. Методы и интерфейсы](topics/13-golang-basics/lecture-05-methods-interfaces.md)<br/>Receiver, structural typing, `any`, type switch, канонические интерфейсы stdlib, дженерики.
+ [Лекция 6. Пакеты и тестирование](topics/13-golang-basics/lecture-06-packages-testing.md)<br/>`internal/`, `go mod`, `go.sum`, `testing`, табличные тесты, бенчмарки, fuzz.

## [Тема 14. Продвинутый Go (Golang Advanced)](topics/14-golang-advanced/index.md)
+ [Лекция 1. Конкурентность](topics/14-golang-advanced/lecture-01-concurrency.md)<br/>Горутины, каналы, `select`, `sync.WaitGroup`/`Mutex`/`Once`, `sync/atomic`, worker pool, fan-in/fan-out, pipeline.
+ [Лекция 2. Context](topics/14-golang-advanced/lecture-02-context.md)<br/>`context.Context`, `WithCancel`/`WithTimeout`/`WithDeadline`/`WithValue`, распространение и отмена.
+ [Лекция 3. HTTP-клиент и HTTP-сервер](topics/14-golang-advanced/lecture-03-http.md)<br/>`net/http`, production-клиент с `Timeout`, роутинг Go 1.22+, middleware, graceful shutdown.
+ [Лекция 4. Файлы, JSON и БД](topics/14-golang-advanced/lecture-04-files-json-db.md)<br/>`os`/`io`/`bufio`, `embed.FS`, `encoding/json` с тегами, `database/sql`, транзакции.
+ [Лекция 5. Бенчмарки и pprof](topics/14-golang-advanced/lecture-05-benchmarks-pprof.md)<br/>`Benchmark*`, `benchstat`, `pprof`, `net/http/pprof`, `-race`, escape analysis, `sync.Pool`.

[Задание на практику](practice/practice.md) — большой сквозной проект «распределённая система оплаты услуг» с GUI, HTTP, БД, логированием и тестами.

Дополнительно: [шпаргалка по типам и scope](practice/cheatsheet.md), [системы контроля версий и Git](practice/version-control.md).

## [Экзаменационные билеты](tickets/index.md)