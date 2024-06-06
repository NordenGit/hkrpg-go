package player

import (
	"github.com/gucooing/hkrpg-go/pkg/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

func newMission() *spb.Mission {
	return &spb.Mission{}
}

func (g *GamePlayer) GetMission() *spb.Mission {
	db := g.GetBasicBin()
	if db.Mission == nil {
		db.Mission = newMission()
	}
	return db.Mission
}

func (g *GamePlayer) GetMainMission() *spb.MainMission {
	db := g.GetMission()
	if db.MainMission == nil {
		db.MainMission = &spb.MainMission{}
	}
	return db.MainMission
}

func (g *GamePlayer) GetMainMissionList() map[uint32]*spb.MissionInfo {
	db := g.GetMainMission()
	if db.MainMissionList == nil {
		db.MainMissionList = make(map[uint32]*spb.MissionInfo)
	}
	return db.MainMissionList
}

func (g *GamePlayer) GetSubMainMissionList() map[uint32]*spb.MissionInfo {
	db := g.GetMainMission()
	if db.SubMissionList == nil {
		db.SubMissionList = make(map[uint32]*spb.MissionInfo)
	}
	return db.SubMissionList
}

func (g *GamePlayer) GetSubMainMissionById(id uint32) *spb.MissionInfo {
	db := g.GetSubMainMissionList()
	return db[id]
}

func (g *GamePlayer) GetFinishMainMissionList() map[uint32]*spb.MissionInfo {
	db := g.GetMainMission()
	if db.FinishMainMissionList == nil {
		db.FinishMainMissionList = make(map[uint32]*spb.MissionInfo)
	}
	return db.FinishMainMissionList
}

func (g *GamePlayer) GetFinishSubMainMissionList() map[uint32]*spb.MissionInfo {
	db := g.GetMainMission()
	if db.FinishSubMissionList == nil {
		db.FinishSubMissionList = make(map[uint32]*spb.MissionInfo)
	}
	return db.FinishSubMissionList
}

func (g *GamePlayer) GetFinishSubMainMissionById(id uint32) *spb.MissionInfo {
	db := g.GetFinishSubMainMissionList()
	return db[id]
}

// 将主任务转成完成状态
func (g *GamePlayer) UpMainMission(mainMissionId uint32) bool {
	mainMissionList := g.GetMainMissionList()
	subMission := mainMissionList[mainMissionId]
	finishMainMissionList := g.GetFinishMainMissionList()
	if subMission == nil {
		return false
	}

	finishMainMissionList[mainMissionId] = &spb.MissionInfo{
		MissionId: subMission.MissionId,
		Progress:  subMission.Progress + 1,
		Status:    spb.MissionStatus_MISSION_FINISH,
	}
	delete(mainMissionList, mainMissionId)
	conf := gdconf.GetMainMissionById(mainMissionId) // 发放奖励
	if conf != nil {
		rewardConf := gdconf.GetRewardDataById(conf.RewardID)
		if rewardConf != nil {
			pileItem := make([]*Material, 0)
			pileItem = append(pileItem, &Material{
				Tid: Hcoin,
				Num: rewardConf.Hcoin,
			})
			for _, data := range rewardConf.Items {
				pileItem = append(pileItem, &Material{
					Tid: data.ItemID,
					Num: data.Count,
				})
			}
			g.AddItem(pileItem)
		}
	}
	return true
}

// 将子任务转成完成状态
func (g *GamePlayer) UpSubMainMission(subMissionId uint32) bool {
	subMainMissionList := g.GetSubMainMissionList()
	subMission := subMainMissionList[subMissionId]
	finishSubMainMissionList := g.GetFinishSubMainMissionList()
	if subMission == nil {
		return false
	}

	finishSubMainMissionList[subMissionId] = &spb.MissionInfo{
		MissionId: subMission.MissionId,
		Progress:  subMission.Progress + 1,
		Status:    spb.MissionStatus_MISSION_FINISH,
	}
	delete(subMainMissionList, subMissionId)

	triggerMissions := map[uint32]uint32{
		100040115: 100040116,
		100040116: 100040115,
		100040121: 100040122,
		100040122: 100040121,
	}
	if triggerID, ok := triggerMissions[subMissionId]; ok && finishSubMainMissionList[triggerID] == nil {
		g.UpSubMainMission(triggerID)
	}

	return true
}

func (g *GamePlayer) TalkStrSubMission(talkStr string) {
	for id := range g.GetSubMainMissionList() {
		conf := gdconf.GetSubMainMissionById(id)
		if conf == nil {
			continue
		}
		if conf.ParamStr1 == talkStr {
			g.FinishSubMission(id) // 完成子任务
		}
	}
}

// 处理战斗任务
func (g *GamePlayer) UpBattleSubMission(req *proto.PVEBattleResultCsReq) {
	db := g.GetBattleBackupById(req.BattleId)
	if db.EventId == 0 {
		return
	}
	for id := range g.GetSubMainMissionList() {
		conf := gdconf.GetSubMainMissionById(id)
		if conf == nil {
			continue
		}
		switch conf.FinishType {
		case "StageWin":
			if req.EndStatus == proto.BattleEndStatus_BATTLE_END_WIN && db.EventId == conf.ParamInt1 {
				g.FinishSubMission(id)
			}
		}
	}
}

// 处理交互任务
func (g *GamePlayer) UpInteractSubMission(db *spb.BlockBin) {
	for id := range g.GetSubMainMissionList() {
		conf := gdconf.GetSubMainMissionById(id)
		if conf == nil {
			continue
		}
		switch conf.FinishType {
		case "PropState":
			propState := g.GetPropState(db, conf.ParamInt1, conf.ParamInt2, "")
			if conf.ParamInt3 == propState {
				g.FinishSubMission(id)
			}
		}
	}
}

// 处理删除实体任务
func (g *GamePlayer) UpKillMonsterSubMission(me *MonsterEntity) {
	for id := range g.GetSubMainMissionList() {
		conf := gdconf.GetSubMainMissionById(id)
		if conf == nil {
			continue
		}
		switch conf.FinishType {
		case "KillMonster":
			if me.GroupId == conf.ParamInt1 && me.InstId == conf.ParamInt2 {
				g.FinishSubMission(id)
			}
		}
	}
}

// 处理创建角色任务
func (g *GamePlayer) CreateCharacterSubMission() {
	for id := range g.GetSubMainMissionList() {
		conf := gdconf.GetSubMainMissionById(id)
		if conf == nil {
			continue
		}
		switch conf.FinishType {
		case "CreateCharacter":
			g.FinishSubMission(id)
		}
	}
}

// 完成列表中的任务即可
func (g *GamePlayer) ListContain(id uint32) {
	db := g.GetSubMainMissionList()[id]
	conf := gdconf.GetSubMainMissionById(id)
	if conf == nil || db == nil {
		return
	}
	db.Progress = 0
	finishSubMainMissionList := g.GetFinishSubMainMissionList()
	finishMainMissionList := g.GetFinishMainMissionList()
	for _, paramInt := range conf.ParamIntList {
		if finishSubMainMissionList[paramInt] != nil || finishMainMissionList[paramInt] != nil {
			db.Progress++
		}
	}
	if db.Progress == uint32(len(conf.ParamIntList)) {
		db.Progress = conf.Progress
		// 完成任务
		g.FinishSubMission(id)
	} else {
		g.MissionPlayerSyncScNotify([]uint32{id}, make([]uint32, 0), make([]uint32, 0))
	}
}

func (g *GamePlayer) MessagePerformSectionFinish(sectionId uint32) { // 处理npc聊天任务
	for id := range g.GetSubMainMissionList() {
		conf := gdconf.GetSubMainMissionById(id)
		if conf == nil {
			continue
		}
		if conf.FinishType == "MessagePerformSectionFinish" {
			if conf.ParamInt1 == sectionId {
				g.FinishSubMission(id)
			}
		}
	}
}

// 完成由服务端完成的任务
func (g *GamePlayer) AutoServerFinishMission() {
	for id := range g.GetSubMainMissionList() {
		conf := gdconf.GetSubMainMissionById(id)
		if conf == nil {
			continue
		}
		switch conf.FinishType {
		case "GetTrialAvatar": // 加载试用角色
			g.GetTrialAvatar(conf.ParamInt1)
			g.FinishSubMission(id)
		case "DelTrialAvatar": // 卸载试用角色
			g.DelTrialAvatar(conf.ParamInt1)
			g.FinishSubMission(id)
		case "EnterFloor": // 传送
			if entryId, ok := gdconf.GetEntryId(id); ok {
				g.EnterSceneByServerScNotify(entryId, 0)
			}
			g.FinishSubMission(id)
		case "SubMissionFinishCnt": // 完成列表中的任务即可
			g.ListContain(id)
		case "MessagePerformSectionFinish": // 对话框显示
			g.AddMessageGroup(conf.ParamInt1)
		}
	}
}

// 接取任务后完成服务端动作（不结束任务
func (g *GamePlayer) AutoServerMissionFinishAction() {
	for id := range g.GetSubMainMissionList() {
		conf := gdconf.GetSubMainMissionById(id)
		if conf == nil {
			continue
		}
		if conf.FinishActionList == nil {
			continue
		}
		for _, finishAction := range conf.FinishActionList {
			switch finishAction.FinishActionType {
			case "ChangeLineup": // 强制更新队伍
				g.NewTrialLine(finishAction.FinishActionPara) // 设置队伍角色
			case "Recover": // 恢复队伍
				g.RecoverLine()
			}
		}
	}
}

// 完成子任务并拉取下一个任务和通知
func (g *GamePlayer) FinishSubMission(missionId uint32) {
	// 先完成子任务
	if !g.UpSubMainMission(missionId) {
		return
	}
	nextList := make([]uint32, 0)
	curFinishMain := make([]uint32, 0)
	finisSub := make([]uint32, 0)
	g.Send(cmd.StartFinishSubMissionScNotify, &proto.StartFinishSubMissionScNotify{SubMissionId: missionId})
	finisSub = append(finisSub, missionId)
	finishSubMainMissionList := g.GetFinishSubMainMissionList()
	subMainMissionList := g.GetSubMainMissionList()
	subMissionConf := gdconf.GetSubMainMissionById(missionId)
	if subMissionConf == nil {
		return
	}
	conf := gdconf.GetGoppMainMissionById(subMissionConf.MainMissionID)
	if conf == nil {
		return
	}
	for _, finishSubMission := range conf.FinishSubMissionList {
		if missionId == finishSubMission {
			//  完成该主线任务，并接取下一个主线任务
			curFinishMain = append(curFinishMain, conf.MainMissionID)
			g.UpMainMission(conf.MainMissionID) // 结束主任务
		}
	}
	for _, confSubMission := range conf.SubMissionList {
		var isNext = false
		if subMainMissionList[confSubMission.ID] != nil || finishSubMainMissionList[confSubMission.ID] != nil {
			continue
		}
		for _, takeParamId := range confSubMission.TakeParamIntList {
			if finishSubMainMissionList[takeParamId] != nil {
				isNext = true
			} else {
				isNext = false
				break
			}
		}
		if isNext {
			nextList = append(nextList, confSubMission.ID)
			subMainMissionList[confSubMission.ID] = &spb.MissionInfo{
				MissionId: confSubMission.ID,
				Progress:  0,
				Status:    spb.MissionStatus_MISSION_DOING,
			}
		}
	}
	// 通知状态
	g.MissionPlayerSyncScNotify(nextList, finisSub, curFinishMain) // 发送通知

	g.ReadyMission()
}

// 任务检查
func (g *GamePlayer) ReadyMission() {
	g.ReadyMainMission()              // 主线检查
	g.AutoServerMissionFinishAction() // 任务自动行为检查
	g.AutoServerFinishMission()       // 检查服务端任务动作
	g.AutoEntryGroup()                // 检查场景上是否有实体需要卸载/加载
}

// 主线检查
func (g *GamePlayer) ReadyMainMission() {
	mainMissionList := g.GetMainMissionList()
	finishMainMissionList := g.GetFinishMainMissionList()
	subMainMissionList := g.GetSubMainMissionList()
	finishSubMainMissionList := g.GetFinishSubMainMissionList()
	conf := gdconf.GetMainMission()
	var nextList []uint32
	for id, mission := range conf {
		if g.IsReceiveMission(mission, mainMissionList, finishMainMissionList) {
			goppConf := gdconf.GetGoppMainMissionById(id)
			if goppConf == nil {
				continue
			}
			// 这里接取了主线
			if id == 1000300 {
				g.AddAvatar(1003, proto.AddAvatarSrcState_ADD_AVATAR_SRC_NONE)
				g.GetTrialAvatar(1003)
			}
			mainMissionList[id] = &spb.MissionInfo{
				MissionId: id,
				Progress:  0,
				Status:    spb.MissionStatus_MISSION_DOING,
			}
			// 接取该主线子任务
			for _, subInfo := range goppConf.SubMissionList {
				if finishSubMainMissionList[subInfo.ID] != nil {
					continue
				}
				if subInfo.TakeType == "Auto" {
					nextList = append(nextList, subInfo.ID)
					subMainMissionList[subInfo.ID] = &spb.MissionInfo{
						MissionId: subInfo.ID,
						Progress:  0,
						Status:    spb.MissionStatus_MISSION_DOING,
					}
				}
			}
			// for _, subId := range goppConf.StartSubMissionList {
			// 	if finishSubMainMissionList[subId] != nil {
			// 		continue
			// 	}
			// 	nextList = append(nextList, subId)
			// 	subMainMissionList[subId] = &spb.MissionInfo{
			// 		MissionId: subId,
			// 		Progress:  0,
			// 		Status:    spb.MissionStatus_MISSION_DOING,
			// 	}
			// }
		}
	}
	// 通知状态
	g.MissionPlayerSyncScNotify(nextList, make([]uint32, 0), make([]uint32, 0)) // 发送通知
}

func (g *GamePlayer) IsReceiveMission(mission *gdconf.MainMission, mainMissionList, finishMainMissionList map[uint32]*spb.MissionInfo) bool {
	var isReceive = false
	if mission == nil || mainMissionList == nil || finishMainMissionList == nil || mission.TakeParam == nil {
		return false
	}
	if mainMissionList[mission.MainMissionID] != nil || finishMainMissionList[mission.MainMissionID] != nil { // 过滤已接取已完成的
		return false
	}
	for _, take := range mission.TakeParam {
		switch take.Type {
		case "Auto":
			isReceive = true
		case "MultiSequence":
			if finishMainMissionList[take.Value] != nil {
				isReceive = true
			}
		default:
			isReceive = false
		}
	}

	return isReceive
}