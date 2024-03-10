# Goditor

Implementation of different simple text editor like kilo, gokilo and hecto; for learning purposes.

Most tutorials out there use [termbox-go](https://github.com/nsf/termbox-go) to handle interaction with the terminal, for this I've decided to use more "up-to-date" `x/term` avaiable in go (only where it makes sense to me, some parts I might want to implement myself), without going for other higher level libraries like [tcell](https://github.com/gdamore/tcell).

**References**:
- [kilo](https://viewsourcecode.org/snaptoken/kilo/index.html)
- [gokilo](https://gokilo.github.io/)
- [hecto](https://www.flenker.blog/hecto/)
