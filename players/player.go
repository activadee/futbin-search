package players

type FilteredPlayer struct {
	ID         int `json:"ID"`
	PlayerID   int `json:"playerid"`
	ResourceID int `json:"resource_id"`
	Image      int `json:"playerimage"`

	Name       string `json:"playername"`
	CommonName string `json:"common_name"`

	MainPosition string `json:"position"`
	Positions    string `json:"pos_all"`
	Accelerate   string `json:"accelerate"`

	Rating   int `json:"rating"`
	RareType int `json:"raretype"`

	Pace        int `json:"pac"`
	Shooting    int `json:"sho"`
	Passing     int `json:"pas"`
	Dribbling   int `json:"dri"`
	Defending   int `json:"def"`
	Physicality int `json:"phy"`

	ClubID   int    `json:"club"`
	Club     string `json:"club_name"`
	Nation   string `json:"nation_name"`
	NationID int    `json:"nation"`
	LeagueID int    `json:"league"`
	League   string `json:"league_name"`

	PsPrice      int `json:"ps_LCPrice"`
	PsPRP        int `json:"ps_PRP"`
	PsMinPrice   int `json:"ps_MinPrice"`
	PsMaxPrice   int `json:"ps_MaxPrice"`
	PcPrice      int `json:"pc_LCPrice"`
	PcPRP        int `json:"pc_PRP"`
	PcMinPrice   int `json:"pc_MinPrice"`
	PcMaxPrice   int `json:"pc_MaxPrice"`
	XboxPrice    int `json:"xbox_LCPrice"`
	XboxPRP      int `json:"xbox_PRP"`
	XboxMinPrice int `json:"xbox_MinPrice"`
	XboxMaxPrice int `json:"xbox_MaxPrice"`
}
type NamePlayer struct {
	ID          string `json:"ID"`
	Playerid    string `json:"playerid"`
	ResourceID  string `json:"resource_id"`
	Playername  string `json:"playername"`
	Rating      string `json:"rating"`
	Club        string `json:"club"`
	League      string `json:"league"`
	Nation      string `json:"nation"`
	Raretype    string `json:"raretype"`
	Rare        string `json:"rare"`
	Playerimage string `json:"playerimage"`
	Position    string `json:"position"`
	PosAlt      string `json:"pos_alt"`
	ClubName    string `json:"club_name"`
	LeagueName  string `json:"league_name"`
	CountryName string `json:"country_name"`
	CommonName  string `json:"common_name"`
	Pac         string `json:"pac"`
	Sho         string `json:"sho"`
	Pas         string `json:"pas"`
	Dri         string `json:"dri"`
	Def         string `json:"def"`
	Phy         string `json:"phy"`
	PlayerImage string `json:"player_image"`
}
type Player struct {
	ID                int         `json:"ID"`
	PlayerID          int         `json:"Player_ID"`
	PlayerResource    int         `json:"Player_Resource"`
	PlayerName        string      `json:"Player_Name"`
	PlayerCommon      string      `json:"Player_Common"`
	PlayerFullname    string      `json:"Player_Fullname"`
	AllNames          string      `json:"AllNames"`
	PlayerClub        int         `json:"Player_Club"`
	PlayerNation      int         `json:"Player_Nation"`
	PlayerLeague      int         `json:"Player_League"`
	PlayerPace        int         `json:"Player_Pace"`
	PlayerShooting    int         `json:"Player_Shooting"`
	PlayerPassing     int         `json:"Player_Passing"`
	PlayerDribbling   int         `json:"Player_Dribbling"`
	PlayerDefending   int         `json:"Player_Defending"`
	PlayerHeading     int         `json:"Player_Heading"`
	PlayerPosition    string      `json:"Player_Position"`
	PlayerFoot        string      `json:"Player_Foot"`
	PlayerHeight      string      `json:"Player_Height"`
	PlayerDOB         string      `json:"Player_DOB"`
	PlayerRating      int         `json:"Player_Rating"`
	Rare              int         `json:"Rare"`
	RareType          int         `json:"Rare_Type"`
	WeakFoot          int         `json:"Weak_Foot"`
	Skills            int         `json:"Skills"`
	DefensiveWorkrate string      `json:"Defensive_Workrate"`
	AttackWorkrate    string      `json:"Attack_Workrate"`
	SQUADID           interface{} `json:"SQUAD_ID"`
	StartDate         string      `json:"Start_Date"`
	Revision          string      `json:"Revision"`
	Specialities      interface{} `json:"Specialities"`
	Traits            string      `json:"Traits"`
	Acceleration      int         `json:"Acceleration"`
	Aggression        int         `json:"Aggression"`
	Agility           int         `json:"Agility"`
	Balance           int         `json:"Balance"`
	Ballcontrol       int         `json:"Ballcontrol"`
	Crossing          int         `json:"Crossing"`
	Curve             int         `json:"Curve"`
	Dribbling         int         `json:"Dribbling"`
	Finishing         int         `json:"Finishing"`
	Headingaccuracy   int         `json:"Headingaccuracy"`
	Interceptions     int         `json:"Interceptions"`
	Jumping           int         `json:"Jumping"`
	Longpassing       int         `json:"Longpassing"`
	Longshots         int         `json:"Longshots"`
	Marking           int         `json:"Marking"`
	Penalties         int         `json:"Penalties"`
	Positioning       int         `json:"Positioning"`
	Reactions         int         `json:"Reactions"`
	Shortpassing      int         `json:"Shortpassing"`
	Freekickaccuracy  int         `json:"Freekickaccuracy"`
	Shotpower         int         `json:"Shotpower"`
	Slidingtackle     int         `json:"Slidingtackle"`
	Sprintspeed       int         `json:"Sprintspeed"`
	Standingtackle    int         `json:"Standingtackle"`
	Stamina           int         `json:"Stamina"`
	Strength          int         `json:"Strength"`
	Vision            int         `json:"Vision"`
	Volleys           int         `json:"Volleys"`
	SpecialImage      int         `json:"Special_Image"`
	BotUpdate         int         `json:"BotUpdate"`
	PlayerWeight      int         `json:"Player_Weight"`
	InPacks           int         `json:"in_packs"`
	Active            int         `json:"active"`
	TotalUp           int         `json:"totalUp"`
	TotalDown         int         `json:"totalDown"`
	Likes             int         `json:"likes"`
	TotalStats        int         `json:"TotalStats"`
	OldSkills         interface{} `json:"Old_Skills"`
	OldWeakFoot       interface{} `json:"Old_Weak_Foot"`
	Composure         int         `json:"Composure"`
	InternationalRep  int         `json:"InternationalRep"`
	Untradeable       int         `json:"untradeable"`
	TotalIGS          int         `json:"TotalIGS"`
	Boost             int         `json:"boost"`
	Realface          interface{} `json:"realface"`
	Groups            string      `json:"Groups"`
	Bodytypecode      int         `json:"bodytypecode"`
	RankID            interface{} `json:"rankID"`
	GuidAssetID       interface{} `json:"guidAssetID"`
	PlayerPosition2   string      `json:"Player_Position2"`
	PlayerPosition3   string      `json:"Player_Position3"`
	PlayerPosition4   string      `json:"Player_Position4"`
	AcceleRATE        string      `json:"AcceleRATE"`
	Inserted          string      `json:"inserted"`
	ClubName          string      `json:"club_name"`
	LeagueName        string      `json:"league_name"`
	CountryName       string      `json:"country_name"`
	GameTotal         int         `json:"game_total"`
	GoalRatio         string      `json:"goal_ratio"`
	AssistsRatio      string      `json:"assists_ratio"`
	YellowCardRatio   string      `json:"yellow_card_ratio"`
	SetInfo           interface{} `json:"setInfo"`
	PriceType         string      `json:"price_type"`
	LCPrice           int         `json:"LCPrice"`
	LCPrice2          int         `json:"LCPrice2"`
	LCPrice3          int         `json:"LCPrice3"`
	LCPrice4          int         `json:"LCPrice4"`
	LCPrice5          int         `json:"LCPrice5"`
	Updatedon         int         `json:"updatedon"`
	MinPrice          int         `json:"MinPrice"`
	MaxPrice          int         `json:"MaxPrice"`
	PBP               int         `json:"PBP"`
	PlayerImage       int         `json:"player_image"`
	TopChemUsed       []ChemUsed  `json:"topChemUsed"`
	TopChemPoll       TopChemPoll `json:"topChemPoll"`
}

