package account

import (
	"fmt"

	"github.com/LuMa2003/riot-games-api-sdk/internal/api"
)

const baseUrl = "https://europe.api.riotgames.com"

func GetAccount(gameName string, tagLine string) (Account, error) {
	return api.Get[Account](fmt.Sprintf("%s/riot/account/v1/accounts/by-riot-id/%s/%s", baseUrl, gameName, tagLine))
}

func GetRegion(game string, puuid string) (Region, error) {
	return api.Get[Region](fmt.Sprintf("%s/riot/account/v1/region/by-game/%s/by-puuid/%s", baseUrl, game, puuid))
}