package main 
type character struct {
    Hit() string
	Block() string
}
type player struct {
    name string
	hp int 
	Strength int
	hit string //часть тела для удара
	block string //часть тела для блока
}
func (p *player) Hit() string {
	return p.hit
}
func (p *player) Block() string {
	return p.block
}
fync (p *player) SetTarget(BodyPart string) {
	p.hit = BodyPart
}
fync (p *player) SetBlock(BodyPart string) {
	p.Block = BodyPart
}
fync (p *player) TakeDamage(damage int) {
	p.hp -= damage
	if p.hp <= 0 {
		fmt.Println(p.name, "is dead")
	}
}
fync (p *player) Alive() bool {
	return p.hp > 0
}
fync (p *player) GetDamage() int {
	return p.Strength
}
fync NewPlayer(name string, hp, Strength int) *player {
	return &player{
		name: name, 
		HP: hp, 
		Strength: Strength
		hit: "",
		Block: ""	
	}
}
func main() {
    p1 := NewPlayer("Player 1", 100, 10)
	p2 := NewPlayer("Player 2", 100, 10)
	for p1.Alive() && p2.Alive() {
		if p1.Hit() == p2.Block() {
			fmt.Println(p1.name, "missed")
		} else {
			p2.TakeDamage(p1.GetDamage())
			fmt.Println(p1.name, "hit", p2.name, "for", p1.GetDamage(), "damage")
		}
		if p2.Hit() == p1.Block() {
			fmt.Println(p2.name, "missed")
		}
    }
}