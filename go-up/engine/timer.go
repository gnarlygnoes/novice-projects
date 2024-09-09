package engine

import rl "github.com/gen2brain/raylib-go/raylib"

type Timer struct {
	StartTime float64
	LifeTime  float64
}

func StartTimer(timer *Timer, lifetime float64) {
	timer.StartTime = rl.GetTime()
	timer.LifeTime = lifetime
}

func TimerDone(timer Timer) bool {
	return rl.GetTime()-timer.StartTime >= timer.LifeTime
}

func GetElapsed(timer Timer) float64 {
	return rl.GetTime() - timer.StartTime
}
