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
