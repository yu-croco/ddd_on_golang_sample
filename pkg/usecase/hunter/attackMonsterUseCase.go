package hunter

import (
	model2 "yu-croco/ddd_on_golang/pkg/domain/model"
	repository2 "yu-croco/ddd_on_golang/pkg/domain/repository"
	service2 "yu-croco/ddd_on_golang/pkg/domain/service"
	errors2 "yu-croco/ddd_on_golang/pkg/errors"
)

type attackMonsterUseCaseImpl struct {
	HunterId          model2.HunterId
	MonsterId         model2.MonsterId
	HunterRepository  repository2.HunterRepository
	MonsterRepository repository2.MonsterRepository
}

type attackMonsterUseCase interface {
	Exec() (*model2.Monster, *errors2.AppError)
}

func NewAttackMonsterUseCaseImpl(hunterId model2.HunterId, monsterId model2.MonsterId,
	hunterRepository repository2.HunterRepository,
	monsterRepository repository2.MonsterRepository) attackMonsterUseCase {

	return attackMonsterUseCaseImpl{
		HunterId:          hunterId,
		MonsterId:         monsterId,
		HunterRepository:  hunterRepository,
		MonsterRepository: monsterRepository,
	}
}

func (impl attackMonsterUseCaseImpl) Exec() (*model2.Monster, *errors2.AppError) {

	hunter, hunterFindErr := impl.HunterRepository.FindById(impl.HunterId)
	if hunterFindErr.HasErrors() {
		return nil, hunterFindErr
	}

	monster, monsterFindErr := impl.MonsterRepository.FindById(impl.MonsterId)
	if monsterFindErr.HasErrors() {
		return nil, monsterFindErr
	}

	hunterAttackDamage := service2.CalculateAttackMonsterDamage(hunter, monster)
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
