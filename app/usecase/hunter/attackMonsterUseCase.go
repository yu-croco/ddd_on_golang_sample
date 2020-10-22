package hunter

import (
	"yu-croco/ddd_on_golang/app/domain/model"
	"yu-croco/ddd_on_golang/app/domain/repository"
	"yu-croco/ddd_on_golang/app/domain/service"
	"yu-croco/ddd_on_golang/app/errors"
)

type attackMonsterUseCaseImpl struct {
	HunterId          model.HunterId
	MonsterId         model.MonsterId
	HunterRepository  repository.HunterRepository
	MonsterRepository repository.MonsterRepository
}

type attackMonsterUseCase interface {
	Exec() (*model.Monster, *errors.AppError)
}

func NewAttackMonsterUseCaseImpl(hunterId model.HunterId, monsterId model.MonsterId,
	hunterRepository repository.HunterRepository,
	monsterRepository repository.MonsterRepository) attackMonsterUseCase {

	return attackMonsterUseCaseImpl{
		HunterId:          hunterId,
		MonsterId:         monsterId,
		HunterRepository:  hunterRepository,
		MonsterRepository: monsterRepository,
	}
}

func (impl attackMonsterUseCaseImpl) Exec() (*model.Monster, *errors.AppError) {

	hunter, hunterFindErr := impl.HunterRepository.FindById(impl.HunterId)
	if hunterFindErr.HasErrors() {
		return nil, hunterFindErr
	}

	monster, monsterFindErr := impl.MonsterRepository.FindById(impl.MonsterId)
	if monsterFindErr.HasErrors() {
		return nil, monsterFindErr
	}

	hunterAttackDamage := service.CalculateAttackMonsterDamage(hunter, monster)
	damagedMonster, attackErr := hunter.Attack(monster, hunterAttackDamage)

	if attackErr.HasErrors() {
		return nil, attackErr
	}

	updatedMonster, updateErr := impl.MonsterRepository.Update(damagedMonster)
	if updateErr.HasErrors() {
		return nil, updateErr
	}

	return updatedMonster, nil
}
