package tgbot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func PrintMessage(message *tgbotapi.Message) {
	if message != nil {
		fmt.Println("===============================")
		fmt.Printf("MessageID: %d\n", message.MessageID)
		if message.From != nil {
			fmt.Printf("From: %+v\n", *message.From)
		}
		if message.SenderChat != nil {
			fmt.Printf("SenderChat: %+v\n", *message.SenderChat)
		}
		fmt.Printf("Date: %d\n", message.Date)
		if message.Chat != nil {
			fmt.Printf("Chat: %+v\n", *message.Chat)
		}
		if message.ForwardFrom != nil {
			fmt.Printf("ForwardFrom: %+v\n", *message.ForwardFrom)
		}
		if message.ForwardFromChat != nil {
			fmt.Printf("ForwardFromChat: %+v\n", *message.ForwardFromChat)
		}
		if message.ForwardFromMessageID != 0 {
			fmt.Printf("ForwardFromMessageID: %d\n", message.ForwardFromMessageID)
		}
		if message.ForwardSignature != "" {
			fmt.Printf("ForwardSignature: %s\n", message.ForwardSignature)
		}
		if message.ForwardSenderName != "" {
			fmt.Printf("ForwardSenderName: %s\n", message.ForwardSenderName)
		}
		if message.ForwardDate != 0 {
			fmt.Printf("ForwardDate: %d\n", message.ForwardDate)
		}
		if message.IsAutomaticForward {
			fmt.Println("IsAutomaticForward: true")
		}
		if message.ReplyToMessage != nil {
			fmt.Printf("ReplyToMessage: %+v\n", *message.ReplyToMessage)
		}
		if message.ViaBot != nil {
			fmt.Printf("ViaBot: %+v\n", *message.ViaBot)
		}
		if message.EditDate != 0 {
			fmt.Printf("EditDate: %d\n", message.EditDate)
		}
		if message.HasProtectedContent {
			fmt.Println("HasProtectedContent: true")
		}
		if message.MediaGroupID != "" {
			fmt.Printf("MediaGroupID: %s\n", message.MediaGroupID)
		}
		if message.AuthorSignature != "" {
			fmt.Printf("AuthorSignature: %s\n", message.AuthorSignature)
		}
		if message.Text != "" {
			fmt.Printf("Text: %s\n", message.Text)
		}
		if len(message.Entities) > 0 {
			fmt.Printf("Entities: %+v\n", message.Entities)
		}
		if message.Animation != nil {
			fmt.Printf("Animation: %+v\n", *message.Animation)
		}
		if message.Audio != nil {
			fmt.Printf("Audio: %+v\n", *message.Audio)
		}
		if message.Document != nil {
			fmt.Printf("Document: %+v\n", *message.Document)
		}
		if len(message.Photo) > 0 {
			fmt.Printf("Photo: %+v\n", message.Photo)
		}
		if message.Sticker != nil {
			fmt.Printf("Sticker: %+v\n", *message.Sticker)
		}
		if message.Video != nil {
			fmt.Printf("Video: %+v\n", *message.Video)
		}
		if message.VideoNote != nil {
			fmt.Printf("VideoNote: %+v\n", *message.VideoNote)
		}
		if message.Voice != nil {
			fmt.Printf("Voice: %+v\n", *message.Voice)
		}
		if message.Caption != "" {
			fmt.Printf("Caption: %s\n", message.Caption)
		}
		if len(message.CaptionEntities) > 0 {
			fmt.Printf("CaptionEntities: %+v\n", message.CaptionEntities)
		}
		if message.Contact != nil {
			fmt.Printf("Contact: %+v\n", *message.Contact)
		}
		if message.Dice != nil {
			fmt.Printf("Dice: %+v\n", *message.Dice)
		}
		if message.Game != nil {
			fmt.Printf("Game: %+v\n", *message.Game)
		}
		if message.Poll != nil {
			fmt.Printf("Poll: %+v\n", *message.Poll)
		}
		if message.Venue != nil {
			fmt.Printf("Venue: %+v\n", *message.Venue)
		}
		if message.Location != nil {
			fmt.Printf("Location: %+v\n", *message.Location)
		}
		if len(message.NewChatMembers) > 0 {
			fmt.Printf("NewChatMembers: %+v\n", message.NewChatMembers)
		}
		if message.LeftChatMember != nil {
			fmt.Printf("LeftChatMember: %+v\n", *message.LeftChatMember)
		}
		if message.NewChatTitle != "" {
			fmt.Printf("NewChatTitle: %s\n", message.NewChatTitle)
		}
		if len(message.NewChatPhoto) > 0 {
			fmt.Printf("NewChatPhoto: %+v\n", message.NewChatPhoto)
		}
		if message.DeleteChatPhoto {
			fmt.Println("DeleteChatPhoto: true")
		}
		if message.GroupChatCreated {
			fmt.Println("GroupChatCreated: true")
		}
		if message.SuperGroupChatCreated {
			fmt.Println("SuperGroupChatCreated: true")
		}
		if message.ChannelChatCreated {
			fmt.Println("ChannelChatCreated: true")
		}
		if message.MessageAutoDeleteTimerChanged != nil {
			fmt.Printf("MessageAutoDeleteTimerChanged: %+v\n", *message.MessageAutoDeleteTimerChanged)
		}
		if message.MigrateToChatID != 0 {
			fmt.Printf("MigrateToChatID: %d\n", message.MigrateToChatID)
		}
		if message.MigrateFromChatID != 0 {
			fmt.Printf("MigrateFromChatID: %d\n", message.MigrateFromChatID)
		}
		if message.PinnedMessage != nil {
			fmt.Printf("PinnedMessage: %+v\n", *message.PinnedMessage)
		}
		if message.Invoice != nil {
			fmt.Printf("Invoice: %+v\n", *message.Invoice)
		}
		if message.SuccessfulPayment != nil {
			fmt.Printf("SuccessfulPayment: %+v\n", *message.SuccessfulPayment)
		}
		if message.ConnectedWebsite != "" {
			fmt.Printf("ConnectedWebsite: %s\n", message.ConnectedWebsite)
		}
		if message.PassportData != nil {
			fmt.Printf("PassportData: %+v\n", *message.PassportData)
		}
		if message.ProximityAlertTriggered != nil {
			fmt.Printf("ProximityAlertTriggered: %+v\n", *message.ProximityAlertTriggered)
		}
		if message.VoiceChatScheduled != nil {
			fmt.Printf("VoiceChatScheduled: %+v\n", *message.VoiceChatScheduled)
		}
		if message.VoiceChatStarted != nil {
			fmt.Printf("VoiceChatStarted: %+v\n", *message.VoiceChatStarted)
		}
		if message.VoiceChatEnded != nil {
			fmt.Printf("VoiceChatEnded: %+v\n", *message.VoiceChatEnded)
		}
		if message.VoiceChatParticipantsInvited != nil {
			fmt.Printf("VoiceChatParticipantsInvited: %+v\n", *message.VoiceChatParticipantsInvited)
		}
		if message.ReplyMarkup != nil {
			fmt.Printf("ReplyMarkup: %+v\n", *message.ReplyMarkup)
		}
		fmt.Println("===============================")
	}
}
