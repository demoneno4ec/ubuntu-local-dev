TODO
# День 3
## Цель, написать bash скрипт
- установлен гит
- сгенерены гит ключи
- ключи проброшены в гитхаб
- спулен скрипт
- скрипт выполняет установку текущих программ
	- браузер
	- gnome-tweaks
	- vim
- скрипт выполняет конфигурацию системы
	- порядок фаворитов (Делать вручную)
	- gnome tweaks настройка клавиш (посмотреть как можно сделать)
	- настройку виртуальной машины (гостевой режим) ??
	- локализация (делать вручную)
## Итог
- фавориты, делать вручную
- gnome tweaks, делать вручную
- настройка виртуальной машины, делать вручную
- локализация, делать вручную

# День 4
- установить
	- fish
	- php7.2 +
	- php7.4 +
	- php8.2
	- docker
	- docker compose plugin
	- composer
	- jetbrains toolbox
	- telegram
	- postman
	- copyq
	- figma linux
	- discord
	- vlc media player
	- sublime text
	- portainer
	- obs
- настроить
	- fish
	- terminal
	- ~/.ssh/config
	- ~/config/fish
- описать
	- расширения для пыхи
	- установку рабочих и медиа программ
	- автозагрузку
	- добавление ppa 


--------

# vm machine
- устройства -> Подключить образ диска Дополнительной гостевой ОС
- /media/demoneno4ec/VBox_GAs_7.0.14/autorun.sh

# Инструкция
## Вариант #1 с bash
```sh
sudo apt install -y git
mkdir -p ~/projects/demoneno4ec
sudo git clone https://github.com/demoneno4ec/ubuntu-local-dev.git ~/projects/demoneno4ec/ubuntu-local-dev
sudo ~/projects/demoneno4ec/ubuntu-local-dev/init.sh
```

## Вариант #2 с go
```sh
wget https://go.dev/dl/go1.22.2.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.22.2.linux-amd64.tar.gz && \
    export PATH=$PATH:/usr/local/go/bin
sudo git clone https://github.com/demoneno4ec/ubuntu-local-dev.git ~/projects/demoneno4ec/ubuntu-local-dev
sudo go run ~/projects/demoneno4ec/ubuntu-local-dev/main.go
```
### Конфигурация config.yaml
- если нужно перенести ssh ключи, то указать в config.yaml->settings->git->ssh_dir путь до папки с приватным и публичным ключём.
	- В противном случае, пара ключей будет сгенерирована с нуля