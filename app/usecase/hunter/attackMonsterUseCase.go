package hunter

import (
	"yu-croco/ddd_on_golang/app/domain/model"
	"yu-croco/ddd_on_golang/app/domain/service"
	"yu-croco/ddd_on_golang/app/errors"
	"yu-croco/ddd_on_golang/app/infrastructure/repositoryImpl"
)

var hunterRepository = repositoryImpl.NewHunterRepositoryImpl()
var monsterRepository = repositoryImpl.NewMonsterRepositoryImpl()

func AttackMonsterUseCase(hunterId int, monsterId int) (*model.Monster, *errors.AppError) {
	hunter, hunterFindErr := hunterRepository.FindById(hunterId)
	if hunterFindErr.HasErrors() {
		return nil, hunterFindErr
	}

	monster, monsterFindErr := monsterRepository.FindById(monsterId)
	if monsterFindErr.HasErrors() {
		return nil, monsterFindErr
	}

	hunterAttackDamage := service.CalculateAttackMonsterDamage(hunter, monster)
	damagedMonster, attackErr := hunter.Attack(monster, hunterAttackDamage)

	if attackErr.HasErrors() {
		return nil, attackErr
	}

	updatedMonster, updateErr := monsterRepository.Update(damagedMonster)
	if updateErr.HasErrors() {
		return nil, updateErr
	}

	return updatedMonster, nil
}
