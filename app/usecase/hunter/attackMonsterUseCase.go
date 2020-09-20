package hunter

import (
	"yu-croco/ddd_on_golang/app/domain/helpers"
	"yu-croco/ddd_on_golang/app/domain/model"
	"yu-croco/ddd_on_golang/app/domain/service"
	"yu-croco/ddd_on_golang/app/infrastructure/repositoryImpl"
)

func AttackMonsterUseCase(hunterId int, monsterId int) (*model.Monster, *helpers.DomainError) {
	hunterRepository := repositoryImpl.HunterRepositoryImpl{}
	monsterRepository := repositoryImpl.MonsterRepositoryImpl{}
	hunter := hunterRepository.FindById(hunterId)
	monster := monsterRepository.FindById(monsterId)

	hunterAttackDamage := service.CalculateAttackMonsterDamage(hunter, monster)
	damagedMonster, attackErr := hunter.Attack(monster, hunterAttackDamage)

	if attackErr.HasErrors() {
		return nil, &attackErr
	}

	updatedMonster := monsterRepository.Update(damagedMonster)

	return updatedMonster, nil
}
