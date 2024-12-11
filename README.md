# VS Code

# Extentions:
- Go
- Code Run

# Setup Go extention
View -> Command Palette
>goinstall
Кликаем на Go: Install/Update Tools, появится список вспомогательных утилит, которые необходимо установить. Для этого отмечаем все флажки и кликаем на OK справа вверху

# Setup IDE
View -> Command Palette
>Preferences: Open User Settings (JSON)

settings.json

{
    "go.useLanguageServer": true, // при наборе кода, активирует подсказки, навигацию по коду, форматирование
    "editor.codeActionsOnSave": {}, // действия при сохранении файла, ничего не указываем чтобы не было конфликтов при сохранении go файлов
    "go.formatTool": "gofumpt", // утилита для форматирования go файлов.https://github.com/mvdan/gofumpt
    "[go]": {
        "editor.formatOnSave": true, // при сохранении файла, будет выполняться форматирование при помощи утилиты gofumpt
        "editor.codeActionsOnSave": {
            "source.organizeImports": "always" // при сохранении файла, будет осуществляться организация, группировка, сортировка импортированных пакетов из раздела import
        }
    },
    "code-runner.saveFileBeforeRun": true, // сохраняет файл перед запуском
    "code-runner.clearPreviousOutput": true, // очищает результат предыдущего запуска перез запуском
    "code-runner.runInTerminal": true // запуск в терминале VSCode
}

