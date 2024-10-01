package race

import "math/rand"

func MoveObstacleCarUpward(car *Car, y int) {
	car.Lane = rand.Intn(3)
	car.Y = y
}

func HandleRivalCars(game *CarRacingGame) {
	for i := range game.RivalCars {
		HandleScoreForOvertake(game, i)
		HandleRivalCarGetsCloser(game, i)
		HandleOvertakenRivalCar(game, i)
	}
}

func HandleOvertakenRivalCar(game *CarRacingGame, i int) {
	if game.RivalCars[i].Y >= game.TrackHeight {
		MoveObstacleCarUpward(&game.RivalCars[i], 0)
	}
}

func HandleRivalCarGetsCloser(game *CarRacingGame, i int) {
	game.RivalCars[i].Y++
}

func HandleScoreForOvertake(game *CarRacingGame, i int) {
	if game.RivalCars[i].Y == game.Player.Y {
		game.Score++
	}
}

func CheckForCollision(game *CarRacingGame) {
	for _, car := range game.RivalCars {
		if car.Lane == game.Player.Lane && car.Y == game.Player.Y {
			game.State = KCollide
			return
		}
	}
}
