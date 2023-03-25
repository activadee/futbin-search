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
