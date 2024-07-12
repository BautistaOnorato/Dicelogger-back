package router

import (
	"database/sql"

	firebase "firebase.google.com/go/v4"
	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/cmd/server/handler"
	"github.com/proyecto-dnd/backend/internal/armor"
	"github.com/proyecto-dnd/backend/internal/armorXCharacterData"
	"github.com/proyecto-dnd/backend/internal/attackEvent"
	"github.com/proyecto-dnd/backend/internal/background"
	backgroundXproficiency "github.com/proyecto-dnd/backend/internal/backgroundXProficiency"
	"github.com/proyecto-dnd/backend/internal/campaign"
	characterdata "github.com/proyecto-dnd/backend/internal/characterData"
	charactertrade "github.com/proyecto-dnd/backend/internal/characterTrade"
	characterXproficiency "github.com/proyecto-dnd/backend/internal/characterXProficiency"
	characterXspell "github.com/proyecto-dnd/backend/internal/characterXSpell"
	classXspell "github.com/proyecto-dnd/backend/internal/classXSpell"
	"github.com/proyecto-dnd/backend/internal/dice_event"
	"github.com/proyecto-dnd/backend/internal/report"
	tradeevent "github.com/proyecto-dnd/backend/internal/tradeEvent"
	"github.com/proyecto-dnd/backend/internal/ws"

	// "github.com/proyecto-dnd/backend/internal/ws"

	// characterXproficiency "github.com/proyecto-dnd/backend/internal/characterXProficiency"
	// characterXspell "github.com/proyecto-dnd/backend/internal/characterXSpell"
	"github.com/proyecto-dnd/backend/internal/character_feature"
	"github.com/proyecto-dnd/backend/internal/class"

	// classXspell "github.com/proyecto-dnd/backend/internal/classXSpell"
	characterXAttackEvent "github.com/proyecto-dnd/backend/internal/characterXAttackEvent"
	"github.com/proyecto-dnd/backend/internal/feature"
	"github.com/proyecto-dnd/backend/internal/friendship"
	"github.com/proyecto-dnd/backend/internal/item"
	itemxcharacterdata "github.com/proyecto-dnd/backend/internal/itemXCharacterData"
	"github.com/proyecto-dnd/backend/internal/proficiency"
	"github.com/proyecto-dnd/backend/internal/proficiencyXclass.go"
	"github.com/proyecto-dnd/backend/internal/race"
	raceXproficiency "github.com/proyecto-dnd/backend/internal/raceXProficiency"
	"github.com/proyecto-dnd/backend/internal/session"
	"github.com/proyecto-dnd/backend/internal/skill"
	skillxcharacterdata "github.com/proyecto-dnd/backend/internal/skillXCharacterData"
	"github.com/proyecto-dnd/backend/internal/spell"
	"github.com/proyecto-dnd/backend/internal/user"
	"github.com/proyecto-dnd/backend/internal/user_campaign"
	"github.com/proyecto-dnd/backend/internal/weapon"
	weaponxcharacterdata "github.com/proyecto-dnd/backend/internal/weaponXCharacterData"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	userFirebaseRepository user.RepositoryUsers
	userFirebaseService    user.ServiceUsers
	userFirebaseHandler    *handler.UserHandler

	campaignRepository campaign.CampaignRepository
	campaignService    campaign.CampaignService
	campaignHandler    *handler.CampaignHandler

	sessionRepository session.SessionRepository
	sessionService    session.SessionService
	sessionHandler    *handler.SessionHandler

	classRepository class.RepositoryCharacterClass
	classService    class.ClassService
	classHandler    *handler.ClassHandler

	proficiencyRepository            proficiency.RepositoryProficiency
	proficiencyService               proficiency.ProficiencyService
	proficiencyHandler               *handler.ProficiencyHandler
	proficiencyXClassRepository      proficiencyXclass.ProficiencyXClassRepository
	proficiencyXClassService         proficiencyXclass.ProficiencyXClassService
	proficiencyXClassHandler         *handler.ProficiencyXClassHandler
	backgroundXProficiencyRepository backgroundXproficiency.BackgroundXProficiencyRepository
	backgroundXProficiencyService    backgroundXproficiency.BackgroundXProficiencyService
	backgroundXProficiencyHandler    *handler.BackgroundXProficiencyHandler

	userCampaignRepository user_campaign.UserCampaignRepository
	userCampaignService    user_campaign.UserCampaignService
	userCampaignHandler    *handler.UserCampaignHandler
	friendshipRepository   friendship.FriendshipRepository
	friendshipService      friendship.FriendshipService
	friendshipHandler      *handler.FriendshipHandler

	featureRepository               feature.FeatureRepository
	featureService                  feature.FeatureService
	featureHandler                  *handler.FeatureHandler
	featureXCharacterDataRepository character_feature.CharacterFeatureRepository
	featureXCharacterDataService    character_feature.CharacterFeatureService
	featureXCharacterDataHandler    handler.CharacterFeature

	characterFeatureRepository character_feature.CharacterFeatureRepository
	characterFeatureService    character_feature.CharacterFeatureService
	characterFeatureHandler    *handler.CharacterFeature

	skillRepository               skill.RepositorySkill
	skillService                  skill.ServiceSkill
	skillHandler                  handler.SkillHandler
	skillXCharacterDataRepository skillxcharacterdata.RepositorySkillXCharacter
	skillXCharacterDataService    skillxcharacterdata.ServiceSkillXCharacter
	skillXCharacterDataHandler    handler.SkillXCharacterHandler

	spellRepository       spell.RepositorySpell
	spellService          spell.ServiceSpell
	spellHandler          *handler.SpellHandler
	classXSpellRepository classXspell.ServiceClassXSpell
	classXSpellService    classXspell.ServiceClassXSpell
	classXSpellHandler    *handler.ClassXSpellHandler

	raceRepository             race.RaceRepository
	raceService                race.RaceService
	raceHandler                *handler.RaceHandler
	raceXProficiencyRepository raceXproficiency.RaceXProficiencyRepository
	raceXProficiencyService    raceXproficiency.RaceXProficiencyService
	raceXProficiencyHandler    *handler.RaceXProficiencyHandler

	characterXSpellRepository characterXspell.RepositoryCharacterXSpell
	characterXSpellService    characterXspell.ServiceCharacterXSpell
	characterXSpellHandler    *handler.CharacterXSpellHandler

	attackEventRepository attackEvent.AttackEventRepository
	attackEventService    attackEvent.AttackEventService
	attackEventHandler    *handler.AttackEventHandler

	characterTradeRepository charactertrade.RepositoryCharacterTrade
	characterTradeService    charactertrade.ServiceCharacterTrade
	tradeEventRepository     tradeevent.RepositoryTradeEvent
	tradeEventService        tradeevent.ServiceTradeEvent
	tradeEventHandler        *handler.TradeEventHandler

	itemRepository               item.RepositoryItem
	itemService                  item.ServiceItem
	itemHandler                  *handler.ItemHandler
	itemXCharacterDataRepository itemxcharacterdata.RepositoryItemXCharacterData
	itemXCharacterDataService    itemxcharacterdata.ServiceItemXCharacterData
	itemXCharacterDataHandler    *handler.ItemXCharacterDataHandler

	weaponRepository               weapon.RepositoryWeapon
	weaponService                  weapon.ServiceWeapon
	weaponHandler                  *handler.WeaponHandler
	weaponXCharacterDataRepository weaponxcharacterdata.RepositoryWeaponXCharacterData
	weaponXCharacterDataService    weaponxcharacterdata.ServiceWeaponXCharacterData
	weaponXCharacterDataHandler    *handler.WeaponXCharacterDataHandler

	armorRepository               armor.ArmorRepository
	armorService                  armor.ArmorService
	armorHandler                  *handler.ArmorHandler
	armorXCharacterDataRepository armorXCharacterData.RepositoryArmorXCharacterData
	armorXCharacterDataService    armorXCharacterData.ServiceArmorXCharacterData
	armorXCharacterDataHandler    *handler.ArmorXCharacterDataHandler

	characterXProficiencyRepository characterXproficiency.CharacterXProficiencyRepository
	characterXProficiencyService    characterXproficiency.CharacterXProficiencyService
	characterXProficiencyHandler    *handler.CharacterXProficiencyHandler

	characterDataRepository characterdata.RepositoryCharacterData
	characterDataService    characterdata.ServiceCharacterData
	characterDataHandler    *handler.CharacterHandler

	characterXAttackEventRepository characterXAttackEvent.CharacterXAttackEventRepository
	characterXAttackEventService    characterXAttackEvent.CharacterXAttackEventService
	characterXAttackEventHandler    *handler.CharacterXAttackEventHandler

	diceEventRepository dice_event.DiceEventRepository
	diceEventService    dice_event.DiceEventService
	diceEventHandler    *handler.DiceEventHandler

	backgroundRepository background.BackgroundRepository
	backgroundService    background.BackgroundService
	backgroundHandler    *handler.BackgroundHandler

	reportGenerator *report.ReportGenerator
	reportHandler   *handler.ReportHandler
)

