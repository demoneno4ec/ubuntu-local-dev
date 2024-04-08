#!/bin/bash
# Настройки для скрипта
## colors
Green='\033[0;32m' # Green
Red='\033[0;31m'   # Red
Cyan='\033[0;36m'  # Cyan

# Сбор пользовательских настроек
echo "# Hello, please introduce yourself."

echo -en "${Cyan} ## Your name (${SUDO_USER}): ${NC}"
read -r name

echo -en "${Cyan} # Your email: ${NC}"
read -r email

# Первоначальное обновление зависимостей и пакетов
printf "${Green}# Обновление зависимостей \n${NC}"
apt update && \
	apt upgrade -y

# установка приложений (apt)
printf "${Green}# Установка приложений (apt) \n${NC}"
apt install -y \
	git \
	vim \
	curl \
	gnome-tweaks \
	software-properties-common

# добавление репозиториев
add-apt-repository -y ppa:ondrej/php && \
	apt update


# Установка приложений не из apt
printf "${Green}# Установка приложений (not apt) \n${NC}"
## Установка браузера Google Chrome
printf "${Green}## Установка Google Chrome \n${NC}"
wget https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb && \
sudo dpkg -i google-chrome-stable_current_amd64.deb
## Установка старых версий php
apt install -y \
	php7.2 \
	php7.2-{common,cli,fpm} \
	php7.4 \
	php7.4-{common,cli,fpm}
## Установка go
wget https://go.dev/dl/go1.22.2.linux-amd64.tar.gz && \
	tar -C /usr/local -xzf go1.22.2.linux-amd64.tar.gz && \
	export PATH=$PATH:/usr/local/go/bin


# Конфигурирование приложений
printf "${Green}# Конфигурирование приложений \n${NC}"
## git
printf "${Green}## git \n${NC}"
git config --global user.email $email
git config --global user.name $name
git config --global core.editor 'vim'
ssh-keygen -b 4096 -t rsa -f ~/.ssh/id_rsa -q -N ''
git config --global --list

## Права доступы
### TODO Можем быть можно переделать на ${SUDO_USER}
chown -R demoneno4ec:demoneno4ec ~/projects
