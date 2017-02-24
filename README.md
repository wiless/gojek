# gojek
Solution to the GCI17 problem


#Loading the Board
The Board is loaded from the input file `input.txt`. Once the board is successfully loaded, the missiles information is also loaded.
```
fbytes = []byte(`5
	5
	1:1,2:0,2:3,3:4,4:2
	1:3,2:4,3:0,4:1,3:2
	5
	1:3,2:4,3:2,3:2,4:4
	3:2,1:0,0:1,2:2,3:4`)
  
b.LoadBoard(fbytes)

ParseMissileActions(fbytes)
```
#Playing
Alternatively, Missile of each player is Launched using the `LaunchMissile` method of the `Board` object. Which keeps track of Hits/Misses internally

Launches the missile as Player:`PlayerID` at Position `Pos`. The method returns `true` on successfull hit on other player's ship
```
	Board.LaunchMissile(playerID, Pos)
```
#Result
The final status of the Board is returned by the `Result()` method

#Log
The application also prints the logs as the game continues
```
2017/02/25 03:43:58 Game Board  Loaded successfully from  input.txt
2017/02/25 03:43:58 Player 0 Launching Missile 0 : HIT ✓
2017/02/25 03:43:58 Player 1 Launching Missile 0 : HIT ✗
2017/02/25 03:43:58 Player 0 Launching Missile 1 : HIT ✓
2017/02/25 03:43:58 Player 1 Launching Missile 1 : HIT ✗
2017/02/25 03:43:58 Player 0 Launching Missile 2 : HIT ✓
2017/02/25 03:43:58 Player 1 Launching Missile 2 : HIT ✗
2017/02/25 03:43:58 Player 0 Launching Missile 3 : HIT ✗
2017/02/25 03:43:58 Player 1 Launching Missile 3 : HIT ✗
2017/02/25 03:43:58 Player 0 Launching Missile 4 : HIT ✗
2017/02/25 03:43:58 Player 1 Launching Missile 4 : HIT ✓
2017/02/25 03:43:58 Game Ended : Player 1 Won
2017/02/25 03:43:58 File Created  output.txt
```

#NOTE

The board can be set to NOT to load if # of ships > GridSize/2 !! Uncomment these lines from `board.go`
```
if b.TotalShips > b.GridSize/2 {
	return fmt.Errorf("No of Ships %d cannot be > %d", b.TotalShips, b.GridSize/2)
}
```
