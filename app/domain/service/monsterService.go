package service

import "yu-croco/ddd_on_golang/app/domain/model"

func CalculateAttackHunterDamage(monster *model.Monster, hunter *model.Hunter) model.MonsterOffensePower {
	if hunter.DefencePower.BiggerOrSameThan(monster.OffensePower) {
		return monster.OffensePower
	} else {
		return monster.OffensePower.Twice().Minus(hunter.DefencePower)
	}
}
