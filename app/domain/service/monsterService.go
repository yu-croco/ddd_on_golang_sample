package service

import "yu-croco/ddd_on_golang/app/domain/model"

func CalculateAttackHunterDamage(monster *model.Monster, hunter *model.Hunter) int {
	if hunter.DefencePower >= monster.OffensePower {
		return monster.OffensePower
	} else {
		return monster.OffensePower*2 - hunter.DefencePower
	}
	return 0
}
