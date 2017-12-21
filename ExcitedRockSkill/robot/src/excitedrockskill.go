package ExcitedRockSkill

import (
	"math"
	"os"
	"strconv"
	"strings"

	"mind/core/framework/drivers/hexabody"
	"mind/core/framework/log"
	"mind/core/framework/skill"
)

type ExcitedRockSkill struct {
	skill.Base
	stop      chan bool
	direction float64
}

func NewSkill() skill.Interface {
	return &ExcitedRockSkill{
		stop: make(chan bool),
	}
}

func (d *ExcitedRockSkill) OnStart() {
	hexabody.Start()
}

func (d *ExcitedRockSkill) OnClose() {
	hexabody.Close()
}

func (d *ExcitedRockSkill) OnDisconnect() {
	os.Exit(0) // Closes the process when remote disconnects
}

func (d *ExcitedRockSkill) OnRecvString(data string) {

	if data == "start" {
		hexabody.WalkContinuously(0, 0.1)
		//		log.Info.Println(data)

	} else if data == "stopWalk" {
		hexabody.StopWalkingContinuously()
		//		log.Info.Println(data)

	} else if data == "excited" {
		//		log.Info.Println(data)
		ready()
		moveLegs()
	} else {
		params := strings.Split(data, ",")
		speed, _ := strconv.ParseFloat(params[1], 32)
		if speed/10000*4 <= 1.2 {
			hexabody.WalkContinuously(0, speed/10000*4)
			log.Info.Println(speed / 10000 * 4)
		}
	}
}

func ready() {
	hexabody.Stand()
	hexabody.MoveHead(0.0, 1)

	hexabody.MoveLeg(2, hexabody.NewLegPosition().SetCoordinates(-100, 50.0, 70.0), 1)
	hexabody.MoveLeg(5, hexabody.NewLegPosition().SetCoordinates(100, 50.0, 70.0), 1)
	hexabody.MoveJoint(0, 1, 90, 1)
	hexabody.MoveJoint(0, 2, 45, 1)
	hexabody.MoveJoint(1, 1, 90, 1)
	hexabody.MoveJoint(1, 2, 45, 1)
}

func moveLegs() {
	hexabody.MoveJoint(0, 1, 45*math.Sin(math.Pi/180)+70, 1)
	hexabody.MoveJoint(0, 0, 35*math.Cos(math.Pi/180)+60, 1)
	hexabody.MoveJoint(1, 1, 45*math.Cos(math.Pi/180)+70, 1)
}
