package service

import "yu-croco/ddd_on_golang/app/domain/model"

func CalculateDamage(hunter *model.Hunter, monster *model.Monster) int {
	if monster.DefencePower >= hunter.OffensePower {
		return hunter.OffensePower
	} else {
		return hunter.OffensePower*2 - monster.DefencePower
	}
}
