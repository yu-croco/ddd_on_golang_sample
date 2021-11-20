package service

import (
	model2 "yu-croco/ddd_on_golang/pkg/domain/model"
)

func CalculateAttackHunterDamage(monster *model2.Monster, hunter *model2.Hunter) model2.MonsterOffensePower {
	if hunter.DefencePower.BiggerOrSameThan(monster.OffensePower) {
		return monster.OffensePower
	} else {
		return monster.OffensePower.Twice().Minus(hunter.DefencePower)
	}
}
