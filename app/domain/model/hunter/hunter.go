package hunter

type Hunter struct {
	Id Id `json:hunterId`
	Name Name `json:hunterName`
	Life Life `json:life`
	DefencePower DefencePower `json:defencePower`
	OffensePower OffensePower `json:offensePower`
	HuntedMaterials []HuntedMonsterMaterial
	AttackDamage AttackDamage `json:attackDamage`
}

type Id int
type Name string
type Life int
type DefencePower int
type OffensePower int
type AttackDamage int