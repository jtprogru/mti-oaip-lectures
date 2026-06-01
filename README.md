# Основы алгоритмизации и программирования

[![Сайт курса](https://img.shields.io/badge/site-jtprogru.github.io%2Fmti--oaip--lectures-2ea44f?logo=github)](https://jtprogru.github.io/mti-oaip-lectures/)
[![MkDocs Material](https://img.shields.io/badge/MkDocs-Material-526CFE?logo=materialformkdocs)](https://squidfunk.github.io/mkdocs-material/)
[![Python 3.14](https://img.shields.io/badge/Python-3.14-3776AB?logo=python&logoColor=white)](https://www.python.org/)
[![Go 1.23](https://img.shields.io/badge/Go-1.23-00ADD8?logo=go&logoColor=white)](https://go.dev/)

Материалы курса «Основы алгоритмизации и программирования» (МТИ, ОАиП). Современная переработка: вместо Pascal/Delphi — параллельные примеры на **Python 3.14** и **Go 1.23**, актуальные инструменты (`uv`, `go mod`, VS Code), сквозной практический проект.

## Где смотреть

- **Сайт курса:** <https://jtprogru.github.io/mti-oaip-lectures/>
- **Исходники лекций:** [`docs/topics/`](docs/topics/)
- **Примеры кода:** [`src/python/`](src/python/), [`src/golang/`](src/golang/)
- **Практическое задание:** [`docs/practice/practice.md`](docs/practice/practice.md)
- **Экзаменационные билеты:** [`docs/tickets/index.md`](docs/tickets/index.md)

## Состав курса

14 тем: алгоритмизация → языки и методы → Python (синтаксис, функции, файлы, библиотеки, ООП) → VS Code и отладка → GUI-разработка → стандартная библиотека и шаблоны → HTTP/regex/SQLite/конкурентность → разработка приложений → качество кода и тестирование → **Go (основы и продвинутые темы)**.

Подробнее — на [главной сайта](https://jtprogru.github.io/mti-oaip-lectures/).

## Локальная сборка

```bash
# зависимости MkDocs (через uv)
uv sync

# просмотр с автоперезагрузкой
uv run mkdocs serve

# одноразовая сборка в ./site
uv run mkdocs build
```

Go-примеры:

```bash
cd src/golang/topic-13-basics && go test ./...
cd src/golang/topic-14-advanced && go test ./...
```

## Атрибуция

Курс основан на материалах [kolei/OAP_backup](https://github.com/kolei/OAP_backup) (исходный курс на Pascal/Delphi). Содержание переработано под современный стек, добавлены параллельные примеры на Go, обновлены инструменты и практические задания.
