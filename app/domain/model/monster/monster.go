package monster

type Monster struct {
	Id           Id           `json:monsterId`
	Name         Name         `json:monsterName`
	Life         Life         `json:life`
	DefencePower DefencePower `json:defencePower`
	OffensePower OffensePower `json:offensePower`
	Materials    []MonsterMaterial
	AttackDamage AttackDamage `json:attackDamage`
}

type Id int
type Name string
type Life int
type DefencePower int
type OffensePower int
type AttackDamage int