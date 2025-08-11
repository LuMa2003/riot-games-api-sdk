package lol

import (
	"fmt"

	"github.com/LuMa2003/riot-games-api-sdk/internal/api"
)

func GetSummoner(region string, puuid string) (Summoner, error) {
	return api.Get[Summoner](fmt.Sprintf("https://%s.api.riotgames.com/riot/summoner/v4/summoners/by-puuid/%s", region, puuid))
}