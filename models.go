package beam

import "time"

type (
	Error struct {
		Error      string `json:"error"`
		StatusCode int    `json:"statusCode"`
	}

	Achievement struct {
		Slug        string `json:"slug"` // maxLength: 100
		Name        string `json:"name"`
		Description string `json:"description"`
		Sparks      uint   `json:"sparks"`
	}

	AchievementEarning struct {
		// The unique ID of the earning.
		ID uint `json:"id"`

		// Indicates whether the achievement was earned.
		Earned bool `json:"earned"`

		// Unit interval indicating the progress.
		Progress int `json:"progress"` // ≤ 1

		// The name of the achievement.
		Name string `json:"achievement"`

		// The ID of the user this earning belongs to.
		User uint `json:"user"`

		Achievement *Achievement `json:"Achievement"`
	}

	Announcement struct {
		// The text of the announcement.
		Text string `json:"text"`

		// A java-script to be executed on the browsers.
		Script string `json:"script"`

		// Unknown.
		Scope interface{} `json:"scope"`
	}

	CPMAnalytic struct {
		Channel uint `json:"channel"`

		// The number of unique ad impressions.
		Impressions uint `json:"impressions"`

		// The amount of money awarded from the impressions, in cents.
		Payout int `json:"payout"`
	}

	Channel struct {
		// The unique ID of the channel.
		ID uint

		// The ID of the user owning the channel.
		UserID uint

		// The name and url of the channel.
		Token string

		// Indicates if the channel is active.
		Online bool

		// True if featureLevel is > 0.
		Featured bool

		// Featuring level.
		FeatureLevel int

		// Indicates if the channel is partnered.
		Partnered bool

		// The id of the transcoding profile.
		TranscodingProfileID uint

		// Indicates if the channel is suspended.
		Suspended bool

		// The title of the channel.
		Name string

		// The target audience of the channel.
		Audience string // (family, teen, 18+)

		// Amount of unique viewers that ever viewed this channel.
		ViewersTotal uint

		// Amount of current viewers.
		ViewersCurrent uint

		// Amount of followers.
		NumFollowers uint

		// The description of the channel, can contain HTML.
		Description string

		// The ID of the game type.
		TypeID uint

		// Indicates if that channel is interactive.
		Interactive bool

		// The ID of the interactive game used.
		InteractiveGameID uint

		// The ftl stream id.
		FTL uint

		// Indicates if the channel has vod saved.
		HasVOD bool

		// ISO 639 language id.
		LanguageID string

		// The ID of the cover resource.
		CoverID uint

		// The resource ID of the thumbnail.
		ThumbnailID uint

		// Indicates if the channel has vod recording enabled.
		VODsEnabled bool

		// The costream that the channel is in, if any.
		CostreamID *UUID
	}

	ChannelAdvanced struct {
		// The game type.
		Type *GameType

		User struct {
			*User
			Groups []UserGroup `json:"groups"`
		}
	}

	ChannelPreferences struct {
		// The text used when sharing the stream. The template parameter %URL% will be replaced with the channel's URL. The template parameter %USER% will be replaced with the channel's name.
		ShareText string

		// Specifies whether links are allowed in the chat.
		ChannelLinksAllowed bool `json:"channel:links:allowed"`

		// Specifies whether links are clickable in the chat.
		ChannelLinksClickable bool `json:"channel:links:clickable"`

		// Interval required between each chat message.
		ChannelSlowchat int `json:"channel:slowchat"`

		// The message to be sent when a user subscribes to the channel. The template parameter %USER% will be replaced with the subscriber's name.
		ChannelNotifySubscribemessage string `json:"channel:notify:subscribemessage"`

		// Indicates whether a notification should be shown upon subscription.
		ChannelNotifySubscribe bool `json:"channel:notify:subscribe"`

		// The message to be sent when a user follows the channel. The template parameter "%USER%" will be replaced with the follower's name.
		ChannelNotifyFollowmessage string `json:"channel:notify:followmessage"`

		// Indicates whether a notification should be shown upon follow.
		ChannelNotifyFollow bool `json:"channel:notify:follow"`

		// The message to be sent when a user hosts the channel. The template parameter "%USER%" will be replaced with the hoster's name.
		ChannelNotifyHostedBy string `json:"channel:notify:hostedBy"`

		// The message to be sent when the channel hosts another. The template parameter "%USER%" will be replaced with the hostee's name.
		ChannelNotifyHosting string `json:"channel:notify:hosting"`

		// The text to be added to the subscription email.
		ChannelPartnerSubMail string `json:"channel:partner:submail"`

		// Indicates whether to mute when the streamer opens his own stream.
		ChannelPlayerMuteOwn bool `json:"channel:player:muteOwn"`

		// Indicates whether the tweet button should be shown.
		ChannelTweetEnabled bool `json:"channel:tweet:enabled"`

		// The message to be used when a user tweets about the channel. The template parameter %URL% will be replaced with the share url.
		ChannelTweetBody string `json:"channel:tweet:body"`
	}

	ChatUser struct {
		// The user ID that this chat user belongs to.
		UserID uint

		// The username.
		UserName string

		// The roles the user has in this chat.
		UserRoles []string

		// Indicates that the user is lurking, requires the chat:view_lurkers permission.
		Lurking bool
	}

	CheckoutItem struct {
		// The ID of the item to be purchased.
		ID uint

		// The amount of items to be purchased.
		Quantity uint
	}

	Confetti struct {
		// The Confetti's ID.
		ID uint

		// The owner of the Confetti; this channel can edit it.
		ChannelID uint

		// The settings configure the confetti. Most options are customizable math expressions parsed with expr-eval.
		Settings struct {
			// Expression to be evaluated to determine the number of generated confetti particles. The window height and width will be passed into the expression. The number of confetti particles will be capped based on the device profile, however. For example: "count": "35 * width / 100".
			Count string

			Particles []struct {
				// Weighted probability that a generated confetti particle will be this particle. Weighting is done based on the sum all probabilities.
				Probability int

				// Expression to be evaluated to determine the vertical velocity, in pixels per millisecond, of the particle. It's called with the window width and height. For example: "velocity": "height / 1000 / 4".
				Velocity string

				// Expression to be evaluated to determine the "depth" of the particle, which will be used to scale its size and velocity to achieve a pseudo-3d effect. It's called with the window width and height. For example: "zdepth": "1 + random(3)".
				ZDepth string

				// Expression to determine the magnitude of the wave, in pixels, that each particle follows on its horizontal axis as it floats down. It's called with the window width and height. For example: "wiggleMagnitude": "15 + random(10)".
				WiggleMagnitude string

				// Expression to determine the period of the wave, in milliseconds, that each particle follows on its horizontal axis as it floats down. It's called with the window width and height. For example: "wigglePeriod": "1000 + random(500)".
				WigglePeriod string

				// Expression to determine how long each particle lives for, in milliseconds, before being destroyed. It's called with the window width and height. For example: "lifetime": "1200 + random(2000)".
				LifeTime string

				// Expression called continuously to determine the opacity. It's called a lifetime values, which starts at zero and goes to one as the time passes, and should return an opacity value in the range from zero to one. For example: "fader": "1 - lifetime^3".
				Fader string

				// Describes the confetti graphic to be drawn.
				Draw interface{} // (ConfettiCircle ∪ ConfettiRectangle ∪ ConfettiImage)
			}
		}
	}

	ConfettiCircle struct {
		// Must be set to circle.
		Shape string

		// Expression to determine the diameter of the confetti circle in pixels. For example: "size": "4 + random(2)"
		Size string

		// Expression to determine the length of time the circle takes to do a vertical, pseudo-3d spin around the x axis. For example: "flipPeriod": "500 + random(1000)"
		FlipPeriod string

		Colors []struct {
			// Weighted probability that a generated confetti particle will be this color. Weighting is done based on the sum all probabilities.
			Probability int

			// Expression used to generate the red color component, should return a value between 0 and 255 inclusive.
			R string

			// Expression used to generate the red color component, should return a value between 0 and 255 inclusive.
			G string

			// Expression used to generate the red color component, should return a value between 0 and 255 inclusive.
			B string
		}
	}

	ConfettiImage struct {
		// Must be set to image.
		Shape string

		// Expression to determine the scale of the image. For example: "scale": "0.5 + random(1)"
		Scale string

		// Expression to determine the length of time the images takes to do a full spin around the z axis. For example: "spinPeriod": "1000 + random(1000)"
		SpinPeriod string

		// Base-64 encoded image data. This will usually be a long string that starts with data:image/png;base64,. Make sure to minify the image before inserting it here to prevent exceeding size limits.
		Data string
	}

	ConfettiRectangle struct {
		// Must be set to rectangle.
		Shape string

		// Expression to determine the width of the confetti circle in pixels. For example: "size": "8 + random(2)"
		Width string

		// Expression to determine the height of the confetti circle in pixels. For example: "size": "4 + random(1)"
		Height string

		// Expression to determine the length of time the circle takes to do a vertical, pseudo-3d spin around the x axis. For example: "flipPeriod": "500 + random(1000)"
		FlipPeriod string

		Colors []struct {
			// Weighted probability that a generated confetti particle will be this color. Weighting is done based on the sum all probabilities.
			Probability int

			// Expression used to generate the red color component, should return a value between 0 and 255 inclusive.
			R string

			// Expression used to generate the red color component, should return a value between 0 and 255 inclusive.
			G string

			// Expression used to generate the red color component, should return a value between 0 and 255 inclusive.
			B string
		}
	}

	DiscordBot struct {
		// The unique ID of the Discord bot.
		ID uint

		// The channel ID linked to this Discord integration.
		ChannelID uint

		// The server ID that this bot belongs to.
		GuildID string

		// Determines which users can get invites to the Discord channel. It can be set to "everyone", "followers", "subs", or "noone" to disable invites. Defaults to "noone".
		InviteSetting string // (EVERYONE, FOLLOWERS, SUBS, NOONE)

		// The ID of the server channel that new users are invited to.
		InviteChannel string

		// Whether the user state should be changed to "Playing GAME on Beam" when the user goes live.
		LiveUpdateState bool

		// Channel ID that chat messages will be mirrored to by the bot. If this is null, chat messages will not be mirrored.
		LiveChatChannel string

		// Channel ID that chat announcements will be mirrored to by the bot. If this is null, chat messages will not be mirrored.
		LiveAnnounceChannel string

		// A list of Discord roles who will be able to use subscriber emotes. (This only has an effect for partnered channels.)
		SyncEmoteRoles []uint

		Roles []struct {
			ID uint

			// Describes what group of people are given this role.
			Type string // (SUBS, FOLLOWERS)

			RoleID string

			// Number of days subscribers have until they're kicked from the server if their subscription expires. It defaults to 14 days and may be set to null to disable kicking. Violence is never the answer.
			GracePeriod uint
		}
	}

	DiscordChannel struct {
		// The unique ID of the Discord channel.
		ID string

		// The human name of the channel.
		Name string

		// Either "voice" or "text".
		Type string

		// Whether the channel is private or not.
		Private bool
	}

	DiscordRole struct {
		// The unique ID of the Discord role.
		ID string

		// The human name of the role.
		Name string

		// Hex color code for the channel, starting with a pound # symbol
		Color string
	}

	EmojiRankAnalytic struct {
		Channel uint
		Emoji   string

		// The number of times this emoji was used in the queried time period.
		Count uint

		Time *IsoDate
	}

	EmoticonGroup map[string]struct {
		// X coordinate of the emote.
		X uint
		// Y coordinate of the emote.
		Y uint
	}

	ExpandedChannel struct {
		// A resource object representing the thumbnail.
		Thumbnail *Resource

		// A resource object representing the cover.
		Cover *Resource

		// A resource object representing the badge.
		Badge *Resource

		// Unused, deprecated.
		Cache []interface{}

		// The channel preferences.
		Preferences *ChannelPreferences
	}

	ExpandedRedeemable struct {
		// Embedded owner.
		Owner *User

		// Embedded redeemer.
		RedeemedBy *User
	}

	ExpandedShare struct {
		// The embedded user the share targets.
		User *User
	}

	ExpandedTranscodingProfile struct {
		// The transcodes for this profile.
		Transcodes []TranscodingProfileTranscode
	}

	FTLVideoManifest struct {
		Resolutions []struct {
			*VideoManifestResolution
			IceServers []struct {
				// The url of the ice server, can be stun or turn.
				URL string
				// Only set when url protocol is turn.
				UserName string
				// Only set when url protocol is turn.
				Credentials string
			}

			// The ID of the channel.
			RtcID uint
		}

		// Time the stream started.
		Since *IsoDate
	}

	FeatureSchedule struct {
		// The feature id.
		ID uint

		// Start time of the featuring.
		StartedAt *IsoDate

		// End time of the featuring
		EndedAt *IsoDate

		// auto means it was created by the featuring service.
		Type string // (auto, manual)

		// The user that scheduled the feature.
		UserID uint

		// The channel to be featured.
		ChannelID uint
	}

	Follow struct {
		// The follower user id.
		User uint

		// The followee channel id.
		Channel uint
	}

	FollowersAnalytic struct {
		Channel uint

		// The total number of followers the channel now has.
		Total uint

		// Increment that the total changed as a result of this event.
		Delta int

		// The ID of the user who followed, may not be present in aggregations.
		User int

		Time *IsoDate
	}

	FrontendVersion struct {
		// The version name.
		Version string

		// The version's display name for the frontend.
		DisplayName string
	}

	GameRankAnalytic struct {
		Channel uint

		// The number of streams playing this game in the time period. Note that this is recorded once per unique stream when it ends.
		Streams uint

		// The number of views the game received in the time period
		Views uint

		// The number of times the game was shared in the time period
		Shared uint

		Time *IsoDate
	}

	GameType struct {
		// The name of the parent type.
		Parent string

		// The description of the type.
		Description string

		// The source where the type has been imported from.
		Source string

		// Total amount of users watching this type of stream.
		ViewersCurrent uint

		// Amount of streams online with this type.
		Online uint
	}

	GameTypeLookup struct {
		// Whether this game type is an exact match to the query.
		Exact bool
	}

	// GameTypeSimple is a base game type.
	GameTypeSimple struct {
		// The unique ID of the game type.
		ID uint

		// The name of the type.
		Name string

		// The url to the types cover.
		CoverUrl string
	}

	// Ingest is an ingest definition.
	Ingest struct {
		// The name and location of the ingest.
		Name string

		// Url of the host.
		Host string

		// A websocket address that is used for ping tests.
		PingTest string

		// List of supported protocols
		Protocols []struct {
			// The protocol name.
			Type string // (ftl, rtmp)
		}
	}

	InteractiveConnectionInfo struct {
		// Websocket url to connect to.
		Address string

		// The connection key required for connecting to interactive.
		Key string

		// Influence multiplier
		Influence int

		Version struct {
			// The interactive game the channel is using.
			*InteractiveVersion

			// The parent game of the version.
			Game *InteractiveGame
		}
	}

	InteractiveControls struct {
		ReportInterval int
		Joysticks      []struct {
			// The unique ID of the control.
			ID uint

			Blueprint []struct {
				// The state the button belongs to.
				State string // maxLength: 100

				Grid string // maxLength: 100

				// X coordinate of the joystick.
				X uint

				// Y coordinate of the joystick.
				Y uint

				// Width of the joystick.
				Width int // 3 ≤ self ≤ 3

				// Height of the joystick.
				Height int // 3 ≤ self ≤ 3
			}

			// The analysis types enabled on this joystick.
			Analysis *InteractiveJoyStickAnalysis
		}
		Screens []struct {
			// The unique ID of the control.
			ID uint

			Blueprint []struct {
				// The state the button belongs to.
				State string // maxLength: 100
			}

			// The analysis types enabled on this screen control.
			Analysis *InteractiveScreenAnalysis
		}
		Tactiles []struct {
			// The unique ID of the control.
			ID uint

			Blueprint []struct {
				// The state the button belongs to.
				State string // maxLength: 100

				Grid string // maxLength: 100

				// X coordinate of the tactile.
				X uint

				// Y coordinate of the tactile.
				Y uint

				// Width of the tactile.
				Width int // 1 ≤ self ≤ 4

				// Height of the tactile.
				Height int // 1 ≤ self ≤ 2
			}

			// The key to be bound to the tactile.
			Key uint

			// The text to show on the tactile.
			Text string

			Cost struct {
				Press struct {
					// Cost, in sparks, for a tactile press.
					Cost uint
				}
			}

			Cooldown struct {
				// Length of cooldown started when a user pushes this tactile. In milliseconds.
				Press int
			}

			// The analysis types enabled on this tactile.
			Analysis *InteractiveTactileAnalysis
		}
	}

	InteractiveGame struct {
		// The unique ID of the game.
		ID uint

		// The user ID of the creator.
		OwnerID uint

		// The name of the game.
		Name string

		// The url to the cover image.
		CoverURL string

		// The description of the game, may contain HTML.
		Description string

		// Indicates whether the game has any public published versions.
		HasPublishedVersions bool

		// Installation instructions, may contain HTML.
		Installation string
	}

	InteractiveGameListing struct {
		Versions []struct {
			// The ID of the version.
			ID uint

			// The semantic version string.
			Version string

			// The current state of the version. draft: Not in review, not published. pending: In review. published: Published and available to the public.
			State string // (draft, pending, published)

			// Value used to help ordering versions.
			VersionOrder uint
		}
		Owner *User
	}

	InteractiveJoyStickAnalysis struct {
		Coords struct {
			// Enable joystick mean analysis on this joystick.
			Mean bool

			// Enable standard deviation analysis on this joystick.
			StdDev bool
		}
	}

	InteractiveScreenAnalysis struct {
		Position struct {
			// Enable mean analysis for this screen control.
			Mean bool

			// Enable standard deviation analysis for this screen control.
			StdDev bool
		}

		// Enable click event for analysis for this screen control.
		Clicks bool
	}

	InteractiveTactileAnalysis struct {
		// Enable holding analysis for the tactile.
		Holding bool

		// Enable frequency analysis for the tactile.
		Frequency bool
	}

	InteractiveVersion struct {
		// The unique ID of the version.
		ID uint

		// The game ID this version belongs to.
		GameID uint

		// The semver of the game.
		Version string

		// Value used to help ordering versions.
		VersionOrder uint

		// The changelog of version.
		Changelog string

		// The current state of the version. draft: Not in review, not published. pending: In review. published: Published and available to the public.
		State string // (draft, pending, published)

		// Installation instructions, may contain HTML.
		Installation string

		// Download instructions, usually a url.
		Download string

		// The controls for the game.
		Controls *InteractiveControls

		// Indicates which version of the Interactive Controls this Interactive Version uses.
		ControlVersion string // (1.0, 2.0)
	}

	// A user invoice.
	Invoice struct {
		// The ID of the invoice.
		ID uint

		// ISO 4217 currency code
		Currency string

		// The invoice status
		Status string // (paid, unpaid, errored, cancelled)

		// The owner of the invoice.
		User uint

		// Creation date
		CreatedAt *IsoDate

		// The total invoice amount.
		Amount int

		// List of invoice items
		Items []InvoiceItem
	}

	// A user invoice.
	InvoiceItem struct {
		// The ID of the invoice item.
		ID uint

		// The purchase type
		Type string // (subscription, redeemable)

		// The channel id
		RelID uint

		// A brief summary of what has been purchased.
		Description string

		// The cost of the item.
		Amount int

		// Item quantity.
		Quantity uint

		// The invoice status
		Status string // (paid, unpaid, errored, cancelled)

		// The user ID of the invoice owner.
		User uint

		// The invoice ID the item belongs to.
		Invoice uint
	}

	IsoDate struct {
		time.Time
	}

	// An object representing a language.
	Language struct {
		// A rfc5646 language code.
		ID string

		// The full language name.
		Name string
	}

	// An object describing a count for a language.
	LanguageCount struct {
		// A rfc5646 language code.
		ID string

		// The count.
		Count uint
	}

	LightVideoManifest struct {
		Resolutions []struct {
			*VideoManifestResolution

			// The source url for the stream.
			URL string
		}

		// Time the stream started.
		Since *IsoDate
	}

	Notification struct {
		// The user ID that this chat user belongs to.
		UserID uint

		// The time at which the notification was sent.
		SentAt *IsoDate

		// The event that triggered the notification.
		Trigger string

		// An generally unstructured object containing information about the event. Events of the same type will share the same structure.
		Payload interface{}
	}

	NotificationPreference struct {
		ID string

		// Whether the user allows general marketing/promotional email.
		AlllowsEmailMarketing bool

		// A list of transports follower notifications will be sent over.
		NotifyFollower []string

		// A list of transports subscriber notifications will be sent over.
		NotifySubscriber []string

		// A list of notification transport wentlive notifications are sent on by default.
		LiveOnByDefault []string

		// The time the user last read their notifications.
		LastRead *IsoDate

		// A list of channels and transports registered on which we'll sent wentlive emails.
		LiveNotifications []struct {
			// May be "*", which will toggle on the associated transport for all channels the user follows.
			ChannelID interface{}

			Kind []string
		}

		// Item type: object
		Transports []interface{}
	}

	OAuthAuthorization struct {
		// List of permissions the authorization has.
		Permissions []string

		// The ID of the other the authorization is for.
		UserID uint

		// The OAuth client linked to this authorization.
		Client *OAuthClient
	}

	OAuthClient struct {
		// Internal client id.
		ID uint

		// The unique ID of the client.
		ClientID string // pattern: ^[a-f0-9]{48}$

		// The name of the client.
		Name string

		// The url of the website using this client.
		Website string

		// The url to the logo of the website.
		Logo string

		// List of wildcard hosts allowed to use the client.
		Hosts []string
	}

	OAuthLink struct {
		// The service this link is to, such as twitter or discord.
		Service string

		// The user's ID on the external service.
		ServiceID string

		// The associated Beam user ID.
		UserID int
	}

	PartnershipApp struct {
		Status string // (applied, accepted, denied)

		// The date of the next possible application, only set if status is denied. If the value is null while status is denied, the channel is banned from re-applying.
		Reapplies *IsoDate

		// The reason of the denial, only set if status is denied.
		Reason string
	}

	// A fully populater user with channel, preferences, groups and private details.
	PrivatePopulatedUser struct {
		// The users channel.
		Channel *Channel

		// The global user groups.
		Groups []UserGroup

		// The preferences the user has.
		Preferences *UserPreferences

		// Two factor related data.
		TwoFactor struct {
			// Indicates whether two factor is enabled.
			Enabled bool
		}
	}

	// A fully populater user with channel, preferences, groups and private details.
	PrivateUser struct {
		// The users email address.
		Email string

		// The users password.
		Password string // minLength: 4

		TwoFactor struct {
			// Indicates whether 2fa is enabled.
			Enabled bool

			// Indicates whether recovery codes have been viewed.
			CodesViewed bool
		}
	}

	Rating struct {
		// Rating system and id. E.g. ESRB:M.
		ID string

		// Localized name for the rating.
		Name string

		// Url for the rating logo.
		LogoURL string
	}

	Recording struct {
		// The unique ID of the recording.
		ID uint

		// The channel ID this recording is a video of.
		ChannelID uint

		State string // (PROCESSING, AVAILABLE, DELETED)

		// The number of users who have viewed this recording.
		ViewsTotal uint

		// The date this recording will be deleted at.
		ExpiresAt *IsoDate

		// The video formats available for this recording.
		VODs []VOD

		// Whether the current user has viewed the recording. This will generally only appear when the recording is looked up from an endpoint with user in the query string.
		Viewed bool

		// Name of the Recording, usually the title of the channel at the time of the recording unless the creator changed it.
		Name string

		// Type id of the recording, usually the type id of the channel at the time of the recording, unless the creator changed it.
		TypeID uint

		// Duration of the recording in seconds.
		Duration int

		// Defines if the video was seen by the user. Only present when the user context was given.
		Seen bool
	}

	RecurringPayment struct {
		// The unique ID of the payment.
		ID uint

		// The type of the product the payment is for.
		Type string

		// The ID of the target the payment is for, can be a channe or a user.
		RelID uint

		// The status of the payment.
		Status string

		// Indicates whether the payment has been cancelled. This is not the same as status being set to cancelled, this property indicates that the user cancelled a running payment.
		Cancelled bool

		// The internal name of the payment processsor.
		Gateway string // (braintree, stripe)

		// Counts the amount of times this recurring payment has recurred.
		TimesPaid uint

		// The ID of the user the payment is for.
		User uint
	}

	RecurringPaymentExpanded struct {
		Subscription struct {
			*Subscription

			// The user of the payment.
			User *User

			// The group obtained through this payment.
			Group *UserGroup
		}
	}

	Redeemable struct {
		// The unique ID of the redeemable.
		ID uint

		// The ID of the owning user.
		OwnerID uint

		State string // (unpaid, paid, redeemed)
		Type  string // (pro)

		// The redeem code.
		Code string

		// The ID of the user that used the code.
		RedeemedByID uint

		// The time the item was redeemed.
		RedeemedAt *IsoDate
	}

	RedeemablePartner struct {
		// The unique ID of the channel.
		ID uint

		// The id of the related redeemable.
		RedeemableID uint

		// The id of the related partner.
		PartnerID uint
	}

	Resource struct {
		// The unique ID of the resource.
		ID uint

		// The type of the resource.
		Type string

		// Id linking to the parent object.
		RelID uint

		// The url of the resource.
		URL string

		// The storage type of the resource.
		Store string // (s3)

		// Relative url to the resource.
		RemotePath string

		// Additional resource information.
		Meta interface{}
	}

	Session struct {
		// The when the session expires.
		Expires *UnixTimestampMillis

		// Session meta data.
		Meta struct {
			// The device type.
			Device string

			// The client name. "Other" is web.
			Client string

			// Semantic version of the client.
			CVersion string

			// The operating system name of the client.
			OS string
		}

		// Last if used for that session, can be v4 and v6.
		IP string

		// The session ID as a UUID.
		ID string
	}

	Share struct {
		// The unique ID of the share.
		ID uint

		// The type of the share.
		Type string // (interactiveGameDraft)

		// The
		RelID string

		// Share code.
		Code string

		// The user ID of the target, can be unset if code is set.
		UserID uint
	}

	SocialInfo struct {
		// Twitter profile URL.
		Twitter string

		// Facebook profile URL.
		Facebook string

		// Youtube profile URL.
		YouTube string

		// Player.me profile URL.
		Player string

		// Discord username and tag.
		Discord string

		// A list of social keys which have been verified via linking the Beam account with the account on the corresponding external service.
		Verified []string
	}

	SparkSpendingAnalytic struct {
		Channel uint

		// The number of sparks spent in this event (always a positive number).
		Sparks int

		// The ID of the user who followed, may not be present in aggregations.
		User uint

		Time *IsoDate
	}

	StreamHostsAnalytic struct {
		Channel uint

		// The channel doing the hoster, may be omitted if the channel is the hoster.
		Hoster uint

		// The channel being hosted, may be omitted if the channel is the hostee.
		Hostee uint

		Time *IsoDate
	}

	StreamSessionsAnalytic struct {
		Channel uint

		// Whether this event is the channel going online, or offline.
		Online bool

		// The length of the last stream, only included in offline events.
		Duration uint

		// The ID of the game the stream was playing, only included in offline events.
		TypeID uint

		Time *IsoDate
	}

	SubRevenueAnalytic struct {
		Channel uint

		// The gateway that the subscription was made with.
		Gateway string // (redeemable, stripe, braintree)

		// The amount of revenue gained through this gateway.
		Total int

		// The amount of money prior to transaction fees being applied.
		Gross int

		// The number of subscriptions made through this gateway.
		Count uint
	}

	Subscription struct {
		// The unique ID of the subscription.
		ID uint

		// Id of the resource type this payment is for, only channel ID at the moment.
		ResourceID uint

		// The resource type the resourceId points to, only channel at the moment.
		ResourceType string

		// The status of the subscription.
		Status string // (active, expired, pending)

		// Indicates whether the payment has been cancelled after completion.
		Cancelled bool

		// The time when the subscription expires.
		ExpiresAt *IsoDate
	}

	SubscriptionWithGroup struct {
		// The group ID of the group the subscription grants the user access to.
		ID uint

		// The embedded group assigned to the user.
		Group *UserGroup
	}

	SubscriptionsAnalytic struct {
		Channel uint

		// The total number of subscribers the channel now has.
		Total uint

		// Increment that the total changed as a result of this event.
		Delta int

		// The ID of the user who subscribed, may not be present in aggregations.
		User uint

		Time *IsoDate
	}

	Team struct {
		// The unique ID of the team.
		ID uint

		// The ID of the owner user.
		OwnerID uint

		// The internal name of the team.
		Token string // minLength: 4, maxLength: 20

		// The display name of the team.
		Name string // minLength: 4, maxLength: 36

		// The description of the team.
		Description string

		// The url pointing to a logo image.
		LogoURL string

		// The url pointing to a background image.
		BackgroundURL string

		// Social info of the team.
		Social *SocialInfo
	}

	TeamMembership struct {
		// The unique ID of the team membership.
		ID uint

		// The team ID of this membership.
		TeamID uint

		// The user ID of this membership.
		UserID uint

		// If false, the user has just been invited. If true, the user has joined the team.
		Accepted bool
	}

	TeamMembershipExpanded struct {
		// Indicates that this team is the users primary.
		IsPrimary bool

		// Embedded user object.
		Owner *User

		// The membership object linking team and user.
		TeamMembership *TeamMembership
	}

	// A type that contains information about creation, update and deletion dates.
	TimeStamped struct {
		// The creation date of the channel.
		CreatedAt *IsoDate

		// The update date of the channel.
		UpdatedAt *IsoDate

		// The deletion date of the channel.
		DeletedAt *IsoDate
	}

	TranscodingProfile struct {
		// The unique id of the transcoding profile.
		ID uint

		// The name of the transcoding profile.
		Name string
	}

	TranscodingProfileTranscode struct {
		// The unique id of the transcode.
		ID uint

		// The name of the transcode.
		Name string

		// The display title of the transcode.
		Title string

		// The width (in pixels) of the transcode.
		Width uint

		// The height (in pixels) of the transcode.
		Height uint

		// The bitrate of the transcode.
		Bitrate uint

		// The FPS of the transcode.
		FPS uint

		// If this transcode requires to be partnered.
		RequiresPartner bool
	}

	UUID struct {
		string
	}

	UnixTimestampMillis struct {
		*time.Time
	}

	UpsellInformation struct {
		// Windows Store product id (aka big id).
		ProductID string

		// Localized title of product.
		Title string

		// Box art url.
		BoxArtURL string

		// Localized product description.
		Description string

		// Price and currency.
		Price string

		// List of rating ids for product.
		Ratings []string

		// Link to store.
		Link string
	}

	User struct {
		// The unique ID of the user.
		ID uint

		// The users experience level, related to experience.
		Level uint

		// Social links.
		Social *SocialInfo

		UserName string // minLength: 4, maxLength: 20, pattern: ^[A-Za-z_][\w-]+$
		Email    string // maxLength: 190

		// Indicates whether the user has verified their e-mail.
		Verified bool

		// The users experience points.
		Experience uint

		// The amount of sparks the user has.
		Sparks uint

		// The users profile URL.
		AVATARURL string

		// The users biography, may contain HTML.
		Bio string

		// The ID of users main team.
		PrimaryTeam uint

		// The ID of the transcoding profile currently active.
		TranscodingProfileID uint

		// Indicates whether the user can choose a transcode profile or not.
		HasTranscodes bool
	}

	UserGroup struct {
		// The unique ID of the group.
		ID uint

		// The name of the group.
		Type string
	}

	UserLog struct {
		// The unique ID of the log entry.
		ID uint

		// User ID of the log entry.
		UserID uint

		// A unique identifier for the log entry.
		Event string

		// Unstructured data describing the event.
		EventData interface{}

		// Source data for the event action. This may be one of:
		Source string // (GUEST, SELF, APP, INTERNAL, ADMIN)

		// Unstructured data containing more information about the event's source.
		SourceData interface{}

		// The date the log entry occurred at.
		CreatedAt *IsoDate
	}

	// Object containing user preferences.
	UserPreferences struct {
		// Use html5 to play chat sounds.
		ChatSoundsHTML5 bool `json:"chat:sounds:html5"`

		// Play sounds in chat.
		ChatSoundsPlay bool `json:"chat:sounds:play"`

		// Allow receiving whispers.
		ChatWhispers bool `json:"chat:whispers"`

		// Show timestamps in chat.
		ChatTimestamps bool `json:"chat:timestamps"`

		// Set the chat color to a certain color so it can be easily replaced.
		ChatChromakey bool `json:"chat:chromakey"`

		// Enable chat tagging via @username.
		ChatTagging bool `json:"chat:tagging"`

		// Chat sound volume as unit interval.
		ChatSoundsVolume int `json:"chat:sounds:volume"`

		// Use colors in chat.
		ChatColors bool `json:"chat:colors"`

		// Hide while in chats.
		ChatLurkmode bool `json:"chat:lurkmode"`

		// Notification settings.
		ChannelNotifications struct {
			// List of sources notifications are allowed from.
			IDs []string

			Transports []string // (notify, email)
		} `json:"channel:notifications"`

		// Confirmed mature channels.
		ChannelMatureAllowed bool `json:"channel:mature:allowed"`

		// Force the player to the flash version.
		ChannelPlayerForceflash bool `json:"channel:player:forceflash"`
	}

	UserWithChannel struct {
		Channel *Channel
	}

	UserWithGroups struct {
		Groups []UserGroup
	}

	VOD struct {
		// The unique ID of the VOD.
		ID uint

		// The base URL of a VOD
		BaseURL string

		// The format of the recording.
		Format string // (hls, raw, dash, thumbnail, chat)

		// Format-specific information about the VOD. Is null when type is chat.
		Data struct {
			// Present for hls, raw, dash and thumbnail.
			Width uint

			// Present for hls, raw, dash and thumbnail.
			Height uint

			// Present for hls, raw and dash.
			FPS int

			// Present for hls, raw and dash.
			Bitrate uint
		}

		// Id of the parent recording.
		RecordingID uint
	}

	VideoManifestResolution struct {
		// The name of the resolution.
		Name string

		// The slug of the resolution.
		Slug string

		// The width of the stream, only set if hasVideo is true.
		Width uint

		// The height of the stream, only set if hasVideo is true.
		Height uint

		// Indicates whether the manifest includes video.
		HasVideo bool

		// The total bitrate.
		Bitrate uint
	}

	ViewerAnalytic struct {
		Channel uint

		// Number of anonymous viewers watching the channel at a point in time.
		Anon uint

		// Number of authenticated viewers, including lurkers, watching the channel at a point in time.
		Authed uint

		Time *IsoDate
	}

	ViewerCount struct {
		// The time of the sample.
		Time *IsoDate

		// Amount of anonymous viewers.
		Anon uint

		// Amount registered viewers.
		Authed uint
	}

	ViewerMetricAnalytic struct {
		Channel uint

		// The viewer's two character country code, if it can be determined..
		Country string

		// The viewer's browser, if it can be determined.
		Browser string // (chr, ff, ie, ios, sf, an)

		// The viewer's playform, if it can be determined.
		Platform string // (desktop, tablet, mobile)?

		Time *IsoDate
	}
)

// Unmarshal implements the json.Unmarshaler interface.
// Time is expected in the format yyyy-mm-dd hh:mm:ss.
func (date *IsoDate) Unmarshal(data []byte) (err error) {
	str := string(data[1 : len(data)-1])
	(*date).Time, err = time.Parse("2006-01-02T15:04:05.999Z", str)
	return
}
