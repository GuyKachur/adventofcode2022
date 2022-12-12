package day9

import "fmt"

type Pointer struct {
	x, y    int
	visited *map[string]int
	moves   []string
}

func (p *Pointer) Down() {
	p.x = p.x + 1
	(*p.visited)[fmt.Sprintf("%d:%d", p.x, p.y)]++
	p.moves = append(p.moves, "⬇️")
}

func (p *Pointer) Left() {
	p.y = p.y - 1
	(*p.visited)[fmt.Sprintf("%d:%d", p.x, p.y)]++
	p.moves = append(p.moves, "⬅️")

}

func (p *Pointer) Up() {
	p.x = p.x - 1
	(*p.visited)[fmt.Sprintf("%d:%d", p.x, p.y)]++
	p.moves = append(p.moves, "⬆️")

}

func (p *Pointer) Right() {
	p.y = p.y + 1
	(*p.visited)[fmt.Sprintf("%d:%d", p.x, p.y)]++
	p.moves = append(p.moves, "➡️")

}

func (p *Pointer) DownLeft() {
	p.x = p.x + 1
	p.y = p.y - 1
	(*p.visited)[fmt.Sprintf("%d:%d", p.x, p.y)]++
	p.moves = append(p.moves, "↙️")

}

func (p *Pointer) DownRight() {
	p.x = p.x + 1
	p.y = p.y + 1
	(*p.visited)[fmt.Sprintf("%d:%d", p.x, p.y)]++
	p.moves = append(p.moves, "↘️")
}

func (p *Pointer) UpLeft() {
	p.x = p.x - 1
	p.y = p.y - 1
	(*p.visited)[fmt.Sprintf("%d:%d", p.x, p.y)]++
	p.moves = append(p.moves, "↖️")
}

func (p *Pointer) UpRight() {
	p.x = p.x - 1
	p.y = p.y + 1
	(*p.visited)[fmt.Sprintf("%d:%d", p.x, p.y)]++
	p.moves = append(p.moves, "↗️")
}

type Leader struct {
	self      *Pointer
	followers []*Leader
}

func (l *Leader) Down() {
	for _, follower := range l.followers {
		if l.self.x-1 == follower.self.x {
			if l.self.y-1 == follower.self.y {
				follower.DownRight()
			} else if l.self.y+1 == follower.self.y {
				follower.DownLeft()
			} else {
				follower.Down()
			}
		}
	}
	l.self.Down()

}

func (l *Leader) DownRight() {
	for _, follower := range l.followers {
		if l.self.x-1 == follower.self.x && l.self.y+1 == follower.self.y {
			follower.Down()
		} else if l.self.x+1 == follower.self.x && l.self.y-1 == follower.self.y {
			follower.Right()
		} else if l.self.x == follower.self.x && l.self.y-1 == follower.self.y || l.self.y-1 == follower.self.y && l.self.x-1 == follower.self.x || l.self.x-1 == follower.self.x && l.self.y == follower.self.y {
			follower.DownRight()
		}
	}
	l.self.DownRight()

}

func (l *Leader) DownLeft() {
	for _, follower := range l.followers {
		if l.self.x-1 == follower.self.x && l.self.y-1 == follower.self.y {
			follower.Down()
		} else if l.self.x+1 == follower.self.x && l.self.y+1 == follower.self.y {
			follower.Left()
		} else if l.self.x == follower.self.x && l.self.y+1 == follower.self.y || l.self.y+1 == follower.self.y && l.self.x-1 == follower.self.x || l.self.x-1 == follower.self.x && l.self.y == follower.self.y {
			follower.DownLeft()
		}
	}
	l.self.DownLeft()

}

func (l *Leader) Left() {
	for _, follower := range l.followers {
		if l.self.y+1 == follower.self.y {
			if l.self.x-1 == follower.self.x {
				follower.DownLeft()
			} else if l.self.x+1 == follower.self.x {
				follower.UpLeft()
			} else {
				follower.Left()
			}
		}
	}
	l.self.Left()

}

func (l *Leader) UpLeft() {
	for _, follower := range l.followers {
		if l.self.x+1 == follower.self.x && l.self.y-1 == follower.self.y {
			follower.Up()
		} else if l.self.x-1 == follower.self.x && l.self.y+1 == follower.self.y {
			follower.Left()
		} else if l.self.x == follower.self.x && l.self.y+1 == follower.self.y || l.self.y+1 == follower.self.y && l.self.x+1 == follower.self.x || l.self.x+1 == follower.self.x && l.self.y == follower.self.y {
			follower.UpLeft()
		}
	}
	l.self.UpLeft()

}

func (l *Leader) Up() {
	for _, follower := range l.followers {
		if l.self.x+1 == follower.self.x {
			if l.self.y-1 == follower.self.y {
				follower.UpRight()
			} else if l.self.y+1 == follower.self.y {
				follower.UpLeft()
			} else {
				follower.Up()
			}
		}
	}
	l.self.Up()
}

func (l *Leader) UpRight() {
	for _, follower := range l.followers {
		if l.self.x+1 == follower.self.x && l.self.y+1 == follower.self.y {
			follower.Up()
		} else if l.self.x-1 == follower.self.x && l.self.y-1 == follower.self.y {
			follower.Right()
		} else if l.self.x == follower.self.x && l.self.y-1 == follower.self.y || l.self.y-1 == follower.self.y && l.self.x+1 == follower.self.x || l.self.x+1 == follower.self.x && l.self.y == follower.self.y {
			follower.UpRight()
		}
	}
	l.self.UpRight()

}

func (l *Leader) Right() {
	for _, follower := range l.followers {
		if l.self.y-1 == follower.self.y {
			if l.self.x-1 == follower.self.x {
				follower.DownRight()
			} else if l.self.x+1 == follower.self.x {
				follower.UpRight()
			} else {
				follower.Right()
			}
		}
	}
	l.self.Right()

}
