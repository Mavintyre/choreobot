package moderator

import "github.com/gempir/go-twitch-irc"

type Moderator interface {
	Moderate(u twitch.User, m twitch.Message)
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
