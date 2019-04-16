package config

func addChannel(ChannelID string) {
	Channels = append(Channels, ChannelID)
}

func addTeamChannel(ChannelID string) {
	TeamChannels = append(TeamChannels, ChannelID)
}

func addTeam(Team Team) {
	i := 0
	exist := false
	for i = 0; i < len(TeamChannels); i++ {
		if TeamChannels[i] == Team.ChannelName {
			exist = true
			break
		}
	}
	if !exist {
		addTeamChannel(Team.ChannelName)
	}

	AllTeams = append(AllTeams, Team)
}

func addRole(Role Role) {
	//Check if the menu exist

	i := 0
	exist := false
	for i = 0; i < len(Channels); i++ {
		if Channels[i] == Role.ChannelName {
			exist = true
			break
		}
	}
	if !exist {
		addChannel(Role.ChannelName)
	}

	i = 0
	exist = false
	for i = 0; i < len(AllTeams); i++ {
		if AllTeams[i].Name == Role.Team.Name {
			exist = true
			break
		}
	}
	if !exist {
		addTeam(Role.Team)
	}

	AllRoles = append(AllRoles, Role)
}

func SetupRole() {
	LoupsGarous := Team{
		Name:        "Loup-Garou",
		ChannelName: "loup-garou",
		HasChannel:  true,
	}

	EnfantSauvages := Team{
		Name:        "Enfant Sauvage",
		ChannelName: "loup-garou",
		HasChannel:  false,
	}

	ChienLoups := Team{
		Name:        "Chien Loup",
		ChannelName: "loup-garou",
		HasChannel:  false,
	}

	Village := Team{
		Name:       "Village",
		HasChannel: false,
	}

	MDJ := Role{
		Name:  "MDJ",
		Image: "mdj.png",
		Team: Team{
			Name:       "MDJ",
			HasChannel: false,
		},
	}

	LG := Role{
		Name:  "Loup-Garou",
		Image: "lg.png",
		Description: `Ton rôle est de tuer tout les villageois, la nuit tu peux accéder à un channel où tu pourra décider de qui tuer.
                Your role is to kill every Villager, during the night you can decide who you're going to kill.`,
		Team: LoupsGarous,
	}

	LGB := Role{
		Name:        "Loup-Garou Blanc",
		Image:       "lgb.png",
		ChannelName: "loup-garou-blanc",
		Description: `Ton rôle est de tuer tout les villageois, la nuit tu peux accéder à un channel où tu pourra décider de qui tuer. Tout les deux nuits, tu peux décider de tuer l'un de tes collégues Loup.
  				Your role is to kill every Villager, during the night you can access a channel where you can decide who you're gonna kill. Every two nights, you can kill one of your Wolves friends.`,
		Team: LoupsGarous,
	}

	IPL := Role{
		Name:        "Infect Père des Loups",
		Image:       "ipl.png",
		ChannelName: "infect-pere-des-loups",
		Description: `Pendant la partie, tu peux transformer une victime de Loup-Garou en Loup-Garou, elle peut gagner avec vous.
  				During the game, you can save a victim of the Wolves into a Wolf, and she can win with you.`,
		Team: LoupsGarous,
	}

	ES := Role{
		Name:        "Enfant Sauvage",
		Image:       "enf.png",
		ChannelName: "enfant-sauvage",
		Description: `Tu décidera d'un maître durant ton premier tour, si ton maître meurt, tu deviendras un loup-garou.
  				You will decide who's your master during the first night, if he dies, you become a Wolf`,
		Team: EnfantSauvages,
	}

	CL := Role{
		Name:        "Chien-Loup",
		Image:       "chlg.png",
		ChannelName: "chien-loup",
		Description: `Durant la premiere nuit tu choisis si tu veux être un chien, ou un Loup-Garou. Si tu deviens un chien, tu rejoins le clan des Villageois.
  				During the first night, you'll choose if you wanna be a good dog (Villager), or a Wolf (Werewolf).`,
		Team: ChienLoups,
	}

	Fr := Role{
		Name:        "Frére",
		Image:       "fr.png",
		ChannelName: "freres",
		Description: `Pendant la nuit tu peux parler à tes frères.
  				During the night, you can talk with your siblings`,
		Team: Village,
	}

	Sorciere := Role{
		Name:        "Sorciére",
		Image:       "sor.png",
		ChannelName: "sorciere",
		Description: `Tu possèdes des potions : une de vie une de mort, durant la nuit tu peux décider d'en utiliser une, ou deux.
  				You have two potions : one to save, one to kill ; During the night you can use one, or two. Those aren't limitless.`,
		Team: Village,
	}

	Salvateur := Role{
		Name:        "Salvateur",
		Image:       "sal.png",
		ChannelName: "salvateur",
		Description: `Pendant la nuit tu peux protèger quelqu'un, cependant tu ne peux te protéger qu'une seule fois et tu ne peux pas protèger la même personne pendant deux tours.
				During the night, you can protect someone, you cannot protect him the next night and you can only protect yourself once.`,
		Team: Village,
	}

	Voyante := Role{
		Name:        "Voyante",
		Image:       "voy.png",
		ChannelName: "voyante",
		Description: `Pendant la nuit, tu peux voir la carte de quelqu'un d'autres.
				During the night, you can see someone else's card.`,
		Team: Village,
	}

	Chasseur := Role{
		Name:  "Chasseur",
		Image: "cha.png",
		Description: `A ton décés, tu peux décider de tirer ou non sur quelqu'un.
				When you die, you can shoot someone, or not.`,
		Team: Village,
	}

	PF := Role{
		Name:        "Petite Fille",
		Image:       "pf.png",
		ChannelName: "petite-fille",
		Description: `Pendant la nuit, tu peux jeter deux dès : un qui te permet de voir des messages de Loup Garou, sans consèquence, l'autre qui te permet de connaitre le pseudo d'un Loup. Si le second rate, un Loup-Garou peut voir ton pseudo.
				During the nice you can roll two dice : one to see Werewolve's message, no consequence if fail. The other one to know a Werewolve's name. If it fails, a Wolf will see yours.`,
		Team: Village,
	}

	Chaman := Role{
		Name:  "Chaman",
		Image: "cham.png",
		Description: `Pendant la nuit, tu peux voir ce que disent les morts.
				During the night, you can see what the Deads have to say.`,
		Team: Village,
	}

	Corbeau := Role{
		Name:        "Corbeau",
		Image:       "cor.png",
		ChannelName: "corbeau",
		Description: `Pendant la nuit, tu peux donner deux votes à une personne, ou pas.
				During the night, you can give two votes on a single person for the next day.`,
		Team: Village,
	}

	Cupidon := Role{
		Name:        "Cupidon",
		Image:       "cup.png",
		ChannelName: "cupidon",
		Description: `Durant le premier tour, tu décide d'un couple, si l'un des membres meurt, le second tombe aussi.
        				During the first night, you'll decide a couple, if one dies, the other one dies too.`,
		Team: Village,
	}

	Ankou := Role{
		Name:  "Ankou",
		Image: "ank.png",
		Description: `Après ta mort, tu peux continuer de voter pendant deux tours, les vivants ne verront pas ton vote.
				After your death, you can continue to vote the next two days. Those who are alive won't see it.`,
		Team: Village,
	}

	Ancien := Role{
		Name:  "Ancien",
		Image: "anc.png",
		Description: `Si le village te tue, le village perd ses pouvoirs.
				If you are killed by the Village, every Villager that has special powers lose it.`,
		Team: Village,
	}

	MO := Role{
		Name:  "Montreur d'ours",
		Image: "mo.png",
		Description: `Si la personne au dessus ou en dessous de toi est Loup ; tu grogne le matin.
        				If the person above, or below you is a Wolf ; you wil roar.`,
		Team: Village,
	}

	Ange := Role{
		Name:  "Ange",
		Image: "ang.png",
		Description: `Ton but est de te suicider pendant le premier tour (Nuit ou Jour).
        				Your goal is to die during the first night or day.`,
		Team: Village,
	}

	addRole(MDJ)
	addRole(LG)
	addRole(LGB)
	addRole(IPL)
	addRole(ES)
	addRole(CL)
	addRole(Fr)
	addRole(Sorciere)
	addRole(Salvateur)
	addRole(Voyante)
	addRole(Chasseur)
	addRole(PF)
	addRole(Chaman)
	addRole(Corbeau)
	addRole(Cupidon)
	addRole(Ankou)
	addRole(Ancien)
	addRole(MO)
	addRole(Ange)

	SpecialChannels = append(SpecialChannels, "lg-gamestats")
	SpecialChannels = append(SpecialChannels, "lg-vote")
	SpecialChannels = append(SpecialChannels, "lg-morts")

}
