package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type SummonUnitDataInfo struct {
	SummonUnitDataMap     map[uint32]*SummonUnitData
	SummonUnitDataJsonMap map[uint32]*SummonUnitDataJson
}

type SummonUnitData struct {
	ID                      uint32 `json:"ID"`
	JsonPath                string `json:"JSONPath"`
	IsClient                bool   `json:"IsClient"`
	IsTeamSummon            bool   `json:"IsTeamSummon"`
	DestroyOnEnterBattle    bool   `json:"DestroyOnEnterBattle"`
	RemoveMazeBuffOnDestroy bool   `json:"RemoveMazeBuffOnDestroy"`
	MaxSummonCount          uint32 `json:"MaxSummonCount"`
	UniqueGroup             string `json:"UniqueGroup"`
}

type SummonUnitDataJson struct {
	GroupConfigName      string                 `json:"GroupConfigName"`
	ConfigEntityPath     string                 `json:"ConfigEntityPath"`
	TickLodTemplateName  string                 `json:"TickLodTemplateName"`
	ResidentEffects      []any                  `json:"ResidentEffects"`
	ShoesType            string                 `json:"ShoesType"`
	ShowShadow           bool                   `json:"ShowShadow"`
	Duration             Duration               `json:"Duration"`
	OnCreate             []any                  `json:"OnCreate"`
	OnDestroy            []OnDestroy            `json:"OnDestroy"`
	OnHide               []OnHide               `json:"OnHide"`
	OnShow               []OnShow               `json:"OnShow"`
	OnGroundInvalid      []any                  `json:"OnGroundInvalid"`
	OnSummonerGroundMove []OnSummonerGroundMove `json:"OnSummonerGroundMove"`
	AnimConfig           AnimConfig             `json:"AnimConfig"`
	MoveConfig           MoveConfig             `json:"MoveConfig"`
	AIConfig             AIConfig               `json:"AIConfig"`
	SkillConfig          SkillConfig            `json:"SkillConfig"`
}
type Duration struct {
	FixedValue FixedValue `json:"fixedValue"`
}
type InstigatorType struct {
	Type  string `json:"$type"`
	Alias string `json:"Alias"`
}
type OnDestroy struct {
	Type               string         `json:"$type"`
	TargetType         TargetType     `json:"TargetType,omitempty"`
	OnNameBoard        bool           `json:"OnNameBoard,omitempty"`
	TaskEnabled        bool           `json:"TaskEnabled"`
	IsClientOnly       bool           `json:"IsClientOnly"`
	EffectPath         string         `json:"EffectPath,omitempty"`
	FadeOutRegionStart int32          `json:"FadeOutRegionStart,omitempty"`
	FadeOutRegionEnd   int32          `json:"FadeOutRegionEnd,omitempty"`
	StimulusName       string         `json:"StimulusName,omitempty"`
	InstigatorType     InstigatorType `json:"InstigatorType,omitempty"`
}
type HangUpKey struct {
	Value string `json:"Value"`
}
type Scale struct {
	X            float32    `json:"x"`
	Y            float32    `json:"y"`
	Z            float32    `json:"z"`
	Normalized   Normalized `json:"normalized"`
	Magnitude    float64    `json:"magnitude"`
	SqrMagnitude uint32     `json:"sqrMagnitude"`
}
type OnHide struct {
	Type                string     `json:"$type"`
	TargetType          TargetType `json:"TargetType"`
	HangUpKey           HangUpKey  `json:"HangUpKey,omitempty"`
	AITickImmediately   bool       `json:"AITickImmediately,omitempty"`
	TaskEnabled         bool       `json:"TaskEnabled"`
	IsClientOnly        bool       `json:"IsClientOnly"`
	EffectPath          string     `json:"EffectPath,omitempty"`
	AttachPoint         string     `json:"AttachPoint,omitempty"`
	Scale               Scale      `json:"Scale,omitempty"`
	ParamEntityUsage    string     `json:"ParamEntityUsage,omitempty"`
	TowardMaxPitchAngle uint32     `json:"TowardMaxPitchAngle,omitempty"`
	MaxMutexCount       uint32     `json:"MaxMutexCount,omitempty"`
	FadeOutRegionStart  int32      `json:"FadeOutRegionStart,omitempty"`
	FadeOutRegionEnd    int32      `json:"FadeOutRegionEnd,omitempty"`
}
type PosTargetType struct {
	Type  string `json:"$type"`
	Alias string `json:"Alias"`
}
type OnShow struct {
	Type                string        `json:"$type"`
	TargetType          TargetType    `json:"TargetType"`
	PosTargetType       PosTargetType `json:"PosTargetType,omitempty"`
	TaskEnabled         bool          `json:"TaskEnabled"`
	IsClientOnly        bool          `json:"IsClientOnly"`
	EffectPath          string        `json:"EffectPath,omitempty"`
	AttachPoint         string        `json:"AttachPoint,omitempty"`
	Scale               Scale         `json:"Scale,omitempty"`
	ParamEntityUsage    string        `json:"ParamEntityUsage,omitempty"`
	TowardMaxPitchAngle uint32        `json:"TowardMaxPitchAngle,omitempty"`
	MaxMutexCount       uint32        `json:"MaxMutexCount,omitempty"`
	HangUpKey           HangUpKey     `json:"HangUpKey,omitempty"`
}
type TargetAlias struct {
	Type  string `json:"$type"`
	Alias string `json:"Alias"`
}
type OnSummonerGroundMove struct {
	Type           string           `json:"$type"`
	Predicate      Predicate        `json:"Predicate"`
	FailedTaskList []FailedTaskList `json:"FailedTaskList"`
	TaskEnabled    bool             `json:"TaskEnabled"`
}
type FailedTaskList struct {
	Type                         string                       `json:"$type"`
	TargetType                   TargetType                   `json:"TargetType"`
	AnimLogicState               string                       `json:"AnimLogicState,omitempty"`
	AnimStateName                string                       `json:"AnimStateName,omitempty"`
	TransitionDuration           float64                      `json:"TransitionDuration,omitempty"`
	StopWhenHitOthers            bool                         `json:"StopWhenHitOthers,omitempty"`
	TaskEnabled                  bool                         `json:"TaskEnabled"`
	IsClientOnly                 bool                         `json:"IsClientOnly"`
	NormalizedTimeEnd            any                          `json:"NormalizedTimeEnd,omitempty"`
	Active                       bool                         `json:"Active,omitempty"`
	Duration                     any                          `json:"Duration,omitempty"`
	BlurX                        float64                      `json:"BlurX,omitempty"`
	BlurY                        float64                      `json:"BlurY,omitempty"`
	BlurRadius                   float64                      `json:"BlurRadius,omitempty"`
	Iteration                    int                          `json:"Iteration,omitempty"`
	BlurStart                    float64                      `json:"BlurStart,omitempty"`
	BlurFeather                  float64                      `json:"BlurFeather,omitempty"`
	Projectile                   Projectile                   `json:"Projectile,omitempty"`
	OnProjectileHit              []OnProjectileHit            `json:"OnProjectileHit,omitempty"`
	OnProjectileLifetimeFinish   []OnProjectileLifetimeFinish `json:"OnProjectileLifetimeFinish,omitempty"`
	PosTargetType                PosTargetType                `json:"PosTargetType,omitempty"`
	EffectPath                   string                       `json:"EffectPath,omitempty"`
	Scale                        Scale                        `json:"Scale,omitempty"`
	ParamEntityUsage             string                       `json:"ParamEntityUsage,omitempty"`
	TowardMaxPitchAngle          uint32                       `json:"TowardMaxPitchAngle,omitempty"`
	MaxMutexCount                uint32                       `json:"MaxMutexCount,omitempty"`
	TargetAlias                  TargetAlias                  `json:"TargetAlias,omitempty"`
	ForceStart                   bool                         `json:"ForceStart,omitempty"`
	NormalizedTransitionDuration float64                      `json:"NormalizedTransitionDuration,omitempty"`
}
type OnProjectileLifetimeFinish struct {
	Type                  string               `json:"$type"`
	AttackTargetType      AttackTargetType     `json:"AttackTargetType"`
	AttackRootTargetType  AttackRootTargetType `json:"AttackRootTargetType"`
	TriggerBattle         bool                 `json:"TriggerBattle"`
	TriggerBattleDelay    float64              `json:"TriggerBattleDelay"`
	AttackDetectConfig    AttackDetectConfig   `json:"AttackDetectConfig"`
	HitConfig             HitConfig            `json:"HitConfig"`
	IncludeProps          bool                 `json:"IncludeProps"`
	AttackDetectCollision bool                 `json:"AttackDetectCollision"`
	TaskEnabled           bool                 `json:"TaskEnabled"`
	IsClientOnly          bool                 `json:"IsClientOnly"`
}
type OnProjectileHit struct {
	Type                  string               `json:"$type"`
	AttackTargetType      AttackTargetType     `json:"AttackTargetType"`
	AttackRootTargetType  AttackRootTargetType `json:"AttackRootTargetType"`
	TriggerBattle         bool                 `json:"TriggerBattle"`
	TriggerBattleDelay    float64              `json:"TriggerBattleDelay"`
	AttackDetectConfig    AttackDetectConfig   `json:"AttackDetectConfig"`
	HitConfig             HitConfig            `json:"HitConfig"`
	IncludeProps          bool                 `json:"IncludeProps"`
	AttackDetectCollision bool                 `json:"AttackDetectCollision"`
	TaskEnabled           bool                 `json:"TaskEnabled"`
	IsClientOnly          bool                 `json:"IsClientOnly"`
}
type HitConfig struct {
	HitAnimation string `json:"HitAnimation"`
}
type Normalized struct {
	X            float64     `json:"x"`
	Y            float64     `json:"y"`
	Z            float64     `json:"z"`
	Normalized   *Normalized `json:"normalized"`
	Magnitude    float64     `json:"magnitude"`
	SqrMagnitude float64     `json:"sqrMagnitude"`
}
type Predicate struct {
	Type         string     `json:"$type"`
	TargetType   TargetType `json:"TargetType"`
	AnimZone     string     `json:"AnimZone"`
	TaskEnabled  bool       `json:"TaskEnabled"`
	IsClientOnly bool       `json:"IsClientOnly"`
}
type AttachOffset struct {
	Y            float64    `json:"y"`
	Normalized   Normalized `json:"normalized"`
	Magnitude    float64    `json:"magnitude"`
	SqrMagnitude float64    `json:"sqrMagnitude"`
}
type LinearPitchAngle struct {
}
type Projectile struct {
	ColliderTemplate               string           `json:"ColliderTemplate"`
	Behavior                       string           `json:"Behavior"`
	FlySpeed                       int              `json:"FlySpeed"`
	EnableRayCast                  bool             `json:"EnableRayCast"`
	FlyTime                        int              `json:"FlyTime"`
	Gravity                        float64          `json:"Gravity"`
	MaxLifeTime                    float64          `json:"MaxLifeTime"`
	AttachPoint                    string           `json:"AttachPoint"`
	AttachOffset                   AttachOffset     `json:"AttachOffset"`
	TargetAttachPoint              string           `json:"TargetAttachPoint"`
	FlyEffect                      string           `json:"FlyEffect"`
	HitEffect                      string           `json:"HitEffect"`
	FlyEffectFadeOut               bool             `json:"FlyEffectFadeOut"`
	LinearPitchAngle               LinearPitchAngle `json:"LinearPitchAngle"`
	BoomerangEccentricity          float64          `json:"BoomerangEccentricity"`
	TriggerHitCallback             bool             `json:"TriggerHitCallback"`
	StaticProjectileCanPassAirWall bool             `json:"StaticProjectileCanPassAirWall"`
	ParabolaUseWorldSpaceUp        bool             `json:"ParabolaUseWorldSpaceUp"`
}
type AttackTargetType struct {
	Type  string `json:"$type"`
	Alias string `json:"Alias"`
}
type AttackRootTargetType struct {
	Type  string `json:"$type"`
	Alias string `json:"Alias"`
}
type Offset struct {
	Y            float64 `json:"y"`
	Magnitude    float64 `json:"magnitude"`
	SqrMagnitude float64 `json:"sqrMagnitude"`
}
type AttackDetectConfig struct {
	Type      string  `json:"$type"`
	MaxRadius float64 `json:"MaxRadius"`
	FanAngle  int     `json:"FanAngle"`
	Hight     int     `json:"Hight"`
	Offset    Offset  `json:"Offset"`
}
type FootScale struct {
	X            uint32     `json:"x"`
	Y            uint32     `json:"y"`
	Z            uint32     `json:"z"`
	Normalized   Normalized `json:"normalized"`
	Magnitude    float64    `json:"magnitude"`
	SqrMagnitude uint32     `json:"sqrMagnitude"`
}
type AnimConfig struct {
	FootScale                             FootScale `json:"FootScale"`
	WalkSpeedRatio                        uint32    `json:"WalkSpeedRatio"`
	RunSpeedRatio                         uint32    `json:"RunSpeedRatio"`
	RootMotionScale                       uint32    `json:"RootMotionScale"`
	MaxWalkAnimSpeedRatio                 float32   `json:"MaxWalkAnimSpeedRatio"`
	MaxRunAnimSpeedRatio                  float32   `json:"MaxRunAnimSpeedRatio"`
	ReferenceWalkSpeed                    float64   `json:"ReferenceWalkSpeed"`
	ReferenceRunSpeed                     float32   `json:"ReferenceRunSpeed"`
	ReferenceFastRunSpeed                 int32     `json:"ReferenceFastRunSpeed"`
	EnableIdleShow                        bool      `json:"EnableIdleShow"`
	AnimEventConfigList                   []string  `json:"AnimEventConfigList"`
	CommonAnimZoneConfigPath              string    `json:"CommonAnimZoneConfigPath"`
	DisableAnimEventLayers                []string  `json:"DisableAnimEventLayers"`
	SummonerAnimEventConfigList           []string  `json:"SummonerAnimEventConfigList"`
	OverrideMovementStepToGroundCurveName string    `json:"OverrideMovementStepToGroundCurveName"`
}
type AdvAIControllerConfig struct {
	HasWalkStop                  bool   `json:"HasWalkStop"`
	HasRunStop                   bool   `json:"HasRunStop"`
	TriggerTurnThresholdOnFacing uint32 `json:"TriggerTurnThresholdOnFacing"`
	TriggerTurnThresholdOnMove   uint32 `json:"TriggerTurnThresholdOnMove"`
	WalkStopTurnThreshold        uint32 `json:"WalkStopTurnThreshold"`
	RunStopTurnThreshold         uint32 `json:"RunStopTurnThreshold"`
}
type MoveConfig struct {
	EnablePlatformMove              bool                  `json:"EnablePlatformMove"`
	NavTurnSpeed                    uint32                `json:"NavTurnSpeed"`
	NavTurnBackTurnSpeed            uint32                `json:"NavTurnBackTurnSpeed"`
	NavMovingSteerTurnSpeed         uint32                `json:"NavMovingSteerTurnSpeed"`
	NavMovingSteerTurnBackTurnSpeed uint32                `json:"NavMovingSteerTurnBackTurnSpeed"`
	AdvAIControllerConfig           AdvAIControllerConfig `json:"AdvAIControllerConfig"`
}
type BeforeFollowTeleport struct {
	Type                string     `json:"$type"`
	TargetType          TargetType `json:"TargetType"`
	EffectPath          string     `json:"EffectPath"`
	AttachPoint         string     `json:"AttachPoint"`
	Scale               Scale      `json:"Scale"`
	ParamEntityUsage    string     `json:"ParamEntityUsage"`
	TowardMaxPitchAngle uint32     `json:"TowardMaxPitchAngle"`
	MaxMutexCount       uint32     `json:"MaxMutexCount"`
	TaskEnabled         bool       `json:"TaskEnabled"`
	IsClientOnly        bool       `json:"IsClientOnly"`
}
type AfterFollowTeleport struct {
	Type                string     `json:"$type"`
	TargetType          TargetType `json:"TargetType"`
	EffectPath          string     `json:"EffectPath"`
	AttachPoint         string     `json:"AttachPoint"`
	Scale               Scale      `json:"Scale"`
	ParamEntityUsage    string     `json:"ParamEntityUsage"`
	TowardMaxPitchAngle uint32     `json:"TowardMaxPitchAngle"`
	MaxMutexCount       uint32     `json:"MaxMutexCount"`
	TaskEnabled         bool       `json:"TaskEnabled"`
	IsClientOnly        bool       `json:"IsClientOnly"`
}
type AdvAIConfig struct {
	AttackRange            uint32                 `json:"AttackRange"`
	AttackAngle            uint32                 `json:"AttackAngle"`
	AttackCD               uint32                 `json:"AttackCD"`
	ChaseTolerance         uint32                 `json:"ChaseTolerance"`
	ChaseAngleTolerance    uint32                 `json:"ChaseAngleTolerance"`
	MaxChaseDistance       uint32                 `json:"MaxChaseDistance"`
	AlertnessIncreaseRatio uint32                 `json:"AlertnessIncreaseRatio"`
	LOSLoseTargetDelay     float64                `json:"LOSLoseTargetDelay"`
	MinWalkScale           float64                `json:"MinWalkScale"`
	MaxWalkScale           float32                `json:"MaxWalkScale"`
	MinRunScale            float64                `json:"MinRunScale"`
	MaxRunScale            uint32                 `json:"MaxRunScale"`
	MinFastRunScale        float64                `json:"MinFastRunScale"`
	MaxFastRunScale        float64                `json:"MaxFastRunScale"`
	BeforeFollowTeleport   []BeforeFollowTeleport `json:"BeforeFollowTeleport"`
	AfterFollowTeleport    []AfterFollowTeleport  `json:"AfterFollowTeleport"`
}
type AIConfig struct {
	AIFile           string      `json:"AIFile"`
	AdvAIConfig      AdvAIConfig `json:"AdvAIConfig"`
	AIMemoryStrategy string      `json:"AIMemoryStrategy"`
}
type SkillList struct {
	Name                  string     `json:"Name"`
	CanCastWithoutTarget  bool       `json:"CanCastWithoutTarget"`
	AdventureSkillType    string     `json:"AdventureSkillType"`
	TargetInfo            TargetInfo `json:"TargetInfo"`
	EntryAbility          string     `json:"EntryAbility"`
	CoolDown              uint32     `json:"CoolDown"`
	CustomSkillAnimStates []string   `json:"CustomSkillAnimStates"`
}
type SkillConfig struct {
	SkillList        []SkillList `json:"SkillList"`
	SkillAbilityList []any       `json:"SkillAbilityList"`
	AbilityList      []string    `json:"AbilityList"`
}

