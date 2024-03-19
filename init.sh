# install ru локализацию 
# TODO надо понять, стоит ли игра свеч и возможно проще делать через gui
# 1. добавить региональные настройки settings -> Reqional & Language
# 2. добавить расскладку клавиатуры settings -> Keyboard
sudo apt install language-pack-ru \
	language-pack-gnome-ru

# Установка браузера
wget https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb && \
sudo dpkg -i google-chrome-stable_current_amd64.deb


# Сделать в UI
- настроить языки
- настроить синхронизацию в google chrome
- настроить favorites
