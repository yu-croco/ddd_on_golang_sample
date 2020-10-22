package hunter

import (
	"yu-croco/ddd_on_golang/app/domain/model"
	"yu-croco/ddd_on_golang/app/domain/repository"
	"yu-croco/ddd_on_golang/app/errors"
)

func GetMaterialFromMonsterUseCase(hunterId model.HunterId, monsterId model.MonsterId,
	hunterRepository repository.HunterRepository,
	monsterRepository repository.MonsterRepository) (*model.MonsterMaterial, *errors.AppError) {

	hunter, hunterFindErr := hunterRepository.FindById(hunterId)
	if hunterFindErr.HasErrors() {
		return nil, hunterFindErr
	}

	monster, monsterFindErr := monsterRepository.FindById(monsterId)
	if monsterFindErr.HasErrors() {
		return nil, monsterFindErr
	}

	takenMaterial, takeErr := hunter.GetMonsterMaterial(monster)
	if takeErr.HasErrors() {
		return nil, takeErr
	}

	hunterRepository.AddMonsterMaterial(hunter, takenMaterial)

	return takenMaterial, nil
}
