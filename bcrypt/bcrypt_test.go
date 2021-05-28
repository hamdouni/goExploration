/*
Le principe de bcrypt c'est de hasher du texte (mot de passe) avec :
- un algo de hashage couteux en temps de process pour rendre infaisable l'attaque par test
- l'utilisation de salt pour rendre impossible l'attaque par dictionnaires : le salt est une chaine random ajoutée lors du hash pour rendre unique le résultat. Si on passe 2 fois de suite bcrypt sur le même texte, on obtient 2 résultats différents

Du coup, comment marche bcrypt pour tester la validité d'un mot de passe si le hash n'est jamais le même ?

Cela vient du fait que le hash contient le salt, et qu'on peut l'utiliser avec le mdp en clair pour voir si ça génère le même résultat. Explication détaillée :
- on prend un mot de passe en claire pwd := "thisisapassword"
- on le bcrypt pour obtenir un hash := "$2a$10$jWpY5H7yvZWr8ut44H.s7Ox7H.rGNgwwqdCRgXLBYm217Ap5.mgfG"
- en prefix on a le hash + des infos de version + le cost utilisé + d'autres caractères de padding :
	$2a$10$jWpY5H7yvZWr8ut44H.s7Ox7H.rGNgwwqdCRgXLBYm217Ap5.mgfG
	_vv_cc_sssssssssssssssssssssshhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh

	- 2a pour version majeur et mineur
	- 10 représente le cost (par défaut)
	- jWpY5H7yvZWr8ut44H.s7O est le salt
	- le reste est le hash du pass qui utilise le salt

	Imaginons qu'on a le hash en bdd, et qu'on veut savoir si le user qui a tapé en clair "thisispassword" correspond au hash en bdd. Comment on fait ?
	A partir du hash en bdd, on extrait le salt, version et cost : cf code source bcrypt.go
		type hashed struct {
			hash  []byte
			salt  []byte
			cost  int // allowed range is MinCost to MaxCost
			major byte
			minor byte
		}
	On lance l'algo bcrypt sur "thisispassowrd" avec ces infos.
	Comme on utilise le meme salt et cost, on obtient le même hash.
*/
package bcrypt

import "testing"

func TestBCrypt(t *testing.T) {

	// façon naturelle de hasher un mot de passe
	hash, err := GenerateFromPassword([]byte("thisisapassword"), 10)
	if err != nil {
		t.Fatalf("error %v\n", err)
	}
	t.Logf("hash %s\n", hash)

	// extrait les infos depuis le mot de passe hashé
	p, err := newFromHash(hash)
	if err != nil {
		t.Fatalf("error %v\n", err)
	}
	t.Logf("salt %s", p.salt) // le salt qui a été utilisé
	t.Logf("cost %d", p.cost) // le cost
	t.Logf("encr %s", p.hash) // l'encryption du mot de passe

	// vérifie que ça donne le résultat attendu en faisant l'encryption
	// nous même en précisant le cost et le salt obtenu juste avant
	pass, err := bcrypt([]byte("thisisapassword"), p.cost, p.salt)
	if err != nil {
		t.Fatalf("error %v\n", err)
	}
	t.Logf("pass %s", pass)
}