func (g *GameDataConfig) loadSummonUnitData() {
	g.SummonUnitDataInfo = &SummonUnitDataInfo{
		SummonUnitDataMap:     make(map[uint32]*SummonUnitData),
		SummonUnitDataJsonMap: make(map[uint32]*SummonUnitDataJson),
	}
	summonUnitDataList := make([]*SummonUnitData, 0)
	playerElementsFilePath := g.excelPrefix + "SummonUnitData.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &summonUnitDataList)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	for _, v := range summonUnitDataList {
		jsonData := new(SummonUnitDataJson)
		confElementsFile, err := os.ReadFile(g.pathPrefix + "/" + v.JsonPath)
		if err != nil {
			logger.Error("open file error: %v", err)
			continue
		}
		err = hjson.Unmarshal(confElementsFile, &jsonData)
		if err != nil {
			logger.Error("parse file error: %v", err)
			continue
		}
		g.SummonUnitDataInfo.SummonUnitDataJsonMap[v.ID] = jsonData
		g.SummonUnitDataInfo.SummonUnitDataMap[v.ID] = v
	}

	logger.Info("load %v SummonUnitData", len(g.SummonUnitDataInfo.SummonUnitDataMap))
	logger.Info("load %v SummonUnitDataJson", len(g.SummonUnitDataInfo.SummonUnitDataJsonMap))
}

func GetSummonUnitData(summonId uint32) *SummonUnitData {
	return CONF.SummonUnitDataInfo.SummonUnitDataMap[summonId]
}