type Router interface {
	MapRoutes()
}

type router struct {
	engine      *gin.Engine
	routerGroup *gin.RouterGroup
	db          *sql.DB
	firebaseApp *firebase.App
	hub         *ws.Hub
}

func NewRouter(engine *gin.Engine, db *sql.DB, firebaseApp *firebase.App) Router {
	userFirebaseRepository = user.NewUserFirebaseRepository(firebaseApp, db)
	userFirebaseService = user.NewServiceUser(userFirebaseRepository)
	userFirebaseHandler = handler.NewUserHandler(&userFirebaseService)

	itemRepository = item.NewItemRepository(db)
	itemService = item.NewItemService(itemRepository)
	itemHandler = handler.NewItemHandler(&itemService)
	itemXCharacterDataRepository = itemxcharacterdata.NewItemXCharacterDataSqlRepository(db)
	itemXCharacterDataService = itemxcharacterdata.NewItemXCharacterDataService(itemXCharacterDataRepository, itemRepository)
	itemXCharacterDataHandler = handler.NewItemXCharacterDataHandler(&itemXCharacterDataService)

	weaponRepository = weapon.NewWeaponRepository(db)
	weaponService = weapon.NewWeaponService(weaponRepository)
	weaponHandler = handler.NewWeaponHandler(&weaponService)
	weaponXCharacterDataRepository = weaponxcharacterdata.NewWeaponXCharacterDataSqlRepository(db)
	weaponXCharacterDataService = weaponxcharacterdata.NewWeaponXCharacterDataService(weaponXCharacterDataRepository, weaponRepository)
	weaponXCharacterDataHandler = handler.NewWeaponXCharacterDataHandler(&weaponXCharacterDataService)

	armorRepository = armor.NewArmorRepository(db)
	armorService = armor.NewArmorService(armorRepository)
	armorHandler = handler.NewArmorHandler(&armorService)
	armorXCharacterDataRepository = armorXCharacterData.NewArmorXCharacterDataSqlRepository(db)
	armorXCharacterDataService = armorXCharacterData.NewServiceArmorXCharacterData(armorXCharacterDataRepository, armorService)
	armorXCharacterDataHandler = handler.NewArmorXCharacterDataHandler(&armorXCharacterDataService) // TO DO Check if armorXCharacterDataHandler works correctly, it was done fast to compile the rest

	classRepository = class.NewClassRepository(db)
	classService = class.NewClassService(classRepository)
	classHandler = handler.NewClassHandler(&classService)

	raceRepository = race.NewRaceRepository(db)
	raceService = race.NewRaceService(raceRepository)
	raceHandler = handler.NewRaceHandler(raceService)
	raceXProficiencyRepository = raceXproficiency.NewRaceXProficiencyRepository(db)
	raceXProficiencyService = raceXproficiency.NewRaceXProficiencyService(raceXProficiencyRepository)
	raceXProficiencyHandler = handler.NewRaceXProficiencyHandler(&raceXProficiencyService)

	proficiencyRepository = proficiency.NewProficiencyRepository(db)
	proficiencyService = proficiency.NewProficiencyService(proficiencyRepository)
	proficiencyHandler = handler.NewProficiencyHandler(&proficiencyService)
	proficiencyXClassRepository = proficiencyXclass.NewProficiencyXClassRepository(db)
	proficiencyXClassService = proficiencyXclass.NewProficiencyXClassService(proficiencyXClassRepository)
	proficiencyXClassHandler = handler.NewProficiencyXClassHandler(proficiencyXClassService)
	characterXProficiencyRepository = characterXproficiency.NewCharacterXProficiencyRepository(db)
	characterXProficiencyService = characterXproficiency.NewCharacterXProficiencyService(characterXProficiencyRepository)
	characterXProficiencyHandler = handler.NewCharacterXProficiencyHandler(characterXProficiencyService)
	userCampaignRepository = user_campaign.NewUserCampaignRepository(db)
	userCampaignService = user_campaign.NewUserCampaignService(userCampaignRepository, userFirebaseService)
	userCampaignHandler = handler.NewUserCampaignHandler(&userCampaignService)
	friendshipRepository = friendship.NewFriendshipRepository(db, userFirebaseRepository, firebaseApp)
	friendshipService = friendship.NewFriendshipService(friendshipRepository)
	friendshipHandler = handler.NewFriendshipHandler(&friendshipService, &userFirebaseService)

	skillRepository = skill.NewSkillRepository(db)
	skillService = skill.NewServiceSkill(skillRepository)
	skillHandler = *handler.NewSkillHandler(&skillService)
	skillXCharacterDataRepository = skillxcharacterdata.NewSkillxCharacterDataRepository(db)
	skillXCharacterDataService = skillxcharacterdata.NewSkillXCharacterService(skillXCharacterDataRepository)
	skillXCharacterDataHandler = *handler.NewSkillXCharacterHandler(&skillXCharacterDataService)

	featureRepository = feature.NewFeatureRepository(db)
	featureService = feature.NewFeatureService(featureRepository)
	featureHandler = handler.NewFeatureHandler(&featureService)
	featureXCharacterDataRepository = character_feature.NewCharacterFeatureRepository(db)
	featureXCharacterDataService = character_feature.NewCharacterFeatureService(featureXCharacterDataRepository)
	featureXCharacterDataHandler = *handler.NewCharacterFeatureHandler(&featureXCharacterDataService)

	spellRepository = spell.NewSpellRepository(db)
	spellService = spell.NewSpellService(spellRepository)
	spellHandler = handler.NewSpellHandler(&spellService)
	classXSpellRepository = classXspell.NewClassXSpellRepository(db)
	classXSpellService = classXspell.NewClassXSpelService(classXSpellRepository)
	classXSpellHandler = handler.NewClassXSpellHandler(&classXSpellService)
	characterXSpellRepository = characterXspell.NewCharacterXSpellRepository(db)
	characterXSpellService = characterXspell.NewCharacterXSpellService(characterXSpellRepository)
	characterXSpellHandler = handler.NewCharacterXSpellHandler(&characterXSpellService)

	sessionRepository = session.NewSessionRepository(db)
	sessionService = session.NewSessionService(sessionRepository)
	sessionHandler = handler.NewSessionHandler(&sessionService)

	backgroundXProficiencyRepository = backgroundXproficiency.NewBackgroundXProficiencyRepository(db)
	backgroundXProficiencyService = backgroundXproficiency.NewBackgroundXProficiencyService(backgroundXProficiencyRepository)
	backgroundXProficiencyHandler = handler.NewBackgroundXProficiencyHandler(backgroundXProficiencyService)
	backgroundRepository = background.NewBackgroundRepository(db)
	backgroundService = background.NewBackgroundService(backgroundRepository)
	backgroundHandler = handler.NewBackgroundHandler(backgroundService)

	characterFeatureRepository = character_feature.NewCharacterFeatureRepository(db)
	characterFeatureService = character_feature.NewCharacterFeatureService(characterFeatureRepository)
	characterFeatureHandler = handler.NewCharacterFeatureHandler(&characterFeatureService)

	characterTradeRepository = charactertrade.NewCharacterTradeMySqlRepository(db)
	characterTradeService = charactertrade.NewCharacterTradeService(characterTradeRepository)
	tradeEventRepository = tradeevent.NewTradeEventMySqlRepository(db)
	tradeEventService = tradeevent.NewTradeEventService(tradeEventRepository, characterTradeService, weaponXCharacterDataService, armorXCharacterDataService, itemXCharacterDataService)
	tradeEventHandler = handler.NewTradeEventHandler(&tradeEventService)

	attackEventRepository = attackEvent.NewAttackEventRepository(db)
	attackEventService = attackEvent.NewAttackEventService(attackEventRepository)
	attackEventHandler = handler.NewAttackEventHandler(&attackEventService)

	diceEventRepository = dice_event.NewDiceEventRepository(db)
	diceEventService = dice_event.NewDiceEventService(diceEventRepository)
	diceEventHandler = handler.NewDiceEventHandler(diceEventService)

	characterDataRepository = characterdata.NewCharacterDataRepository(db)
	characterDataService = characterdata.NewServiceCharacterData(characterDataRepository, itemXCharacterDataService, weaponXCharacterDataService, armorXCharacterDataService, skillService, skillXCharacterDataService, featureService, featureXCharacterDataService, spellService, characterXSpellService, proficiencyService, characterXProficiencyService, tradeEventService, attackEventService, diceEventService, userFirebaseService)
	characterDataHandler = handler.NewCharacterHandler(&characterDataService)

	campaignRepository = campaign.NewCampaignRepository(db)
	campaignService = campaign.NewCampaignService(campaignRepository, sessionService, userCampaignService, characterDataService, userFirebaseService)
	campaignHandler = handler.NewCampaignHandler(&campaignService, &userFirebaseService)

	characterXAttackEventRepository = characterXAttackEvent.NewCharacterXAttackEventRepository(db)
	characterXAttackEventService = characterXAttackEvent.NewCharacterXAttackEventService(characterXAttackEventRepository)
	characterXAttackEventHandler = handler.NewCharacterXAttackEventHandler(characterXAttackEventService)

	reportGenerator = report.NewReportGenerator(tradeEventService, attackEventService, diceEventService, characterDataService)
	reportHandler = handler.NewReportHandler(reportGenerator)

	hub := ws.NewHub(tradeEventService, attackEventService, diceEventService)
	go hub.Run()
	return &router{
		engine:      engine,
		db:          db,
		firebaseApp: firebaseApp,
		hub:         hub,
	}
}

