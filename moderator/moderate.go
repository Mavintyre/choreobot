package moderator

import "github.com/gempir/go-twitch-irc"

type Moderator interface {
	Moderate(u twitch.User, m twitch.Message)
}

type ConfigurableModerator interface {
	Moderator
	AddRule(r Rule)
	DisableRule(r Rule)
	ModifyRule(r Rule)
	GetRules() []Rule
	GetRuleByLabel(n string) Rule
}

type Rule interface {
}

type ruleManager struct {
}

func (m *ruleManager) Moderate(user twitch.User, msg twitch.Message) {

}

// Features to be implemented
//   Blacklisting
//	 Whitelist
//	 Permit
//	 Profanity avoidance
//// Quotas and Limits
//	 Caps Limit
//	 Repetition Spam
//	 Emote Limit
//   Symbols (?) Limit

// Limits should have decaying quotas and should probably take some sort of karma score into account.
