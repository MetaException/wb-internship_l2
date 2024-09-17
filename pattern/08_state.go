package pattern

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

/*

 */

type State interface {
	PressPlayButton(player *Player)
}

type PlayingState struct{}

func (s *PlayingState) PressPlayButton(player *Player) {
	fmt.Println("Ставим на паузу")
	player.SetState(&PausedState{})
}

type PausedState struct{}

func (s *PausedState) PressPlayButton(player *Player) {
	fmt.Println("Возобновляем")
	player.SetState(&PlayingState{})
}

type StoppedState struct{}

func (s *StoppedState) PressPlayButton(player *Player) {
	fmt.Println("Проигрываем")
	player.SetState(&PlayingState{})
}

type Player struct {
	state State
}

func NewPlayer() *Player {
	return &Player{state: &StoppedState{}}
}
func (p *Player) SetState(state State) {
	p.state = state
}

func (p *Player) PressPlayButton() {
	p.state.PressPlayButton(p)
}

func main() {
	player := NewPlayer()

	player.PressPlayButton()

	player.PressPlayButton()

	player.PressPlayButton()

	player.PressPlayButton()
}
