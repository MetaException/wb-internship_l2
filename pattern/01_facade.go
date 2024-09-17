package pattern

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

/*
	Преимущества:
		* Предоставляет простую реализацию работы сложной системой
		* Скрытие внутренней реализации системы
		* Более понятный код

	Недостатки:
		* Сложность поддержки при большом числе методов

	Примеры:
		* Библиотеки для работы с графикой и звуком
		* Работа с файловыми системами
		* Работа с API
*/

type AudioPlayer struct {
	//... some fields
}

func (p *AudioPlayer) Play(path string) {

}

type VideoPlayer struct {
	//... some other fields
}

func (p *VideoPlayer) Play(path string) {

}

type ImagePlayer struct {
	//... some other fields
}

func (p *ImagePlayer) Play(path string) {

}

type PlayerFacade struct {
	audioPlayer *AudioPlayer
	videoPlayer *VideoPlayer
	imagePlayer *ImagePlayer
}

func NewPlayerFacade() *PlayerFacade {
	return &PlayerFacade{
		audioPlayer: &AudioPlayer{},
		videoPlayer: &VideoPlayer{},
		imagePlayer: &ImagePlayer{},
	}
}

func (pf *PlayerFacade) Play(path string, fileType string) {
	if fileType == "image" {
		pf.imagePlayer.Play(path)
	} else if fileType == "audio" {
		pf.audioPlayer.Play(path)
	} else {
		pf.videoPlayer.Play(path)
	}
}
