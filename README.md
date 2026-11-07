# Анализатор размера диска (Go)

[![hexlet-check](https://github.com/danilawarrior/go-project-242/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/danilawarrior/go-project-242/actions)

Консольная утилита, которая считает размер файла или директории с поддержкой рекурсивного обхода, человекочитаемого формата и учёта скрытых файлов.

Учебный проект Хекслета: https://ru.hexlet.io/programs/go

## Стек

- Go
- urfave/cli
- golangci-lint

## Установка

Склонируйте репозиторий:

```bash
git clone https://github.com/danilawarrior/go-project-242.git
cd go-project-242
```

Соберите приложение:

```bash
make build
```

После сборки исполняемый файл будет доступен по пути:

```bash
./bin/hexlet-path-size
```

## Использование

Базовый запуск:

```bash
./bin/hexlet-path-size <path>
```

Доступные флаги:

- `-H`, `--human` — выводить размер в человекочитаемом формате
- `-a`, `--all` — учитывать скрытые файлы и директории
- `-r`, `--recursive` — рекурсивно считать содержимое вложенных директорий

Примеры:

```bash
./bin/hexlet-path-size README.md
```

```bash
./bin/hexlet-path-size --human README.md
```

```bash
./bin/hexlet-path-size --recursive --all .
```

## Демонстрация

[Запись установки и запуска в asciinema](https://asciinema.org/a/pLuDeqfl9j5SF3OP)

---

<details>
<summary>Автоматические тесты Хекслета</summary>

Тесты запускаются на каждый коммит. За запуск отвечает файл `.github/workflows/hexlet-check.yml`.

</details>

## О Хекслете

[Хекслет](https://ru.hexlet.io/) — школа программирования с практическими учебными проектами.