func (r *router) MapRoutes() {
	r.setGroup()
	r.setSwaggerRoute()
	r.buildUserRoutes()
	r.buildAttackEventRoutes()
	r.buildCampaignRoutes()
	r.buildSessionRoutes()
	r.buildClassRoutes()
	r.buildProficiencyRoutes()
	r.buildProficiencyXClassRoutes()
	r.buildUserCampaignRoutes()
	r.buildFriendshipRoutes()
	r.buildSpellRoutes()
	r.buildClassXSpellRoutes()
	r.buildRaceXProficiencyRoutes()
	r.buildBackgroundXProficiencyRoutes()
	r.buildCharacterXSpellRoutes()
	r.buildFeatureRoutes()
	r.buildItemRoutes()
	r.buildItemXCharacterDataRoutes()
	r.buildWeaponRoutes()
	r.buildWeaponXCharacterDataRoutes()
	r.buildCharacterXProficiencyRoutes()
	r.buildSkillRoutes()
	r.buildRaceRoutes()
	r.buildCharacterDataRoutes()
	r.buildCharacterFeatureRoutes()
	r.buildArmorRoutes()
	r.buildArmorXCharacterDataRoutes()
	r.buildCharacterXAttackEventRoutes()
	r.buildTradeEventRoutes()
	r.buildDiceEventRoutes()
	r.buildSkillXCharacterDataRoutes()
	r.buildWebsocketRoutes()
	r.buildReportRoutes()
	r.buildBackgroundRoutes()
	// TODO Add other builders here	and write their functions

}