type ChemUsed struct {
	Playstyle int    `json:"playstyle"`
	Pct       int    `json:"pct"`
	Name      string `json:"name"`
}

type TopChemPoll struct {
	PollGraphData []PollGraphData `json:"pollGraphData"`
}

type PollGraphData struct {
	Artist     int `json:"artist"`
	Architect  int `json:"architect"`
	Anchor     int `json:"anchor"`
	Backbone   int `json:"backbone"`
	Basic      int `json:"basic"`
	Catalyst   int `json:"catalyst"`
	Deadeye    int `json:"deadeye"`
	Engine     int `json:"engine"`
	Finisher   int `json:"finisher"`
	Gladiator  int `json:"gladiator"`
	Guardian   int `json:"guardian"`
	Hawk       int `json:"hawk"`
	Hunter     int `json:"hunter"`
	Maestro    int `json:"maestro"`
	Marksman   int `json:"marksman"`
	Sniper     int `json:"sniper"`
	Powerhouse int `json:"powerhouse"`
	Sentinel   int `json:"sentinel"`
	Shadow     int `json:"shadow"`
}

func (p FilteredPlayer) Price(platform string) int {
	switch platform {
	case "PC":
		return p.PcPrice
	case "XB":
		return p.XboxPrice
	default:
		return p.PsPrice
	}
}

func (p FilteredPlayer) MinPrice(platform string) int {
	switch platform {
	case "PC":
		return p.PcMinPrice
	case "XB":
		return p.XboxMinPrice
	default:
		return p.PsMinPrice
	}
}

func (p FilteredPlayer) MaxPrice(platform string) int {
	switch platform {
	case "PC":
		return p.PcMaxPrice
	case "XB":
		return p.XboxMaxPrice
	default:
		return p.PsMaxPrice
	}
}

func (p FilteredPlayer) PRP(platform string) int {
	switch platform {
	case "PC":
		return p.PcPRP
	case "XB":
		return p.XboxPRP
	default:
		return p.PsPRP
	}
}

type playersDataFilter struct {
	Players []FilteredPlayer `json:"data"`
}
type playersDataName struct {
	Players []NamePlayer `json:"data"`
}
type playersData struct {
	Players []Player `json:"data"`
}
