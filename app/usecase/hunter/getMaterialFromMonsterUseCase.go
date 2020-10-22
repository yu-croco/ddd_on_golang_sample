package hunter

import (
	"yu-croco/ddd_on_golang/app/domain/model"
	"yu-croco/ddd_on_golang/app/domain/repository"
	"yu-croco/ddd_on_golang/app/errors"
)

type getMaterialFromMonsterUseCaseImpl struct {
	HunterId          model.HunterId
	MonsterId         model.MonsterId
	HunterRepository  repository.HunterRepository
	MonsterRepository repository.MonsterRepository
}
type getMaterialFromMonsterUseCase interface {
	Exec() (*model.MonsterMaterial, *errors.AppError)
}

func NewGetMaterialFromMonsterUseCaseImpl(hunterId model.HunterId, monsterId model.MonsterId,
	hunterRepository repository.HunterRepository,
	monsterRepository repository.MonsterRepository) getMaterialFromMonsterUseCase {

	return getMaterialFromMonsterUseCaseImpl{
		HunterId:          hunterId,
		MonsterId:         monsterId,
		HunterRepository:  hunterRepository,
		MonsterRepository: monsterRepository,
	}
}

func (impl getMaterialFromMonsterUseCaseImpl) Exec() (*model.MonsterMaterial, *errors.AppError) {

	hunter, hunterFindErr := impl.HunterRepository.FindById(impl.HunterId)
	if hunterFindErr.HasErrors() {
		return nil, hunterFindErr
	}

	monster, monsterFindErr := impl.MonsterRepository.FindById(impl.MonsterId)
	if monsterFindErr.HasErrors() {
		return nil, monsterFindErr
	}

	takenMaterial, takeErr := hunter.GetMonsterMaterial(monster)
	if takeErr.HasErrors() {
		return nil, takeErr
	}

	impl.HunterRepository.AddMonsterMaterial(hunter, takenMaterial)

	return takenMaterial, nil
}
