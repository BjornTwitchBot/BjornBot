package main

import (
	"github.com/BjornTwitchBot/BjornBot/Godeps/_workspace/src/github.com/fabioxgn/go-bot"
	"math/rand"
	"time"
)

func eelfacts(command *bot.Cmd) (msg string, err error) {
	eelFacts := [25]string{
		"There are more than 700 different kinds, or species, of eels",
		"Depending on the species, eels can grow to be anywhere between 4 inches to 11 1/2 feet long",
		"Eels are smooth",
		"They eat a variety of animals such as worms, snails, frogs, shrimp, mussels, lizards and other small fish. They generally hunt for food at night.",
		"The moray eel is the most widespread eel in the world and all of the species live in tropical seas",
		"Electric eels are not related to eels, but are more closely related to catfish and carp",
		"an Electric Eels' attack is about five times the amount of power that is in a standard wall socket",
		"In 2010, Greenpeace International has added the American eel to its seafood red list, a list of fish that are commonly sold in supermarkets around the world, and which have a very high risk of being sourced from unsustainable fisheries",
		"American eel's hatch as Leptocephali, then become “glass eels”, then “elvers.” Upon reaching their fresh water destination, they transform one more time into “yellow eels.” American eels reach sexual maturity in approximately 5 to 25 years. They die after spawning",
		"the American eel is at very high risk of extinction in the wild",
		"Captive European Eels have lived as long as 80 years, with some claims of living as long as 155",
		"While many eels are farm fraised, they have no been breed in captivity",
		"The Japanese freshwater eel produces a fluorescent protein. This protein is the basis of a new test to assess dangerous blood toxins that can trigger liver disease",
		"The aptly named 'Giant marbled eel' can grow up to 2 meters (6.6 ft) for females and 1.5 meters (4.9 ft) for males and can weigh up to 20.5 kilograms (45 lb), making it the largest species of anguillid eels.",
		"In 1876, as a young student in Austria, Sigmund Freud dissected hundreds of eels in search of the male sex organs. He had to concede failure in his first major published research paper, and turned to other issues in frustration",
		"The electric eel is a South American electric fish. Despite the name, it is not an eel, but rather a knifefish.",
		"Garden eel's live in burrows on the sea floor and get their name from their practice of poking their heads from their burrows while most of their bodies remain hidden. Since they tend to live in groups, the many eel heads 'growing' from the sea floor resemble the plants in a garden. The largest can be as much as an acre!",
		"Reef-associated roving coral groupers have been observed to recruit giant morays to help them hunt. The invitation to hunt is initiated by head-shaking. This style of hunting may allow morays to flush prey from niches not accessible to groupers",
		"Ribbon eels are carnivores, preying on small fish and other marine creatures.  They can attract their prey with their flared nostrils and then clamp down on them with their strong jaws and retreat into their burrows.",
		"When you swim in a creek, and an eel bites your cheek, that’s a moray",
		"In recent studies, it's been proven that most American eels will fail the marshmallow test when administered in a controlled environment",
		"The Japanese freshwater eel self identifies as Yandere",
		"In 1887 scientists had catalogued over 1000 species of eel. By the turn of the 20th century, the consensus was that most of them were 'just long fish'.",
		"Eels and Platypus share a common ancestor.",
		"The worlds longest recorded eel is named Philip and lives in West Virginia.",
	}

	rand.Seed(time.Now().UnixNano())
	msg = eelFacts[rand.Intn(len(eelFacts))]
	return
}

func init() {
	bot.RegisterCommand(
		"eelfacts",
		"Provides random facts about eels",
		"",
		eelfacts)
}
