package oauth

const (
	// ScopeAchievementsViewSelf allow view your earned achievements.
	ScopeAchievementsViewSelf = "achievement:view:self"

	// ScopeChannelAnalytics allow view analytics for a channel.
	ScopeChannelAnalytics = "channel:analytics"

	// ScopeChannelAnalyticsSelf allow view your channel analytics.
	ScopeChannelAnalyticsSelf = "channel:analytics:self"

	// ScopeChannelCoStreamSelf allow manage your costreaming requests.
	ScopeChannelCoStreamSelf = "channel:costream:self"

	// ScopeChannelDetailsSelf allow view your channel details.
	ScopeChannelDetailsSelf = "channel:details:self"

	// ScopeChannelFollowSelf allow follow and unfollow other channels.
	ScopeChannelFollowSelf = "channel:follow:self"

	// ScopeChannelPartnership allow create and view partnership applications.
	ScopeChannelPartnership = "channel:partnership"

	// ScopeChannelPartnershipSelf allow manage your partnership status.
	ScopeChannelPartnershipSelf = "channel:partnership:self"

	// ScopeChannelStreamKeySelf allow view your channel's stream key.
	ScopeChannelStreamKeySelf = "channel:streamKey:self"

	// ScopeChannelUpdateSelf allow update your channel settings
	ScopeChannelUpdateSelf = "channel:update:self"

	// ScopeChatByPassLinks allow bypass links being disallowed in chat.
	ScopeChatByPassLinks = "chat:bypass_links"

	// ScopeChatByPassSlowChat allow bypass slowchat settings on channels.
	ScopeChatByPassSlowChat = "chat:bypass_slowchat"

	// ScopeChatChangeBan allow manage bans in chats.
	ScopeChatChangeBan = "chat:change_ban"

	// ScopeChatChangeRole allow manage roles in chats.
	ScopeChatChangeRole = "chat:change_role"

	// ScopeChat allow interact with chats on your behalf.
	ScopeChat = "chat:chat"

	// ScopeChatClearMessages allow clear messages in chats where authorized.
	ScopeChatClearMessages = "chat:clear_messages"

	// ScopeChatConnect allow connect to chat.
	ScopeChatConnect = "chat:connect"

	// ScopeChatEditOptions allow edit chat options, including links settings and slowchat.
	ScopeChatEditOptions = "chat:edit_options"

	// ScopeChatGiveawayStart allow start a giveaway in chats where authorized.
	ScopeChatGiveawayStart = "chat:giveaway_start"

	// ScopeChatPollStart allow start a poll in chats where authorized.
	ScopeChatPollStart = "chat:poll_start"

	// ScopeChatPollVote allow vote in chat polls.
	ScopeChatPollVote = "chat:poll_vote"

	// ScopeChatPurge allow clear all messages from a specific user in chat.
	ScopeChatPurge = "chat:purge"

	// ScopeChatRemoveMessage allow remove own and other's messages in chat.
	ScopeChatRemoveMessage = "chat:remove_message"

	// ScopeChatTimeout allow change timeout settings in chats.
	ScopeChatTimeout = "chat:timeout"

	// ScopeChatViewDeleted allow view deleted messages in chat.
	ScopeChatViewDeleted = "chat:view_deleted"

	// ScopeChatWhisper allow gives the ability to whisper in a channel
	ScopeChatWhisper = "chat:whisper"

	// ScopeInteractiveManageSelf allow create, update and delete the interactive games in your account.
	ScopeInteractiveManageSelf = "interactive:manage:self"

	// ScopeInteractiveRobotSelf allow run as an interactive game in your channel.
	ScopeInteractiveRobotSelf = "interactive:robot:self"

	// ScopeInvoiceSelf allow view the users invoices.
	ScopeInvoiceSelf = "invoice:view:self"

	// ScopeLogViewSelf allow view and manage the your security log.
	ScopeLogViewSelf = "log:view:self"

	// ScopeNotificationUpdateSelf allow create and manage your notifications.
	ScopeNotificationUpdateSelf = "notification:update:self"

	// ScopeNotificationViewSelf allow view your notifications.
	ScopeNotificationViewSelf = "notification:view:self"

	// ScopeRecordingManageSelf allow manage the users VODs.
	ScopeRecordingManageSelf = "recording:manage:self"

	// ScopeRedeemableCreateSelf allow create redeemables after performing a purchase.
	ScopeRedeemableCreateSelf = "redeemable:create:self"

	// ScopeRedeemableRedeemSelf allow use users redeemable.
	ScopeRedeemableRedeemSelf = "redeemable:redeem:self"

	// ScopeRedeemableViewSelf allow view users redeemables.
	ScopeRedeemableViewSelf = "redeemable:view:self"

	// ScopeResourceFindSelf allow view emoticons and other graphical resources you have access to.
	ScopeResourceFindSelf = "resource:find:self"

	// ScopeSubscriptionCancelSelf allow cancel your subscriptions.
	ScopeSubscriptionCancelSelf = "subscription:cancel:self"

	// ScopeSubscriptionCreateSelf allow create new subscriptions.
	ScopeSubscriptionCreateSelf = "subscription:create:self"

	// ScopeSubscriptionReNewSelf allow renew your existing subscriptions.
	ScopeSubscriptionReNewSelf = "subscription:renew:self"

	// ScopeSubscriptionViewSelf allow view who you're subscribed to.
	ScopeSubscriptionViewSelf = "subscription:view:self"

	// ScopeTeamAdminister allow administrate teams the user has rights in.
	ScopeTeamAdminister = "team:administer"

	// ScopeTeamManageSelf allow create, join, leave teams and set the users primary team.
	ScopeTeamManageSelf = "team:manage:self"

	// ScopeTransactionCancelSelf allow cancel pending transactions.
	ScopeTransactionCancelSelf = "transaction:cancel:self"

	// ScopeTransactionViewSelf allow view your pending transactions.
	ScopeTransactionViewSelf = "transaction:view:self"

	// ScopeUserAnalyticsSelf allow view your user analytics
	ScopeUserAnalyticsSelf = "user:analytics:self"

	// ScopeUserDetailsSelf allow view your email address and other private details.
	ScopeUserDetailsSelf = "user:details:self"

	// ScopeUserGetDiscordInviteSelf allow view users discord invites.
	ScopeUserGetDiscordInviteSelf = "user:getDiscordInvite:self"

	// ScopeUserLogSelf allow view your user security log.
	ScopeUserLogSelf = "user:log:self"

	// ScopeUserNotificationSelf allow view and manage your notifications.
	ScopeUserNotificationSelf = "user:notification:self"

	// ScopeUserSeenSelf allow mark a VOD as seen for the user.
	ScopeUserSeenSelf = "user:seen:self"

	// ScopeUserUpdateSelf allow update your account, including your email but not your password.
	ScopeUserUpdateSelf = "user:update:self"

	// ScopeUserUpdatePasswordSelf allow update your password.
	ScopeUserUpdatePasswordSelf = "user:updatePassword:self"
)
