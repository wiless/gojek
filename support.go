package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

type Pos struct {
	X, Y int
}

func (p Pos) String() string {
	return fmt.Sprintf("%d:%d", p.X, p.Y)
}

func (p *Pos) ParseXY(str string) error {
	// log.Print("Scanning  ", str)

	n, err := fmt.Sscanf(str, "%d:%d", &p.X, &p.Y)

	if n != 2 {
		return errors.New("Missing format for Position (x:y)")
	}
	return err
}

func ParseInt(str string) (v int) {
	_, err := fmt.Sscanf(str, "%d", &v)
	if err != nil {
		log.Panic("Unable to Parse INT value")
	}
	return v
}

func ParseXYs(s string) (p []Pos, e error) {

	strs := strings.Split(s, ",")
	p = make([]Pos, len(strs))
	for i, v := range strs {
		n, err := fmt.Sscanf(v, "%d:%d", &p[i].X, &p[i].Y)
		if n != 2 {
			return nil, errors.New("Missing format for Position (x:y)")
		}
		if err != nil {
			return nil, err
		}
	}
	return p, nil
}
