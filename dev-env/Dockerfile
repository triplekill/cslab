FROM archlinux/base
# init image
RUN curl -fsSL https://www.archlinux.org/mirrorlist/?country=CN -o /etc/pacman.d/mirrorlist \
    && pacman -Syy \
    && pacman -S --force pacman-mirrorlist

WORKDIR /home
