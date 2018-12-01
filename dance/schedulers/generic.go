package schedulers

import (
	"danser/beatmap/objects"
	"danser/bmath"
	"danser/render"
)

//type GenericScheduler struct {
//	cursor *render.Cursor
//	queue  []objects.BaseObject
//	mover  movers.MultiPointMover
//}
//
//func NewGenericScheduler(mover func() movers.MultiPointMover) Scheduler {
//	return &GenericScheduler{mover: mover()}
//}
//
//func (sched *GenericScheduler) Init(objs []objects.BaseObject, cursor *render.Cursor) {
//	sched.cursor = cursor
//	sched.queue = objs
//	sched.mover.Reset()
//	sched.queue = PreprocessQueue(0, sched.queue, settings.Dance.SliderDance)
//	/////////////////////////////////////////////////////////////////////////////////////////////////////
//	// 初始位置 (100, 100)
//	sched.mover.SetObjects([]objects.BaseObject{objects.DummyCircle(bmath.NewVec2d(100, 100), 0), sched.queue[0]})
//	/////////////////////////////////////////////////////////////////////////////////////////////////////
//}
//
//func (sched *GenericScheduler) Update(time int64) {
//	if len(sched.queue) > 0 {
//		move := true
//		for i := 0; i < len(sched.queue); i++ {
//			g := sched.queue[i]
//			if g.GetBasicData().StartTime > time {
//				break
//			}
//
//			move = false
//
//			if time >= g.GetBasicData().StartTime && time <= g.GetBasicData().EndTime {
//
//				sched.cursor.SetPos(g.GetPosition())
//			} else if time > g.GetBasicData().EndTime {
//				if i < len(sched.queue)-1 {
//					sched.queue = append(sched.queue[:i], sched.queue[i+1:]...)
//				} else if i < len(sched.queue) {
//					sched.queue = sched.queue[:i]
//				}
//				i--
//
//				if len(sched.queue) > 0 {
//					sched.queue = PreprocessQueue(i+1, sched.queue, settings.Dance.SliderDance)
//					sched.mover.SetObjects([]objects.BaseObject{g, sched.queue[i+1]})
//				}
//
//				move = true
//			}
//		}
//
//		if move && sched.mover.GetEndTime() >= time {
//			sched.cursor.SetPos(sched.mover.Update(time))
//		}
//
//	}
//	//log.Println(time, sched.cursor.Position)
//}



/////////////////////////////////////////////////////////////////////////////////////////////////////
// 修改Scheduler
type ReplayScheduler struct {
	cursor *render.Cursor
}

func NewReplayScheduler() Scheduler {
	return &ReplayScheduler{}
}

func (sched *ReplayScheduler) Init(objs []objects.BaseObject, cursor *render.Cursor) {
	sched.cursor = cursor
}

func (sched *ReplayScheduler) Update(time int64, position bmath.Vector2d) {
	sched.cursor.SetPos(position)
}
/////////////////////////////////////////////////////////////////////////////////////////////////////
