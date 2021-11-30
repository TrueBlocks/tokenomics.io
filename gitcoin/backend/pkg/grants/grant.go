package grants

import (
	"encoding/json"
)

// Grant is one of the Gitcoin Grants
type Grant struct {
	Id                            uint64   `json:"id"`
	Active                        bool     `json:"active"`
	Title                         string   `json:"title"`
	Slug                          string   `json:"slug"`
	Description                   string   `json:"description"`
	ReferenceUrl                  string   `json:"reference_url"`
	Logo                          *string  `json:"logo"`
	AdminAddress                  string   `json:"admin_address"`
	AmountReceived                string   `json:"amount_received"`
	TokenAddress                  string   `json:"token_address"`
	TokenSymbol                   string   `json:"token_symbol"`
	ContractAddress               string   `json:"contract_address"`
	Meta                          MetaData `json:"metadata"`
	Network                       string   `json:"network"`
	RequiredGasPrice              string   `json:"required_gas_price"`
	AdminProfile                  Profile  `json:"admin_profile"`
	TeamMembers                   Profiles `json:"team_members"`
	ClrPredictionCurve            Points3d `json:"clr_prediction_curve"`
	ClrRoundNum                   string   `json:"clr_round_num"`
	IsClrActive                   bool     `json:"is_clr_active"`
	AmountReceivedInRound         string   `json:"amount_received_in_round"`
	PositiveRoundContributorCount uint64   `json:"positive_round_contributor_count"`
}
type Grants []Grant

func (g *Grant) ToJSON() string {
	str, err := json.Marshal(g)
	if err != nil {
		return ""
	}
	return string(str)
}

// Points are 2d and 3d arrays of floats
type Point2d [2]float64
type Points2d []Point2d
type Point3d [3]float64
type Points3d []Point3d

// MetaData carries information about the grant
type MetaData struct {
	Gem                                int64            `json:"gem"`
	Related                            *Points2d        `json:"related,omitempty"`
	Upcoming                           int64            `json:"upcoming"`
	WallOfLove                         [][2]interface{} `json:"wall_of_love"`
	UnsubscribedProfiles               []uint64         `json:"unsubscribed_profiles,omitempty"`
	LastCalcTimeRelated                *float64         `json:"last_calc_time_related,omitempty"`
	LastCalcTimeContributorCounts      float64          `json:"last_calc_time_contributor_counts"`
	LastCalcTimeSybilAndContribAmounts float64          `json:"last_calc_time_sybil_and_contrib_amounts"`
}

// Profile is a Github user
type Profile struct {
	Id           uint64            `json:"id"`
	Url          string            `json:"url"`
	Name         string            `json:"name"`
	Handle       string            `json:"handle"`
	Keywords     []string          `json:"keywords"`
	Position     uint64            `json:"position"`
	AvatarUrl    string            `json:"avatar_url"`
	GithubUrl    string            `json:"github_url"`
	TotalEarned  float64           `json:"total_earned"`
	Oranizations map[string]uint64 `json:"organizations"`
}
type Profiles []Profile

// Love is donor comments on a Grant
type Love map[string]uint64
type Loves []Love
