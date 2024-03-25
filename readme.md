TODO
# Цель, написать bash скрипт
- установлен гит
- сгенерены гит ключи
- клюичи проброшены в гитхаб
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

# Инструкция
```sh
sudo apt install -y git
mkdir -p ~/projects/demoneno4ec
sudo git clone https://github.com/demoneno4ec/ubuntu-local-dev.git ~/projects/demoneno4ec/ubuntu-local-dev
sudo ~/projects/demoneno4ec/ubuntu-local-dev/init.sh
```

# vm machine
- устройства -> Подключить образ диска Дополнительной гостевой ОС
- /media/demoneno4ec/VBox_GAs_7.0.14/autorun.sh