func (r *router) setGroup() {
	r.routerGroup = r.engine.Group("/api/v1")
}

func (r *router) setSwaggerRoute() {
	r.routerGroup.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func (r *router) buildUserRoutes() {
	userGroup := r.routerGroup.Group("/user")
	{
		// TODO Add Middlewares if needed
		userGroup.POST("/register", userFirebaseHandler.HandlerCreate())
		userGroup.POST("/login", userFirebaseHandler.HandlerLogin())
		userGroup.POST("/subscribe/:months", userFirebaseHandler.HandlerSubPremium())
		userGroup.GET("", userFirebaseHandler.HandlerGetAll())
		userGroup.GET("/:id", userFirebaseHandler.HandlerGetById())
		userGroup.GET("/checkSub", userFirebaseHandler.HandlerCheckSubscriptionExpDate())
		userGroup.GET("/jwt", userFirebaseHandler.HandlerGetJwtInfo())
		userGroup.POST("/sendEmailVerification", userFirebaseHandler.HandlerSendEmailVerification())
		userGroup.PUT("/:id", userFirebaseHandler.HandlerUpdate())
		userGroup.PATCH("/", userFirebaseHandler.HandlerPatch())
		userGroup.DELETE("/:id", userFirebaseHandler.HandlerDelete())
	}
}

func (r *router) buildAttackEventRoutes() {
	eventGroup := r.routerGroup.Group("/attackevent")
	{
		eventGroup.POST("", attackEventHandler.HandlerCreate())
		eventGroup.GET("", attackEventHandler.HandlerGetAll())
		eventGroup.GET("/:id", attackEventHandler.HandlerGetById())
		eventGroup.GET("/session/:id", attackEventHandler.HandlerGetBySessionId())
		eventGroup.GET("/protagonist/:id", attackEventHandler.HandlerGetByProtagonistId())
		eventGroup.GET("/affected/:id", attackEventHandler.HandlerGetByAffectedId())
		eventGroup.GET("/prot/:protagonistid/aff/:affectedid", attackEventHandler.HandlerGetByProtagonistIdAndAffectedId())
		eventGroup.PUT("/:id", attackEventHandler.HandlerUpdate())
		eventGroup.DELETE("/:id", attackEventHandler.HandlerDelete())
	}
}

func (r *router) buildCampaignRoutes() {
	campaignGroup := r.routerGroup.Group("/campaign")
	{
		campaignGroup.POST("", campaignHandler.HandlerCreate())
		campaignGroup.GET("", campaignHandler.HandlerGetAll())
		campaignGroup.GET("/:id", campaignHandler.HandlerGetById())
		campaignGroup.GET("/user", campaignHandler.HandlerGetByUserId())
		campaignGroup.PUT("/:id", campaignHandler.HandlerUpdate())
		campaignGroup.DELETE("/:id", campaignHandler.HandlerDelete())
	}
}

func (r *router) buildSessionRoutes() {
	sessionGroup := r.routerGroup.Group("/session")
	{
		sessionGroup.POST("", sessionHandler.HandlerCreate())
		sessionGroup.GET("", sessionHandler.HandlerGetAll())
		sessionGroup.GET("/:id", sessionHandler.HandlerGetById())
		sessionGroup.GET("/campaign/:id", sessionHandler.HandlerGetByCampaignId())
		sessionGroup.PUT("/:id", sessionHandler.HandlerUpdate())
		sessionGroup.DELETE("/:id", sessionHandler.HandlerDelete())
	}
}

func (r *router) buildClassRoutes() {
	classGroup := r.routerGroup.Group("/class")
	{
		classGroup.POST("", classHandler.HandlerCreate())
		classGroup.GET("", classHandler.HandlerGetAll())
		classGroup.GET("/:id", classHandler.HandlerGetById())
		classGroup.PUT("/:id", classHandler.HandlerUpdate())
		classGroup.DELETE("/:id", classHandler.HandlerDelete())
	}
}

func (r *router) buildProficiencyRoutes() {
	proficiencyGroup := r.routerGroup.Group("/proficiency")
	{
		proficiencyGroup.POST("", proficiencyHandler.HandlerCreate())
		proficiencyGroup.GET("", proficiencyHandler.HandlerGetAll())
		proficiencyGroup.GET("/:id", proficiencyHandler.HandlerGetById())
		proficiencyGroup.GET("/character/:characterId", proficiencyHandler.HandlerGetByCharacterId())
		proficiencyGroup.PUT("/:id", proficiencyHandler.HandlerUpdate())
		proficiencyGroup.DELETE("/:id", proficiencyHandler.HandlerDelete())
	}
}

func (r *router) buildBackgroundRoutes() {
	backgroundGroup := r.routerGroup.Group("/background")
	{
		backgroundGroup.GET("", backgroundHandler.HandlerGetAll())
		backgroundGroup.GET("/:id", backgroundHandler.HandlerGetById())
	}
}

func (r *router) buildProficiencyXClassRoutes() {
	proficiencyXClassGroup := r.routerGroup.Group("/proficiencyxclass")
	{
		proficiencyXClassGroup.POST("", proficiencyXClassHandler.HandlerCreate())
		proficiencyXClassGroup.DELETE("", proficiencyXClassHandler.HandlerDelete())
	}
}

func (r *router) buildUserCampaignRoutes() {
	userCampaignGroup := r.routerGroup.Group("/user_campaign")
	{
		userCampaignGroup.POST("", userCampaignHandler.HandlerCreate())
		userCampaignGroup.POST("/friends/:id", userCampaignHandler.HandlerAddFriendsToCampaign())
		userCampaignGroup.PUT("/addCharacter", userCampaignHandler.HandlerAddCharacterToCampaign())
		userCampaignGroup.GET("", userCampaignHandler.HandlerGetAll())
		userCampaignGroup.GET("/:id", userCampaignHandler.HandlerGetById())
		userCampaignGroup.GET("/campaign/:id", userCampaignHandler.HandlerGetByCampaignId())
		userCampaignGroup.GET("/user/:id", userCampaignHandler.HandlerGetByUserId())
		userCampaignGroup.DELETE("/:id", userCampaignHandler.HandlerDelete())

	}
}

func (r *router) buildFriendshipRoutes() {
	friendshipGroup := r.routerGroup.Group("/friendship")
	{
		friendshipGroup.POST("/:friend", friendshipHandler.HandlerCreate())
		friendshipGroup.GET("", friendshipHandler.HandlerGetAllFriends())
		friendshipGroup.GET("/search/:name", friendshipHandler.HandlerGetBySimilarName())
		friendshipGroup.GET("/find/:name", friendshipHandler.HandlerSearchFollowers())
		friendshipGroup.DELETE("/:friend", friendshipHandler.HandlerDelete())
	}
}

func (r *router) buildFeatureRoutes() {
	featureGroup := r.routerGroup.Group("/feature")
	{
		featureGroup.POST("", featureHandler.HandlerCreate())
		featureGroup.GET("", featureHandler.HandlerGetAll())
		featureGroup.GET("/character/:id", featureHandler.HandlerGetAllFeaturesByCharacterId())
		featureGroup.GET("/:id", featureHandler.HandlerGetById())
		featureGroup.PUT("/:id", featureHandler.HandlerUpdate())
		featureGroup.DELETE("/:id", featureHandler.HandlerDelete())
	}
}

func (r *router) buildSpellRoutes() {
	spellGroup := r.routerGroup.Group("/spell")
	{
		spellGroup.POST("", spellHandler.HandlerCreate())
		spellGroup.GET("", spellHandler.HandlergetAll())
		spellGroup.GET("/:id", spellHandler.HandlerGetById())
		spellGroup.GET("/character/:id", spellHandler.HandlerGetByCharacterId())
		spellGroup.GET("/class/:id", spellHandler.HandlerGetByClassId())
		spellGroup.PUT("/:id", spellHandler.HandlerUpdate())
		spellGroup.DELETE("/:id", spellHandler.HandlerDelete())
	}
}

func (r *router) buildClassXSpellRoutes() {
	classXSpellGroup := r.routerGroup.Group("/classxspell")
	{
		classXSpellGroup.POST("", classXSpellHandler.HandlerCreate())
		classXSpellGroup.DELETE("", classXSpellHandler.HandlerDelete())
	}
}

func (r *router) buildRaceXProficiencyRoutes() {
	raceXProficiencyGroup := r.routerGroup.Group("/raceXproficiency")
	{
		raceXProficiencyGroup.POST("", raceXProficiencyHandler.HandlerCreate())
		raceXProficiencyGroup.DELETE("", raceXProficiencyHandler.HandlerDelete())
	}
}

func (r *router) buildBackgroundXProficiencyRoutes() {
	backgroundXProficiencyGroup := r.routerGroup.Group("/backgroundXproficiency")
	{
		backgroundXProficiencyGroup.POST("", backgroundXProficiencyHandler.HandlerCreate())
		backgroundXProficiencyGroup.DELETE("", backgroundXProficiencyHandler.HandlerDelete())
	}
}

func (r *router) buildCharacterXSpellRoutes() {
	characterXSpellGroup := r.routerGroup.Group("/characterXspell")
	{
		characterXSpellGroup.POST("", characterXSpellHandler.HandlerCreate())
		characterXSpellGroup.DELETE("/:id", characterXSpellHandler.HandlerDelete())
		characterXSpellGroup.DELETE("/delete", characterXSpellHandler.HandlerDeleteParams())
	}
}

func (r *router) buildCharacterFeatureRoutes() {
	characterFeatureGroup := r.routerGroup.Group("/character_feature")
	{
		characterFeatureGroup.POST("", characterFeatureHandler.HandlerCreate())
		characterFeatureGroup.GET("", characterFeatureHandler.HandlerGetAll())
		characterFeatureGroup.GET("/feature/:id", characterFeatureHandler.HandlerGetByFeatureId())
		characterFeatureGroup.GET("/character/:id", characterFeatureHandler.HandlerGetByCharacterId())
		characterFeatureGroup.DELETE("/", characterFeatureHandler.HandlerDeleteQueries())
	}
}

func (r *router) buildItemRoutes() {
	itemGroup := r.routerGroup.Group("/item")
	{
		itemGroup.POST("", itemHandler.HandlerCreate())
		itemGroup.DELETE("/:id", itemHandler.HandlerDelete())
		itemGroup.GET("", itemHandler.HandlerGetAll())
		itemGroup.GET("/generic", itemHandler.HandlerGetAllGeneric())
		itemGroup.GET("/:id", itemHandler.HandlerGetById())
		itemGroup.GET("/campaign/:id", itemHandler.HandlerGetByCampaignId())
		itemGroup.PUT("/:id", itemHandler.HandlerUpdate())
	}
}

func (r *router) buildItemXCharacterDataRoutes() {
	itemXCharacterDataGroup := r.routerGroup.Group("/item_character")
	{
		itemXCharacterDataGroup.POST("", itemXCharacterDataHandler.HandlerCreate())
		itemXCharacterDataGroup.DELETE("/:id", itemXCharacterDataHandler.HandlerDelete())
		itemXCharacterDataGroup.DELETE("/character/:id", itemXCharacterDataHandler.HandlerDeleteByCharacterId())
		itemXCharacterDataGroup.GET("", itemXCharacterDataHandler.HandlerGetAll())
		itemXCharacterDataGroup.GET("/:id", itemXCharacterDataHandler.HandlerGetById())
		itemXCharacterDataGroup.GET("/character/:id", itemXCharacterDataHandler.HandlerGetByCharacterDataId())
		itemXCharacterDataGroup.PUT("/:id", itemXCharacterDataHandler.HandlerUpdate())
	}
}

func (r *router) buildWeaponRoutes() {
	weaponGroup := r.routerGroup.Group("/weapon")
	{
		weaponGroup.POST("", weaponHandler.HandlerCreate())
		weaponGroup.GET("/generic", weaponHandler.HandlerGetAllGeneric())
		weaponGroup.DELETE("/:id", weaponHandler.HandlerDelete())
		weaponGroup.GET("", weaponHandler.HandlerGetAll())
		weaponGroup.GET("/:id", weaponHandler.HandlerGetById())
		weaponGroup.GET("/campaign/:id", weaponHandler.HandlerGetByCampaignId())
		weaponGroup.PUT("/:id", weaponHandler.HandlerUpdate())
	}
}

func (r *router) buildWeaponXCharacterDataRoutes() {
	weaponXCharacterDataGroup := r.routerGroup.Group("/weapon_character")
	{
		weaponXCharacterDataGroup.POST("", weaponXCharacterDataHandler.HandlerCreate())
		weaponXCharacterDataGroup.DELETE("/:id", weaponXCharacterDataHandler.HandlerDelete())
		weaponXCharacterDataGroup.DELETE("/character/:id", weaponXCharacterDataHandler.HandlerDeleteByCharacterDataId())
		weaponXCharacterDataGroup.GET("", weaponXCharacterDataHandler.HandlerGetAll())
		weaponXCharacterDataGroup.GET("/:id", weaponXCharacterDataHandler.HandlerGetById())
		weaponXCharacterDataGroup.GET("/character/:id", weaponXCharacterDataHandler.HandlerGetByCharacterDataId())
		weaponXCharacterDataGroup.PUT("/:id", weaponXCharacterDataHandler.HandlerUpdate())
	}
}

func (r *router) buildCharacterXProficiencyRoutes() {
	characterXProficiencyGroup := r.routerGroup.Group("/character_proficiency")
	{
		characterXProficiencyGroup.POST("", characterXProficiencyHandler.HandlerCreate())
		characterXProficiencyGroup.DELETE("/:id", characterXProficiencyHandler.HandlerDelete())
	}
}

func (r *router) buildSkillRoutes() {
	skillGroup := r.routerGroup.Group("/skill")
	{
		skillGroup.POST("", skillHandler.HandlerCreate())
		skillGroup.GET("", skillHandler.HandlerGetAll())
		skillGroup.GET("/:id", skillHandler.HandlerGetById())
		skillGroup.GET("/character/:characterId", skillHandler.HandlerGetByCharacterId())
		skillGroup.PUT("/:id", skillHandler.HandlerUpdate())
		skillGroup.DELETE("/:id", skillHandler.HandlerDelete())
	}
}

func (r *router) buildRaceRoutes() {
	raceGroup := r.routerGroup.Group("/race")
	{
		raceGroup.POST("", raceHandler.HandlerCreate())
		raceGroup.GET("", raceHandler.HandlerGetAll())
		raceGroup.GET("/:id", raceHandler.HandlerGetById())
		raceGroup.PUT("/:id", raceHandler.HandlerUpdate())
		raceGroup.DELETE("/:id", raceHandler.HandlerDelete())
	}
}

func (r *router) buildCharacterDataRoutes() {
	characterDataGroup := r.routerGroup.Group("/character")
	{
		characterDataGroup.POST("", characterDataHandler.HandlerCreate())
		characterDataGroup.GET("", characterDataHandler.HandlerGetAll())
		characterDataGroup.GET("/filter", characterDataHandler.HandlerGetByCampaignIdAndUserId())
		characterDataGroup.GET("/:id", characterDataHandler.HandlerGetById())
		characterDataGroup.GET("/event/:eventid", characterDataHandler.HandlerGetByAttackEventId())
		characterDataGroup.GET("/generic", characterDataHandler.HandlerGetGenerics())
		characterDataGroup.GET("/user", characterDataHandler.HandlerGetByUser())
		characterDataGroup.PUT("/:id", characterDataHandler.HandlerUpdate())
		characterDataGroup.DELETE("/:id", characterDataHandler.HandlerDelete())
	}
}

func (r *router) buildArmorRoutes() {
	armorGroup := r.routerGroup.Group("/armor")
	{
		armorGroup.POST("", armorHandler.HandlerCreate())
		armorGroup.GET("", armorHandler.HandlerGetAll())
		armorGroup.GET("/:id", armorHandler.HandlerGetById())
		armorGroup.PUT("/:id", armorHandler.HandlerUpdate())
		armorGroup.DELETE("/:id", armorHandler.HandlerDelete())
	}
}

func (r *router) buildArmorXCharacterDataRoutes() {
	armorXCharacterDataGroup := r.routerGroup.Group("/armor_character")
	{
		armorXCharacterDataGroup.POST("", armorXCharacterDataHandler.HandlerCreate())
		armorXCharacterDataGroup.DELETE("/:id", armorXCharacterDataHandler.HandlerDelete())
		armorXCharacterDataGroup.DELETE("/character/:id", armorXCharacterDataHandler.HandlerDeleteByCharacterId())
		armorXCharacterDataGroup.GET("", armorXCharacterDataHandler.HandlerGetAll())
		armorXCharacterDataGroup.GET("/:id", armorXCharacterDataHandler.HandlerGetById())
		armorXCharacterDataGroup.GET("/character/:id", armorXCharacterDataHandler.HandlerGetByCharacterDataId())
		armorXCharacterDataGroup.PUT("/:id", armorXCharacterDataHandler.HandlerUpdate())
	}
}

func (r *router) buildCharacterXAttackEventRoutes() {
	characterXAttackEventGroup := r.routerGroup.Group("/characterXattackevent")
	{
		characterXAttackEventGroup.POST("", characterXAttackEventHandler.HandlerCreate())
		characterXAttackEventGroup.GET("", characterXAttackEventHandler.HandlerGetAll())
		characterXAttackEventGroup.GET("/:id", characterXAttackEventHandler.HandlerGetById())
		characterXAttackEventGroup.GET("/character/:id", characterXAttackEventHandler.HandlerGetByCharacterId())
		characterXAttackEventGroup.GET("/attackevent/:id", characterXAttackEventHandler.HandlerGetByEventId())
		characterXAttackEventGroup.DELETE("/:id", characterXAttackEventHandler.HandlerDelete())
	}
}

func (r *router) buildDiceEventRoutes() {
	diceEventGroup := r.routerGroup.Group("/diceevent")
	{
		diceEventGroup.POST("", diceEventHandler.HandlerCreate())
		diceEventGroup.GET("", diceEventHandler.HandlerGetAll())
		diceEventGroup.GET("/:id", diceEventHandler.HandlerGetById())
		diceEventGroup.PUT("/:id", diceEventHandler.HandlerUpdate())
		diceEventGroup.DELETE("/:id", diceEventHandler.HandlerDelete())
	}
}

func (r *router) buildSkillXCharacterDataRoutes() {
	skillXCharacterDataGroup := r.routerGroup.Group("/skill_character")
	{
		skillXCharacterDataGroup.POST("", skillXCharacterDataHandler.HandlerCreate())
		skillXCharacterDataGroup.DELETE("/:id", skillXCharacterDataHandler.HandlerDelete())

	}
}

func (r *router) buildTradeEventRoutes() {
	tradeEventGroup := r.routerGroup.Group("/tradeevent")
	{
		tradeEventGroup.POST("", tradeEventHandler.HandlerCreate())
		tradeEventGroup.GET("/session/:id", tradeEventHandler.HandlerGetBySessionId())
		tradeEventGroup.GET("/sender/:id", tradeEventHandler.HandlerGetBySender())
		tradeEventGroup.GET("/receiver/:id", tradeEventHandler.HandlerGetByReceiver())
		tradeEventGroup.GET("/character/:id", tradeEventHandler.HandlerGetBySenderOrReceiver())
		tradeEventGroup.DELETE("/:id", tradeEventHandler.HandlerDelete())
	}
}

func (r *router) buildWebsocketRoutes() {
	wsGroup := r.routerGroup.Group("/ws")
	{
		wsGroup.GET("/:session_id", r.hub.ServeWs)
	}
}

func (r *router) buildReportRoutes() {
	reportGroup := r.routerGroup.Group("/report")
	{
		reportGroup.GET("/session/:id", reportHandler.HandlerGetSessionReport())
		reportGroup.GET("/character/campaign/:id", reportHandler.HandlerGetCharacterCampaignReport())
	}
}
