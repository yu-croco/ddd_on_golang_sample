package service

import (
	model2 "yu-croco/ddd_on_golang/pkg/domain/model"
)

func CalculateAttackMonsterDamage(hunter *model2.Hunter, monster *model2.Monster) model2.HunterOffensePower {
	if monster.DefencePower.BiggerOrSameThan(hunter.OffensePower) {
		return hunter.OffensePower
	} else {
		return hunter.OffensePower.Twice().Minus(monster.DefencePower)
	}
}
