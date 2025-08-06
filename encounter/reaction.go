package encounter

import (
	"becmi/dice"
	"becmi/localization"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"log"
)

const (
	ReactionAttack = iota
	ReactionAggressive
	ReactionCautious
	ReactionNeutral
	ReactionFriendly
)

type Reaction struct {
	Roll     int
	Reaction int
}

func reactionRoll(modifier int) []Reaction {
	var rollmodifier int
	var reactions []Reaction

rollloop:
	for {
		roll := dice.Roll("2d6") + modifier + rollmodifier
		switch {
		case roll < 4:
			reactions = append(reactions, Reaction{roll, ReactionAttack})
			break rollloop
		case roll < 7:
			reactions = append(reactions, Reaction{roll, ReactionAggressive})
			rollmodifier = -4
		case roll < 10:
			reactions = append(reactions, Reaction{roll, ReactionCautious})
			rollmodifier = 0
		case roll < 12:
			reactions = append(reactions, Reaction{roll, ReactionNeutral})
			rollmodifier = 4
		case roll >= 12:
			reactions = append(reactions, Reaction{roll, ReactionFriendly})
			break rollloop
		default:
			log.Fatalf("Error: Reaction roll result %d out of range.", roll)
		}
	}

	return reactions
}

func (r Reaction) String() string {
	var reactionMsg *i18n.Message

	switch r.Reaction {
	case ReactionAttack:
		reactionMsg = &i18n.Message{
			ID:    "Reaction Attack",
			Other: "Monster attacks.",
		}
	case ReactionAggressive:
		reactionMsg = &i18n.Message{
			ID:    "Reaction Aggressive",
			Other: "Monster is aggressive (growls, threatens, etc.).",
		}
	case ReactionCautious:
		reactionMsg = &i18n.Message{
			ID:    "Reaction Cautious",
			Other: "Monster is cautious.",
		}
	case ReactionNeutral:
		reactionMsg = &i18n.Message{
			ID:    "Reaction Neutral",
			Other: "Monster is neutral.",
		}
	case ReactionFriendly:
		reactionMsg = &i18n.Message{
			ID:    "Reaction Friendly",
			Other: "Monster is friendly.",
		}
	default:
		log.Fatalf("Error: Reaction %d out of range.", r.Reaction)
	}

	reaction := localization.Locale[localization.LanguageSetting].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: reactionMsg})
	return reaction

}

func reactionString(reactions []Reaction) string {
	var reactionstring string
	for idx, reaction := range reactions {

		roundMsg := &i18n.Message{
			ID:    "Round",
			Other: "Round",
		}

		round := localization.Locale[localization.LanguageSetting].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: roundMsg})

		reactionstring += fmt.Sprintf("%s %2d: %s (%d)\n", round, idx+1, reaction, reaction.Roll)
	}
	return reactionstring
}

// RollReaction makes a Monster Reaction Roll, if necessary over multiple rounds and returns the Result as Text.
func RollReaction(modifier int) string {
	reactions := reactionRoll(modifier)
	return reactionString(reactions)
}
